package ir_test

import (
	"fmt"
	"log"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
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

	// Create a new LLVM IR module.
	m := ir.NewModule()

	// Create an external function declaration and append it to the module.
	//
	//    int abs(int x);
	abs := m.NewFunction("abs", i32, ir.NewParam("x", i32))

	// Create a global variable definition and append it to the module.
	//
	//    int seed = 0;
	seed := m.NewGlobalDef("seed", zero)

	// Create a function definition and append it to the module.
	//
	//    int rand(void) { ... }
	rand := m.NewFunction("rand", i32)

	// Create an unnamed entry basic block and append it to the `rand` function.
	entry := rand.NewBlock("")

	// Create instructions and append them to the entry basic block.
	tmp1 := entry.NewLoad(seed)
	tmp2 := entry.NewMul(tmp1, a)
	tmp3 := entry.NewAdd(tmp2, c)
	entry.NewStore(tmp3, seed)
	tmp4 := entry.NewCall(abs, tmp3)
	entry.NewRet(tmp4)

	// Print the LLVM IR assembly of the module.
	fmt.Println(m)

	// Output:
	//
	// @seed = global i32 0
	//
	// declare i32 @abs(i32 %x)
	//
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

func Example_evaluator() {
	// Parse the LLVM IR assembly file `eval.ll`.
	m, err := asm.ParseFile("testdata/eval.ll")
	if err != nil {
		log.Fatal(err)
	}
	// Evalute and print the return value of the `main` function.
	for _, f := range m.Funcs {
		if f.Name == "main" {
			e := newEvaluator(f)
			fmt.Println("result:", e.eval())
			break
		}
	}

	// Output:
	//
	// result: 42
}

// evaluator is a function evaluator.
type evaluator struct {
	// Function.
	f *ir.Function
	// Function arguments.
	args []value.Value
}

// newEvaluator returns a new function evaluator, for evaluating the result of
// invoking f with args.
func newEvaluator(f *ir.Function, args ...value.Value) *evaluator {
	return &evaluator{f: f, args: args}
}

// eval evalutes f and returns the corresponding 64-bit integer.
func (e *evaluator) eval() uint32 {
	f := e.f
	if !types.Equal(f.Sig.Ret, types.I32) {
		panic(fmt.Errorf("support for function return type %v not yet implemented", f.Sig.Ret))
	}
	for _, block := range f.Blocks {
		switch term := block.Term.(type) {
		case *ir.TermRet:
			// NOTE: support for functions with more than one RET terminator not
			// yet implemented.
			if term.X != nil {
				// Evaluate the result of the first return value of a function is
				// evaluated.
				return e.evalValue(term.X)
			}
		}
	}
	panic(fmt.Errorf("unable to locate RET terminator in function %q", f.Name))
}

// evalInst evaluates inst and returns the corresponding 64-bit integer.
func (e *evaluator) evalInst(inst ir.Instruction) uint32 {
	switch inst := inst.(type) {
	// Binary instructions.
	case *ir.InstAdd:
		return e.evalValue(inst.X) + e.evalValue(inst.Y)
	case *ir.InstSub:
		return e.evalValue(inst.X) - e.evalValue(inst.Y)
	case *ir.InstMul:
		return e.evalValue(inst.X) * e.evalValue(inst.Y)
	case *ir.InstUDiv:
		return e.evalValue(inst.X) / e.evalValue(inst.Y)
	case *ir.InstSDiv:
		return e.evalValue(inst.X) / e.evalValue(inst.Y)
	case *ir.InstURem:
		return e.evalValue(inst.X) % e.evalValue(inst.Y)
	case *ir.InstSRem:
		return e.evalValue(inst.X) % e.evalValue(inst.Y)
	// Bitwise instructions.
	case *ir.InstShl:
		return e.evalValue(inst.X) << e.evalValue(inst.Y)
	case *ir.InstLShr:
		return e.evalValue(inst.X) >> e.evalValue(inst.Y)
	case *ir.InstAShr:
		x, y := e.evalValue(inst.X), e.evalValue(inst.Y)
		result := x >> y
		if x&0x80000000 != 0 {
			for i := uint32(31); i >= 0; i-- {
				mask := uint32(1 << i)
				if result&mask != 0 {
					break
				}
				result |= mask
			}
		}
		return result
	case *ir.InstAnd:
		return e.evalValue(inst.X) & e.evalValue(inst.Y)
	case *ir.InstOr:
		return e.evalValue(inst.X) | e.evalValue(inst.Y)
	case *ir.InstXor:
		return e.evalValue(inst.X) ^ e.evalValue(inst.Y)
	// Other instructions.
	case *ir.InstCall:
		callee, ok := inst.Callee.(*ir.Function)
		if !ok {
			panic(fmt.Errorf("support for callee of type %T not yet implemented", inst.Callee))
		}
		ee := newEvaluator(callee, inst.Args...)
		return ee.eval()
	default:
		panic(fmt.Errorf("support for instruction type %T not yet implemented", inst))
	}
}

// evalValue evalutes v and returns the corresponding 64-bit integer.
func (e *evaluator) evalValue(v value.Value) uint32 {
	switch v := v.(type) {
	case ir.Instruction:
		return e.evalInst(v)
	case *constant.Int:
		return uint32(v.X.Int64())
	case *types.Param:
		if len(v.Name) == 0 {
			panic("support for unnamed parameters not yet implemented")
		}
		f := e.f
		for i, param := range f.Sig.Params {
			if v.Name == param.Name {
				return e.evalValue(e.args[i])
			}
		}
		panic(fmt.Errorf("unable to locate paramater %q of function %q", v.Name, f.Name))
	default:
		panic(fmt.Errorf("support for value type %T not yet implemented", v))
	}
}
