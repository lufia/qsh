package main

import (
	"strings"
	"testing"

	"github.com/lufia/qsh/ast"
)

type Token struct {
	Type int
	tree *ast.Node
}

func TestLexWords(t *testing.T) {
	tab := []struct {
		Source string
		Wants  []*Token
	}{
		{
			Source: "a bc d\n",
			Wants: []*Token{
				&Token{Type: WORD, tree: &ast.Node{Str: "a"}},
				&Token{Type: WORD, tree: &ast.Node{Str: "bc"}},
				&Token{Type: WORD, tree: &ast.Node{Str: "d"}},
				&Token{Type: '\n'},
			},
		},
		{
			Source: "'a b'\n",
			Wants: []*Token{
				&Token{Type: WORD, tree: &ast.Node{Str: "a b"}},
				&Token{Type: '\n'},
			},
		},
		{
			Source: "'a''b'\n",
			Wants: []*Token{
				&Token{Type: WORD, tree: &ast.Node{Str: "a'b"}},
				&Token{Type: '\n'},
			},
		},
		{
			Source: "a'b c'\n",
			Wants: []*Token{
				&Token{Type: WORD, tree: &ast.Node{Str: "a"}},
				&Token{Type: WORD, tree: &ast.Node{Str: "b c"}},
				&Token{Type: '\n'},
			},
		},
		{
			Source: "$a\n",
			Wants: []*Token{
				&Token{Type: '$'},
				&Token{Type: WORD, tree: &ast.Node{Str: "a"}},
				&Token{Type: '\n'},
			},
		},
		{
			Source: "a=1\n",
			Wants: []*Token{
				&Token{Type: WORD, tree: &ast.Node{Str: "a"}},
				&Token{Type: '='},
				&Token{Type: WORD, tree: &ast.Node{Str: "1"}},
				&Token{Type: '\n'},
			},
		},
	}
	for _, v := range tab {
		var l Lexer
		r := strings.NewReader(v.Source)
		l.Init(r, "-")

		t.Run(v.Source, func(t *testing.T) {
			var s yySymType
			for _, want := range v.Wants {
				k := l.Lex(&s)
				if k < 0 {
					t.Errorf("Lex() = EOF; want %d", want.Type)
					break
				}
				if k != want.Type {
					t.Errorf("typ = %d; want %d", k, want.Type)
				}
				if want.tree == nil {
					if s.tree != nil {
						t.Errorf("tree = %v; want nil", s.tree)
					}
				} else {
					if s.tree.Str != want.tree.Str {
						t.Errorf("str = %q; want %q", s.tree.Str, want.tree.Str)
					}
				}
			}
			if k := l.Lex(&s); k >= 0 {
				t.Errorf("Lex() = %d; want EOF", k)
			}
		})
	}
}
