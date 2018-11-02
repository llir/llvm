// === [ Vector instructions ] =================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#vector-operations

package ast

// --- [ extractelement ] ------------------------------------------------------

// InstExtractElement represents an extractelement instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#extractelement-instruction
type InstExtractElement struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Vector.
	X Value
	// Index.
	Index Value
	// Metadata attached to the instruction.
	Metadata []*AttachedMD
}

// GetName returns the name of the value.
func (inst *InstExtractElement) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstExtractElement) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstExtractElement) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstExtractElement) isInst() {}

// --- [ insertelement ] -------------------------------------------------------

// InstInsertElement represents an insertelement instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#insertelement-instruction
type InstInsertElement struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Vector.
	X Value
	// Element to insert.
	Elem Value
	// Index.
	Index Value
	// Metadata attached to the instruction.
	Metadata []*AttachedMD
}

// GetName returns the name of the value.
func (inst *InstInsertElement) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstInsertElement) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstInsertElement) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstInsertElement) isInst() {}

// --- [ shufflevector ] -------------------------------------------------------

// InstShuffleVector represents a shufflevector instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#shufflevector-instruction
type InstShuffleVector struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Vector 1.
	X Value
	// Vector 2.
	Y Value
	// Shuffle mask.
	Mask Value
	// Metadata attached to the instruction.
	Metadata []*AttachedMD
}

// GetName returns the name of the value.
func (inst *InstShuffleVector) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstShuffleVector) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstShuffleVector) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstShuffleVector) isInst() {}
