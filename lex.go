package main

import (
	"bytes"
	"io"
	"log"
	"unicode"

	"github.com/lufia/qsh/ast"
)

const (
	EOF = -1
)

var keywords = map[string]int{
	"if": IF,
}

type Lexer struct {
	s        io.RuneScanner
	err      error
	filename string
	lineno   int
	buf      bytes.Buffer
}

func (l *Lexer) getc() rune {
	if l.err != nil {
		return EOF
	}
	var c rune
	c, _, l.err = l.s.ReadRune()
	if l.err != nil {
		return EOF
	}
	return c
}

func (l *Lexer) ungetc() {
	if l.err != nil {
		return
	}
	l.err = l.s.UnreadRune()
}

func (l *Lexer) Init(s io.RuneScanner, filename string) {
	l.s = s
	l.filename = filename
	l.lineno = 1
}

func (l *Lexer) Lex(lval *yySymType) int {
	l.buf.Reset()
	lval.tree = nil

	c := l.getc()
	for isSpace(c) {
		c = l.getc()
	}
	switch c {
	case EOF:
		return -1
	case '=', '&', ';', '$', '{', '}', '(', ')':
		return int(c)
	case '\n':
		l.lineno++
		return int(c)
	case '\'':
		lval.tree = ast.Token(l.scanQuotedText())
		lval.tree.Quoted = true
		return WORD
	default:
		l.buf.WriteRune(c)
		s := l.scanText()
		if k, ok := keywords[s]; ok {
			return k
		}
		lval.tree = ast.Token(s)
		return WORD
	}
}

func (l *Lexer) scanQuotedText() string {
	for {
		c := l.getc()
		if c == EOF {
			break
		}
		if c == '\'' {
			if c1 := l.getc(); c1 != '\'' {
				l.ungetc()
				break
			}
		}
		l.buf.WriteRune(c)
	}
	return l.buf.String()
}

func (l *Lexer) scanText() string {
	for {
		c := l.getc()
		if c == EOF || unicode.IsSpace(c) || isDelim(c) {
			l.ungetc()
			break
		}
		l.buf.WriteRune(c)
	}
	return l.buf.String()
}

func isDelim(c rune) bool {
	switch c {
	case '=', '&', ';', '$', '{', '}', '(', ')', '\'':
		return true
	}
	return false
}

func isSpace(c rune) bool {
	switch c {
	case '\t', '\v', '\f', ' ', 0x85, 0xa0:
		return true
	}
	return false
}

func (l *Lexer) Error(msg string) {
	s := l.buf.String()
	log.Printf("%s:%d: %s near %q\n", l.filename, l.lineno, msg, s)
}
