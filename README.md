# llvm

[![Build Status](https://travis-ci.org/llir/llvm.svg?branch=master)](https://travis-ci.org/llir/llvm)
[![Coverage Status](https://coveralls.io/repos/github/llir/llvm/badge.svg?branch=master)](https://coveralls.io/github/llir/llvm?branch=master)
[![GoDoc](https://godoc.org/github.com/llir/llvm?status.svg)](https://godoc.org/github.com/llir/llvm)

Library for interacting with [LLVM IR](http://llvm.org/docs/LangRef.html) in pure Go.

## Installation

```bash
go get -u github.com/llir/llvm/...
```

## Users

* [blessedvirginmary](https://github.com/NateGraff/blessedvirginmary): LLVM IR to Bash transpiler by [@NateGraff](https://github.com/NateGraff).
* [decomp](https://github.com/decomp/decomp): LLVM IR to Go decompiler.
* [tre](https://github.com/zegl/tre): Go to LLVM IR compiler by [@zegl](https://github.com/zegl).
* [uc](https://github.com/mewmew/uc): µC to LLVM IR compiler by [@sangisos](https://github.com/sangisos) and [@mewmew](https://github.com/mewmew).

## Usage

### Parse LLVM IR assembly

[Example usage in GoDoc](https://godoc.org/github.com/llir/llvm/asm#example-package).

```go
// Parse the LLVM IR assembly file `rand.ll`.
m, err := asm.ParseFile("testdata/rand.ll")
if err != nil {
	log.Fatalf("%+v", err)
}
// Pretty-print the data types of the parsed LLVM IR module.
pretty.Println(m)
```

### Output LLVM IR assembly

[Example usage in GoDoc](https://godoc.org/github.com/llir/llvm/ir#example-package).

```go
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
zero := constant.NewInt(i32, 0)
a := constant.NewInt(i32, 0x15A4E35) // multiplier of the PRNG.
c := constant.NewInt(i32, 1)         // increment of the PRNG.

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
```

### Process LLVM IR

[Example usage in GoDoc](https://godoc.org/github.com/llir/llvm/ir#example-package--Evaluator).

The following example program parses [eval.ll](ir/testdata/eval.ll), evaluates the return value of the `@main` function and prints the result to standard output. The result should be 42.

```go
package main

import (
	"fmt"
	"log"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func main() {
	// Parse the LLVM IR assembly file `eval.ll`.
	m, err := asm.ParseFile("testdata/eval.ll")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	// Evalute and print the return value of the `@main` function.
	for _, f := range m.Funcs {
		if f.Name() == "main" {
			e := newEvaluator(f)
			fmt.Println("result:", e.eval())
			break
		}
	}
}

// evaluator is a function evaluator.
type evaluator struct {
	// Function being evaluated.
	f *ir.Function
	// Function arguments.
	args []value.Value
}

// newEvaluator returns a new function evaluator, for evaluating the result of
// invoking f with args.
func newEvaluator(f *ir.Function, args ...value.Value) *evaluator {
	return &evaluator{f: f, args: args}
}

// eval evalutes f and returns the corresponding 32-bit integer.
func (e *evaluator) eval() uint32 {
	f := e.f
	if !types.Equal(f.Sig.RetType, types.I32) {
		panic(fmt.Errorf("support for function return type %s not yet implemented", f.Sig.RetType))
	}
	for _, block := range f.Blocks {
		switch term := block.Term.(type) {
		case *ir.TermRet:
			// Note: support for functions with more than one ret terminator not
			// yet implemented.
			if term.X != nil {
				// The result of the first return value of a function is evaluated.
				return e.evalValue(term.X)
			}
		}
	}
	panic(fmt.Errorf("unable to locate ret terminator in function %q", f.Ident()))
}

// evalInst evaluates inst and returns the corresponding 32-bit integer.
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
		// sign extend.
		if x&0x80000000 != 0 {
			result = signExt(result)
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
			panic(fmt.Errorf("support for callee type %T not yet implemented", inst.Callee))
		}
		ee := newEvaluator(callee, inst.Args...)
		return ee.eval()
	default:
		panic(fmt.Errorf("support for instruction type %T not yet implemented", inst))
	}
}

// evalValue evalutes v and returns the corresponding 32-bit integer.
func (e *evaluator) evalValue(v value.Value) uint32 {
	switch v := v.(type) {
	case ir.Instruction:
		return e.evalInst(v)
	case *constant.Int:
		return uint32(v.X.Int64())
	case *ir.Param:
		f := e.f
		for i, param := range f.Params {
			if v.Ident() == param.Ident() {
				return e.evalValue(e.args[i])
			}
		}
		panic(fmt.Errorf("unable to locate paramater %q of function %q", v.Ident(), f.Ident()))
	default:
		panic(fmt.Errorf("support for value type %T not yet implemented", v))
	}
}

// signExt sign extends x.
func signExt(x uint32) uint32 {
	for i := uint32(31); i >= 0; i-- {
		mask := uint32(1 << i)
		if x&mask != 0 {
			break
		}
		x |= mask
	}
	return x
}
```
