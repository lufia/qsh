package ast

//go:generate stringer -type=LexType
type LexType int

//go:generate stringer -type=Direction
type Direction int

const (
	WORD LexType = iota
	SIMPLE
	LIST
	BLOCK
	ASYNC
	VAR
	TUPLE
	ASSIGN
	IF
	FOR
	REDIR
)

const (
	READ Direction = iota
	WRITE
	APPEND
	HERE
)

type Node struct {
	Type   LexType
	Str    string
	Quoted bool
	Left   *Node
	Right  *Node
	Dir    Direction
}

func New(t LexType, p1, p2 *Node) *Node {
	return &Node{
		Type:  t,
		Left:  p1,
		Right: p2,
	}
}

func Token(s string) *Node {
	return &Node{
		Type: WORD,
		Str:  s,
	}
}

func Redir(dir Direction) *Node {
	return &Node{
		Type: REDIR,
		Dir:  dir,
	}
}

func Simple(p *Node) *Node {
	n := New(SIMPLE, p, nil)
	return n
}

func Block(p *Node) *Node {
	return New(BLOCK, p, nil)
}

func Async(p *Node) *Node {
	return New(ASYNC, p, nil)
}

func Var(p *Node) *Node {
	return New(VAR, p, nil)
}

func Tuple(p *Node) *Node {
	return New(TUPLE, p, nil)
}

func Assign(p1, p2 *Node) *Node {
	return New(ASSIGN, p1, p2)
}

func Redirect(p1, p2 *Node) *Node {
	if p1.Type != REDIR {
		panic("first argument of redirect must be REDIR type")
	}
	p1.Left = p2
	return p1
}
