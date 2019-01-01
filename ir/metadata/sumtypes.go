package metadata

import "fmt"

// TODO: constraint what types may be assigned to Node, MDNode, etc (i.e. make
// them sum types).

// Node is a metadata node.
//
// A Node has one of the following underlying types.
//
//    *metadata.Def            // https://godoc.org/github.com/llir/llvm/ir/metadata#Def
//    *metadata.DIExpression   // https://godoc.org/github.com/llir/llvm/ir/metadata#DIExpression
type Node interface {
	// Ident returns the identifier associated with the metadata node.
	Ident() string
}

// Definition is a metadata definition.
//
// A Definition has one of the following underlying types.
//
//    *metadata.Def     // https://godoc.org/github.com/llir/llvm/ir/metadata#Def
//    metadata.MDNode   // https://godoc.org/github.com/llir/llvm/ir/metadata#MDNode
type Definition interface {
	// String returns the LLVM syntax representation of the metadata.
	fmt.Stringer
	// Ident returns the identifier associated with the metadata definition.
	Ident() string
	// ID returns the ID of the metadata definition.
	ID() int64
	// SetID sets the ID of the metadata definition.
	SetID(id int64)
	// LLString returns the LLVM syntax representation of the metadata
	// definition.
	LLString() string
}

// MDNode is a metadata node.
//
// A MDNode has one of the following underlying types.
//
//    *metadata.Tuple            // https://godoc.org/github.com/llir/llvm/ir/metadata#Tuple
//    *metadata.Def              // https://godoc.org/github.com/llir/llvm/ir/metadata#Def
//    metadata.SpecializedNode   // https://godoc.org/github.com/llir/llvm/ir/metadata#SpecializedNode
type MDNode interface {
	// Ident returns the identifier associated with the metadata node.
	Ident() string
	// LLString returns the LLVM syntax representation of the metadata node.
	LLString() string
}

// Field is a metadata field.
//
// A Field has one of the following underlying types.
//
//    *metadata.NullLit   // https://godoc.org/github.com/llir/llvm/ir/metadata#NullLit
//    metadata.Metadata   // https://godoc.org/github.com/llir/llvm/ir/metadata#Metadata
type Field interface {
	// String returns the LLVM syntax representation of the metadata field.
	fmt.Stringer
}

// SpecializedNode is a specialized metadata node.
//
// A SpecializedNode has one of the following underlying types.
//
//    *metadata.DIBasicType                  // https://godoc.org/github.com/llir/llvm/ir/metadata#DIBasicType
//    *metadata.DICompileUnit                // https://godoc.org/github.com/llir/llvm/ir/metadata#DICompileUnit
//    *metadata.DICompositeType              // https://godoc.org/github.com/llir/llvm/ir/metadata#DICompositeType
//    *metadata.DIDerivedType                // https://godoc.org/github.com/llir/llvm/ir/metadata#DIDerivedType
//    *metadata.DIEnumerator                 // https://godoc.org/github.com/llir/llvm/ir/metadata#DIEnumerator
//    *metadata.DIExpression                 // https://godoc.org/github.com/llir/llvm/ir/metadata#DIExpression
//    *metadata.DIFile                       // https://godoc.org/github.com/llir/llvm/ir/metadata#DIFile
//    *metadata.DIGlobalVariable             // https://godoc.org/github.com/llir/llvm/ir/metadata#DIGlobalVariable
//    *metadata.DIGlobalVariableExpression   // https://godoc.org/github.com/llir/llvm/ir/metadata#DIGlobalVariableExpression
//    *metadata.DIImportedEntity             // https://godoc.org/github.com/llir/llvm/ir/metadata#DIImportedEntity
//    *metadata.DILabel                      // https://godoc.org/github.com/llir/llvm/ir/metadata#DILabel
//    *metadata.DILexicalBlock               // https://godoc.org/github.com/llir/llvm/ir/metadata#DILexicalBlock
//    *metadata.DILexicalBlockFile           // https://godoc.org/github.com/llir/llvm/ir/metadata#DILexicalBlockFile
//    *metadata.DILocalVariable              // https://godoc.org/github.com/llir/llvm/ir/metadata#DILocalVariable
//    *metadata.DILocation                   // https://godoc.org/github.com/llir/llvm/ir/metadata#DILocation
//    *metadata.DIMacro                      // https://godoc.org/github.com/llir/llvm/ir/metadata#DIMacro
//    *metadata.DIMacroFile                  // https://godoc.org/github.com/llir/llvm/ir/metadata#DIMacroFile
//    *metadata.DIModule                     // https://godoc.org/github.com/llir/llvm/ir/metadata#DIModule
//    *metadata.DINamespace                  // https://godoc.org/github.com/llir/llvm/ir/metadata#DINamespace
//    *metadata.DIObjCProperty               // https://godoc.org/github.com/llir/llvm/ir/metadata#DIObjCProperty
//    *metadata.DISubprogram                 // https://godoc.org/github.com/llir/llvm/ir/metadata#DISubprogram
//    *metadata.DISubrange                   // https://godoc.org/github.com/llir/llvm/ir/metadata#DISubrange
//    *metadata.DISubroutineType             // https://godoc.org/github.com/llir/llvm/ir/metadata#DISubroutineType
//    *metadata.DITemplateTypeParameter      // https://godoc.org/github.com/llir/llvm/ir/metadata#DITemplateTypeParameter
//    *metadata.DITemplateValueParameter     // https://godoc.org/github.com/llir/llvm/ir/metadata#DITemplateValueParameter
//    *metadata.GenericDINode                // https://godoc.org/github.com/llir/llvm/ir/metadata#GenericDINode
type SpecializedNode interface {
	Definition
}

// FieldOrInt is a metadata field or integer.
//
// A FieldOrInt has one of the following underlying types.
//
//    metadata.Field    // https://godoc.org/github.com/llir/llvm/ir/metadata#Field
//    metadata.IntLit   // https://godoc.org/github.com/llir/llvm/ir/metadata#IntLit
type FieldOrInt interface {
	Field
}

// DIExpressionField is a metadata DIExpression field.
//
// A DIExpressionField has one of the following underlying types.
//
//    metadata.UintLit   // https://godoc.org/github.com/llir/llvm/ir/metadata#UintLit
//    enum.DwarfOp       // https://godoc.org/github.com/llir/llvm/ir/enum#DwarfOp
type DIExpressionField interface {
	fmt.Stringer
	// IsDIExpressionField ensures that only DIExpression fields can be assigned
	// to the metadata.DIExpressionField interface.
	IsDIExpressionField()
}

// IsDIExpressionField ensures that only DIExpression fields can be assigned to
// the metadata.DIExpressionField interface.
func (UintLit) IsDIExpressionField() {}

// Metadata is a sumtype of metadata.
//
// A Metadata has one of the following underlying types.
//
//    value.Value                // https://godoc.org/github.com/llir/llvm/ir/value#Value
//    *metadata.String           // https://godoc.org/github.com/llir/llvm/ir/metadata#String
//    *metadata.Tuple            // https://godoc.org/github.com/llir/llvm/ir/metadata#Tuple
//    *metadata.Def              // https://godoc.org/github.com/llir/llvm/ir/metadata#Def
//    metadata.SpecializedNode   // https://godoc.org/github.com/llir/llvm/ir/metadata#SpecializedNode
type Metadata interface {
	// String returns the LLVM syntax representation of the metadata.
	fmt.Stringer
}
