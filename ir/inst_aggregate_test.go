package ir

import (
	"testing"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func TestTypeCheckInstExtractValue(t *testing.T) {
	structType := types.NewStruct(types.I32, types.I64)

	// Should succeed.
	var v value.Value = constant.NewUndef(structType)
	v.String()
	v = NewInsertValue(v, constant.NewInt(types.I32, 1), 0)
	v.String()
	v = NewInsertValue(v, constant.NewInt(types.I64, 1), 1)
	v.String()

	var panicErr error
	func() {
		defer func() { panicErr = recover().(error) }()
		// Should panic because index 1 is I64, not I32.
		v = NewInsertValue(v, constant.NewInt(types.I32, 1), 1)
		t.Fatal("unreachable")
	}()
	expected := "insertvalue elem type mismatch, expected i64, got i32"
	got := panicErr.Error()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}

	func() {
		defer func() { panicErr = recover().(error) }()
		// Should panic because index 0 is I32, not I64.
		v = NewInsertValue(v, constant.NewInt(types.I64, 1), 0)
		t.Fatal("unreachable")
	}()
	expected = "insertvalue elem type mismatch, expected i32, got i64"
	got = panicErr.Error()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
