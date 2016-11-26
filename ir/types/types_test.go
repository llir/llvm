package types_test

import (
	"testing"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func TestVoidTypeString(t *testing.T) {
	const want = "void"
	got := types.Void.String()
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestLabelTypeString(t *testing.T) {
	const want = "label"
	got := types.Label.String()
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestIntTypeString(t *testing.T) {
	golden := []struct {
		want string
		typ  *types.IntType
	}{
		{want: "i1", typ: types.NewInt(1)},
		{want: "i8", typ: types.NewInt(8)},
		{want: "i16", typ: types.NewInt(16)},
		{want: "i32", typ: types.NewInt(32)},
		{want: "i64", typ: types.NewInt(64)},
		{want: "i128", typ: types.NewInt(128)},
		{want: "i1", typ: types.I1},
		{want: "i8", typ: types.I8},
		{want: "i16", typ: types.I16},
		{want: "i32", typ: types.I32},
		{want: "i64", typ: types.I64},
		{want: "i128", typ: types.I128},
	}
	for i, g := range golden {
		got := g.typ.String()
		if got != g.want {
			t.Errorf("i=%d; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestFloatTypeString(t *testing.T) {
	golden := []struct {
		want string
		typ  *types.FloatType
	}{
		{want: "half", typ: types.Half},
		{want: "float", typ: types.Float},
		{want: "double", typ: types.Double},
		{want: "fp128", typ: types.FP128},
		{want: "x86_fp80", typ: types.X86_FP80},
		{want: "ppc_fp128", typ: types.PPC_FP128},
	}
	for i, g := range golden {
		got := g.typ.String()
		if got != g.want {
			t.Errorf("i=%d; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestFuncTypeString(t *testing.T) {
	formatParam := types.NewParam("format", types.NewPointer(types.I8))
	printfSig := types.NewFunc(types.I32, formatParam)
	printfSig.SetVariadic(true)
	golden := []struct {
		want string
		typ  *types.FuncType
	}{
		{want: "void ()", typ: types.NewFunc(types.Void)},
		{want: "i32 ()", typ: types.NewFunc(types.I32)},
		{want: "i32 (i8*, ...)", typ: printfSig},
	}
	for i, g := range golden {
		got := g.typ.String()
		if got != g.want {
			t.Errorf("i=%d; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestPointerTypeString(t *testing.T) {
	golden := []struct {
		want string
		typ  *types.PointerType
	}{
		{want: "i8*", typ: types.NewPointer(types.I8)},
	}
	for i, g := range golden {
		got := g.typ.String()
		if got != g.want {
			t.Errorf("i=%d; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestVectorTypeString(t *testing.T) {
	golden := []struct {
		want string
		typ  *types.VectorType
	}{
		{want: "<10 x i8>", typ: types.NewVector(types.I8, 10)},
		{want: "<42 x i8*>", typ: types.NewVector(types.NewPointer(types.I8), 42)},
	}
	for i, g := range golden {
		got := g.typ.String()
		if got != g.want {
			t.Errorf("i=%d; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestArrayTypeString(t *testing.T) {
	golden := []struct {
		want string
		typ  *types.ArrayType
	}{
		{want: "[10 x i8]", typ: types.NewArray(types.I8, 10)},
		{want: "[42 x i8*]", typ: types.NewArray(types.NewPointer(types.I8), 42)},
	}
	for i, g := range golden {
		got := g.typ.String()
		if got != g.want {
			t.Errorf("i=%d; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestStructTypeString(t *testing.T) {
	golden := []struct {
		want string
		typ  *types.StructType
	}{
		{want: "{i32, i8*}", typ: types.NewStruct(types.I32, types.NewPointer(types.I8))},
		{want: "{i32, i16, i8}", typ: types.NewStruct(types.I32, types.I16, types.I8)},
		{want: "{}", typ: types.NewStruct()},
	}
	for i, g := range golden {
		got := g.typ.String()
		if got != g.want {
			t.Errorf("i=%d; expected %q, got %q", i, g.want, got)
		}
	}
}

func TestIsVoid(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: true, typ: types.Void},
		{want: true, typ: &types.VoidType{}},
		{want: false, typ: types.Label},
		{want: false, typ: &types.LabelType{}},
		{want: false, typ: &types.IntType{}},
		{want: false, typ: types.I1},
		{want: false, typ: types.I8},
		{want: false, typ: types.I16},
		{want: false, typ: types.I32},
		{want: false, typ: types.I64},
		{want: false, typ: types.I128},
		{want: false, typ: &types.FloatType{}},
		{want: false, typ: types.Half},
		{want: false, typ: types.Float},
		{want: false, typ: types.Double},
		{want: false, typ: types.FP128},
		{want: false, typ: types.X86_FP80},
		{want: false, typ: types.PPC_FP128},
		{want: false, typ: &types.FuncType{}},
		{want: false, typ: &types.PointerType{}},
		{want: false, typ: &types.VectorType{}},
		{want: false, typ: &types.ArrayType{}},
		{want: false, typ: &types.StructType{}},
	}
	for i, g := range golden {
		got := types.IsVoid(g.typ)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestIsLabel(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: types.Void},
		{want: false, typ: &types.VoidType{}},
		{want: true, typ: types.Label},
		{want: true, typ: &types.LabelType{}},
		{want: false, typ: &types.IntType{}},
		{want: false, typ: types.I1},
		{want: false, typ: types.I8},
		{want: false, typ: types.I16},
		{want: false, typ: types.I32},
		{want: false, typ: types.I64},
		{want: false, typ: types.I128},
		{want: false, typ: &types.FloatType{}},
		{want: false, typ: types.Half},
		{want: false, typ: types.Float},
		{want: false, typ: types.Double},
		{want: false, typ: types.FP128},
		{want: false, typ: types.X86_FP80},
		{want: false, typ: types.PPC_FP128},
		{want: false, typ: &types.FuncType{}},
		{want: false, typ: &types.PointerType{}},
		{want: false, typ: &types.VectorType{}},
		{want: false, typ: &types.ArrayType{}},
		{want: false, typ: &types.StructType{}},
	}
	for i, g := range golden {
		got := types.IsLabel(g.typ)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestIsBool(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: types.Void},
		{want: false, typ: &types.VoidType{}},
		{want: false, typ: types.Label},
		{want: false, typ: &types.LabelType{}},
		{want: false, typ: &types.IntType{}},
		{want: true, typ: types.I1},
		{want: false, typ: types.I8},
		{want: false, typ: types.I16},
		{want: false, typ: types.I32},
		{want: false, typ: types.I64},
		{want: false, typ: types.I128},
		{want: false, typ: &types.FloatType{}},
		{want: false, typ: types.Half},
		{want: false, typ: types.Float},
		{want: false, typ: types.Double},
		{want: false, typ: types.FP128},
		{want: false, typ: types.X86_FP80},
		{want: false, typ: types.PPC_FP128},
		{want: false, typ: &types.FuncType{}},
		{want: false, typ: &types.PointerType{}},
		{want: false, typ: &types.VectorType{}},
		{want: false, typ: &types.ArrayType{}},
		{want: false, typ: &types.StructType{}},
	}
	for i, g := range golden {
		got := types.IsBool(g.typ)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestIsInt(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: types.Void},
		{want: false, typ: &types.VoidType{}},
		{want: false, typ: types.Label},
		{want: false, typ: &types.LabelType{}},
		{want: true, typ: &types.IntType{}},
		{want: true, typ: types.I1},
		{want: true, typ: types.I8},
		{want: true, typ: types.I16},
		{want: true, typ: types.I32},
		{want: true, typ: types.I64},
		{want: true, typ: types.I128},
		{want: false, typ: &types.FloatType{}},
		{want: false, typ: types.Half},
		{want: false, typ: types.Float},
		{want: false, typ: types.Double},
		{want: false, typ: types.FP128},
		{want: false, typ: types.X86_FP80},
		{want: false, typ: types.PPC_FP128},
		{want: false, typ: &types.FuncType{}},
		{want: false, typ: &types.PointerType{}},
		{want: false, typ: &types.VectorType{}},
		{want: false, typ: &types.ArrayType{}},
		{want: false, typ: &types.StructType{}},
	}
	for i, g := range golden {
		got := types.IsInt(g.typ)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestIsFloat(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: types.Void},
		{want: false, typ: &types.VoidType{}},
		{want: false, typ: types.Label},
		{want: false, typ: &types.LabelType{}},
		{want: false, typ: &types.IntType{}},
		{want: false, typ: types.I1},
		{want: false, typ: types.I8},
		{want: false, typ: types.I16},
		{want: false, typ: types.I32},
		{want: false, typ: types.I64},
		{want: false, typ: types.I128},
		{want: true, typ: &types.FloatType{}},
		{want: true, typ: types.Half},
		{want: true, typ: types.Float},
		{want: true, typ: types.Double},
		{want: true, typ: types.FP128},
		{want: true, typ: types.X86_FP80},
		{want: true, typ: types.PPC_FP128},
		{want: false, typ: &types.FuncType{}},
		{want: false, typ: &types.PointerType{}},
		{want: false, typ: &types.VectorType{}},
		{want: false, typ: &types.ArrayType{}},
		{want: false, typ: &types.StructType{}},
	}
	for i, g := range golden {
		got := types.IsFloat(g.typ)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestIsFunc(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: types.Void},
		{want: false, typ: &types.VoidType{}},
		{want: false, typ: types.Label},
		{want: false, typ: &types.LabelType{}},
		{want: false, typ: &types.IntType{}},
		{want: false, typ: types.I1},
		{want: false, typ: types.I8},
		{want: false, typ: types.I16},
		{want: false, typ: types.I32},
		{want: false, typ: types.I64},
		{want: false, typ: types.I128},
		{want: false, typ: &types.FloatType{}},
		{want: false, typ: types.Half},
		{want: false, typ: types.Float},
		{want: false, typ: types.Double},
		{want: false, typ: types.FP128},
		{want: false, typ: types.X86_FP80},
		{want: false, typ: types.PPC_FP128},
		{want: true, typ: &types.FuncType{}},
		{want: false, typ: &types.PointerType{}},
		{want: false, typ: &types.VectorType{}},
		{want: false, typ: &types.ArrayType{}},
		{want: false, typ: &types.StructType{}},
	}
	for i, g := range golden {
		got := types.IsFunc(g.typ)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestIsPointer(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: types.Void},
		{want: false, typ: &types.VoidType{}},
		{want: false, typ: types.Label},
		{want: false, typ: &types.LabelType{}},
		{want: false, typ: &types.IntType{}},
		{want: false, typ: types.I1},
		{want: false, typ: types.I8},
		{want: false, typ: types.I16},
		{want: false, typ: types.I32},
		{want: false, typ: types.I64},
		{want: false, typ: types.I128},
		{want: false, typ: &types.FloatType{}},
		{want: false, typ: types.Half},
		{want: false, typ: types.Float},
		{want: false, typ: types.Double},
		{want: false, typ: types.FP128},
		{want: false, typ: types.X86_FP80},
		{want: false, typ: types.PPC_FP128},
		{want: false, typ: &types.FuncType{}},
		{want: true, typ: &types.PointerType{}},
		{want: false, typ: &types.VectorType{}},
		{want: false, typ: &types.ArrayType{}},
		{want: false, typ: &types.StructType{}},
	}
	for i, g := range golden {
		got := types.IsPointer(g.typ)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestIsVector(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: types.Void},
		{want: false, typ: &types.VoidType{}},
		{want: false, typ: types.Label},
		{want: false, typ: &types.LabelType{}},
		{want: false, typ: &types.IntType{}},
		{want: false, typ: types.I1},
		{want: false, typ: types.I8},
		{want: false, typ: types.I16},
		{want: false, typ: types.I32},
		{want: false, typ: types.I64},
		{want: false, typ: types.I128},
		{want: false, typ: &types.FloatType{}},
		{want: false, typ: types.Half},
		{want: false, typ: types.Float},
		{want: false, typ: types.Double},
		{want: false, typ: types.FP128},
		{want: false, typ: types.X86_FP80},
		{want: false, typ: types.PPC_FP128},
		{want: false, typ: &types.FuncType{}},
		{want: false, typ: &types.PointerType{}},
		{want: true, typ: &types.VectorType{}},
		{want: false, typ: &types.ArrayType{}},
		{want: false, typ: &types.StructType{}},
	}
	for i, g := range golden {
		got := types.IsVector(g.typ)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestIsArray(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: types.Void},
		{want: false, typ: &types.VoidType{}},
		{want: false, typ: types.Label},
		{want: false, typ: &types.LabelType{}},
		{want: false, typ: &types.IntType{}},
		{want: false, typ: types.I1},
		{want: false, typ: types.I8},
		{want: false, typ: types.I16},
		{want: false, typ: types.I32},
		{want: false, typ: types.I64},
		{want: false, typ: types.I128},
		{want: false, typ: &types.FloatType{}},
		{want: false, typ: types.Half},
		{want: false, typ: types.Float},
		{want: false, typ: types.Double},
		{want: false, typ: types.FP128},
		{want: false, typ: types.X86_FP80},
		{want: false, typ: types.PPC_FP128},
		{want: false, typ: &types.FuncType{}},
		{want: false, typ: &types.PointerType{}},
		{want: false, typ: &types.VectorType{}},
		{want: true, typ: &types.ArrayType{}},
		{want: false, typ: &types.StructType{}},
	}
	for i, g := range golden {
		got := types.IsArray(g.typ)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestIsStruct(t *testing.T) {
	golden := []struct {
		want bool
		typ  types.Type
	}{
		{want: false, typ: types.Void},
		{want: false, typ: &types.VoidType{}},
		{want: false, typ: types.Label},
		{want: false, typ: &types.LabelType{}},
		{want: false, typ: &types.IntType{}},
		{want: false, typ: types.I1},
		{want: false, typ: types.I8},
		{want: false, typ: types.I16},
		{want: false, typ: types.I32},
		{want: false, typ: types.I64},
		{want: false, typ: types.I128},
		{want: false, typ: &types.FloatType{}},
		{want: false, typ: types.Half},
		{want: false, typ: types.Float},
		{want: false, typ: types.Double},
		{want: false, typ: types.FP128},
		{want: false, typ: types.X86_FP80},
		{want: false, typ: types.PPC_FP128},
		{want: false, typ: &types.FuncType{}},
		{want: false, typ: &types.PointerType{}},
		{want: false, typ: &types.VectorType{}},
		{want: false, typ: &types.ArrayType{}},
		{want: true, typ: &types.StructType{}},
	}
	for i, g := range golden {
		got := types.IsStruct(g.typ)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

// Valutate that the relevant types satisfy the value.Value interface.
var (
	_ value.Value = &types.Param{}
)

// Valutate that the relevant types satisfy the types.Type interface.
var (
	_ types.Type = &types.VoidType{}
	_ types.Type = &types.LabelType{}
	_ types.Type = &types.IntType{}
	_ types.Type = &types.FloatType{}
	_ types.Type = &types.FuncType{}
	_ types.Type = &types.PointerType{}
	_ types.Type = &types.VectorType{}
	_ types.Type = &types.ArrayType{}
	_ types.Type = &types.StructType{}
)
