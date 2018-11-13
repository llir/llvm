package asm_test

import (
	"log"

	"github.com/kr/pretty"
	"github.com/llir/llvm/asm"
)

func Example() {
	// Parse the LLVM IR assembly file `rand.ll`.
	m, err := asm.ParseFile("testdata/rand.ll")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	// Pretty-print the data types of the parsed LLVM IR module.
	pretty.Println(m)
	// Output:
	//
	// &ir.Module{
	//     TypeDefs: nil,
	//     Globals:  {
	//         &ir.Global{
	//             GlobalName:  "seed",
	//             Immutable:   false,
	//             ContentType: &types.IntType{Alias:"", BitSize:32},
	//             Init:        &constant.Int{
	//                 Typ: &types.IntType{(CYCLIC REFERENCE)},
	//                 X:   &big.Int{},
	//             },
	//             Typ: &types.PointerType{
	//                 Alias:     "",
	//                 ElemType:  &types.IntType{(CYCLIC REFERENCE)},
	//                 AddrSpace: 0,
	//             },
	//             Linkage:               0x0,
	//             Preemption:            0x0,
	//             Visibility:            0x0,
	//             DLLStorageClass:       0x0,
	//             TLSModel:              0x0,
	//             UnnamedAddr:           0x0,
	//             ExternallyInitialized: false,
	//             Section:               "",
	//             Comdat:                (*ir.ComdatDef)(nil),
	//             Align:                 0,
	//             FuncAttrs:             nil,
	//             Metadata:              nil,
	//         },
	//     },
	//     Funcs: {
	//         &ir.Function{
	//             GlobalName: "abs",
	//             Sig:        &types.FuncType{
	//                 Alias:   "",
	//                 RetType: &types.IntType{Alias:"", BitSize:32},
	//                 Params:  {
	//                     &types.IntType{Alias:"", BitSize:32},
	//                 },
	//                 Variadic: false,
	//             },
	//             Params: {
	//                 &ir.Param{
	//                     LocalName: "x",
	//                     Typ:       &types.IntType{Alias:"", BitSize:32},
	//                     Attrs:     nil,
	//                 },
	//             },
	//             Blocks: nil,
	//             Typ:    &types.PointerType{
	//                 Alias:     "",
	//                 ElemType:  &types.FuncType{(CYCLIC REFERENCE)},
	//                 AddrSpace: 0,
	//             },
	//             Linkage:         0x0,
	//             Preemption:      0x0,
	//             Visibility:      0x0,
	//             DLLStorageClass: 0x0,
	//             CallingConv:     0x0,
	//             ReturnAttrs:     nil,
	//             UnnamedAddr:     0x0,
	//             FuncAttrs:       nil,
	//             Section:         "",
	//             Comdat:          (*ir.ComdatDef)(nil),
	//             GC:              "",
	//             Prefix:          nil,
	//             Prologue:        nil,
	//             Personality:     nil,
	//             UseListOrders:   nil,
	//             Metadata:        nil,
	//             mu:              sync.Mutex{},
	//         },
	//         &ir.Function{
	//             GlobalName: "rand",
	//             Sig:        &types.FuncType{
	//                 Alias:    "",
	//                 RetType:  &types.IntType{Alias:"", BitSize:32},
	//                 Params:   nil,
	//                 Variadic: false,
	//             },
	//             Params: nil,
	//             Blocks: {
	//                 &ir.BasicBlock{
	//                     LocalName: "0",
	//                     Insts:     {
	//                         &ir.InstLoad{
	//                             LocalName: "1",
	//                             Src:       &ir.Global{(CYCLIC REFERENCE)},
	//                             Typ:       &types.IntType{Alias:"", BitSize:32},
	//                             Atomic:    false,
	//                             Volatile:  false,
	//                             SyncScope: "",
	//                             Ordering:  0x0,
	//                             Align:     0,
	//                             Metadata:  nil,
	//                         },
	//                         &ir.InstMul{
	//                             LocalName: "2",
	//                             X:         &ir.InstLoad{(CYCLIC REFERENCE)},
	//                             Y:         &constant.Int{
	//                                 Typ: &types.IntType{(CYCLIC REFERENCE)},
	//                                 X:   &big.Int{
	//                                     neg: false,
	//                                     abs: {0x15a4e35},
	//                                 },
	//                             },
	//                             Typ:           &types.IntType{Alias:"", BitSize:32},
	//                             OverflowFlags: nil,
	//                             Metadata:      nil,
	//                         },
	//                         &ir.InstAdd{
	//                             LocalName: "3",
	//                             X:         &ir.InstMul{(CYCLIC REFERENCE)},
	//                             Y:         &constant.Int{
	//                                 Typ: &types.IntType{(CYCLIC REFERENCE)},
	//                                 X:   &big.Int{
	//                                     neg: false,
	//                                     abs: {0x1},
	//                                 },
	//                             },
	//                             Typ:           &types.IntType{Alias:"", BitSize:32},
	//                             OverflowFlags: nil,
	//                             Metadata:      nil,
	//                         },
	//                         &ir.InstStore{
	//                             Src:       &ir.InstAdd{(CYCLIC REFERENCE)},
	//                             Dst:       &ir.Global{(CYCLIC REFERENCE)},
	//                             Atomic:    false,
	//                             Volatile:  false,
	//                             SyncScope: "",
	//                             Ordering:  0x0,
	//                             Align:     0,
	//                             Metadata:  nil,
	//                         },
	//                         &ir.InstCall{
	//                             LocalName: "4",
	//                             Callee:    &ir.Function{(CYCLIC REFERENCE)},
	//                             Args:      {
	//                                 &ir.InstAdd{(CYCLIC REFERENCE)},
	//                             },
	//                             Typ:            &types.IntType{Alias:"", BitSize:32},
	//                             Tail:           0x0,
	//                             FastMathFlags:  nil,
	//                             CallingConv:    0x0,
	//                             ReturnAttrs:    nil,
	//                             AddrSpace:      0,
	//                             FuncAttrs:      nil,
	//                             OperandBundles: nil,
	//                             Metadata:       nil,
	//                         },
	//                     },
	//                     Term: &ir.TermRet{
	//                         X:  &ir.InstCall{
	//                             LocalName: "4",
	//                             Callee:    &ir.Function{(CYCLIC REFERENCE)},
	//                             Args:      {
	//                                 &ir.InstAdd{(CYCLIC REFERENCE)},
	//                             },
	//                             Typ:            &types.IntType{Alias:"", BitSize:32},
	//                             Tail:           0x0,
	//                             FastMathFlags:  nil,
	//                             CallingConv:    0x0,
	//                             ReturnAttrs:    nil,
	//                             AddrSpace:      0,
	//                             FuncAttrs:      nil,
	//                             OperandBundles: nil,
	//                             Metadata:       nil,
	//                         },
	//                         Metadata: nil,
	//                     },
	//                 },
	//             },
	//             Typ: &types.PointerType{
	//                 Alias:     "",
	//                 ElemType:  &types.FuncType{(CYCLIC REFERENCE)},
	//                 AddrSpace: 0,
	//             },
	//             Linkage:         0x0,
	//             Preemption:      0x0,
	//             Visibility:      0x0,
	//             DLLStorageClass: 0x0,
	//             CallingConv:     0x0,
	//             ReturnAttrs:     nil,
	//             UnnamedAddr:     0x0,
	//             FuncAttrs:       nil,
	//             Section:         "",
	//             Comdat:          (*ir.ComdatDef)(nil),
	//             GC:              "",
	//             Prefix:          nil,
	//             Prologue:        nil,
	//             Personality:     nil,
	//             UseListOrders:   nil,
	//             Metadata:        nil,
	//             mu:              sync.Mutex{},
	//         },
	//     },
	//     SourceFilename:    "",
	//     DataLayout:        "",
	//     TargetTriple:      "",
	//     ModuleAsms:        nil,
	//     ComdatDefs:        nil,
	//     Aliases:           nil,
	//     IFuncs:            nil,
	//     AttrGroupDefs:     nil,
	//     NamedMetadataDefs: nil,
	//     MetadataDefs:      nil,
	//     UseListOrders:     nil,
	//     UseListOrderBBs:   nil,
	// }
}
