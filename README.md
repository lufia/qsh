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

### Redirection

```
# output
echo hello >out

# append
echo hello >>out

# input
echo hello <in
```

## Advent calendar

* https://blog.zoncoen.net/2015/12/22/cli-toml-processor-with-goyacc/
* https://qiita.com/draftcode/items/c9f2422fca14133c7f6a
* https://dev.classmethod.jp/etc/goyacc-json-generator/
