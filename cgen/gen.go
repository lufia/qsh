package cgen

import (
	"fmt"

	"github.com/lufia/qsh/ast"
)

type Cmd struct {
	pc    int
	words []string
}

func Simple(cmd *Cmd) {
	fmt.Println(cmd.words)
}

type String string

func (s String) Push(cmd *Cmd) {
	cmd.words = append(cmd.words, string(s))
}

type Code struct {
	steps []func(cmd *Cmd)
}

func (c *Code) emit(f func(cmd *Cmd)) {
	c.steps = append(c.steps, f)
}

/*
Simple Code:

op:mark
op:word("arg2")
op:word("arg1")
op:simple
*/

func Compile(p *ast.Node) error {
	var c Code
	walk(&c, p)

	var cmd Cmd
	for cmd.pc < len(c.steps) {
		c.steps[cmd.pc](&cmd)
		cmd.pc++
	}
	return nil
}

func walk(c *Code, p *ast.Node) error {
	if p == nil {
		return nil
	}
	switch p.Type {
	case ast.SIMPLE:
		walk(c, p.Left)
		c.emit(Simple)
	case ast.LIST:
		walk(c, p.Left)
		walk(c, p.Right)
	case ast.WORD:
		s := String(p.Str)
		c.emit(s.Push)
	}
	return nil
}
