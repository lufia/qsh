package main

import (
	"io"
	"log"
	"text/scanner"
)

type Lexer struct {
	s        scanner.Scanner
	filename string
	lineno   int
	EOF      bool
}

func (l *Lexer) Init(r io.Reader, filename string) {
	l.filename = filename
	l.s.Init(r)
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok := l.s.Scan()
	if tok == scanner.EOF {
		l.EOF = true
		return 0
	}
	lval.tree = &Tree{
		typ: WORD,
		str: l.s.TokenText(),
	}
	return WORD
}

func (l *Lexer) Error(msg string) {
	log.Printf("%s:%d: %s\n", l.filename, l.lineno, msg)
}
