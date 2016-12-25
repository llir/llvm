// Note, the LLVM IR traversal implementation of this package is heavily
// inspired by go fix; which is governed by a BSD license.

// Package irutil provides LLVM IR utility functions.
package irutil

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Walk traverses the AST x, calling visit(y) for each node y in the tree but
// also with a pointer to each types.Type, value.Value, and *ir.BasicBlock, in a
// bottom-up traversal.
func Walk(x interface{}, visit func(interface{})) {
	WalkBeforeAfter(x, nop, visit)
}

// WalkFunc traverses the AST of the given function, calling visit(y) for each
// node y in the tree but also with a pointer to each types.Type, value.Value,
// and *ir.BasicBlock, in a bottom-up traversal.
func WalkFunc(f *ir.Function, visit func(interface{})) {
	WalkFuncBeforeAfter(f, nop, visit)
}

// nop performs no operation on the given AST.
func nop(x interface{}) {
}

// WalkBeforeAfter traverses the AST x, calling before(y) before traversing y's
// children and after(y) afterward for each node y in the tree but also with a
// pointer to each types.Type, value.Value, and *ir.BasicBlock, in a bottom-up
// traversal.
func WalkBeforeAfter(x interface{}, before, after func(interface{})) {
	w := &walker{
		visited: make(map[interface{}]bool),
	}
	w.walkBeforeAfter(x, before, after)
}

