package ast

// --- [ alloca ] --------------------------------------------------------------

// InstAlloca represents an alloca instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#alloca-instruction
type InstAlloca struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Type Type
	// Element type.
	Elem Type
	// Number of elements; or nil if one element.
	NElems Value
}

// GetName returns the name of the value.
func (inst *InstAlloca) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstAlloca) SetName(name string) {
	inst.Name = name
}

// --- [ load ] ----------------------------------------------------------------

// InstLoad represents a load instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#load-instruction
type InstLoad struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Type Type
	// Element type.
	Elem Type
	// Source address.
	Src Value
}

// GetName returns the name of the value.
func (inst *InstLoad) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstLoad) SetName(name string) {
	inst.Name = name
}

// --- [ store ] ---------------------------------------------------------------

// InstStore represents a store instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#store-instruction
type InstStore struct {
	// Source value.
	Src Value
	// Destination address.
	Dst Value
}

// --- [ fence ] ---------------------------------------------------------------

// --- [ cmpxchg ] -------------------------------------------------------------

// --- [ atomicrmw ] -----------------------------------------------------------

// --- [ getelementptr ] -------------------------------------------------------

// InstGetElementPtr represents a getelementptr instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#getelementptr-instruction
type InstGetElementPtr struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Type Type
	// Source address element type.
	Elem Type
	// Source address.
	Src Value
	// Element indices.
	Indices []Value
}

// GetName returns the name of the value.
func (inst *InstGetElementPtr) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstGetElementPtr) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstAlloca) isValue()        {}
func (*InstLoad) isValue()          {}
func (*InstStore) isValue()         {}
func (*InstGetElementPtr) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstAlloca) isInst()        {}
func (*InstLoad) isInst()          {}
func (*InstStore) isInst()         {}
func (*InstGetElementPtr) isInst() {}
