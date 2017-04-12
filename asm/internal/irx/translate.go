// Translates AST values as follows.
//
// Per module.
//
//    1. Index type definitions.
//    2. Index global variables.
//       - Store preliminary content type.
//    3. Index function.
//       - Store type.
//    4. Fix type definitions.
//    5. Fix globals.
//    6. Fix functions.
//
// Per function.
//
//    1. Index function parameters.
//    2. Index basic blocks.
//    3. Index local variables produced by instructions.
//       - Store preliminary type.
//    4. Fix basic blocks.

package irx

import (
	"fmt"

	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
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
		typ := newEmptyNamedType(old.Def)
		typ.SetName(name)
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
		// Store preliminary content type.
		content := m.irType(old.Content)
		global.Typ = types.NewPointer(content)
		global.Content = content
		m.Globals = append(m.Globals, global)
		m.globals[name] = global
	}

	// Index functions.
	for _, old := range module.Funcs {
		name := old.Name
		if _, ok := m.globals[name]; ok {
			panic(fmt.Errorf("global identifier %q already present; old `%v`, new `%v`", name, m.globals[name], old))
		}
		// Store type.
		oldSig := m.irType(old.Sig)
		sig, ok := oldSig.(*types.FuncType)
		if !ok {
			panic(fmt.Errorf("invalid function signature type, expected *types.FuncType, got %T", oldSig))
		}
		typ := types.NewPointer(sig)
		f := &ir.Function{
			Parent:   m.Module,
			Name:     name,
			Typ:      typ,
			Sig:      sig,
			Metadata: make(map[string]*ir.Metadata),
		}
		m.Funcs = append(m.Funcs, f)
		m.globals[name] = f
	}

	// Index metadata.
	for _, old := range module.Metadata {
		id := old.ID
		if _, ok := m.metadata[id]; ok {
			panic(fmt.Errorf("metadata ID %q already present; old `%v`, new `%v`", id, m.metadata[id], old))
		}
		md := &ir.Metadata{
			ID: id,
		}
		m.Metadata = append(m.Metadata, md)
		m.metadata[id] = md
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

	// Fix named metadata definitions.
	for _, old := range module.NamedMetadata {
		md := &ir.NamedMetadata{
			Name: old.Name,
		}
		for _, oldMetadata := range old.Metadata {
			old, ok := oldMetadata.(*ast.Metadata)
			if !ok {
				panic(fmt.Errorf("invalid metadata type; expected *ast.Metadata, got %T", oldMetadata))
			}
			metadata := m.getMetadata(old.ID)
			md.Metadata = append(md.Metadata, metadata)
		}
		m.NamedMetadata = append(m.NamedMetadata, md)
	}

	// Fix metadata definition.
	for _, md := range module.Metadata {
		m.metadataDef(md)
	}

	if len(m.errs) > 0 {
		// TODO: Return a list of all errors.
		return nil, m.errs[0]
	}
	return m.Module, nil
}

// newEmptyNamedType returns an empty type definition for the given named type.
func newEmptyNamedType(old ast.Type) types.Type {
	switch old := old.(type) {
	case *ast.VoidType:
		return &types.VoidType{}
	case *ast.FuncType:
		return &types.FuncType{}
	case *ast.IntType:
		return &types.IntType{}
	case *ast.FloatType:
		return &types.FloatType{}
	case *ast.PointerType:
		return &types.PointerType{}
	case *ast.VectorType:
		return &types.VectorType{}
	case *ast.LabelType:
		return &types.LabelType{}
	case *ast.MetadataType:
		return &types.MetadataType{}
	case *ast.ArrayType:
		return &types.ArrayType{}
	case *ast.StructType:
		return &types.StructType{}
	case *ast.NamedType:
		return newEmptyNamedType(old.Def)
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", old))
	}
}

// === [ Type definitions ] ====================================================

