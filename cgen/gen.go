package cgen

import (
	"github.com/lufia/qsh/ast"
)

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

type Code struct {
	steps []func(cmd *Cmd)
}

func (c *Code) emit(f func(cmd *Cmd)) {
	c.steps = append(c.steps, f)
}

func Compile(p *ast.Node) error {
	var c Code
	walk(&c, p)

	Start(&c)
	return nil
}

func walk(c *Code, p *ast.Node) error {
	if p == nil {
		return nil
	}
	switch p.Type {
	case ast.WORD:
		s := String(p.Str)
		c.emit(s.Push)
	case ast.SIMPLE:
		walk(c, p.Left)
		c.emit(Simple)
	case ast.LIST:
		walk(c, p.Left)
		walk(c, p.Right)
	case ast.VAR:
		walk(c, p.Left)
		c.emit(Var)
	case ast.ASSIGN:
		walk(c, p.Right)
		walk(c, p.Left)
		c.emit(Assign)
	}
	return nil
}
