// Translates AST values as follows.
//
// Per module.
//
//    1. Index type definitions.
//    2. Index global variables.
//       - Store preliminary content type.
//    3. Index function.
//       - Store type.
//    4. Index metadata.
//    5. Fix type definitions.
//    6. Fix globals.
//    7. Fix named metadata definition.
//    8. Fix metadata definition.
//    9. Fix functions.
//
// Per function.
//
//    1. Index function parameters.
//    2. Index basic blocks.
//    3. Index local variables produced by instructions.
//    4. Resolve locals.
//    5. Fix basic blocks.

package irx

import (
	"fmt"
	"sort"

	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// === [ Modules ] =============================================================

// Translate translates the AST of the given module to an equivalent LLVM IR
// module.
func Translate(module *ast.Module) (*ir.Module, error) {
	m := NewModule()

	// Set target specifiers.
	m.DataLayout = module.DataLayout
	m.TargetTriple = module.TargetTriple

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
		// Store preliminary content type (for circular dependencies).
		content := m.irType(old.Content)
		global := &ir.Global{
			Name:     name,
			Typ:      types.NewPointer(content),
			Content:  content,
			Metadata: make(map[string]*metadata.Metadata),
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
			Metadata: make(map[string]*metadata.Metadata),
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
		md := &metadata.Metadata{
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

	// Fix named metadata definitions.
	for _, old := range module.NamedMetadata {
		md := &metadata.Named{
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
		typ.AddrSpace = d.AddrSpace
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
		global.Content = init.Type()
		global.Init = init
	} else {
		global.Content = m.irType(old.Content)
	}
	typ := types.NewPointer(global.Content)
	typ.AddrSpace = old.AddrSpace
	global.Typ = typ
	global.IsConst = old.Immutable

	// Fix attached metadata.
	global.Metadata = m.irMetadata(old.Metadata)
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

	// Fix calling convention.
	f.CallConv = ir.CallConv(oldFunc.CallConv)

	// Fix attached metadata.
	f.Metadata = m.irMetadata(oldFunc.Metadata)

	// Early exit if function declaration.
	if len(oldFunc.Blocks) < 1 {
		return
	}

	// Reset locals.
	m.locals = make(map[string]value.Named)
	// track resolved and unresolved local values.
	var (
		resolved   = make(map[ast.NamedValue]value.Named)
		unresolved = make(map[ast.NamedValue]value.Named)
	)

	// Index function parameters.
	for i, param := range f.Params() {
		name := param.Name
		if _, ok := m.locals[name]; ok {
			panic(fmt.Errorf("local identifier %q already present for function %s; old `%v`, new `%v`", name, f.Ident(), m.locals[name], param))
		}
		m.locals[name] = param
		resolved[oldFunc.Sig.Params[i]] = param
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
		resolved[old] = block
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
					switch t := oldInst.Type.(type) {
					case *ast.VoidType:
						continue
					case *ast.FuncType:
						if _, ok := t.Ret.(*ast.VoidType); ok {
							continue
						}
					}
				}
				m.locals[inst.GetName()] = inst
				oi, ok := oldInst.(ast.NamedValue)
				if !ok {
					panic(fmt.Errorf("invalid old instruction type; expected ast.NamedValue, got %T", oldInst))
				}
				unresolved[oi] = inst
			}
		}
	}

	// Resolve locals (as they can have circular dependencies).
	for len(unresolved) > 0 {
		prev := len(unresolved)
		for old, local := range unresolved {
			if m.resolveInst(old, resolved, unresolved) {
				delete(unresolved, old)
				resolved[old] = local
			}
		}
		if len(unresolved) == prev {
			var names []string
			for old := range unresolved {
				names = append(names, old.GetName())
			}
			sort.Strings(names)
			panic(fmt.Errorf("unable to resolve %d local values; %v", len(unresolved), names))
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
func (m *Module) metadataNode(oldNode ast.MetadataNode) metadata.Node {
	switch oldNode := oldNode.(type) {
	case *ast.Metadata:
		if len(oldNode.ID) > 0 {
			return m.getMetadata(oldNode.ID)
		}
		// Unnamed metadata literal.
		md := &metadata.Metadata{}
		for _, node := range oldNode.Nodes {
			n := m.metadataNode(node)
			md.Nodes = append(md.Nodes, n)
		}
		return md
	case *ast.MetadataString:
		return &metadata.String{
			Val: oldNode.Val,
		}
	case *ast.MetadataValue:
		return &metadata.Value{
			X: m.irValue(oldNode.X),
		}
	case ast.Constant:
		c := m.irConstant(oldNode)
		md, ok := c.(metadata.Node)
		if !ok {
			panic(fmt.Sprintf("invalid metadata node type; expected metadata.Node, got %T", c))
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

// --- [ Aggregate expressions ] -----------------------------------------------

// --- [ Vector expressions ] --------------------------------------------------

// --- [ Memory expressions ] --------------------------------------------------

// --- [ Conversion expressions ] ----------------------------------------------

// --- [ Other expressions ] ---------------------------------------------------

// === [ Basic blocks ] ========================================================

// basicBlock translates the given basic block to LLVM IR, emitting code to m.
func (m *Module) basicBlock(oldBlock *ast.BasicBlock, block *ir.BasicBlock) {
	// Fix instructions not producing values.
	for i := 0; i < len(oldBlock.Insts); i++ {
		old := oldBlock.Insts[i]
		v := block.Insts[i]
		switch old := old.(type) {
		// Memory instructions
		case *ast.InstStore:
			inst, ok := v.(*ir.InstStore)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstStore, got %T", v))
			}
			inst.Src = m.irValue(old.Src)
			inst.Dst = m.irValue(old.Dst)
			inst.Metadata = m.irMetadata(old.Metadata)

		// Other instructions
		case *ast.InstCall:
			inst, ok := v.(*ir.InstCall)
			if !ok {
				panic(fmt.Errorf("invalid instruction type; expected *ir.InstCall, got %T", v))
			}
			// Handle calls to void functions.
			switch t := old.Type.(type) {
			case *ast.VoidType:
			case *ast.FuncType:
				if _, ok := t.Ret.(*ast.VoidType); !ok {
					// handled by resolveInst.
					continue
				}
			default:
				// handled by resolveInst.
				continue
			}
			// TODO: Call m.instCall to reuse code.
			callee := m.irValue(old.Callee)
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
			// TODO: Validate old.Type against inst.Sig.
			for _, oldArg := range old.Args {
				arg := m.irValue(oldArg)
				inst.Args = append(inst.Args, arg)
			}
			inst.CallConv = ir.CallConv(old.CallConv)
			inst.Metadata = m.irMetadata(old.Metadata)
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
		term.Metadata = m.irMetadata(oldTerm.Metadata)
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
		term.Metadata = m.irMetadata(oldTerm.Metadata)
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
		term.Metadata = m.irMetadata(oldTerm.Metadata)
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
		term.Metadata = m.irMetadata(oldTerm.Metadata)
		block.Term = term
	case *ast.TermUnreachable:
		term := &ir.TermUnreachable{
			Parent: block,
		}
		term.Metadata = m.irMetadata(oldTerm.Metadata)
		block.Term = term
	default:
		panic(fmt.Errorf("support for terminator %T not yet implemented", oldTerm))
	}
}

// === [ Instructions ] ========================================================

// resolveInst resolves the given local value, by recursively resolving its
// operands.
func (m *Module) resolveInst(old ast.NamedValue, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	switch old := old.(type) {
	// Binary instructions
	case *ast.InstAdd:
		return m.instAdd(old, resolved, unresolved)
	case *ast.InstFAdd:
		return m.instFAdd(old, resolved, unresolved)
	case *ast.InstSub:
		return m.instSub(old, resolved, unresolved)
	case *ast.InstFSub:
		return m.instFSub(old, resolved, unresolved)
	case *ast.InstMul:
		return m.instMul(old, resolved, unresolved)
	case *ast.InstFMul:
		return m.instFMul(old, resolved, unresolved)
	case *ast.InstUDiv:
		return m.instUDiv(old, resolved, unresolved)
	case *ast.InstSDiv:
		return m.instSDiv(old, resolved, unresolved)
	case *ast.InstFDiv:
		return m.instFDiv(old, resolved, unresolved)
	case *ast.InstURem:
		return m.instURem(old, resolved, unresolved)
	case *ast.InstSRem:
		return m.instSRem(old, resolved, unresolved)
	case *ast.InstFRem:
		return m.instFRem(old, resolved, unresolved)

	// Bitwise instructions
	case *ast.InstShl:
		return m.instShl(old, resolved, unresolved)
	case *ast.InstLShr:
		return m.instLShr(old, resolved, unresolved)
	case *ast.InstAShr:
		return m.instAShr(old, resolved, unresolved)
	case *ast.InstAnd:
		return m.instAnd(old, resolved, unresolved)
	case *ast.InstOr:
		return m.instOr(old, resolved, unresolved)
	case *ast.InstXor:
		return m.instXor(old, resolved, unresolved)

	// Vector instructions
	case *ast.InstExtractElement:
		return m.instExtractElement(old, resolved, unresolved)
	case *ast.InstInsertElement:
		return m.instInsertElement(old, resolved, unresolved)
	case *ast.InstShuffleVector:
		return m.instShuffleVector(old, resolved, unresolved)

	// Aggregate instructions
	case *ast.InstExtractValue:
		return m.instExtractValue(old, resolved, unresolved)
	case *ast.InstInsertValue:
		return m.instInsertValue(old, resolved, unresolved)

	// Memory instructions
	case *ast.InstAlloca:
		return m.instAlloca(old, resolved, unresolved)
	case *ast.InstLoad:
		return m.instLoad(old, resolved, unresolved)
	case *ast.InstGetElementPtr:
		return m.instGetElementPtr(old, resolved, unresolved)

	// Conversion instructions
	case *ast.InstTrunc:
		return m.instTrunc(old, resolved, unresolved)
	case *ast.InstZExt:
		return m.instZExt(old, resolved, unresolved)
	case *ast.InstSExt:
		return m.instSExt(old, resolved, unresolved)
	case *ast.InstFPTrunc:
		return m.instFPTrunc(old, resolved, unresolved)
	case *ast.InstFPExt:
		return m.instFPExt(old, resolved, unresolved)
	case *ast.InstFPToUI:
		return m.instFPToUI(old, resolved, unresolved)
	case *ast.InstFPToSI:
		return m.instFPToSI(old, resolved, unresolved)
	case *ast.InstUIToFP:
		return m.instUIToFP(old, resolved, unresolved)
	case *ast.InstSIToFP:
		return m.instSIToFP(old, resolved, unresolved)
	case *ast.InstPtrToInt:
		return m.instPtrToInt(old, resolved, unresolved)
	case *ast.InstIntToPtr:
		return m.instIntToPtr(old, resolved, unresolved)
	case *ast.InstBitCast:
		return m.instBitCast(old, resolved, unresolved)
	case *ast.InstAddrSpaceCast:
		return m.instAddrSpaceCast(old, resolved, unresolved)

	// Other instructions
	case *ast.InstICmp:
		return m.instICmp(old, resolved, unresolved)
	case *ast.InstFCmp:
		return m.instFCmp(old, resolved, unresolved)
	case *ast.InstPhi:
		return m.instPhi(old, resolved, unresolved)
	case *ast.InstSelect:
		return m.instSelect(old, resolved, unresolved)
	case *ast.InstCall:
		return m.instCall(old, resolved, unresolved)

	default:
		panic(fmt.Errorf("support for instruction %T not yet implemented", old))
	}
}

// === [ Terminators ] =========================================================
