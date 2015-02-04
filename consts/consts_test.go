package consts

import (
	"log"
	"strings"
	"testing"

	"github.com/mewlang/llvm/types"
)

var (
	// i1, i3, i5, i8, i32, i64
	i1Typ, i3Typ, i5Typ, i8Typ, i32Typ, i64Typ *types.Int
	// float, double
	f32Typ, f64Typ *types.Float
	// <2 x i32>
	i32x2VectorTyp *types.Vector
	// <2 x float>
	f32x2VectorTyp *types.Vector
	// [2 x i32]
	i32x2ArrayTyp *types.Array
	// {i32, i8}
	i32i8StructTyp *types.Struct
	// i1 1
	i1One Constant
	// i8 3
	i8Three Constant
	// i32 -13
	i32MinusThirteen Constant
	// i32 -4
	i32MinusFour Constant
	// i32 -3
	i32MinusThree Constant
	// i32 3
	i32Three Constant
	// i32 4
	i32Four Constant
	// i32 15
	i32Fifteen Constant
	// i32 42
	i32FortyTwo Constant
	// float -3.0
	f32MinusThree Constant
	// float -4.0
	f32MinusFour Constant
	// float 3.0
	f32Three Constant
	// float 4.0
	f32Four Constant
	// double 4.0
	f64Four Constant
	// <2 x float> <float 3.0, float 4.0>
	f32x2VectorThreeFour Constant
	// <2 x float> <float -3.0, float 4.0>
	f32x2VectorMinusThreeFour Constant
	// <2 x i32> <i32 3, i32 42>
	i32x2VectorThreeFortyTwo Constant
	// <2 x i32> <i32 -3, i32 15>
	i32x2VectorMinusThreeFifteen Constant
)

func init() {
	var err error
	i1Typ, err = types.NewInt(1)
	if err != nil {
		log.Fatalln(err)
	}
	i3Typ, err = types.NewInt(3)
	if err != nil {
		log.Fatalln(err)
	}
	i5Typ, err = types.NewInt(5)
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
	f32x2VectorTyp, err = types.NewVector(f32Typ, 2)
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

	// i1 1
	i1One, err = NewInt(i1Typ, "1")
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
	// i32 -4
	i32MinusFour, err = NewInt(i32Typ, "-4")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 -3
	i32MinusThree, err = NewInt(i32Typ, "-3")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 3
	i32Three, err = NewInt(i32Typ, "3")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 4
	i32Four, err = NewInt(i32Typ, "4")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 15
	i32Fifteen, err = NewInt(i32Typ, "15")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 42
	i32FortyTwo, err = NewInt(i32Typ, "42")
	if err != nil {
		log.Fatalln(err)
	}
	// float -3.0
	f32MinusThree, err = NewFloat(f32Typ, "-3.0")
	if err != nil {
		log.Fatalln(err)
	}
	// float -4.0
	f32MinusFour, err = NewFloat(f32Typ, "-4.0")
	if err != nil {
		log.Fatalln(err)
	}
	// float 3.0
	f32Three, err = NewFloat(f32Typ, "3.0")
	if err != nil {
		log.Fatalln(err)
	}
	// float 4.0
	f32Four, err = NewFloat(f32Typ, "4.0")
	if err != nil {
		log.Fatalln(err)
	}
	// double 4.0
	f64Four, err = NewFloat(f64Typ, "4.0")
	if err != nil {
		log.Fatalln(err)
	}
	// <2 x float> <float 3.0, float 4.0>
	f32x2VectorThreeFour, err = NewVector(f32x2VectorTyp, []Constant{f32Three, f32Four})
	if err != nil {
		log.Fatalln(err)
	}
	// <2 x float> <float -3.0, float 4.0>
	f32x2VectorMinusThreeFour, err = NewVector(f32x2VectorTyp, []Constant{f32MinusThree, f32Four})
	if err != nil {
		log.Fatalln(err)
	}
	// <2 x i32> <i32 3, i32 42>
	i32x2VectorThreeFortyTwo, err = NewVector(i32x2VectorTyp, []Constant{i32Three, i32FortyTwo})
	if err != nil {
		log.Fatalln(err)
	}
	// <2 x i32> <i32 -3, i32 15>
	i32x2VectorMinusThreeFifteen, err = NewVector(i32x2VectorTyp, []Constant{i32MinusThree, i32Fifteen})
	if err != nil {
		log.Fatalln(err)
	}
}

