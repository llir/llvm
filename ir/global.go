package ir

import (
	"fmt"
	"strings"

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

	// Pointer type to global variable, including an optional address space. If
	// Typ is nil, the first invocation of Type stores a pointer type with
	// ContentType as element.
	Typ *types.PointerType
	// (optional) Linkage; zero value if not present.
	Linkage ll.Linkage
	// (optional) Preemption; zero value if not present.
	Preemption ll.Preemption
	// (optional) Visibility; zero value if not present.
	Visibility ll.Visibility
	// (optional) DLL storage class; zero value if not present.
	DLLStorageClass ll.DLLStorageClass
	// (optional) Thread local storage model; zero value if not present.
	TLSModel ll.TLSModel
	// (optional) Unnamed address; zero value if not present.
	UnnamedAddr ll.UnnamedAddr
	// (optional) Externally initialized; false if not present.
	ExternallyInitialized bool
	// (optional) Section name; empty if not present.
	Section string
	// (optional) Comdat definition; nil if not present.
	Comdat *ComdatDef
	// (optional) Alignment; zero if not present.
	Align int64
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
	return fmt.Sprintf("%s %s", g.Type(), g.Ident())
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
	// GlobalIdent "=" OptLinkage OptPreemptionSpecifier OptVisibility
	// OptDLLStorageClass OptThreadLocal OptUnnamedAddr OptAddrSpace
	// OptExternallyInitialized Immutable Type Constant GlobalAttrs FuncAttrs
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s =", g.Ident())
	if g.Linkage != ll.LinkageNone {
		fmt.Fprintf(buf, " %s", g.Linkage)
	}
	if g.Preemption != ll.PreemptionNone {
		fmt.Fprintf(buf, " %s", g.Preemption)
	}
	if g.Visibility != ll.VisibilityNone {
		fmt.Fprintf(buf, " %s", g.Visibility)
	}
	if g.DLLStorageClass != ll.DLLStorageClassNone {
		fmt.Fprintf(buf, " %s", g.DLLStorageClass)
	}
	if g.TLSModel != ll.TLSModelNone {
		fmt.Fprintf(buf, " %s", g.TLSModel)
	}
	if g.UnnamedAddr != ll.UnnamedAddrNone {
		fmt.Fprintf(buf, " %s", g.UnnamedAddr)
	}
	if g.Typ.AddrSpace != 0 {
		fmt.Fprintf(buf, " %s", g.Typ.AddrSpace)
	}
	if g.ExternallyInitialized {
		buf.WriteString(" externallyinitialized")
	}
	if g.Immutable {
		buf.WriteString(" constant")
	} else {
		buf.WriteString(" global")
	}
	fmt.Fprintf(buf, " %s", g.ContentType)
	if g.Init != nil {
		fmt.Fprintf(buf, " %s", g.Init.Ident())
	}
	if g.Section != "" {
		fmt.Fprintf(buf, ", section %s", quote(g.Section))
	}
	if g.Comdat != nil {
		fmt.Fprintf(buf, ", %s", g.Comdat)
	}
	if g.Align != 0 {
		fmt.Fprintf(buf, ", align %d", g.Align)
	}
	// TODO: add metadata.
	//for _, md := range g.Metadata {
	//	fmt.Fprintf(buf, ", %s", md)
	//}
	// TODO: add function attributes.
	//for _, attr := range g.FuncAttrs {
	//	fmt.Fprintf(buf, " %s", attr)
	//}
	return buf.String()
}
