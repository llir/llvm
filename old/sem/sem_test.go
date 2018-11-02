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
		// Global variables.
		{
			path: "testdata/global.ll",
			//errs: []string{
			//	"invalid global variable content type; expected single value or aggregate type, got *types.LabelType",
			//	"invalid global variable content type; expected single value or aggregate type, got *types.MetadataType",
			//},
		},

		// Types.
		{
			path: "testdata/type_func.ll",
			errs: []string{
				"invalid function return type; expected void, single value or aggregate type, got *types.FuncType",
				"invalid function return type; expected void, single value or aggregate type, got *types.LabelType",
				"invalid function return type; expected void, single value or aggregate type, got *types.MetadataType",
			},
		},
		{
			path: "testdata/type_int.ll",
			errs: []string{
				"invalid integer type bit width; expected > 0, got 0",
				"invalid integer type bit width; expected < 2^24, got 8388608",
			},
		},
		{
			path: "testdata/type_pointer.ll",
			errs: []string{
				"invalid pointer element type; expected function, single value or aggregate type, got *types.VoidType",
				"invalid pointer element type; expected function, single value or aggregate type, got *types.LabelType",
				"invalid pointer element type; expected function, single value or aggregate type, got *types.MetadataType",
			},
		},
		{
			path: "testdata/type_vector.ll",
			errs: []string{
				"invalid vector element type; expected integer, floating-point or pointer type, got *types.VectorType",
				"invalid vector element type; expected integer, floating-point or pointer type, got *types.LabelType",
				//"invalid vector element type; expected integer, floating-point or pointer type, got *types.MetadataType",
				"invalid vector element type; expected integer, floating-point or pointer type, got *types.ArrayType",
				"invalid vector element type; expected integer, floating-point or pointer type, got *types.StructType",
				"invalid vector element type; expected integer, floating-point or pointer type, got *types.StructType",
			},
		},
		{
			path: "testdata/type_array.ll",
			errs: []string{
				"invalid array element type; expected single value or aggregate type, got *types.LabelType",
				//"invalid array element type; expected single value or aggregate type, got *types.MetadataType",
			},
		},
		{
			path: "testdata/type_struct.ll",
			errs: []string{
				"invalid struct field type; expected single value or aggregate type, got *types.LabelType",
				//"invalid struct field type; expected single value or aggregate type, got *types.MetadataType",
			},
		},

		// Constants.
		{
			path: "testdata/const_vector.ll",
			errs: []string{
				"vector element type `i32` and element type `i8` mismatch",
			},
		},
		{
			path: "testdata/const_array.ll",
			errs: []string{
				"array element type `i32` and element type `i8` mismatch",
			},
		},
		{
			path: "testdata/const_struct.ll",
			errs: nil,
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
				t.Errorf("want:")
				for _, err := range g.errs {
					t.Errorf("\t%s", err)
				}
				t.Errorf("got:")
				for _, err := range errs {
					t.Errorf("\t%s", err)
				}
				continue
			}
			for i := range g.errs {
				want, got := g.errs[i], errs[i].Error()
				if got != want {
					t.Errorf("%q: error mismatch; expected `%v`, got `%v`", g.path, want, got)
				}
			}
		} else if g.errs != nil {
			t.Errorf("%q: expected semantic error, got nil", g.path)
			continue
		}
	}
}
