// Fix dummy values as follows.
//
// Per module.
//
//    1. Index type definitions.
//    2. Index global variables.
//    3. Index functions.
//    4. Fix type definitions.
//    5. Resolve named types.
//    6. Resolve global identifiers.
//
// Per function.
//
//    1. Assign unique local IDs to unnamed basic blocks and instructions.
//    2. Index basic blocks.
//    3. Index function parameters.
//    4. Index local variables produced by instructions.
//    5. Resolve local identifiers.

package astx

import (
	"fmt"

	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/asm/internal/ast/astutil"
	"github.com/llir/llvm/internal/enc"
)

// === [ Modules ] =============================================================

// fixModule replaces dummy values within the given module with their real
// values.
func fixModule(m *ast.Module) *ast.Module {
	fix := &fixer{
		globals:  make(map[string]ast.NamedValue),
		types:    make(map[string]*ast.NamedType),
		metadata: make(map[string]*ast.Metadata),
	}

	// Index type definitions.
	for _, typ := range m.Types {
		name := typ.Name
		if _, ok := fix.types[name]; ok {
			panic(fmt.Errorf("type name %q already present; old `%v`, new `%v`", name, fix.types[name], typ))
		}
		fix.types[name] = typ
	}

	// Index global variables.
	for _, global := range m.Globals {
		name := global.Name
		if _, ok := fix.globals[name]; ok {
			panic(fmt.Errorf("global identifier %q already present; old `%v`, new `%v`", name, fix.globals[name], global))
		}
		fix.globals[name] = global
	}

	// Index functions.
	for _, f := range m.Funcs {
		name := f.Name
		if _, ok := fix.globals[name]; ok {
			panic(fmt.Errorf("global identifier %q already present; old `%v`, new `%v`", name, fix.globals[name], f))
		}
		fix.globals[name] = f
	}

	// Index metadata.
	for _, md := range m.Metadata {
		id := md.ID
		if _, ok := fix.metadata[id]; ok {
			panic(fmt.Errorf("metadata ID %q already present; old `%v`, new `%v`", id, fix.metadata[id], md))
		}
		fix.metadata[id] = md
	}

	// Fix type definitions.
	for _, typ := range m.Types {
		typ.Def = fix.fixType(typ.Def)
	}

	// Resolve named types.
	resolveTypes := func(node interface{}) {
		p, ok := node.(*ast.Type)
		if !ok {
			return
		}
		old, ok := (*p).(*ast.NamedTypeDummy)
		if !ok {
			return
		}
		typ := fix.getType(old.Name)
		if typ.Def == nil {
			panic(fmt.Errorf("invalid type definition %q; expected underlying definition, got nil", typ.Name))
		}
		*p = typ
	}
	astutil.Walk(m, resolveTypes)

	// Resolve values of global identifiers.
	resolveGlobals := func(node interface{}) {
		p, ok := node.(*ast.Value)
		if !ok {
			return
		}
		old, ok := (*p).(*ast.GlobalDummy)
		if !ok {
			return
		}
		global := fix.getGlobal(old.Name)
		// TODO: Validate type of old and new global.
		*p = global
	}
	astutil.Walk(m, resolveGlobals)

	// Resolve named values of global identifiers.
	resolveNamedGlobals := func(node interface{}) {
		p, ok := node.(*ast.NamedValue)
		if !ok {
			return
		}
		old, ok := (*p).(*ast.GlobalDummy)
		if !ok {
			return
		}
		global := fix.getGlobal(old.Name)
		// TODO: Validate type of old and new global.
		*p = global
	}
	astutil.Walk(m, resolveNamedGlobals)

	// Resolve constants of global identifiers.
	resolveConstantGlobals := func(node interface{}) {
		p, ok := node.(*ast.Constant)
		if !ok {
			return
		}
		old, ok := (*p).(*ast.GlobalDummy)
		if !ok {
			return
		}
		global := fix.getGlobal(old.Name)
		g, ok := global.(ast.Constant)
		if !ok {
			panic(fmt.Errorf("invalid global type of %q; expected ast.Constant, got %T", global.GetName(), global))
		}
		// TODO: Validate type of old and new global.
		*p = g
	}
	astutil.Walk(m, resolveConstantGlobals)

	// Fix functions.
	for _, f := range m.Funcs {
		fix.fixFunc(f)
	}

	// Resolve metadata nodes.
	resolveMetadataNodes := func(node interface{}) {
		p, ok := node.(*ast.MetadataNode)
		if !ok {
			return
		}
		old, ok := (*p).(*ast.MetadataIDDummy)
		if !ok {
			return
		}
		metadata := fix.getMetadata(old.ID)
		*p = metadata
	}
	astutil.Walk(m, resolveMetadataNodes)

	return m
}

// === [ Type definitions ] ====================================================

