package sem_test

import (
	"testing"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/sem"
)

func TestCheck(t *testing.T) {
	golden := []struct {
		path string
		errs []string
	}{
		{
			path: "testdata/type_int.ll",
			errs: []string{
				"invalid integer type bit width; expected > 0, got 0",
				"invalid integer type bit width; expected < 2^24, got 8388608",
			},
		},
	}
	for _, g := range golden {
		m, err := asm.ParseFile(g.path)
		if err != nil {
			t.Errorf("%q: unable to parse file; %v", g.path, err)
			continue
		}
		if err := sem.Check(m); err != nil {
			if g.errs == nil {
				t.Errorf("%q: unexpected semantic error; %v", g.path, err)
				continue
			}
			errs := err.(sem.ErrorList)
			if len(errs) != len(g.errs) {
				t.Errorf("%q: number of errors mismatch; expected %d, got %d", g.path, len(g.errs), len(errs))
				continue
			}
			for i := range g.errs {
				want, got := g.errs[i], errs[i].Error()
				if got != want {
					t.Errorf("%q: error mismatch; expected `%v`, got `%v`", g.path, want, got)
					continue
				}
			}
		} else if g.errs != nil {
			t.Errorf("%q: expected semantic error, got nil", g.path)
			continue
		}
	}
}
