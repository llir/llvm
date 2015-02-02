package types_test

import (
	"log"
	"testing"

	"github.com/mewlang/llvm/types"
)

// Types used by test cases.
var (
	// Void type.
	voidTyp *types.Void // void

	// Integer types.
	i1Typ  *types.Int // i1
	i8Typ  *types.Int // i8
	i32Typ *types.Int // i32

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
	funcTyp *types.Func // i32 (i32)

	// Pointer types.
	i8PtrTyp   *types.Pointer // i8*
	f16PtrTyp  *types.Pointer // half*
	mmxPtrTyp  *types.Pointer // x86_mmx*
	funcPtrTyp *types.Pointer // i32 (i32)*

	// Vector types.
	i8VecTyp       *types.Vector // [1 x i8]
	i32VecTyp      *types.Vector // [2 x i32]
	f16VecTyp      *types.Vector // [3 x half]
	f32VecTyp      *types.Vector // [4 x float]
	f64VecTyp      *types.Vector // [5 x double]
	f128VecTyp     *types.Vector // [6 x fp128]
	f80_x86VecTyp  *types.Vector // [7 x x86_fp80]
	f128_ppcVecTyp *types.Vector // [8 x ppc_fp128]
	i8PtrVecTyp    *types.Vector // [9 x i8*]
	f16PtrVecTyp   *types.Vector // [10 x half*]

	// Array types.
	i8ArrTyp       *types.Array // <1 x i8>
	i32ArrTyp      *types.Array // <2 x i32>
	f16ArrTyp      *types.Array // <3 x half>
	f32ArrTyp      *types.Array // <4 x float>
	f64ArrTyp      *types.Array // <5 x double>
	f128ArrTyp     *types.Array // <6 x fp128>
	f80_x86ArrTyp  *types.Array // <7 x x86_fp80>
	f128_ppcArrTyp *types.Array // <8 x ppc_fp128>
	i8PtrArrTyp    *types.Array // <9 x i8*>
	f16PtrArrTyp   *types.Array // <10 x half*>

	// Struct types.
	structTyp *types.Struct // {i1, float, x86_mmx, i32 (i32)*, [1 x i8], <3 x half>}
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
	i8VecTyp, err = types.NewVector(i8Typ, 1)
	if err != nil {
		log.Fatalln(err)
	}
	// <2 x i32>
	i32VecTyp, err = types.NewVector(i32Typ, 2)
	if err != nil {
		log.Fatalln(err)
	}
	// <3 x half>
	f16VecTyp, err = types.NewVector(f16Typ, 3)
	if err != nil {
		log.Fatalln(err)
	}
	// <4 x float>
	f32VecTyp, err = types.NewVector(f32Typ, 4)
	if err != nil {
		log.Fatalln(err)
	}
	// <5 x double>
	f64VecTyp, err = types.NewVector(f64Typ, 5)
	if err != nil {
		log.Fatalln(err)
	}
	// <6 x fp128>
	f128VecTyp, err = types.NewVector(f128Typ, 6)
	if err != nil {
		log.Fatalln(err)
	}
	// <7 x x86_fp80>
	f80_x86VecTyp, err = types.NewVector(f80_x86Typ, 7)
	if err != nil {
		log.Fatalln(err)
	}
	// <8 x ppc_fp128>
	f128_ppcVecTyp, err = types.NewVector(f128_ppcTyp, 8)
	if err != nil {
		log.Fatalln(err)
	}
	// <9 x i8*>
	i8PtrVecTyp, err = types.NewVector(i8PtrTyp, 9)
	if err != nil {
		log.Fatalln(err)
	}
	// <10 x half*>
	f16PtrVecTyp, err = types.NewVector(f16PtrTyp, 10)
	if err != nil {
		log.Fatalln(err)
	}

	// Array types.
	// [1 x i8]
	i8ArrTyp, err = types.NewArray(i8Typ, 1)
	if err != nil {
		log.Fatalln(err)
	}
	// [2 x i32]
	i32ArrTyp, err = types.NewArray(i32Typ, 2)
	if err != nil {
		log.Fatalln(err)
	}
	// [3 x half]
	f16ArrTyp, err = types.NewArray(f16Typ, 3)
	if err != nil {
		log.Fatalln(err)
	}
	// [4 x float]
	f32ArrTyp, err = types.NewArray(f32Typ, 4)
	if err != nil {
		log.Fatalln(err)
	}
	// [5 x double]
	f64ArrTyp, err = types.NewArray(f64Typ, 5)
	if err != nil {
		log.Fatalln(err)
	}
	// [6 x fp128]
	f128ArrTyp, err = types.NewArray(f128Typ, 6)
	if err != nil {
		log.Fatalln(err)
	}
	// [7 x x86_fp80]
	f80_x86ArrTyp, err = types.NewArray(f80_x86Typ, 7)
	if err != nil {
		log.Fatalln(err)
	}
	// [8 x ppc_fp128]
	f128_ppcArrTyp, err = types.NewArray(f128_ppcTyp, 8)
	if err != nil {
		log.Fatalln(err)
	}
	// [9 x i8*]
	i8PtrArrTyp, err = types.NewArray(i8PtrTyp, 9)
	if err != nil {
		log.Fatalln(err)
	}
	// [10 x half*]
	f16PtrArrTyp, err = types.NewArray(f16PtrTyp, 10)
	if err != nil {
		log.Fatalln(err)
	}

	// Structure types.
	// {i1, float, x86_mmx, i32 (i32)*, <1 x i8>, [3 x half]}
	fields := []types.Type{i1Typ, f32Typ, mmxTyp, funcPtrTyp, i8VecTyp, f16ArrTyp}
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
		want string
		n    int
		err  string
	}{
		{want: "i1", n: 1},
		{want: "i32", n: 32},
		{want: "", n: -1, err: "invalid integer size (-1)"},
		{want: "", n: 0, err: "invalid integer size (0)"},
		{want: "", n: 1 << 23, err: "invalid integer size (8388608)"},
		{want: "i8388607", n: 1<<23 - 1},
	}

	for i, g := range golden {
		typ, err := types.NewInt(g.n)
		if err != nil {
			if err.Error() != g.err {
				t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			}
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestFloatString(t *testing.T) {
	golden := []struct {
		want string
		kind types.FloatKind
	}{
		{want: "half", kind: types.Float16},
		{want: "float", kind: types.Float32},
		{want: "double", kind: types.Float64},
		{want: "fp128", kind: types.Float128},
		{want: "x86_fp80", kind: types.Float80_x86},
		{want: "ppc_fp128", kind: types.Float128_PPC},
	}

	for i, g := range golden {
		typ, err := types.NewFloat(g.kind)
		if err != nil {
			t.Errorf("i=%d; %v", i, err)
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
		want     string
		result   types.Type
		params   []types.Type
		variadic bool
		err      string
	}{
		// i: 0
		{want: "void ()", result: voidTyp, params: nil},
		// i: 1
		{want: "i32 (i32)", result: i32Typ, params: []types.Type{i32Typ}},
		// i: 2
		{want: "void (i32, i8)", result: voidTyp, params: []types.Type{i32Typ, i8Typ}},
		// i: 3
		{want: "i32 (i8*, ...)", result: i32Typ, params: []types.Type{i8PtrTyp}, variadic: true},
		// i: 4
		// i32 (void)
		{want: "", result: i32Typ, params: []types.Type{voidTyp}, err: "invalid function parameter type; void type only allowed for function results"},
		// i: 5
		// i32 (i32 (i32))
		{want: "", result: i32Typ, params: []types.Type{funcTyp}, err: `invalid function parameter type "i32 (i32)"`},
		// i: 6
		// label ()
		{want: "", result: labelTyp, params: nil, err: `invalid result parameter type "label"`},
	}

	for i, g := range golden {
		typ, err := types.NewFunc(g.result, g.params, g.variadic)
		if err != nil {
			if err.Error() != g.err {
				t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			}
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
		want string
		elem types.Type
		err  string
	}{
		{want: "i32*", elem: i32Typ},
		{want: "half*", elem: f16Typ},
		{want: "i32 (i32)*", elem: funcTyp},
		{want: "i8**", elem: i8PtrTyp},
		// void*
		{want: "", elem: voidTyp, err: `invalid pointer to "void"; use i8* instead`},
		{want: "", elem: labelTyp, err: `invalid pointer to "label"`},
	}

	for i, g := range golden {
		typ, err := types.NewPointer(g.elem)
		if err != nil {
			if err.Error() != g.err {
				t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			}
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
		want string
		elem types.Type
		n    int
		err  string
	}{
		{want: "<1 x i32>", elem: i32Typ, n: 1},
		{want: "<5 x i8*>", elem: i8PtrTyp, n: 5},
		{want: "<10 x i8>", elem: i8Typ, n: 10},
		{want: "<6 x double>", elem: f64Typ, n: 6},
		// <-1 x i8>
		{want: "", elem: i8Typ, n: -1, err: "invalid vector length (-1)"},
		// <5 x void>
		{want: "", elem: voidTyp, n: 5, err: "invalid vector element type; void type only allowed for function results"},
		// <3 x label>
		{want: "", elem: labelTyp, n: 3, err: `invalid vector element type "label"`},
		// <7 x label>
		{want: "", elem: mmxTyp, n: 7, err: `invalid vector element type "x86_mmx"`},
		// <2 x i32 (i32)>
		{want: "", elem: funcTyp, n: 2, err: `invalid vector element type "i32 (i32)"`},
	}

	for i, g := range golden {
		typ, err := types.NewVector(g.elem, g.n)
		if err != nil {
			if err.Error() != g.err {
				t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			}
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
		want string
		elem types.Type
		n    int
		err  string
	}{
		{want: "[1 x i32]", elem: i32Typ, n: 1},
		{want: "[5 x i8*]", elem: i8PtrTyp, n: 5},
		{want: "[10 x i8]", elem: i8Typ, n: 10},
		{want: "[6 x double]", elem: f64Typ, n: 6},
		// [-1 x i8]
		{want: "", elem: i8Typ, n: -1, err: "invalid array length (-1)"},
		// [5 x void]
		{want: "", elem: voidTyp, n: 5, err: "invalid array element type; void type only allowed for function results"},
		// [3 x label]
		{want: "", elem: labelTyp, n: 3, err: `invalid array element type "label"`},
		{want: "[7 x x86_mmx]", elem: mmxTyp, n: 7},
		// [2 x i32 (i32)]
		{want: "", elem: funcTyp, n: 2, err: `invalid array element type "i32 (i32)"`},
	}

	for i, g := range golden {
		typ, err := types.NewArray(g.elem, g.n)
		if err != nil {
			if err.Error() != g.err {
				t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			}
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
		want   string
		fields []types.Type
		packed bool
		err    string
	}{
		{want: "{i32}", fields: []types.Type{i32Typ}},
		{want: "{i8, i8}", fields: []types.Type{i8Typ, i8Typ}},
		// {void}
		{want: "", fields: []types.Type{voidTyp}, err: "invalid structure field type; void type only allowed for function results"},
		{want: "", fields: []types.Type{funcTyp}, err: `invalid structure field type "i32 (i32)"`},
		{want: "<{i8*}>", fields: []types.Type{i8PtrTyp}, packed: true},
		// {label}
		{want: "", fields: []types.Type{labelTyp}, err: `invalid structure field type "label"`},
		{want: "{x86_mmx}", fields: []types.Type{mmxTyp}},
	}

	for i, g := range golden {
		typ, err := types.NewStruct(g.fields, g.packed)
		if err != nil {
			if err.Error() != g.err {
				t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			}
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

func TestIsInt(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: voidTyp},        // void
		{want: true, typ: i1Typ},           // i1
		{want: true, typ: i8Typ},           // i8
		{want: true, typ: i32Typ},          // i32
		{want: false, typ: f16Typ},         // half
		{want: false, typ: f32Typ},         // float
		{want: false, typ: f64Typ},         // double
		{want: false, typ: f128Typ},        // fp128
		{want: false, typ: f80_x86Typ},     // x86_fp80
		{want: false, typ: f128_ppcTyp},    // ppc_fp128
		{want: false, typ: mmxTyp},         // x86_mmx
		{want: false, typ: labelTyp},       // label
		{want: false, typ: metadataTyp},    // metadata
		{want: false, typ: funcTyp},        // i32 (i32)
		{want: false, typ: i8PtrTyp},       // i8*
		{want: false, typ: f16PtrTyp},      // half*
		{want: false, typ: mmxPtrTyp},      // x86_mmx*
		{want: false, typ: funcPtrTyp},     // i32 (i32)*
		{want: false, typ: i8VecTyp},       // [1 x i8]
		{want: false, typ: i32VecTyp},      // [2 x i32]
		{want: false, typ: f16VecTyp},      // [3 x half]
		{want: false, typ: f32VecTyp},      // [4 x float]
		{want: false, typ: f64VecTyp},      // [5 x double]
		{want: false, typ: f128VecTyp},     // [6 x fp128]
		{want: false, typ: f80_x86VecTyp},  // [7 x x86_fp80]
		{want: false, typ: f128_ppcVecTyp}, // [8 x ppc_fp128]
		{want: false, typ: i8PtrVecTyp},    // [9 x i8*]
		{want: false, typ: f16PtrVecTyp},   // [10 x half*]
		{want: false, typ: i8ArrTyp},       // <1 x i8>
		{want: false, typ: i32ArrTyp},      // <2 x i32>
		{want: false, typ: f16ArrTyp},      // <3 x half>
		{want: false, typ: f32ArrTyp},      // <4 x float>
		{want: false, typ: f64ArrTyp},      // <5 x double>
		{want: false, typ: f128ArrTyp},     // <6 x fp128>
		{want: false, typ: f80_x86ArrTyp},  // <7 x x86_fp80>
		{want: false, typ: f128_ppcArrTyp}, // <8 x ppc_fp128>
		{want: false, typ: i8PtrArrTyp},    // <9 x i8*>
		{want: false, typ: f16PtrArrTyp},   // <10 x half*>
		{want: false, typ: structTyp},      // {i1, float, x86_mmx, i32 (i32)*, [1 x i8], <3 x half>}
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
		{want: false, typ: voidTyp},        // void
		{want: true, typ: i1Typ},           // i1
		{want: true, typ: i8Typ},           // i8
		{want: true, typ: i32Typ},          // i32
		{want: false, typ: f16Typ},         // half
		{want: false, typ: f32Typ},         // float
		{want: false, typ: f64Typ},         // double
		{want: false, typ: f128Typ},        // fp128
		{want: false, typ: f80_x86Typ},     // x86_fp80
		{want: false, typ: f128_ppcTyp},    // ppc_fp128
		{want: false, typ: mmxTyp},         // x86_mmx
		{want: false, typ: labelTyp},       // label
		{want: false, typ: metadataTyp},    // metadata
		{want: false, typ: funcTyp},        // i32 (i32)
		{want: false, typ: i8PtrTyp},       // i8*
		{want: false, typ: f16PtrTyp},      // half*
		{want: false, typ: mmxPtrTyp},      // x86_mmx*
		{want: false, typ: funcPtrTyp},     // i32 (i32)*
		{want: true, typ: i8VecTyp},        // [1 x i8]
		{want: true, typ: i32VecTyp},       // [2 x i32]
		{want: false, typ: f16VecTyp},      // [3 x half]
		{want: false, typ: f32VecTyp},      // [4 x float]
		{want: false, typ: f64VecTyp},      // [5 x double]
		{want: false, typ: f128VecTyp},     // [6 x fp128]
		{want: false, typ: f80_x86VecTyp},  // [7 x x86_fp80]
		{want: false, typ: f128_ppcVecTyp}, // [8 x ppc_fp128]
		{want: false, typ: i8PtrVecTyp},    // [9 x i8*]
		{want: false, typ: f16PtrVecTyp},   // [10 x half*]
		{want: false, typ: i8ArrTyp},       // <1 x i8>
		{want: false, typ: i32ArrTyp},      // <2 x i32>
		{want: false, typ: f16ArrTyp},      // <3 x half>
		{want: false, typ: f32ArrTyp},      // <4 x float>
		{want: false, typ: f64ArrTyp},      // <5 x double>
		{want: false, typ: f128ArrTyp},     // <6 x fp128>
		{want: false, typ: f80_x86ArrTyp},  // <7 x x86_fp80>
		{want: false, typ: f128_ppcArrTyp}, // <8 x ppc_fp128>
		{want: false, typ: i8PtrArrTyp},    // <9 x i8*>
		{want: false, typ: f16PtrArrTyp},   // <10 x half*>
		{want: false, typ: structTyp},      // {i1, float, x86_mmx, i32 (i32)*, [1 x i8], <3 x half>}
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
		{want: false, typ: voidTyp},        // void
		{want: false, typ: i1Typ},          // i1
		{want: false, typ: i8Typ},          // i8
		{want: false, typ: i32Typ},         // i32
		{want: true, typ: f16Typ},          // half
		{want: true, typ: f32Typ},          // float
		{want: true, typ: f64Typ},          // double
		{want: true, typ: f128Typ},         // fp128
		{want: true, typ: f80_x86Typ},      // x86_fp80
		{want: true, typ: f128_ppcTyp},     // ppc_fp128
		{want: false, typ: mmxTyp},         // x86_mmx
		{want: false, typ: labelTyp},       // label
		{want: false, typ: metadataTyp},    // metadata
		{want: false, typ: funcTyp},        // i32 (i32)
		{want: false, typ: i8PtrTyp},       // i8*
		{want: false, typ: f16PtrTyp},      // half*
		{want: false, typ: mmxPtrTyp},      // x86_mmx*
		{want: false, typ: funcPtrTyp},     // i32 (i32)*
		{want: false, typ: i8VecTyp},       // [1 x i8]
		{want: false, typ: i32VecTyp},      // [2 x i32]
		{want: false, typ: f16VecTyp},      // [3 x half]
		{want: false, typ: f32VecTyp},      // [4 x float]
		{want: false, typ: f64VecTyp},      // [5 x double]
		{want: false, typ: f128VecTyp},     // [6 x fp128]
		{want: false, typ: f80_x86VecTyp},  // [7 x x86_fp80]
		{want: false, typ: f128_ppcVecTyp}, // [8 x ppc_fp128]
		{want: false, typ: i8PtrVecTyp},    // [9 x i8*]
		{want: false, typ: f16PtrVecTyp},   // [10 x half*]
		{want: false, typ: i8ArrTyp},       // <1 x i8>
		{want: false, typ: i32ArrTyp},      // <2 x i32>
		{want: false, typ: f16ArrTyp},      // <3 x half>
		{want: false, typ: f32ArrTyp},      // <4 x float>
		{want: false, typ: f64ArrTyp},      // <5 x double>
		{want: false, typ: f128ArrTyp},     // <6 x fp128>
		{want: false, typ: f80_x86ArrTyp},  // <7 x x86_fp80>
		{want: false, typ: f128_ppcArrTyp}, // <8 x ppc_fp128>
		{want: false, typ: i8PtrArrTyp},    // <9 x i8*>
		{want: false, typ: f16PtrArrTyp},   // <10 x half*>
		{want: false, typ: structTyp},      // {i1, float, x86_mmx, i32 (i32)*, [1 x i8], <3 x half>}
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
		{want: false, typ: voidTyp},        // void
		{want: false, typ: i1Typ},          // i1
		{want: false, typ: i8Typ},          // i8
		{want: false, typ: i32Typ},         // i32
		{want: true, typ: f16Typ},          // half
		{want: true, typ: f32Typ},          // float
		{want: true, typ: f64Typ},          // double
		{want: true, typ: f128Typ},         // fp128
		{want: true, typ: f80_x86Typ},      // x86_fp80
		{want: true, typ: f128_ppcTyp},     // ppc_fp128
		{want: false, typ: mmxTyp},         // x86_mmx
		{want: false, typ: labelTyp},       // label
		{want: false, typ: metadataTyp},    // metadata
		{want: false, typ: funcTyp},        // i32 (i32)
		{want: false, typ: i8PtrTyp},       // i8*
		{want: false, typ: f16PtrTyp},      // half*
		{want: false, typ: mmxPtrTyp},      // x86_mmx*
		{want: false, typ: funcPtrTyp},     // i32 (i32)*
		{want: false, typ: i8VecTyp},       // [1 x i8]
		{want: false, typ: i32VecTyp},      // [2 x i32]
		{want: true, typ: f16VecTyp},       // [3 x half]
		{want: true, typ: f32VecTyp},       // [4 x float]
		{want: true, typ: f64VecTyp},       // [5 x double]
		{want: true, typ: f128VecTyp},      // [6 x fp128]
		{want: true, typ: f80_x86VecTyp},   // [7 x x86_fp80]
		{want: true, typ: f128_ppcVecTyp},  // [8 x ppc_fp128]
		{want: false, typ: i8PtrVecTyp},    // [9 x i8*]
		{want: false, typ: f16PtrVecTyp},   // [10 x half*]
		{want: false, typ: i8ArrTyp},       // <1 x i8>
		{want: false, typ: i32ArrTyp},      // <2 x i32>
		{want: false, typ: f16ArrTyp},      // <3 x half>
		{want: false, typ: f32ArrTyp},      // <4 x float>
		{want: false, typ: f64ArrTyp},      // <5 x double>
		{want: false, typ: f128ArrTyp},     // <6 x fp128>
		{want: false, typ: f80_x86ArrTyp},  // <7 x x86_fp80>
		{want: false, typ: f128_ppcArrTyp}, // <8 x ppc_fp128>
		{want: false, typ: i8PtrArrTyp},    // <9 x i8*>
		{want: false, typ: f16PtrArrTyp},   // <10 x half*>
		{want: false, typ: structTyp},      // {i1, float, x86_mmx, i32 (i32)*, [1 x i8], <3 x half>}
	}

	for i, g := range golden {
		got := types.IsFloats(g.typ)
		if got != g.want {
			t.Errorf("i=%d: expected %v, got %v for type %q", i, g.want, got, g.typ)
		}
	}
}
