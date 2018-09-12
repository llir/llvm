package enc

import (
	"testing"
)

func TestGlobal(t *testing.T) {
	golden := []struct {
		s    string
		want string
	}{
		// i=0
		{s: "foo", want: "@foo"},
		// i=1
		{s: "a b", want: `@"a\20b"`},
		// i=2
		{s: "$a", want: "@$a"},
		// i=3
		{s: "-a", want: "@-a"},
		// i=4
		{s: ".a", want: "@.a"},
		// i=5
		{s: "_a", want: "@_a"},
		// i=6
		{s: "#a", want: `@"\23a"`},
		// i=7
		{s: "a b#c", want: `@"a\20b\23c"`},
		// i=8
		{s: "2", want: "@2"},
		// i=9
		{s: "foo世bar", want: `@"foo\E4\B8\96bar"`},
	}
	for i, g := range golden {
		got := Global(g.s)
		if g.want != got {
			t.Errorf("i=%d: name mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestLocal(t *testing.T) {
	golden := []struct {
		s    string
		want string
	}{
		// i=0
		{s: "foo", want: "%foo"},
		// i=1
		{s: "a b", want: `%"a\20b"`},
		// i=2
		{s: "$a", want: "%$a"},
		// i=3
		{s: "-a", want: "%-a"},
		// i=4
		{s: ".a", want: "%.a"},
		// i=5
		{s: "_a", want: "%_a"},
		// i=6
		{s: "#a", want: `%"\23a"`},
		// i=7
		{s: "a b#c", want: `%"a\20b\23c"`},
		// i=8
		{s: "2", want: "%2"},
		// i=9
		{s: "foo世bar", want: `%"foo\E4\B8\96bar"`},
	}
	for i, g := range golden {
		got := Local(g.s)
		if g.want != got {
			t.Errorf("i=%d: name mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestLabel(t *testing.T) {
	golden := []struct {
		s    string
		want string
	}{
		// i=0
		{s: "foo", want: "foo:"},
		// i=1
		{s: "a b", want: `"a\20b":`},
		// i=2
		{s: "$a", want: "$a:"},
		// i=3
		{s: "-a", want: "-a:"},
		// i=4
		{s: ".a", want: ".a:"},
		// i=5
		{s: "_a", want: "_a:"},
		// i=6
		{s: "#a", want: `"\23a":`},
		// i=7
		{s: "a b#c", want: `"a\20b\23c":`},
		// i=8
		{s: "2", want: "2:"},
		// i=9
		{s: "foo世bar", want: `"foo\E4\B8\96bar":`},
	}
	for i, g := range golden {
		got := Label(g.s)
		if g.want != got {
			t.Errorf("i=%d: name mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestAttrGroupID(t *testing.T) {
	golden := []struct {
		s    string
		want string
	}{
		// i=0
		{s: "42", want: "#42"},
	}
	for i, g := range golden {
		got := AttrGroupID(g.s)
		if g.want != got {
			t.Errorf("i=%d: name mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestComdat(t *testing.T) {
	golden := []struct {
		s    string
		want string
	}{
		// i=0
		{s: "foo", want: "$foo"},
		// i=1
		{s: "a b", want: `$"a\20b"`},
		// i=2
		{s: "$a", want: "$$a"},
		// i=3
		{s: "-a", want: "$-a"},
		// i=4
		{s: ".a", want: "$.a"},
		// i=5
		{s: "_a", want: "$_a"},
		// i=6
		{s: "#a", want: `$"\23a"`},
		// i=7
		{s: "a b#c", want: `$"a\20b\23c"`},
		// i=8
		{s: "2", want: "$2"},
		// i=9
		{s: "foo世bar", want: `$"foo\E4\B8\96bar"`},
	}
	for i, g := range golden {
		got := Comdat(g.s)
		if g.want != got {
			t.Errorf("i=%d: name mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestMetadata(t *testing.T) {
	golden := []struct {
		s    string
		want string
	}{
		// i=0
		{s: "foo", want: "!foo"},
		// i=1
		{s: "a b", want: `!a\20b`},
		// i=2
		{s: "$a", want: "!$a"},
		// i=3
		{s: "-a", want: "!-a"},
		// i=4
		{s: ".a", want: "!.a"},
		// i=5
		{s: "_a", want: "!_a"},
		// i=6
		{s: "#a", want: `!\23a`},
		// i=7
		{s: "a b#c", want: `!a\20b\23c`},
		// i=8
		{s: "2", want: "!2"},
		// i=9
		{s: "foo世bar", want: `!foo\E4\B8\96bar`},
	}
	for i, g := range golden {
		got := Metadata(g.s)
		if g.want != got {
			t.Errorf("i=%d: name mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestEscapeString(t *testing.T) {
	golden := []struct {
		s    string
		want string
	}{
		// i=0
		{s: "foo", want: "foo"},
		// i=1
		{s: "a b", want: `a b`},
		// i=2
		{s: "$a", want: "$a"},
		// i=3
		{s: "-a", want: "-a"},
		// i=4
		{s: ".a", want: ".a"},
		// i=5
		{s: "_a", want: "_a"},
		// i=6
		{s: "#a", want: `#a`},
		// i=7
		{s: "a b#c", want: `a b#c`},
		// i=8
		{s: "2", want: "2"},
		// i=9
		{s: "foo世bar", want: `foo\E4\B8\96bar`},
		// i=10
		{s: `foo \ bar`, want: `foo \5C bar`},
	}
	for i, g := range golden {
		got := EscapeString(g.s)
		if g.want != got {
			t.Errorf("i=%d: string mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestEscape(t *testing.T) {
	golden := []struct {
		s    string
		want string
	}{
		// i=0
		{s: "foo", want: "foo"},
		// i=1
		{s: "a b", want: `a b`},
		// i=2
		{s: "$a", want: "$a"},
		// i=3
		{s: "-a", want: "-a"},
		// i=4
		{s: ".a", want: ".a"},
		// i=5
		{s: "_a", want: "_a"},
		// i=6
		{s: "#a", want: `#a`},
		// i=7
		{s: "a b#c", want: `a b#c`},
		// i=8
		{s: "2", want: "2"},
		// i=9
		{s: "foo世bar", want: `foo\E4\B8\96bar`},
		// i=10
		{s: `foo \ bar`, want: `foo \5C bar`},
	}
	// isPrint reports whether the given byte is printable in ASCII.
	isPrint := func(b byte) bool {
		return ' ' <= b && b <= '~' && b != '"' && b != '\\'
	}
	for i, g := range golden {
		got := Escape(g.s, isPrint)
		if g.want != got {
			t.Errorf("i=%d: string mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestUnescape(t *testing.T) {
	golden := []struct {
		s    string
		want string
	}{
		// i=0
		{s: "foo", want: "foo"},
		// i=1
		{s: `a\20b`, want: "a b"},
		// i=2
		{s: "$a", want: "$a"},
		// i=3
		{s: "-a", want: "-a"},
		// i=4
		{s: ".a", want: ".a"},
		// i=5
		{s: "_a", want: "_a"},
		// i=6
		{s: `\23a`, want: "#a"},
		// i=7
		{s: `a\20b\23c`, want: "a b#c"},
		// i=8
		{s: "2", want: "2"},
		// i=9
		{s: `foo\E4\B8\96bar`, want: "foo世bar"},
		// i=10
		{s: `foo \5C bar`, want: `foo \ bar`},
		// i=11
		{s: `foo \\ bar`, want: `foo \ bar`},
	}
	for i, g := range golden {
		got := Unescape(g.s)
		if g.want != got {
			t.Errorf("i=%d: string mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestQuote(t *testing.T) {
	golden := []struct {
		s    string
		want string
	}{
		// i=0
		{s: "foo", want: `"foo"`},
		// i=1
		{s: "a b", want: `"a b"`},
		// i=2
		{s: "$a", want: `"$a"`},
		// i=3
		{s: "-a", want: `"-a"`},
		// i=4
		{s: ".a", want: `".a"`},
		// i=5
		{s: "_a", want: `"_a"`},
		// i=6
		{s: "#a", want: `"#a"`},
		// i=7
		{s: "a b#c", want: `"a b#c"`},
		// i=8
		{s: "2", want: `"2"`},
		// i=9
		{s: "foo世bar", want: `"foo\E4\B8\96bar"`},
		// i=10
		{s: `foo \ bar`, want: `"foo \5C bar"`},
	}
	for i, g := range golden {
		got := Quote(g.s)
		if g.want != got {
			t.Errorf("i=%d: string mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestUnquote(t *testing.T) {
	golden := []struct {
		s    string
		want string
	}{
		// i=0
		{s: `"foo"`, want: "foo"},
		// i=1
		{s: `"a\20b"`, want: "a b"},
		// i=2
		{s: `"$a"`, want: "$a"},
		// i=3
		{s: `"-a"`, want: "-a"},
		// i=4
		{s: `".a"`, want: ".a"},
		// i=5
		{s: `"_a"`, want: "_a"},
		// i=6
		{s: `"\23a"`, want: "#a"},
		// i=7
		{s: `"a\20b\23c"`, want: "a b#c"},
		// i=8
		{s: `"2"`, want: "2"},
		// i=9
		{s: `"foo\E4\B8\96bar"`, want: "foo世bar"},
		// i=10
		{s: `"foo \5C bar"`, want: `foo \ bar`},
		// i=11
		{s: `"foo \\ bar"`, want: `foo \ bar`},
	}
	for i, g := range golden {
		got := Unquote(g.s)
		if g.want != got {
			t.Errorf("i=%d: string mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func BenchmarkGlobalNoReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global("$foo_bar_baz")
	}
}

func BenchmarkGlobalReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global("$foo bar#baz")
	}
}

func BenchmarkLocalNoReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Local("$foo_bar_baz")
	}
}

func BenchmarkLocalReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Local("$foo bar#baz")
	}
}
