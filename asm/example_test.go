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
	//     types:   nil,
	//     globals: {
	//         &ir.Global{
	//             name: "seed",
	//             typ:  &types.PointerType{
	//                 elem:  &types.IntType{size:32},
	//                 space: 0,
	//             },
	//             content: &types.IntType{size:32},
	//             init:    &constant.Int{
	//                 typ: &types.IntType{(CYCLIC REFERENCE)},
	//                 x:   &big.Int{},
	//             },
	//             isConst: false,
	//         },
	//     },
	//     funcs: {
	//         &ir.Function{
	//             parent: &ir.Module{(CYCLIC REFERENCE)},
	//             name:   "abs",
	//             typ:    &types.PointerType{
	//                 elem: &types.FuncType{
	//                     ret:    &types.IntType{size:32},
	//                     params: {
	//                         &types.Param{
	//                             name: "x",
	//                             typ:  &types.IntType{size:32},
	//                         },
	//                     },
	//                     variadic: false,
	//                 },
	//                 space: 0,
	//             },
	//             sig: &types.FuncType{
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
	//         },
	//         &ir.Function{
	//             parent: &ir.Module{(CYCLIC REFERENCE)},
	//             name:   "rand",
	//             typ:    &types.PointerType{
	//                 elem: &types.FuncType{
	//                     ret:      &types.IntType{size:32},
	//                     params:   nil,
	//                     variadic: false,
	//                 },
	//                 space: 0,
	//             },
	//             sig: &types.FuncType{
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
	//                             name:   "1",
	//                             typ:    &types.IntType{size:32},
	//                             src:    &ir.Global{(CYCLIC REFERENCE)},
	//                         },
	//                         &ir.InstMul{
	//                             parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                             name:   "2",
	//                             x:      &ir.InstLoad{(CYCLIC REFERENCE)},
	//                             y:      &constant.Int{
	//                                 typ: &types.IntType{size:32},
	//                                 x:   &big.Int{
	//                                     neg: false,
	//                                     abs: {0x15a4e35},
	//                                 },
	//                             },
	//                         },
	//                         &ir.InstAdd{
	//                             parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                             name:   "3",
	//                             x:      &ir.InstMul{(CYCLIC REFERENCE)},
	//                             y:      &constant.Int{
	//                                 typ: &types.IntType{size:32},
	//                                 x:   &big.Int{
	//                                     neg: false,
	//                                     abs: {0x1},
	//                                 },
	//                             },
	//                         },
	//                         &ir.InstStore{
	//                             parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                             src:    &ir.InstAdd{(CYCLIC REFERENCE)},
	//                             dst:    &ir.Global{(CYCLIC REFERENCE)},
	//                         },
	//                         &ir.InstCall{
	//                             parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                             name:   "4",
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
	//                             name:   "4",
	//                             callee: &ir.Function{(CYCLIC REFERENCE)},
	//                             args:   {
	//                                 &ir.InstAdd{(CYCLIC REFERENCE)},
	//                             },
	//                         },
	//                     },
	//                 },
	//             },
	//         },
	//     },
	// }

}
