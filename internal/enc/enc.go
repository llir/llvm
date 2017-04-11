package enc

import (
	"fmt"
	"strings"
)

// Global encodes a global name to its LLVM IR assembly representation.
//
// Examples:
//    "foo" -> "@foo"
//    "a b" -> `@"a\20b"`
//    "世" -> `@"\E4\B8\96"`
//
// References:
//    http://www.llvm.org/docs/LangRef.html#identifiers
func Global(name string) string {
	return "@" + EscapeIdent(name)
}

// Local encodes a local name to its LLVM IR assembly representation.
//
// Examples:
//    "foo" -> "%foo"
//    "a b" -> `%"a\20b"`
//    "世" -> `%"\E4\B8\96"`
//
// References:
//    http://www.llvm.org/docs/LangRef.html#identifiers
func Local(name string) string {
	return "%" + EscapeIdent(name)
}

// Metadata encodes a metadata name to its LLVM IR assembly representation.
//
// Examples:
//    "foo" -> "!foo"
//    "a b" -> `!a\20b`
//    "世" -> `!\E4\B8\96`
//
// References:
//    http://www.llvm.org/docs/LangRef.html#identifiers
func Metadata(name string) string {
	valid := func(b byte) bool {
		const metadataChar = tail + `\`
		return strings.IndexByte(metadataChar, b) != -1
	}
	return "!" + Escape(name, valid)
}

const (
	// decimal specifies the decimal digit characters.
	decimal = "0123456789"
	// upper specifies the uppercase letters.
	upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// lower specifies the lowercase letters.
	lower = "abcdefghijklmnopqrstuvwxyz"
	// alpha specifies the alphabetic characters.
	alpha = upper + lower
	// head is the set of valid characters for the first character of an
	// identifier.
	head = alpha + "$-._"
	// tail is the set of valid characters for the remaining characters of an
	// identifier (i.e. all characters in the identifier except the first). All
	// characters of a label may be from the tail set, even the first character.
	tail = head + decimal
)

// EscapeIdent replaces any characters which are not valid in identifiers with
// corresponding hexadecimal escape sequence (\XX).
func EscapeIdent(s string) string {
	// Check if a replacement is required.
	extra := 0
	for i := 0; i < len(s); i++ {
		if strings.IndexByte(tail, s[i]) == -1 {
			// Two extra bytes are required for each invalid byte; e.g.
			//    "#" -> `\23`
			//    "世" -> `\E4\B8\96`
			extra += 2
		}
	}
	if extra == 0 {
		return s
	}

	// Replace invalid characters.
	const hextable = "0123456789ABCDEF"
	buf := make([]byte, len(s)+extra)
	j := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if strings.IndexByte(tail, b) != -1 {
			buf[j] = b
			j++
			continue
		}
		buf[j] = '\\'
		buf[j+1] = hextable[b>>4]
		buf[j+2] = hextable[b&0x0F]
		j += 3
	}
	// Add surrounding quotes.
	return `"` + string(buf) + `"`
}

// EscapeString replaces any characters categorized as invalid in string
// literals with corresponding hexadecimal escape sequence (\XX).
func EscapeString(s string) string {
	valid := func(b byte) bool {
		return ' ' <= b && b <= '~' && b != '"' && b != '\\'
	}
	return Escape(s, valid)
}

// Escape replaces any characters categorized as invalid by the valid function
// with corresponding hexadecimal escape sequence (\XX).
func Escape(s string, valid func(b byte) bool) string {
	// Check if a replacement is required.
	extra := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		switch {
		case !valid(b):
			// Two extra bytes are required for each invalid byte; e.g.
			//    "#" -> `\23`
			//    "世" -> `\E4\B8\96`
			extra += 2
		default:
			// no extra bytes required.
		}
	}
	if extra == 0 {
		return s
	}

	// Replace invalid characters.
	const hextable = "0123456789ABCDEF"
	buf := make([]byte, len(s)+extra)
	j := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		switch {
		case !valid(b):
			buf[j] = '\\'
			buf[j+1] = hextable[b>>4]
			buf[j+2] = hextable[b&0x0F]
			j += 3
		default:
			buf[j] = b
			j++
		}
	}
	return string(buf)
}

// Unquote interprets s as a double-quoted string literal, returning the string
// value that s quotes.
func Unquote(s string) string {
	if len(s) < 2 {
		panic(fmt.Errorf("invalid length of quoted string; expected >= 2, got %d", len(s)))
	}
	if !strings.HasPrefix(s, `"`) {
		panic(fmt.Errorf("invalid quoted string `%s`; missing quote character prefix", s))
	}
	if !strings.HasSuffix(s, `"`) {
		panic(fmt.Errorf("invalid quoted string `%s`; missing quote character suffix", s))
	}
	// Skip double-quotes.
	s = s[1 : len(s)-1]
	return Unescape(s)
}

// Unescape replaces hexadecimal escape sequences (\xx) in s with their
// corresponding characters.
func Unescape(s string) string {
	if !strings.ContainsRune(s, '\\') {
		return s
	}
	j := 0
	buf := []byte(s)
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == '\\' && i+2 < len(s) {
			if s[i+1] == '\\' {
				b = '\\'
				i++
			} else {
				x1, ok := unhex(s[i+1])
				if ok {
					x2, ok := unhex(s[i+2])
					if ok {
						b = x1<<4 | x2
						i += 2
					}
				}
			}
		}
		if i != j {
			buf[j] = b
		}
		j++
	}
	return string(buf[:j])
}

// unhex returns the numeric value represented by the hexadecimal digit b. It
// returns false if b is not a hexadecimal digit.
func unhex(b byte) (v byte, ok bool) {
	// This is an adapted copy of the unhex function from the strconv package,
	// which is goverend by a BSD-style license.
	switch {
	case '0' <= b && b <= '9':
		return b - '0', true
	case 'a' <= b && b <= 'f':
		return b - 'a' + 10, true
	case 'A' <= b && b <= 'F':
		return b - 'A' + 10, true
	}
	return 0, false
}
