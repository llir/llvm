package ast

// A Value represents an LLVM IR value, which may be used as an operand of
// instructions and terminators.
//
// Value may have one of the following underlying types.
//
//    ast.Constant
//    ast.NamedValue
type Value interface {
	// isValue ensures that only values can be assigned to the ast.Value
	// interface.
	isValue()
}

// NamedValue represents a named LLVM IR value.
//
// NamedValue may have one of the following underlying types.
//
//    *ast.Global
//    *ast.GlobalDummy
//    *ast.Function
//    *ast.Param
//    *ast.BasicBlock
//    *ast.LocalDummy
//    ast.Instruction
type NamedValue interface {
	Value
	// GetName returns the name of the value.
	GetName() string
	// SetName sets the name of the value.
	SetName(name string)
}
