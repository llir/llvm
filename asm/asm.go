// Package asm provides support for reading and writing the assembly language
// representation of LLVM IR.
//
// ref: http://llvm.org/docs/LangRef.html
package asm

import (
	"bytes"
	"encoding/hex"
	"strings"
	"unicode/utf8"
)

// EncGlobalName encodes a global name to its LLVM IR assembly representation.
//
// Examples:
//    "foo" -> "@foo"
//    "a b" -> `@"a\20b"`
func EncGlobalName(name string) string {
	return "@" + escape(name)
}

// EncLocalName encodes a local name to its LLVM IR assembly representation.
//
// Examples:
//    "foo" -> "%foo"
//    "a b" -> `%"a\20b"`
func EncLocalName(name string) string {
	return "%" + escape(name)
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

// escape replaces any characters which are not valid in identifiers with their
// hexadecimal escape sequence (\XX).
func escape(s string) string {
	// Check if a replacement is required.
	replace := false
	for _, r := range s {
		if !strings.ContainsRune(tail, r) {
			replace = true
			break
		}
	}
	if !replace {
		return s
	}

	// Replace invalid characters.
	buf := new(bytes.Buffer)
	var src [utf8.MaxRune]byte
	var dst [2 * utf8.MaxRune]byte
	for _, r := range s {
		if strings.ContainsRune(tail, r) {
			buf.WriteRune(r)
			continue
		}
		n := utf8.EncodeRune(src[:], r)
		m := hex.Encode(dst[:], src[:n])
		buf.WriteByte('\\')
		buf.Write(dst[:m])
	}
	return buf.String()
}
