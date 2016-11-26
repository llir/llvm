package ir_test

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func Example() {
	// This example produces LLVM IR code equivalent to the following C code,
	// which implements a pseudo-random number generator.
	//
	//    int abs(int x);
	//
	//    int seed = 0;
	//
	//    // ref: https://en.wikipedia.org/wiki/Linear_congruential_generator
	//    //    a = 0x15A4E35
	//    //    c = 1
	//    int rand(void) {
	//       seed = seed*0x15A4E35 + 1;
	//       return abs(seed);
	//    }

	// Create convenience types and constants.
	i32 := types.I32
	zero := constant.NewInt(0, i32)
	a := constant.NewInt(0x15A4E35, i32) // multiplier of the PRNG.
	c := constant.NewInt(1, i32)         // increment of the PRNG.

	// Create a new module.
	m := ir.NewModule()

	// Create an external function declaration and append it to the module.
	//
	//    int abs(int x);
	abs := m.NewFunction("abs", i32, ir.NewParam("x", i32))

	// Create a global variable definition and append it to the module.
	//
	//    int seed = 0;
	seedptr := m.NewGlobalDef("seed", zero)

	// Create a function definition and append it to the module.
	//
	//    int rand(void) { ... }
	rand := m.NewFunction("rand", i32)

	// Create an unnamed entry basic block and append it to the `rand` function.
	entry := rand.NewBlock("")

	// Create instructions and append them to the entry basic block.
	seed := entry.NewLoad(seedptr)
	tmp1 := entry.NewMul(seed, a)
	tmp2 := entry.NewAdd(tmp1, c)
	entry.NewStore(tmp2, seedptr)
	tmp3 := entry.NewCall(abs, tmp2)
	entry.NewRet(tmp3)

	// Print the LLVM IR assembly of the module.
	fmt.Println(m)

	// Output:
	// @seed = global i32 0
	// declare i32 @abs(i32 %x)
	// define i32 @rand() {
	// ; <label>:0
	// 	%1 = load i32, i32* @seed
	// 	%2 = mul i32 %1, 22695477
	// 	%3 = add i32 %2, 1
	// 	store i32 %3, i32* @seed
	// 	%4 = call i32 @abs(i32 %3)
	// 	ret i32 %4
	// }
}
