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
	//                 typ:  &types.IntType{(CYCLIC REFERENCE)},
	//                 x:    &big.Int{},
	//                 used: constant.used{
	//                     uses: {
	//                         &ir.constantTracker{
	//                             orig: &&constant.Int{(CYCLIC REFERENCE)},
	//                             user: &ir.Global{(CYCLIC REFERENCE)},
	//                         },
	//                     },
	//                 },
	//             },
	//             isConst: false,
	//             used:    ir.used{
	//                 uses: {
	//                     &ir.valueTracker{
	//                         orig: &&ir.Global{(CYCLIC REFERENCE)},
	//                         user: &ir.InstLoad{
	//                             parent: &ir.BasicBlock{
	//                                 parent: &ir.Function{
	//                                     parent: &ir.Module{(CYCLIC REFERENCE)},
	//                                     name:   "rand",
	//                                     typ:    &types.PointerType{
	//                                         elem: &types.FuncType{
	//                                             ret:      &types.IntType{size:32},
	//                                             params:   nil,
	//                                             variadic: false,
	//                                         },
	//                                         space: 0,
	//                                     },
	//                                     sig: &types.FuncType{
	//                                         ret:      &types.IntType{size:32},
	//                                         params:   nil,
	//                                         variadic: false,
	//                                     },
	//                                     params: nil,
	//                                     blocks: {
	//                                         &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                     },
	//                                     used: ir.used{},
	//                                 },
	//                                 name:  "0",
	//                                 insts: {
	//                                     &ir.InstLoad{(CYCLIC REFERENCE)},
	//                                     &ir.InstMul{
	//                                         parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                         name:   "2",
	//                                         x:      &ir.InstLoad{(CYCLIC REFERENCE)},
	//                                         y:      &constant.Int{
	//                                             typ: &types.IntType{size:32},
	//                                             x:   &big.Int{
	//                                                 neg: false,
	//                                                 abs: {0x15a4e35},
	//                                             },
	//                                             used: constant.used{
	//                                                 uses: {
	//                                                     &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                         used: ir.used{
	//                                             uses: {
	//                                                 &ir.valueTracker{
	//                                                     orig: &!%v(DEPTH EXCEEDED),
	//                                                     user: &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                     },
	//                                     &ir.InstAdd{
	//                                         parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                         name:   "3",
	//                                         x:      &ir.InstMul{(CYCLIC REFERENCE)},
	//                                         y:      &constant.Int{
	//                                             typ: &types.IntType{size:32},
	//                                             x:   &big.Int{
	//                                                 neg: false,
	//                                                 abs: {0x1},
	//                                             },
	//                                             used: constant.used{
	//                                                 uses: {
	//                                                     &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                         used: ir.used{
	//                                             uses: {
	//                                                 &ir.valueTracker{
	//                                                     orig: &!%v(DEPTH EXCEEDED),
	//                                                     user: &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                                 &ir.valueTracker{
	//                                                     orig: &!%v(DEPTH EXCEEDED),
	//                                                     user: &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                     },
	//                                     &ir.InstStore{
	//                                         parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                         src:    &ir.InstAdd{(CYCLIC REFERENCE)},
	//                                         dst:    &ir.Global{(CYCLIC REFERENCE)},
	//                                     },
	//                                     &ir.InstCall{
	//                                         parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                         name:   "4",
	//                                         callee: &ir.Function{
	//                                             parent: &ir.Module{(CYCLIC REFERENCE)},
	//                                             name:   "abs",
	//                                             typ:    &types.PointerType{
	//                                                 elem:  &!%v(DEPTH EXCEEDED),
	//                                                 space: 0,
	//                                             },
	//                                             sig: &types.FuncType{
	//                                                 ret:    &!%v(DEPTH EXCEEDED),
	//                                                 params: {
	//                                                     &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                                 variadic: false,
	//                                             },
	//                                             params: {
	//                                                 &ir.Param{
	//                                                     Param: &!%v(DEPTH EXCEEDED),
	//                                                     used:  ir.used{},
	//                                                 },
	//                                             },
	//                                             blocks: nil,
	//                                             used:   ir.used{
	//                                                 uses: {
	//                                                     &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                         sig: &types.FuncType{
	//                                             ret:    &types.IntType{size:32},
	//                                             params: {
	//                                                 &types.Param{
	//                                                     name: "x",
	//                                                     typ:  &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                             variadic: false,
	//                                         },
	//                                         args: {
	//                                             &ir.InstAdd{(CYCLIC REFERENCE)},
	//                                         },
	//                                         used: ir.used{
	//                                             uses: {
	//                                                 &ir.valueTracker{
	//                                                     orig: &!%v(DEPTH EXCEEDED),
	//                                                     user: &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                     },
	//                                 },
	//                                 term: &ir.TermRet{
	//                                     parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                     x:      &ir.InstCall{
	//                                         parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                         name:   "4",
	//                                         callee: &ir.Function{
	//                                             parent: &ir.Module{(CYCLIC REFERENCE)},
	//                                             name:   "abs",
	//                                             typ:    &types.PointerType{
	//                                                 elem:  &!%v(DEPTH EXCEEDED),
	//                                                 space: 0,
	//                                             },
	//                                             sig:    &types.FuncType{(CYCLIC REFERENCE)},
	//                                             params: {
	//                                                 &ir.Param{
	//                                                     Param: &!%v(DEPTH EXCEEDED),
	//                                                     used:  ir.used{},
	//                                                 },
	//                                             },
	//                                             blocks: nil,
	//                                             used:   ir.used{
	//                                                 uses: {
	//                                                     &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                         sig: &types.FuncType{
	//                                             ret:    &types.IntType{size:32},
	//                                             params: {
	//                                                 &types.Param{
	//                                                     name: "x",
	//                                                     typ:  &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                             variadic: false,
	//                                         },
	//                                         args: {
	//                                             &ir.InstAdd{(CYCLIC REFERENCE)},
	//                                         },
	//                                         used: ir.used{
	//                                             uses: {
	//                                                 &ir.valueTracker{
	//                                                     orig: &!%v(DEPTH EXCEEDED),
	//                                                     user: &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                     },
	//                                 },
	//                                 used: ir.used{},
	//                             },
	//                             name: "1",
	//                             typ:  &types.IntType{size:32},
	//                             src:  &ir.Global{(CYCLIC REFERENCE)},
	//                             used: ir.used{
	//                                 uses: {
	//                                     &ir.valueTracker{
	//                                         orig: &&ir.InstLoad{(CYCLIC REFERENCE)},
	//                                         user: &ir.InstMul{
	//                                             parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                             name:   "2",
	//                                             x:      &ir.InstLoad{(CYCLIC REFERENCE)},
	//                                             y:      &constant.Int{
	//                                                 typ: &types.IntType{size:32},
	//                                                 x:   &big.Int{
	//                                                     neg: false,
	//                                                     abs: {0x15a4e35},
	//                                                 },
	//                                                 used: constant.used{
	//                                                     uses: {
	//                                                         &!%v(DEPTH EXCEEDED),
	//                                                     },
	//                                                 },
	//                                             },
	//                                             used: ir.used{
	//                                                 uses: {
	//                                                     &ir.valueTracker{
	//                                                         orig: &!%v(DEPTH EXCEEDED),
	//                                                         user: &!%v(DEPTH EXCEEDED),
	//                                                     },
	//                                                 },
	//                                             },
	//                                         },
	//                                     },
	//                                 },
	//                             },
	//                         },
	//                     },
	//                     &ir.valueTracker{
	//                         orig: &&ir.Global{(CYCLIC REFERENCE)},
	//                         user: &ir.InstStore{
	//                             parent: &ir.BasicBlock{
	//                                 parent: &ir.Function{
	//                                     parent: &ir.Module{(CYCLIC REFERENCE)},
	//                                     name:   "rand",
	//                                     typ:    &types.PointerType{
	//                                         elem:  &types.FuncType{(CYCLIC REFERENCE)},
	//                                         space: 0,
	//                                     },
	//                                     sig: &types.FuncType{
	//                                         ret:      &types.IntType{size:32},
	//                                         params:   nil,
	//                                         variadic: false,
	//                                     },
	//                                     params: nil,
	//                                     blocks: {
	//                                         &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                     },
	//                                     used: ir.used{},
	//                                 },
	//                                 name:  "0",
	//                                 insts: {
	//                                     &ir.InstLoad{(CYCLIC REFERENCE)},
	//                                     &ir.InstMul{
	//                                         parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                         name:   "2",
	//                                         x:      &ir.InstLoad{(CYCLIC REFERENCE)},
	//                                         y:      &constant.Int{
	//                                             typ: &types.IntType{size:32},
	//                                             x:   &big.Int{
	//                                                 neg: false,
	//                                                 abs: {0x15a4e35},
	//                                             },
	//                                             used: constant.used{
	//                                                 uses: {
	//                                                     &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                         used: ir.used{
	//                                             uses: {
	//                                                 &ir.valueTracker{
	//                                                     orig: &!%v(DEPTH EXCEEDED),
	//                                                     user: &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                     },
	//                                     &ir.InstAdd{
	//                                         parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                         name:   "3",
	//                                         x:      &ir.InstMul{(CYCLIC REFERENCE)},
	//                                         y:      &constant.Int{
	//                                             typ: &types.IntType{size:32},
	//                                             x:   &big.Int{
	//                                                 neg: false,
	//                                                 abs: {0x1},
	//                                             },
	//                                             used: constant.used{
	//                                                 uses: {
	//                                                     &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                         used: ir.used{
	//                                             uses: {
	//                                                 &ir.valueTracker{
	//                                                     orig: &!%v(DEPTH EXCEEDED),
	//                                                     user: &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                                 &ir.valueTracker{
	//                                                     orig: &!%v(DEPTH EXCEEDED),
	//                                                     user: &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                     },
	//                                     &ir.InstStore{(CYCLIC REFERENCE)},
	//                                     &ir.InstCall{
	//                                         parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                         name:   "4",
	//                                         callee: &ir.Function{
	//                                             parent: &ir.Module{(CYCLIC REFERENCE)},
	//                                             name:   "abs",
	//                                             typ:    &types.PointerType{
	//                                                 elem:  &!%v(DEPTH EXCEEDED),
	//                                                 space: 0,
	//                                             },
	//                                             sig:    &types.FuncType{(CYCLIC REFERENCE)},
	//                                             params: {
	//                                                 &ir.Param{
	//                                                     Param: &!%v(DEPTH EXCEEDED),
	//                                                     used:  ir.used{},
	//                                                 },
	//                                             },
	//                                             blocks: nil,
	//                                             used:   ir.used{
	//                                                 uses: {
	//                                                     &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                         sig: &types.FuncType{
	//                                             ret:    &types.IntType{size:32},
	//                                             params: {
	//                                                 &types.Param{
	//                                                     name: "x",
	//                                                     typ:  &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                             variadic: false,
	//                                         },
	//                                         args: {
	//                                             &ir.InstAdd{(CYCLIC REFERENCE)},
	//                                         },
	//                                         used: ir.used{
	//                                             uses: {
	//                                                 &ir.valueTracker{
	//                                                     orig: &!%v(DEPTH EXCEEDED),
	//                                                     user: &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                     },
	//                                 },
	//                                 term: &ir.TermRet{
	//                                     parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                     x:      &ir.InstCall{
	//                                         parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                         name:   "4",
	//                                         callee: &ir.Function{
	//                                             parent: &ir.Module{(CYCLIC REFERENCE)},
	//                                             name:   "abs",
	//                                             typ:    &types.PointerType{
	//                                                 elem:  &!%v(DEPTH EXCEEDED),
	//                                                 space: 0,
	//                                             },
	//                                             sig:    &types.FuncType{(CYCLIC REFERENCE)},
	//                                             params: {
	//                                                 &ir.Param{
	//                                                     Param: &!%v(DEPTH EXCEEDED),
	//                                                     used:  ir.used{},
	//                                                 },
	//                                             },
	//                                             blocks: nil,
	//                                             used:   ir.used{
	//                                                 uses: {
	//                                                     &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                         sig: &types.FuncType{
	//                                             ret:    &types.IntType{size:32},
	//                                             params: {
	//                                                 &types.Param{
	//                                                     name: "x",
	//                                                     typ:  &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                             variadic: false,
	//                                         },
	//                                         args: {
	//                                             &ir.InstAdd{(CYCLIC REFERENCE)},
	//                                         },
	//                                         used: ir.used{
	//                                             uses: {
	//                                                 &ir.valueTracker{
	//                                                     orig: &!%v(DEPTH EXCEEDED),
	//                                                     user: &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                     },
	//                                 },
	//                                 used: ir.used{},
	//                             },
	//                             src: &ir.InstAdd{
	//                                 parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                 name:   "3",
	//                                 x:      &ir.InstMul{
	//                                     parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                     name:   "2",
	//                                     x:      &ir.InstLoad{(CYCLIC REFERENCE)},
	//                                     y:      &constant.Int{
	//                                         typ: &types.IntType{size:32},
	//                                         x:   &big.Int{
	//                                             neg: false,
	//                                             abs: {0x15a4e35},
	//                                         },
	//                                         used: constant.used{
	//                                             uses: {
	//                                                 &ir.valueTracker{
	//                                                     orig: &!%v(DEPTH EXCEEDED),
	//                                                     user: &!%v(DEPTH EXCEEDED),
	//                                                 },
	//                                             },
	//                                         },
	//                                     },
	//                                     used: ir.used{
	//                                         uses: {
	//                                             &ir.valueTracker{
	//                                                 orig: &!%v(DEPTH EXCEEDED),
	//                                                 user: &ir.InstAdd{(CYCLIC REFERENCE)},
	//                                             },
	//                                         },
	//                                     },
	//                                 },
	//                                 y:  &constant.Int{
	//                                     typ: &types.IntType{size:32},
	//                                     x:   &big.Int{
	//                                         neg: false,
	//                                         abs: {0x1},
	//                                     },
	//                                     used: constant.used{
	//                                         uses: {
	//                                             &ir.valueTracker{
	//                                                 orig: &!%v(DEPTH EXCEEDED),
	//                                                 user: &ir.InstAdd{(CYCLIC REFERENCE)},
	//                                             },
	//                                         },
	//                                     },
	//                                 },
	//                                 used: ir.used{
	//                                     uses: {
	//                                         &ir.valueTracker{
	//                                             orig: &&!%v(DEPTH EXCEEDED),
	//                                             user: &ir.InstStore{(CYCLIC REFERENCE)},
	//                                         },
	//                                         &ir.valueTracker{
	//                                             orig: &&!%v(DEPTH EXCEEDED),
	//                                             user: &ir.InstCall{(CYCLIC REFERENCE)},
	//                                         },
	//                                     },
	//                                 },
	//                             },
	//                             dst: &ir.Global{(CYCLIC REFERENCE)},
	//                         },
	//                     },
	//                 },
	//             },
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
	//             params: {
	//                 &ir.Param{
	//                     Param: &types.Param{
	//                         name: "x",
	//                         typ:  &types.IntType{size:32},
	//                     },
	//                     used: ir.used{},
	//                 },
	//             },
	//             blocks: nil,
	//             used:   ir.used{
	//                 uses: {
	//                     &ir.namedTracker{
	//                         orig: &&ir.Function{(CYCLIC REFERENCE)},
	//                         user: &ir.InstCall{
	//                             parent: &ir.BasicBlock{
	//                                 parent: &ir.Function{
	//                                     parent: &ir.Module{(CYCLIC REFERENCE)},
	//                                     name:   "rand",
	//                                     typ:    &types.PointerType{
	//                                         elem:  &types.FuncType{(CYCLIC REFERENCE)},
	//                                         space: 0,
	//                                     },
	//                                     sig: &types.FuncType{
	//                                         ret:      &types.IntType{size:32},
	//                                         params:   nil,
	//                                         variadic: false,
	//                                     },
	//                                     params: nil,
	//                                     blocks: {
	//                                         &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                     },
	//                                     used: ir.used{},
	//                                 },
	//                                 name:  "0",
	//                                 insts: {
	//                                     &ir.InstLoad{(CYCLIC REFERENCE)},
	//                                     &ir.InstMul{(CYCLIC REFERENCE)},
	//                                     &ir.InstAdd{(CYCLIC REFERENCE)},
	//                                     &ir.InstStore{(CYCLIC REFERENCE)},
	//                                     &ir.InstCall{(CYCLIC REFERENCE)},
	//                                 },
	//                                 term: &ir.TermRet{
	//                                     parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                                     x:      &ir.InstCall{(CYCLIC REFERENCE)},
	//                                 },
	//                                 used: ir.used{},
	//                             },
	//                             name:   "4",
	//                             callee: &ir.Function{(CYCLIC REFERENCE)},
	//                             sig:    &types.FuncType{(CYCLIC REFERENCE)},
	//                             args:   {
	//                                 &ir.InstAdd{(CYCLIC REFERENCE)},
	//                             },
	//                             used: ir.used{
	//                                 uses: {
	//                                     &ir.valueTracker{
	//                                         orig: &&ir.InstCall{(CYCLIC REFERENCE)},
	//                                         user: &ir.TermRet{(CYCLIC REFERENCE)},
	//                                     },
	//                                 },
	//                             },
	//                         },
	//                     },
	//                 },
	//             },
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
	//             params: nil,
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
	//                             used:   ir.used{
	//                                 uses: {
	//                                     &ir.valueTracker{
	//                                         orig: &&ir.InstLoad{(CYCLIC REFERENCE)},
	//                                         user: &ir.InstMul{(CYCLIC REFERENCE)},
	//                                     },
	//                                 },
	//                             },
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
	//                                 used: constant.used{
	//                                     uses: {
	//                                         &ir.valueTracker{
	//                                             orig: &&!%v(DEPTH EXCEEDED),
	//                                             user: &ir.InstMul{(CYCLIC REFERENCE)},
	//                                         },
	//                                     },
	//                                 },
	//                             },
	//                             used: ir.used{
	//                                 uses: {
	//                                     &ir.valueTracker{
	//                                         orig: &&ir.InstMul{(CYCLIC REFERENCE)},
	//                                         user: &ir.InstAdd{(CYCLIC REFERENCE)},
	//                                     },
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
	//                                 used: constant.used{
	//                                     uses: {
	//                                         &ir.valueTracker{
	//                                             orig: &&!%v(DEPTH EXCEEDED),
	//                                             user: &ir.InstAdd{(CYCLIC REFERENCE)},
	//                                         },
	//                                     },
	//                                 },
	//                             },
	//                             used: ir.used{
	//                                 uses: {
	//                                     &ir.valueTracker{
	//                                         orig: &&ir.InstAdd{(CYCLIC REFERENCE)},
	//                                         user: &ir.InstStore{(CYCLIC REFERENCE)},
	//                                     },
	//                                     &ir.valueTracker{
	//                                         orig: &&ir.InstAdd{(CYCLIC REFERENCE)},
	//                                         user: &ir.InstCall{(CYCLIC REFERENCE)},
	//                                     },
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
	//                             sig:    &types.FuncType{(CYCLIC REFERENCE)},
	//                             args:   {
	//                                 &ir.InstAdd{(CYCLIC REFERENCE)},
	//                             },
	//                             used: ir.used{
	//                                 uses: {
	//                                     &ir.valueTracker{
	//                                         orig: &&ir.InstCall{(CYCLIC REFERENCE)},
	//                                         user: &ir.TermRet{(CYCLIC REFERENCE)},
	//                                     },
	//                                 },
	//                             },
	//                         },
	//                     },
	//                     term: &ir.TermRet{
	//                         parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                         x:      &ir.InstCall{
	//                             parent: &ir.BasicBlock{(CYCLIC REFERENCE)},
	//                             name:   "4",
	//                             callee: &ir.Function{(CYCLIC REFERENCE)},
	//                             sig:    &types.FuncType{(CYCLIC REFERENCE)},
	//                             args:   {
	//                                 &ir.InstAdd{(CYCLIC REFERENCE)},
	//                             },
	//                             used: ir.used{
	//                                 uses: {
	//                                     &ir.valueTracker{
	//                                         orig: &&ir.InstCall{(CYCLIC REFERENCE)},
	//                                         user: &ir.TermRet{(CYCLIC REFERENCE)},
	//                                     },
	//                                 },
	//                             },
	//                         },
	//                     },
	//                     used: ir.used{},
	//                 },
	//             },
	//             used: ir.used{},
	//         },
	//     },
	// }
}
