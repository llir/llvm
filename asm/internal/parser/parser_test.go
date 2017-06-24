package parser_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/llir/llvm/asm"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestRoundTrip(t *testing.T) {
	// Round-trip test of the parser.
	golden := []struct {
		path string
	}{
		// Empty module.
		{path: "../../testdata/empty.ll"},
		// Top-level declarations.
		{path: "../../testdata/module.ll"},
		{path: "../../testdata/global.ll"},
		{path: "../../testdata/func.ll"},
		{path: "../../testdata/metadata.ll"},
		// Types.
		{path: "../../testdata/type.ll"},
		// Constants.
		{path: "../../testdata/const.ll"},
		// Constant expressions.
		{path: "../../testdata/expr_binary.ll"},
		{path: "../../testdata/expr_bitwise.ll"},
		{path: "../../testdata/expr_vector.ll"},
		{path: "../../testdata/expr_aggregate.ll"},
		{path: "../../testdata/expr_memory.ll"},
		{path: "../../testdata/expr_conversion.ll"},
		{path: "../../testdata/expr_other.ll"},
		// Instructions.
		{path: "../../testdata/inst_binary.ll"},
		{path: "../../testdata/inst_bitwise.ll"},
		{path: "../../testdata/inst_vector.ll"},
		{path: "../../testdata/inst_aggregate.ll"},
		{path: "../../testdata/inst_memory.ll"},
		{path: "../../testdata/inst_conversion.ll"},
		{path: "../../testdata/inst_other.ll"},
		// Terminators.
		{path: "../../testdata/term.ll"},
		// Pseudo-random number generator.
		{path: "../../testdata/rand.ll"},
	}
	dmp := diffmatchpatch.New()
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
			diffs := dmp.DiffMain(want, got, false)
			fmt.Println(dmp.DiffPrettyText(diffs))
			t.Errorf("%q: module mismatch; expected `%v`, got `%v`", g.path, want, got)
			continue
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
