package llvm

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
	//             GlobalIdent: ir.GlobalIdent{GlobalName:"seed", GlobalID:0},
	//             Immutable:   false,
	//             ContentType: &types.IntType{TypeName:"", BitSize:0x20},
	//             Init:        &constant.Int{
	//                 Typ: &types.IntType{(CYCLIC REFERENCE)},
	//                 X:   &big.Int{},
	//             },
	//             Typ: &types.PointerType{
	//                 TypeName:  "",
	//                 ElemType:  &types.IntType{(CYCLIC REFERENCE)},
	//                 AddrSpace: 0x0,
	//             },
	//             Linkage:               0x0,
	//             Preemption:            0x0,
	//             Visibility:            0x0,
	//             DLLStorageClass:       0x0,
	//             TLSModel:              0x0,
	//             UnnamedAddr:           0x0,
	//             AddrSpace:             0x0,
	//             ExternallyInitialized: false,
	//             Section:               "",
	//             Partition:             "",
	//             Comdat:                (*ir.ComdatDef)(nil),
	//             Align:                 0x0,
	//             FuncAttrs:             nil,
	//             Metadata:              nil,
	//         },
	//     },
	//     Funcs: {
	//         &ir.Func{
	//             GlobalIdent: ir.GlobalIdent{GlobalName:"abs", GlobalID:0},
	//             Sig:         &types.FuncType{
	//                 TypeName: "",
	//                 RetType:  &types.IntType{TypeName:"", BitSize:0x20},
	//                 Params:   {
	//                     &types.IntType{TypeName:"", BitSize:0x20},
	//                 },
	//                 Variadic: false,
	//             },
	//             Params: {
	//                 &ir.Param{
	//                     LocalIdent: ir.LocalIdent{LocalName:"x", LocalID:0},
	//                     Typ:        &types.IntType{TypeName:"", BitSize:0x20},
	//                     Attrs:      nil,
	//                 },
	//             },
	//             Blocks: nil,
	//             Typ:    &types.PointerType{
	//                 TypeName:  "",
	//                 ElemType:  &types.FuncType{(CYCLIC REFERENCE)},
	//                 AddrSpace: 0x0,
	//             },
	//             Linkage:         0x0,
	//             Preemption:      0x0,
	//             Visibility:      0x0,
	//             DLLStorageClass: 0x0,
	//             CallingConv:     0x0,
	//             ReturnAttrs:     nil,
	//             UnnamedAddr:     0x0,
	//             AddrSpace:       0x0,
	//             FuncAttrs:       nil,
	//             Section:         "",
	//             Partition:       "",
	//             Comdat:          (*ir.ComdatDef)(nil),
	//             Align:           0x0,
	//             GC:              "",
	//             Prefix:          nil,
	//             Prologue:        nil,
	//             Personality:     nil,
	//             UseListOrders:   nil,
	//             Metadata:        nil,
	//             Parent:          &ir.Module{(CYCLIC REFERENCE)},
	//             mu:              sync.Mutex{},
	//         },
	//         &ir.Func{
	//             GlobalIdent: ir.GlobalIdent{GlobalName:"rand", GlobalID:0},
	//             Sig:         &types.FuncType{
	//                 TypeName: "",
	//                 RetType:  &types.IntType{TypeName:"", BitSize:0x20},
	//                 Params:   nil,
	//                 Variadic: false,
	//             },
	//             Params: nil,
	//             Blocks: {
	//                 &ir.Block{
	//                     LocalIdent: ir.LocalIdent{},
	//                     Insts:      {
	//                         &ir.InstLoad{
	//                             LocalIdent: ir.LocalIdent{LocalName:"", LocalID:1},
	//                             ElemType:   &types.IntType{TypeName:"", BitSize:0x20},
	//                             Src:        &ir.Global{(CYCLIC REFERENCE)},
	//                             Atomic:     false,
	//                             Volatile:   false,
	//                             SyncScope:  "",
	//                             Ordering:   0x0,
	//                             Align:      0x0,
	//                             Metadata:   nil,
	//                         },
	//                         &ir.InstMul{
	//                             LocalIdent: ir.LocalIdent{LocalName:"", LocalID:2},
	//                             X:          &ir.InstLoad{(CYCLIC REFERENCE)},
	//                             Y:          &constant.Int{
	//                                 Typ: &types.IntType{(CYCLIC REFERENCE)},
	//                                 X:   &big.Int{
	//                                     neg: false,
	//                                     abs: {0x15a4e35},
	//                                 },
	//                             },
	//                             Typ:           &types.IntType{TypeName:"", BitSize:0x20},
	//                             OverflowFlags: nil,
	//                             Metadata:      nil,
	//                         },
	//                         &ir.InstAdd{
	//                             LocalIdent: ir.LocalIdent{LocalName:"", LocalID:3},
	//                             X:          &ir.InstMul{(CYCLIC REFERENCE)},
	//                             Y:          &constant.Int{
	//                                 Typ: &types.IntType{(CYCLIC REFERENCE)},
	//                                 X:   &big.Int{
	//                                     neg: false,
	//                                     abs: {0x1},
	//                                 },
	//                             },
	//                             Typ:           &types.IntType{TypeName:"", BitSize:0x20},
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
	//                             Align:     0x0,
	//                             Metadata:  nil,
	//                         },
	//                         &ir.InstCall{
	//                             LocalIdent: ir.LocalIdent{LocalName:"", LocalID:4},
	//                             Callee:     &ir.Func{(CYCLIC REFERENCE)},
	//                             Args:       {
	//                                 &ir.InstAdd{(CYCLIC REFERENCE)},
	//                             },
	//                             Typ:            &types.IntType{TypeName:"", BitSize:0x20},
	//                             Tail:           0x0,
	//                             FastMathFlags:  nil,
	//                             CallingConv:    0x0,
	//                             ReturnAttrs:    nil,
	//                             AddrSpace:      0x0,
	//                             FuncAttrs:      nil,
	//                             OperandBundles: nil,
	//                             Metadata:       nil,
	//                         },
	//                     },
	//                     Term: &ir.TermRet{
	//                         X:  &ir.InstCall{
	//                             LocalIdent: ir.LocalIdent{LocalName:"", LocalID:4},
	//                             Callee:     &ir.Func{(CYCLIC REFERENCE)},
	//                             Args:       {
	//                                 &ir.InstAdd{(CYCLIC REFERENCE)},
	//                             },
	//                             Typ:            &types.IntType{TypeName:"", BitSize:0x20},
	//                             Tail:           0x0,
	//                             FastMathFlags:  nil,
	//                             CallingConv:    0x0,
	//                             ReturnAttrs:    nil,
	//                             AddrSpace:      0x0,
	//                             FuncAttrs:      nil,
	//                             OperandBundles: nil,
	//                             Metadata:       nil,
	//                         },
	//                         Metadata: nil,
	//                     },
	//                     Parent: &ir.Func{(CYCLIC REFERENCE)},
	//                 },
	//             },
	//             Typ: &types.PointerType{
	//                 TypeName:  "",
	//                 ElemType:  &types.FuncType{(CYCLIC REFERENCE)},
	//                 AddrSpace: 0x0,
	//             },
	//             Linkage:         0x0,
	//             Preemption:      0x0,
	//             Visibility:      0x0,
	//             DLLStorageClass: 0x0,
	//             CallingConv:     0x0,
	//             ReturnAttrs:     nil,
	//             UnnamedAddr:     0x0,
	//             AddrSpace:       0x0,
	//             FuncAttrs:       nil,
	//             Section:         "",
	//             Partition:       "",
	//             Comdat:          (*ir.ComdatDef)(nil),
	//             Align:           0x0,
	//             GC:              "",
	//             Prefix:          nil,
	//             Prologue:        nil,
	//             Personality:     nil,
	//             UseListOrders:   nil,
	//             Metadata:        nil,
	//             Parent:          &ir.Module{(CYCLIC REFERENCE)},
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
	//     NamedMetadataDefs: {
	//     },
	//     MetadataDefs:    nil,
	//     UseListOrders:   nil,
	//     UseListOrderBBs: nil,
	// }
}
