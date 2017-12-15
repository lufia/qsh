# qsh
q shell

## INSTALLATION

```console
$ go get github.com/lufia/qsh
```

## Language

### Comment

```
# comment
```

### Variables

Simple declaration.

```
a=1     # scalar; same as a=(1)
a=(1 2) # array

echo $a
# Output:
# 1 2
```

if variable name all are capital letter, it handles as environment variables.

```
RESULT_CODE=1
bash -c 'echo $RESULT_CODE'
# Output: 1
```

Indirect reference

```
arch=lib_amd64
lib_amd64=lib64
lib_x86=lib

echo $$arch
# Output:
# lib64
```

### If statement

```
if { true } {
	echo ok
}
# Output:
# ok
```

### For statement

```
for i in 1 2 3 {
	echo $i
}
# Output:
# 1
# 2
# 3
```

### Modules

```
ackage main

var SampleModule = map[string]string{
	"hello": "Hello",
}

func Hello(args []string) ([]string, error) {
	return []string{"hello"}, nil
}
```

shell

```
load sample

echo ${hello a b c}
```

### Expression

```
true && echo true

false || echo false
```

### Redirection

```
# output
echo hello >out

# append
echo hello >>out

# input
echo hello <in

# pipe
echo hello | wc
```

## TODO

- [ ] Basic
	- [x] comments
	- [x] command execution
	- [ ] background execution
	- [ ] inline execution
	- [ ] glob
- [ ] Variable
	- [x] assign
	- [ ] indexing
	- [ ] concat
	- [ ] environment
- [ ] Redirection
	- [x] read
	- [x] write
	- [x] append
	- [ ] error
	- [x] pipe
	- [ ] dup
- [ ] Statements
	- [x] if
	- [ ] if-else
	- [x] for
	- [ ] switch
	- [ ] load
- [x] Expression
	- [x] `&&`
	- [x] `||`

## Advent calendar

* https://blog.zoncoen.net/2015/12/22/cli-toml-processor-with-goyacc/
* https://qiita.com/draftcode/items/c9f2422fca14133c7f6a
* https://dev.classmethod.jp/etc/goyacc-json-generator/
