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
	cmd.pc++
}

type String string

func (s String) Push(cmd *Cmd) {
	cmd.words = append(cmd.words, string(s))
	cmd.pc++
}

var (
	vtab = make(map[string]string)
)

func Var(cmd *Cmd) {
	n := len(cmd.words)
	name := cmd.words[n-1]
	cmd.words[n-1] = vtab[name]
	cmd.pc++
}

func Assign(cmd *Cmd) {
	n := len(cmd.words)
	name := cmd.words[n-1]
	value := cmd.words[n-2]
	vtab[name] = value
	cmd.words = cmd.words[0 : n-2]
	cmd.pc++
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

Variable:

op:mark
op:word("name")
op:var

Assign:

op:mark
op:word("value")
op:mark
op:word("name")
op:assign
*/

func Compile(p *ast.Node) error {
	var c Code
	walk(&c, p)

	var cmd Cmd
	for cmd.pc < len(c.steps) {
		c.steps[cmd.pc](&cmd)
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
	case ast.VAR:
		walk(c, p.Left)
		c.emit(Var)
	case ast.ASSIGN:
		walk(c, p.Right)
		walk(c, p.Left)
		c.emit(Assign)
	case ast.LIST:
		walk(c, p.Left)
		walk(c, p.Right)
	case ast.WORD:
		s := String(p.Str)
		c.emit(s.Push)
	}
	return nil
}
