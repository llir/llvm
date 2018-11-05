// Package ir declares the types used to represent LLVM IR modules.
package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
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
		ComdatDefs []*enum.ComdatDef
		// (optional) Indirect symbol definitions (aliases and IFuncs).
		// TODO: figure out how to represent aliases and IFuncs.
		//IndirectSymbols []*IndirectSymbol
		// (optional) Attribute group definitions.
		AttrGroupDefs []*enum.AttrGroupDef
		// (optional) Named metadata definitions.
		// TODO: figure out how to represent metadata.
		//NamedMetadataDefs []*metadata.NamedMetadataDef
		// (optional) Metadata definitions.
		// TODO: figure out how to represent metadata.
		//MetadataDefs []*metadata.MetadataDef
		// (optional) Use-list order directives.
		UseListOrders []*enum.UseListOrder
		// (optional) Basic block specific use-list order directives.
		UseListOrderBBs []*UseListOrderBB
	*/
}

// Def returns the LLVM syntax representation of the module.
func (m *Module) Def() string {
	buf := &strings.Builder{}
	// Type definitions.
	if len(m.TypeDefs) > 0 && buf.Len() > 0 {
		buf.WriteString("\n")
	}
	for _, t := range m.TypeDefs {
		// LocalIdent "=" "type" OpaqueType
		// LocalIdent "=" "type" Type
		fmt.Fprintf(buf, "%s = type %s\n", t, t.Def())
	}
	// TODO: implement Module.Def.
	// Global declarations and definitions.
	if len(m.Globals) > 0 && buf.Len() > 0 {
		buf.WriteString("\n")
	}
	for _, g := range m.Globals {
		fmt.Fprintln(buf, g.Def())
	}
	// Function declarations and definitions.
	if len(m.Funcs) > 0 && buf.Len() > 0 {
		buf.WriteString("\n")
	}
	for i, f := range m.Funcs {
		if i != 0 {
			buf.WriteString("\n")
		}
		fmt.Fprintln(buf, f.Def())
	}
	// TODO: implement Module.Def.
	return buf.String()
}

// ~~~ [ Comdat Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ComdatDef is a comdat definition top-level entity.
type ComdatDef struct {
	// Comdat name (without '$' prefix).
	Name string
	// Comdat kind.
	Kind enum.SelectionKind
}

// String returns the string representation of the Comdat definition.
func (c *ComdatDef) String() string {
	return c.Name
}

// Def returns the LLVM syntax representation of the Comdat definition.
func (c *ComdatDef) Def() string {
	// ComdatName "=" "comdat" SelectionKind
	return fmt.Sprintf("%s = comdat %s", enc.Comdat(c.Name), c.Kind)
}
