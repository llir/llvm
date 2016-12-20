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
		WalkBeforeAfter(&n.Init, before, after)
	case []*ast.Function:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.Function:
		WalkBeforeAfter(&n.Blocks, before, after)
		WalkBeforeAfter(&n.Sig, before, after)
	case []*ast.Param:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.Param:
		panic("not yet implemented")
	case *ast.GlobalDummy:
		panic("not yet implemented")
	case *ast.LocalDummy:
		panic("not yet implemented")
	// Types
	case []ast.Type:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.VoidType:
		panic("not yet implemented")
	case *ast.LabelType:
		panic("not yet implemented")
	case *ast.IntType:
		panic("not yet implemented")
	case *ast.FloatType:
		panic("not yet implemented")
	case *ast.FuncType:
		panic("not yet implemented")
	case *ast.PointerType:
		panic("not yet implemented")
	case *ast.VectorType:
		panic("not yet implemented")
	case *ast.ArrayType:
		panic("not yet implemented")
	case *ast.StructType:
		panic("not yet implemented")
	case []*ast.NamedType:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.NamedType:
		panic("not yet implemented")
	case *ast.NamedTypeDummy:
		panic("not yet implemented")
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
		panic("not yet implemented")
	case *ast.FloatConst:
		panic("not yet implemented")
	case *ast.NullConst:
		panic("not yet implemented")
	case *ast.VectorConst:
		panic("not yet implemented")
	case *ast.ArrayConst:
		panic("not yet implemented")
	case *ast.StructConst:
		panic("not yet implemented")
	case *ast.ZeroInitializerConst:
		panic("not yet implemented")
	// Constant expressions
	case *ast.ExprAdd:
		panic("not yet implemented")
	case *ast.ExprFAdd:
		panic("not yet implemented")
	case *ast.ExprSub:
		panic("not yet implemented")
	case *ast.ExprFSub:
		panic("not yet implemented")
	case *ast.ExprMul:
		panic("not yet implemented")
	case *ast.ExprFMul:
		panic("not yet implemented")
	case *ast.ExprUDiv:
		panic("not yet implemented")
	case *ast.ExprSDiv:
		panic("not yet implemented")
	case *ast.ExprFDiv:
		panic("not yet implemented")
	case *ast.ExprURem:
		panic("not yet implemented")
	case *ast.ExprSRem:
		panic("not yet implemented")
	case *ast.ExprFRem:
		panic("not yet implemented")
	case *ast.ExprShl:
		panic("not yet implemented")
	case *ast.ExprLShr:
		panic("not yet implemented")
	case *ast.ExprAShr:
		panic("not yet implemented")
	case *ast.ExprAnd:
		panic("not yet implemented")
	case *ast.ExprOr:
		panic("not yet implemented")
	case *ast.ExprXor:
		panic("not yet implemented")
	case *ast.ExprGetElementPtr:
		panic("not yet implemented")
	case *ast.ExprTrunc:
		panic("not yet implemented")
	case *ast.ExprZExt:
		panic("not yet implemented")
	case *ast.ExprSExt:
		panic("not yet implemented")
	case *ast.ExprFPTrunc:
		panic("not yet implemented")
	case *ast.ExprFPExt:
		panic("not yet implemented")
	case *ast.ExprFPToUI:
		panic("not yet implemented")
	case *ast.ExprFPToSI:
		panic("not yet implemented")
	case *ast.ExprUIToFP:
		panic("not yet implemented")
	case *ast.ExprSIToFP:
		panic("not yet implemented")
	case *ast.ExprPtrToInt:
		panic("not yet implemented")
	case *ast.ExprIntToPtr:
		panic("not yet implemented")
	case *ast.ExprBitCast:
		panic("not yet implemented")
	case *ast.ExprAddrSpaceCast:
		panic("not yet implemented")
	case *ast.ExprICmp:
		panic("not yet implemented")
	case *ast.ExprFCmp:
		panic("not yet implemented")
	case *ast.ExprSelect:
		panic("not yet implemented")
	case []*ast.BasicBlock:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.BasicBlock:
		panic("not yet implemented")
	// Instructions
	case []ast.Instruction:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.InstAdd:
		panic("not yet implemented")
	case *ast.InstFAdd:
		panic("not yet implemented")
	case *ast.InstSub:
		panic("not yet implemented")
	case *ast.InstFSub:
		panic("not yet implemented")
	case *ast.InstMul:
		panic("not yet implemented")
	case *ast.InstFMul:
		panic("not yet implemented")
	case *ast.InstUDiv:
		panic("not yet implemented")
	case *ast.InstSDiv:
		panic("not yet implemented")
	case *ast.InstFDiv:
		panic("not yet implemented")
	case *ast.InstURem:
		panic("not yet implemented")
	case *ast.InstSRem:
		panic("not yet implemented")
	case *ast.InstFRem:
		panic("not yet implemented")
	case *ast.InstShl:
		panic("not yet implemented")
	case *ast.InstLShr:
		panic("not yet implemented")
	case *ast.InstAShr:
		panic("not yet implemented")
	case *ast.InstAnd:
		panic("not yet implemented")
	case *ast.InstOr:
		panic("not yet implemented")
	case *ast.InstXor:
		panic("not yet implemented")
	case *ast.InstAlloca:
		panic("not yet implemented")
	case *ast.InstLoad:
		panic("not yet implemented")
	case *ast.InstStore:
		panic("not yet implemented")
	case *ast.InstGetElementPtr:
		panic("not yet implemented")
	case *ast.InstTrunc:
		panic("not yet implemented")
	case *ast.InstZExt:
		panic("not yet implemented")
	case *ast.InstSExt:
		panic("not yet implemented")
	case *ast.InstFPTrunc:
		panic("not yet implemented")
	case *ast.InstFPExt:
		panic("not yet implemented")
	case *ast.InstFPToUI:
		panic("not yet implemented")
	case *ast.InstFPToSI:
		panic("not yet implemented")
	case *ast.InstUIToFP:
		panic("not yet implemented")
	case *ast.InstSIToFP:
		panic("not yet implemented")
	case *ast.InstPtrToInt:
		panic("not yet implemented")
	case *ast.InstIntToPtr:
		panic("not yet implemented")
	case *ast.InstBitCast:
		panic("not yet implemented")
	case *ast.InstAddrSpaceCast:
		panic("not yet implemented")
	case *ast.InstICmp:
		panic("not yet implemented")
	case *ast.InstFCmp:
		panic("not yet implemented")
	case *ast.InstPhi:
		panic("not yet implemented")
	case []*ast.Incoming:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.Incoming:
		panic("not yet implemented")
	case *ast.InstSelect:
		panic("not yet implemented")
	case *ast.InstCall:
		panic("not yet implemented")
	// Terminators
	case *ast.TermRet:
		panic("not yet implemented")
	case *ast.TermBr:
		panic("not yet implemented")
	case *ast.TermCondBr:
		panic("not yet implemented")
	case *ast.TermSwitch:
		panic("not yet implemented")
	case []*ast.Case:
		for i := range n {
			WalkBeforeAfter(&n[i], before, after)
		}
	case *ast.Case:
		panic("not yet implemented")
	case *ast.TermUnreachable:
		panic("not yet implemented")

	default:
		panic(fmt.Errorf("support for type %T not yet implemented", x))
	}

	after(x)
}
