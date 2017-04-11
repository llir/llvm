// === [ Metadata ] ============================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#metadata

package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// A MetadataNode represents an LLVM IR metadata node.
//
// MetadataNode may have one of the following underlying types.
//
//    *ir.Metadata         (https://godoc.org/github.com/llir/llvm/ir#Metadata)
//    *ir.MetadataString   (https://godoc.org/github.com/llir/llvm/ir#MetadataString)
//    constant.Constant    (https://godoc.org/github.com/llir/llvm/ir/constant#Constant)
type MetadataNode interface {
	value.Value
	// MetadataNode ensures that only metadata nodes can be assigned to the
	// ir.MetadataNode interface.
	MetadataNode()
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

// Type returns the type of the metadata.
func (md *Metadata) Type() types.Type {
	return types.Metadata
}

// Ident returns the identifier associated with the metadata.
func (md *Metadata) Ident() string {
	if len(md.ID) > 0 {
		return enc.Metadata(md.ID)
	}
	return md.Def()
}

// Def returns the LLVM syntax representation of the definition of the metadata.
func (md *Metadata) Def() string {
	buf := &bytes.Buffer{}
	buf.WriteString("!{")
	for i, node := range md.Nodes {
		if i != 0 {
			buf.WriteString(", ")
		}
		if _, ok := node.(constant.Constant); ok {
			fmt.Fprintf(buf, "%s ", node.Type())
		}
		buf.WriteString(node.Ident())
	}
	buf.WriteString("}")
	return buf.String()
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*Metadata) MetadataNode() {}

// --- [ metadata string ] -----------------------------------------------------

// A MetadataString represents an LLVM IR metadata string.
type MetadataString struct {
	// String value.
	Val string
}

// Ident returns the identifier associated with the metadata.
func (md *MetadataString) Ident() string {
	return fmt.Sprintf(`!"%s"`, enc.EscapeString(md.Val))
}

// Type returns the type of the metadata.
func (md *MetadataString) Type() types.Type {
	return types.Metadata
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*MetadataString) MetadataNode() {}

// --- [ named metadata ] ------------------------------------------------------

// NamedMetadata represents a named collection of metadata, which belongs to a
// module.
type NamedMetadata struct {
	// Metadata name.
	Name string
	// Associated metadata.
	Metadata []*Metadata
}

// Def returns the LLVM syntax representation of the definition of the named
// metadata.
func (md *NamedMetadata) Def() string {
	buf := &bytes.Buffer{}
	buf.WriteString("!{")
	for i, metadata := range md.Metadata {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(metadata.Ident())
	}
	buf.WriteString("}")
	return buf.String()
}
