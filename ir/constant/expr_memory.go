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
	Typ *types.PointerType
	// Source address element type.
	Elem types.Type
	// Source address.
	Src Constant
	// Element indices.
	Indices []Constant
}

// NewGetElementPtr returns a new getelementptr expression based on the given
// source address and element indices.
func NewGetElementPtr(src Constant, indices ...Constant) *ExprGetElementPtr {
	srcType, ok := src.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid source address type; expected *types.PointerType, got %T", src.Type()))
	}
	elem := srcType.Elem
	e := elem
	for i, index := range indices {
		if i == 0 {
			// Ignore checking the 0th index as it simply follows the pointer of
			// src.
			//
			// ref: http://llvm.org/docs/GetElementPtr.html#why-is-the-extra-0-index-required
			continue
		}
		switch t := e.(type) {
		case *types.PointerType:
			// ref: http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep
			panic("unable to index into element of pointer type; for more information, see http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep")
		case *types.ArrayType:
			e = t.Elem
		case *types.StructType:
			idx, ok := index.(*Int)
			if !ok {
				panic(fmt.Errorf("invalid index type for structure element; expected *constant.Int, got %T", index))
			}
			e = t.Fields[idx.Int64()]
		default:
			panic(fmt.Errorf("support for indexing element type %T not yet implemented", e))
		}
	}
	typ := types.NewPointer(e)
	return &ExprGetElementPtr{
		Typ:     typ,
		Elem:    elem,
		Src:     src,
		Indices: indices,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprGetElementPtr) Type() types.Type {
	return expr.Typ
}

// Ident returns the string representation of the constant expression.
func (expr *ExprGetElementPtr) Ident() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "getelementptr (%s, %s %s",
		expr.Elem,
		expr.Src.Type(),
		expr.Src.Ident())
	for _, index := range expr.Indices {
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

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprGetElementPtr) MetadataNode() {}
