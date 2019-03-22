package ir_test

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func Example_main() {
	// This example produces LLVM IR code equivalent to the following C code:
	//
	//    int main() {
	//      int a = 32;
	//      int b = 16;
	//      return a + b;
	//    }
	//
	// Read: https://blog.felixangell.com/an-introduction-to-llvm-in-go for inspiration
	//
	// Output:
	//define i32 @main() {
	//; <label>:0
	//       %1 = alloca i32, align 4
	//       %2 = alloca i32, align 4
	//       %3 = alloca i32, align 4
	//       store i32 0, i32* %1, align 4
	//       store i32 32, i32* %2, align 4
	//       store i32 16, i32* %3, align 4
	//       %4 = load i32, i32* %2, align 4
	//       %5 = load i32, i32* %3, align 4
	//       %6 = add i32 %4, %5
	//       ret i32 %6
	//}

	i32 := types.I32

	// Create a new LLVM IR module.
	mod := ir.NewModule()
	//    int main() { ... }
	main := mod.NewFunc("main", i32)

	// Create an unnamed entry basic block and append it to the `main` function.
	entry := main.NewBlock("")
	// Create instructions and append them to the entry basic

	// %1 = alloca i32, align 4
	zero := entry.NewAlloca(i32)
	zero.Align = 4

	// %a = alloca i32, align 4
	a := entry.NewAlloca(i32)
	a.Align = 4

	// %b = alloca i32, align 4
	b := entry.NewAlloca(i32)
	b.Align = 4

	// store i32 0, i32* %1
	s0 := entry.NewStore(constant.NewInt(i32, 0), zero)
	s0.Align = 4

	// store i32 32, i32* %a, align 4
	s1 := entry.NewStore(constant.NewInt(i32, 32), a)
	s1.Align = 4

	// store i32 16, i32* %b, align 4
	s2 := entry.NewStore(constant.NewInt(i32, 16), b)
	s2.Align = 4

	// %2 = load i32, i32* %a, align 4
	tmpA := entry.NewLoad(a)
	tmpA.Align = 4

	// %3 = load i32, i32* %b, align 4
	tmpB := entry.NewLoad(b)
	tmpB.Align = 4

	// %4 = add nsw i32 %2, %3
	tmpC := entry.NewAdd(tmpA, tmpB)

	entry.NewRet(tmpC)

	// Print the LLVM IR assembly of the module.
	fmt.Println(mod)
}
