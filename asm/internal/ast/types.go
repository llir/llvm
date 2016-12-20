package ast

// A Type represents an LLVM IR type.
//
// Type may have one of the following underlying types.
//
//    *ast.VoidType
//    *ast.LabelType
//    *ast.IntType
//    *ast.FloatType
//    *ast.FuncType
//    *ast.PointerType
//    *ast.VectorType
//    *ast.ArrayType
//    *ast.StructType
//    *ast.NamedType
type Type interface {
	// isType ensures that only types can be assigned to the ast.Type interface.
	isType()
}

// TypeDummy represents a dummy type; used when a type is unknown during
// parsing.
type TypeDummy struct {
}

// isType ensures that only types can be assigned to the ast.Type interface.
func (*TypeDummy) isType() {}
