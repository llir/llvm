// === [ Memory expressions ] ==================================================

package dummyconstant

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// --- [ getelementptr ] -------------------------------------------------------

// ExprGetElementPtr represents a dummy getelementptr expression.
type ExprGetElementPtr struct {
	// Source address element type.
	elem types.Type
	// Source address.
	src constant.Constant
	// Element indices.
	indices []constant.Constant
	// Track uses of the value.
	used
}

// NewGetElementPtr returns a dummy new getelementptr expression based on the
// given source address element type, source address and element indices.
func NewGetElementPtr(elem types.Type, src constant.Constant, indices ...constant.Constant) *ExprGetElementPtr {
	return &ExprGetElementPtr{elem: elem, src: src, indices: indices}
}

// Type returns the type of the constant expression.
func (expr *ExprGetElementPtr) Type() types.Type {
	panic("dummy implementation")
}

// Ident returns the string representation of the constant expression.
func (expr *ExprGetElementPtr) Ident() string {
	buf := &bytes.Buffer{}
	src := expr.Src()
	fmt.Fprintf(buf, "getelementptr (%s, %s %s",
		expr.ElemType(),
		src.Type(),
		src.Ident())
	for _, index := range expr.Indices() {
		fmt.Fprintf(buf, ", %s %s",
			index.Type(),
			index.Ident())
	}
	buf.WriteString(")")
	return buf.String()
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprGetElementPtr) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprGetElementPtr) Simplify() constant.Constant {
	panic("not yet implemented")
}

// ElemType returns the source address element type of the getelementptr
// expression.
func (expr *ExprGetElementPtr) ElemType() types.Type {
	return expr.elem
}

// Src returns the source address of the getelementptr expression.
func (expr *ExprGetElementPtr) Src() constant.Constant {
	return expr.src
}

// Indices returns the element indices of the getelementptr expression.
func (expr *ExprGetElementPtr) Indices() []constant.Constant {
	return expr.indices
}
