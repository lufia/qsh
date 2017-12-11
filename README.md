# qsh
q shell

## INSTALLATION

```console
$ make
```

## Language

### Comment

```
# comment
```

### Variables

Simple declaration.

```
a=1		# single value
a=(1 2)	# array

echo $a
```

Indirect reference

```
arch=lib_amd64
lib_amd64=lib64
lib_x86=lib

echo $$arch		# echo lib64
```

### If statement

```
if { true } {
	echo ok
}
```

### For statement

```
for i in 1 2 3 {
	echo $i
}
```

## Advent calendar

* https://blog.zoncoen.net/2015/12/22/cli-toml-processor-with-goyacc/
* https://qiita.com/draftcode/items/c9f2422fca14133c7f6a
* https://dev.classmethod.jp/etc/goyacc-json-generator/
