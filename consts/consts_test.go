package consts_test

import (
	"log"
	"strings"
	"testing"

	"github.com/mewlang/llvm/consts"
	"github.com/mewlang/llvm/types"
)

var (
	// i1, i3, i5, i8, i32, i64
	i1Typ, i3Typ, i5Typ, i8Typ, i32Typ, i64Typ *types.Int
	// float, double, f128, ppc_f128
	f32Typ, f64Typ, f128Typ, f128_ppcTyp *types.Float
	// <2 x i32>
	i32x2VecTyp *types.Vector
	// <3 x i32>
	i32x3VecTyp *types.Vector
	// <2 x float>
	f32x2VecTyp *types.Vector
	// <3 x float>
	f32x3VecTyp *types.Vector
	// [2 x i32]
	i32x2ArrTyp *types.Array
	// {i32, i8}
	i32i8StructTyp *types.Struct
	// [2 x {i32, i8}]
	i32i8x2ArrTyp *types.Array
	// i1 1
	i1One consts.Constant
	// i8 3
	i8Three consts.Constant
	// i8 4
	i8Four consts.Constant
	// i32 -13
	i32MinusThirteen consts.Constant
	// i32 -4
	i32MinusFour consts.Constant
	// i32 -3
	i32MinusThree consts.Constant
	// i32 1
	i32One consts.Constant
	// i32 2
	i32Two consts.Constant
	// i32 3
	i32Three consts.Constant
	// i32 4
	i32Four consts.Constant
	// i32 15
	i32Fifteen consts.Constant
	// i32 42
	i32FortyTwo consts.Constant
	// float -3.0
	f32MinusThree consts.Constant
	// float -4.0
	f32MinusFour consts.Constant
	// float 1.0
	f32One consts.Constant
	// float 2.0
	f32Two consts.Constant
	// float 3.0
	f32Three consts.Constant
	// float 4.0
	f32Four consts.Constant
	// double 4.0
	f64Four consts.Constant
	// <3 x i32> <i32 1, i32 2, i32 3>
	i32x3OneTwoThree consts.Constant
	// <2 x i32> <i32 3, i32 42>
	i32x2VecThreeFortyTwo consts.Constant
	// <2 x i32> <i32 -3, i32 15>
	i32x2VecMinusThreeFifteen consts.Constant
	// <2 x float> <float 3.0, float 4.0>
	f32x2VecThreeFour consts.Constant
	// <2 x float> <float -3.0, float 4.0>
	f32x2VecMinusThreeFour consts.Constant
	// <3 x float> <float 3.0, float 2.0, float 1.0>
	f32x3VecThreeFourFifteen consts.Constant
	// {i32, i8} {i32 4, i8 3}
	i32i8FourThree consts.Constant
	// {i32, i8} {i32 3, i8 4}
	i32i8ThreeFour consts.Constant
	// TODO: Uncomment when fp128 and ppc_fp128 are supported.
	/*
		// fp128 3.0
		f128Three consts.Constant
		// ppc_fp128 4.0
		f128_ppcFour consts.Constant
	*/
)

