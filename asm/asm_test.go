package asm

import (
	"io/ioutil"
	"testing"

	"github.com/mewkiz/pkg/diffutil"
	"github.com/mewkiz/pkg/osutil"
)

func TestParseFile(t *testing.T) {
	golden := []struct {
		path string
	}{
		{path: "testdata/inst_binary.ll"},
		{path: "testdata/inst_bitwise.ll"},
		{path: "testdata/inst_vector.ll"},
		{path: "testdata/inst_aggregate.ll"},
		{path: "testdata/inst_memory.ll"},
		{path: "testdata/inst_conversion.ll"},
		{path: "testdata/inst_other.ll"},
		{path: "testdata/terminator.ll"},
		// LLVM Features.
		{path: "testdata/Feature/exception.ll"},
	}
	for _, g := range golden {
		m, err := ParseFile(g.path)
		if err != nil {
			t.Errorf("unable to parse %q into AST; %v", g.path, err)
			continue
		}
		path := g.path
		if osutil.Exists(g.path + ".golden") {
			path = g.path + ".golden"
		}
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("unable to read %q; %v", path, err)
			continue
		}
		want := string(buf)
		got := m.Def()
		if want != got {
			if err := diffutil.Diff(want, got, false); err != nil {
				panic(err)
			}
			t.Errorf("module mismatch; expected `%s`, got `%s`", want, got)
			continue
		}
	}
}
