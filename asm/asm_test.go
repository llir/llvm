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
		// i=0
		{s: "foo", want: "@foo"},
		// i=1
		{s: "a b", want: `@a\20b`},
		// i=2
		{s: "$a", want: "@$a"},
		// i=3
		{s: "-a", want: "@-a"},
		// i=4
		{s: ".a", want: "@.a"},
		// i=5
		{s: "_a", want: "@_a"},
		// i=6
		{s: "#a", want: `@\23a`},
		// i=7
		{s: "a b#c", want: `@a\20b\23c`},
		// i=8
		{s: "2", want: "@2"},
		// i=9
		{s: "foo世bar", want: `@foo\E4\B8\96bar`},
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
		// i=0
		{s: "foo", want: "%foo"},
		// i=1
		{s: "a b", want: `%a\20b`},
		// i=2
		{s: "$a", want: "%$a"},
		// i=3
		{s: "-a", want: "%-a"},
		// i=4
		{s: ".a", want: "%.a"},
		// i=5
		{s: "_a", want: "%_a"},
		// i=6
		{s: "#a", want: `%\23a`},
		// i=7
		{s: "a b#c", want: `%a\20b\23c`},
		// i=8
		{s: "2", want: "%2"},
		// i=9
		{s: "foo世bar", want: `%foo\E4\B8\96bar`},
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
