TARG=qsh

YACC=$(GOPATH)/bin/goyacc
STRINGER=$(GOPATH)/bin/stringer
GOFILES=\
	$(wildcard *.go */*.go)\
	ast/lextype_string.go\
	ast/direction_string.go\

.PHONY: all test clean

all: $(YACC) $(TARG)

qsh: y.go $(GOFILES)
	go build

test: y.go $(GOFILES)
	go test ./...

$(YACC):
	go install golang.org/x/tools/cmd/goyacc@latest

$(STRINGER):
	go install golang.org/x/tools/cmd/stringer@latest

ast/lextype_string.go ast/direction_string.go: ast/node.go $(STRINGER)
	go generate $<

y.go: gram.y
	goyacc $<

clean:
	rm -f $(TARG) y.output
