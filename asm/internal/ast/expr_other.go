package ast

import "fmt"

// --- [ icmp ] ----------------------------------------------------------------

// ExprICmp represents an icmp expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#icmp-instruction
type ExprICmp struct {
	// Type of the constant expression.
	Type Type
	// Integer predicate.
	Pred IntPred
	// Operands.
	X, Y Constant
}

// IntPred represents the set of integer predicate of the icmp expression.
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

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*ExprICmp) isValue() {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*ExprICmp) isConstant() {}

// isConstExpr ensures that only constant expressions can be assigned to the
// ast.ConstExpr interface.
func (*ExprICmp) isConstExpr() {}

// --- [ fcmp ] ----------------------------------------------------------------

// ExprFCmp represents an fcmp expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#fcmp-instruction
type ExprFCmp struct {
	// Type of the constant expression.
	Type Type
	// Floating-point predicate.
	Pred FloatPred
	// Operands.
	X, Y Constant
}

// FloatPred represents the set of predicates of the fcmp expression.
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

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*ExprFCmp) isValue() {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*ExprFCmp) isConstant() {}

// isConstExpr ensures that only constant expressions can be assigned to the
// ast.ConstExpr interface.
func (*ExprFCmp) isConstExpr() {}

// --- [ select ] --------------------------------------------------------------

// ExprSelect represents a select expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#select-instruction
type ExprSelect struct {
	// Type of the constant expression.
	Type Type
	// Selection condition.
	Cond Constant
	// Operands.
	X, Y Constant
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*ExprSelect) isValue() {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*ExprSelect) isConstant() {}

// isConstExpr ensures that only constant expressions can be assigned to the
// ast.ConstExpr interface.
func (*ExprSelect) isConstExpr() {}
