package main

import (
	"bufio"
	"flag"
	"io"
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
	fin := bufio.NewReader(os.Stdin)
	l.Init(fin, "<stdin>")
	for yyParse(&l) == 0 {
	}
	if l.err != nil && l.err != io.EOF {
		l.Error(l.err.Error())
	}
}
