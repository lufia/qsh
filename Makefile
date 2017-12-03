TARG=qsh

GOFILES=$(wildcard *.go)

.PHONY: all clean

all: $(TARG)

qsh: y.go $(GOFILES)
	go build

y.go: gram.y
	goyacc $<

clean:
	rm -f $(TARG) y.*
