package build

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	vtab = make(map[string][]string)
)

const (
	ListSeparator  = string(filepath.ListSeparator)
	FieldSeparator = " "
)

var listTab = map[string]struct{}{
	"PATH": struct{}{},
}

func ImportVar(name string) {
	s, ok := os.LookupEnv(name)
	if !ok {
		return
	}
	if isExportedList(name) {
		vtab[name] = strings.Split(s, ListSeparator)
		return
	}
	vtab[name] = strings.Split(s, FieldSeparator)
}

func isExportedList(name string) bool {
	return isExportedVar(name) && strings.Contains(name, "PATH")
}

func isExportedVar(name string) bool {
	return strings.ToUpper(name) == name
}

func LookupVar(name string) ([]string, bool) {
	v, ok := vtab[name]
	if ok {
		return v, true
	}
	return nil, false
}

func UpdateVar(name string, v []string) {
	if isExportedList(name) {
		s := strings.Join(v, ListSeparator)
		os.Setenv(name, s)
	} else if isExportedVar(name) {
		s := strings.Join(v, FieldSeparator)
		os.Setenv(name, s)
	}
	vtab[name] = v
}

func UnsetVar(name string) {
	if isExportedVar(name) {
		os.Unsetenv(name)
	}
	delete(vtab, name)
}

func lastStatus() string {
	v, ok := LookupVar("status")
	if !ok {
		return ""
	}
	return strings.Join(v, FieldSeparator)
}

func isSuccess() bool {
	return lastStatus() == ""
}

func updateStatus(err error) {
	if err == nil {
		UnsetVar("status")
	} else {
		UpdateVar("status", []string{err.Error()})
	}
}
