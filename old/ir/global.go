// === [ Global variables ] ====================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#global-variables

package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
)

// A Global represents an LLVM IR global variable definition or external global
// variable declaration.
//
// Global variables always define a pointer to their "content" type because they
// describe a region of memory, and all memory objects in LLVM are accessed
// through pointers.
//
// Global variables may be referenced from instructions (e.g. load), and are
// thus considered LLVM IR values of pointer type.
type Global struct {
	// Global variable name.
	Name string
	// Global variable type.
	Typ *types.PointerType
	// Content type.
	Content types.Type
	// Initial value; or nil if defined externally.
	Init constant.Constant
	// Immutability of the global variable.
	IsConst bool
	// Map from metadata identifier (e.g. !dbg) to metadata associated with the
	// global.
	Metadata map[string]*metadata.Metadata
}

// NewGlobalDecl returns a new external global variable declaration based on the
// given global variable name and content type.
func NewGlobalDecl(name string, content types.Type) *Global {
	typ := types.NewPointer(content)
	return &Global{
		Name:     name,
		Typ:      typ,
		Content:  content,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// NewGlobalDef returns a new global variable definition based on the given
// global variable name and initial value.
func NewGlobalDef(name string, init constant.Constant) *Global {
	content := init.Type()
	typ := types.NewPointer(content)
	return &Global{
		Name:     name,
		Typ:      typ,
		Content:  content,
		Init:     init,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// Type returns the type of the global variable.
func (global *Global) Type() types.Type {
	return global.Typ
}

// Ident returns the identifier associated with the global variable.
func (global *Global) Ident() string {
	return enc.Global(global.Name)
}

// GetName returns the name of the global variable.
func (global *Global) GetName() string {
	return global.Name
}

// SetName sets the name of the global variable.
func (global *Global) SetName(name string) {
	global.Name = name
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Global) Immutable() {}

// MetadataNode ensures that only metadata nodes can be assigned to the
// metadata.Node interface.
func (*Global) MetadataNode() {}

// String returns the LLVM syntax representation of the global variable.
func (global *Global) String() string {
	imm := "global"
	if global.IsConst {
		imm = "constant"
	}
	md := metadataString(global.Metadata, ",")
	addrspace := &bytes.Buffer{}
	if global.Typ.AddrSpace != 0 {
		fmt.Fprintf(addrspace, " addrspace(%d)", global.Typ.AddrSpace)
	}
	if global.Init != nil {
		// Global variable definition.
		return fmt.Sprintf("%s =%s %s %s %s%s",
			global.Ident(),
			addrspace,
			imm,
			global.Init.Type(),
			global.Init.Ident(),
			md)
	}
	// External global variable declaration.
	return fmt.Sprintf("%s = external%s %s %s%s",
		global.Ident(),
		addrspace,
		imm,
		global.Content,
		md)
}
