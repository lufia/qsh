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
	flag.Parse()
	log.SetFlags(0)
	log.SetPrefix("qsh: ")

	if *flagDebug {
		yyDebug = 2
	}
	if flag.NArg() == 0 {
		Run(os.Stdin, "<stdin>", nil)
	} else {
		args := flag.Args()
		f, err := os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}
		Run(f, args[0], args[1:])
		f.Close()
	}
}

func Run(f io.Reader, name string, args []string) {
	var l Lexer
	fin := bufio.NewReader(f)
	l.Init(fin, name)
	for yyParse(&l) == 0 {
	}
	if l.err != nil && l.err != io.EOF {
		l.Error(l.err.Error())
	}
}
