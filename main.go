package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/lufia/qsh/build"
)

var (
	flagDebug = flag.Bool("d", false, "debug")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: qsh [-d] [file [arg ...]]\n")
	os.Exit(2)
}

func init() {
	a := os.Environ()
	for _, s := range a {
		kv := strings.SplitN(s, "=", 2)
		if len(kv) != 2 {
			continue
		}
		build.ImportVar(kv[0])
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("qsh: ")
	flag.Usage = usage
	flag.Parse()

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
