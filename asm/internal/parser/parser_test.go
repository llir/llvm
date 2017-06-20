package parser_test

import (
	"io/ioutil"
	"testing"

	"github.com/llir/llvm/asm"
)

func TestParseString(t *testing.T) {
	// Round-trip test of the parser.
	golden := []struct {
		path string
	}{
		// Empty module.
		{path: "../../testdata/empty.ll"},
		// Instructions.
		{path: "../../testdata/inst_binary.ll"},
		{path: "../../testdata/inst_bitwise.ll"},
		{path: "../../testdata/inst_vector.ll"},
	}
	for _, g := range golden {
		buf, err := ioutil.ReadFile(g.path)
		if err != nil {
			t.Errorf("%q: unable to read file; %v", g.path, err)
			continue
		}
		input := string(buf)
		m, err := asm.ParseString(input)
		if err != nil {
			t.Errorf("%q: unable to parse file; %v", g.path, err)
			continue
		}
		want := input
		// Read foo.ll.golden if present.
		if buf, err := ioutil.ReadFile(g.path + ".golden"); err == nil {
			want = string(buf)
		}
		got := m.String()
		if want != got {
			t.Errorf("%q: module mismatch; expected `%v`, got `%v`", g.path, want, got)
		}
	}
}

func BenchmarkParseFile(b *testing.B) {
	const path = "../testdata/sqlite/sqlite.ll"
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		b.Errorf("%q: unable to read file; %v", path, err)
		return
	}
	input := string(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := asm.ParseString(input)
		if err != nil {
			b.Errorf("%q: unable to parse file; %v", path, err)
			return
		}
	}
}
