// === [ Other expressions ] ===================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#other-operations

package constant

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// --- [ icmp ] ----------------------------------------------------------------

// ExprICmp represents an icmp expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#icmp-instruction
type ExprICmp struct {
	// Type of the constant expression.
	Typ types.Type
	// Integer predicate.
	Pred IntPred
	// Operands.
	X, Y Constant
}

// NewICmp returns a new icmp expression based on the given integer predicate
// and operands.
func NewICmp(pred IntPred, x, y Constant) *ExprICmp {
	var typ types.Type = types.I1
	if t, ok := x.Type().(*types.VectorType); ok {
		typ = types.NewVector(types.I1, t.Len)
	}
	return &ExprICmp{
		Typ:  typ,
		Pred: pred,
		X:    x,
		Y:    y,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprICmp) Type() types.Type {
	return expr.Typ
}

// Ident returns the string representation of the constant expression.
func (expr *ExprICmp) Ident() string {
	return fmt.Sprintf("icmp %s (%s %s, %s %s)",
		expr.Pred,
		expr.X.Type(),
		expr.X.Ident(),
		expr.Y.Type(),
		expr.Y.Ident())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprICmp) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprICmp) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprICmp) MetadataNode() {}

// IntPred represents the set of integer predicates of the icmp expression.
type IntPred int

// Integer predicates.
const (
	IntEQ  IntPred = iota + 1 // eq: equal
	IntNE                     // ne: not equal
	IntUGT                    // ugt: unsigned greater than
	IntUGE                    // uge: unsigned greater than or equal
	IntULT                    // ult: unsigned less than
	IntULE                    // ule: unsigned less than or equal
	IntSGT                    // sgt: signed greater than
	IntSGE                    // sge: signed greater than or equal
	IntSLT                    // slt: signed less than
	IntSLE                    // sle: signed less than or equal
)

// String returns the LLVM syntax representation of the integer predicate.
func (pred IntPred) String() string {
	m := map[IntPred]string{
		IntEQ:  "eq",
		IntNE:  "ne",
		IntUGT: "ugt",
		IntUGE: "uge",
		IntULT: "ult",
		IntULE: "ule",
		IntSGT: "sgt",
		IntSGE: "sge",
		IntSLT: "slt",
		IntSLE: "sle",
	}
	if s, ok := m[pred]; ok {
		return s
	}
	return fmt.Sprintf("<unknown integer predicate %d>", int(pred))
}

// --- [ fcmp ] ----------------------------------------------------------------

// ExprFCmp represents an fcmp expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#fcmp-instruction
type ExprFCmp struct {
	// Type of the constant expression.
	Typ types.Type
	// Floating-point predicate.
	Pred FloatPred
	// Operands.
	X, Y Constant
}

// NewFCmp returns a new fcmp expression based on the given floating-point
// predicate and operands.
func NewFCmp(pred FloatPred, x, y Constant) *ExprFCmp {
	var typ types.Type = types.I1
	if t, ok := x.Type().(*types.VectorType); ok {
		typ = types.NewVector(types.I1, t.Len)
	}
	return &ExprFCmp{
		Typ:  typ,
		Pred: pred,
		X:    x,
		Y:    y,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprFCmp) Type() types.Type {
	return expr.Typ
}

// Ident returns the string representation of the constant expression.
func (expr *ExprFCmp) Ident() string {
	return fmt.Sprintf("fcmp %s (%s %s, %s %s)",
		expr.Pred,
		expr.X.Type(),
		expr.X.Ident(),
		expr.Y.Type(),
		expr.Y.Ident())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFCmp) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprFCmp) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprFCmp) MetadataNode() {}

// FloatPred represents the set of floating-point predicates of the fcmp
// expression.
type FloatPred int

// Floating-point predicates.
const (
	FloatFalse FloatPred = iota + 1 // false: no comparison, always returns false
	FloatOEQ                        // oeq: ordered and equal
	FloatOGT                        // ogt: ordered and greater than
	FloatOGE                        // oge: ordered and greater than or equal
	FloatOLT                        // olt: ordered and less than
	FloatOLE                        // ole: ordered and less than or equal
	FloatONE                        // one: ordered and not equal
	FloatORD                        // ord: ordered (no nans)
	FloatUEQ                        // ueq: unordered or equal
	FloatUGT                        // ugt: unordered or greater than
	FloatUGE                        // uge: unordered or greater than or equal
	FloatULT                        // ult: unordered or less than
	FloatULE                        // ule: unordered or less than or equal
	FloatUNE                        // une: unordered or not equal
	FloatUNO                        // uno: unordered (either nans)
	FloatTrue                       // true: no comparison, always returns true
)

// String returns the LLVM syntax representation of the floating-point
// predicate.
func (pred FloatPred) String() string {
	m := map[FloatPred]string{
		FloatFalse: "false",
		FloatOEQ:   "oeq",
		FloatOGT:   "ogt",
		FloatOGE:   "oge",
		FloatOLT:   "olt",
		FloatOLE:   "ole",
		FloatONE:   "one",
		FloatORD:   "ord",
		FloatUEQ:   "ueq",
		FloatUGT:   "ugt",
		FloatUGE:   "uge",
		FloatULT:   "ult",
		FloatULE:   "ule",
		FloatUNE:   "une",
		FloatUNO:   "uno",
		FloatTrue:  "true",
	}
	if s, ok := m[pred]; ok {
		return s
	}
	return fmt.Sprintf("<unknown floating-point predicate %d>", int(pred))
}

// --- [ select ] --------------------------------------------------------------

// ExprSelect represents a select expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#select-instruction
type ExprSelect struct {
	// Selection condition.
	Cond Constant
	// Operands.
	X, Y Constant
}

// NewSelect returns a new select expression based on the given selection
// condition and operands.
func NewSelect(cond, x, y Constant) *ExprSelect {
	return &ExprSelect{
		Cond: cond,
		X:    x,
		Y:    y,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprSelect) Type() types.Type {
	return expr.X.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprSelect) Ident() string {
	return fmt.Sprintf("select (%s %s, %s %s, %s %s)",
		expr.Cond.Type(),
		expr.Cond.Ident(),
		expr.X.Type(),
		expr.X.Ident(),
		expr.Y.Type(),
		expr.Y.Ident())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSelect) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprSelect) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprSelect) MetadataNode() {}
