package parser_test

import (
	"io/ioutil"
	"testing"

	"github.com/llir/llvm/asm"
)

func TestParseFile(t *testing.T) {
	// Round-trip test of the parser.
	golden := []struct {
		path string
	}{
		{path: "../testdata/rand.ll"},
		//{path: "../testdata/loop.ll"},
		{path: "../testdata/frem.ll"},
		{path: "../testdata/binary.ll"},
	}
	for i, g := range golden {
		buf, err := ioutil.ReadFile(g.path)
		if err != nil {
			t.Errorf("i=%d: unable to read file; %v", i, err)
			continue
		}
		want := string(buf)
		m, err := asm.ParseString(want)
		if err != nil {
			t.Errorf("i=%d: unable to parse file; %v", i, err)
			continue
		}
		got := m.String()
		if want != got {
			t.Errorf("i=%d: module mismatch; expected `%v`, got `%v`", i, want, got)
		}
	}
}
