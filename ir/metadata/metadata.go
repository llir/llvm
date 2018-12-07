// Package metadata provides access to LLVM IR metadata.
package metadata

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/types"
)

// Convenience literals.
var (
	// Null metadata literal.
	Null = &NullLit{}
)

// ~~~ [ Named metadata definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NamedDef is a named metadata definition.
type NamedDef struct {
	// Metadata definition name (without '!' prefix).
	Name string
	// Metadata definition nodes.
	Nodes []Node
}

// String returns the string representation of the named metadata definition.
func (md *NamedDef) String() string {
	return enc.MetadataName(md.Name)
}

// Def returns the LLVM syntax representation of the named metadata definition.
func (md *NamedDef) Def() string {
	// Name=MetadataName '=' '!' '{' MDNodes=(MetadataNode separator ',')* '}'
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = !{", enc.MetadataName(md.Name))
	for i, node := range md.Nodes {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(node.String())
	}
	buf.WriteString("}")
	return buf.String()
}

// ~~~ [ Metadata definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// TODO: check if metadata.Def should implement value.Value.

// Def is a metadata definition.
type Def struct {
	// Metadata definition ID (without '!' prefix).
	ID int64
	// Metadata definition node.
	Node MDNode // Tuple or SpecializedNode

	// extra.

	// (optional) Distinct.
	Distinct bool
}

// String returns the string representation of the metadata definition.
func (md *Def) String() string {
	return enc.MetadataID(md.ID)
}

// Def returns the LLVM syntax representation of the metadata definition.
func (md *Def) Def() string {
	// ID=MetadataID '=' Distinctopt MDNode=MDTuple
	//
	// ID=MetadataID '=' Distinctopt MDNode=SpecializedMDNode
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", enc.MetadataID(md.ID))
	if md.Distinct {
		buf.WriteString("distinct ")
	}
	buf.WriteString(md.Node.String())
	return buf.String()
}

// === [ Metadata nodes and metadata strings ] =================================

// --- [ Metadata tuple ] ------------------------------------------------------

// Tuple is a metadata node tuple.
type Tuple struct {
	// Metadata tuple fields.
	Fields []Field
}

// String returns the string representation of the metadata node tuple.
func (md *Tuple) String() string {
	// '!' MDFields
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

// --- [ Metadata value ] ------------------------------------------------------

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

// --- [ Metadata string ] -----------------------------------------------------

// String is a metadata string.
type String struct {
	// Metadata string value.
	Value string
}

// String returns the string representation of the metadata string.
func (md *String) String() string {
	// '!' Val=StringLit
	return fmt.Sprintf("!%s", quote(md.Value))
}

// --- [ Metadata attachment ] -------------------------------------------------

// Attachment is a metadata attachment.
type Attachment struct {
	// Metadata attachment name (without '!' prefix); e.g. !dbg.
	Name string
	// Metadata attachment node.
	Node MDNode
}

// String returns the string representation of the metadata attachment.
func (m *Attachment) String() string {
	// Name=MetadataName MDNode
	return fmt.Sprintf("%s %s", enc.MetadataName(m.Name), m.Node)
}

// --- [ Integer literals ] -----------------------------------------------------

// IntLit is an integer literal.
type IntLit int64

// String returns the string representation of the integer literal.
func (i IntLit) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// UintLit is an unsigned integer literal.
type UintLit uint64

// String returns the string representation of the unsigned integer literal.
func (i UintLit) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

// --- [ Null literal ] --------------------------------------------------------

// NullLit is a metadata null literal.
type NullLit struct{}

// String returns the string representation of the metadata null literal.
func (i *NullLit) String() string {
	return "null"
}
