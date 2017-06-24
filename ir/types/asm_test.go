package types_test

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
		{path: "../../asm/testdata/empty.ll"},
		// Top-level declarations.
		{path: "../../asm/testdata/module.ll"},
		{path: "../../asm/testdata/global.ll"},
		{path: "../../asm/testdata/func.ll"},
		{path: "../../asm/testdata/metadata.ll"},
		// Types.
		{path: "../../asm/testdata/type.ll"},
		// Constants.
		{path: "../../asm/testdata/const.ll"},
		// Constant expressions.
		{path: "../../asm/testdata/expr_binary.ll"},
		{path: "../../asm/testdata/expr_bitwise.ll"},
		{path: "../../asm/testdata/expr_vector.ll"},
		{path: "../../asm/testdata/expr_aggregate.ll"},
		{path: "../../asm/testdata/expr_memory.ll"},
		{path: "../../asm/testdata/expr_conversion.ll"},
		{path: "../../asm/testdata/expr_other.ll"},
		// Instructions.
		{path: "../../asm/testdata/inst_binary.ll"},
		{path: "../../asm/testdata/inst_bitwise.ll"},
		{path: "../../asm/testdata/inst_vector.ll"},
		{path: "../../asm/testdata/inst_aggregate.ll"},
		{path: "../../asm/testdata/inst_memory.ll"},
		{path: "../../asm/testdata/inst_conversion.ll"},
		{path: "../../asm/testdata/inst_other.ll"},
		// Terminators.
		{path: "../../asm/testdata/term.ll"},
		// Pseudo-random number generator.
		{path: "../../asm/testdata/rand.ll"},
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
