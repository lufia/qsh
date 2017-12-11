package cgen

type Cmd struct {
	pc    int
	words []string
}

var (
	vtab = make(map[string][]string)
)

type String string

func (s String) Push(cmd *Cmd) {
	cmd.words = append(cmd.words, string(s))
	cmd.pc++
}

func Simple(cmd *Cmd) {
	fmt.Println(cmd.words)
	cmd.pc++
}

func Var(cmd *Cmd) {
	n := len(cmd.words)
	name := cmd.words[n-1]
	v := vtab[name]
	if len(v) != 1 {
		// TODO: variable name is not singleton
		return
	}
	cmd.words[n-1] = v[0]
	cmd.pc++
}

func Assign(cmd *Cmd) {
	n := len(cmd.words)
	name := cmd.words[n-1]
	value := cmd.words[n-2]
	vtab[name] = []string{value}
	cmd.words = cmd.words[0 : n-2]
	cmd.pc++
}
