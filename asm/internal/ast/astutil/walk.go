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

// A walker traverses ASTs of LLVM IR while preventing infinite loops.
type walker struct {
	// visited keeps track of visited nodes to prevent infinite loops.
	visited map[interface{}]bool
}

// walkBeforeAfter traverses the AST x, calling before(y) before traversing y's
// children and after(y) afterward for each node y in the tree but also with a
// pointer to each ast.Type, ast.Value, and *ast.BasicBlock, in a bottom-up
// traversal.
func (w *walker) walkBeforeAfter(x interface{}, before, after func(interface{})) {
	// Prevent infinite loops.
	if w.visited[x] {
		return
	}
	w.visited[x] = true

	before(x)

	switch n := x.(type) {
	// pointers to interfaces
	case *ast.Type:
		WalkBeforeAfter(*n, before, after)
	case *ast.Value:
		WalkBeforeAfter(*n, before, after)
	case *ast.NamedValue:
		WalkBeforeAfter(*n, before, after)
	case *ast.Constant:
		WalkBeforeAfter(*n, before, after)
	case *ast.Instruction:
		WalkBeforeAfter(*n, before, after)
	case *ast.Terminator:
		WalkBeforeAfter(*n, before, after)

	// pointers to struct pointers
	case **ast.Module:
		WalkBeforeAfter(*n, before, after)
	case **ast.Global:
		WalkBeforeAfter(*n, before, after)
	case **ast.Function:
		WalkBeforeAfter(*n, before, after)
	case **ast.Param:
		WalkBeforeAfter(*n, before, after)
	case **ast.GlobalDummy:
		WalkBeforeAfter(*n, before, after)
	case **ast.LocalDummy:
		WalkBeforeAfter(*n, before, after)
	// Types
	case **ast.VoidType:
		WalkBeforeAfter(*n, before, after)
	case **ast.LabelType:
		WalkBeforeAfter(*n, before, after)
	case **ast.IntType:
		WalkBeforeAfter(*n, before, after)
	case **ast.FloatType:
		WalkBeforeAfter(*n, before, after)
	case **ast.FuncType:
		WalkBeforeAfter(*n, before, after)
	case **ast.PointerType:
		WalkBeforeAfter(*n, before, after)
	case **ast.VectorType:
		WalkBeforeAfter(*n, before, after)
	case **ast.ArrayType:
		WalkBeforeAfter(*n, before, after)
	case **ast.StructType:
		WalkBeforeAfter(*n, before, after)
	case **ast.NamedType:
		WalkBeforeAfter(*n, before, after)
	case **ast.NamedTypeDummy:
		WalkBeforeAfter(*n, before, after)
	// Constants
	case **ast.IntConst:
		WalkBeforeAfter(*n, before, after)
	case **ast.FloatConst:
		WalkBeforeAfter(*n, before, after)
	case **ast.NullConst:
		WalkBeforeAfter(*n, before, after)
	case **ast.VectorConst:
		WalkBeforeAfter(*n, before, after)
	case **ast.ArrayConst:
		WalkBeforeAfter(*n, before, after)
	case **ast.StructConst:
		WalkBeforeAfter(*n, before, after)
	case **ast.ZeroInitializerConst:
		WalkBeforeAfter(*n, before, after)
	// Constant expressions
	case **ast.ExprAdd:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprFAdd:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprSub:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprFSub:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprMul:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprFMul:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprUDiv:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprSDiv:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprFDiv:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprURem:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprSRem:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprFRem:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprShl:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprLShr:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprAShr:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprAnd:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprOr:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprXor:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprGetElementPtr:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprTrunc:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprZExt:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprSExt:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprFPTrunc:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprFPExt:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprFPToUI:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprFPToSI:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprUIToFP:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprSIToFP:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprPtrToInt:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprIntToPtr:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprBitCast:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprAddrSpaceCast:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprICmp:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprFCmp:
		WalkBeforeAfter(*n, before, after)
	case **ast.ExprSelect:
		WalkBeforeAfter(*n, before, after)
	case **ast.BasicBlock:
		WalkBeforeAfter(*n, before, after)
	// Instructions
	case **ast.InstAdd:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstFAdd:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstSub:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstFSub:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstMul:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstFMul:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstUDiv:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstSDiv:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstFDiv:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstURem:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstSRem:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstFRem:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstShl:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstLShr:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstAShr:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstAnd:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstOr:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstXor:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstAlloca:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstLoad:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstStore:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstGetElementPtr:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstTrunc:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstZExt:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstSExt:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstFPTrunc:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstFPExt:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstFPToUI:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstFPToSI:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstUIToFP:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstSIToFP:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstPtrToInt:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstIntToPtr:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstBitCast:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstAddrSpaceCast:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstICmp:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstFCmp:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstPhi:
		WalkBeforeAfter(*n, before, after)
	case **ast.Incoming:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstSelect:
		WalkBeforeAfter(*n, before, after)
	case **ast.InstCall:
		WalkBeforeAfter(*n, before, after)
	// Terminators
	case **ast.TermRet:
		WalkBeforeAfter(*n, before, after)
	case **ast.TermBr:
		WalkBeforeAfter(*n, before, after)
	case **ast.TermCondBr:
		WalkBeforeAfter(*n, before, after)
	case **ast.TermSwitch:
		WalkBeforeAfter(*n, before, after)
	case **ast.Case:
		WalkBeforeAfter(*n, before, after)
	case **ast.TermUnreachable:
		WalkBeforeAfter(*n, before, after)

	// pointers to slices
	case *[]ast.Type:
		WalkBeforeAfter(*n, before, after)
	case *[]*ast.NamedType:
		WalkBeforeAfter(*n, before, after)
	case *[]*ast.Global:
		WalkBeforeAfter(*n, before, after)
	case *[]ast.Value:
		WalkBeforeAfter(*n, before, after)
	case *[]ast.Constant:
		WalkBeforeAfter(*n, before, after)
	case *[]*ast.Function:
		WalkBeforeAfter(*n, before, after)
	case *[]*ast.Param:
		WalkBeforeAfter(*n, before, after)
	case *[]*ast.BasicBlock:
		WalkBeforeAfter(*n, before, after)
	case *[]ast.Instruction:
		WalkBeforeAfter(*n, before, after)
	case *[]*ast.Case:
		WalkBeforeAfter(*n, before, after)
	case *[]*ast.Incoming:
		WalkBeforeAfter(*n, before, after)

	// These are ordered and grouped to match ../../ll.bnf
	case *ast.Module:
		WalkBeforeAfter(&n.Types, before, after)
		WalkBeforeAfter(&n.Globals, before, after)
		WalkBeforeAfter(&n.Funcs, before, after)
	case []*ast.Global:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.Global:
		WalkBeforeAfter(&n.Content, before, after)
		if n.Init != nil {
			WalkBeforeAfter(&n.Init, before, after)
		}
	case []*ast.Function:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.Function:
		WalkBeforeAfter(&n.Sig, before, after)
		if n.Blocks != nil {
			WalkBeforeAfter(&n.Blocks, before, after)
		}
	case []*ast.Param:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.Param:
		WalkBeforeAfter(&n.Type, before, after)
	case *ast.GlobalDummy:
		WalkBeforeAfter(&n.Type, before, after)
	case *ast.LocalDummy:
		WalkBeforeAfter(&n.Type, before, after)
	// Types
	case []ast.Type:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.VoidType:
		// nothing to do.
	case *ast.LabelType:
		// nothing to do.
	case *ast.IntType:
		// nothing to do.
	case *ast.FloatType:
		// nothing to do.
	case *ast.FuncType:
		WalkBeforeAfter(&n.Ret, before, after)
		WalkBeforeAfter(&n.Params, before, after)
	case *ast.PointerType:
		WalkBeforeAfter(&n.Elem, before, after)
	case *ast.VectorType:
		WalkBeforeAfter(&n.Elem, before, after)
	case *ast.ArrayType:
		WalkBeforeAfter(&n.Elem, before, after)
	case *ast.StructType:
		WalkBeforeAfter(&n.Fields, before, after)
	case []*ast.NamedType:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.NamedType:
		WalkBeforeAfter(&n.Def, before, after)
	case *ast.NamedTypeDummy:
		// nothing to do.
	// Constants
	case []ast.Value:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case []ast.Constant:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.IntConst:
		WalkBeforeAfter(&n.Type, before, after)
	case *ast.FloatConst:
		WalkBeforeAfter(&n.Type, before, after)
	case *ast.NullConst:
		WalkBeforeAfter(&n.Type, before, after)
	case *ast.VectorConst:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.Elems, before, after)
	case *ast.ArrayConst:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.Elems, before, after)
	case *ast.StructConst:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.Fields, before, after)
	case *ast.ZeroInitializerConst:
		WalkBeforeAfter(&n.Type, before, after)
	// Constant expressions
	case *ast.ExprAdd:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprFAdd:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprSub:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprFSub:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprMul:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprFMul:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprUDiv:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprSDiv:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprFDiv:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprURem:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprSRem:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprFRem:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprShl:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprLShr:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprAShr:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprAnd:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprOr:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprXor:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprGetElementPtr:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.Elem, before, after)
		WalkBeforeAfter(&n.Src, before, after)
		WalkBeforeAfter(&n.Indices, before, after)
	case *ast.ExprTrunc:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.ExprZExt:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.ExprSExt:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.ExprFPTrunc:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.ExprFPExt:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.ExprFPToUI:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.ExprFPToSI:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.ExprUIToFP:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.ExprSIToFP:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.ExprPtrToInt:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.ExprIntToPtr:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.ExprBitCast:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.ExprAddrSpaceCast:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.ExprICmp:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprFCmp:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.ExprSelect:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.Cond, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case []*ast.BasicBlock:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.BasicBlock:
		WalkBeforeAfter(&n.Insts, before, after)
		WalkBeforeAfter(&n.Term, before, after)
	// Instructions
	case []ast.Instruction:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.InstAdd:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstFAdd:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstSub:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstFSub:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstMul:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstFMul:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstUDiv:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstSDiv:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstFDiv:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstURem:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstSRem:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstFRem:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstShl:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstLShr:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstAShr:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstAnd:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstOr:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstXor:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstAlloca:
		WalkBeforeAfter(&n.Elem, before, after)
		WalkBeforeAfter(&n.NElems, before, after)
	case *ast.InstLoad:
		WalkBeforeAfter(&n.Elem, before, after)
		WalkBeforeAfter(&n.Src, before, after)
	case *ast.InstStore:
		WalkBeforeAfter(&n.Src, before, after)
		WalkBeforeAfter(&n.Dst, before, after)
	case *ast.InstGetElementPtr:
		WalkBeforeAfter(&n.Elem, before, after)
		WalkBeforeAfter(&n.Src, before, after)
		WalkBeforeAfter(&n.Indices, before, after)
	case *ast.InstTrunc:
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.InstZExt:
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.InstSExt:
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.InstFPTrunc:
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.InstFPExt:
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.InstFPToUI:
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.InstFPToSI:
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.InstUIToFP:
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.InstSIToFP:
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.InstPtrToInt:
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.InstIntToPtr:
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.InstBitCast:
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.InstAddrSpaceCast:
		WalkBeforeAfter(&n.From, before, after)
		WalkBeforeAfter(&n.To, before, after)
	case *ast.InstICmp:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstFCmp:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstPhi:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.Incs, before, after)
	case []*ast.Incoming:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.Incoming:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Pred, before, after)
	case *ast.InstSelect:
		WalkBeforeAfter(&n.Cond, before, after)
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Y, before, after)
	case *ast.InstCall:
		WalkBeforeAfter(&n.Type, before, after)
		WalkBeforeAfter(&n.Callee, before, after)
		WalkBeforeAfter(&n.Args, before, after)
	// Terminators
	case *ast.TermRet:
		if n.X != nil {
			WalkBeforeAfter(&n.X, before, after)
		}
	case *ast.TermBr:
		WalkBeforeAfter(&n.Target, before, after)
	case *ast.TermCondBr:
		WalkBeforeAfter(&n.Cond, before, after)
		WalkBeforeAfter(&n.TargetTrue, before, after)
		WalkBeforeAfter(&n.TargetFalse, before, after)
	case *ast.TermSwitch:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.TargetDefault, before, after)
		WalkBeforeAfter(&n.Cases, before, after)
	case []*ast.Case:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.Case:
		WalkBeforeAfter(&n.X, before, after)
		WalkBeforeAfter(&n.Target, before, after)
	case *ast.TermUnreachable:
		// nothing to do.

	default:
		panic(fmt.Errorf("support for type %T not yet implemented", x))
	}

	after(x)
}
