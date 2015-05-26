package types_test

import (
	"log"
	"strings"
	"testing"

	"github.com/llir/llvm/types"
)

// Types used by test cases.
var (
	// Void type.
	voidTyp *types.Void // void

	// Integer types.
	i1Typ  *types.Int // i1
	i8Typ  *types.Int // i8
	i32Typ *types.Int // i32
	i64Typ *types.Int // i64

	// Floating point types.
	f16Typ      *types.Float // half
	f32Typ      *types.Float // float
	f64Typ      *types.Float // double
	f128Typ     *types.Float // fp128
	f80_x86Typ  *types.Float // x86_fp80
	f128_ppcTyp *types.Float // ppc_fp128

	// MMX type.
	mmxTyp *types.MMX // x86_mmx

	// Label type.
	labelTyp *types.Label // label

	// Metadata type.
	metadataTyp *types.Metadata // metadata

	// Function types.
	voidFuncTyp            *types.Func // void ()
	i32FuncTyp             *types.Func // i32 ()
	voidFunci32Typ         *types.Func // void (i32)
	voidFuncf32Typ         *types.Func // void (float)
	voidFunci32EllipsisTyp *types.Func // void (i32, ...)
	funcTyp                *types.Func // i32 (i32)

	// Pointer types.
	i8PtrTyp   *types.Pointer // i8*
	f16PtrTyp  *types.Pointer // half*
	mmxPtrTyp  *types.Pointer // x86_mmx*
	funcPtrTyp *types.Pointer // i32 (i32)*

	// Vector types.
	i8x1VecTyp       *types.Vector // <1 x i8>
	i32x2VecTyp      *types.Vector // <2 x i32>
	i32x3VecTyp      *types.Vector // <3 x i32>
	f16x3VecTyp      *types.Vector // <3 x half>
	f32x4VecTyp      *types.Vector // <4 x float>
	f64x5VecTyp      *types.Vector // <5 x double>
	f128x6VecTyp     *types.Vector // <6 x fp128>
	f80_x86x7VecTyp  *types.Vector // <7 x x86_fp80>
	f128_ppcx8VecTyp *types.Vector // <8 x ppc_fp128>
	i8Ptrx9VecTyp    *types.Vector // <9 x i8*>
	f16Ptrx10VecTyp  *types.Vector // <10 x half*>

	// Array types.
	i8x1ArrTyp       *types.Array // [1 x i8]
	i32x2ArrTyp      *types.Array // [2 x i32]
	i32x3ArrTyp      *types.Array // [3 x i32]
	f16x3ArrTyp      *types.Array // [3 x half]
	f32x4ArrTyp      *types.Array // [4 x float]
	f64x5ArrTyp      *types.Array // [5 x double]
	f128x6ArrTyp     *types.Array // [6 x fp128]
	f80_x86x7ArrTyp  *types.Array // [7 x x86_fp80]
	f128_ppcx8ArrTyp *types.Array // [8 x ppc_fp128]
	i8Ptrx9ArrTyp    *types.Array // [9 x i8*]
	f16Ptrx10ArrTyp  *types.Array // [10 x half*]

	// Structure types.
	i32i8structTyp   *types.Struct // {i32, i8}
	i32i32structTyp  *types.Struct // {i32, i32}
	i32i8i8structTyp *types.Struct // {i32, i8, i8}
	structTyp        *types.Struct // {i1, float, x86_mmx, i32 (i32)*, [1 x i8], <3 x half>}
)

