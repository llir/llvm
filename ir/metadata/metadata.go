// Package metadata provides access to LLVM IR metadata.
package metadata

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/types"
)

// ~~~ [ Named Metadata Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NamedMetadataDef is a named metadata definition.
type NamedMetadataDef struct {
	// Metadata definition name (without '!' prefix).
	Name string
	// Metadata definition nodes.
	Nodes []MetadataNode
}

// String returns the string representation of the named metadata definition.
func (md *NamedMetadataDef) String() string {
	return enc.Metadata(md.Name)
}

// Def returns the LLVM syntax representation of the named metadata definition.
func (md *NamedMetadataDef) Def() string {
	// MetadataName "=" "!" "{" MetadataNodes "}"
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = !{", enc.Metadata(md.Name))
	for i, node := range md.Nodes {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(node.String())
	}
	buf.WriteString("}")
	return buf.String()
}

// ~~~ [ Metadata Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// MetadataDef is a metadata definition.
type MetadataDef struct {
	// Metadata definition ID (without '!' prefix).
	ID string
	// Metadata definition node.
	Node MDNode // MDTuple or SpecializedMDNode

	// extra.

	// (optional) Distinct.
	Distinct bool
}

// String returns the string representation of the metadata definition.
func (md *MetadataDef) String() string {
	return enc.Metadata(md.ID)
}

// Def returns the LLVM syntax representation of the metadata definition.
func (md *MetadataDef) Def() string {
	// MetadataID "=" OptDistinct MDTuple
	// MetadataID "=" OptDistinct SpecializedMDNode
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", enc.Metadata(md.ID))
	if md.Distinct {
		buf.WriteString("distinct ")
	}
	buf.WriteString(md.Node.String())
	return buf.String()
}

// === [ Metadata Nodes and Metadata Strings ] =================================

// --- [ Metadata Tuple ] ------------------------------------------------------

// MDTuple is a metadata node tuple.
type MDTuple struct {
	// Metadata tuple fields.
	Fields []MDField
}

// String returns the string representation of the metadata node tuple.
func (md *MDTuple) String() string {
	// "!" MDFields
	buf := &strings.Builder{}
	buf.WriteString("!{")
	for i, field := range md.Fields {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(field.String())
	}
	buf.WriteString("}")
	return buf.String()
}

// --- [ Metadata Value ] ------------------------------------------------------

// A Value is a metadata value.
type Value struct {
	// Metadata value.
	Value Metadata
}

// String returns the LLVM syntax representation of the metadata value as a
// type-value pair.
func (md *Value) String() string {
	return fmt.Sprintf("%s %s", md.Type(), md.Ident())
}

// Type returns the type of the metadata value.
func (md *Value) Type() types.Type {
	return types.Metadata
}

// Ident returns the identifier associated with the metadata value.
func (md *Value) Ident() string {
	return md.Value.String()
}

// --- [ Metadata String ] -----------------------------------------------------

// MDString is a metadata string.
type MDString struct {
	// Metadata string value.
	Value string
}

// String returns the string representation of the metadata string.
func (md *MDString) String() string {
	// "!" StringLit
	return fmt.Sprintf("!%s", quote(md.Value))
}

// --- [ Metadata Attachment ] -------------------------------------------------

// MetadataAttachment is a metadata attachment.
type MetadataAttachment struct {
	// Metadata attachment name (without '!' prefix); e.g. !dbg.
	Name string
	// Metadata attachment node.
	Node MDNode
}

// String returns the string representation of the metadata attachment.
func (m *MetadataAttachment) String() string {
	// !dbg !42
	return fmt.Sprintf("%s %s", enc.Metadata(m.Name), m.Node)
}
