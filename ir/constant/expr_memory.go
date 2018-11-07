package constant

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/types"
)

// --- [ Memory expressions ] --------------------------------------------------

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprGetElementPtr is an LLVM IR getelementptr expression.
type ExprGetElementPtr struct {
	// Element type.
	ElemType types.Type
	// Source address.
	Src Constant
	// Element indicies.
	Indices []*Index

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
	// (optional) The result is a poison value if the calculated pointer is not
	// an in bounds address of the allocated source object.
	InBounds bool
}

// NewGetElementPtr returns a new getelementptr expression based on the given
// element type, source address and element indices.
func NewGetElementPtr(elemType types.Type, src Constant, indices ...*Index) *ExprGetElementPtr {
	return &ExprGetElementPtr{ElemType: elemType, Src: src, Indices: indices}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprGetElementPtr) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprGetElementPtr) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = gepType(e.ElemType, e.Indices)
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprGetElementPtr) Ident() string {
	// "getelementptr" OptInBounds "(" Type "," Type Constant "," GEPConstIndices ")"
	buf := &strings.Builder{}
	buf.WriteString("getelementptr")
	if e.InBounds {
		buf.WriteString(" inbounds")
	}
	fmt.Fprintf(buf, " (%v, %v", e.ElemType, e.Src)
	for _, index := range e.Indices {
		fmt.Fprintf(buf, ", %v", index)
	}
	buf.WriteString(")")
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprGetElementPtr) Simplify() Constant {
	panic("not yet implemented")
}

// ___ [ gep indices ] _________________________________________________________

// Index is an index of a getelementptr constant expression.
type Index struct {
	// Element index.
	Index Constant

	// extra.

	// (optional) States that the element index is not out the bounds of the
	// allocated object. If inrange is stated but the element index is out of
	// bounds, the behaviour is undefined.
	InRange bool
}

// NewIndex returns a new gep element index.
func NewIndex(index Constant) *Index {
	return &Index{Index: index}
}

// String returns a string representation of the getelementptr index.
func (index *Index) String() string {
	// OptInrange Type Constant
	if index.InRange {
		return fmt.Sprintf("inrange %v", index.Index)
	}
	return index.Index.String()
}

// ### [ Helper functions ] ####################################################

// gepType returns the pointer type to the element at the position in the type
// specified by the given indices, as calculated by the getelementptr
// instruction.
func gepType(elemType types.Type, indices []*Index) *types.PointerType {
	e := elemType
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
			panic(fmt.Errorf("unable to index into element of pointer type `%v`; for more information, see http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep", elemType))
		case *types.VectorType:
			e = t.ElemType
		case *types.ArrayType:
			e = t.ElemType
		case *types.StructType:
			idx, ok := index.Index.(*Int)
			if !ok {
				panic(fmt.Errorf("invalid index type for structure element; expected *constant.Int, got %T", index))
			}
			e = t.Fields[idx.X.Int64()]
		default:
			panic(fmt.Errorf("support for indexing element type %T not yet implemented", e))
		}
	}
	return types.NewPointer(e)
}
