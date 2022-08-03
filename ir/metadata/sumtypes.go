package metadata

import (
	"fmt"

	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/value"
)

// TODO: constraint what types may be assigned to Node, MDNode, etc (i.e. make
// them sum types).

// Node is a metadata node.
//
// A Node has one of the following underlying types.
//
//   - [metadata.Definition]
//   - [*metadata.DIExpression]
type Node interface {
	// Ident returns the identifier associated with the metadata node.
	Ident() string
}

// Definition is a metadata definition.
//
// A Definition has one of the following underlying types.
//
//   - [metadata.MDNode]
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
//   - [*metadata.Tuple]
//   - [metadata.Definition]
//   - [metadata.SpecializedNode]
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
//   - [*metadata.NullLit]
//   - [metadata.Metadata]
type Field interface {
	// String returns the LLVM syntax representation of the metadata field.
	fmt.Stringer
}

// SpecializedNode is a specialized metadata node.
//
// A SpecializedNode has one of the following underlying types.
//
//   - [*metadata.DIBasicType]
//   - [*metadata.DICommonBlock]
//   - [*metadata.DICompileUnit]
//   - [*metadata.DICompositeType]
//   - [*metadata.DIDerivedType]
//   - [*metadata.DIEnumerator]
//   - [*metadata.DIExpression]
//   - [*metadata.DIFile]
//   - [*metadata.DIGlobalVariable]
//   - [*metadata.DIGlobalVariableExpression]
//   - [*metadata.DIImportedEntity]
//   - [*metadata.DILabel]
//   - [*metadata.DILexicalBlock]
//   - [*metadata.DILexicalBlockFile]
//   - [*metadata.DILocalVariable]
//   - [*metadata.DILocation]
//   - [*metadata.DIMacro]
//   - [*metadata.DIMacroFile]
//   - [*metadata.DIModule]
//   - [*metadata.DINamespace]
//   - [*metadata.DIObjCProperty]
//   - [*metadata.DISubprogram]
//   - [*metadata.DISubrange]
//   - [*metadata.DISubroutineType]
//   - [*metadata.DITemplateTypeParameter]
//   - [*metadata.DITemplateValueParameter]
//   - [*metadata.GenericDINode]
type SpecializedNode interface {
	Definition
}

// FieldOrInt is a metadata field or integer.
//
// A FieldOrInt has one of the following underlying types.
//
//   - [metadata.Field]
//   - [metadata.IntLit]
type FieldOrInt interface {
	Field
}

// DIExpressionField is a metadata DIExpression field.
//
// A DIExpressionField has one of the following underlying types.
//
//   - [metadata.UintLit]
//   - [enum.DwarfAttEncoding]
//   - [enum.DwarfOp]
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
//   - [value.Value]
//   - [*metadata.String]
//   - [*metadata.Tuple]
//   - [metadata.Definition]
//   - [metadata.SpecializedNode]
//   - [*metadata.DIArgList]
type Metadata interface {
	// String returns the LLVM syntax representation of the metadata.
	fmt.Stringer
}

// NOTE: used for creating godoc links for enum.Foo and value.Foo identifiers.
var (
	_ = enum.DwarfOp(0)
	_ = value.Value(nil)
)
