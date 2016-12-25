package ast_test

import (
	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/ir/constant"
)

// Valutate that the relevant types satisfy the ast.Constant interface.
var (
	// Simple constants.
	_ ast.Constant = &ast.IntConst{}
	_ ast.Constant = &ast.FloatConst{}
	_ ast.Constant = &ast.NullConst{}
	// Complex constants.
	_ ast.Constant = &ast.VectorConst{}
	_ ast.Constant = &ast.ArrayConst{}
	_ ast.Constant = &ast.CharArrayConst{}
	_ ast.Constant = &ast.StructConst{}
	_ ast.Constant = &ast.ZeroInitializerConst{}
	// Global variable and function addresses
	_ ast.Constant = &ast.Global{}
	_ ast.Constant = &ast.Function{}
)

// Valutate that the relevant types satisfy the constant.Expr interface.
var (
	// Constant expressions.
	// Binary instructions
	_ constant.Expr = &constant.ExprAdd{}
	_ constant.Expr = &constant.ExprFAdd{}
	_ constant.Expr = &constant.ExprSub{}
	_ constant.Expr = &constant.ExprFSub{}
	_ constant.Expr = &constant.ExprMul{}
	_ constant.Expr = &constant.ExprFMul{}
	_ constant.Expr = &constant.ExprUDiv{}
	_ constant.Expr = &constant.ExprSDiv{}
	_ constant.Expr = &constant.ExprFDiv{}
	_ constant.Expr = &constant.ExprURem{}
	_ constant.Expr = &constant.ExprSRem{}
	_ constant.Expr = &constant.ExprFRem{}
	// Bitwise instructions
	_ constant.Expr = &constant.ExprShl{}
	_ constant.Expr = &constant.ExprLShr{}
	_ constant.Expr = &constant.ExprAShr{}
	_ constant.Expr = &constant.ExprAnd{}
	_ constant.Expr = &constant.ExprOr{}
	_ constant.Expr = &constant.ExprXor{}
	// Memory instructions
	_ constant.Expr = &constant.ExprGetElementPtr{}
	// Conversion instructions
	_ constant.Expr = &constant.ExprTrunc{}
	_ constant.Expr = &constant.ExprZExt{}
	_ constant.Expr = &constant.ExprSExt{}
	_ constant.Expr = &constant.ExprFPTrunc{}
	_ constant.Expr = &constant.ExprFPExt{}
	_ constant.Expr = &constant.ExprFPToUI{}
	_ constant.Expr = &constant.ExprFPToSI{}
	_ constant.Expr = &constant.ExprUIToFP{}
	_ constant.Expr = &constant.ExprSIToFP{}
	_ constant.Expr = &constant.ExprPtrToInt{}
	_ constant.Expr = &constant.ExprIntToPtr{}
	_ constant.Expr = &constant.ExprBitCast{}
	_ constant.Expr = &constant.ExprAddrSpaceCast{}
	// Other instructions
	_ constant.Expr = &constant.ExprICmp{}
	_ constant.Expr = &constant.ExprFCmp{}
	_ constant.Expr = &constant.ExprSelect{}
)

// Valutate that the relevant types satisfy the ast.Instruction interface.
var (
	// Binary instructions
	_ ast.Instruction = &ast.InstAdd{}
	_ ast.Instruction = &ast.InstFAdd{}
	_ ast.Instruction = &ast.InstSub{}
	_ ast.Instruction = &ast.InstFSub{}
	_ ast.Instruction = &ast.InstMul{}
	_ ast.Instruction = &ast.InstFMul{}
	_ ast.Instruction = &ast.InstUDiv{}
	_ ast.Instruction = &ast.InstSDiv{}
	_ ast.Instruction = &ast.InstFDiv{}
	_ ast.Instruction = &ast.InstURem{}
	_ ast.Instruction = &ast.InstSRem{}
	_ ast.Instruction = &ast.InstFRem{}
	// Bitwise instructions
	_ ast.Instruction = &ast.InstShl{}
	_ ast.Instruction = &ast.InstLShr{}
	_ ast.Instruction = &ast.InstAShr{}
	_ ast.Instruction = &ast.InstAnd{}
	_ ast.Instruction = &ast.InstOr{}
	_ ast.Instruction = &ast.InstXor{}
	// Vector instructions
	// Aggregate instructions
	// Memory instructions
	_ ast.Instruction = &ast.InstAlloca{}
	_ ast.Instruction = &ast.InstLoad{}
	_ ast.Instruction = &ast.InstStore{}
	_ ast.Instruction = &ast.InstGetElementPtr{}
	// Conversion instructions
	_ ast.Instruction = &ast.InstTrunc{}
	_ ast.Instruction = &ast.InstZExt{}
	_ ast.Instruction = &ast.InstSExt{}
	_ ast.Instruction = &ast.InstFPTrunc{}
	_ ast.Instruction = &ast.InstFPExt{}
	_ ast.Instruction = &ast.InstFPToUI{}
	_ ast.Instruction = &ast.InstFPToSI{}
	_ ast.Instruction = &ast.InstUIToFP{}
	_ ast.Instruction = &ast.InstSIToFP{}
	_ ast.Instruction = &ast.InstPtrToInt{}
	_ ast.Instruction = &ast.InstIntToPtr{}
	_ ast.Instruction = &ast.InstBitCast{}
	_ ast.Instruction = &ast.InstAddrSpaceCast{}
	// Other instructions
	_ ast.Instruction = &ast.InstICmp{}
	_ ast.Instruction = &ast.InstFCmp{}
	_ ast.Instruction = &ast.InstPhi{}
	_ ast.Instruction = &ast.InstSelect{}
	_ ast.Instruction = &ast.InstCall{}
)

