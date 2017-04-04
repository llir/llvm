// === [ Aggregate instructions ] ==============================================
//
// References:
//    http://llvm.org/docs/LangRef.html#aggregate-operations

package ast

// --- [ extractvalue ] ------------------------------------------------------

// InstExtractValue represents an extract value instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#extractvalue-instruction
type InstExtractValue struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Aggregate value.
	X Value
	// Indices.
	Indices []int64
}

// GetName returns the name of the value.
func (inst *InstExtractValue) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstExtractValue) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstExtractValue) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstExtractValue) isInst() {}

// --- [ insertvalue ] -------------------------------------------------------

// InstInsertValue represents an insert value instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#insertvalue-instruction
type InstInsertValue struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Aggregate value.
	X Value
	// Element to insert.
	Elem Value
	// Indices.
	Indices []int64
}

// GetName returns the name of the value.
func (inst *InstInsertValue) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstInsertValue) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstInsertValue) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstInsertValue) isInst() {}
