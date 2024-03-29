// Code generated by "string2enum -linecomment -type Tail ../../ir/enum"; DO NOT EDIT.

package enum

import (
	"fmt"

	"github.com/llir/llvm/ir/enum"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.TailNone-0]
	_ = x[enum.TailMustTail-1]
	_ = x[enum.TailNoTail-2]
	_ = x[enum.TailTail-3]
}

const _Tail_name = "nonemusttailnotailtail"

var _Tail_index = [...]uint8{0, 4, 12, 18, 22}

// TailFromString returns the Tail enum corresponding to s.
func TailFromString(s string) enum.Tail {
	if len(s) == 0 {
		return 0
	}
	for i := range _Tail_index[:len(_Tail_index)-1] {
		if s == _Tail_name[_Tail_index[i]:_Tail_index[i+1]] {
			return enum.Tail(i)
		}
	}
	panic(fmt.Errorf("unable to locate Tail enum corresponding to %q", s))
}

func _(s string) {
	// Check for duplicate string values in type "Tail".
	switch s {
	// 0
	case "none":
	// 1
	case "musttail":
	// 2
	case "notail":
	// 3
	case "tail":
	}
}
