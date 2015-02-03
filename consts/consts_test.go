package consts

import (
	"log"
	"testing"

	"github.com/mewlang/llvm/types"
)

var (
	i1Typ, i32Typ, i64Typ *types.Int
	f32Typ, f64Typ        *types.Float
)

func init() {
	var err error
	i1Typ, err = types.NewInt(1)
	if err != nil {
		log.Fatalln(err)
	}
	i32Typ, err = types.NewInt(32)
	if err != nil {
		log.Fatalln(err)
	}
	i64Typ, err = types.NewInt(64)
	if err != nil {
		log.Fatalln(err)
	}
	f32Typ, err = types.NewFloat(types.Float32)
	if err != nil {
		log.Fatalln(err)
	}
	f64Typ, err = types.NewFloat(types.Float64)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestIntString(t *testing.T) {
	golden := []struct {
		want  string
		input string
		typ   *types.Int
	}{
		{
			want:  "true",
			input: "true", typ: i1Typ,
		},
		{
			want:  "true",
			input: "1", typ: i1Typ,
		},
		{
			want:  "false",
			input: "false", typ: i1Typ,
		},
		{
			want:  "false",
			input: "0", typ: i1Typ,
		},
		{
			want:  "42",
			input: "42", typ: i32Typ,
		},
		{
			want:  "-137438953472",
			input: "-137438953472", typ: i64Typ,
		},
	}

	for i, g := range golden {
		v, err := NewInt(g.typ, g.input)
		if err != nil {
			t.Errorf("i=%d: %v", i, err)
			continue
		}
		got := v.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestFloatString(t *testing.T) {
	golden := []struct {
		input string
		typ   *types.Float
		want  string
		err   string
	}{
		// i=0
		{
			input: "3.14", typ: f32Typ,
			want: "3.14", err: `invalid floating point constant "3.14" for type "float"; precision loss`,
		},
		// i=0
		{
			input: "3.14", typ: f64Typ,
			want: "3.14",
		},
		// i=1
		{
			input: "-25000000000.0", typ: f64Typ,
			want: "-2.5e+10",
		},
		//{want: "3.14159265358979323846264338327950288419716939937510", input: "3.14159265358979323846264338327950288419716939937510"},
	}

	for i, g := range golden {
		v, err := NewFloat(g.typ, g.input)
		if !errEqual(err, g.err) {
			t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			continue
		} else if err != nil {
			// Expected error match, check next test case.
			continue
		}
		got := v.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

// errEqual returns true if err is represented by the string s, and false
// otherwise.
func errEqual(err error, s string) bool {
	t := ""
	if err != nil {
		t = err.Error()
	}
	return s == t
}
