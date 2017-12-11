package cgen

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

type Cmd struct {
	pc    int
	words []string
}

var (
	vtab = make(map[string][]string)
)

type String string

func (s String) Push(cmd *Cmd) {
	cmd.words = append(cmd.words, string(s))
	cmd.pc++
}

func Simple(cmd *Cmd) {
	cmd.pc++
	p := cmd.words[0]
	if !filepath.IsAbs(p) {
		var err error
		p, err = resolvePath(p)
		if err != nil {
			return // TODO: catch an error
		}
	}
	c := exec.Command(p, cmd.words[1:]...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		// TODO
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
		// TODO: variable name is not singleton
		return
	}
	cmd.words[n-1] = v[0]
	cmd.pc++
}

func Assign(cmd *Cmd) {
	n := len(cmd.words)
	name := cmd.words[n-1]
	value := cmd.words[n-2]
	vtab[name] = []string{value}
	cmd.words = cmd.words[0 : n-2]
	cmd.pc++
}
