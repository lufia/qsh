package cgen

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

type Stack struct {
	words []string
}

type Cmd struct {
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
		code: code,
		ret:  runq,
	}
	for runq != nil && runq.pc < len(runq.code.steps) {
		runq.pc++
		code.steps[runq.pc-1](runq)
	}
	Return()
}

type Goto int

func (g Goto) Jump(cmd *Cmd) {
	cmd.pc = int(g)
}

func Return() {
	runq = runq.ret
}

func Error(err error) {
	log.Println(err)
}

type String string

func (s String) Push(cmd *Cmd) {
	p := cmd.currentStack()
	p.words = append(p.words, string(s))
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
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
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
