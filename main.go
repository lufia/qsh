package main

import (
	"flag"
	"log"
	"os"
)

var (
	flagDebug = flag.Bool("d", false, "debug")
)

func main() {
	log.SetFlags(0)

	if *flagDebug {
		yyDebug = 2
	}
	var l Lexer
	l.Init(os.Stdin, "<stdin>")
	for yyParse(&l) == 0 {
	}
}
