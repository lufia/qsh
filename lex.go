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
	"if":   IF,
	"for":  FOR,
	"in":   IN,
	"load": LOAD,
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
	for {
		if c == '#' {
			l.skipLine()
			l.lineno++
			c = l.getc()
			continue
		}
		if !isSpace(c) {
			break
		}
		c = l.getc()
	}
	switch c {
	case EOF:
		return -1
	case '=', ';', '$', '{', '}', '(', ')':
		return int(c)
	case '&':
		c = l.getc()
		if c == '&' {
			return ANDAND
		}
		l.ungetc()
		return '&'
	case '|':
		c = l.getc()
		if c == '|' {
			return OROR
		}
		l.ungetc()
		return '|'
	case '<':
		lval.tree = ast.Redir(ast.READ)
		return REDIR
	case '>':
		c = l.getc()
		if c == '>' {
			lval.tree = ast.Redir(ast.APPEND)
			return REDIR
		}
		l.ungetc()
		lval.tree = ast.Redir(ast.WRITE)
		return REDIR
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

func (l *Lexer) skipLine() {
	for {
		c := l.getc()
		if c == EOF || c == '\n' {
			break
		}
	}
}

func isDelim(c rune) bool {
	switch c {
	case '=', '&', '|', ';', '$', '{', '}', '(', ')', '<', '>', '\'':
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
