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
		{path: "../testdata/loop.ll"},
		{path: "../testdata/frem.ll"},
		{path: "../testdata/binary.ll"},
		{path: "../testdata/bitwise.ll"},
		{path: "../testdata/fcmp.ll"},
		{path: "../testdata/select.ll"},
		{path: "../testdata/alloca.ll"},
		{path: "../testdata/getelementptr.ll"},
		{path: "../testdata/unreachable.ll"},
		{path: "../testdata/switch.ll"},
		{path: "../testdata/conversion.ll"},
		{path: "../testdata/bitcast.ll"},
		{path: "../testdata/addrspacecast.ll"},
		{path: "../testdata/extern.ll"},
		{path: "../testdata/const.ll"},
		{path: "../testdata/va_args.ll"},
		{path: "../testdata/array.ll"},
	}
	for _, g := range golden {
		buf, err := ioutil.ReadFile(g.path)
		if err != nil {
			t.Errorf("%q: unable to read file; %v", g.path, err)
			continue
		}
		want := string(buf)
		m, err := asm.ParseString(want)
		if err != nil {
			t.Errorf("%q: unable to parse file; %v", g.path, err)
			continue
		}
		got := m.String()
		if want != got {
			t.Errorf("%q: module mismatch; expected `%v`, got `%v`", g.path, want, got)
		}
	}
}
