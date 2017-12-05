package ast

//go:generate stringer -type=LexType
type LexType int

const (
	WORD LexType = iota
	SIMPLE
	LIST
)

type Node struct {
	Type  LexType
	Str   string
	Left  *Node
	Right *Node
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

func Simple(p *Node) *Node {
	n := New(SIMPLE, p, nil)
	return n
}
