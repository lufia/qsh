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
	mtab = make(map[string]func([]string) ([]string, error))
	runq *Cmd
)

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

func (cmd *Cmd) popStack() *Stack {
	s := cmd.currentStack()
	cmd.stack = cmd.stack[0 : len(cmd.stack)-1]
	return s
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
	s := cmd.popStack()
	if len(s.words) != 1 {
		Error(errors.New("< requires singleton"))
		return
	}

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
	s := cmd.popStack()
	if len(s.words) != 1 {
		Error(errors.New("> requires singleton"))
		return
	}

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
	s := cmd.popStack()
	if len(s.words) != 1 {
		Error(errors.New(">> requires singleton"))
		return
	}

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
	s := cmd.popStack()
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
	paths, _ := LookupVar("PATH")
	for _, dir := range paths {
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
	s := cmd.popStack()
	if len(s.words) != 1 {
		Error(errors.New("variable name is not singleton"))
		return
	}
	v, _ := LookupVar(s.words[0])
	s1 := cmd.currentStack()
	s1.words = append(s1.words, v...)
}

func Assign(cmd *Cmd) {
	s := cmd.popStack()
	if len(s.words) != 1 {
		Error(errors.New("variable name is not singleton"))
		return
	}
	name := s.words[0]

	s = cmd.popStack()
	if len(s.words) == 0 {
		UnsetVar(name)
	} else {
		UpdateVar(name, s.words)
	}
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
	UpdateVar(name, []string{p.words[0]})
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

func Load(cmd *Cmd) {
	s := cmd.currentStack()
	if len(s.words) != 1 {
		Error(errors.New("load name is not singleton"))
		return
	}
	if err := load(s.words[0]); err != nil {
		Error(err)
		return
	}
}

func Module(cmd *Cmd) {
	s := cmd.popStack()
	if len(s.words) == 0 {
		Error(errors.New("module requires one or more args"))
		return
	}
	f, ok := mtab[s.words[0]]
	if !ok {
		Error(errors.New(s.words[0] + ": module not found"))
		return
	}
	v, err := f(s.words[1:])
	if err != nil {
		Error(err)
		return
	}

	s1 := cmd.currentStack()
	s1.words = append(s1.words, v...)
}
