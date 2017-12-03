package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	var l Lexer
	l.Init(os.Stdin, "<stdin>")
	for !l.EOF {
		yyParse(&l)
	}
}
