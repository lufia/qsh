package ast

import (
	"fmt"
	"io"
	"os"
)

func Dump(p *Node) {
	dump1(p, os.Stderr, 0)
}

func dump1(p *Node, w io.Writer, indent int) {
	if p == nil {
		return
	}
	switch p.Type {
	case WORD:
		fmt.Fprintf(w, "%*s%s:%q\n", indent, "", p.Type, p.Str)
	case SIMPLE:
		fmt.Fprintf(w, "%*s%s:\n", indent, "", p.Type)
		fmt.Fprintf(w, "%*sLeft:\n", indent+1, "")
		dump1(p.Left, w, indent+2)
	case LIST:
		fmt.Fprintf(w, "%*s%s:\n", indent, "", p.Type)
		fmt.Fprintf(w, "%*sLeft:\n", indent+1, "")
		dump1(p.Left, w, indent+2)
		fmt.Fprintf(w, "%*sRight:\n", indent+1, "")
		dump1(p.Right, w, indent+2)
	case BLOCK:
		fmt.Fprintf(w, "%*s%s:\n", indent, "", p.Type)
		dump1(p.Left, w, indent+2)
	case ASYNC:
		fmt.Fprintf(w, "%*s%s:\n", indent, "", p.Type)
		dump1(p.Left, w, indent+2)
	case VAR:
		fmt.Fprintf(w, "%*s%s:\n", indent, "", p.Type)
		dump1(p.Left, w, indent+2)
	case TUPLE:
		fmt.Fprintf(w, "%*s%s:\n", indent, "", p.Type)
		dump1(p.Left, w, indent+2)
	case ASSIGN:
		fmt.Fprintf(w, "%*s%s:\n", indent, "", p.Type)
		fmt.Fprintf(w, "%*sLeft:\n", indent+1, "")
		dump1(p.Left, w, indent+2)
		fmt.Fprintf(w, "%*sRight:\n", indent+1, "")
		dump1(p.Right, w, indent+2)
	case IF:
		fmt.Fprintf(w, "%*s%s:\n", indent, "", p.Type)
		fmt.Fprintf(w, "%*sLeft:\n", indent+1, "")
		dump1(p.Left, w, indent+2)
		fmt.Fprintf(w, "%*sRight:\n", indent+1, "")
		dump1(p.Right, w, indent+2)
	case FOR:
		fmt.Fprintf(w, "%*s%s:\n", indent, "", p.Type)
		fmt.Fprintf(w, "%*sLeft:\n", indent+1, "")
		dump1(p.Left, w, indent+2)
		fmt.Fprintf(w, "%*sRight:\n", indent+1, "")
		dump1(p.Right, w, indent+2)
	}
}
