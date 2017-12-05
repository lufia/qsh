TARG=qsh

YACC=$(GOPATH)/bin/goyacc
GOFILES=$(wildcard *.go */*.go)

.PHONY: all clean

all: $(YACC) $(TARG)

qsh: y.go $(GOFILES)
	go build

$(YACC):
	go get -u golang.org/x/tools/cmd/goyacc

y.go: gram.y
	goyacc $<

clean:
	rm -f $(TARG) y.*
