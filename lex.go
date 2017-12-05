package main

import (
	"bytes"
	"io"
	"log"
	"text/scanner"
	"unicode"

	"github.com/lufia/qsh/ast"
)

type Lexer struct {
	s        scanner.Scanner
	filename string
	lineno   int
	buf      bytes.Buffer
}

func (l *Lexer) Init(r io.Reader, filename string) {
	l.filename = filename
	l.s.Error = func(s *scanner.Scanner, msg string) {
		l.Error(msg)
	}
	l.s.Init(r)
}

func (l *Lexer) Lex(lval *yySymType) int {
	l.buf.Reset()
	lval.tree = nil

	c := l.s.Next()
	for isSpace(c) {
		c = l.s.Next()
	}
	switch c {
	case scanner.EOF:
		return -1
	case '\n':
		return int(c)
	case '\'':
		lval.tree = ast.Token(l.scanQuotedText())
		lval.tree.Quoted = true
		return WORD
	default:
		l.buf.WriteRune(c)
		lval.tree = ast.Token(l.scanText())
		return WORD
	}
}

func (l *Lexer) scanQuotedText() string {
	for {
		c := l.s.Next()
		if c == scanner.EOF {
			break
		}
		if c == '\'' {
			if l.s.Peek() != '\'' {
				break
			}
			l.s.Next()
		}
		l.buf.WriteRune(c)
	}
	return l.buf.String()
}

func (l *Lexer) scanText() string {
	for {
		c := l.s.Peek()
		if c == scanner.EOF || unicode.IsSpace(c) || c == '\'' {
			break
		}
		l.buf.WriteRune(l.s.Next())
	}
	return l.buf.String()
}

func isSpace(c rune) bool {
	switch c {
	case '\t', '\v', '\f', ' ', 0x85, 0xa0:
		return true
	}
	return false
}

func (l *Lexer) Error(msg string) {
	log.Printf("%s:%d: %s\n", l.filename, l.lineno, msg)
}
