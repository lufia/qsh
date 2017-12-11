package cgen

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	vtab = make(map[string][]string)
	runq *Cmd
)

type Cmd struct {
	code  *Code
	pc    int
	words []string
	ret   *Cmd
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

type Label int

func (t Label) Jump(cmd *Cmd) {
	cmd.pc = int(t)
}

func Return() {
	runq = runq.ret
}

func Error(err error) {
	log.Println(err)
}

type String string

func (s String) Push(cmd *Cmd) {
	cmd.words = append(cmd.words, string(s))
}

func Simple(cmd *Cmd) {
	p := cmd.words[0]
	if !filepath.IsAbs(p) {
		var err error
		p, err = resolvePath(p)
		if err != nil {
			Error(err)
			return
		}
	}
	c := exec.Command(p, cmd.words[1:]...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		Error(err)
		return
	}
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
	return "", errors.New("command not found")
}

func Var(cmd *Cmd) {
	n := len(cmd.words)
	name := cmd.words[n-1]
	v := vtab[name]
	if len(v) != 1 {
		Error(errors.New("variable name is not singleton"))
		return
	}
	cmd.words[n-1] = v[0]
}

func Assign(cmd *Cmd) {
	n := len(cmd.words)
	name := cmd.words[n-1]
	value := cmd.words[n-2]
	vtab[name] = []string{value}
	cmd.words = cmd.words[0 : n-2]
}

func If(cmd *Cmd) {
}
