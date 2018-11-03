package enc

import (
	"reflect"
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
		s    []byte
		want string
	}{
		// i=0
		{s: []byte("foo"), want: "foo"},
		// i=1
		{s: []byte("a b"), want: `a b`},
		// i=2
		{s: []byte("$a"), want: "$a"},
		// i=3
		{s: []byte("-a"), want: "-a"},
		// i=4
		{s: []byte(".a"), want: ".a"},
		// i=5
		{s: []byte("_a"), want: "_a"},
		// i=6
		{s: []byte("#a"), want: `#a`},
		// i=7
		{s: []byte("a b#c"), want: `a b#c`},
		// i=8
		{s: []byte("2"), want: "2"},
		// i=9
		{s: []byte("foo世bar"), want: `foo\E4\B8\96bar`},
		// i=10
		{s: []byte(`foo \ bar`), want: `foo \5C bar`},
		// i=11 (arbitrary data, invalid UTF-8)
		{s: []byte{'f', 'o', 'o', 0x81, 0x82, 'b', 'a', 'r'}, want: `foo\81\82bar`},
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
		s    []byte
		want string
	}{
		// i=0
		{s: []byte("foo"), want: "foo"},
		// i=1
		{s: []byte("a b"), want: `a b`},
		// i=2
		{s: []byte("$a"), want: "$a"},
		// i=3
		{s: []byte("-a"), want: "-a"},
		// i=4
		{s: []byte(".a"), want: ".a"},
		// i=5
		{s: []byte("_a"), want: "_a"},
		// i=6
		{s: []byte("#a"), want: `#a`},
		// i=7
		{s: []byte("a b#c"), want: `a b#c`},
		// i=8
		{s: []byte("2"), want: "2"},
		// i=9
		{s: []byte("foo世bar"), want: `foo\E4\B8\96bar`},
		// i=10
		{s: []byte(`foo \ bar`), want: `foo \5C bar`},
		// i=11 (arbitrary data, invalid UTF-8)
		{s: []byte{'f', 'o', 'o', 0x81, 0x82, 'b', 'a', 'r'}, want: `foo\81\82bar`},
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
		want []byte
	}{
		// i=0
		{s: "foo", want: []byte("foo")},
		// i=1
		{s: `a\20b`, want: []byte("a b")},
		// i=2
		{s: "$a", want: []byte("$a")},
		// i=3
		{s: "-a", want: []byte("-a")},
		// i=4
		{s: ".a", want: []byte(".a")},
		// i=5
		{s: "_a", want: []byte("_a")},
		// i=6
		{s: `\23a`, want: []byte("#a")},
		// i=7
		{s: `a\20b\23c`, want: []byte("a b#c")},
		// i=8
		{s: "2", want: []byte("2")},
		// i=9
		{s: `foo\E4\B8\96bar`, want: []byte("foo世bar")},
		// i=10
		{s: `foo \5C bar`, want: []byte(`foo \ bar`)},
		// i=11
		{s: `foo \\ bar`, want: []byte(`foo \ bar`)},
		// i=12 (arbitrary data, invalid UTF-8)
		{s: `foo\81\82bar`, want: []byte{'f', 'o', 'o', 0x81, 0x82, 'b', 'a', 'r'}},
	}
	for i, g := range golden {
		got := Unescape(g.s)
		if !reflect.DeepEqual(g.want, got) {
			t.Errorf("i=%d: string mismatch; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestQuote(t *testing.T) {
	golden := []struct {
		s    []byte
		want string
	}{
		// i=0
		{s: []byte("foo"), want: `"foo"`},
		// i=1
		{s: []byte("a b"), want: `"a b"`},
		// i=2
		{s: []byte("$a"), want: `"$a"`},
		// i=3
		{s: []byte("-a"), want: `"-a"`},
		// i=4
		{s: []byte(".a"), want: `".a"`},
		// i=5
		{s: []byte("_a"), want: `"_a"`},
		// i=6
		{s: []byte("#a"), want: `"#a"`},
		// i=7
		{s: []byte("a b#c"), want: `"a b#c"`},
		// i=8
		{s: []byte("2"), want: `"2"`},
		// i=9
		{s: []byte("foo世bar"), want: `"foo\E4\B8\96bar"`},
		// i=10
		{s: []byte(`foo \ bar`), want: `"foo \5C bar"`},
		// i=11 (arbitrary data, invalid UTF-8)
		{s: []byte{'f', 'o', 'o', 0x81, 0x82, 'b', 'a', 'r'}, want: `"foo\81\82bar"`},
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
		want []byte
	}{
		// i=0
		{s: `"foo"`, want: []byte("foo")},
		// i=1
		{s: `"a\20b"`, want: []byte("a b")},
		// i=2
		{s: `"$a"`, want: []byte("$a")},
		// i=3
		{s: `"-a"`, want: []byte("-a")},
		// i=4
		{s: `".a"`, want: []byte(".a")},
		// i=5
		{s: `"_a"`, want: []byte("_a")},
		// i=6
		{s: `"\23a"`, want: []byte("#a")},
		// i=7
		{s: `"a\20b\23c"`, want: []byte("a b#c")},
		// i=8
		{s: `"2"`, want: []byte("2")},
		// i=9
		{s: `"foo\E4\B8\96bar"`, want: []byte("foo世bar")},
		// i=10
		{s: `"foo \5C bar"`, want: []byte(`foo \ bar`)},
		// i=11
		{s: `"foo \\ bar"`, want: []byte(`foo \ bar`)},
		// i=12 (arbitrary data, invalid UTF-8)
		{s: `"foo\81\82bar"`, want: []byte{'f', 'o', 'o', 0x81, 0x82, 'b', 'a', 'r'}},
	}
	for i, g := range golden {
		got := Unquote(g.s)
		if !reflect.DeepEqual(g.want, got) {
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
