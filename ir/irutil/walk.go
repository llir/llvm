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
	for _, typ := range m.Types() {
		WalkType(typ, visit)
	}
	for _, global := range m.Globals() {
		WalkValue(global, visit)
	}
	for _, f := range m.Funcs() {
		WalkValue(f, visit)
	}
}

// WalkType traverses the given LLVM IR type in depth-first order, calling visit
// for each type and value of the type.
func WalkType(t types.Type, visit func(node interface{})) {
	visit(t)
	switch t := t.(type) {
	case *types.VoidType:
		// nothing to do; no child nodes.
	case *types.LabelType:
		// nothing to do; no child nodes.
	case *types.IntType:
		// nothing to do; no child nodes.
	case *types.FloatType:
		// nothing to do; no child nodes.
	case *types.FuncType:
		WalkType(t.RetType(), visit)
		for _, param := range t.Params() {
			WalkValue(param, visit)
		}
	case *types.PointerType:
		WalkType(t.Elem(), visit)
	case *types.VectorType:
		WalkType(t.Elem(), visit)
	case *types.ArrayType:
		WalkType(t.Elem(), visit)
	case *types.StructType:
		for _, field := range t.Fields() {
			WalkType(field, visit)
		}
	case *types.NamedType:
		// TODO: Figure out how to prevent loops in recursive named type
		// defintions.
		if def, ok := t.Def(); ok {
			WalkType(def, visit)
		}
	default:
		panic(fmt.Sprintf("support for walking type %T not yet implemented", t))
	}
}

// WalkValue traverses the given LLVM IR value in depth-first order, calling
// visit for each type, value, instruction and terminator of the value.
func WalkValue(v value.Value, visit func(node interface{})) {
	visit(v)
	switch v := v.(type) {
	case constant.Constant:
		WalkConstant(v, visit)
	case *types.Param:
		WalkType(v.Type(), visit)
	case *ir.BasicBlock:
		for _, inst := range v.Insts() {
			WalkInst(inst, visit)
		}
		WalkTerm(v.Term(), visit)
	case ir.Instruction:
		WalkInst(v, visit)
	// Dummy values.
	case *dummy.Local:
		WalkType(v.Type(), visit)
	default:
		panic(fmt.Sprintf("support for walking value %T not yet implemented", v))
	}
}

// WalkConstant traverses the given LLVM IR constant in depth-first order,
// calling visit for each type and constant of the constant.
func WalkConstant(c constant.Constant, visit func(node interface{})) {
	visit(c)
	switch c := c.(type) {
	// Simple constants
	case *constant.Int:
		WalkType(c.Type(), visit)
	case *constant.Float:
		WalkType(c.Type(), visit)
	case *constant.Null:
		WalkType(c.Type(), visit)

	// Complex constants
	case *constant.Vector:
		WalkType(c.Type(), visit)
		for _, elem := range c.Elems() {
			WalkConstant(elem, visit)
		}
	case *constant.Array:
		WalkType(c.Type(), visit)
		for _, elem := range c.Elems() {
			WalkConstant(elem, visit)
		}
	case *constant.Struct:
		WalkType(c.Type(), visit)
		for _, field := range c.Fields() {
			WalkConstant(field, visit)
		}
	case *constant.ZeroInitializer:
		WalkType(c.Type(), visit)

	// Global variable and function addresses
	case *ir.Global:
		WalkType(c.Type(), visit)
		WalkType(c.ContentType(), visit)
		if init, ok := c.Init(); ok {
			WalkConstant(init, visit)
		}
	case *ir.Function:
		WalkType(c.Type(), visit)
		WalkType(c.Sig(), visit)
		for _, block := range c.Blocks() {
			WalkValue(block, visit)
		}

	// Constant expressions
	case constant.Expr:
		WalkExpr(c, visit)

	// Dummy constants.
	case *dummy.Global:
		WalkType(c.Type(), visit)
	default:
		panic(fmt.Sprintf("support for walking constant %T not yet implemented", c))
	}
}

// WalkExpr traverses the given LLVM IR constant expression in depth-first
// order, calling visit for each type and constant of the constant expression.
func WalkExpr(expr constant.Expr, visit func(node interface{})) {
	visit(expr)
	switch expr := expr.(type) {
	// Binary instructions
	case *constant.ExprAdd:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprFAdd:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprSub:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprFSub:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprMul:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprFMul:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprUDiv:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprSDiv:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprFDiv:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprURem:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprSRem:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprFRem:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)

	// Bitwise instructions
	case *constant.ExprShl:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprLShr:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprAShr:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprAnd:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprOr:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprXor:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)

	// Memory instructions
	case *constant.ExprGetElementPtr:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.Src(), visit)
		for _, index := range expr.Indices() {
			WalkConstant(index, visit)
		}

	// Conversion instructions
	case *constant.ExprTrunc:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.From(), visit)
	case *constant.ExprZExt:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.From(), visit)
	case *constant.ExprSExt:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.From(), visit)
	case *constant.ExprFPTrunc:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.From(), visit)
	case *constant.ExprFPExt:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.From(), visit)
	case *constant.ExprFPToUI:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.From(), visit)
	case *constant.ExprFPToSI:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.From(), visit)
	case *constant.ExprUIToFP:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.From(), visit)
	case *constant.ExprSIToFP:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.From(), visit)
	case *constant.ExprPtrToInt:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.From(), visit)
	case *constant.ExprIntToPtr:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.From(), visit)
	case *constant.ExprBitCast:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.From(), visit)
	case *constant.ExprAddrSpaceCast:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.From(), visit)

	// Other instructions
	case *constant.ExprICmp:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprFCmp:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	case *constant.ExprSelect:
		WalkType(expr.Type(), visit)
		WalkConstant(expr.Cond(), visit)
		WalkConstant(expr.X(), visit)
		WalkConstant(expr.Y(), visit)
	default:
		panic(fmt.Sprintf("support for walking constant expression %T not yet implemented", expr))
	}
}

