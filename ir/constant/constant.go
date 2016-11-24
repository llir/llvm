// === [ Constants ] ===========================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#constants

// Package constant implements values representing immutable LLVM IR constants.
package constant

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// A Constant represents an LLVM IR constant; a value that is immutable at
// runtime, such as an integer or a floating point literal.
//
// Constant may have one of the following underlying types.
//
//    TODO
type Constant interface {
	value.Value
}

// --- [ integer ] -------------------------------------------------------------

// Int represents an integer constant.
type Int struct {
	// Constant value.
	x int64
	// Constant type.
	typ *types.IntType
}

// NewInt returns a new integer constant of the given value and type.
func NewInt(x int64, typ types.Type) *Int {
	if typ, ok := typ.(*types.IntType); ok {
		return &Int{x: x, typ: typ}
	}
	panic(fmt.Sprintf("invalid integer constant type; expected *types.IntType, got %T", typ))
}

// Type returns the type of the constant.
func (c *Int) Type() types.Type {
	return c.typ
}

// Ident returns the value of the constant.
func (c *Int) Ident() string {
	if c.typ.Size() == 1 {
		switch c.x {
		case 0:
			return "false"
		case 1:
			return "true"
		default:
			panic(fmt.Sprintf("invalid integer constant value; expected 0 or 1, got %d", c.x))
		}
	}
	return fmt.Sprintf("%d", c.x)
}

// X returns the value of the integer constant.
func (c *Int) X() int64 {
	return c.x
}

// --- [ floating-point ] ------------------------------------------------------

// Float represents a floating-point constant.
type Float struct {
	// Constant value.
	x float64
	// Constant type.
	typ types.Type
}

// NewFloat returns a new floating-point constant of the given value and type.
func NewFloat(x float64, typ types.Type) *Float {
	return &Float{x: x, typ: typ}
}

// Type returns the type of the constant.
func (c *Float) Type() types.Type {
	return c.typ
}

// Ident returns the value of the constant.
func (c *Float) Ident() string {
	return fmt.Sprintf("%g", c.x)
}

// X returns the value of the integer constant.
func (c *Float) X() float64 {
	return c.x
}
