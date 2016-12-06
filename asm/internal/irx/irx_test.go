package irx_test

import (
	"strings"
	"testing"

	"github.com/kr/pretty"
	"github.com/llir/llvm/asm"
)

func TestFix(t *testing.T) {
	// Verify that the fixer has replaced all dummy values.
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
	}
	for _, g := range golden {
		m, err := asm.ParseFile(g.path)
		if err != nil {
			t.Errorf("%q: unable to parse file; %v", g.path, err)
			continue
		}
		// Hack :) Rather than using reflect, pretty-print the module and search
		// for occurances of globalDummy, localDummy, instPhiDummy, incomingDummy,
		// instCallDummy, termBrDummy and termCondBrDummy.
		s := pretty.Sprint(m)
		if strings.Contains(s, "Dummy") {
			t.Errorf("%q: module contains dummy value; `%v`", g.path, s)
		}
	}
}
