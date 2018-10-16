package ir

import (
	"fmt"

	"github.com/llir/l/internal/enc"
	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/types"
)

// === [ Global variables ] ====================================================

// Global is a global variable declaration or definition.
type Global struct {
	// Global variable name.
	GlobalName string
	// Immutability of global variable (constant or global).
	Immutable bool
	// Content type.
	ContentType types.Type
	// Initial value; or nil if declaration.
	Init Constant

	// extra.

	// Pointer type to global variable, including optional address space.
	Typ *types.PointerType
	// (optional) Linkage.
	Linkage ll.Linkage
	// (optional) Preemption; zero value if not present.
	Preemption ll.Preemption
	// (optional) Visibility; zero value if not present.
	Visibility ll.Visibility
	// (optional) DLL storage class; zero value if not present.
	DLLStorageClass ll.DLLStorageClass
	// (optional) Thread local storage model; zero value if not present.
	TLSModel ll.TLSModel
	// (optional) Unnamed address.
	UnnamedAddr ll.UnnamedAddr
	// (optional) Externally initialized.
	ExternallyInitialized bool
	// (optional) Section name; empty if not present.
	Section string
	// (optional) Comdat definition; nil if not present.
	// TODO: define ComdatDef.
	//Comdat *ComdatDef
	// (optional) Alignment; zero if not present.
	Alignment int
	// (optional) Function attributes.
	FuncAttrs []ll.FuncAttribute
	// (optional) Metadata attachments.
	// TODO: add support for metadata.
	//Metadata []*metadata.MetadataAttachment
}

// NewGlobalDecl returns a new global variable declaration based on the given
// global variable name and content type.
func NewGlobalDecl(name string, contentType types.Type) *Global {
	return &Global{GlobalName: name, ContentType: contentType}
}

// NewGlobalDef returns a new global variable definition based on the given
// global variable name and initial value.
func NewGlobalDef(name string, init Constant) *Global {
	return &Global{GlobalName: name, ContentType: init.Type(), Init: init}
}

// String returns the LLVM syntax representation of the global variable as a
// type-value pair.
func (g *Global) String() string {
	return fmt.Sprintf("%v %v", g.Type(), g.Ident())
}

// Type returns the type of the global variable.
func (g *Global) Type() types.Type {
	// Cache type if not present.
	if g.Typ == nil {
		g.Typ = types.NewPointer(g.ContentType)
	}
	return g.Typ
}

// Ident returns the identifier associated with the global variable.
func (g *Global) Ident() string {
	return enc.Global(g.GlobalName)
}

// Name returns the name of the global variable.
func (g *Global) Name() string {
	return g.GlobalName
}

// SetName sets the name of the global variable.
func (g *Global) SetName(name string) {
	g.GlobalName = name
}

// Def returns the LLVM syntax representation of the global variable definition
// or declaration.
func (g *Global) Def() string {
	panic("not yet implemented")
}
