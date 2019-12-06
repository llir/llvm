package constant

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/gep"
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
	Indices []Constant // *Int, *Vector or *Index

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type // *types.PointerType or *types.VectorType (with elements of pointer type)
	// (optional) The result is a poison value if the calculated pointer is not
	// an in bounds address of the allocated source object.
	InBounds bool
}

// TODO: re-work NewGetElementPtr to take elemType as argument, as this is
// really the type used to compute the result type of gep.

// NewGetElementPtr returns a new getelementptr expression based on the given
// source address and element indices.
func NewGetElementPtr(src Constant, indices ...Constant) *ExprGetElementPtr {
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
	// TODO: remove e.ElemType computation once NewGetElementPtr takes elemType
	// as argument.
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
			panic(fmt.Errorf("support for source type %T not yet implemented", typ))
		}
	}
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = gepExprType(e.ElemType, e.Src.Type(), e.Indices)
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
	Constant

	// extra.

	// (optional) States that the element index is not out the bounds of the
	// allocated object. If inrange is stated but the element index is out of
	// bounds, the behaviour is undefined.
	InRange bool
}

// NewIndex returns a new gep element index.
func NewIndex(index Constant) *Index {
	return &Index{Constant: index}
}

// String returns a string representation of the getelementptr index.
func (index *Index) String() string {
	// OptInrange Type Constant
	if index.InRange {
		return fmt.Sprintf("inrange %s", index.Constant)
	}
	return index.Constant.String()
}

// ### [ Helper functions ] ####################################################

// gepExprType computes the result type of a getelementptr constant expression.
//
//    getelementptr (ElemType, Src, Indices)
func gepExprType(elemType, src types.Type, indices []Constant) types.Type {
	var idxs []gep.Index
	for _, index := range indices {
		idx := getIntValue(index)
		idxs = append(idxs, idx)
	}
	return gep.ResultType(elemType, src, idxs)
}

// getIntValue returns the concrete integer value of the given index.
func getIntValue(index Constant) gep.Index {
	// unpack inrange indices.
	if idx, ok := index.(*Index); ok {
		index = idx.Constant
	}
	// TODO: use index.Simplify() to simplify the constant expression to a
	// concrete integer constant.
	switch index := index.(type) {
	case *Int:
		val := index.X.Int64()
		return gep.NewIndex(val)
	case *ZeroInitializer:
		return gep.NewIndex(0)
	case *Vector:
		// ref: https://llvm.org/docs/LangRef.html#getelementptr-instruction
		//
		// > The getelementptr returns a vector of pointers, instead of a single
		// > address, when one or more of its arguments is a vector. In such
		// > cases, all vector arguments should have the same number of elements,
		// > and every scalar argument will be effectively broadcast into a vector
		// > during address calculation.

		// Sanity check. All vector elements must be integers, and must have the
		// same value.
		var val int64
		if len(index.Elems) < 1 {
			return gep.Index{HasVal: false}
		}
		for i, elem := range index.Elems {
			switch elem := elem.(type) {
			case *Int:
				x := elem.X.Int64()
				if i == 0 {
					val = x
				} else if x != val {
					// since all elements were not identical, we must conclude that
					// the index vector does not have a concrete value.
					return gep.Index{HasVal: false}
				}
			default:
				// TODO: remove debug output.
				panic(fmt.Errorf("support for gep index vector element type %T not yet implemented", elem))
				return gep.Index{HasVal: false}
			}
		}
		return gep.Index{
			HasVal:    true,
			Val:       val,
			VectorLen: uint64(len(index.Elems)),
		}
	case *ExprPtrToInt:
		return gep.Index{HasVal: false}
	default:
		// TODO: add support for more constant expressions.
		// TODO: remove debug output.
		panic(fmt.Errorf("support for gep index type %T not yet implemented", index))
		return gep.Index{HasVal: false}
	}
}

// Example from dir.ll:
//    %113 = getelementptr inbounds %struct.fileinfo, %struct.fileinfo* %96, <2 x i64> %110, !dbg !4736
//    %116 = bitcast i8** %115 to <2 x %struct.fileinfo*>*, !dbg !4738
//    store <2 x %struct.fileinfo*> %113, <2 x %struct.fileinfo*>* %116, align 8, !dbg !4738, !tbaa !1793