// Valutate that the relevant types satisfy the ast.Terminator interface.
var (
	// Terminators
	_ ast.Terminator = &ast.TermRet{}
	_ ast.Terminator = &ast.TermBr{}
	_ ast.Terminator = &ast.TermCondBr{}
	_ ast.Terminator = &ast.TermSwitch{}
	_ ast.Terminator = &ast.TermUnreachable{}
)

// Valutate that the relevant types satisfy the ast.NamedValue interface.
var (
	_ ast.NamedValue = &ast.Global{}
	_ ast.NamedValue = &ast.GlobalDummy{}
	_ ast.NamedValue = &ast.Function{}
	_ ast.NamedValue = &ast.Param{}
	_ ast.NamedValue = &ast.BasicBlock{}
	_ ast.NamedValue = &ast.LocalDummy{}
	// Binary instructions
	_ ast.NamedValue = &ast.InstAdd{}
	_ ast.NamedValue = &ast.InstFAdd{}
	_ ast.NamedValue = &ast.InstSub{}
	_ ast.NamedValue = &ast.InstFSub{}
	_ ast.NamedValue = &ast.InstMul{}
	_ ast.NamedValue = &ast.InstFMul{}
	_ ast.NamedValue = &ast.InstUDiv{}
	_ ast.NamedValue = &ast.InstSDiv{}
	_ ast.NamedValue = &ast.InstFDiv{}
	_ ast.NamedValue = &ast.InstURem{}
	_ ast.NamedValue = &ast.InstSRem{}
	_ ast.NamedValue = &ast.InstFRem{}
	// Bitwise instructions
	_ ast.NamedValue = &ast.InstShl{}
	_ ast.NamedValue = &ast.InstLShr{}
	_ ast.NamedValue = &ast.InstAShr{}
	_ ast.NamedValue = &ast.InstAnd{}
	_ ast.NamedValue = &ast.InstOr{}
	_ ast.NamedValue = &ast.InstXor{}
	// Vector instructions
	// Aggregate instructions
	// Memory instructions
	_ ast.NamedValue = &ast.InstAlloca{}
	_ ast.NamedValue = &ast.InstLoad{}
	_ ast.NamedValue = &ast.InstGetElementPtr{}
	// Conversion instructions
	_ ast.NamedValue = &ast.InstTrunc{}
	_ ast.NamedValue = &ast.InstZExt{}
	_ ast.NamedValue = &ast.InstSExt{}
	_ ast.NamedValue = &ast.InstFPTrunc{}
	_ ast.NamedValue = &ast.InstFPExt{}
	_ ast.NamedValue = &ast.InstFPToUI{}
	_ ast.NamedValue = &ast.InstFPToSI{}
	_ ast.NamedValue = &ast.InstUIToFP{}
	_ ast.NamedValue = &ast.InstSIToFP{}
	_ ast.NamedValue = &ast.InstPtrToInt{}
	_ ast.NamedValue = &ast.InstIntToPtr{}
	_ ast.NamedValue = &ast.InstBitCast{}
	_ ast.NamedValue = &ast.InstAddrSpaceCast{}
	// Other instructions
	_ ast.NamedValue = &ast.InstICmp{}
	_ ast.NamedValue = &ast.InstFCmp{}
	_ ast.NamedValue = &ast.InstPhi{}
	_ ast.NamedValue = &ast.InstSelect{}
	_ ast.NamedValue = &ast.InstCall{}
)

// Valutate that the relevant types satisfy the ast.Type interface.
var (
	_ ast.Type = &ast.VoidType{}
	_ ast.Type = &ast.LabelType{}
	_ ast.Type = &ast.IntType{}
	_ ast.Type = &ast.FloatType{}
	_ ast.Type = &ast.FuncType{}
	_ ast.Type = &ast.PointerType{}
	_ ast.Type = &ast.VectorType{}
	_ ast.Type = &ast.ArrayType{}
	_ ast.Type = &ast.StructType{}
)
