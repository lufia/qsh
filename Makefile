TARG=qsh

.PHONY: all clean

all: $(TARG)

qsh: y.go
	go build

y.go: gram.y
	goyacc $<

clean:
	rm -f $(TARG) y.*
