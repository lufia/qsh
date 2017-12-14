// +build !linux

package build

import (
	"errors"
)

func load(s string) error {
	return errors.New("load is not implement in this OS")
}
