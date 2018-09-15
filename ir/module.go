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
