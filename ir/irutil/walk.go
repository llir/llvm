// Package irutil provides LLVM IR utility functions.
package irutil

import (
	"fmt"

	"github.com/llir/llvm/internal/dummy"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Walk traverses the given LLVM IR module in depth-first order, calling visit
// for each type, value, instruction and terminator of the module.
func Walk(m *ir.Module, visit func(node interface{})) {
	w := newWalker()
	w.walk(m, visit)
}

// WalkType traverses the given LLVM IR type in depth-first order, calling visit
// for each type and value of the type.
func WalkType(t types.Type, visit func(node interface{})) {
	w := newWalker()
	w.walkType(t, visit)
}

// WalkValue traverses the given LLVM IR value in depth-first order, calling
// visit for each type, value, instruction and terminator of the value.
func WalkValue(v value.Value, visit func(node interface{})) {
	w := newWalker()
	w.walkValue(v, visit)
}

// WalkConstant traverses the given LLVM IR constant in depth-first order,
// calling visit for each type and constant of the constant.
func WalkConstant(c constant.Constant, visit func(node interface{})) {
	w := newWalker()
	w.walkConstant(c, visit)
}

// WalkExpr traverses the given LLVM IR constant expression in depth-first
// order, calling visit for each type and constant of the constant expression.
func WalkExpr(expr constant.Expr, visit func(node interface{})) {
	w := newWalker()
	w.walkExpr(expr, visit)
}

// WalkInst traverses the given LLVM IR instruction in depth-first order,
// calling visit for each type, value and instruction of the instruction.
func WalkInst(inst ir.Instruction, visit func(node interface{})) {
	w := newWalker()
	w.walkInst(inst, visit)
}

// WalkTerm traverses the given LLVM IR terminator in depth-first order, calling
// visit for each type, value and instruction of the terminator.
func WalkTerm(term ir.Terminator, visit func(node interface{})) {
	w := newWalker()
	w.walkTerm(term, visit)
}

// A walker traverses a given LLVM IR module in depth-first order.
type walker struct {
	// visited keeps track of visited nodes to prevent infinite loops.
	visited map[interface{}]bool
}

// newWalker returns a new walker for travering LLVM IR modules in depth-first
// order.
func newWalker() *walker {
	return &walker{
		visited: make(map[interface{}]bool),
	}
}

// walk traverses the given LLVM IR module in depth-first order, calling visit
// for each type, value, instruction and terminator of the module.
func (w *walker) walk(m *ir.Module, visit func(node interface{})) {
	for _, typ := range m.Types() {
		w.walkType(typ, visit)
	}
	for _, global := range m.Globals() {
		w.walkValue(global, visit)
	}
	for _, f := range m.Funcs() {
		w.walkValue(f, visit)
	}
}

// walkType traverses the given LLVM IR type in depth-first order, calling visit
// for each type and value of the type.
func (w *walker) walkType(t types.Type, visit func(node interface{})) {
	// Prevent infinite loops.
	if w.visited[t] {
		return
	}
	w.visited[t] = true

	visit(t)
	switch t := t.(type) {
	// Basic types.
	case *types.VoidType:
		// nothing to do; no child nodes.
	case *types.LabelType:
		// nothing to do; no child nodes.
	case *types.IntType:
		// nothing to do; no child nodes.
	case *types.FloatType:
		// nothing to do; no child nodes.

	// Derived types.
	case *types.FuncType:
		w.walkType(t.RetType(), visit)
		for _, param := range t.Params() {
			w.walkValue(param, visit)
		}
	case *types.PointerType:
		w.walkType(t.Elem(), visit)
	case *types.VectorType:
		w.walkType(t.Elem(), visit)
	case *types.ArrayType:
		w.walkType(t.Elem(), visit)
	case *types.StructType:
		for _, field := range t.Fields() {
			w.walkType(field, visit)
		}

	// Named types.
	case *types.NamedType:
		if def, ok := t.Def(); ok {
			w.walkType(def, visit)
		}
	default:
		panic(fmt.Sprintf("support for walking type %T not yet implemented", t))
	}
}

// walkValue traverses the given LLVM IR value in depth-first order, calling
// visit for each type, value, instruction and terminator of the value.
func (w *walker) walkValue(v value.Value, visit func(node interface{})) {
	// Prevent infinite loops.
	if w.visited[v] {
		return
	}
	w.visited[v] = true

	visit(v)
	switch v := v.(type) {
	case constant.Constant:
		// Let walkConstant handle the subset of constant values. Mark visited as
		// false for now.
		w.visited[v] = false
		w.walkConstant(v, visit)
	case *types.Param:
		w.walkType(v.Type(), visit)
	case *ir.BasicBlock:
		for _, inst := range v.Insts() {
			w.walkInst(inst, visit)
		}
		w.walkTerm(v.Term(), visit)
	case ir.Instruction:
		// Let walkInst handle the subset of instruction values. Mark visited as
		// false for now.
		w.visited[v] = false
		w.walkInst(v, visit)
	// Dummy values.
	case *dummy.Local:
		w.walkType(v.Type(), visit)
	default:
		panic(fmt.Sprintf("support for walking value %T not yet implemented", v))
	}
}

// walkConstant traverses the given LLVM IR constant in depth-first order,
// calling visit for each type and constant of the constant.
func (w *walker) walkConstant(c constant.Constant, visit func(node interface{})) {
	// Prevent infinite loops.
	if w.visited[c] {
		return
	}
	w.visited[c] = true

	visit(c)
	switch c := c.(type) {
	// Simple constants
	case *constant.Int:
		w.walkType(c.Type(), visit)
	case *constant.Float:
		w.walkType(c.Type(), visit)
	case *constant.Null:
		w.walkType(c.Type(), visit)

	// Complex constants
	case *constant.Vector:
		w.walkType(c.Type(), visit)
		for _, elem := range c.Elems() {
			w.walkConstant(elem, visit)
		}
	case *constant.Array:
		w.walkType(c.Type(), visit)
		for _, elem := range c.Elems() {
			w.walkConstant(elem, visit)
		}
	case *constant.Struct:
		w.walkType(c.Type(), visit)
		for _, field := range c.Fields() {
			w.walkConstant(field, visit)
		}
	case *constant.ZeroInitializer:
		w.walkType(c.Type(), visit)

	// Global variable and function addresses
	case *ir.Global:
		w.walkType(c.Type(), visit)
		w.walkType(c.ContentType(), visit)
		if init, ok := c.Init(); ok {
			w.walkConstant(init, visit)
		}
	case *ir.Function:
		w.walkType(c.Type(), visit)
		w.walkType(c.Sig(), visit)
		for _, block := range c.Blocks() {
			w.walkValue(block, visit)
		}

	// Constant expressions
	case constant.Expr:
		// Let walkExpr handle the subset of constant expressions. Mark visited as
		// false for now.
		w.visited[c] = false
		w.walkExpr(c, visit)

	// Dummy constants.
	case *dummy.Global:
		w.walkType(c.Type(), visit)
	default:
		panic(fmt.Sprintf("support for walking constant %T not yet implemented", c))
	}
}

// walkExpr traverses the given LLVM IR constant expression in depth-first
// order, calling visit for each type and constant of the constant expression.
func (w *walker) walkExpr(expr constant.Expr, visit func(node interface{})) {
	// Prevent infinite loops.
	if w.visited[expr] {
		return
	}
	w.visited[expr] = true

	visit(expr)
	switch expr := expr.(type) {
	// Binary instructions
	case *constant.ExprAdd:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprFAdd:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprSub:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprFSub:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprMul:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprFMul:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprUDiv:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprSDiv:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprFDiv:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprURem:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprSRem:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprFRem:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)

	// Bitwise instructions
	case *constant.ExprShl:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprLShr:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprAShr:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprAnd:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprOr:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprXor:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)

	// Memory instructions
	case *constant.ExprGetElementPtr:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.Src(), visit)
		for _, index := range expr.Indices() {
			w.walkConstant(index, visit)
		}

	// Conversion instructions
	case *constant.ExprTrunc:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.From(), visit)
	case *constant.ExprZExt:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.From(), visit)
	case *constant.ExprSExt:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.From(), visit)
	case *constant.ExprFPTrunc:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.From(), visit)
	case *constant.ExprFPExt:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.From(), visit)
	case *constant.ExprFPToUI:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.From(), visit)
	case *constant.ExprFPToSI:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.From(), visit)
	case *constant.ExprUIToFP:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.From(), visit)
	case *constant.ExprSIToFP:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.From(), visit)
	case *constant.ExprPtrToInt:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.From(), visit)
	case *constant.ExprIntToPtr:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.From(), visit)
	case *constant.ExprBitCast:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.From(), visit)
	case *constant.ExprAddrSpaceCast:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.From(), visit)

	// Other instructions
	case *constant.ExprICmp:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprFCmp:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	case *constant.ExprSelect:
		w.walkType(expr.Type(), visit)
		w.walkConstant(expr.Cond(), visit)
		w.walkConstant(expr.X(), visit)
		w.walkConstant(expr.Y(), visit)
	default:
		panic(fmt.Sprintf("support for walking constant expression %T not yet implemented", expr))
	}
}

