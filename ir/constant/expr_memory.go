// === [ Memory expressions ] ==================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations

package constant

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// --- [ getelementptr ] -------------------------------------------------------

// ExprGetElementPtr represents a getelementptr expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#getelementptr-instruction
type ExprGetElementPtr struct {
	// Type of the constant expression.
	typ types.Type
	// Source address element type.
	elem types.Type
	// Source address.
	src Constant
	// Element indices.
	indices []Constant
}

// NewGetElementPtr returns a new getelementptr expression based on the given
// source address and element indices.
func NewGetElementPtr(src Constant, indices ...Constant) *ExprGetElementPtr {
	srcType, ok := src.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Sprintf("invalid source address type; expected *types.PointerType, got %T", src.Type()))
	}
	elem := srcType.Elem()
	e := elem
	for i, index := range indices {
		if i == 0 {
			// Ignore checking the 0th index as it simply follows the pointer of
			// src.
			//
			// ref: http://llvm.org/docs/GetElementPtr.html#why-is-the-extra-0-index-required
			continue
		}
		if t, ok := e.(*types.NamedType); ok {
			e, ok = t.Def()
			if !ok {
				panic(fmt.Sprintf("invalid named type %q; expected underlying type definition, got nil", t.Name()))
			}
		}
		switch t := e.(type) {
		case *types.PointerType:
			// ref: http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep
			panic("unable to index into element of pointer type; for more information, see http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep")
		case *types.ArrayType:
			e = t.Elem()
		case *types.StructType:
			idx, ok := index.(*Int)
			if !ok {
				panic(fmt.Sprintf("invalid index type for structure element; expected *constant.Int, got %T", index))
			}
			e = t.Fields()[idx.Int64()]
		default:
			panic(fmt.Sprintf("support for indexing element type %T not yet implemented", e))
		}
	}
	typ := types.NewPointer(e)
	return &ExprGetElementPtr{typ: typ, elem: elem, src: src, indices: indices}
}

// Type returns the type of the constant expression.
func (expr *ExprGetElementPtr) Type() types.Type {
	return expr.typ
}

// Ident returns the string representation of the constant expression.
func (expr *ExprGetElementPtr) Ident() string {
	buf := &bytes.Buffer{}
	src := expr.Src()
	fmt.Fprintf(buf, "getelementptr (%s, %s %s",
		expr.elem,
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
func (expr *ExprGetElementPtr) Simplify() Constant {
	panic("not yet implemented")
}

// Src returns the source address of the getelementptr expression.
func (expr *ExprGetElementPtr) Src() Constant {
	return expr.src
}

// SetSrc sets the source address of the getelementptr expression.
func (expr *ExprGetElementPtr) SetSrc(src Constant) {
	expr.src = src
}

// Indices returns the element indices of the getelementptr expression.
func (expr *ExprGetElementPtr) Indices() []Constant {
	return expr.indices
}

// SetIndices sets the element indices of the getelementptr expression.
func (expr *ExprGetElementPtr) SetIndices(indices []Constant) {
	expr.indices = indices
}
