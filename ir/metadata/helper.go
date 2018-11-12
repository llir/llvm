package metadata

import (
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/enum"
)

// diFlagsString returns the string representation of the given debug
// information flags.
func diFlagsString(flags enum.DIFlag) string {
	if flags == enum.DIFlagZero {
		return flags.String()
	}
	var ss []string
	if flag := flags & 0x3; flag != 0 {
		ss = append(ss, flag.String())
	}
	for mask := enum.DIFlagFirst; mask <= enum.DIFlagLast; mask <<= 1 {
		if flags&mask != 0 {
			ss = append(ss, mask.String())
		}
	}
	return strings.Join(ss, " | ")
}

// quote returns s as a double-quoted string literal.
func quote(s string) string {
	return enc.Quote([]byte(s))
}
