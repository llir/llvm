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
		{path: "../testdata/hello.ll"},
		{path: "../testdata/empty.ll"},
		{path: "../testdata/struct.ll"},
		{path: "../testdata/recursive_types.ll"},
		{path: "../testdata/call_local_func.ll"},
		{path: "../testdata/ret.ll"},
		{path: "../testdata/gep_forward_reference.ll"},
		{path: "../testdata/const_struct.ll"},
		{path: "../testdata/float16.ll"},
		//{path: "../testdata/float128.ll"},
		//{path: "../testdata/hex_float.ll"},
		//{path: "../testdata/float_literals.ll"},
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
