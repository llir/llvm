package metadata

import "fmt"

// TODO: constraint what types may be assigned to Node, MDNode, etc (i.e. make
// them sum types).

// Node is a metadata node.
//
// A Node has one of the following underlying types.
//
//    metadata.Definition      // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#Definition
//    *metadata.DIExpression   // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIExpression
type Node interface {
	// Ident returns the identifier associated with the metadata node.
	Ident() string
}

// Definition is a metadata definition.
//
// A Definition has one of the following underlying types.
//
//    metadata.MDNode   // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#MDNode
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
	// SetDistinct specifies whether the metadata definition is dinstict.
	SetDistinct(distinct bool)
}

// MDNode is a metadata node.
//
// A MDNode has one of the following underlying types.
//
//    *metadata.Tuple            // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#Tuple
//    metadata.Definition        // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#Definition
//    metadata.SpecializedNode   // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#SpecializedNode
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
//    *metadata.NullLit   // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#NullLit
//    metadata.Metadata   // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#Metadata
type Field interface {
	// String returns the LLVM syntax representation of the metadata field.
	fmt.Stringer
}

// SpecializedNode is a specialized metadata node.
//
// A SpecializedNode has one of the following underlying types.
//
//    *metadata.DIBasicType                  // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIBasicType
//    *metadata.DICommonBlock                // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DICommonBlock
//    *metadata.DICompileUnit                // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DICompileUnit
//    *metadata.DICompositeType              // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DICompositeType
//    *metadata.DIDerivedType                // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIDerivedType
//    *metadata.DIEnumerator                 // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIEnumerator
//    *metadata.DIExpression                 // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIExpression
//    *metadata.DIFile                       // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIFile
//    *metadata.DIGlobalVariable             // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIGlobalVariable
//    *metadata.DIGlobalVariableExpression   // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIGlobalVariableExpression
//    *metadata.DIImportedEntity             // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIImportedEntity
//    *metadata.DILabel                      // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DILabel
//    *metadata.DILexicalBlock               // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DILexicalBlock
//    *metadata.DILexicalBlockFile           // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DILexicalBlockFile
//    *metadata.DILocalVariable              // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DILocalVariable
//    *metadata.DILocation                   // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DILocation
//    *metadata.DIMacro                      // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIMacro
//    *metadata.DIMacroFile                  // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIMacroFile
//    *metadata.DIModule                     // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIModule
//    *metadata.DINamespace                  // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DINamespace
//    *metadata.DIObjCProperty               // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIObjCProperty
//    *metadata.DISubprogram                 // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DISubprogram
//    *metadata.DISubrange                   // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DISubrange
//    *metadata.DISubroutineType             // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DISubroutineType
//    *metadata.DITemplateTypeParameter      // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DITemplateTypeParameter
//    *metadata.DITemplateValueParameter     // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DITemplateValueParameter
//    *metadata.GenericDINode                // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#GenericDINode
type SpecializedNode interface {
	Definition
}

// FieldOrInt is a metadata field or integer.
//
// A FieldOrInt has one of the following underlying types.
//
//    metadata.Field    // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#Field
//    metadata.IntLit   // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#IntLit
type FieldOrInt interface {
	Field
}

// DIExpressionField is a metadata DIExpression field.
//
// A DIExpressionField has one of the following underlying types.
//
//    metadata.UintLit        // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#UintLit
//    enum.DwarfAttEncoding   // https://pkg.go.dev/github.com/llir/llvm/ir/enum#DwarfAttEncoding
//    enum.DwarfOp            // https://pkg.go.dev/github.com/llir/llvm/ir/enum#DwarfOp
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
//    value.Value                // https://pkg.go.dev/github.com/llir/llvm/ir/value#Value
//    *metadata.String           // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#String
//    *metadata.Tuple            // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#Tuple
//    metadata.Definition        // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#Definition
//    metadata.SpecializedNode   // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#SpecializedNode
//    *metadata.DIArgList        // https://pkg.go.dev/github.com/llir/llvm/ir/metadata#DIArgList
type Metadata interface {
	// String returns the LLVM syntax representation of the metadata.
	fmt.Stringer
}
