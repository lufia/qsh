package ast

//go:generate stringer -type=LexType
type LexType int

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
)

type Node struct {
	Type   LexType
	Str    string
	Quoted bool
	Left   *Node
	Right  *Node
}

func (n *Node) String() string {
	switch n.Type {
	case WORD:
		return n.Str
	default:
		return "Node"
	}
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