// WalkInst traverses the given LLVM IR instruction in depth-first order,
// calling visit for each type, value and instruction of the instruction.
func WalkInst(inst ir.Instruction, visit func(node interface{})) {
	visit(inst)
	switch inst := inst.(type) {
	// Binary instructions
	case *ir.InstAdd:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstFAdd:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstSub:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstFSub:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstMul:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstFMul:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstUDiv:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstSDiv:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstFDiv:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstURem:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstSRem:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstFRem:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)

	// Bitwise instructions
	case *ir.InstShl:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstLShr:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstAShr:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstAnd:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstOr:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstXor:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)

	// Memory instructions
	case *ir.InstAlloca:
		WalkType(inst.Type(), visit)
		WalkType(inst.ElemType(), visit)
		if nelems, ok := inst.NElems(); ok {
			WalkValue(nelems, visit)
		}
	case *ir.InstLoad:
		WalkType(inst.Type(), visit)
		WalkValue(inst.Src(), visit)
	case *ir.InstStore:
		WalkValue(inst.Src(), visit)
		WalkValue(inst.Dst(), visit)
	case *ir.InstGetElementPtr:
		WalkType(inst.Type(), visit)
		WalkValue(inst.Src(), visit)
		for _, index := range inst.Indices() {
			WalkValue(index, visit)
		}

	// Conversion instructions
	case *ir.InstTrunc:
		WalkType(inst.Type(), visit)
		WalkValue(inst.From(), visit)
	case *ir.InstZExt:
		WalkType(inst.Type(), visit)
		WalkValue(inst.From(), visit)
	case *ir.InstSExt:
		WalkType(inst.Type(), visit)
		WalkValue(inst.From(), visit)
	case *ir.InstFPTrunc:
		WalkType(inst.Type(), visit)
		WalkValue(inst.From(), visit)
	case *ir.InstFPExt:
		WalkType(inst.Type(), visit)
		WalkValue(inst.From(), visit)
	case *ir.InstFPToUI:
		WalkType(inst.Type(), visit)
		WalkValue(inst.From(), visit)
	case *ir.InstFPToSI:
		WalkType(inst.Type(), visit)
		WalkValue(inst.From(), visit)
	case *ir.InstUIToFP:
		WalkType(inst.Type(), visit)
		WalkValue(inst.From(), visit)
	case *ir.InstSIToFP:
		WalkType(inst.Type(), visit)
		WalkValue(inst.From(), visit)
	case *ir.InstPtrToInt:
		WalkType(inst.Type(), visit)
		WalkValue(inst.From(), visit)
	case *ir.InstIntToPtr:
		WalkType(inst.Type(), visit)
		WalkValue(inst.From(), visit)
	case *ir.InstBitCast:
		WalkType(inst.Type(), visit)
		WalkValue(inst.From(), visit)
	case *ir.InstAddrSpaceCast:
		WalkType(inst.Type(), visit)
		WalkValue(inst.From(), visit)

	// Other instructions
	case *ir.InstICmp:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstFCmp:
		WalkType(inst.Type(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstPhi:
		WalkType(inst.Type(), visit)
		for _, inc := range inst.Incs() {
			WalkValue(inc.X(), visit)
			WalkValue(inc.Pred(), visit)
		}
	case *ir.InstSelect:
		WalkType(inst.Type(), visit)
		WalkValue(inst.Cond(), visit)
		WalkValue(inst.X(), visit)
		WalkValue(inst.Y(), visit)
	case *ir.InstCall:
		WalkType(inst.Type(), visit)
		WalkValue(inst.Callee(), visit)
		for _, arg := range inst.Args() {
			WalkValue(arg, visit)
		}

	// Dummy instructions
	case *dummy.InstCall:
		WalkType(inst.Type(), visit)
		for _, arg := range inst.Args() {
			WalkValue(arg, visit)
		}
	default:
		panic(fmt.Sprintf("support for walking instruction %T not yet implemented", inst))
	}
}

// WalkTerm traverses the given LLVM IR terminator in depth-first order, calling
// visit for each type, value and instruction of the terminator.
func WalkTerm(term ir.Terminator, visit func(node interface{})) {
	visit(term)
	switch term := term.(type) {
	case *ir.TermRet:
		if x, ok := term.X(); ok {
			WalkValue(x, visit)
		}
	case *ir.TermBr:
		WalkValue(term.Target(), visit)
	case *ir.TermCondBr:
		WalkValue(term.Cond(), visit)
		WalkValue(term.TargetTrue(), visit)
		WalkValue(term.TargetFalse(), visit)
	case *ir.TermSwitch:
		WalkValue(term.X(), visit)
		WalkValue(term.TargetDefault(), visit)
		for _, c := range term.Cases() {
			WalkValue(c.X(), visit)
			WalkValue(c.Target(), visit)
		}
	case *ir.TermUnreachable:
		// nothing to do; no child nodes.
	default:
		panic(fmt.Sprintf("support for walking terminator %T not yet implemented", term))
	}
}
