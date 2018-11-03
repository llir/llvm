package asm

import (
	"io/ioutil"
	"testing"
)

func TestParseFile(t *testing.T) {
	golden := []struct {
		path string
	}{
		{path: "testdata/inst_binary.ll"},
		{path: "testdata/inst_bitwise.ll"},
		{path: "testdata/inst_vector.ll"},
		{path: "testdata/inst_aggregate.ll"},
		//{path: "testdata/inst_memory.ll"},
		{path: "testdata/inst_conversion.ll"},
		//{path: "testdata/inst_other.ll"},
	}
	for _, g := range golden {
		m, err := ParseFile(g.path)
		buf, err := ioutil.ReadFile(g.path)
		if err != nil {
			t.Errorf("unable to read %q; %v", g.path, err)
		}
		want := string(buf)
		if err != nil {
			t.Errorf("unable to parse %q into AST; %v", g.path, err)
			continue
		}
		got := m.Def()
		if want != got {
			t.Errorf("module mismatch; expected `%s`, got `%s`", want, got)
			continue
		}
	}
}