// WalkFuncBeforeAfter traverses the AST of the given function, calling
// before(y) before traversing y's children and after(y) afterward for each node
// y in the tree but also with a pointer to each types.Type, value.Value, and
// *ir.BasicBlock, in a bottom-up traversal.
//
// Special precausion is taken during traversal to stay within the scope of the
// function.
func WalkFuncBeforeAfter(f *ir.Function, before, after func(interface{})) {
	w := &walker{
		funcScope: true,
		visited:   make(map[interface{}]bool),
	}
	// Traverse child nodes of function, instead of f directly, as *ir.Function
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
// pointer to each types.Type, value.Value, and *ir.BasicBlock, in a bottom-up
// traversal.
func (w *walker) walkBeforeAfter(x interface{}, before, after func(interface{})) {
	switch x.(type) {
	case []*ir.Global, []*ir.Function, []*ir.Param, []types.Type, []*types.NamedType, []value.Value, []constant.Constant, []*ir.BasicBlock, []ir.Instruction, []*ir.Incoming, []*ir.Case:
		// unhashable type.
	case *ir.Function:
		if w.funcScope {
			// *ir.Function nodes are not traversed when staying within the scope
			// *of the function.
			return
		}
	default:
		// Prevent infinite loops.

		// TODO: Check if it is enough to only track *types.NamedType to prevent
		// inf loops.
		if w.visited[x] {
			return
		}
		w.visited[x] = true
	}

	before(x)

	switch n := x.(type) {
	// pointers to interfaces
	case *types.Type:
		w.walkBeforeAfter(*n, before, after)
	case *value.Value:
		w.walkBeforeAfter(*n, before, after)
	case *value.Named:
		w.walkBeforeAfter(*n, before, after)
	case *constant.Constant:
		w.walkBeforeAfter(*n, before, after)
	case *ir.Instruction:
		w.walkBeforeAfter(*n, before, after)
	case *ir.Terminator:
		w.walkBeforeAfter(*n, before, after)

	// pointers to struct pointers
	case **ir.Global:
		w.walkBeforeAfter(*n, before, after)
	case **ir.Function:
		w.walkBeforeAfter(*n, before, after)
	case **ir.Param:
		w.walkBeforeAfter(*n, before, after)
	// Types
	case **types.VoidType:
		w.walkBeforeAfter(*n, before, after)
	case **types.LabelType:
		w.walkBeforeAfter(*n, before, after)
	case **types.IntType:
		w.walkBeforeAfter(*n, before, after)
	case **types.FloatType:
		w.walkBeforeAfter(*n, before, after)
	case **types.FuncType:
		w.walkBeforeAfter(*n, before, after)
	case **types.PointerType:
		w.walkBeforeAfter(*n, before, after)
	case **types.VectorType:
		w.walkBeforeAfter(*n, before, after)
	case **types.ArrayType:
		w.walkBeforeAfter(*n, before, after)
	case **types.StructType:
		w.walkBeforeAfter(*n, before, after)
	case **types.NamedType:
		w.walkBeforeAfter(*n, before, after)
	// Constants
	case **constant.Int:
		w.walkBeforeAfter(*n, before, after)
	case **constant.Float:
		w.walkBeforeAfter(*n, before, after)
	case **constant.Null:
		w.walkBeforeAfter(*n, before, after)
	case **constant.Vector:
		w.walkBeforeAfter(*n, before, after)
	case **constant.Array:
		w.walkBeforeAfter(*n, before, after)
	case **constant.Struct:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ZeroInitializer:
		w.walkBeforeAfter(*n, before, after)
	// Constant expressions
	case **constant.ExprAdd:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprFAdd:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprSub:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprFSub:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprMul:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprFMul:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprUDiv:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprSDiv:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprFDiv:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprURem:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprSRem:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprFRem:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprShl:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprLShr:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprAShr:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprAnd:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprOr:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprXor:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprGetElementPtr:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprTrunc:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprZExt:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprSExt:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprFPTrunc:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprFPExt:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprFPToUI:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprFPToSI:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprUIToFP:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprSIToFP:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprPtrToInt:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprIntToPtr:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprBitCast:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprAddrSpaceCast:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprICmp:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprFCmp:
		w.walkBeforeAfter(*n, before, after)
	case **constant.ExprSelect:
		w.walkBeforeAfter(*n, before, after)
	// Basic blocks.
	case **ir.BasicBlock:
		w.walkBeforeAfter(*n, before, after)
	// Instructions
	case **ir.InstAdd:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstFAdd:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstSub:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstFSub:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstMul:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstFMul:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstUDiv:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstSDiv:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstFDiv:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstURem:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstSRem:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstFRem:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstShl:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstLShr:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstAShr:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstAnd:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstOr:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstXor:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstAlloca:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstLoad:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstStore:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstGetElementPtr:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstTrunc:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstZExt:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstSExt:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstFPTrunc:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstFPExt:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstFPToUI:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstFPToSI:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstUIToFP:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstSIToFP:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstPtrToInt:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstIntToPtr:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstBitCast:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstAddrSpaceCast:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstICmp:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstFCmp:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstPhi:
		w.walkBeforeAfter(*n, before, after)
	case **ir.Incoming:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstSelect:
		w.walkBeforeAfter(*n, before, after)
	case **ir.InstCall:
		w.walkBeforeAfter(*n, before, after)
	// Terminators
	case **ir.TermRet:
		w.walkBeforeAfter(*n, before, after)
	case **ir.TermBr:
		w.walkBeforeAfter(*n, before, after)
	case **ir.TermCondBr:
		w.walkBeforeAfter(*n, before, after)
	case **ir.TermSwitch:
		w.walkBeforeAfter(*n, before, after)
	case **ir.Case:
		w.walkBeforeAfter(*n, before, after)
	case **ir.TermUnreachable:
		w.walkBeforeAfter(*n, before, after)

	// pointers to slices
	case *[]types.Type:
		w.walkBeforeAfter(*n, before, after)
	case *[]*types.NamedType:
		w.walkBeforeAfter(*n, before, after)
	case *[]*ir.Global:
		w.walkBeforeAfter(*n, before, after)
	case *[]value.Value:
		w.walkBeforeAfter(*n, before, after)
	case *[]constant.Constant:
		w.walkBeforeAfter(*n, before, after)
	case *[]*ir.Function:
		w.walkBeforeAfter(*n, before, after)
	case *[]*ir.Param:
		w.walkBeforeAfter(*n, before, after)
	case *[]*ir.BasicBlock:
		w.walkBeforeAfter(*n, before, after)
	case *[]ir.Instruction:
		w.walkBeforeAfter(*n, before, after)
	case *[]*ir.Case:
		w.walkBeforeAfter(*n, before, after)
	case *[]*ir.Incoming:
		w.walkBeforeAfter(*n, before, after)

	// These are ordered and grouped to match ../../ll.bnf
	case *ir.Module:
		if n.Types != nil {
			w.walkBeforeAfter(&n.Types, before, after)
		}
		if n.Globals != nil {
			w.walkBeforeAfter(&n.Globals, before, after)
		}
		if n.Funcs != nil {
			w.walkBeforeAfter(&n.Funcs, before, after)
		}
	case []*ir.Global:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ir.Global:
		w.walkBeforeAfter(&n.Content, before, after)
		if n.Init != nil {
			w.walkBeforeAfter(&n.Init, before, after)
		}
	case []*ir.Function:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ir.Function:
		w.walkBeforeAfter(&n.Sig, before, after)
		if n.Blocks != nil {
			w.walkBeforeAfter(&n.Blocks, before, after)
		}
	case []*ir.Param:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ir.Param:
		w.walkBeforeAfter(&n.Typ, before, after)
	// Types
	case []types.Type:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *types.VoidType:
		// nothing to do.
	case *types.LabelType:
		// nothing to do.
	case *types.IntType:
		// nothing to do.
	case *types.FloatType:
		// nothing to do.
	case *types.FuncType:
		w.walkBeforeAfter(&n.Ret, before, after)
		if n.Params != nil {
			w.walkBeforeAfter(&n.Params, before, after)
		}
	case *types.PointerType:
		w.walkBeforeAfter(&n.Elem, before, after)
	case *types.VectorType:
		w.walkBeforeAfter(&n.Elem, before, after)
	case *types.ArrayType:
		w.walkBeforeAfter(&n.Elem, before, after)
	case *types.StructType:
		if n.Fields != nil {
			w.walkBeforeAfter(&n.Fields, before, after)
		}
	case []*types.NamedType:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *types.NamedType:
		if n.Def != nil {
			w.walkBeforeAfter(&n.Def, before, after)
		}
	// Constants
	case []value.Value:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case []constant.Constant:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *constant.Int:
		w.walkBeforeAfter(&n.Typ, before, after)
	case *constant.Float:
		w.walkBeforeAfter(&n.Typ, before, after)
	case *constant.Null:
		w.walkBeforeAfter(&n.Typ, before, after)
	case *constant.Vector:
		w.walkBeforeAfter(&n.Typ, before, after)
		if n.Elems != nil {
			w.walkBeforeAfter(&n.Elems, before, after)
		}
	case *constant.Array:
		w.walkBeforeAfter(&n.Typ, before, after)
		if n.Elems != nil {
			w.walkBeforeAfter(&n.Elems, before, after)
		}
	case *constant.Struct:
		w.walkBeforeAfter(&n.Typ, before, after)
		if n.Fields != nil {
			w.walkBeforeAfter(&n.Fields, before, after)
		}
	case *constant.ZeroInitializer:
		w.walkBeforeAfter(&n.Typ, before, after)
	// Constant expressions
	case *constant.ExprAdd:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprFAdd:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprSub:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprFSub:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprMul:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprFMul:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprUDiv:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprSDiv:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprFDiv:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprURem:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprSRem:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprFRem:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprShl:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprLShr:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprAShr:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprAnd:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprOr:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprXor:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprGetElementPtr:
		w.walkBeforeAfter(&n.Typ, before, after)
		w.walkBeforeAfter(&n.Elem, before, after)
		w.walkBeforeAfter(&n.Src, before, after)
		if n.Indices != nil {
			w.walkBeforeAfter(&n.Indices, before, after)
		}
	case *constant.ExprTrunc:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *constant.ExprZExt:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *constant.ExprSExt:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *constant.ExprFPTrunc:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *constant.ExprFPExt:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *constant.ExprFPToUI:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *constant.ExprFPToSI:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *constant.ExprUIToFP:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *constant.ExprSIToFP:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *constant.ExprPtrToInt:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *constant.ExprIntToPtr:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *constant.ExprBitCast:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *constant.ExprAddrSpaceCast:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *constant.ExprICmp:
		w.walkBeforeAfter(&n.Typ, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprFCmp:
		w.walkBeforeAfter(&n.Typ, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *constant.ExprSelect:
		w.walkBeforeAfter(&n.Cond, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	// Basic blocks.
	case []*ir.BasicBlock:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ir.BasicBlock:
		if n.Insts != nil {
			w.walkBeforeAfter(&n.Insts, before, after)
		}
		w.walkBeforeAfter(&n.Term, before, after)
	// Instructions
	case []ir.Instruction:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ir.InstAdd:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstFAdd:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstSub:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstFSub:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstMul:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstFMul:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstUDiv:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstSDiv:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstFDiv:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstURem:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstSRem:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstFRem:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstShl:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstLShr:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstAShr:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstAnd:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstOr:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstXor:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstAlloca:
		w.walkBeforeAfter(&n.Elem, before, after)
		if n.NElems != nil {
			w.walkBeforeAfter(&n.NElems, before, after)
		}
	case *ir.InstLoad:
		w.walkBeforeAfter(&n.Typ, before, after)
		w.walkBeforeAfter(&n.Src, before, after)
	case *ir.InstStore:
		w.walkBeforeAfter(&n.Src, before, after)
		w.walkBeforeAfter(&n.Dst, before, after)
	case *ir.InstGetElementPtr:
		w.walkBeforeAfter(&n.Elem, before, after)
		w.walkBeforeAfter(&n.Src, before, after)
		if n.Indices != nil {
			w.walkBeforeAfter(&n.Indices, before, after)
		}
	case *ir.InstTrunc:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ir.InstZExt:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ir.InstSExt:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ir.InstFPTrunc:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ir.InstFPExt:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ir.InstFPToUI:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ir.InstFPToSI:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ir.InstUIToFP:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ir.InstSIToFP:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ir.InstPtrToInt:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ir.InstIntToPtr:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ir.InstBitCast:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ir.InstAddrSpaceCast:
		w.walkBeforeAfter(&n.From, before, after)
		w.walkBeforeAfter(&n.To, before, after)
	case *ir.InstICmp:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstFCmp:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstPhi:
		w.walkBeforeAfter(&n.Typ, before, after)
		if n.Incs != nil {
			w.walkBeforeAfter(&n.Incs, before, after)
		}
	case []*ir.Incoming:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ir.Incoming:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Pred, before, after)
	case *ir.InstSelect:
		w.walkBeforeAfter(&n.Cond, before, after)
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Y, before, after)
	case *ir.InstCall:
		w.walkBeforeAfter(&n.Callee, before, after)
		w.walkBeforeAfter(&n.Sig, before, after)
		if n.Args != nil {
			w.walkBeforeAfter(&n.Args, before, after)
		}
	// Terminators
	case *ir.TermRet:
		if n.X != nil {
			w.walkBeforeAfter(&n.X, before, after)
		}
	case *ir.TermBr:
		w.walkBeforeAfter(&n.Target, before, after)
	case *ir.TermCondBr:
		w.walkBeforeAfter(&n.Cond, before, after)
		w.walkBeforeAfter(&n.TargetTrue, before, after)
		w.walkBeforeAfter(&n.TargetFalse, before, after)
	case *ir.TermSwitch:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.TargetDefault, before, after)
		if n.Cases != nil {
			w.walkBeforeAfter(&n.Cases, before, after)
		}
	case []*ir.Case:
		for i := range n {
			w.walkBeforeAfter(&n[i], before, after)
		}
	case *ir.Case:
		w.walkBeforeAfter(&n.X, before, after)
		w.walkBeforeAfter(&n.Target, before, after)
	case *ir.TermUnreachable:
		// nothing to do.

	default:
		panic(fmt.Errorf("support for type %T not yet implemented", x))
	}

	after(x)
}
