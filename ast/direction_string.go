// Code generated by "stringer -type=Direction"; DO NOT EDIT.

package ast

import "strconv"

const _Direction_name = "READWRITEAPPENDHERE"

var _Direction_index = [...]uint8{0, 4, 9, 15, 19}

func (i Direction) String() string {
	if i < 0 || i >= Direction(len(_Direction_index)-1) {
		return "Direction(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Direction_name[_Direction_index[i]:_Direction_index[i+1]]
}