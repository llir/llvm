package consts

import (
	"log"
	"testing"

	"github.com/mewlang/llvm/types"
)

var (
	// i1, i8, i32, i64
	i1Typ, i8Typ, i32Typ, i64Typ *types.Int
	// float, double
	f32Typ, f64Typ *types.Float
	// <2 x i32>
	i32x2VectorTyp *types.Vector
	// [2 x i32]
	i32x2ArrayTyp *types.Array
	// {i32, i8}
	i32i8StructTyp *types.Struct
	// i8 3
	i8Three Constant
	// i32 -13
	i32MinusThirteen Constant
	// i32 42
	i32FortyTwo Constant
)

func init() {
	var err error
	i1Typ, err = types.NewInt(1)
	if err != nil {
		log.Fatalln(err)
	}
	i8Typ, err = types.NewInt(8)
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
	i32x2VectorTyp, err = types.NewVector(i32Typ, 2)
	if err != nil {
		log.Fatalln(err)
	}
	i32x2ArrayTyp, err = types.NewArray(i32Typ, 2)
	if err != nil {
		log.Fatalln(err)
	}
	i32i8StructTyp, err = types.NewStruct([]types.Type{i32Typ, i8Typ}, false)
	if err != nil {
		log.Fatalln(err)
	}

	// i8 3
	i8Three, err = NewInt(i8Typ, "3")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 -13
	i32MinusThirteen, err = NewInt(i32Typ, "-13")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 42
	i32FortyTwo, err = NewInt(i32Typ, "42")
	if err != nil {
		log.Fatalln(err)
	}
}

func TestIntString(t *testing.T) {
	golden := []struct {
		input string
		typ   *types.Int
		want  string
		err   string
	}{
		{
			input: "true", typ: i1Typ,
			want: "i1 true",
		},
		{
			input: "1", typ: i1Typ,
			want: "i1 true",
		},
		{
			input: "false", typ: i1Typ,
			want: "i1 false",
		},
		{
			input: "0", typ: i1Typ,
			want: "i1 false",
		},
		{
			input: "2", typ: i1Typ,
			want: "", err: `invalid integer constant "2" for boolean type`,
		},
		{
			input: "true", typ: i32Typ,
			want: "", err: `integer constant "true" type mismatch; expected i1, got i32`,
		},
		{
			input: "false", typ: i32Typ,
			want: "", err: `integer constant "false" type mismatch; expected i1, got i32`,
		},
		{
			input: "42", typ: i32Typ,
			want: "i32 42",
		},
		{
			input: "-137438953472", typ: i64Typ,
			want: "i64 -137438953472",
		},
	}

	for i, g := range golden {
		v, err := NewInt(g.typ, g.input)
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
			want: "", err: `invalid floating point constant "3.14" for type "float"; precision loss`,
		},
		// i=1
		{
			input: "2.0", typ: f32Typ,
			want: "float 2.0",
		},
		// i=2
		{
			input: "3.14", typ: f64Typ,
			want: "double 3.14",
		},
		// i=3
		{
			input: "-25000000000.0", typ: f64Typ,
			want: "double -2.5e10",
		},
		// i=4
		{
			input: "3e14", typ: f64Typ,
			want: "double 3.0e14",
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

func TestVectorString(t *testing.T) {
	golden := []struct {
		elems []Constant
		typ   *types.Vector
		want  string
		err   string
	}{
		// i=0
		{
			elems: []Constant{i32FortyTwo, i32MinusThirteen}, typ: i32x2VectorTyp,
			want: "<2 x i32> <i32 42, i32 -13>",
		},
	}

	for i, g := range golden {
		v, err := NewVector(g.typ, g.elems)
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

func TestArrayString(t *testing.T) {
	golden := []struct {
		elems []Constant
		typ   *types.Array
		want  string
		err   string
	}{
		// i=0
		{
			elems: []Constant{i32MinusThirteen, i32FortyTwo}, typ: i32x2ArrayTyp,
			want: "[2 x i32] [i32 -13, i32 42]",
		},
	}

	for i, g := range golden {
		v, err := NewArray(g.typ, g.elems)
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

func TestStructString(t *testing.T) {
	golden := []struct {
		fields []Constant
		typ    *types.Struct
		want   string
		err    string
	}{
		// i=0
		{
			fields: []Constant{i32MinusThirteen, i8Three}, typ: i32i8StructTyp,
			want: "{i32, i8} {i32 -13, i8 3}",
		},
	}

	for i, g := range golden {
		v, err := NewStruct(g.typ, g.fields)
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
