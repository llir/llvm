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
