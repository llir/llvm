package irx

import (
	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/ir"
)

// === [ Modules ] =============================================================

// Translate translates the AST of the given module to an equivalent LLVM IR
// module.
func Translate(module *ast.Module) (*ir.Module, error) {
	m := NewModule()
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
func (m *Module) typeDef(def *ast.NamedType) {
	panic("not yet implemented")
}

// === [ Global variables ] ====================================================

// globalDecl translates the given global variable declaration to LLVM IR,
// emitting code to m.
func (m *Module) globalDecl(n *ast.Global) {
	panic("not yet implemented")
}

// === [ Functions ] ===========================================================

// funcDecl translates the given function declaration to LLVM IR, emitting code
// to m.
func (m *Module) funcDecl(n *ast.Function) {
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
