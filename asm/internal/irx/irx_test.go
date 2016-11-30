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
	}
	for i, g := range golden {
		m, err := asm.ParseFile(g.path)
		if err != nil {
			t.Errorf("i=%d: unable to parse file; %v", i, err)
			continue
		}
		// Hack :) Rather than using reflect, pretty-print the module and search
		// for occurances of globalDummy, localDummy, instPhiDummy, incomingDummy,
		// instCallDummy, termBrDummy and termCondBrDummy.
		s := pretty.Sprint(m)
		if strings.Contains(s, "Dummy") {
			t.Errorf("i=%d: module contains dummy value; `%v`", i, s)
		}
	}
}
