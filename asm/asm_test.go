package asm_test

import (
	"testing"

	"github.com/llir/llvm/asm"
)

func TestEncGlobalName(t *testing.T) {
	golden := []struct {
		s    string
		want string
	}{
		{s: "foo", want: "@foo"},
		{s: "a b", want: `@a\20b`},
		{s: "$a", want: "@$a"},
		{s: "-a", want: "@-a"},
		{s: ".a", want: "@.a"},
		{s: "_a", want: "@_a"},
		{s: "#a", want: `@\23a`},
		{s: "a b#c", want: `@a\20b\23c`},
		{s: "2", want: "@2"},
	}

	for i, g := range golden {
		got := asm.EncGlobalName(g.s)
		if got != g.want {
			t.Errorf("i=%d: name mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestEncLocalName(t *testing.T) {
	golden := []struct {
		s    string
		want string
	}{
		{s: "foo", want: "%foo"},
		{s: "a b", want: `%a\20b`},
		{s: "$a", want: "%$a"},
		{s: "-a", want: "%-a"},
		{s: ".a", want: "%.a"},
		{s: "_a", want: "%_a"},
		{s: "#a", want: `%\23a`},
		{s: "a b#c", want: `%a\20b\23c`},
		{s: "2", want: "%2"},
	}

	for i, g := range golden {
		got := asm.EncLocalName(g.s)
		if got != g.want {
			t.Errorf("i=%d: name mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func BenchmarkEncGlobalNameNoReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		asm.EncGlobalName("$foo_bar_baz")
	}
}

func BenchmarkEncGlobalNameReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		asm.EncGlobalName("$foo bar#baz")
	}
}

func BenchmarkEncLocalNameNoReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		asm.EncLocalName("$foo_bar_baz")
	}
}

func BenchmarkEncLocalNameReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		asm.EncLocalName("$foo bar#baz")
	}
}
