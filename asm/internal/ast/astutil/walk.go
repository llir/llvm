// Note, the AST traversal implementation of this package is heavily inspired by
// go fix; which is governed by a BSD license.

// Package astutil provides utility functions for interacting with ASTs of LLVM
// IR assembly.
package astutil

import (
	"fmt"

	"github.com/llir/llvm/asm/internal/ast"
)

// Walk traverses the AST x, calling visit(y) for each node y in the tree but
// also with a pointer to each ast.Type, ast.Value, and *ast.BasicBlock, in a
// bottom-up traversal.
func Walk(x interface{}, visit func(interface{})) {
	WalkBeforeAfter(x, nop, visit)
}

// WalkFunc traverses the AST of the given function, calling visit(y) for each
// node y in the tree but also with a pointer to each ast.Type, ast.Value, and
// *ast.BasicBlock, in a bottom-up traversal.
func WalkFunc(f *ast.Function, visit func(interface{})) {
	WalkFuncBeforeAfter(f, nop, visit)
}

// nop performs no operation on the given AST.
func nop(x interface{}) {
}

// WalkBeforeAfter traverses the AST x, calling before(y) before traversing y's
// children and after(y) afterward for each node y in the tree but also with a
// pointer to each ast.Type, ast.Value, and *ast.BasicBlock, in a bottom-up
// traversal.
func WalkBeforeAfter(x interface{}, before, after func(interface{})) {
	w := &walker{
		visited: make(map[interface{}]bool),
	}
	w.walkBeforeAfter(x, before, after)
}

// WalkFuncBeforeAfter traverses the AST of the given function, calling
// before(y) before traversing y's children and after(y) afterward for each node
// y in the tree but also with a pointer to each ast.Type, ast.Value, and
// *ast.BasicBlock, in a bottom-up traversal.
//
// Special precausion is taken during traversal to stay within the scope of the
// function.
func WalkFuncBeforeAfter(f *ast.Function, before, after func(interface{})) {
	w := &walker{
		funcScope: true,
		visited:   make(map[interface{}]bool),
	}
	// Traverse child nodes of function, instead of f directly, as *ast.Function
	// nodes are not traversed when staying within the scope of the function.
	w.walkBeforeAfter(&f.Sig, before, after)
	if f.Blocks != nil {
		w.walkBeforeAfter(&f.Blocks, before, after)
	}
}

// A walker traverses ASTs of LLVM IR while preventing infinite loops.
type walker struct {
	// Specifies whether to stay within the scope of the function during
	// traversal.
	funcScope bool
	// visited keeps track of visited nodes to prevent infinite loops.
	visited map[interface{}]bool
}

