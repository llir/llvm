// Package sem implements a static semantic analysis checker of LLVM IR modules.
package sem

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/irutil"
	"github.com/llir/llvm/ir/types"
)

// Check performs static semantic analysis on the given LLVM IR module.
func Check(m *ir.Module) error {
	// List of identified errors.
	var errs []error
	// check performs static semantic analysis on the given LLVM IR node.
	check := func(n interface{}) {
		var err error
		switch n := n.(type) {
		case *ir.Global:
			err = checkGlobal(n)
		case *ir.Function:
			err = checkFunc(n)
		case *ir.Param:
			err = checkParam(n)
		case *ir.BasicBlock:
			err = checkBlock(n)
		case types.Type:
			err = checkType(n)
		case constant.Constant:
			err = checkConst(n)
		case ir.Instruction:
			err = checkInst(n)
		case ir.Terminator:
			err = checkTerm(n)
		}
		if err != nil {
			errs = append(errs, err)
		}
	}
	irutil.Walk(m, check)
	if len(errs) > 0 {
		// TODO: Return the full list of identified errors.
		return errs[0]
	}
	return nil
}

// checkGlobal validates the semantics of the given global.
func checkGlobal(global *ir.Global) error {
	panic("not yet implemented")
}

// checkFunc validates the semantics of the given function.
func checkFunc(f *ir.Function) error {
	panic("not yet implemented")
}

// checkParam validates the semantics of the given function parameter.
func checkParam(param *ir.Param) error {
	panic("not yet implemented")
}

// checkBlock validates the semantics of the given basic block.
func checkBlock(param *ir.BasicBlock) error {
	panic("not yet implemented")
}

// checkType validates the semantics of the given type.
func checkType(t types.Type) error {
	switch t := t.(type) {
	case *types.VoidType:
		panic("not yet implemented")
	case *types.LabelType:
		panic("not yet implemented")
	case *types.IntType:
		panic("not yet implemented")
	case *types.FloatType:
		panic("not yet implemented")
	case *types.FuncType:
		panic("not yet implemented")
	case *types.PointerType:
		panic("not yet implemented")
	case *types.VectorType:
		panic("not yet implemented")
	case *types.ArrayType:
		panic("not yet implemented")
	case *types.StructType:
		panic("not yet implemented")
	case *types.NamedType:
		panic("not yet implemented")
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", t))
	}
}

// checkConst validates the semantics of the given constant.
func checkConst(c constant.Constant) error {
	switch c := c.(type) {
	// Simple constants.
	case *constant.Int:
		panic("not yet implemented")
	case *constant.Float:
		panic("not yet implemented")
	case *constant.Null:
		panic("not yet implemented")
	// Complex constants.
	case *constant.Vector:
		panic("not yet implemented")
	case *constant.Array:
		panic("not yet implemented")
	case *constant.Struct:
		panic("not yet implemented")
	case *constant.ZeroInitializer:
		panic("not yet implemented")
	// Binary expressions.
	case *constant.ExprAdd:
		panic("not yet implemented")
	case *constant.ExprFAdd:
		panic("not yet implemented")
	case *constant.ExprSub:
		panic("not yet implemented")
	case *constant.ExprFSub:
		panic("not yet implemented")
	case *constant.ExprMul:
		panic("not yet implemented")
	case *constant.ExprFMul:
		panic("not yet implemented")
	case *constant.ExprUDiv:
		panic("not yet implemented")
	case *constant.ExprSDiv:
		panic("not yet implemented")
	case *constant.ExprFDiv:
		panic("not yet implemented")
	case *constant.ExprURem:
		panic("not yet implemented")
	case *constant.ExprSRem:
		panic("not yet implemented")
	case *constant.ExprFRem:
		panic("not yet implemented")
	// Bitwise expressions.
	case *constant.ExprShl:
		panic("not yet implemented")
	case *constant.ExprLShr:
		panic("not yet implemented")
	case *constant.ExprAShr:
		panic("not yet implemented")
	case *constant.ExprAnd:
		panic("not yet implemented")
	case *constant.ExprOr:
		panic("not yet implemented")
	case *constant.ExprXor:
		panic("not yet implemented")
	// Memory expressions.
	case *constant.ExprGetElementPtr:
		panic("not yet implemented")
	// Conversion expressions.
	case *constant.ExprTrunc:
		panic("not yet implemented")
	case *constant.ExprZExt:
		panic("not yet implemented")
	case *constant.ExprSExt:
		panic("not yet implemented")
	case *constant.ExprFPTrunc:
		panic("not yet implemented")
	case *constant.ExprFPExt:
		panic("not yet implemented")
	case *constant.ExprFPToUI:
		panic("not yet implemented")
	case *constant.ExprFPToSI:
		panic("not yet implemented")
	case *constant.ExprUIToFP:
		panic("not yet implemented")
	case *constant.ExprSIToFP:
		panic("not yet implemented")
	case *constant.ExprPtrToInt:
		panic("not yet implemented")
	case *constant.ExprIntToPtr:
		panic("not yet implemented")
	case *constant.ExprBitCast:
		panic("not yet implemented")
	case *constant.ExprAddrSpaceCast:
		panic("not yet implemented")
	// Other expressions.
	case *constant.ExprICmp:
		panic("not yet implemented")
	case *constant.ExprFCmp:
		panic("not yet implemented")
	case *constant.ExprSelect:
		panic("not yet implemented")
	default:
		panic(fmt.Errorf("support for constant %T not yet implemented", c))
	}
}

