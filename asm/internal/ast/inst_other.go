package ast

// --- [ icmp ] ----------------------------------------------------------------

// InstICmp represents an icmp instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#icmp-instruction
type InstICmp struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Type Type
	// Integer condition code.
	Cond IntPred
	// Operands.
	X, Y Value
}

// --- [ fcmp ] ----------------------------------------------------------------

// InstFCmp represents an fcmp instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fcmp-instruction
type InstFCmp struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Type Type
	// Floating-point condition code.
	Cond FloatPred
	// Operands.
	X, Y Value
}

// --- [ phi ] -----------------------------------------------------------------

// InstPhi represents a phi instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#phi-instruction
type InstPhi struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Type Type
	// Incoming values.
	Incs []*Incoming
}

// Incoming represents an incoming value of a phi instruction.
type Incoming struct {
	// Incoming value.
	//
	// Initially holds *astx.IntLit, *astx.LocalIdent, ... when created from
	// astx.NewIncoming since the type is not yet known. The astx.NewPhiInst
	// later replaces this with a value (e.g. *ast.IntConst, *ast.LocalDummy,
	// ...).
	X interface{}
	// Predecessor basic block of the incoming value.
	Pred string
}

// --- [ select ] --------------------------------------------------------------

// InstSelect represents a select instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#select-instruction
type InstSelect struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Type Type
	// Selection condition.
	Cond Value
	// Operands.
	X, Y Value
}

// --- [ call ] ----------------------------------------------------------------

// InstCall represents a call instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#call-instruction
type InstCall struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Type Type
	// Callee.
	Callee string
	// Specifies whether the callee is a local identifier.
	CalleeLocal bool
	// Callee signature.
	Sig Type
	// Function arguments.
	Args []Value
}

// --- [ va_arg ] --------------------------------------------------------------

// --- [ landingpad ] ----------------------------------------------------------

// --- [ catchpad ] ------------------------------------------------------------

// --- [ cleanuppad ] ----------------------------------------------------------

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstICmp) isValue()   {}
func (*InstFCmp) isValue()   {}
func (*InstPhi) isValue()    {}
func (*InstSelect) isValue() {}
func (*InstCall) isValue()   {}

// isNamedValue ensures that only named values can be assigned to the
// ast.NamedValue interface.
func (*InstICmp) isNamedValue()   {}
func (*InstFCmp) isNamedValue()   {}
func (*InstPhi) isNamedValue()    {}
func (*InstSelect) isNamedValue() {}
func (*InstCall) isNamedValue()   {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstICmp) isInst()   {}
func (*InstFCmp) isInst()   {}
func (*InstPhi) isInst()    {}
func (*InstSelect) isInst() {}
func (*InstCall) isInst()   {}
