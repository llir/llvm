// Package ir declares the types used to represent LLVM IR modules.
package ir

import (
	"fmt"
	"strings"

	"github.com/llir/l/ir/types"
)

// === [ Modules ] =============================================================

// Module is an LLVM IR module.
type Module struct {
	// Type definitions.
	TypeDefs []types.Type
	// Global variable declarations and definitions.
	Globals []*Global
	// Function declarations and definitions.
	Funcs []*Function

	// extra.

	// (optional) Source filename; or empty if not present.
	SourceFilename string
	/*
		// (optional) Data layout; or empty if not present.
		DataLayout string
		// (optional) Target triple; or empty if not present.
		TargetTriple string
		// (optional) Module-level inline assembly.
		ModuleAsms []string
		// (optional) Comdat definitions.
		ComdatDefs []*ll.ComdatDef
		// (optional) Indirect symbol definitions (aliases and IFuncs).
		// TODO: figure out how to represent aliases and IFuncs.
		//IndirectSymbols []*IndirectSymbol
		// (optional) Attribute group definitions.
		AttrGroupDefs []*ll.AttrGroupDef
		// (optional) Named metadata definitions.
		// TODO: figure out how to represent metadata.
		//NamedMetadataDefs []*metadata.NamedMetadataDef
		// (optional) Metadata definitions.
		// TODO: figure out how to represent metadata.
		//MetadataDefs []*metadata.MetadataDef
		// (optional) Use-list order directives.
		UseListOrders []*ll.UseListOrder
		// (optional) Basic block specific use-list order directives.
		UseListOrderBBs []*UseListOrderBB
	*/
}

// Def returns the LLVM syntax representation of the module.
func (m *Module) Def() string {
	buf := &strings.Builder{}
	// Type definitions.
	for _, t := range m.TypeDefs {
		// LocalIdent "=" "type" OpaqueType
		// LocalIdent "=" "type" Type
		fmt.Fprintf(buf, "%v = type %v\n", t, t.Def())
	}
	// TODO: implement Module.Def.
	return buf.String()
}