// checkInst validates the semantics of the given instruction.
func checkInst(inst ir.Instruction) error {
	switch inst := inst.(type) {
	// Binary instructions.
	case *ir.InstAdd:
		panic("not yet implemented")
	case *ir.InstFAdd:
		panic("not yet implemented")
	case *ir.InstSub:
		panic("not yet implemented")
	case *ir.InstFSub:
		panic("not yet implemented")
	case *ir.InstMul:
		panic("not yet implemented")
	case *ir.InstFMul:
		panic("not yet implemented")
	case *ir.InstUDiv:
		panic("not yet implemented")
	case *ir.InstSDiv:
		panic("not yet implemented")
	case *ir.InstFDiv:
		panic("not yet implemented")
	case *ir.InstURem:
		panic("not yet implemented")
	case *ir.InstSRem:
		panic("not yet implemented")
	case *ir.InstFRem:
		panic("not yet implemented")
	// Bitwise instructions.
	case *ir.InstShl:
		panic("not yet implemented")
	case *ir.InstLShr:
		panic("not yet implemented")
	case *ir.InstAShr:
		panic("not yet implemented")
	case *ir.InstAnd:
		panic("not yet implemented")
	case *ir.InstOr:
		panic("not yet implemented")
	case *ir.InstXor:
		panic("not yet implemented")
	// Memory instructions.
	case *ir.InstAlloca:
		panic("not yet implemented")
	case *ir.InstLoad:
		panic("not yet implemented")
	case *ir.InstStore:
		panic("not yet implemented")
	case *ir.InstGetElementPtr:
		panic("not yet implemented")
	// Conversion instructions.
	case *ir.InstTrunc:
		panic("not yet implemented")
	case *ir.InstZExt:
		panic("not yet implemented")
	case *ir.InstSExt:
		panic("not yet implemented")
	case *ir.InstFPTrunc:
		panic("not yet implemented")
	case *ir.InstFPExt:
		panic("not yet implemented")
	case *ir.InstFPToUI:
		panic("not yet implemented")
	case *ir.InstFPToSI:
		panic("not yet implemented")
	case *ir.InstUIToFP:
		panic("not yet implemented")
	case *ir.InstSIToFP:
		panic("not yet implemented")
	case *ir.InstPtrToInt:
		panic("not yet implemented")
	case *ir.InstIntToPtr:
		panic("not yet implemented")
	case *ir.InstBitCast:
		panic("not yet implemented")
	case *ir.InstAddrSpaceCast:
		panic("not yet implemented")
	// Other instructions.
	case *ir.InstICmp:
		panic("not yet implemented")
	case *ir.InstFCmp:
		panic("not yet implemented")
	case *ir.InstPhi:
		panic("not yet implemented")
	case *ir.InstSelect:
		panic("not yet implemented")
	case *ir.InstCall:
		panic("not yet implemented")
	default:
		panic(fmt.Errorf("support for instruction %T not yet implemented", inst))
	}
}

// checkTerm validates the semantics of the given terminator.
func checkTerm(term ir.Terminator) error {
	switch term := term.(type) {
	case *ir.TermRet:
		panic("not yet implemented")
	case *ir.TermBr:
		panic("not yet implemented")
	case *ir.TermCondBr:
		panic("not yet implemented")
	case *ir.TermSwitch:
		panic("not yet implemented")
	case *ir.TermUnreachable:
		panic("not yet implemented")
	default:
		panic(fmt.Errorf("support for instruction %T not yet implemented", term))
	}
}
