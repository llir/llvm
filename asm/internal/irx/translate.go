// Translates AST values as follows.
//
// Per module.
//
//    1. Index type definitions.
//    2. Index global variables.
//    3. Index functions.
//    4. Fix type definitions.
//    5. Fix globals.
//    6. Fix functions.

package irx

import (
	"fmt"

	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
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
		typ := &types.NamedType{
			Name: name,
		}
		m.Types = append(m.Types, typ)
		m.types[name] = typ
	}

	// Index global variables.
	for _, old := range module.Globals {
		name := old.Name
		if _, ok := m.globals[name]; ok {
			panic(fmt.Errorf("global identifier %q already present; old `%v`, new `%v`", name, m.globals[name], old))
		}
		global := &ir.Global{
			Name: name,
		}
		m.Globals = append(m.Globals, global)
		m.globals[name] = global
	}

	// Index functions.
	for _, old := range module.Funcs {
		name := old.Name
		if _, ok := m.globals[name]; ok {
			panic(fmt.Errorf("global identifier %q already present; old `%v`, new `%v`", name, m.globals[name], old))
		}
		f := &ir.Function{
			Parent: m.Module,
			Name:   name,
		}
		m.Funcs = append(m.Funcs, f)
		m.globals[name] = f
	}

	// Fix type definitions.
	for _, typ := range module.Types {
		m.typeDef(typ)
	}

	// Fix globals.
	for _, global := range module.Globals {
		m.globalDecl(global)
	}

	// Fix functions.
	for _, f := range module.Funcs {
		m.funcDecl(f)
	}

	if len(m.errs) > 0 {
		// TODO: Return a list of all errors.
		return nil, m.errs[0]
	}
	return m.Module, nil
}

// === [ Type definitions ] ====================================================

// typeDef translates the given type definition to LLVM IR, emitting code to m.
func (m *Module) typeDef(old *ast.NamedType) {
	typ := m.getType(old.Name)
	def := m.irType(old.Def)
	typ.Def = def
}

// === [ Global variables ] ====================================================

// globalDecl translates the given global variable declaration to LLVM IR,
// emitting code to m.
func (m *Module) globalDecl(old *ast.Global) {
	v := m.getGlobal(old.Name)
	global, ok := v.(*ir.Global)
	if !ok {
		panic(fmt.Errorf("invalid global type; expected *ir.Global, got %T", v))
	}
	if old.Init != nil {
		init := m.irConstant(old.Init)
		global.Content = init.Type()
		global.Init = init
	} else {
		global.Content = m.irType(old.Content)
	}
	global.Typ = types.NewPointer(global.Content)
	global.IsConst = old.Immutable
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
