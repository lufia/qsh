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

If statement:
	op:word("ls")
	op:simple
	op:if
	op:int(&END)
	op:word("pwd")
	op:simple
	op:wasTrue
	op:END
*/

type Code struct {
	steps []func(cmd *Cmd)
}

// Pos returns next position.
func (c *Code) Pos() int {
	return len(c.steps)
}

func (c *Code) emit(f func(cmd *Cmd)) {
	c.steps = append(c.steps, f)
}

func (c *Code) alloc() *addr {
	pos := c.Pos()
	c.steps = append(c.steps, c.nop)
	return &addr{
		code: c,
		slot: pos,
	}
}

func (*Code) nop(*Cmd) {
}

type addr struct {
	code *Code
	slot int
}

func (a *addr) Set(f func(cmd *Cmd)) {
	a.code.steps[a.slot] = f
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
	case ast.BLOCK:
		walk(c, p.Left)
	case ast.ASYNC:
		// TODO: func do not fork.
		walk(c, p.Left)
	case ast.VAR:
		walk(c, p.Left)
		c.emit(Var)
	case ast.ASSIGN:
		walk(c, p.Right)
		walk(c, p.Left)
		c.emit(Assign)
	case ast.IF:
		walk(c, p.Left)
		c.emit(If)
		op := c.alloc()
		walk(c, p.Right)
		label := Label(c.Pos())
		op.Set(label.Jump)
	}
	return nil
}
