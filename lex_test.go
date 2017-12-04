package main

import (
	"strings"
	"testing"
)

type Token struct {
	Type int
	tree *Tree
}

func TestLexWords(t *testing.T) {
	tab := []struct {
		Source string
		Wants  []*Token
	}{
		{
			Source: "a bc d\n",
			Wants: []*Token{
				&Token{Type: WORD, tree: &Tree{str: "a"}},
				&Token{Type: WORD, tree: &Tree{str: "bc"}},
				&Token{Type: WORD, tree: &Tree{str: "d"}},
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
					if s.tree.str != want.tree.str {
						t.Errorf("str = %q; want %q", s.tree.str, want.tree.str)
					}
				}
			}
			if k := l.Lex(&s); k >= 0 {
				t.Errorf("Lex() = %d; want EOF", k)
			}
		})
	}
}