func TestIntString(t *testing.T) {
	golden := []struct {
		input string
		typ   types.Type
		want  string
		err   string
	}{
		// i=0
		{
			input: "true", typ: i1Typ,
			want: "i1 true",
		},
		// i=1
		{
			input: "1", typ: i1Typ,
			want: "i1 true",
		},
		// i=2
		{
			input: "false", typ: i1Typ,
			want: "i1 false",
		},
		// i=3
		{
			input: "0", typ: i1Typ,
			want: "i1 false",
		},
		// i=4
		{
			input: "2", typ: i1Typ,
			want: "", err: `invalid integer constant "2" for boolean type`,
		},
		// i=5
		{
			input: "true", typ: i32Typ,
			want: "", err: `integer constant "true" type mismatch; expected i1, got i32`,
		},
		// i=6
		{
			input: "false", typ: i32Typ,
			want: "", err: `integer constant "false" type mismatch; expected i1, got i32`,
		},
		// i=7
		{
			input: "42", typ: i32Typ,
			want: "i32 42",
		},
		// i=8
		{
			input: "-137438953472", typ: i64Typ,
			want: "i64 -137438953472",
		},
		// i=9
		{
			input: "3.0", typ: f32Typ,
			want: "", err: `invalid type "float" for integer constant`,
		},
		// i=10
		{
			input: "foo", typ: i64Typ,
			want: "", err: `unable to parse integer constant "foo"`,
		},
	}

	for i, g := range golden {
		v, err := NewInt(g.typ, g.input)
		if !sameError(err, g.err) {
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
		typ   types.Type
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
		// i=5
		{
			input: "12", typ: i32Typ,
			want: "", err: `invalid type "i32" for floating point constant`,
		},
		// i=6
		{
			input: "foo", typ: f32Typ,
			want: "", err: `unable to parse floating point constant "foo"`,
		},
		//{want: "3.14159265358979323846264338327950288419716939937510", input: "3.14159265358979323846264338327950288419716939937510"},
	}

	for i, g := range golden {
		v, err := NewFloat(g.typ, g.input)
		if !sameError(err, g.err) {
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
		if !sameError(err, g.err) {
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
		if !sameError(err, g.err) {
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
		if !sameError(err, g.err) {
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

func TestIntTruncString(t *testing.T) {
	golden := []struct {
		orig Constant
		to   *types.Int
		want string
		err  string
	}{
		// i=0
		{
			orig: i32Fifteen, to: i3Typ,
			want: "i3 trunc(i32 15 to i3)",
		},
	}

	for i, g := range golden {
		v, err := NewIntTrunc(g.orig, g.to)
		if !sameError(err, g.err) {
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

func TestIntZeroExtString(t *testing.T) {
	golden := []struct {
		orig Constant
		to   *types.Int
		want string
		err  string
	}{
		// i=0
		{
			orig: i1One, to: i5Typ,
			want: "i5 zext(i1 true to i5)",
		},
	}

	for i, g := range golden {
		v, err := NewIntZeroExt(g.orig, g.to)
		if !sameError(err, g.err) {
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

func TestIntSignExtString(t *testing.T) {
	golden := []struct {
		orig Constant
		to   *types.Int
		want string
		err  string
	}{
		// i=0
		{
			orig: i1One, to: i5Typ,
			want: "i5 sext(i1 true to i5)",
		},
	}

	for i, g := range golden {
		v, err := NewIntSignExt(g.orig, g.to)
		if !sameError(err, g.err) {
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

func TestFloatTruncString(t *testing.T) {
	golden := []struct {
		orig Constant
		to   *types.Float
		want string
		err  string
	}{
		// i=0
		{
			orig: f64Four, to: f32Typ,
			want: "float fptrunc(double 4.0 to float)",
		},
	}

	for i, g := range golden {
		v, err := NewFloatTrunc(g.orig, g.to)
		if !sameError(err, g.err) {
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

func TestFloatExtString(t *testing.T) {
	golden := []struct {
		orig Constant
		to   *types.Float
		want string
		err  string
	}{
		// i=0
		{
			orig: f32Four, to: f64Typ,
			want: "double fpext(float 4.0 to double)",
		},
	}

	for i, g := range golden {
		v, err := NewFloatExt(g.orig, g.to)
		if !sameError(err, g.err) {
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

func TestFloatToUintString(t *testing.T) {
	golden := []struct {
		orig Constant
		to   types.Type
		want string
		err  string
	}{
		// i=0
		{
			orig: f32Four, to: i32Typ,
			want: "i32 fptoui(float 4.0 to i32)",
		},
		// i=1
		{
			orig: f32x2VectorThreeFour, to: i32x2VectorTyp,
			want: "<2 x i32> fptoui(<2 x float> <float 3.0, float 4.0> to <2 x i32>)",
		},
	}

	for i, g := range golden {
		v, err := NewFloatToUint(g.orig, g.to)
		if !sameError(err, g.err) {
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

func TestFloatToIntString(t *testing.T) {
	golden := []struct {
		orig Constant
		to   types.Type
		want string
		err  string
	}{
		// i=0
		{
			orig: f32MinusFour, to: i32Typ,
			want: "i32 fptosi(float -4.0 to i32)",
		},
		// i=1
		{
			orig: f32x2VectorMinusThreeFour, to: i32x2VectorTyp,
			want: "<2 x i32> fptosi(<2 x float> <float -3.0, float 4.0> to <2 x i32>)",
		},
	}

	for i, g := range golden {
		v, err := NewFloatToInt(g.orig, g.to)
		if !sameError(err, g.err) {
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

func TestUintToFloatString(t *testing.T) {
	golden := []struct {
		orig Constant
		to   types.Type
		want string
		err  string
	}{
		// i=0
		{
			orig: i32Four, to: f32Typ,
			want: "float uitofp(i32 4 to float)",
		},
		// i=1
		{
			orig: i32x2VectorThreeFortyTwo, to: f32x2VectorTyp,
			want: "<2 x float> uitofp(<2 x i32> <i32 3, i32 42> to <2 x float>)",
		},
	}

	for i, g := range golden {
		v, err := NewUintToFloat(g.orig, g.to)
		if !sameError(err, g.err) {
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

func TestIntToFloatString(t *testing.T) {
	golden := []struct {
		orig Constant
		to   types.Type
		want string
		err  string
	}{
		// i=0
		{
			orig: i32MinusFour, to: f32Typ,
			want: "float sitofp(i32 -4 to float)",
		},
		// i=1
		{
			orig: i32x2VectorMinusThreeFifteen, to: f32x2VectorTyp,
			want: "<2 x float> sitofp(<2 x i32> <i32 -3, i32 15> to <2 x float>)",
		},
	}

	for i, g := range golden {
		v, err := NewIntToFloat(g.orig, g.to)
		if !sameError(err, g.err) {
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

// sameError returns true if err is represented by the string s, and false
// otherwise. Some error messages constants suffixes from external functions,
// e.g. the strconv error in:
//
//    unable to parse integer constant "foo"; strconv.ParseInt: parsing "foo": invalid syntax`
//
// To avoid depending on the error of external functions, s matches the error if
// it is a non-empty prefix of err.
func sameError(err error, s string) bool {
	t := ""
	if err != nil {
		if len(s) == 0 {
			return false
		}
		t = err.Error()
	}
	return strings.HasPrefix(t, s)
}
