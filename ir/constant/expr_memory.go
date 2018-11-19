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
	Typ types.Type // *types.PointerType or *types.VectorType (with elements of pointer type)
	// (optional) The result is a poison value if the calculated pointer is not
	// an in bounds address of the allocated source object.
	InBounds bool
}

// NewGetElementPtr returns a new getelementptr expression based on the given
// source address and element indices.
func NewGetElementPtr(src Constant, indices ...*Index) *ExprGetElementPtr {
	e := &ExprGetElementPtr{Src: src, Indices: indices}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprGetElementPtr) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprGetElementPtr) Type() types.Type {
	// Cache element type if not present.
	if e.ElemType == nil {
		switch typ := e.Src.Type().(type) {
		case *types.PointerType:
			e.ElemType = typ.ElemType
		case *types.VectorType:
			t, ok := typ.ElemType.(*types.PointerType)
			if !ok {
				panic(fmt.Errorf("invalid vector element type; expected *types.Pointer, got %T", typ.ElemType))
			}
			e.ElemType = t.ElemType
		default:
			panic(fmt.Errorf("support for souce type %T not yet implemented", typ))
		}
	}
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = gepType(e.ElemType, e.Indices)
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprGetElementPtr) Ident() string {
	// 'getelementptr' InBoundsopt '(' ElemType=Type ',' Src=TypeConst
	// Indices=(',' GEPIndex)* ')'
	buf := &strings.Builder{}
	buf.WriteString("getelementptr")
	if e.InBounds {
		buf.WriteString(" inbounds")
	}
	fmt.Fprintf(buf, " (%s, %s", e.ElemType, e.Src)
	for _, index := range e.Indices {
		fmt.Fprintf(buf, ", %s", index)
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
		return fmt.Sprintf("inrange %s", index.Index)
	}
	return index.Index.String()
}

// ### [ Helper functions ] ####################################################

// gepType returns the pointer type or vector of pointers type to the element at
// the position in the type specified by the given indices, as calculated by the
// getelementptr instruction.
func gepType(elemType types.Type, indices []*Index) types.Type {
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
			panic(fmt.Errorf("unable to index into element of pointer type `%s`; for more information, see http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep", elemType))
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
	// TODO: Validate how index vectors in gep are supposed to work.
	//
	// Example from dir.ll:
	//    %113 = getelementptr inbounds %struct.fileinfo, %struct.fileinfo* %96, <2 x i64> %110, !dbg !4736
	//    %116 = bitcast i8** %115 to <2 x %struct.fileinfo*>*, !dbg !4738
	//    store <2 x %struct.fileinfo*> %113, <2 x %struct.fileinfo*>* %116, align 8, !dbg !4738, !tbaa !1793
	if len(indices) > 0 {
		if t, ok := indices[0].Index.Type().(*types.VectorType); ok {
			return types.NewVector(t.Len, types.NewPointer(e))
		}
	}
	return types.NewPointer(e)
}
