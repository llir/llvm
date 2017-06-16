// === [ Aggregate expressions ] ===============================================
//
// References:
//    http://llvm.org/docs/LangRef.html#aggregate-operations

package constant

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// --- [ extractvalue ] --------------------------------------------------------

// ExprExtractValue represents an extractvalue expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#extractvalue-instruction
type ExprExtractValue struct {
	// Type of the constant expression.
	Typ types.Type
	// Vector.
	X Constant
	// Indices.
	Indices []int64
}

// NewExtractValue returns a new extractvalue expression based on the given
// vector and indices.
func NewExtractValue(x Constant, indices []int64) *ExprExtractValue {
	typ, err := aggregateElemType(x.Type(), indices)
	if err != nil {
		panic(err)
	}
	return &ExprExtractValue{
		Typ:     typ,
		X:       x,
		Indices: indices,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprExtractValue) Type() types.Type {
	return expr.Typ
}

// Ident returns the string representation of the constant expression.
func (expr *ExprExtractValue) Ident() string {
	buf := &bytes.Buffer{}
	for _, index := range expr.Indices {
		fmt.Fprintf(buf, ", %d", index)
	}
	return fmt.Sprintf("extractvalue (%s %s%s)",
		expr.X.Type(),
		expr.X.Ident(),
		buf)
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprExtractValue) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprExtractValue) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprExtractValue) MetadataNode() {}

// --- [ insertvalue ] ---------------------------------------------------------

// ExprInsertValue represents an insertvalue expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#insertvalue-instruction
type ExprInsertValue struct {
	// Vector.
	X Constant
	// Element to insert.
	Elem Constant
	// Indices.
	Indices []int64
}

// NewInsertValue returns a new insertvalue expression based on the given
// vector, element and indices.
func NewInsertValue(x, elem Constant, indices []int64) *ExprInsertValue {
	return &ExprInsertValue{
		X:       x,
		Elem:    elem,
		Indices: indices,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprInsertValue) Type() types.Type {
	return expr.X.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprInsertValue) Ident() string {
	buf := &bytes.Buffer{}
	for _, index := range expr.Indices {
		fmt.Fprintf(buf, ", %d", index)
	}
	return fmt.Sprintf("insertvalue (%s %s, %s %s%s)",
		expr.X.Type(),
		expr.X.Ident(),
		expr.Elem.Type(),
		expr.Elem.Ident(),
		buf)
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprInsertValue) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprInsertValue) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprInsertValue) MetadataNode() {}

// ### [ Helper functions ] ####################################################

// aggregateElemType returns the element type of the given aggregate type, based
// on the specified indices.
func aggregateElemType(t types.Type, indices []int64) (types.Type, error) {
	if len(indices) == 0 {
		return t, nil
	}
	index := indices[0]
	switch t := t.(type) {
	case *types.ArrayType:
		if index >= t.Len {
			return nil, errors.Errorf("invalid index (%d); exceeds array length (%d)", index, t.Len)
		}
		return aggregateElemType(t.Elem, indices[1:])
	case *types.StructType:
		if index >= int64(len(t.Fields)) {
			return nil, errors.Errorf("invalid index (%d); exceeds struct field count (%d)", index, len(t.Fields))
		}
		return aggregateElemType(t.Fields[index], indices[1:])
	default:
		return nil, errors.Errorf("invalid aggregate value type; expected *types.ArrayType or *types.StructType, got %T", t)
	}
}
