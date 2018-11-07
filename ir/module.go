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
	// (optional) Data layout; or empty if not present.
	DataLayout string
	// (optional) Target triple; or empty if not present.
	TargetTriple string
	// (optional) Module-level inline assembly.
	ModuleAsms []string
	// (optional) Comdat definitions.
	ComdatDefs []*ComdatDef
	// (optional) Aliases.
	Aliases []*Alias
	// (optional) IFuncs.
	IFuncs []*IFunc
	/*
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
	// Source filename.
	if len(m.SourceFilename) > 0 {
		// "source_filename" "=" StringLit
		fmt.Fprintf(buf, "source_filename = %s\n", quote(m.SourceFilename))
	}
	// Data layout.
	if len(m.DataLayout) > 0 {
		// "target" "datalayout" "=" StringLit
		fmt.Fprintf(buf, "target datalayout = %s\n", quote(m.DataLayout))
	}
	// Target triple.
	if len(m.TargetTriple) > 0 {
		// "target" "triple" "=" StringLit
		fmt.Fprintf(buf, "target triple = %s\n", quote(m.TargetTriple))
	}
	// Module-level inline assembly.
	for _, asm := range m.ModuleAsms {
		// "module" "asm" StringLit
		fmt.Fprintf(buf, "module asm %s\n", quote(asm))
	}
	// Type definitions.
	if len(m.TypeDefs) > 0 && buf.Len() > 0 {
		buf.WriteString("\n")
	}
	for _, t := range m.TypeDefs {
		// LocalIdent "=" "type" OpaqueType
		// LocalIdent "=" "type" Type
		fmt.Fprintf(buf, "%s = type %s\n", t, t.Def())
	}
	// Global declarations and definitions.
	if len(m.ComdatDefs) > 0 && buf.Len() > 0 {
		buf.WriteString("\n")
	}
	for i, c := range m.ComdatDefs {
		if i != 0 {
			buf.WriteString("\n")
		}
		fmt.Fprintln(buf, c.Def())
	}
	// Global declarations and definitions.
	if len(m.Globals) > 0 && buf.Len() > 0 {
		buf.WriteString("\n")
	}
	for _, g := range m.Globals {
		fmt.Fprintln(buf, g.Def())
	}
	// Aliases.
	if len(m.Aliases) > 0 && buf.Len() > 0 {
		buf.WriteString("\n")
	}
	for _, alias := range m.Aliases {
		fmt.Fprintln(buf, alias.Def())
	}
	// IFuncs.
	if len(m.IFuncs) > 0 && buf.Len() > 0 {
		buf.WriteString("\n")
	}
	for _, ifunc := range m.IFuncs {
		fmt.Fprintln(buf, ifunc.Def())
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
	// Attribute group definitions.
	// TODO: add support for AttrGoupDefs.
	//for _, a := range m.AttrGroupDefs {
	//	fmt.Fprintln(buf, a.Def())
	//}
	// Named metadata definitions.
	// TODO: add support for named metadata definitions.
	//for _, md := range m.NamedMetadataDefs {
	//	fmt.Fprintln(buf, md.Def())
	//}
	// Metadata definitions.
	// TODO: add support for metadata definitions.
	//for _, md := range m.MetadataDefs {
	//	fmt.Fprintln(buf, md.Def())
	//}
	// Use-list orders.
	// TODO: add support for use-list orders.
	//for _, u := range m.UseListOrders {
	//	fmt.Fprintln(buf, u.Def())
	//}
	// Basic block specific use-list orders.
	// TODO: add support for basic block specific use-list orders.
	//for _, u := range m.UseListOrderBBs {
	//	fmt.Fprintln(buf, u.Def())
	//}
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
	return fmt.Sprintf("comdat(%s)", enc.Comdat(c.Name))
}

// Def returns the LLVM syntax representation of the Comdat definition.
func (c *ComdatDef) Def() string {
	// ComdatName "=" "comdat" SelectionKind
	return fmt.Sprintf("%s = comdat %s", enc.Comdat(c.Name), c.Kind)
}
