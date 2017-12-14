package build

import (
	"errors"
	"path/filepath"
	"plugin"
	"strings"
	"unicode"
)

func load(s string) error {
	p, err := plugin.Open(s + ".so")
	if err != nil {
		return err
	}
	sym, err := p.Lookup(capitalize(s) + "Module")
	if err != nil {
		return err
	}
	tab, ok := sym.(*map[string]string)
	if !ok {
		return errors.New("no exported functions")
	}
	for key, v := range *tab {
		sym, err = p.Lookup(v)
		if err != nil {
			return err
		}
		f, ok := sym.(func([]string) ([]string, error))
		if !ok {
			continue
		}
		mtab[key] = f
	}
	return nil
}

func capitalize(s string) string {
	s = filepath.Base(s)
	ext := filepath.Ext(s)
	s = s[0 : len(s)-len(ext)]
	a := strings.Split(s, "_")
	t := ""
	for _, p := range a {
		if p == "" {
			continue
		}
		r := []rune(p)
		t += string(unicode.ToUpper(r[0])) + string(r[1:])
	}
	return t
}
