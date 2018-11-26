package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
)

// === [ Global variables ] ====================================================

// Global is a global variable declaration or definition.
type Global struct {
	// Global variable name (without '@' prefix).
	GlobalIdent
	// Immutability of global variable (constant or global).
	Immutable bool
	// Content type.
	ContentType types.Type
	// Initial value; or nil if declaration.
	Init constant.Constant

	// extra.

	// Pointer type to global variable, including an optional address space. If
	// Typ is nil, the first invocation of Type stores a pointer type with
	// ContentType as element.
	Typ *types.PointerType
	// (optional) Linkage; zero value if not present.
	Linkage enum.Linkage
	// (optional) Preemption; zero value if not present.
	Preemption enum.Preemption
	// (optional) Visibility; zero value if not present.
	Visibility enum.Visibility
	// (optional) DLL storage class; zero value if not present.
	DLLStorageClass enum.DLLStorageClass
	// (optional) Thread local storage model; zero value if not present.
	TLSModel enum.TLSModel
	// (optional) Unnamed address; zero value if not present.
	UnnamedAddr enum.UnnamedAddr
	// (optional) Externally initialized; false if not present.
	ExternallyInitialized bool
	// (optional) Section name; empty if not present.
	Section string
	// (optional) Comdat; nil if not present.
	Comdat *ComdatDef
	// (optional) Alignment; zero if not present.
	Align Align
	// (optional) Function attributes.
	FuncAttrs []FuncAttribute
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewGlobalDecl returns a new global variable declaration based on the given
// global variable name and content type.
func NewGlobalDecl(name string, contentType types.Type) *Global {
	global := &Global{ContentType: contentType}
	global.SetName(name)
	// Compute type.
	global.Type()
	return global
}

// NewGlobalDef returns a new global variable definition based on the given
// global variable name and initial value.
func NewGlobalDef(name string, init constant.Constant) *Global {
	global := &Global{ContentType: init.Type(), Init: init}
	global.SetName(name)
	// Compute type.
	global.Type()
	return global
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

// Def returns the LLVM syntax representation of the global variable definition
// or declaration.
func (g *Global) Def() string {
	// Global declaration.
	//
	//    Name=GlobalIdent '=' ExternLinkage Preemptionopt Visibilityopt
	//    DLLStorageClassopt ThreadLocalopt UnnamedAddropt AddrSpaceopt
	//    ExternallyInitializedopt Immutable ContentType=Type (',' Section)? (','
	//    Comdat)? (',' Alignment)? Metadata=(',' MetadataAttachment)+?
	//    FuncAttrs=(',' FuncAttribute)+?
	//
	// Global definition.
	//
	//    Name=GlobalIdent '=' Linkageopt Preemptionopt Visibilityopt
	//    DLLStorageClassopt ThreadLocalopt UnnamedAddropt AddrSpaceopt
	//    ExternallyInitializedopt Immutable ContentType=Type Init=Constant (','
	//    Section)? (',' Comdat)? (',' Alignment)? Metadata=(','
	//    MetadataAttachment)+? FuncAttrs=(',' FuncAttribute)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s =", g.Ident())
	if g.Linkage != enum.LinkageNone {
		fmt.Fprintf(buf, " %s", g.Linkage)
	}
	if g.Preemption != enum.PreemptionNone {
		fmt.Fprintf(buf, " %s", g.Preemption)
	}
	if g.Visibility != enum.VisibilityNone {
		fmt.Fprintf(buf, " %s", g.Visibility)
	}
	if g.DLLStorageClass != enum.DLLStorageClassNone {
		fmt.Fprintf(buf, " %s", g.DLLStorageClass)
	}
	if g.TLSModel != enum.TLSModelNone {
		fmt.Fprintf(buf, " %s", tlsModelString(g.TLSModel))
	}
	if g.UnnamedAddr != enum.UnnamedAddrNone {
		fmt.Fprintf(buf, " %s", g.UnnamedAddr)
	}
	if t, ok := g.Type().(*types.PointerType); ok {
		if t.AddrSpace != 0 {
			fmt.Fprintf(buf, " %s", t.AddrSpace)
		}
	}
	if g.ExternallyInitialized {
		buf.WriteString(" externally_initialized")
	}
	if g.Immutable {
		buf.WriteString(" constant")
	} else {
		buf.WriteString(" global")
	}
	fmt.Fprintf(buf, " %s", g.ContentType)
	if g.Init != nil {
		// Global definition.
		fmt.Fprintf(buf, " %s", g.Init.Ident())
	}
	if g.Section != "" {
		fmt.Fprintf(buf, ", section %s", quote(g.Section))
	}
	if g.Comdat != nil {
		if g.Comdat.Name == g.Name() {
			buf.WriteString(", comdat")
		} else {
			fmt.Fprintf(buf, ", %s", g.Comdat)
		}
	}
	if g.Align != 0 {
		fmt.Fprintf(buf, ", %s", g.Align)
	}
	for _, md := range g.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	for _, attr := range g.FuncAttrs {
		fmt.Fprintf(buf, " %s", attr)
	}
	return buf.String()
}
