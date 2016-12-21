// Translates AST values as follows.
//
// Per module.
//
//    1. Index type definitions.
//    2. Fix type definitions.
//    3. Index global variables.
//    4. Index functions.

package irx

import (
	"fmt"

	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/ir"
)

// === [ Modules ] =============================================================

// Translate translates the AST of the given module to an equivalent LLVM IR
// module.
func Translate(module *ast.Module) (*ir.Module, error) {
	m := NewModule()

	// Index type definitions.
	for _, old := range module.Types {
		name := old.Name
		if _, ok := m.types[name]; ok {
			panic(fmt.Errorf("type name %q already present; old `%v`, new `%v`", name, m.types[name], old))
		}
		typ := m.NewType(name, nil)
		m.types[name] = typ
	}

	// Fix type definitions.
	for _, typ := range module.Types {
		m.typeDef(typ)
	}

	for _, global := range module.Globals {
		m.globalDecl(global)
	}
	for _, f := range module.Funcs {
		m.funcDecl(f)
	}
	return m.Module, nil
}

// === [ Type definitions ] ====================================================

// typeDef translates the given type definition to LLVM IR, emitting code to m.
func (m *Module) typeDef(old *ast.NamedType) {
	typ := m.getType(old.Name)
	def := m.irtype(old.Def)
	typ.Def = def
}

// === [ Global variables ] ====================================================

// globalDecl translates the given global variable declaration to LLVM IR,
// emitting code to m.
func (m *Module) globalDecl(old *ast.Global) {
	panic("not yet implemented")
}

// === [ Functions ] ===========================================================

// funcDecl translates the given function declaration to LLVM IR, emitting code
// to m.
func (m *Module) funcDecl(old *ast.Function) {
	panic("not yet implemented")
}

// === [ Identifiers ] =========================================================

// === [ Types ] ===============================================================

// === [ Values ] ==============================================================

// === [ Constants ] ===========================================================

// --- [ Binary expressions ] --------------------------------------------------

// --- [ Bitwise expressions ] -------------------------------------------------

// --- [ Memory expressions ] --------------------------------------------------

// --- [ Conversion expressions ] ----------------------------------------------

// --- [ Other expressions ] ---------------------------------------------------

// === [ Basic blocks ] ========================================================

// === [ Instructions ] ========================================================

// --- [ Binary instructions ] -------------------------------------------------

// --- [ Bitwise instructions ] ------------------------------------------------

// --- [ Memory instructions ] -------------------------------------------------

// --- [ Conversion instructions ] ---------------------------------------------

// --- [ Other instructions ] --------------------------------------------------

// === [ Terminators ] =========================================================