func init() {
	// i1
	var err error
	i1Typ, err = types.NewInt(1)
	if err != nil {
		log.Fatalln(err)
	}
	// i3
	i3Typ, err = types.NewInt(3)
	if err != nil {
		log.Fatalln(err)
	}
	// i5
	i5Typ, err = types.NewInt(5)
	if err != nil {
		log.Fatalln(err)
	}
	// i8
	i8Typ, err = types.NewInt(8)
	if err != nil {
		log.Fatalln(err)
	}
	// i32
	i32Typ, err = types.NewInt(32)
	if err != nil {
		log.Fatalln(err)
	}
	// i64
	i64Typ, err = types.NewInt(64)
	if err != nil {
		log.Fatalln(err)
	}
	// float
	f32Typ, err = types.NewFloat(types.Float32)
	if err != nil {
		log.Fatalln(err)
	}
	// double
	f64Typ, err = types.NewFloat(types.Float64)
	if err != nil {
		log.Fatalln(err)
	}
	// fp128
	f128Typ, err = types.NewFloat(types.Float128)
	if err != nil {
		log.Fatalln(err)
	}
	// ppc_fp128
	f128_ppcTyp, err = types.NewFloat(types.Float128_PPC)
	if err != nil {
		log.Fatalln(err)
	}
	// <2 x i32>
	i32x2VecTyp, err = types.NewVector(i32Typ, 2)
	if err != nil {
		log.Fatalln(err)
	}
	// <3 x i32>
	i32x3VecTyp, err = types.NewVector(i32Typ, 3)
	if err != nil {
		log.Fatalln(err)
	}
	// <2 x float>
	f32x2VecTyp, err = types.NewVector(f32Typ, 2)
	if err != nil {
		log.Fatalln(err)
	}
	// <3 x float>
	f32x3VecTyp, err = types.NewVector(f32Typ, 3)
	if err != nil {
		log.Fatalln(err)
	}
	// [2 x i32]
	i32x2ArrTyp, err = types.NewArray(i32Typ, 2)
	if err != nil {
		log.Fatalln(err)
	}
	// {i32, i8}
	i32i8StructTyp, err = types.NewStruct([]types.Type{i32Typ, i8Typ}, false)
	if err != nil {
		log.Fatalln(err)
	}
	// [2 x {i32, i8}]
	i32i8x2ArrTyp, err = types.NewArray(i32i8StructTyp, 2)
	if err != nil {
		log.Fatalln(err)
	}

	// i1 1
	i1One, err = consts.NewInt(i1Typ, "1")
	if err != nil {
		log.Fatalln(err)
	}
	// i8 3
	i8Three, err = consts.NewInt(i8Typ, "3")
	if err != nil {
		log.Fatalln(err)
	}
	// i8 4
	i8Four, err = consts.NewInt(i8Typ, "4")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 -13
	i32MinusThirteen, err = consts.NewInt(i32Typ, "-13")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 -4
	i32MinusFour, err = consts.NewInt(i32Typ, "-4")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 -3
	i32MinusThree, err = consts.NewInt(i32Typ, "-3")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 1
	i32One, err = consts.NewInt(i32Typ, "1")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 2
	i32Two, err = consts.NewInt(i32Typ, "2")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 3
	i32Three, err = consts.NewInt(i32Typ, "3")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 4
	i32Four, err = consts.NewInt(i32Typ, "4")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 15
	i32Fifteen, err = consts.NewInt(i32Typ, "15")
	if err != nil {
		log.Fatalln(err)
	}
	// i32 42
	i32FortyTwo, err = consts.NewInt(i32Typ, "42")
	if err != nil {
		log.Fatalln(err)
	}
	// float -3.0
	f32MinusThree, err = consts.NewFloat(f32Typ, "-3.0")
	if err != nil {
		log.Fatalln(err)
	}
	// float -4.0
	f32MinusFour, err = consts.NewFloat(f32Typ, "-4.0")
	if err != nil {
		log.Fatalln(err)
	}
	// float 1.0
	f32One, err = consts.NewFloat(f32Typ, "1.0")
	if err != nil {
		log.Fatalln(err)
	}
	// float 2.0
	f32Two, err = consts.NewFloat(f32Typ, "2.0")
	if err != nil {
		log.Fatalln(err)
	}
	// float 3.0
	f32Three, err = consts.NewFloat(f32Typ, "3.0")
	if err != nil {
		log.Fatalln(err)
	}
	// float 4.0
	f32Four, err = consts.NewFloat(f32Typ, "4.0")
	if err != nil {
		log.Fatalln(err)
	}
	// double 4.0
	f64Four, err = consts.NewFloat(f64Typ, "4.0")
	if err != nil {
		log.Fatalln(err)
	}
	// TODO: Uncomment when fp128 and ppc_fp128 are supported.
	/*
		// fp128 3.0
		f128Three, err = consts.NewFloat(f128Typ, "3.0")
		if err != nil {
			log.Fatalln(err)
		}
		// ppc_fp128 4.0
		f128_ppcFour, err = consts.NewFloat(f128_ppcTyp, "4.0")
		if err != nil {
			log.Fatalln(err)
		}
	*/
	// <3 x i32> <i32 1, i32 2, i32 3>
	i32x3OneTwoThree, err = consts.NewVector(i32x3VecTyp, []consts.Constant{i32One, i32Two, i32Three})
	if err != nil {
		log.Fatalln(err)
	}
	// <2 x i32> <i32 3, i32 42>
	i32x2VecThreeFortyTwo, err = consts.NewVector(i32x2VecTyp, []consts.Constant{i32Three, i32FortyTwo})
	if err != nil {
		log.Fatalln(err)
	}
	// <2 x i32> <i32 -3, i32 15>
	i32x2VecMinusThreeFifteen, err = consts.NewVector(i32x2VecTyp, []consts.Constant{i32MinusThree, i32Fifteen})
	if err != nil {
		log.Fatalln(err)
	}
	// <2 x float> <float 3.0, float 4.0>
	f32x2VecThreeFour, err = consts.NewVector(f32x2VecTyp, []consts.Constant{f32Three, f32Four})
	if err != nil {
		log.Fatalln(err)
	}
	// <2 x float> <float -3.0, float 4.0>
	f32x2VecMinusThreeFour, err = consts.NewVector(f32x2VecTyp, []consts.Constant{f32MinusThree, f32Four})
	if err != nil {
		log.Fatalln(err)
	}
	// <3 x float> <float 3.0, float 2.0, float 1.0>
	f32x3VecThreeFourFifteen, err = consts.NewVector(f32x3VecTyp, []consts.Constant{f32Three, f32Two, f32One})
	if err != nil {
		log.Fatalln(err)
	}
	// {i32, i8} {i32 4, i8 3}
	i32i8FourThree, err = consts.NewStruct(i32i8StructTyp, []consts.Constant{i32Four, i8Three})
	if err != nil {
		log.Fatalln(err)
	}
	// {i32, i8} {i32 3, i8 4}
	i32i8ThreeFour, err = consts.NewStruct(i32i8StructTyp, []consts.Constant{i32Three, i8Four})
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
		v, err := consts.NewInt(g.typ, g.input)
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
		v, err := consts.NewFloat(g.typ, g.input)
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
		elems []consts.Constant
		typ   types.Type
		want  string
		err   string
	}{
		// i=0
		{
			elems: []consts.Constant{i32FortyTwo, i32MinusThirteen}, typ: i32x2VecTyp,
			want: "<2 x i32> <i32 42, i32 -13>",
		},
		// i=1
		{
			elems: nil, typ: f64Typ,
			want: "", err: `invalid type "double" for vector constant`,
		},
		// i=2
		{
			elems: []consts.Constant{f32Three, f32Four}, typ: f32x2VecTyp,
			want: "<2 x float> <float 3.0, float 4.0>",
		},
		// i=3
		{
			elems: []consts.Constant{f32Three, i32Four}, typ: f32x2VecTyp,
			want: "", err: `invalid vector element type; expected "float", got "i32"`,
		},
	}

	for i, g := range golden {
		v, err := consts.NewVector(g.typ, g.elems)
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
		elems []consts.Constant
		typ   types.Type
		want  string
		err   string
	}{
		// i=0
		{
			elems: []consts.Constant{i32MinusThirteen, i32FortyTwo}, typ: i32x2ArrTyp,
			want: "[2 x i32] [i32 -13, i32 42]",
		},
		// i=1
		{
			elems: nil, typ: i32x2VecTyp,
			want: "", err: `invalid type "<2 x i32>" for array constant`,
		},
		// i=2
		{
			elems: []consts.Constant{i32i8FourThree, i32i8ThreeFour}, typ: i32i8x2ArrTyp,
			want: "[2 x {i32, i8}] [{i32, i8} {i32 4, i8 3}, {i32, i8} {i32 3, i8 4}]",
		},
		// i=3
		{
			elems: []consts.Constant{i32i8FourThree, i32Four}, typ: i32i8x2ArrTyp,
			want: "", err: `invalid array element type; expected "{i32, i8}", got "i32"`,
		},
	}

	for i, g := range golden {
		v, err := consts.NewArray(g.typ, g.elems)
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
		fields []consts.Constant
		typ    types.Type
		want   string
		err    string
	}{
		// i=0
		{
			fields: []consts.Constant{i32MinusThirteen, i8Three}, typ: i32i8StructTyp,
			want: "{i32, i8} {i32 -13, i8 3}",
		},
		// i=1
		{
			fields: nil, typ: i32x2VecTyp,
			want: "", err: `invalid type "<2 x i32>" for structure constant`,
		},
		// i=2
		{
			fields: []consts.Constant{i32Three, i32Fifteen, i8Three}, typ: i32i8StructTyp,
			want: "", err: "incorrect number of fields in structure constant; expected 2, got 3",
		},
		// i=2
		{
			fields: []consts.Constant{i32Four, i32Three}, typ: i32i8StructTyp,
			want: "", err: `invalid structure field (1) type; expected "i8", got "i32"`,
		},
	}

	for i, g := range golden {
		v, err := consts.NewStruct(g.typ, g.fields)
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
		orig consts.Constant
		to   types.Type
		want string
		err  string
	}{
		// i=0
		{
			orig: i32Fifteen, to: i3Typ,
			want: "i3 trunc(i32 15 to i3)",
		},
		// i=1
		{
			orig: f32Four, to: i3Typ,
			want: "", err: `invalid integer truncation; expected integer constant for orig, got "float"`,
		},
		// i=2
		{
			orig: i32Four, to: f64Typ,
			want: "", err: `invalid integer truncation; expected integer target type, got "double"`,
		},
		// i=3
		{
			orig: i32Four, to: i64Typ,
			want: "", err: `invalid integer truncation; target size (64) larger than original size (32)`,
		},
	}

	for i, g := range golden {
		v, err := consts.NewIntTrunc(g.orig, g.to)
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
		orig consts.Constant
		to   types.Type
		want string
		err  string
	}{
		// i=0
		{
			orig: i1One, to: i5Typ,
			want: "i5 zext(i1 true to i5)",
		},
		// i=1
		{
			orig: f64Four, to: i3Typ,
			want: "", err: `invalid integer zero extension; expected integer constant for orig, got "double"`,
		},
		// i=2
		{
			orig: i32Four, to: i32x2VecTyp,
			want: "", err: `invalid integer zero extension; expected integer target type, got "<2 x i32>"`,
		},
		// i=3
		{
			orig: i32Four, to: i8Typ,
			want: "", err: `invalid integer zero extension; target size (8) smaller than original size (32)`,
		},
	}

	for i, g := range golden {
		v, err := consts.NewIntZeroExt(g.orig, g.to)
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
		orig consts.Constant
		to   types.Type
		want string
		err  string
	}{
		// i=0
		{
			orig: i1One, to: i5Typ,
			want: "i5 sext(i1 true to i5)",
		},
		// i=1
		{
			orig: i32i8FourThree, to: i32Typ,
			want: "", err: `invalid integer sign extension; expected integer constant for orig, got "{i32, i8}"`,
		},
		// i=2
		{
			orig: i32Four, to: i32i8x2ArrTyp,
			want: "", err: `invalid integer sign extension; expected integer target type, got "[2 x {i32, i8}]"`,
		},
		// i=3
		{
			orig: i32Four, to: i3Typ,
			want: "", err: `invalid integer sign extension; target size (3) smaller than original size (32)`,
		},
	}

	for i, g := range golden {
		v, err := consts.NewIntSignExt(g.orig, g.to)
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
		orig consts.Constant
		to   types.Type
		want string
		err  string
	}{
		// i=0
		{
			orig: f64Four, to: f32Typ,
			want: "float fptrunc(double 4.0 to float)",
		},
		// i=1
		{
			orig: i32Four, to: f32Typ,
			want: "", err: `invalid floating point truncation; expected floating point constant for orig, got "i32"`,
		},
		// i=2
		{
			orig: f32Three, to: i32Typ,
			want: "", err: `invalid floating point truncation; expected floating point target type, got "i32"`,
		},
		// i=3
		{
			orig: f32Three, to: f64Typ,
			want: "", err: `invalid floating point truncation; target size (64) larger than original size (32)`,
		},
		// TODO: Uncomment when fp128 and ppc_fp128 are supported.
		/*
			// i=4
			{
				orig: f128Three, to: f128_ppcTyp,
				want: "", err: `invalid floating point truncation; cannot convert from "fp128" to "ppc_fp128"`,
			},
			// i=5
			{
				orig: f128_ppcFour, to: f128Typ,
				want: "", err: `invalid floating point truncation; cannot convert from "ppc_fp128" to "fp128"`,
			},
		*/
	}

	for i, g := range golden {
		v, err := consts.NewFloatTrunc(g.orig, g.to)
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
		orig consts.Constant
		to   types.Type
		want string
		err  string
	}{
		// i=0
		{
			orig: f32Four, to: f64Typ,
			want: "double fpext(float 4.0 to double)",
		},
		// i=1
		{
			orig: i8Three, to: f32Typ,
			want: "", err: `invalid floating point extension; expected floating point constant for orig, got "i8"`,
		},
		// i=2
		{
			orig: f32Three, to: i64Typ,
			want: "", err: `invalid floating point extension; expected floating point target type, got "i64"`,
		},
		// i=3
		{
			orig: f64Four, to: f32Typ,
			want: "", err: `invalid floating point extension; target size (32) smaller than original size (64)`,
		},
		// TODO: Uncomment when fp128 and ppc_fp128 are supported.
		/*
			// i=4
			{
				orig: f128Three, to: f128_ppcTyp,
				want: "", err: `invalid floating point extension; cannot convert from "fp128" to "ppc_fp128"`,
			},
			// i=5
			{
				orig: f128_ppcFour, to: f128Typ,
				want: "", err: `invalid floating point extension; cannot convert from "ppc_fp128" to "fp128"`,
			},
		*/
	}

	for i, g := range golden {
		v, err := consts.NewFloatExt(g.orig, g.to)
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
		orig consts.Constant
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
			orig: f32x2VecThreeFour, to: i32x2VecTyp,
			want: "<2 x i32> fptoui(<2 x float> <float 3.0, float 4.0> to <2 x i32>)",
		},
		// i=2
		{
			orig: i32x2VecMinusThreeFifteen, to: i32x2VecTyp,
			want: "", err: `invalid floating point to unsigned integer conversion; expected floating point constant (or constant vector) for orig, got "<2 x i32>"`,
		},
		// i=3
		{
			orig: f32x2VecThreeFour, to: f32x2VecTyp,
			want: "", err: `invalid floating point to unsigned integer conversion; expected integer (or integer vector) target type, got "<2 x float>"`,
		},
		// i=4
		{
			orig: f32x2VecThreeFour, to: i32Typ,
			want: "", err: `invalid floating point to unsigned integer conversion; cannot convert from "<2 x float>" to "i32"`,
		},
		// i=5
		{
			orig: f64Four, to: i32x2VecTyp,
			want: "", err: `invalid floating point to unsigned integer conversion; cannot convert from "double" to "<2 x i32>"`,
		},
		// i=6
		{
			orig: f32x2VecThreeFour, to: i32x3VecTyp,
			want: "", err: `invalid floating point to unsigned integer conversion; cannot convert from "<2 x float>" to "<3 x i32>"`,
		},
	}

	for i, g := range golden {
		v, err := consts.NewFloatToUint(g.orig, g.to)
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
		orig consts.Constant
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
			orig: f32x2VecMinusThreeFour, to: i32x2VecTyp,
			want: "<2 x i32> fptosi(<2 x float> <float -3.0, float 4.0> to <2 x i32>)",
		},
		// i=2
		{
			orig: i32Four, to: i32Typ,
			want: "", err: `invalid floating point to signed integer conversion; expected floating point constant (or constant vector) for orig, got "i32"`,
		},
		// i=3
		{
			orig: f32Four, to: f32Typ,
			want: "", err: `invalid floating point to signed integer conversion; expected integer (or integer vector) target type, got "float"`,
		},
		// i=4
		{
			orig: f32x2VecThreeFour, to: i64Typ,
			want: "", err: `invalid floating point to signed integer conversion; cannot convert from "<2 x float>" to "i64"`,
		},
		// i=5
		{
			orig: f32Three, to: i32x2VecTyp,
			want: "", err: `invalid floating point to signed integer conversion; cannot convert from "float" to "<2 x i32>"`,
		},
		// i=6
		{
			orig: f32x3VecThreeFourFifteen, to: i32x2VecTyp,
			want: "", err: `invalid floating point to signed integer conversion; cannot convert from "<3 x float>" to "<2 x i32>"`,
		},
	}

	for i, g := range golden {
		v, err := consts.NewFloatToInt(g.orig, g.to)
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
		orig consts.Constant
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
			orig: i32x2VecThreeFortyTwo, to: f32x2VecTyp,
			want: "<2 x float> uitofp(<2 x i32> <i32 3, i32 42> to <2 x float>)",
		},
		// i=2
		{
			orig: i32i8FourThree, to: f32Typ,
			want: "", err: `invalid unsigned integer to floating point conversion; expected integer constant (or constant vector) for orig, got "{i32, i8}"`,
		},
		// i=3
		{
			orig: i32x2VecMinusThreeFifteen, to: i32x2VecTyp,
			want: "", err: `invalid unsigned integer to floating point conversion; expected floating point (or floating point vector) target type, got "<2 x i32>"`,
		},
		// i=4
		{
			orig: i32x2VecThreeFortyTwo, to: f32Typ,
			want: "", err: `invalid unsigned integer to floating point conversion; cannot convert from "<2 x i32>" to "float"`,
		},
		// i=5
		{
			orig: i32Fifteen, to: f32x2VecTyp,
			want: "", err: `invalid unsigned integer to floating point conversion; cannot convert from "i32" to "<2 x float>"`,
		},
		// i=6
		{
			orig: i32x2VecMinusThreeFifteen, to: f32x3VecTyp,
			want: "", err: `invalid unsigned integer to floating point conversion; cannot convert from "<2 x i32>" to "<3 x float>"`,
		},
	}

	for i, g := range golden {
		v, err := consts.NewUintToFloat(g.orig, g.to)
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
		orig consts.Constant
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
			orig: i32x2VecMinusThreeFifteen, to: f32x2VecTyp,
			want: "<2 x float> sitofp(<2 x i32> <i32 -3, i32 15> to <2 x float>)",
		},
		// i=2
		{
			orig: f64Four, to: f32Typ,
			want: "", err: `invalid signed integer to floating point conversion; expected integer constant (or constant vector) for orig, got "double"`,
		},
		// i=3
		{
			orig: i32Four, to: i32Typ,
			want: "", err: `invalid signed integer to floating point conversion; expected floating point (or floating point vector) target type, got "i32"`,
		},
		// i=4
		{
			orig: i32x2VecMinusThreeFifteen, to: f32Typ,
			want: "", err: `invalid signed integer to floating point conversion; cannot convert from "<2 x i32>" to "float"`,
		},
		// i=5
		{
			orig: i32Three, to: f32x2VecTyp,
			want: "", err: `invalid signed integer to floating point conversion; cannot convert from "i32" to "<2 x float>"`,
		},
		// i=6
		{
			orig: i32x3OneTwoThree, to: f32x2VecTyp,
			want: "", err: `invalid signed integer to floating point conversion; cannot convert from "<3 x i32>" to "<2 x float>"`,
		},
	}

	for i, g := range golden {
		v, err := consts.NewIntToFloat(g.orig, g.to)
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
