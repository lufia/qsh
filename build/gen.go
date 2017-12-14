package build

import (
	"github.com/lufia/qsh/ast"
)

/*
Simple Code:
	op:mark
	op:word("arg2")
	op:word("arg1")
	op:simple

Redirect:
	op:mark
	op:word("ls")
	op:mark
	op:word("out")
	op:stdout
	op:simple

Variable:
	op:mark
	op:word("name")
	op:var

Assign(single):
	op:mark
	op:word("value")
	op:mark
	op:word("name")
	op:assign

Assign(tuple):
	op:mark
	op:word("arg1")
	op:word("arg2")
	op:mark
	op:word("name")
	op:assign

If statement:
	op:mark
	op:word("ls")
	op:simple
	op:if
	op:goto(END)
	op:mark
	op:word("pwd")
	op:simple
	op:wasTrue
END:

For statement:
	op:mark
	op:word("a")
	op:word("b")
	op:mark
	op:word("i")
	op:for
	op:goto(END)
	op:word("ls")
	op:simple
	op:jump(&for)
END:

And operator:
	op:mark
	op:word("a")
	op:simple
	op:true
	op:goto(END)
	op:mark
	op:word("b")
	op:simple
END:

Or operator:
	op:mark
	op:word("a")
	op:simple
	op:false
	op:goto(END)
	op:mark
	op:word("b")
	op:simple
END:

Pipe operator:
	op:pipe
	op:int(fd0)
	op:int(fd1)
	op:goto(EXIT)
	op:goto(END)

	op:mark
	op:word("ls")
	op:simple
	op:exit

EXIT:
	op:mark
	op:work("wc")
	op:simple
	op:return

END:
	op:pipewait
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
	case ast.REDIR:
		c.emit(Mark)
		walk(c, p.Left)
		switch p.Dir {
		case ast.READ:
			c.emit(SetStdin)
		case ast.WRITE:
			c.emit(SetStdout)
		case ast.APPEND:
			c.emit(SetStdoutAppend)
		}
		walk(c, p.Right)
		c.emit(RevertRedir)
	case ast.SIMPLE:
		c.emit(Mark)
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
		c.emit(Mark)
		walk(c, p.Left)
		c.emit(Var)
	case ast.TUPLE:
		walk(c, p.Left)
	case ast.ASSIGN:
		c.emit(Mark)
		walk(c, p.Right)
		c.emit(Mark)
		walk(c, p.Left)
		c.emit(Assign)
	case ast.IF:
		walk(c, p.Left)
		c.emit(If)
		op := c.alloc()
		walk(c, p.Right)
		g := Goto(c.Pos())
		op.Set(g.Jump)
	case ast.FOR:
		c.emit(Mark)
		walk(c, p.Left.Right)
		c.emit(Mark)
		walk(c, p.Left.Left)
		loop := c.Pos()
		c.emit(For)
		op := c.alloc()
		walk(c, p.Right)
		c.emit(Goto(loop).Jump)
		g := Goto(c.Pos())
		op.Set(g.Jump)
	case ast.ANDAND:
		walk(c, p.Left)
		c.emit(ContinueIf)
		op := c.alloc()
		walk(c, p.Right)
		g := Goto(c.Pos())
		op.Set(g.Jump)
	case ast.OROR:
		walk(c, p.Left)
		c.emit(ContinueUnless)
		op := c.alloc()
		walk(c, p.Right)
		g := Goto(c.Pos())
		op.Set(g.Jump)
	case ast.PIPE:
		c.emit(Pipe)
		parent := c.alloc()
		end := c.alloc()
		walk(c, p.Left)
		c.emit(Exit)
		parent.Set(Goto(c.Pos()).Jump)
		walk(c, p.Right)
		c.emit(Return)
		end.Set(Goto(c.Pos()).Jump)
		c.emit(Wait)
	case ast.LOAD:
		c.emit(Mark)
		walk(c, p.Left)
		c.emit(Load)
	}
	return nil
}