// walkBeforeAfter traverses the AST x, calling before(y) before traversing y's
// children and after(y) afterward for each node y in the tree but also with a
// pointer to each ast.Type, ast.Value, and *ast.BasicBlock, in a bottom-up
// traversal.
func (w *walker) walkBeforeAfter(x interface{}, before, after func(interface{})) {
	switch x.(type) {
	case []*ast.Global, []*ast.Function, []*ast.Param, []ast.Type, []*ast.NamedType, []ast.Value, []ast.Constant, []*ast.BasicBlock, []ast.Instruction, []*ast.Incoming, []*ast.Case:
		// unhashable type.
	case *ast.Function:
		if w.funcScope {
			// *ast.Function nodes are not traversed when staying within the scope
			// *of the function.
			return
		}
	default:
		// Prevent infinite loops.

		// TODO: Check if it is enough to only track *ast.NamedType to prevent inf
		// loops.
		if w.visited[x] {
			return
		}
		w.visited[x] = true
	}

	before(x)

	switch n := x.(type) {
	// pointers to interfaces
	case *ast.Type:
		w.walkBeforeAfter(*n, before, after)
	case *ast.Value:
		w.walkBeforeAfter(*n, before, after)
	case *ast.NamedValue:
		w.walkBeforeAfter(*n, before, after)
	case *ast.Constant:
		w.walkBeforeAfter(*n, before, after)
	case *ast.Instruction:
		w.walkBeforeAfter(*n, before, after)
	case *ast.Terminator:
		w.walkBeforeAfter(*n, before, after)

	// pointers to struct pointers
	case **ast.Global:
		w.walkBeforeAfter(*n, before, after)
	case **ast.Function:
		w.walkBeforeAfter(*n, before, after)
	case **ast.Param:
		w.walkBeforeAfter(*n, before, after)
	case **ast.GlobalDummy:
		w.walkBeforeAfter(*n, before, after)
	case **ast.LocalDummy:
		w.walkBeforeAfter(*n, before, after)
	// Types
	case **ast.VoidType:
		w.walkBeforeAfter(*n, before, after)
	case **ast.LabelType:
		w.walkBeforeAfter(*n, before, after)
	case **ast.MetadataType:
		w.walkBeforeAfter(*n, before, after)
	case **ast.IntType:
		w.walkBeforeAfter(*n, before, after)
	case **ast.FloatType:
		w.walkBeforeAfter(*n, before, after)
	case **ast.FuncType:
		w.walkBeforeAfter(*n, before, after)
	case **ast.PointerType:
		w.walkBeforeAfter(*n, before, after)
	case **ast.VectorType:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ArrayType:
		w.walkBeforeAfter(*n, before, after)
	case **ast.StructType:
		w.walkBeforeAfter(*n, before, after)
	case **ast.NamedType:
		w.walkBeforeAfter(*n, before, after)
	case **ast.NamedTypeDummy:
		w.walkBeforeAfter(*n, before, after)
	// Constants
	case **ast.IntConst:
		w.walkBeforeAfter(*n, before, after)
	case **ast.FloatConst:
		w.walkBeforeAfter(*n, before, after)
	case **ast.NullConst:
		w.walkBeforeAfter(*n, before, after)
	case **ast.VectorConst:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ArrayConst:
		w.walkBeforeAfter(*n, before, after)
	case **ast.CharArrayConst:
		w.walkBeforeAfter(*n, before, after)
	case **ast.StructConst:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ZeroInitializerConst:
		w.walkBeforeAfter(*n, before, after)
	// Constant expressions
	case **ast.ExprAdd:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprFAdd:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprSub:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprFSub:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprMul:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprFMul:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprUDiv:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprSDiv:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprFDiv:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprURem:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprSRem:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprFRem:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprShl:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprLShr:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprAShr:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprAnd:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprOr:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprXor:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprGetElementPtr:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprTrunc:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprZExt:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprSExt:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprFPTrunc:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprFPExt:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprFPToUI:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprFPToSI:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprUIToFP:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprSIToFP:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprPtrToInt:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprIntToPtr:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprBitCast:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprAddrSpaceCast:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprICmp:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprFCmp:
		w.walkBeforeAfter(*n, before, after)
	case **ast.ExprSelect:
		w.walkBeforeAfter(*n, before, after)
	// Basic blocks.
	case **ast.BasicBlock:
		w.walkBeforeAfter(*n, before, after)
	// Instructions
	case **ast.InstAdd:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstFAdd:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstSub:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstFSub:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstMul:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstFMul:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstUDiv:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstSDiv:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstFDiv:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstURem:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstSRem:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstFRem:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstShl:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstLShr:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstAShr:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstAnd:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstOr:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstXor:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstAlloca:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstLoad:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstStore:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstGetElementPtr:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstTrunc:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstZExt:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstSExt:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstFPTrunc:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstFPExt:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstFPToUI:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstFPToSI:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstUIToFP:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstSIToFP:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstPtrToInt:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstIntToPtr:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstBitCast:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstAddrSpaceCast:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstICmp:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstFCmp:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstPhi:
		w.walkBeforeAfter(*n, before, after)
	case **ast.Incoming:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstSelect:
		w.walkBeforeAfter(*n, before, after)
	case **ast.InstCall:
		w.walkBeforeAfter(*n, before, after)
	// Terminators
	case **ast.TermRet:
		w.walkBeforeAfter(*n, before, after)
	case **ast.TermBr:
		w.walkBeforeAfter(*n, before, after)
	case **ast.TermCondBr:
		w.walkBeforeAfter(*n, before, after)
	case **ast.TermSwitch:
		w.walkBeforeAfter(*n, before, after)
	case **ast.Case:
		w.walkBeforeAfter(*n, before, after)
	case **ast.TermUnreachable:
		w.walkBeforeAfter(*n, before, after)

	// pointers to slices
	case *[]ast.Type:
		w.walkBeforeAfter(*n, before, after)
	case *[]*ast.NamedType:
		w.walkBeforeAfter(*n, before, after)
	case *[]*ast.Global:
		w.walkBeforeAfter(*n, before, after)
	case *[]ast.Value:
		w.walkBeforeAfter(*n, before, after)
	case *[]ast.Constant:
		w.walkBeforeAfter(*n, before, after)
	case *[]*ast.Function:
		w.walkBeforeAfter(*n, before, after)
	case *[]*ast.Param:
		w.walkBeforeAfter(*n, before, after)
	case *[]*ast.BasicBlock:
		w.walkBeforeAfter(*n, before, after)
	case *[]ast.Instruction:
		w.walkBeforeAfter(*n, before, after)
	case *[]*ast.Case:
		w.walkBeforeAfter(*n, before, after)
	case *[]*ast.Incoming:
		w.walkBeforeAfter(*n, before, after)

	// These are ordered and grouped to match ../../ll.bnf
	case *ast.Module:
		if n.Types != nil {
			w.walkBeforeAfter(&n.Types, before, after)
		}
		if n.Globals != nil {
			w.walkBeforeAfter(&n.Globals, before, after)
		}
		if n.Funcs != nil {
			w.walkBeforeAfter(&n.Funcs, before, after)
		}
	case []*ast.Global:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ast.Global:
		w.walkBeforeAfter(&n.Content, before, after)
		if n.Init != nil {
			w.walkBeforeAfter(&n.Init, before, after)
		}
	case []*ast.Function:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ast.Function:
		w.walkBeforeAfter(&n.Sig, before, after)
		if n.Blocks != nil {
			w.walkBeforeAfter(&n.Blocks, before, after)
		}
	case []*ast.Param:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ast.Param:
		w.walkBeforeAfter(&n.Type, before, after)
	case *ast.GlobalDummy:
		w.walkBeforeAfter(&n.Type, before, after)
	case *ast.LocalDummy:
		w.walkBeforeAfter(&n.Type, before, after)
	// Types
	case []ast.Type:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ast.TypeDummy:
		// nothing to do.
	case *ast.VoidType:
		// nothing to do.
	case *ast.LabelType:
		// nothing to do.
	case *ast.MetadataType:
		// nothing to do.
	case *ast.IntType:
		// nothing to do.
	case *ast.FloatType:
		// nothing to do.
	case *ast.FuncType:
		w.walkBeforeAfter(&n.Ret, before, after)
		if n.Params != nil {
			w.walkBeforeAfter(&n.Params, before, after)
		}
	case *ast.PointerType:
		w.walkBeforeAfter(&n.Elem, before, after)
	case *ast.VectorType:
		w.walkBeforeAfter(&n.Elem, before, after)
	case *ast.ArrayType:
		w.walkBeforeAfter(&n.Elem, before, after)
	case *ast.StructType:
		if n.Fields != nil {
			w.walkBeforeAfter(&n.Fields, before, after)
		}
	case []*ast.NamedType:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ast.NamedType:
		if n.Def != nil {
			w.walkBeforeAfter(&n.Def, before, after)
		}
	case *ast.NamedTypeDummy:
		// nothing to do.
	// Constants
	case []ast.Value:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case []ast.Constant:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ast.IntConst:
		w.walkBeforeAfter(&n.Type, before, after)
	case *ast.FloatConst:
		w.walkBeforeAfter(&n.Type, before, after)
	case *ast.NullConst:
		w.walkBeforeAfter(&n.Type, before, after)
	case *ast.VectorConst:
		w.walkBeforeAfter(&n.Type, before, after)
		if n.Elems != nil {
			w.walkBeforeAfter(&n.Elems, before, after)
		}
	case *ast.ArrayConst:
		w.walkBeforeAfter(&n.Type, before, after)
		if n.Elems != nil {
			w.walkBeforeAfter(&n.Elems, before, after)
		}
	case *ast.CharArrayConst:
		w.walkBeforeAfter(&n.Type, before, after)
	case *ast.StructConst:
		w.walkBeforeAfter(&n.Type, before, after)
		if n.Fields != nil {
			w.walkBeforeAfter(&n.Fields, before, after)
		}
	case *ast.ZeroInitializerConst:
		w.walkBeforeAfter(&n.Type, before, after)
	// Constant expressions
	case *ast.ExprAdd:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprFAdd:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprSub:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprFSub:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprMul:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprFMul:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprUDiv:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprSDiv:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprFDiv:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprURem:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprSRem:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprFRem:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprShl:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprLShr:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprAShr:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprAnd:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprOr:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprXor:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprGetElementPtr:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.Elem, before, after)
		w.walkBeforeAfter(&n.Src, before, after)
		if n.Indices != nil {
			w.walkBeforeAfter(&n.Indices, before, after)
		}
	case *ast.ExprTrunc:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.ExprZExt:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.ExprSExt:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.ExprFPTrunc:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.ExprFPExt:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.ExprFPToUI:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.ExprFPToSI:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.ExprUIToFP:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.ExprSIToFP:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.ExprPtrToInt:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.ExprIntToPtr:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.ExprBitCast:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.ExprAddrSpaceCast:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.ExprICmp:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprFCmp:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.ExprSelect:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.Cond, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	// Basic blocks.
	case []*ast.BasicBlock:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ast.BasicBlock:
		if n.Insts != nil {
			w.walkBeforeAfter(&n.Insts, before, after)
		}
		w.walkBeforeAfter(&n.Term, before, after)
	// Instructions
	case []ast.Instruction:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ast.InstAdd:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstFAdd:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstSub:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstFSub:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstMul:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstFMul:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstUDiv:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstSDiv:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstFDiv:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstURem:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstSRem:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstFRem:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstShl:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstLShr:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstAShr:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstAnd:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstOr:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstXor:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstAlloca:
		w.walkBeforeAfter(&n.Elem, before, after)
		if n.NElems != nil {
			w.walkBeforeAfter(&n.NElems, before, after)
		}
	case *ast.InstLoad:
		w.walkBeforeAfter(&n.Elem, before, after)
		w.walkBeforeAfter(&n.Src, before, after)
	case *ast.InstStore:
		w.walkBeforeAfter(&n.Src, before, after)
		w.walkBeforeAfter(&n.Dst, before, after)
	case *ast.InstGetElementPtr:
		w.walkBeforeAfter(&n.Elem, before, after)
		w.walkBeforeAfter(&n.Src, before, after)
		if n.Indices != nil {
			w.walkBeforeAfter(&n.Indices, before, after)
		}
	case *ast.InstTrunc:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.InstZExt:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.InstSExt:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.InstFPTrunc:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.InstFPExt:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.InstFPToUI:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.InstFPToSI:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.InstUIToFP:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.InstSIToFP:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.InstPtrToInt:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.InstIntToPtr:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.InstBitCast:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.InstAddrSpaceCast:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ast.InstICmp:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstFCmp:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstPhi:
		w.walkBeforeAfter(&n.Type, before, after)
		if n.Incs != nil {
			w.walkBeforeAfter(&n.Incs, before, after)
		}
	case []*ast.Incoming:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ast.Incoming:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Pred, before, after)
	case *ast.InstSelect:
		w.walkBeforeAfter(&n.Cond, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ast.InstCall:
		w.walkBeforeAfter(&n.Type, before, after)
		w.walkBeforeAfter(&n.Callee, before, after)
		if n.Args != nil {
			w.walkBeforeAfter(&n.Args, before, after)
		}
	// Terminators
	case *ast.TermRet:
		if n.X != nil {
			w.walkBeforeAfter(&n.X, before, after)
		}
	case *ast.TermBr:
		w.walkBeforeAfter(&n.Target, before, after)
	case *ast.TermCondBr:
		w.walkBeforeAfter(&n.Cond, before, after)
		w.walkBeforeAfter(&n.TargetTrue, before, after)
		w.walkBeforeAfter(&n.TargetFalse, before, after)
	case *ast.TermSwitch:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.TargetDefault, before, after)
		if n.Cases != nil {
			w.walkBeforeAfter(&n.Cases, before, after)
		}
	case []*ast.Case:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ast.Case:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Target, before, after)
	case *ast.TermUnreachable:
		// nothing to do.

	default:
		panic(fmt.Errorf("support for type %T not yet implemented", x))
	}

	after(x)
}
