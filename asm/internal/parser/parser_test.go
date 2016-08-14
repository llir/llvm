package parser_test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/spec/gocc/lexer"
	"github.com/llir/spec/gocc/parser"
	"github.com/mewkiz/pkg/errutil"
)

func TestParser(t *testing.T) {
	var golden = []struct {
		path string
		want string
	}{
		{
			path: "../../testdata/uc/noisy/advanced/8queens.ll",
			want: "../../testdata/uc/noisy/advanced/8queens.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/advanced/bubble.ll",
			want: "../../testdata/uc/noisy/advanced/bubble.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/advanced/eval.ll",
			want: "../../testdata/uc/noisy/advanced/eval.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/advanced/primes.ll",
			want: "../../testdata/uc/noisy/advanced/primes.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/advanced/quick.ll",
			want: "../../testdata/uc/noisy/advanced/quick.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/medium/circle.ll",
			want: "../../testdata/uc/noisy/medium/circle.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/medium/fac.ll",
			want: "../../testdata/uc/noisy/medium/fac.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/medium/fac-b.ll",
			want: "../../testdata/uc/noisy/medium/fac-b.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/medium/fib.ll",
			want: "../../testdata/uc/noisy/medium/fib.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/simple/sim01.ll",
			want: "../../testdata/uc/noisy/simple/sim01.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/simple/sim02.ll",
			want: "../../testdata/uc/noisy/simple/sim02.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/simple/sim03.ll",
			want: "../../testdata/uc/noisy/simple/sim03.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/simple/sim04.ll",
			want: "../../testdata/uc/noisy/simple/sim04.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/simple/sim05.ll",
			want: "../../testdata/uc/noisy/simple/sim05.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/simple/sim06.ll",
			want: "../../testdata/uc/noisy/simple/sim06.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/simple/sim07.ll",
			want: "../../testdata/uc/noisy/simple/sim07.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/simple/sim08.ll",
			want: "../../testdata/uc/noisy/simple/sim08.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/simple/sim09.ll",
			want: "../../testdata/uc/noisy/simple/sim09.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/simple/sim10.ll",
			want: "../../testdata/uc/noisy/simple/sim10.ll.golden",
		},
		{
			path: "../../testdata/uc/noisy/simple/sim11.ll",
			want: "../../testdata/uc/noisy/simple/sim11.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/lexer/l01.ll",
			want: "../../testdata/uc/quiet/lexer/l01.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/lexer/l02.ll",
			want: "../../testdata/uc/quiet/lexer/l02.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/lexer/l03.ll",
			want: "../../testdata/uc/quiet/lexer/l03.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/lexer/l04.ll",
			want: "../../testdata/uc/quiet/lexer/l04.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/lexer/l05.ll",
			want: "../../testdata/uc/quiet/lexer/l05.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/lexer/l06.ll",
			want: "../../testdata/uc/quiet/lexer/l06.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/mips/m01.ll",
			want: "../../testdata/uc/quiet/mips/m01.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/mips/m02.ll",
			want: "../../testdata/uc/quiet/mips/m02.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/mips/m03.ll",
			want: "../../testdata/uc/quiet/mips/m03.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/parser/p01.ll",
			want: "../../testdata/uc/quiet/parser/p01.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/parser/p02.ll",
			want: "../../testdata/uc/quiet/parser/p02.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/parser/p03.ll",
			want: "../../testdata/uc/quiet/parser/p03.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/parser/p04.ll",
			want: "../../testdata/uc/quiet/parser/p04.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/parser/p05.ll",
			want: "../../testdata/uc/quiet/parser/p05.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/parser/p06.ll",
			want: "../../testdata/uc/quiet/parser/p06.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/parser/p07.ll",
			want: "../../testdata/uc/quiet/parser/p07.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/parser/p08.ll",
			want: "../../testdata/uc/quiet/parser/p08.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/rtl/r01.ll",
			want: "../../testdata/uc/quiet/rtl/r01.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/rtl/r02.ll",
			want: "../../testdata/uc/quiet/rtl/r02.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/rtl/r03.ll",
			want: "../../testdata/uc/quiet/rtl/r03.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/rtl/r04.ll",
			want: "../../testdata/uc/quiet/rtl/r04.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/rtl/r05.ll",
			want: "../../testdata/uc/quiet/rtl/r05.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/rtl/r06.ll",
			want: "../../testdata/uc/quiet/rtl/r06.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/semantic/s01.ll",
			want: "../../testdata/uc/quiet/semantic/s01.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/semantic/s02.ll",
			want: "../../testdata/uc/quiet/semantic/s02.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/semantic/s03.ll",
			want: "../../testdata/uc/quiet/semantic/s03.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/semantic/s04.ll",
			want: "../../testdata/uc/quiet/semantic/s04.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/semantic/s05.ll",
			want: "../../testdata/uc/quiet/semantic/s05.ll.golden",
		},
		{
			path: "../../testdata/uc/quiet/semantic/s06.ll",
			want: "../../testdata/uc/quiet/semantic/s06.ll.golden",
		},
	}

	for _, g := range golden {
		log.Println("path:", g.path)
		s, err := lexer.NewLexerFile(g.path)
		if err != nil {
			t.Errorf("%q: error lexing file; %v", g.path, errutil.Err(err))
			continue
		}
		p := parser.NewParser()
		module, err := p.Parse(s)
		if err != nil {
			t.Errorf("%q: error parsing file; %v", g.path, errutil.Err(err))
			continue
		}
		m := module.(*ir.Module)
		buf, err := ioutil.ReadFile(g.want)
		if err != nil {
			t.Errorf("%q: error reading file; %v", g.path, errutil.Err(err))
			continue
		}
		got := m.String()
		want := string(buf)
		if got != want {
			t.Errorf("%q: module string representation mismatch; expected %v, got %v", g.path, want, got)
		}
	}
}