// typeDef translates the given type definition to LLVM IR, emitting code to m.
func (m *Module) typeDef(old *ast.NamedType) {
	typ := m.getType(old.Name)
	def := m.irType(old.Def)
	switch typ := typ.(type) {
	case *types.VoidType:
		_, ok := def.(*types.VoidType)
		if !ok {
			panic(fmt.Errorf("invalid type; expected *types.VoidType, got %T", def))
		}
		// nothing to do.
	case *types.FuncType:
		d, ok := def.(*types.FuncType)
		if !ok {
			panic(fmt.Errorf("invalid type; expected *types.FuncType, got %T", def))
		}
		typ.Ret = d.Ret
		typ.Params = d.Params
		typ.Variadic = d.Variadic
	case *types.IntType:
		d, ok := def.(*types.IntType)
		if !ok {
			panic(fmt.Errorf("invalid type; expected *types.IntType, got %T", def))
		}
		typ.Size = d.Size
	case *types.FloatType:
		d, ok := def.(*types.FloatType)
		if !ok {
			panic(fmt.Errorf("invalid type; expected *types.FloatType, got %T", def))
		}
		typ.Kind = d.Kind
	case *types.PointerType:
		d, ok := def.(*types.PointerType)
		if !ok {
			panic(fmt.Errorf("invalid type; expected *types.PointerType, got %T", def))
		}
		typ.Elem = d.Elem
	case *types.VectorType:
		d, ok := def.(*types.VectorType)
		if !ok {
			panic(fmt.Errorf("invalid type; expected *types.VectorType, got %T", def))
		}
		typ.Elem = d.Elem
		typ.Len = d.Len
	case *types.LabelType:
		_, ok := def.(*types.LabelType)
		if !ok {
			panic(fmt.Errorf("invalid type; expected *types.LabelType, got %T", def))
		}
		// nothing to do.
	case *types.MetadataType:
		_, ok := def.(*types.MetadataType)
		if !ok {
			panic(fmt.Errorf("invalid type; expected *types.MetadataType, got %T", def))
		}
		// nothing to do.
	case *types.ArrayType:
		d, ok := def.(*types.ArrayType)
		if !ok {
			panic(fmt.Errorf("invalid type; expected *types.ArrayType, got %T", def))
		}
		typ.Elem = d.Elem
		typ.Len = d.Len
	case *types.StructType:
		d, ok := def.(*types.StructType)
		if !ok {
			panic(fmt.Errorf("invalid type; expected *types.StructType, got %T", def))
		}
		typ.Fields = d.Fields
		typ.Opaque = d.Opaque
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", typ))
	}
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
		// TODO: Verify that two circularly referential globals both get the
		// correct type; more specifically that neither get global.Content == nil
		// after resolution.
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
func (m *Module) funcDecl(oldFunc *ast.Function) {
	v := m.getGlobal(oldFunc.Name)
	f, ok := v.(*ir.Function)
	if !ok {
		panic(fmt.Errorf("invalid function type for function %s; expected *ir.Function, got %T", enc.Global(oldFunc.Name), v))
	}

	// Fix attached metadata.
	for _, oldMetadata := range oldFunc.Metadata {
		key := oldMetadata.Name
		metadata := m.metadataNode(oldMetadata.Metadata)
		if prev, ok := f.Metadata[key]; ok {
			panic(fmt.Errorf("metadata for metadata name %q already present; previous `%v`, new `%v`", key, prev, m.Metadata))
		}
		md, ok := metadata.(*ir.Metadata)
		if !ok {
			panic(fmt.Errorf("invalid metadata type; expected *ir.Metadata, got %T", metadata))
		}
		f.Metadata[key] = md
	}

	// Early exit if function declaration.
	if len(oldFunc.Blocks) < 1 {
		return
	}

	// Reset locals.
	m.locals = make(map[string]value.Named)

	// Index function parameters.
	for _, param := range f.Params() {
		name := param.Name
		if _, ok := m.locals[name]; ok {
			panic(fmt.Errorf("local identifier %q already present for function %s; old `%v`, new `%v`", name, f.Ident(), m.locals[name], param))
		}
		m.locals[name] = param
	}

	// Index basic blocks.
	for _, old := range oldFunc.Blocks {
		name := old.Name
		if _, ok := m.locals[name]; ok {
			panic(fmt.Errorf("local identifier %q already present for function %s; old `%v`, new `%v`", name, f.Ident(), m.locals[name], old))
		}
		block := &ir.BasicBlock{
			Name:   name,
			Parent: f,
		}
		f.Blocks = append(f.Blocks, block)
		m.locals[name] = block
	}

	// Index local variables produced by instructions.
	for i := 0; i < len(oldFunc.Blocks); i++ {
		oldBlock := oldFunc.Blocks[i]
		block := f.Blocks[i]
		for _, oldInst := range oldBlock.Insts {
			var inst ir.Instruction
			switch oldInst := oldInst.(type) {
			// Binary instructions
			case *ast.InstAdd:
				inst = &ir.InstAdd{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstFAdd:
				inst = &ir.InstFAdd{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstSub:
				inst = &ir.InstSub{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstFSub:
				inst = &ir.InstFSub{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstMul:
				inst = &ir.InstMul{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstFMul:
				inst = &ir.InstFMul{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstUDiv:
				inst = &ir.InstUDiv{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstSDiv:
				inst = &ir.InstSDiv{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstFDiv:
				inst = &ir.InstFDiv{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstURem:
				inst = &ir.InstURem{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstSRem:
				inst = &ir.InstSRem{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstFRem:
				inst = &ir.InstFRem{
					Parent: block,
					Name:   oldInst.Name,
				}

			// Bitwise instructions
			case *ast.InstShl:
				inst = &ir.InstShl{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstLShr:
				inst = &ir.InstLShr{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstAShr:
				inst = &ir.InstAShr{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstAnd:
				inst = &ir.InstAnd{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstOr:
				inst = &ir.InstOr{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstXor:
				inst = &ir.InstXor{
					Parent: block,
					Name:   oldInst.Name,
				}

			// Vector instructions
			case *ast.InstExtractElement:
				inst = &ir.InstExtractElement{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstInsertElement:
				inst = &ir.InstInsertElement{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstShuffleVector:
				inst = &ir.InstShuffleVector{
					Parent: block,
					Name:   oldInst.Name,
				}

			// Aggregate instructions
			case *ast.InstExtractValue:
				inst = &ir.InstExtractValue{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstInsertValue:
				inst = &ir.InstInsertValue{
					Parent: block,
					Name:   oldInst.Name,
				}

			// Memory instructions
			case *ast.InstAlloca:
				inst = &ir.InstAlloca{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstLoad:
				inst = &ir.InstLoad{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstStore:
				// Store instructions produce no value, and are thus not assigned
				// names.
				inst = &ir.InstStore{
					Parent: block,
				}
			case *ast.InstGetElementPtr:
				inst = &ir.InstGetElementPtr{
					Parent: block,
					Name:   oldInst.Name,
				}

			// Conversion instructions
			case *ast.InstTrunc:
				inst = &ir.InstTrunc{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstZExt:
				inst = &ir.InstZExt{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstSExt:
				inst = &ir.InstSExt{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstFPTrunc:
				inst = &ir.InstFPTrunc{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstFPExt:
				inst = &ir.InstFPExt{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstFPToUI:
				inst = &ir.InstFPToUI{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstFPToSI:
				inst = &ir.InstFPToSI{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstUIToFP:
				inst = &ir.InstUIToFP{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstSIToFP:
				inst = &ir.InstSIToFP{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstPtrToInt:
				inst = &ir.InstPtrToInt{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstIntToPtr:
				inst = &ir.InstIntToPtr{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstBitCast:
				inst = &ir.InstBitCast{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstAddrSpaceCast:
				inst = &ir.InstAddrSpaceCast{
					Parent: block,
					Name:   oldInst.Name,
				}

			// Other instructions
			case *ast.InstICmp:
				inst = &ir.InstICmp{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstFCmp:
				inst = &ir.InstFCmp{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstPhi:
				inst = &ir.InstPhi{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstSelect:
				inst = &ir.InstSelect{
					Parent: block,
					Name:   oldInst.Name,
				}
			case *ast.InstCall:
				inst = &ir.InstCall{
					Parent: block,
					Name:   oldInst.Name,
				}

			default:
				panic(fmt.Errorf("support for instruction %T not yet implemented", oldInst))
			}
			block.Insts = append(block.Insts, inst)

			// TODO: Validate if it is required to store a preliminary type of
			// instructions prior to local variable resolution.
			//
			// What happens if a getelementptr instruction refers to the value
			// produced by an instruction which cannot be calculated prior to its
			// operands being resolved?
			//
			//// Store preliminary type.

			// Index local variable.
			if inst, ok := inst.(value.Named); ok {
				// Ignore local value if of type void.
				if oldInst, ok := oldInst.(*ast.InstCall); ok {
					if _, ok := oldInst.Type.(*ast.VoidType); ok {
						continue
					}
				}
				m.locals[inst.GetName()] = inst
			}
		}
	}

	// Fix basic blocks.
	for i := 0; i < len(oldFunc.Blocks); i++ {
		oldBlock := oldFunc.Blocks[i]
		block := f.Blocks[i]
		m.basicBlock(oldBlock, block)
	}
}

// === [ Metadata definitions ] ================================================

// metadataDef translates the given metadata definition to LLVM IR, emitting
// code to m.
func (m *Module) metadataDef(oldMetadata *ast.Metadata) {
	md := m.getMetadata(oldMetadata.ID)
	for _, oldNode := range oldMetadata.Nodes {
		node := m.metadataNode(oldNode)
		md.Nodes = append(md.Nodes, node)
	}
}

// metadataNode returns the corresponding LLVM IR metadata node of the given
// metadata node.
func (m *Module) metadataNode(oldNode ast.MetadataNode) ir.MetadataNode {
	switch oldNode := oldNode.(type) {
	case *ast.Metadata:
		if len(oldNode.ID) > 0 {
			return m.getMetadata(oldNode.ID)
		}
		// Unnamed metadata literal.
		md := &ir.Metadata{}
		for _, node := range oldNode.Nodes {
			n := m.metadataNode(node)
			md.Nodes = append(md.Nodes, n)
		}
		return md
	case *ast.MetadataString:
		return &ir.MetadataString{
			Val: oldNode.Val,
		}
	case ast.Constant:
		c := m.irConstant(oldNode)
		md, ok := c.(ir.MetadataNode)
		if !ok {
			panic(fmt.Sprintf("invalid metadata node type; expected ir.MetadataNode, got %T", c))
		}
		return md
	default:
		panic(fmt.Errorf("support for metadata node type %T not yet implemented", oldNode))
	}
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

// basicBlock translates the given basic block to LLVM IR, emitting code to m.
func (m *Module) basicBlock(oldBlock *ast.BasicBlock, block *ir.BasicBlock) {
	// Fix instructions.
	for i := 0; i < len(oldBlock.Insts); i++ {
		oldInst := oldBlock.Insts[i]
		v := block.Insts[i]
		switch oldInst := oldInst.(type) {
		// Binary instructions
		case *ast.InstAdd:
			inst, ok := v.(*ir.InstAdd)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstAdd, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstFAdd:
			inst, ok := v.(*ir.InstFAdd)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFAdd, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstSub:
			inst, ok := v.(*ir.InstSub)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstSub, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstFSub:
			inst, ok := v.(*ir.InstFSub)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFSub, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstMul:
			inst, ok := v.(*ir.InstMul)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstMul, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstFMul:
			inst, ok := v.(*ir.InstFMul)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFMul, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstUDiv:
			inst, ok := v.(*ir.InstUDiv)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstUDiv, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstSDiv:
			inst, ok := v.(*ir.InstSDiv)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstSDiv, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstFDiv:
			inst, ok := v.(*ir.InstFDiv)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFDiv, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstURem:
			inst, ok := v.(*ir.InstURem)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstURem, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstSRem:
			inst, ok := v.(*ir.InstSRem)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstSRem, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstFRem:
			inst, ok := v.(*ir.InstFRem)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFRem, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)

		// Bitwise instructions
		case *ast.InstShl:
			inst, ok := v.(*ir.InstShl)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstShl, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstLShr:
			inst, ok := v.(*ir.InstLShr)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstLShr, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstAShr:
			inst, ok := v.(*ir.InstAShr)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstAShr, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstAnd:
			inst, ok := v.(*ir.InstAnd)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstAnd, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstOr:
			inst, ok := v.(*ir.InstOr)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstOr, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstXor:
			inst, ok := v.(*ir.InstXor)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstXor, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)

		// Vector instructions
		case *ast.InstExtractElement:
			inst, ok := v.(*ir.InstExtractElement)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstExtractElement, got %T", v))
			}
			x := m.irValue(oldInst.X)
			t, ok := x.Type().(*types.VectorType)
			if !ok {
				panic(fmt.Errorf("invalid vector type; expected *types.VectorType, got %T", x.Type()))
			}
			inst.Typ = t.Elem
			inst.X = x
			inst.Index = m.irValue(oldInst.Index)
		case *ast.InstInsertElement:
			inst, ok := v.(*ir.InstInsertElement)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstInsertElement, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Elem = m.irValue(oldInst.Elem)
			inst.Index = m.irValue(oldInst.Index)
		case *ast.InstShuffleVector:
			inst, ok := v.(*ir.InstShuffleVector)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstShuffleVector, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
			inst.Mask = m.irValue(oldInst.Mask)

		// Aggregate instructions
		case *ast.InstExtractValue:
			inst, ok := v.(*ir.InstExtractValue)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstExtractValue, got %T", v))
			}
			x := m.irValue(oldInst.X)
			typ := aggregateElemType(x.Type(), oldInst.Indices)
			inst.Typ = typ
			inst.X = x
			inst.Indices = oldInst.Indices
		case *ast.InstInsertValue:
			inst, ok := v.(*ir.InstInsertValue)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstInsertValue, got %T", v))
			}
			inst.X = m.irValue(oldInst.X)
			inst.Elem = m.irValue(oldInst.Elem)
			inst.Indices = oldInst.Indices

		// Memory instructions
		case *ast.InstAlloca:
			inst, ok := v.(*ir.InstAlloca)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstAlloca, got %T", v))
			}
			elem := m.irType(oldInst.Elem)
			typ := types.NewPointer(elem)
			inst.Typ = typ
			inst.Elem = elem
			if oldInst.NElems != nil {
				inst.NElems = m.irValue(oldInst.NElems)
			}
		case *ast.InstLoad:
			inst, ok := v.(*ir.InstLoad)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstLoad, got %T", v))
			}
			src := m.irValue(oldInst.Src)
			srcType, ok := src.Type().(*types.PointerType)
			if !ok {
				panic(fmt.Errorf("invalid source type; expected *types.PointerType, got %T", src.Type()))
			}
			typ := srcType.Elem
			if got, want := typ, m.irType(oldInst.Elem); !got.Equal(want) {
				m.errs = append(m.errs, errors.Errorf("source element type mismatch; expected `%v`, got `%v`", want, got))
			}
			inst.Typ = typ
			inst.Src = src
		case *ast.InstStore:
			inst, ok := v.(*ir.InstStore)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstStore, got %T", v))
			}
			inst.Src = m.irValue(oldInst.Src)
			inst.Dst = m.irValue(oldInst.Dst)
		case *ast.InstGetElementPtr:
			inst, ok := v.(*ir.InstGetElementPtr)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstGetElementPtr, got %T", v))
			}
			src := m.irValue(oldInst.Src)
			srcType, ok := src.Type().(*types.PointerType)
			if !ok {
				m.errs = append(m.errs, errors.Errorf("invalid source type; expected *types.PointerType, got %T", src.Type()))
			}
			elem := srcType.Elem
			if got, want := elem, m.irType(oldInst.Elem); !got.Equal(want) {
				m.errs = append(m.errs, errors.Errorf("source element type mismatch; expected `%v`, got `%v`", want, got))
			}
			var indices []value.Value
			for _, oldIndex := range oldInst.Indices {
				index := m.irValue(oldIndex)
				indices = append(indices, index)
			}
			e := elem
			for i, index := range indices {
				if i == 0 {
					// Ignore checking the 0th index as it simply follows the pointer of
					// src.
					//
					// ref: http://llvm.org/docs/GetElementPtr.html#why-is-the-extra-0-index-required
					continue
				}
				switch t := e.(type) {
				case *types.PointerType:
					// ref: http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep
					panic("unable to index into element of pointer type; for more information, see http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep")
				case *types.ArrayType:
					e = t.Elem
				case *types.StructType:
					idx, ok := index.(*constant.Int)
					if !ok {
						panic(fmt.Errorf("invalid index type for structure element; expected *constant.Int, got %T", index))
					}
					e = t.Fields[idx.Int64()]
				default:
					panic(fmt.Errorf("support for indexing element type %T not yet implemented", e))
				}
			}
			typ := types.NewPointer(e)
			inst.Typ = typ
			inst.Elem = elem
			inst.Src = src
			inst.Indices = indices

		// Conversion instructions
		case *ast.InstTrunc:
			inst, ok := v.(*ir.InstTrunc)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstTrunc, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstZExt:
			inst, ok := v.(*ir.InstZExt)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstZExt, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstSExt:
			inst, ok := v.(*ir.InstSExt)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstSExt, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstFPTrunc:
			inst, ok := v.(*ir.InstFPTrunc)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFPTrunc, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstFPExt:
			inst, ok := v.(*ir.InstFPExt)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFPExt, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstFPToUI:
			inst, ok := v.(*ir.InstFPToUI)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFPToUI, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstFPToSI:
			inst, ok := v.(*ir.InstFPToSI)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFPToSI, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstUIToFP:
			inst, ok := v.(*ir.InstUIToFP)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstUIToFP, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstSIToFP:
			inst, ok := v.(*ir.InstSIToFP)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstSIToFP, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstPtrToInt:
			inst, ok := v.(*ir.InstPtrToInt)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstPtrToInt, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstIntToPtr:
			inst, ok := v.(*ir.InstIntToPtr)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstIntToPtr, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstBitCast:
			inst, ok := v.(*ir.InstBitCast)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstBitCast, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)
		case *ast.InstAddrSpaceCast:
			inst, ok := v.(*ir.InstAddrSpaceCast)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstAddrSpaceCast, got %T", v))
			}
			inst.From = m.irValue(oldInst.From)
			inst.To = m.irType(oldInst.To)

		// Other instructions
		case *ast.InstICmp:
			inst, ok := v.(*ir.InstICmp)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstICmp, got %T", v))
			}
			pred := irIntPred(oldInst.Pred)
			x := m.irValue(oldInst.X)
			y := m.irValue(oldInst.Y)
			var typ types.Type = types.I1
			if t, ok := x.Type().(*types.VectorType); ok {
				typ = types.NewVector(types.I1, t.Len)
			}
			inst.Typ = typ
			inst.Pred = pred
			inst.X = x
			inst.Y = y
		case *ast.InstFCmp:
			inst, ok := v.(*ir.InstFCmp)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstFCmp, got %T", v))
			}
			pred := irFloatPred(oldInst.Pred)
			x := m.irValue(oldInst.X)
			y := m.irValue(oldInst.Y)
			var typ types.Type = types.I1
			if t, ok := x.Type().(*types.VectorType); ok {
				typ = types.NewVector(types.I1, t.Len)
			}
			inst.Typ = typ
			inst.Pred = pred
			inst.X = x
			inst.Y = y
		case *ast.InstPhi:
			inst, ok := v.(*ir.InstPhi)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstPhi, got %T", v))
			}
			inst.Typ = m.irType(oldInst.Type)
			for _, oldInc := range oldInst.Incs {
				x := m.irValue(oldInc.X)
				v := m.getLocal(oldInc.Pred.GetName())
				pred, ok := v.(*ir.BasicBlock)
				if !ok {
					panic(fmt.Errorf("invalid basic block type; expected *ir.BasicBlock, got %T", v))
				}
				inc := &ir.Incoming{
					X:    x,
					Pred: pred,
				}
				inst.Incs = append(inst.Incs, inc)
			}
		case *ast.InstSelect:
			inst, ok := v.(*ir.InstSelect)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstSelect, got %T", v))
			}
			inst.Cond = m.irValue(oldInst.Cond)
			inst.X = m.irValue(oldInst.X)
			inst.Y = m.irValue(oldInst.Y)
		case *ast.InstCall:
			inst, ok := v.(*ir.InstCall)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstCall, got %T", v))
			}
			callee := m.irValue(oldInst.Callee)
			typ, ok := callee.Type().(*types.PointerType)
			if !ok {
				panic(fmt.Errorf("invalid callee type, expected *types.PointerType, got %T", callee.Type()))
			}
			sig, ok := typ.Elem.(*types.FuncType)
			if !ok {
				panic(fmt.Errorf("invalid callee signature type, expected *types.FuncType, got %T", typ.Elem))
			}
			inst.Callee = callee
			inst.Sig = sig
			// TODO: Validate oldInst.Type against inst.Sig.
			for _, oldArg := range oldInst.Args {
				arg := m.irValue(oldArg)
				inst.Args = append(inst.Args, arg)
			}

		default:
			panic(fmt.Errorf("support for instruction %T not yet implemented", oldInst))
		}
	}

	// Fix terminator.
	switch oldTerm := oldBlock.Term.(type) {
	case *ast.TermRet:
		term := &ir.TermRet{
			Parent: block,
		}
		if oldTerm.X != nil {
			term.X = m.irValue(oldTerm.X)
		}
		block.Term = term
	case *ast.TermBr:
		term := &ir.TermBr{
			Parent: block,
		}
		v := m.irValue(oldTerm.Target)
		target, ok := v.(*ir.BasicBlock)
		if !ok {
			panic(fmt.Errorf("invalid target branch type, expected *ir.BasicBlock, got %T", v))
		}
		term.Target = target
		block.Term = term
	case *ast.TermCondBr:
		term := &ir.TermCondBr{
			Parent: block,
		}
		tTrue := m.irValue(oldTerm.TargetTrue)
		targetTrue, ok := tTrue.(*ir.BasicBlock)
		if !ok {
			panic(fmt.Errorf("invalid true target branch type, expected *ir.BasicBlock, got %T", tTrue))
		}
		tFalse := m.irValue(oldTerm.TargetFalse)
		targetFalse, ok := tFalse.(*ir.BasicBlock)
		if !ok {
			panic(fmt.Errorf("invalid false target branch type, expected *ir.BasicBlock, got %T", tFalse))
		}
		successors := []*ir.BasicBlock{targetTrue, targetFalse}
		term.Cond = m.irValue(oldTerm.Cond)
		term.TargetTrue = targetTrue
		term.TargetFalse = targetFalse
		term.Successors = successors
		block.Term = term
	case *ast.TermSwitch:
		term := &ir.TermSwitch{
			Parent: block,
		}
		term.X = m.irValue(oldTerm.X)
		v := m.getLocal(oldTerm.TargetDefault.GetName())
		targetDefault, ok := v.(*ir.BasicBlock)
		if !ok {
			panic(fmt.Errorf("invalid default target branch type, expected *ir.BasicBlock, got %T", v))
		}
		term.TargetDefault = targetDefault
		successors := []*ir.BasicBlock{targetDefault}
		for _, oldCase := range oldTerm.Cases {
			xx := m.irConstant(oldCase.X)
			x, ok := xx.(*constant.Int)
			if !ok {
				panic(fmt.Errorf("invalid x type, expected *constant.Int, got %T", xx))
			}
			v := m.getLocal(oldCase.Target.GetName())
			target, ok := v.(*ir.BasicBlock)
			if !ok {
				panic(fmt.Errorf("invalid target branch type, expected *ir.BasicBlock, got %T", v))
			}
			c := &ir.Case{
				X:      x,
				Target: target,
			}
			term.Cases = append(term.Cases, c)
			successors = append(successors, target)
		}
		term.Successors = successors
		block.Term = term
	case *ast.TermUnreachable:
		term := &ir.TermUnreachable{
			Parent: block,
		}
		block.Term = term
	default:
		panic(fmt.Errorf("support for terminator %T not yet implemented", oldTerm))
	}
}

// === [ Instructions ] ========================================================

// --- [ Binary instructions ] -------------------------------------------------

// --- [ Bitwise instructions ] ------------------------------------------------

// --- [ Memory instructions ] -------------------------------------------------

// --- [ Conversion instructions ] ---------------------------------------------

// --- [ Other instructions ] --------------------------------------------------

// === [ Terminators ] =========================================================
