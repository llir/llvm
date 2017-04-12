// === [ Metadata ] ============================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#metadata

package ast

// A MetadataNode represents an LLVM IR metadata node.
//
// MetadataNode may have one of the following underlying types.
//
//    *ast.Metadata
//    *ast.MetadataString
//    ast.Constant
type MetadataNode interface {
	Value
	// isMetadataNode ensures that only metadata nodes can be assigned to the
	// ast.MetadataNode interface.
	isMetadataNode()
}

// --- [ metadata ] ------------------------------------------------------------

// Metadata represents a set of LLVM IR metadata nodes.
//
// Metadata may be referenced from instructions (e.g. call), and are thus
// considered LLVM IR values of metadata type.
type Metadata struct {
	// Metadata ID; or empty if metadata literal.
	ID string
	// Metadata nodes.
	Nodes []MetadataNode
}

// --- [ metadata string ] -----------------------------------------------------

// A MetadataString represents an LLVM IR metadata string.
type MetadataString struct {
	// String value.
	Val string
}

// --- [ named metadata ] ------------------------------------------------------

// NamedMetadata represents a named collection of metadata, which belongs to a
// module.
type NamedMetadata struct {
	// Metadata name.
	Name string
	// Associated metadata; initially *ast.MetadataIDDummy and replaced with
	// corresponding *ast.Metadata by astx.fixModule.
	Metadata []MetadataNode
}

// AttachedMD represents attached metadata.
type AttachedMD struct {
	// Name associated with the attached metadata (e.g. !dbg).
	Name string
	// Metadata; may be *ast.MetadataIDDummy or *ast.Metadata during translation,
	// *ast.MetadataIDDummy are later replaced with corresponding *ast.Metadata
	// by astx.fixModule.
	Metadata MetadataNode
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*Metadata) isValue()       {}
func (*MetadataString) isValue() {}

// isMetadataNode ensures that only metadata nodes can be assigned to the
// ast.MetadataNode interface.
func (*Metadata) isMetadataNode()       {}
func (*MetadataString) isMetadataNode() {}

// ### [ dummy ] ###############################################################

// MetadataIDDummy represents a dummy metadata ID.
type MetadataIDDummy struct {
	// Metadata ID.
	ID string
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*MetadataIDDummy) isValue() {}

// isMetadataNode ensures that only metadata nodes can be assigned to the
// ast.MetadataNode interface.
func (*MetadataIDDummy) isMetadataNode() {}
