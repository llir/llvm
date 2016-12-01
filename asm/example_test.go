package asm_test

import (
	"log"

	"github.com/kr/pretty"
	"github.com/llir/llvm/asm"
)

func Example() {
	// Parse the LLVM IR assembly file `rand.ll`.
	m, err := asm.ParseFile("internal/testdata/rand.ll")
	if err != nil {
		log.Fatal(err)
	}
	// Pretty-print the data types of the parsed LLVM IR module.
	pretty.Println(m)
	// Output:
	//
	// &ir.Module{
	//     globals: {
	//         &ir.Global{
	//             name:    "seed",
	//             content: &types.IntType{size:32},
	//             init:    &constant.Int{
	//                 x:   &big.Int{},
	//                 typ: &types.IntType{(CYCLIC REFERENCE)},
	//             },
	//             typ: &types.PointerType{
	//                 elem:  &types.IntType{(CYCLIC REFERENCE)},
	//                 space: 0,
	//             },
	//             immutable: false,
	//         },
	//     },
	//     funcs: {
	//         &ir.Function{
	//             parent: &ir.Module{(CYCLIC REFERENCE)},
	//             name:   "abs",
	//             sig:    &types.FuncType{
	//                 ret:    &types.IntType{size:32},
	//                 params: {
	//                     &types.Param{
	//                         name: "x",
	//                         typ:  &types.IntType{size:32},
	//                     },
	//                 },
	//                 variadic: false,
	//             },
	//             blocks: nil,
	//             typ:    &types.PointerType{
	//                 elem:  &types.FuncType{(CYCLIC REFERENCE)},
	//                 space: 0,
	//             },
	//         },
	//         &ir.Function{
	//             parent: &ir.Module{(CYCLIC REFERENCE)},
	//             name:   "rand",
	//             sig:    &types.FuncType{
	//                 ret:      &types.IntType{size:32},
	//                 params:   nil,
	//                 variadic: false,
	//             },
	//             blocks: {
	//                 &ir.BasicBlock{
	//                     parent: &ir.Function{(CYCLIC REFERENCE)},
	//                     name:   "0",
	//                     insts:  {
	//                         &ir.InstLoad{
	//                             parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                             ident:  "1",
	//                             src:    &ir.Global{(CYCLIC REFERENCE)},
	//                             typ:    &types.IntType{size:32},
	//                         },
	//                         &ir.InstMul{
	//                             parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                             ident:  "2",
	//                             x:      &ir.InstLoad{(CYCLIC REFERENCE)},
	//                             y:      &constant.Int{
	//                                 x:  &big.Int{
	//                                     neg: false,
	//                                     abs: {0x15a4e35},
	//                                 },
	//                                 typ: &types.IntType{size:32},
	//                             },
	//                         },
	//                         &ir.InstAdd{
	//                             parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                             ident:  "3",
	//                             x:      &ir.InstMul{(CYCLIC REFERENCE)},
	//                             y:      &constant.Int{
	//                                 x:  &big.Int{
	//                                     neg: false,
	//                                     abs: {0x1},
	//                                 },
	//                                 typ: &types.IntType{size:32},
	//                             },
	//                         },
	//                         &ir.InstStore{
	//                             parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                             src:    &ir.InstAdd{(CYCLIC REFERENCE)},
	//                             dst:    &ir.Global{(CYCLIC REFERENCE)},
	//                         },
	//                         &ir.InstCall{
	//                             parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                             ident:  "4",
	//                             callee: &ir.Function{(CYCLIC REFERENCE)},
	//                             args:   {
	//                                 &ir.InstAdd{(CYCLIC REFERENCE)},
	//                             },
	//                         },
	//                     },
	//                     term: &ir.TermRet{
	//                         parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                         x:      &ir.InstCall{
	//                             parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                             ident:  "4",
	//                             callee: &ir.Function{(CYCLIC REFERENCE)},
	//                             args:   {
	//                                 &ir.InstAdd{(CYCLIC REFERENCE)},
	//                             },
	//                         },
	//                     },
	//                 },
	//             },
	//             typ: &types.PointerType{
	//                 elem:  &types.FuncType{(CYCLIC REFERENCE)},
	//                 space: 0,
	//             },
	//         },
	//     },
	// }
}
