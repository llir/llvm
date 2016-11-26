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

func TestVoidEqual(t *testing.T) {
	golden := []struct {
		want bool
		t, u types.Type
	}{
		{want: true, t: types.Void, u: &types.VoidType{}},
		{want: false, t: types.Void, u: &types.LabelType{}},
		{want: false, t: types.Void, u: &types.IntType{}},
		{want: false, t: types.Void, u: &types.FloatType{}},
		{want: false, t: types.Void, u: &types.FuncType{}},
		{want: false, t: types.Void, u: &types.PointerType{}},
		{want: false, t: types.Void, u: &types.VectorType{}},
		{want: false, t: types.Void, u: &types.ArrayType{}},
		{want: false, t: types.Void, u: &types.StructType{}},
	}
	for i, g := range golden {
		got := g.t.Equal(g.u)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestLabelEqual(t *testing.T) {
	golden := []struct {
		want bool
		t, u types.Type
	}{
		{want: false, t: types.Label, u: &types.VoidType{}},
		{want: true, t: types.Label, u: &types.LabelType{}},
		{want: false, t: types.Label, u: &types.IntType{}},
		{want: false, t: types.Label, u: &types.FloatType{}},
		{want: false, t: types.Label, u: &types.FuncType{}},
		{want: false, t: types.Label, u: &types.PointerType{}},
		{want: false, t: types.Label, u: &types.VectorType{}},
		{want: false, t: types.Label, u: &types.ArrayType{}},
		{want: false, t: types.Label, u: &types.StructType{}},
	}
	for i, g := range golden {
		got := g.t.Equal(g.u)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestIntEqual(t *testing.T) {
	golden := []struct {
		want bool
		t, u types.Type
	}{
		{want: false, t: &types.IntType{}, u: &types.VoidType{}},
		{want: false, t: &types.IntType{}, u: &types.LabelType{}},
		{want: true, t: types.I1, u: types.NewInt(1)},
		{want: false, t: types.I1, u: types.I8},
		{want: false, t: types.I1, u: types.I16},
		{want: false, t: types.I1, u: types.I32},
		{want: false, t: types.I1, u: types.I64},
		{want: false, t: types.I1, u: types.I128},
		{want: false, t: types.I8, u: types.I1},
		{want: true, t: types.I8, u: types.NewInt(8)},
		{want: false, t: types.I8, u: types.I16},
		{want: false, t: types.I8, u: types.I32},
		{want: false, t: types.I8, u: types.I64},
		{want: false, t: types.I8, u: types.I128},
		{want: false, t: types.I16, u: types.I1},
		{want: false, t: types.I16, u: types.I8},
		{want: true, t: types.I16, u: types.NewInt(16)},
		{want: false, t: types.I16, u: types.I32},
		{want: false, t: types.I16, u: types.I64},
		{want: false, t: types.I16, u: types.I128},
		{want: false, t: types.I32, u: types.I1},
		{want: false, t: types.I32, u: types.I8},
		{want: false, t: types.I32, u: types.I16},
		{want: true, t: types.I32, u: types.NewInt(32)},
		{want: false, t: types.I32, u: types.I64},
		{want: false, t: types.I32, u: types.I128},
		{want: false, t: types.I64, u: types.I1},
		{want: false, t: types.I64, u: types.I8},
		{want: false, t: types.I64, u: types.I16},
		{want: false, t: types.I64, u: types.I32},
		{want: true, t: types.I64, u: types.NewInt(64)},
		{want: false, t: types.I64, u: types.I128},
		{want: false, t: types.I128, u: types.I1},
		{want: false, t: types.I128, u: types.I8},
		{want: false, t: types.I128, u: types.I16},
		{want: false, t: types.I128, u: types.I32},
		{want: false, t: types.I128, u: types.I64},
		{want: true, t: types.I128, u: types.NewInt(128)},
		{want: false, t: &types.IntType{}, u: &types.FloatType{}},
		{want: false, t: &types.IntType{}, u: &types.FuncType{}},
		{want: false, t: &types.IntType{}, u: &types.PointerType{}},
		{want: false, t: &types.IntType{}, u: &types.VectorType{}},
		{want: false, t: &types.IntType{}, u: &types.ArrayType{}},
		{want: false, t: &types.IntType{}, u: &types.StructType{}},
	}
	for i, g := range golden {
		got := g.t.Equal(g.u)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestFloatEqual(t *testing.T) {
	golden := []struct {
		want bool
		t, u types.Type
	}{
		{want: false, t: &types.FloatType{}, u: &types.VoidType{}},
		{want: false, t: &types.FloatType{}, u: &types.LabelType{}},
		{want: false, t: &types.FloatType{}, u: &types.IntType{}},
		{want: true, t: types.Half, u: types.Half},
		{want: false, t: types.Half, u: types.Float},
		{want: false, t: types.Half, u: types.Double},
		{want: false, t: types.Half, u: types.FP128},
		{want: false, t: types.Half, u: types.X86_FP80},
		{want: false, t: types.Half, u: types.PPC_FP128},
		{want: false, t: types.Float, u: types.Half},
		{want: true, t: types.Float, u: types.Float},
		{want: false, t: types.Float, u: types.Double},
		{want: false, t: types.Float, u: types.FP128},
		{want: false, t: types.Float, u: types.X86_FP80},
		{want: false, t: types.Float, u: types.PPC_FP128},
		{want: false, t: types.Double, u: types.Half},
		{want: false, t: types.Double, u: types.Float},
		{want: true, t: types.Double, u: types.Double},
		{want: false, t: types.Double, u: types.FP128},
		{want: false, t: types.Double, u: types.X86_FP80},
		{want: false, t: types.Double, u: types.PPC_FP128},
		{want: false, t: types.FP128, u: types.Half},
		{want: false, t: types.FP128, u: types.Float},
		{want: false, t: types.FP128, u: types.Double},
		{want: true, t: types.FP128, u: types.FP128},
		{want: false, t: types.FP128, u: types.X86_FP80},
		{want: false, t: types.FP128, u: types.PPC_FP128},
		{want: false, t: types.X86_FP80, u: types.Half},
		{want: false, t: types.X86_FP80, u: types.Float},
		{want: false, t: types.X86_FP80, u: types.Double},
		{want: false, t: types.X86_FP80, u: types.FP128},
		{want: true, t: types.X86_FP80, u: types.X86_FP80},
		{want: false, t: types.X86_FP80, u: types.PPC_FP128},
		{want: false, t: types.PPC_FP128, u: types.Half},
		{want: false, t: types.PPC_FP128, u: types.Float},
		{want: false, t: types.PPC_FP128, u: types.Double},
		{want: false, t: types.PPC_FP128, u: types.FP128},
		{want: false, t: types.PPC_FP128, u: types.X86_FP80},
		{want: true, t: types.PPC_FP128, u: types.PPC_FP128},
		{want: false, t: &types.FloatType{}, u: &types.FuncType{}},
		{want: false, t: &types.FloatType{}, u: &types.PointerType{}},
		{want: false, t: &types.FloatType{}, u: &types.VectorType{}},
		{want: false, t: &types.FloatType{}, u: &types.ArrayType{}},
		{want: false, t: &types.FloatType{}, u: &types.StructType{}},
	}
	for i, g := range golden {
		got := g.t.Equal(g.u)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestFuncEqual(t *testing.T) {
	i8, i32 := types.I8, types.I32
	// void ()
	voidSig := types.NewFunc(types.Void)
	// i32 ()
	i32Sig := types.NewFunc(i32)
	// i32 (i32)
	i32i32Sig := types.NewFunc(i32, types.NewParam("x", i32))
	// i32 (i32, ...)
	i32i32VariadicSig := types.NewFunc(i32, types.NewParam("x", i32))
	i32i32VariadicSig.SetVariadic(true)
	// i32 (i32, i32)
	i32i32i32Sig := types.NewFunc(i32, types.NewParam("x", i32), types.NewParam("y", i32))
	// i32 (i32, float)
	i32i32FloatSig := types.NewFunc(i32, types.NewParam("x", i32), types.NewParam("y", types.Float))
	// i32 (i8*, ...)
	formatParam := types.NewParam("format", types.NewPointer(i8))
	printfSig := types.NewFunc(i32, formatParam)
	printfSig.SetVariadic(true)

	golden := []struct {
		want bool
		t, u types.Type
	}{
		{want: false, t: &types.FuncType{}, u: &types.VoidType{}},
		{want: false, t: &types.FuncType{}, u: &types.LabelType{}},
		{want: false, t: &types.FuncType{}, u: &types.IntType{}},
		{want: false, t: &types.FuncType{}, u: &types.FloatType{}},
		{want: true, t: voidSig, u: voidSig},
		{want: true, t: i32Sig, u: i32Sig},
		{want: true, t: i32i32Sig, u: i32i32Sig},
		{want: true, t: i32i32VariadicSig, u: i32i32VariadicSig},
		{want: true, t: i32i32FloatSig, u: i32i32FloatSig},
		{want: true, t: printfSig, u: printfSig},
		{want: true, t: i32i32i32Sig, u: i32i32i32Sig},
		{want: false, t: voidSig, u: i32Sig},
		{want: false, t: i32Sig, u: i32i32Sig},
		{want: false, t: i32i32Sig, u: i32i32VariadicSig},
		{want: false, t: i32i32Sig, u: i32i32i32Sig},
		{want: false, t: i32i32Sig, u: i32i32FloatSig},
		{want: false, t: i32i32VariadicSig, u: printfSig},
		{want: false, t: &types.FuncType{}, u: &types.PointerType{}},
		{want: false, t: &types.FuncType{}, u: &types.VectorType{}},
		{want: false, t: &types.FuncType{}, u: &types.ArrayType{}},
		{want: false, t: &types.FuncType{}, u: &types.StructType{}},
	}
	for i, g := range golden {
		got := g.t.Equal(g.u)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestPointerEqual(t *testing.T) {
	i8ptr := types.NewPointer(types.I8)
	i32ptr := types.NewPointer(types.I32)
	voidptr := types.NewPointer(types.Void)

	golden := []struct {
		want bool
		t, u types.Type
	}{
		{want: false, t: &types.PointerType{}, u: &types.VoidType{}},
		{want: false, t: &types.PointerType{}, u: &types.LabelType{}},
		{want: false, t: &types.PointerType{}, u: &types.IntType{}},
		{want: false, t: &types.PointerType{}, u: &types.FloatType{}},
		{want: false, t: &types.PointerType{}, u: &types.FuncType{}},
		{want: true, t: i8ptr, u: i8ptr},
		{want: false, t: i8ptr, u: i32ptr},
		{want: false, t: i8ptr, u: voidptr},
		{want: false, t: i32ptr, u: i8ptr},
		{want: true, t: i32ptr, u: i32ptr},
		{want: false, t: i32ptr, u: voidptr},
		{want: false, t: voidptr, u: i8ptr},
		{want: false, t: voidptr, u: i32ptr},
		{want: true, t: voidptr, u: voidptr},
		{want: false, t: &types.PointerType{}, u: &types.VectorType{}},
		{want: false, t: &types.PointerType{}, u: &types.ArrayType{}},
		{want: false, t: &types.PointerType{}, u: &types.StructType{}},
	}
	for i, g := range golden {
		got := g.t.Equal(g.u)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestVectorEqual(t *testing.T) {
	i32x8vec := types.NewVector(types.I32, 8)
	i32x10vec := types.NewVector(types.I32, 10)
	i8x8vec := types.NewVector(types.I8, 8)
	i8x10vec := types.NewVector(types.I8, 10)

	golden := []struct {
		want bool
		t, u types.Type
	}{
		{want: false, t: &types.VectorType{}, u: &types.VoidType{}},
		{want: false, t: &types.VectorType{}, u: &types.LabelType{}},
		{want: false, t: &types.VectorType{}, u: &types.IntType{}},
		{want: false, t: &types.VectorType{}, u: &types.FloatType{}},
		{want: false, t: &types.VectorType{}, u: &types.FuncType{}},
		{want: false, t: &types.VectorType{}, u: &types.PointerType{}},
		{want: true, t: i32x8vec, u: i32x8vec},
		{want: false, t: i32x8vec, u: i32x10vec},
		{want: false, t: i32x8vec, u: i8x8vec},
		{want: false, t: i32x8vec, u: i8x10vec},
		{want: false, t: i32x10vec, u: i32x8vec},
		{want: true, t: i32x10vec, u: i32x10vec},
		{want: false, t: i32x10vec, u: i8x8vec},
		{want: false, t: i32x10vec, u: i8x10vec},
		{want: false, t: i8x8vec, u: i32x8vec},
		{want: false, t: i8x8vec, u: i32x10vec},
		{want: true, t: i8x8vec, u: i8x8vec},
		{want: false, t: i8x8vec, u: i8x10vec},
		{want: false, t: i8x10vec, u: i32x8vec},
		{want: false, t: i8x10vec, u: i32x10vec},
		{want: false, t: i8x10vec, u: i8x8vec},
		{want: true, t: i8x10vec, u: i8x10vec},
		{want: false, t: &types.VectorType{}, u: &types.ArrayType{}},
		{want: false, t: &types.VectorType{}, u: &types.StructType{}},
	}
	for i, g := range golden {
		got := g.t.Equal(g.u)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestArrayEqual(t *testing.T) {
	i32x8arr := types.NewArray(types.I32, 8)
	i32x10arr := types.NewArray(types.I32, 10)
	i8x8arr := types.NewArray(types.I8, 8)
	i8x10arr := types.NewArray(types.I8, 10)

	golden := []struct {
		want bool
		t, u types.Type
	}{
		{want: false, t: &types.ArrayType{}, u: &types.VoidType{}},
		{want: false, t: &types.ArrayType{}, u: &types.LabelType{}},
		{want: false, t: &types.ArrayType{}, u: &types.IntType{}},
		{want: false, t: &types.ArrayType{}, u: &types.FloatType{}},
		{want: false, t: &types.ArrayType{}, u: &types.FuncType{}},
		{want: false, t: &types.ArrayType{}, u: &types.PointerType{}},
		{want: false, t: &types.ArrayType{}, u: &types.VectorType{}},
		{want: true, t: i32x8arr, u: i32x8arr},
		{want: false, t: i32x8arr, u: i32x10arr},
		{want: false, t: i32x8arr, u: i8x8arr},
		{want: false, t: i32x8arr, u: i8x10arr},
		{want: false, t: i32x10arr, u: i32x8arr},
		{want: true, t: i32x10arr, u: i32x10arr},
		{want: false, t: i32x10arr, u: i8x8arr},
		{want: false, t: i32x10arr, u: i8x10arr},
		{want: false, t: i8x8arr, u: i32x8arr},
		{want: false, t: i8x8arr, u: i32x10arr},
		{want: true, t: i8x8arr, u: i8x8arr},
		{want: false, t: i8x8arr, u: i8x10arr},
		{want: false, t: i8x10arr, u: i32x8arr},
		{want: false, t: i8x10arr, u: i32x10arr},
		{want: false, t: i8x10arr, u: i8x8arr},
		{want: true, t: i8x10arr, u: i8x10arr},
		{want: false, t: &types.ArrayType{}, u: &types.StructType{}},
	}
	for i, g := range golden {
		got := g.t.Equal(g.u)
		if got != g.want {
			t.Errorf("i=%d; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestStructEqual(t *testing.T) {
	i8, i32 := types.I8, types.I32
	i8i8struct := types.NewStruct(i8, i8)
	i8i32struct := types.NewStruct(i8, i32)
	i8struct := types.NewStruct(i8)
	i32struct := types.NewStruct(i32)
	emptystruct := types.NewStruct()

	golden := []struct {
		want bool
		t, u types.Type
	}{
		{want: false, t: &types.StructType{}, u: &types.VoidType{}},
		{want: false, t: &types.StructType{}, u: &types.LabelType{}},
		{want: false, t: &types.StructType{}, u: &types.IntType{}},
		{want: false, t: &types.StructType{}, u: &types.FloatType{}},
		{want: false, t: &types.StructType{}, u: &types.FuncType{}},
		{want: false, t: &types.StructType{}, u: &types.PointerType{}},
		{want: false, t: &types.StructType{}, u: &types.VectorType{}},
		{want: false, t: &types.StructType{}, u: &types.ArrayType{}},
		{want: true, t: i8i8struct, u: i8i8struct},
		{want: false, t: i8i8struct, u: i8i32struct},
		{want: false, t: i8i8struct, u: i8struct},
		{want: false, t: i8i8struct, u: i32struct},
		{want: false, t: i8i8struct, u: emptystruct},
		{want: false, t: i8i32struct, u: i8i8struct},
		{want: true, t: i8i32struct, u: i8i32struct},
		{want: false, t: i8i32struct, u: i8struct},
		{want: false, t: i8i32struct, u: i32struct},
		{want: false, t: i8i32struct, u: emptystruct},
		{want: false, t: i8struct, u: i8i8struct},
		{want: false, t: i8struct, u: i8i32struct},
		{want: true, t: i8struct, u: i8struct},
		{want: false, t: i8struct, u: i32struct},
		{want: false, t: i8struct, u: emptystruct},
		{want: false, t: i32struct, u: i8i8struct},
		{want: false, t: i32struct, u: i8i32struct},
		{want: false, t: i32struct, u: i8struct},
		{want: true, t: i32struct, u: i32struct},
		{want: false, t: i32struct, u: emptystruct},
		{want: false, t: emptystruct, u: i8i8struct},
		{want: false, t: emptystruct, u: i8i32struct},
		{want: false, t: emptystruct, u: i8struct},
		{want: false, t: emptystruct, u: i32struct},
		{want: true, t: emptystruct, u: emptystruct},
	}
	for i, g := range golden {
		got := g.t.Equal(g.u)
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
