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
	go get -u golang.org/x/tools/cmd/goyacc

$(STRINGER):
	go get -u golang.org/x/tools/cmd/stringer

ast/lextype_string.go ast/direction_string.go: ast/node.go $(STRINGER)
	go generate $<

y.go: gram.y
	goyacc $<

clean:
	rm -f $(TARG) y.output