// walkInst traverses the given LLVM IR instruction in depth-first order,
// calling visit for each type, value and instruction of the instruction.
func (w *walker) walkInst(inst ir.Instruction, visit func(node interface{})) {
	// Prevent infinite loops.
	if w.visited[inst] {
		return
	}
	w.visited[inst] = true

	visit(inst)
	switch inst := inst.(type) {
	// Binary instructions
	case *ir.InstAdd:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstFAdd:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstSub:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstFSub:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstMul:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstFMul:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstUDiv:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstSDiv:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstFDiv:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstURem:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstSRem:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstFRem:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)

	// Bitwise instructions
	case *ir.InstShl:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstLShr:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstAShr:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstAnd:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstOr:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstXor:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)

	// Memory instructions
	case *ir.InstAlloca:
		w.walkType(inst.Type(), visit)
		w.walkType(inst.ElemType(), visit)
		if nelems, ok := inst.NElems(); ok {
			w.walkValue(nelems, visit)
		}
	case *ir.InstLoad:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.Src(), visit)
	case *ir.InstStore:
		w.walkValue(inst.Src(), visit)
		w.walkValue(inst.Dst(), visit)
	case *ir.InstGetElementPtr:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.Src(), visit)
		for _, index := range inst.Indices() {
			w.walkValue(index, visit)
		}

	// Conversion instructions
	case *ir.InstTrunc:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.From(), visit)
	case *ir.InstZExt:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.From(), visit)
	case *ir.InstSExt:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.From(), visit)
	case *ir.InstFPTrunc:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.From(), visit)
	case *ir.InstFPExt:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.From(), visit)
	case *ir.InstFPToUI:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.From(), visit)
	case *ir.InstFPToSI:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.From(), visit)
	case *ir.InstUIToFP:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.From(), visit)
	case *ir.InstSIToFP:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.From(), visit)
	case *ir.InstPtrToInt:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.From(), visit)
	case *ir.InstIntToPtr:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.From(), visit)
	case *ir.InstBitCast:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.From(), visit)
	case *ir.InstAddrSpaceCast:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.From(), visit)

	// Other instructions
	case *ir.InstICmp:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstFCmp:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstPhi:
		w.walkType(inst.Type(), visit)
		for _, inc := range inst.Incs() {
			w.walkValue(inc.X(), visit)
			w.walkValue(inc.Pred(), visit)
		}
	case *ir.InstSelect:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.Cond(), visit)
		w.walkValue(inst.X(), visit)
		w.walkValue(inst.Y(), visit)
	case *ir.InstCall:
		w.walkType(inst.Type(), visit)
		w.walkValue(inst.Callee(), visit)
		for _, arg := range inst.Args() {
			w.walkValue(arg, visit)
		}

	// Dummy instructions
	case *dummy.InstGetElementPtr:
		w.walkType(inst.ElemType(), visit)
		w.walkValue(inst.Src(), visit)
		for _, index := range inst.Indices() {
			w.walkValue(index, visit)
		}
	case *dummy.InstPhi:
		w.walkType(inst.Type(), visit)
		for _, inc := range inst.Incs() {
			x, ok := inc.X().(value.Value)
			if !ok {
				panic(fmt.Sprintf("invalid x type, expected value.Value, got %T", inc.X()))
			}
			w.walkValue(x, visit)
		}
	case *dummy.InstCall:
		w.walkType(inst.Type(), visit)
		for _, arg := range inst.Args() {
			w.walkValue(arg, visit)
		}
	default:
		panic(fmt.Sprintf("support for walking instruction %T not yet implemented", inst))
	}
}

