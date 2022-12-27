package ir

import "github.com/llir/llvm/ir/types"

// --- [ Type definitions ] ----------------------------------------------------

// NewTypeDef appends a new type definition to the module based on the given
// type name and underlying type.
//
// Note, the name of the given type is set by invoking typ.SetName. As such,
// users are advised to refrain from creating type definitions from the
// convenience global variables of the types package (e.g. types.I64), since
// doing so would change the LLVM IR name of all uses of the types.I64 global
// variable. Instead, create a new type (e.g. types.NewInt(64)) for this
// purpose.
func (m *Module) NewTypeDef(name string, typ types.Type) types.Type {
	typ.SetName(name)
	m.TypeDefs = append(m.TypeDefs, typ)
	return typ
}