func init() {
	var err error
	// Void type.
	// void
	voidTyp = types.NewVoid()

	// Integer types.
	// i1
	i1Typ, err = types.NewInt(1)
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

	// Floating point types.
	// half
	f16Typ, err = types.NewFloat(types.Float16)
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
	// x86_fp80
	f80_x86Typ, err = types.NewFloat(types.Float80_x86)
	if err != nil {
		log.Fatalln(err)
	}
	// ppc_fp128
	f128_ppcTyp, err = types.NewFloat(types.Float128_PPC)

	// MMX type.
	// x86_mmx
	mmxTyp = types.NewMMX()

	// Label type.
	// label
	labelTyp = types.NewLabel()

	// Metadata type.
	// metadata
	metadataTyp = types.NewMetadata()

	// Function types.
	// void ()
	voidFuncTyp, err = types.NewFunc(voidTyp, nil, false)
	if err != nil {
		log.Fatalln(err)
	}
	// i32 ()
	i32FuncTyp, err = types.NewFunc(i32Typ, nil, false)
	if err != nil {
		log.Fatalln(err)
	}
	// void (i32)
	voidFunci32Typ, err = types.NewFunc(voidTyp, []types.Type{i32Typ}, false)
	if err != nil {
		log.Fatalln(err)
	}
	// void (float)
	voidFuncf32Typ, err = types.NewFunc(voidTyp, []types.Type{f32Typ}, false)
	if err != nil {
		log.Fatalln(err)
	}
	// void (i32, ...)
	voidFunci32EllipsisTyp, err = types.NewFunc(voidTyp, []types.Type{i32Typ}, true)
	if err != nil {
		log.Fatalln(err)
	}
	// i32 (i32)
	funcTyp, err = types.NewFunc(i32Typ, []types.Type{i32Typ}, false)
	if err != nil {
		log.Fatalln(err)
	}

	// Pointer types.
	// i8*
	i8PtrTyp, err = types.NewPointer(i8Typ)
	if err != nil {
		log.Fatalln(err)
	}
	// half*
	f16PtrTyp, err = types.NewPointer(f16Typ)
	if err != nil {
		log.Fatalln(err)
	}
	// x86_mmx*
	mmxPtrTyp, err = types.NewPointer(mmxTyp)
	if err != nil {
		log.Fatalln(err)
	}
	// i32 (i32)*
	funcPtrTyp, err = types.NewPointer(funcTyp)
	if err != nil {
		log.Fatalln(err)
	}

	// Vector types.
	// <1 x i8>
	i8x1VecTyp, err = types.NewVector(i8Typ, 1)
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
	// <3 x half>
	f16x3VecTyp, err = types.NewVector(f16Typ, 3)
	if err != nil {
		log.Fatalln(err)
	}
	// <4 x float>
	f32x4VecTyp, err = types.NewVector(f32Typ, 4)
	if err != nil {
		log.Fatalln(err)
	}
	// <5 x double>
	f64x5VecTyp, err = types.NewVector(f64Typ, 5)
	if err != nil {
		log.Fatalln(err)
	}
	// <6 x fp128>
	f128x6VecTyp, err = types.NewVector(f128Typ, 6)
	if err != nil {
		log.Fatalln(err)
	}
	// <7 x x86_fp80>
	f80_x86x7VecTyp, err = types.NewVector(f80_x86Typ, 7)
	if err != nil {
		log.Fatalln(err)
	}
	// <8 x ppc_fp128>
	f128_ppcx8VecTyp, err = types.NewVector(f128_ppcTyp, 8)
	if err != nil {
		log.Fatalln(err)
	}
	// <9 x i8*>
	i8Ptrx9VecTyp, err = types.NewVector(i8PtrTyp, 9)
	if err != nil {
		log.Fatalln(err)
	}
	// <10 x half*>
	f16Ptrx10VecTyp, err = types.NewVector(f16PtrTyp, 10)
	if err != nil {
		log.Fatalln(err)
	}

	// Array types.
	// [1 x i8]
	i8x1ArrTyp, err = types.NewArray(i8Typ, 1)
	if err != nil {
		log.Fatalln(err)
	}
	// [2 x i32]
	i32x2ArrTyp, err = types.NewArray(i32Typ, 2)
	if err != nil {
		log.Fatalln(err)
	}
	// [3 x i32]
	i32x3ArrTyp, err = types.NewArray(i32Typ, 3)
	if err != nil {
		log.Fatalln(err)
	}
	// [3 x half]
	f16x3ArrTyp, err = types.NewArray(f16Typ, 3)
	if err != nil {
		log.Fatalln(err)
	}
	// [4 x float]
	f32x4ArrTyp, err = types.NewArray(f32Typ, 4)
	if err != nil {
		log.Fatalln(err)
	}
	// [5 x double]
	f64x5ArrTyp, err = types.NewArray(f64Typ, 5)
	if err != nil {
		log.Fatalln(err)
	}
	// [6 x fp128]
	f128x6ArrTyp, err = types.NewArray(f128Typ, 6)
	if err != nil {
		log.Fatalln(err)
	}
	// [7 x x86_fp80]
	f80_x86x7ArrTyp, err = types.NewArray(f80_x86Typ, 7)
	if err != nil {
		log.Fatalln(err)
	}
	// [8 x ppc_fp128]
	f128_ppcx8ArrTyp, err = types.NewArray(f128_ppcTyp, 8)
	if err != nil {
		log.Fatalln(err)
	}
	// [9 x i8*]
	i8Ptrx9ArrTyp, err = types.NewArray(i8PtrTyp, 9)
	if err != nil {
		log.Fatalln(err)
	}
	// [10 x half*]
	f16Ptrx10ArrTyp, err = types.NewArray(f16PtrTyp, 10)
	if err != nil {
		log.Fatalln(err)
	}

	// Structure types.
	// {i32, i8}
	fields := []types.Type{i32Typ, i8Typ}
	i32i8structTyp, err = types.NewStruct(fields, false)
	if err != nil {
		log.Fatalln(err)
	}
	// {i32, i32}
	fields = []types.Type{i32Typ, i32Typ}
	i32i32structTyp, err = types.NewStruct(fields, false)
	if err != nil {
		log.Fatalln(err)
	}
	// {i32, i8, i8}
	fields = []types.Type{i32Typ, i8Typ, i8Typ}
	i32i8i8structTyp, err = types.NewStruct(fields, false)
	if err != nil {
		log.Fatalln(err)
	}
	// {i1, float, x86_mmx, i32 (i32)*, <1 x i8>, [3 x half]}
	fields = []types.Type{i1Typ, f32Typ, mmxTyp, funcPtrTyp, i8x1VecTyp, f16x3ArrTyp}
	structTyp, err = types.NewStruct(fields, false)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestVoidString(t *testing.T) {
	const want = "void"
	typ := types.NewVoid()
	got := typ.String()
	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestIntString(t *testing.T) {
	golden := []struct {
		n    int
		want string
		err  string
	}{
		// i=0
		{
			n:    1,
			want: "i1",
		},
		// i=1
		{
			n:    32,
			want: "i32",
		},
		// i=2
		{
			n:    1<<23 - 1,
			want: "i8388607",
		},
		// i=3
		{
			n:    -1,
			want: "", err: "invalid integer size (-1)",
		},
		// i=4
		{
			n:    0,
			want: "", err: "invalid integer size (0)",
		},
		// i=5
		{
			n:    1 << 23,
			want: "", err: "invalid integer size (8388608)",
		},
	}

	for i, g := range golden {
		typ, err := types.NewInt(g.n)
		if !sameError(err, g.err) {
			t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			continue
		} else if err != nil {
			// Expected error match, check next test case.
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestFloatSize(t *testing.T) {
	golden := []struct {
		kind types.FloatKind
		want int
	}{
		// i=0
		{
			kind: types.Float16,
			want: 16,
		},
		// i=1
		{
			kind: types.Float32,
			want: 32,
		},
		// i=2
		{
			kind: types.Float64,
			want: 64,
		},
		// i=3
		{
			kind: types.Float128,
			want: 128,
		},
		// i=4
		{
			kind: types.Float80_x86,
			want: 80,
		},
		// i=5
		{
			kind: types.Float128_PPC,
			want: 128,
		},
	}

	for i, g := range golden {
		typ, err := types.NewFloat(g.kind)
		if err != nil {
			t.Errorf("i=%d: %v", i, err)
			continue
		}
		got := typ.Size()
		if got != g.want {
			t.Errorf("i=%d: size mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestFloatString(t *testing.T) {
	golden := []struct {
		kind types.FloatKind
		want string
		err  string
	}{
		// i=0
		{
			kind: types.Float16,
			want: "half",
		},
		// i=1
		{
			kind: types.Float32,
			want: "float",
		},
		// i=2
		{
			kind: types.Float64,
			want: "double",
		},
		// i=3
		{
			kind: types.Float128,
			want: "fp128",
		},
		// i=4
		{
			kind: types.Float80_x86,
			want: "x86_fp80",
		},
		// i=5
		{
			kind: types.Float128_PPC,
			want: "ppc_fp128",
		},
		// i=6
		{
			kind: -1,
			want: "", err: "invalid floating point kind (-1)",
		},
	}

	for i, g := range golden {
		typ, err := types.NewFloat(g.kind)
		if !sameError(err, g.err) {
			t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			continue
		} else if err != nil {
			// Expected error match, check next test case.
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestMMXString(t *testing.T) {
	const want = "x86_mmx"
	typ := types.NewMMX()
	got := typ.String()
	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestLabelString(t *testing.T) {
	const want = "label"
	typ := types.NewLabel()
	got := typ.String()
	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestMetadataString(t *testing.T) {
	const want = "metadata"
	typ := types.NewMetadata()
	got := typ.String()
	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestFuncString(t *testing.T) {
	golden := []struct {
		result   types.Type
		params   []types.Type
		variadic bool
		want     string
		err      string
	}{
		// i=0
		{
			result: voidTyp, params: nil,
			want: "void ()",
		},
		// i=1
		{
			result: i32Typ, params: []types.Type{i32Typ},
			want: "i32 (i32)",
		},
		// i=2
		{
			result: voidTyp, params: []types.Type{i32Typ, i8Typ},
			want: "void (i32, i8)",
		},
		// i=3
		{
			result: i32Typ, params: []types.Type{i8PtrTyp}, variadic: true,
			want: "i32 (i8*, ...)",
		},
		// i=4
		{
			result: i32Typ, params: []types.Type{voidTyp}, // i32 (void)
			want: "", err: "invalid function parameter type; void type only allowed for function results",
		},
		// i=5
		{
			result: i32Typ, params: []types.Type{funcTyp}, // i32 (i32 (i32))
			want: "", err: `invalid function parameter type "i32 (i32)"`,
		},
		// i=6
		{
			result: labelTyp, params: nil, // label ()
			want: "", err: `invalid result parameter type "label"`,
		},
	}

	for i, g := range golden {
		typ, err := types.NewFunc(g.result, g.params, g.variadic)
		if !sameError(err, g.err) {
			t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			continue
		} else if err != nil {
			// Expected error match, check next test case.
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestPointerString(t *testing.T) {
	golden := []struct {
		elem types.Type
		want string
		err  string
	}{
		// i=0
		{
			elem: i32Typ,
			want: "i32*",
		},
		// i=1
		{
			elem: f16Typ,
			want: "half*",
		},
		// i=2
		{
			elem: funcTyp,
			want: "i32 (i32)*",
		},
		// i=3
		{
			elem: i8PtrTyp,
			want: "i8**",
		},
		// i=4
		{
			elem: voidTyp, // void*
			want: "", err: `invalid pointer to "void"; use i8* instead`,
		},
		// i=5
		{
			elem: labelTyp, // label*
			want: "", err: `invalid pointer to "label"`,
		},
	}

	for i, g := range golden {
		typ, err := types.NewPointer(g.elem)
		if !sameError(err, g.err) {
			t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			continue
		} else if err != nil {
			// Expected error match, check next test case.
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestVectorString(t *testing.T) {
	golden := []struct {
		elem types.Type
		n    int
		want string
		err  string
	}{
		// i=0
		{
			elem: i32Typ, n: 1,
			want: "<1 x i32>",
		},
		// i=1
		{
			elem: i8PtrTyp, n: 5,
			want: "<5 x i8*>",
		},
		// i=2
		{
			elem: i8Typ, n: 10,
			want: "<10 x i8>",
		},
		// i=3
		{
			elem: f64Typ, n: 6,
			want: "<6 x double>",
		},
		// i=4
		{
			elem: i8Typ, n: -1, // <-1 x i8>
			want: "", err: "invalid vector length (-1)",
		},
		// i=5
		{
			elem: voidTyp, n: 5, // <5 x void>
			want: "", err: "invalid vector element type; void type only allowed for function results",
		},
		// i=6
		{
			elem: labelTyp, n: 3, // <3 x label>
			want: "", err: `invalid vector element type "label"`,
		},
		// i=7
		{
			elem: mmxTyp, n: 7, // <7 x label>
			want: "", err: `invalid vector element type "x86_mmx"`,
		},
		// i=8
		{
			elem: funcTyp, n: 2, // <2 x i32 (i32)>
			want: "", err: `invalid vector element type "i32 (i32)"`,
		},
	}

	for i, g := range golden {
		typ, err := types.NewVector(g.elem, g.n)
		if !sameError(err, g.err) {
			t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			continue
		} else if err != nil {
			// Expected error match, check next test case.
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestArrayString(t *testing.T) {
	golden := []struct {
		elem types.Type
		n    int
		want string
		err  string
	}{
		// i=0
		{
			elem: i32Typ, n: 1,
			want: "[1 x i32]",
		},
		// i=1
		{
			elem: i8PtrTyp, n: 5,
			want: "[5 x i8*]",
		},
		// i=2
		{
			elem: i8Typ, n: 10,
			want: "[10 x i8]",
		},
		// i=3
		{
			elem: f64Typ, n: 6,
			want: "[6 x double]",
		},
		// i=4
		{
			elem: mmxTyp, n: 7,
			want: "[7 x x86_mmx]",
		},
		// i=5
		{
			elem: i8Typ, n: -1, // [-1 x i8]
			want: "", err: "invalid array length (-1)",
		},
		// i=6
		{
			elem: voidTyp, n: 5, // [5 x void]
			want: "", err: "invalid array element type; void type only allowed for function results",
		},
		// i=7
		{
			elem: labelTyp, n: 3, // [3 x label]
			want: "", err: `invalid array element type "label"`,
		},
		// i=8
		{
			elem: funcTyp, n: 2, // [2 x i32 (i32)]
			want: "", err: `invalid array element type "i32 (i32)"`,
		},
	}

	for i, g := range golden {
		typ, err := types.NewArray(g.elem, g.n)
		if !sameError(err, g.err) {
			t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			continue
		} else if err != nil {
			// Expected error match, check next test case.
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestStructString(t *testing.T) {
	golden := []struct {
		fields []types.Type
		packed bool
		want   string
		err    string
	}{
		// i=0
		{
			fields: []types.Type{i32Typ},
			want:   "{i32}",
		},
		// i=1
		{
			fields: []types.Type{i8Typ, i8Typ},
			want:   "{i8, i8}",
		},
		// i=2
		{
			fields: []types.Type{i8PtrTyp}, packed: true,
			want: "<{i8*}>",
		},
		// i=3
		{
			fields: []types.Type{mmxTyp},
			want:   "{x86_mmx}",
		},
		// i=4
		{
			fields: []types.Type{voidTyp}, // {void}
			want:   "", err: "invalid structure field type; void type only allowed for function results",
		},
		// i=5
		{
			fields: []types.Type{funcTyp}, // {i32 (i32)}
			want:   "", err: `invalid structure field type "i32 (i32)"`,
		},
		// i=6
		{
			fields: []types.Type{labelTyp}, // {label}
			want:   "", err: `invalid structure field type "label"`,
		},
	}

	for i, g := range golden {
		typ, err := types.NewStruct(g.fields, g.packed)
		if !sameError(err, g.err) {
			t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			continue
		} else if err != nil {
			// Expected error match, check next test case.
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}

	const want = "{i1, float, x86_mmx, i32 (i32)*, <1 x i8>, [3 x half]}"
	got := structTyp.String()
	if got != want {
		t.Errorf("string mismatch; expected %v, got %v", want, got)
	}
}

func TestNamedStructString(t *testing.T) {
	golden := []struct {
		name string
		want string
	}{
		// i=0
		{
			name: "foo",
			want: "%foo",
		},
	}

	ctx := types.NewContext()
	for i, g := range golden {
		typ, err := ctx.Struct(g.name)
		if err != nil {
			t.Errorf("i=%d: unexpected error; %v", i, err)
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestEqual(t *testing.T) {
	golden := []struct {
		want bool
		a, b types.Type
	}{
		{want: true, a: voidTyp, b: voidTyp},
		{want: true, a: i1Typ, b: i1Typ},
		{want: true, a: i8Typ, b: i8Typ},
		{want: true, a: i32Typ, b: i32Typ},
		{want: true, a: f16Typ, b: f16Typ},
		{want: true, a: f32Typ, b: f32Typ},
		{want: true, a: f64Typ, b: f64Typ},
		{want: true, a: f128Typ, b: f128Typ},
		{want: true, a: f80_x86Typ, b: f80_x86Typ},
		{want: true, a: f128_ppcTyp, b: f128_ppcTyp},
		{want: true, a: mmxTyp, b: mmxTyp},
		{want: true, a: labelTyp, b: labelTyp},
		{want: true, a: metadataTyp, b: metadataTyp},
		{want: true, a: voidFuncTyp, b: voidFuncTyp},
		{want: true, a: i32FuncTyp, b: i32FuncTyp},
		{want: true, a: voidFunci32Typ, b: voidFunci32Typ},
		{want: true, a: voidFuncf32Typ, b: voidFuncf32Typ},
		{want: true, a: voidFunci32EllipsisTyp, b: voidFunci32EllipsisTyp},
		{want: true, a: funcTyp, b: funcTyp},
		{want: true, a: i8PtrTyp, b: i8PtrTyp},
		{want: true, a: f16PtrTyp, b: f16PtrTyp},
		{want: true, a: mmxPtrTyp, b: mmxPtrTyp},
		{want: true, a: funcPtrTyp, b: funcPtrTyp},
		{want: true, a: i8x1VecTyp, b: i8x1VecTyp},
		{want: true, a: i32x2VecTyp, b: i32x2VecTyp},
		{want: true, a: f16x3VecTyp, b: f16x3VecTyp},
		{want: true, a: f32x4VecTyp, b: f32x4VecTyp},
		{want: true, a: f64x5VecTyp, b: f64x5VecTyp},
		{want: true, a: f128x6VecTyp, b: f128x6VecTyp},
		{want: true, a: f80_x86x7VecTyp, b: f80_x86x7VecTyp},
		{want: true, a: f128_ppcx8VecTyp, b: f128_ppcx8VecTyp},
		{want: true, a: i8Ptrx9VecTyp, b: i8Ptrx9VecTyp},
		{want: true, a: f16Ptrx10VecTyp, b: f16Ptrx10VecTyp},
		{want: true, a: i8x1ArrTyp, b: i8x1ArrTyp},
		{want: true, a: i32x2ArrTyp, b: i32x2ArrTyp},
		{want: true, a: f16x3ArrTyp, b: f16x3ArrTyp},
		{want: true, a: f32x4ArrTyp, b: f32x4ArrTyp},
		{want: true, a: f64x5ArrTyp, b: f64x5ArrTyp},
		{want: true, a: f128x6ArrTyp, b: f128x6ArrTyp},
		{want: true, a: f80_x86x7ArrTyp, b: f80_x86x7ArrTyp},
		{want: true, a: f128_ppcx8ArrTyp, b: f128_ppcx8ArrTyp},
		{want: true, a: i8Ptrx9ArrTyp, b: i8Ptrx9ArrTyp},
		{want: true, a: f16Ptrx10ArrTyp, b: f16Ptrx10ArrTyp},
		{want: true, a: i32i8structTyp, b: i32i8structTyp},
		{want: true, a: i32i32structTyp, b: i32i32structTyp},
		{want: true, a: i32i8i8structTyp, b: i32i8i8structTyp},
		{want: true, a: structTyp, b: structTyp},
		{want: false, a: voidTyp, b: structTyp},
		{want: false, a: i1Typ, b: voidTyp},
		{want: false, a: i8Typ, b: i1Typ},
		{want: false, a: i32Typ, b: i8Typ},
		{want: false, a: f16Typ, b: i32Typ},
		{want: false, a: f32Typ, b: f16Typ},
		{want: false, a: f64Typ, b: f32Typ},
		{want: false, a: f128Typ, b: f64Typ},
		{want: false, a: f80_x86Typ, b: f128Typ},
		{want: false, a: f128_ppcTyp, b: f80_x86Typ},
		{want: false, a: mmxTyp, b: f128_ppcTyp},
		{want: false, a: labelTyp, b: mmxTyp},
		{want: false, a: metadataTyp, b: labelTyp},
		{want: false, a: voidFuncTyp, b: i32FuncTyp},
		{want: false, a: voidFuncTyp, b: voidFunci32Typ},
		{want: false, a: voidFunci32Typ, b: voidFuncf32Typ},
		{want: false, a: voidFunci32Typ, b: voidFunci32EllipsisTyp},
		{want: false, a: funcTyp, b: metadataTyp},
		{want: false, a: i8PtrTyp, b: funcTyp},
		{want: false, a: f16PtrTyp, b: i8PtrTyp},
		{want: false, a: mmxPtrTyp, b: f16PtrTyp},
		{want: false, a: funcPtrTyp, b: mmxPtrTyp},
		{want: false, a: i8x1VecTyp, b: funcPtrTyp},
		{want: false, a: i32x2VecTyp, b: i8x1VecTyp},
		{want: false, a: f16x3VecTyp, b: i32x2VecTyp},
		{want: false, a: f32x4VecTyp, b: f16x3VecTyp},
		{want: false, a: f64x5VecTyp, b: f32x4VecTyp},
		{want: false, a: f128x6VecTyp, b: f64x5VecTyp},
		{want: false, a: f80_x86x7VecTyp, b: f128x6VecTyp},
		{want: false, a: f128_ppcx8VecTyp, b: f80_x86x7VecTyp},
		{want: false, a: i8Ptrx9VecTyp, b: f128_ppcx8VecTyp},
		{want: false, a: f16Ptrx10VecTyp, b: i8Ptrx9VecTyp},
		{want: false, a: i8x1ArrTyp, b: f16Ptrx10VecTyp},
		{want: false, a: i32x2ArrTyp, b: i8x1ArrTyp},
		{want: false, a: f16x3ArrTyp, b: i32x2ArrTyp},
		{want: false, a: f32x4ArrTyp, b: f16x3ArrTyp},
		{want: false, a: f64x5ArrTyp, b: f32x4ArrTyp},
		{want: false, a: f128x6ArrTyp, b: f64x5ArrTyp},
		{want: false, a: f80_x86x7ArrTyp, b: f128x6ArrTyp},
		{want: false, a: f128_ppcx8ArrTyp, b: f80_x86x7ArrTyp},
		{want: false, a: i8Ptrx9ArrTyp, b: f128_ppcx8ArrTyp},
		{want: false, a: f16Ptrx10ArrTyp, b: i8Ptrx9ArrTyp},
		{want: false, a: i32i8structTyp, b: i32i8i8structTyp},
		{want: false, a: i32i8structTyp, b: i32i32structTyp},
		{want: false, a: structTyp, b: f16Ptrx10ArrTyp},
	}

	for i, g := range golden {
		got := types.Equal(g.a, g.b)
		if got != g.want {
			t.Errorf("i=%d: expected %v, got %v", i, g.want, got)
		}
	}
}

func TestIsInt(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: voidTyp},          // void
		{want: true, typ: i1Typ},             // i1
		{want: true, typ: i8Typ},             // i8
		{want: true, typ: i32Typ},            // i32
		{want: false, typ: f16Typ},           // half
		{want: false, typ: f32Typ},           // float
		{want: false, typ: f64Typ},           // double
		{want: false, typ: f128Typ},          // fp128
		{want: false, typ: f80_x86Typ},       // x86_fp80
		{want: false, typ: f128_ppcTyp},      // ppc_fp128
		{want: false, typ: mmxTyp},           // x86_mmx
		{want: false, typ: labelTyp},         // label
		{want: false, typ: metadataTyp},      // metadata
		{want: false, typ: funcTyp},          // i32 (i32)
		{want: false, typ: i8PtrTyp},         // i8*
		{want: false, typ: f16PtrTyp},        // half*
		{want: false, typ: mmxPtrTyp},        // x86_mmx*
		{want: false, typ: funcPtrTyp},       // i32 (i32)*
		{want: false, typ: i8x1VecTyp},       // <1 x i8>
		{want: false, typ: i32x2VecTyp},      // <2 x i32>
		{want: false, typ: f16x3VecTyp},      // <3 x half>
		{want: false, typ: f32x4VecTyp},      // <4 x float>
		{want: false, typ: f64x5VecTyp},      // <5 x double>
		{want: false, typ: f128x6VecTyp},     // <6 x fp128>
		{want: false, typ: f80_x86x7VecTyp},  // <7 x x86_fp80>
		{want: false, typ: f128_ppcx8VecTyp}, // <8 x ppc_fp128>
		{want: false, typ: i8Ptrx9VecTyp},    // <9 x i8*>
		{want: false, typ: f16Ptrx10VecTyp},  // <10 x half*>
		{want: false, typ: i8x1ArrTyp},       // [1 x i8]
		{want: false, typ: i32x2ArrTyp},      // [2 x i32]
		{want: false, typ: f16x3ArrTyp},      // [3 x half]
		{want: false, typ: f32x4ArrTyp},      // [4 x float]
		{want: false, typ: f64x5ArrTyp},      // [5 x double]
		{want: false, typ: f128x6ArrTyp},     // [6 x fp128]
		{want: false, typ: f80_x86x7ArrTyp},  // [7 x x86_fp80]
		{want: false, typ: f128_ppcx8ArrTyp}, // [8 x ppc_fp128]
		{want: false, typ: i8Ptrx9ArrTyp},    // [9 x i8*]
		{want: false, typ: f16Ptrx10ArrTyp},  // [10 x half*]
		{want: false, typ: structTyp},        // {i1, float, x86_mmx, i32 (i32)*, [1 x i8], <3 x half>}
	}

	for i, g := range golden {
		got := types.IsInt(g.typ)
		if got != g.want {
			t.Errorf("i=%d: expected %v, got %v for type %q", i, g.want, got, g.typ)
		}
	}
}

func TestIsInts(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: voidTyp},          // void
		{want: true, typ: i1Typ},             // i1
		{want: true, typ: i8Typ},             // i8
		{want: true, typ: i32Typ},            // i32
		{want: false, typ: f16Typ},           // half
		{want: false, typ: f32Typ},           // float
		{want: false, typ: f64Typ},           // double
		{want: false, typ: f128Typ},          // fp128
		{want: false, typ: f80_x86Typ},       // x86_fp80
		{want: false, typ: f128_ppcTyp},      // ppc_fp128
		{want: false, typ: mmxTyp},           // x86_mmx
		{want: false, typ: labelTyp},         // label
		{want: false, typ: metadataTyp},      // metadata
		{want: false, typ: funcTyp},          // i32 (i32)
		{want: false, typ: i8PtrTyp},         // i8*
		{want: false, typ: f16PtrTyp},        // half*
		{want: false, typ: mmxPtrTyp},        // x86_mmx*
		{want: false, typ: funcPtrTyp},       // i32 (i32)*
		{want: true, typ: i8x1VecTyp},        //<[1 x i8>
		{want: true, typ: i32x2VecTyp},       //<[2 x i32>
		{want: false, typ: f16x3VecTyp},      //<[3 x half>
		{want: false, typ: f32x4VecTyp},      //<[4 x float>
		{want: false, typ: f64x5VecTyp},      //<[5 x double>
		{want: false, typ: f128x6VecTyp},     //<[6 x fp128>
		{want: false, typ: f80_x86x7VecTyp},  //<[7 x x86_fp80>
		{want: false, typ: f128_ppcx8VecTyp}, //<[8 x ppc_fp128>
		{want: false, typ: i8Ptrx9VecTyp},    //<[9 x i8*>
		{want: false, typ: f16Ptrx10VecTyp},  //<[10 x half*>
		{want: false, typ: i8x1ArrTyp},       // [1 x i8]
		{want: false, typ: i32x2ArrTyp},      // [2 x i32]
		{want: false, typ: f16x3ArrTyp},      // [3 x half]
		{want: false, typ: f32x4ArrTyp},      // [4 x float]
		{want: false, typ: f64x5ArrTyp},      // [5 x double]
		{want: false, typ: f128x6ArrTyp},     // [6 x fp128]
		{want: false, typ: f80_x86x7ArrTyp},  // [7 x x86_fp80]
		{want: false, typ: f128_ppcx8ArrTyp}, // [8 x ppc_fp128]
		{want: false, typ: i8Ptrx9ArrTyp},    // [9 x i8*]
		{want: false, typ: f16Ptrx10ArrTyp},  // [10 x half*]
		{want: false, typ: structTyp},        // {i1, float, x86_mmx, i32 (i32)*, [1 x i8], <3 x half>}
	}

	for i, g := range golden {
		got := types.IsInts(g.typ)
		if got != g.want {
			t.Errorf("i=%d: expected %v, got %v for type %q", i, g.want, got, g.typ)
		}
	}
}

func TestIsFloat(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: voidTyp},          // void
		{want: false, typ: i1Typ},            // i1
		{want: false, typ: i8Typ},            // i8
		{want: false, typ: i32Typ},           // i32
		{want: true, typ: f16Typ},            // half
		{want: true, typ: f32Typ},            // float
		{want: true, typ: f64Typ},            // double
		{want: true, typ: f128Typ},           // fp128
		{want: true, typ: f80_x86Typ},        // x86_fp80
		{want: true, typ: f128_ppcTyp},       // ppc_fp128
		{want: false, typ: mmxTyp},           // x86_mmx
		{want: false, typ: labelTyp},         // label
		{want: false, typ: metadataTyp},      // metadata
		{want: false, typ: funcTyp},          // i32 (i32)
		{want: false, typ: i8PtrTyp},         // i8*
		{want: false, typ: f16PtrTyp},        // half*
		{want: false, typ: mmxPtrTyp},        // x86_mmx*
		{want: false, typ: funcPtrTyp},       // i32 (i32)*
		{want: false, typ: i8x1VecTyp},       // <1 x i8>
		{want: false, typ: i32x2VecTyp},      // <2 x i32>
		{want: false, typ: f16x3VecTyp},      // <3 x half>
		{want: false, typ: f32x4VecTyp},      // <4 x float>
		{want: false, typ: f64x5VecTyp},      // <5 x double>
		{want: false, typ: f128x6VecTyp},     // <6 x fp128>
		{want: false, typ: f80_x86x7VecTyp},  // <7 x x86_fp80>
		{want: false, typ: f128_ppcx8VecTyp}, // <8 x ppc_fp128>
		{want: false, typ: i8Ptrx9VecTyp},    // <9 x i8*>
		{want: false, typ: f16Ptrx10VecTyp},  // <10 x half*>
		{want: false, typ: i8x1ArrTyp},       // [1 x i8]
		{want: false, typ: i32x2ArrTyp},      // [2 x i32]
		{want: false, typ: f16x3ArrTyp},      // [3 x half]
		{want: false, typ: f32x4ArrTyp},      // [4 x float]
		{want: false, typ: f64x5ArrTyp},      // [5 x double]
		{want: false, typ: f128x6ArrTyp},     // [6 x fp128]
		{want: false, typ: f80_x86x7ArrTyp},  // [7 x x86_fp80]
		{want: false, typ: f128_ppcx8ArrTyp}, // [8 x ppc_fp128]
		{want: false, typ: i8Ptrx9ArrTyp},    // [9 x i8*]
		{want: false, typ: f16Ptrx10ArrTyp},  // [10 x half*]
		{want: false, typ: structTyp},        // {i1, float, x86_mmx, i32 (i32)*, [1 x i8], <3 x half>}
	}

	for i, g := range golden {
		got := types.IsFloat(g.typ)
		if got != g.want {
			t.Errorf("i=%d: expected %v, got %v for type %q", i, g.want, got, g.typ)
		}
	}
}

func TestIsFloats(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: voidTyp},          // void
		{want: false, typ: i1Typ},            // i1
		{want: false, typ: i8Typ},            // i8
		{want: false, typ: i32Typ},           // i32
		{want: true, typ: f16Typ},            // half
		{want: true, typ: f32Typ},            // float
		{want: true, typ: f64Typ},            // double
		{want: true, typ: f128Typ},           // fp128
		{want: true, typ: f80_x86Typ},        // x86_fp80
		{want: true, typ: f128_ppcTyp},       // ppc_fp128
		{want: false, typ: mmxTyp},           // x86_mmx
		{want: false, typ: labelTyp},         // label
		{want: false, typ: metadataTyp},      // metadata
		{want: false, typ: funcTyp},          // i32 (i32)
		{want: false, typ: i8PtrTyp},         // i8*
		{want: false, typ: f16PtrTyp},        // half*
		{want: false, typ: mmxPtrTyp},        // x86_mmx*
		{want: false, typ: funcPtrTyp},       // i32 (i32)*
		{want: false, typ: i8x1VecTyp},       // <1 x i8>
		{want: false, typ: i32x2VecTyp},      // <2 x i32>
		{want: true, typ: f16x3VecTyp},       // <3 x half>
		{want: true, typ: f32x4VecTyp},       // <4 x float>
		{want: true, typ: f64x5VecTyp},       // <5 x double>
		{want: true, typ: f128x6VecTyp},      // <6 x fp128>
		{want: true, typ: f80_x86x7VecTyp},   // <7 x x86_fp80>
		{want: true, typ: f128_ppcx8VecTyp},  // <8 x ppc_fp128>
		{want: false, typ: i8Ptrx9VecTyp},    // <9 x i8*>
		{want: false, typ: f16Ptrx10VecTyp},  // <10 x half*>
		{want: false, typ: i8x1ArrTyp},       // [1 x i8]
		{want: false, typ: i32x2ArrTyp},      // [2 x i32]
		{want: false, typ: f16x3ArrTyp},      // [3 x half]
		{want: false, typ: f32x4ArrTyp},      // [4 x float]
		{want: false, typ: f64x5ArrTyp},      // [5 x double]
		{want: false, typ: f128x6ArrTyp},     // [6 x fp128]
		{want: false, typ: f80_x86x7ArrTyp},  // [7 x x86_fp80]
		{want: false, typ: f128_ppcx8ArrTyp}, // [8 x ppc_fp128]
		{want: false, typ: i8Ptrx9ArrTyp},    // [9 x i8*]
		{want: false, typ: f16Ptrx10ArrTyp},  // [10 x half*]
		{want: false, typ: structTyp},        // {i1, float, x86_mmx, i32 (i32)*, [1 x i8], <3 x half>}
	}

	for i, g := range golden {
		got := types.IsFloats(g.typ)
		if got != g.want {
			t.Errorf("i=%d: expected %v, got %v for type %q", i, g.want, got, g.typ)
		}
	}
}

func TestSameLength(t *testing.T) {
	golden := []struct {
		a, b types.Type
		want bool
	}{
		{want: false, a: i32x2VecTyp, b: i32x3VecTyp},
		{want: false, a: i32x3VecTyp, b: i32x2VecTyp},
		{want: true, a: i32x2VecTyp, b: i32x2VecTyp},
		{want: false, a: i32x2ArrTyp, b: i32x3ArrTyp},
		{want: false, a: i32x3ArrTyp, b: i32x2ArrTyp},
		{want: true, a: i32x2ArrTyp, b: i32x2ArrTyp},
		{want: true, a: i32Typ, b: i64Typ},
		{want: true, a: i64Typ, b: i32Typ},
		{want: false, a: f32Typ, b: i32x2VecTyp},
		{want: false, a: i32x2VecTyp, b: f32Typ},
		{want: false, a: f32Typ, b: i32x2ArrTyp},
		{want: false, a: i32x2ArrTyp, b: f32Typ},
		{want: true, a: structTyp, b: f32Typ},
	}

	for i, g := range golden {
		got := types.SameLength(g.a, g.b)
		if got != g.want {
			t.Errorf("i=%d: expected %v, got %v", i, g.want, got)
		}
	}
}

// sameError returns true if err is represented by the string s, and false
// otherwise. Some error messages contains suffixes from external functions,
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