// walkTerm traverses the given LLVM IR terminator in depth-first order, calling
// visit for each type, value and instruction of the terminator.
func (w *walker) walkTerm(term ir.Terminator, visit func(node interface{})) {
	// Prevent infinite loops.
	if w.visited[term] {
		return
	}
	w.visited[term] = true

	visit(term)
	switch term := term.(type) {
	// Terminators.
	case *ir.TermRet:
		if x, ok := term.X(); ok {
			w.walkValue(x, visit)
		}
	case *ir.TermBr:
		w.walkValue(term.Target(), visit)
	case *ir.TermCondBr:
		w.walkValue(term.Cond(), visit)
		w.walkValue(term.TargetTrue(), visit)
		w.walkValue(term.TargetFalse(), visit)
	case *ir.TermSwitch:
		w.walkValue(term.X(), visit)
		w.walkValue(term.TargetDefault(), visit)
		for _, c := range term.Cases() {
			w.walkValue(c.X(), visit)
			w.walkValue(c.Target(), visit)
		}
	case *ir.TermUnreachable:
		// nothing to do; no child nodes.

	// Dummy terminators
	case *dummy.TermBr:
		// nothing to do; no child nodes.
	case *dummy.TermCondBr:
		w.walkValue(term.Cond(), visit)
	case *dummy.TermSwitch:
		w.walkValue(term.X(), visit)
		for _, c := range term.Cases() {
			w.walkValue(c.X(), visit)
		}
	default:
		panic(fmt.Sprintf("support for walking terminator %T not yet implemented", term))
	}
}
