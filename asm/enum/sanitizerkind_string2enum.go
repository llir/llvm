// Code generated by "string2enum -linecomment -type SanitizerKind ../../ir/enum"; DO NOT EDIT.

package enum

import (
	"fmt"

	"github.com/llir/llvm/ir/enum"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.SanitizerKindNoSanitizeAddress-1]
	_ = x[enum.SanitizerKindNoSanitizeHwAddress-2]
	_ = x[enum.SanitizerKindSanitizeAddressDynInit-3]
	_ = x[enum.SanitizerKindSanitizeMemTag-4]
}

const _SanitizerKind_name = "no_sanitize_addressno_sanitize_hwaddresssanitize_address_dyninitsanitize_memtag"

var _SanitizerKind_index = [...]uint8{0, 19, 40, 64, 79}

// SanitizerKindFromString returns the SanitizerKind enum corresponding to s.
func SanitizerKindFromString(s string) enum.SanitizerKind {
	if len(s) == 0 {
		return 0
	}
	for i := range _SanitizerKind_index[:len(_SanitizerKind_index)-1] {
		if s == _SanitizerKind_name[_SanitizerKind_index[i]:_SanitizerKind_index[i+1]] {
			return enum.SanitizerKind(i + 1)
		}
	}
	panic(fmt.Errorf("unable to locate SanitizerKind enum corresponding to %q", s))
}

func _(s string) {
	// Check for duplicate string values in type "SanitizerKind".
	switch s {
	// 1
	case "no_sanitize_address":
	// 2
	case "no_sanitize_hwaddress":
	// 3
	case "sanitize_address_dyninit":
	// 4
	case "sanitize_memtag":
	}
}