// fixType replaces dummy types within the given type with their real types.
func (fix *fixer) fixType(old ast.Type) ast.Type {
	switch old := old.(type) {
	case *ast.VoidType:
		// nothing to do.
	case *ast.FuncType:
		old.Ret = fix.fixType(old.Ret)
		for _, param := range old.Params {
			param.Type = fix.fixType(param.Type)
		}
	case *ast.IntType:
		// nothing to do.
	case *ast.FloatType:
		// nothing to do.
	case *ast.PointerType:
		old.Elem = fix.fixType(old.Elem)
	case *ast.VectorType:
		old.Elem = fix.fixType(old.Elem)
	case *ast.LabelType:
		// nothing to do.
	case *ast.MetadataType:
		// nothing to do.
	case *ast.ArrayType:
		old.Elem = fix.fixType(old.Elem)
	case *ast.StructType:
		for i, field := range old.Fields {
			old.Fields[i] = fix.fixType(field)
		}
	case *ast.NamedType:
		if old.Def == nil {
			old.Def = fix.getType(old.Name)
		}
	case *ast.NamedTypeDummy:
		return fix.getType(old.Name)
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", old))
	}
	return old
}

// === [ Functions ] ===========================================================

// fixFunc replaces dummy values within the given function with their real
// values.
func (fix *fixer) fixFunc(f *ast.Function) {
	// Early exit if function declaration.
	if len(f.Blocks) < 1 {
		return
	}

	// Reset locals.
	fix.locals = make(map[string]ast.NamedValue)

	// Assign unique local IDs to unnamed basic blocks and instructions.
	f.AssignIDs()

	// Index basic blocks.
	for _, block := range f.Blocks {
		name := block.Name
		if _, ok := fix.locals[name]; ok {
			panic(fmt.Errorf("basic block label %q already present for function %s; old `%v`, new `%v`", name, enc.Global(f.Name), fix.locals[name], block))
		}
		fix.locals[name] = block
	}

	// Index function parameters.
	for _, param := range f.Sig.Params {
		name := param.Name
		if _, ok := fix.locals[name]; ok {
			panic(fmt.Errorf("function parameter name %q already present for function %s; old `%v`, new `%v`", name, enc.Global(f.Name), fix.locals[name], param))
		}
		fix.locals[name] = param
	}

	// Index local variables produced by instructions.
	for _, block := range f.Blocks {
		for _, inst := range block.Insts {
			if inst, ok := inst.(ast.NamedValue); ok {
				// Ignore local value if of type void.
				if inst, ok := inst.(*ast.InstCall); ok {
					if _, ok := inst.Type.(*ast.VoidType); ok {
						continue
					}
					if sig, ok := inst.Type.(*ast.FuncType); ok {
						if _, ok := sig.Ret.(*ast.VoidType); ok {
							continue
						}
					}
				}
				name := inst.GetName()
				if _, ok := fix.locals[name]; ok {
					panic(fmt.Errorf("instruction name %q already present for function %s; old `%v`, new `%v`", name, enc.Global(f.Name), fix.locals[name], inst))
				}
				fix.locals[name] = inst
			}
		}
	}

	// Resolve values of local identifiers.
	resolveLocals := func(node interface{}) {
		p, ok := node.(*ast.Value)
		if !ok {
			return
		}
		old, ok := (*p).(*ast.LocalDummy)
		if !ok {
			return
		}
		local := fix.getLocal(old.Name)
		// TODO: Validate type of old and new local.
		*p = local
	}
	astutil.WalkFunc(f, resolveLocals)

	// Resolve named values of local identifiers.
	resolveNamedLocals := func(node interface{}) {
		p, ok := node.(*ast.NamedValue)
		if !ok {
			return
		}
		old, ok := (*p).(*ast.LocalDummy)
		if !ok {
			return
		}
		local := fix.getLocal(old.Name)
		// TODO: Validate type of old and new local.
		*p = local
	}
	astutil.WalkFunc(f, resolveNamedLocals)
}

// ### [ Helper functions ] ####################################################

// A fixer keeps track of global and local identifiers to replace dummy values
// with their real values.
type fixer struct {
	// Per module.

	// types maps from type identifiers to their real types.
	types map[string]*ast.NamedType
	// globals maps global identifiers to their real values.
	globals map[string]ast.NamedValue
	// metadata maps metadata IDs to their real metadata.
	metadata map[string]*ast.Metadata

	// Per function.

	// locals maps local identifiers to their real values.
	locals map[string]ast.NamedValue
}

// getType returns the type of the given type name.
func (fix *fixer) getType(name string) *ast.NamedType {
	typ, ok := fix.types[name]
	if !ok {
		panic(fmt.Errorf("unable to locate type name %q", name))
	}
	return typ
}

// getGlobal returns the global value of the given global identifier.
func (fix *fixer) getGlobal(name string) ast.NamedValue {
	global, ok := fix.globals[name]
	if !ok {
		panic(fmt.Errorf("unable to locate global identifier %q", name))
	}
	return global
}

// getMetadata returns the metadata of the given metadata ID.
func (fix *fixer) getMetadata(id string) *ast.Metadata {
	metadata, ok := fix.metadata[id]
	if !ok {
		panic(fmt.Errorf("unable to locate metadata ID %q", id))
	}
	return metadata
}

// getLocal returns the local value of the given local identifier.
func (fix *fixer) getLocal(name string) ast.NamedValue {
	local, ok := fix.locals[name]
	if !ok {
		panic(fmt.Errorf("unable to locate local identifier %q", name))
	}
	return local
}
