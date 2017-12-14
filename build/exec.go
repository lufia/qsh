package build

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var (
	vtab = make(map[string][]string)
	runq *Cmd
)

func lastStatus() string {
	v, ok := vtab["status"]
	if !ok || len(v) != 1 {
		return ""
	}
	return v[0]
}

func isSuccess() bool {
	return lastStatus() == ""
}

func updateStatus(err error) {
	if err == nil {
		delete(vtab, "status")
	} else {
		vtab["status"] = []string{err.Error()}
	}
}

type file struct {
	io.ReadWriteCloser
	opened bool
}

func openFile(name string, flag int, perm os.FileMode) (*file, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return &file{ReadWriteCloser: f, opened: true}, nil
}

func (f *file) Close() error {
	if f.opened {
		return f.ReadWriteCloser.Close()
	}
	f.opened = false
	return nil
}

type Redir struct {
	stdin  io.ReadCloser
	stdout io.WriteCloser
	stderr io.WriteCloser
	next   *Redir
}

func NewRedir() *Redir {
	return &Redir{
		stdin:  os.Stdin,
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
}

func (r *Redir) Stdout() io.Writer {
	for p := r; p != nil; p = p.next {
		if p.stdout != nil {
			return p.stdout
		}
	}
	panic("no stdout")
}

type Stack struct {
	words []string
}

type Cmd struct {
	redir *Redir
	code  *Code
	pc    int
	stack []*Stack
	ret   *Cmd
}

func Mark(cmd *Cmd) {
	var s Stack
	cmd.stack = append(cmd.stack, &s)
}

func (cmd *Cmd) currentStack() *Stack {
	return cmd.stack[len(cmd.stack)-1]
}

func (cmd *Cmd) popStack() {
	cmd.stack = cmd.stack[0 : len(cmd.stack)-1]
}

func Start(code *Code) {
	runq = &Cmd{
		redir: NewRedir(),
		code:  code,
		ret:   runq,
	}
	start(runq)
	runq = runq.ret
}

func start(cmd *Cmd) {
	for cmd != nil && cmd.pc < len(cmd.code.steps) {
		cmd.pc++
		cmd.code.steps[cmd.pc-1](cmd)
	}
}

type Goto int

func (g Goto) Jump(cmd *Cmd) {
	cmd.pc = int(g)
}

func Return(cmd *Cmd) {
	RevertRedir(cmd)
}

func Error(err error) {
	log.Println(err)
}

type String string

func (s String) Push(cmd *Cmd) {
	p := cmd.currentStack()
	p.words = append(p.words, string(s))
}

func SetStdin(cmd *Cmd) {
	s := cmd.currentStack()
	if len(s.words) != 1 {
		Error(errors.New("< requires singleton"))
		return
	}
	cmd.popStack()

	f, err := openFile(s.words[0], os.O_RDONLY, 0)
	if err != nil {
		Error(err)
		return
	}
	redir := &Redir{
		stdin: f,
		next:  cmd.redir,
	}
	cmd.redir = redir
}

func SetStdout(cmd *Cmd) {
	s := cmd.currentStack()
	if len(s.words) != 1 {
		Error(errors.New("> requires singleton"))
		return
	}
	cmd.popStack()

	f, err := openFile(s.words[0], os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		Error(err)
		return
	}
	redir := &Redir{
		stdout: f,
		next:   cmd.redir,
	}
	cmd.redir = redir
}

func SetStdoutAppend(cmd *Cmd) {
	s := cmd.currentStack()
	if len(s.words) != 1 {
		Error(errors.New(">> requires singleton"))
		return
	}
	cmd.popStack()

	f, err := openFile(s.words[0], os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		Error(err)
		return
	}
	redir := &Redir{
		stdout: f,
		next:   cmd.redir,
	}
	cmd.redir = redir
}

func RevertRedir(cmd *Cmd) {
	r := cmd.redir
	cmd.redir = r.next
	if r.stdin != nil {
		r.stdin.Close()
	}
	if r.stdout != nil {
		r.stdout.Close()
	}
	if r.stderr != nil {
		r.stderr.Close()
	}
}

func Simple(cmd *Cmd) {
	defer cmd.popStack()

	s := cmd.currentStack()
	p := s.words[0]
	if !filepath.IsAbs(p) {
		var err error
		p, err = resolvePath(p)
		if err != nil {
			Error(err)
			return
		}
	}
	c := exec.Command(p, s.words[1:]...)
	c.Stdin = cmd.redir.stdin
	c.Stdout = cmd.redir.Stdout()
	c.Stderr = cmd.redir.stderr
	if err := c.Run(); err != nil {
		updateStatus(err)
		return
	}
	updateStatus(nil)
	return
}

func resolvePath(p string) (string, error) {
	for _, dir := range vtab["PATH"] {
		f := filepath.Join(dir, p)
		info, err := os.Stat(f)
		if err != nil {
			continue
		}
		m := info.Mode()
		if m.IsRegular() && (m.Perm()&0111) != 0 {
			return f, nil
		}
	}
	return "", fmt.Errorf("%s: command not found", p)
}

func Var(cmd *Cmd) {
	s := cmd.currentStack()
	if len(s.words) != 1 {
		Error(errors.New("variable name is not singleton"))
		return
	}
	v := vtab[s.words[0]]
	cmd.popStack()
	s1 := cmd.currentStack()
	s1.words = append(s1.words, v...)
}

func Assign(cmd *Cmd) {
	s := cmd.currentStack()
	if len(s.words) != 1 {
		Error(errors.New("variable name is not singleton"))
		return
	}
	cmd.popStack()
	name := s.words[0]

	s = cmd.currentStack()
	if len(s.words) == 0 {
		delete(vtab, name)
	} else {
		vtab[name] = s.words
	}
	cmd.popStack()
}

func If(cmd *Cmd) {
	if isSuccess() {
		cmd.pc++
	}
}

func For(cmd *Cmd) {
	p := cmd.stack[len(cmd.stack)-2]
	if len(p.words) == 0 {
		cmd.popStack()
		return
	}

	s := cmd.currentStack()
	if len(s.words) != 1 {
		Error(errors.New("variable name is not singleton"))
		return
	}

	name := s.words[0]
	vtab[name] = []string{p.words[0]}
	p.words = p.words[1:]
	cmd.pc++ // skip goto op
}

func ContinueIf(cmd *Cmd) {
	if isSuccess() {
		cmd.pc++
	}
}

func ContinueUnless(cmd *Cmd) {
	if !isSuccess() {
		cmd.pc++
	}
}

func Pipe(cmd *Cmd) {
	pr, pw := io.Pipe()
	redir := Redir{
		stdout: pw,
		next:   cmd.redir,
	}
	cmd1 := Cmd{
		redir: &redir,
		code:  cmd.code,
		pc:    cmd.pc + 2,
		ret:   cmd,
	}
	go func() {
		start(&cmd1)
		panic("cannot reach here")
	}()

	cmd.redir = &Redir{
		stdin: pr,
		next:  cmd.redir,
	}
}

func Wait(cmd *Cmd) {
}

func Exit(cmd *Cmd) {
	RevertRedir(cmd)
	runtime.Goexit()
}
