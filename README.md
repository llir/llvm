# llvm

[![Build Status](https://github.com/llir/llvm/workflows/Go/badge.svg?branch=master)](https://github.com/llir/llvm/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/llir/llvm/badge.svg?branch=master)](https://coveralls.io/github/llir/llvm?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/llir/llvm)](https://goreportcard.com/report/github.com/llir/llvm)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/llir/llvm)

Library for interacting with [LLVM IR](http://llvm.org/docs/LangRef.html) in pure Go.

## Introduction

* [Introductory blog post "LLVM IR and Go"](https://blog.gopheracademy.com/advent-2018/llvm-ir-and-go/)
* [Our Document](https://llir.github.io/document/)

## Installation

```bash
go get -u github.com/llir/llvm/...
```

## Versions

Map between `llir/llvm` tagged releases and LLVM release versions.

* [llir/llvm v0.3.2](https://github.com/llir/llvm/tree/v0.3.2): LLVM 10.0
* [llir/llvm v0.3.0](https://github.com/llir/llvm/tree/v0.3.0): LLVM 9.0

## Users

* [decomp](https://github.com/decomp/decomp): LLVM IR to Go decompiler by [@decomp](https://github.com/decomp).
* [geode](https://github.com/geode-lang/geode): Geode to LLVM IR compiler by [@nickwanninger](https://github.com/nickwanninger).
* [leaven](https://github.com/andybalholm/leaven): LLVM IR to Go decompiler by [@andybalholm](https://github.com/andybalholm).
* [slate](https://github.com/nektro/slate): Slate to LLVM IR compiler by [@nektro](https://github.com/nektro).
* [tre](https://github.com/zegl/tre): Go to LLVM IR compiler by [@zegl](https://github.com/zegl).
* [uc](https://github.com/mewmew/uc): µC to LLVM IR compiler by [@sangisos](https://github.com/sangisos) and [@mewmew](https://github.com/mewmew).

## Usage

### Input example - Parse LLVM IR assembly

[Example usage in GoDoc](https://godoc.org/github.com/llir/llvm/asm#example-package).

```go
// This example parses an LLVM IR assembly file and pretty-prints the data types
// of the parsed module to standard output.
package main

import (
	"log"

	"github.com/kr/pretty"
	"github.com/llir/llvm/asm"
)

func main() {
	// Parse the LLVM IR assembly file `foo.ll`.
	m, err := asm.ParseFile("foo.ll")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	// Pretty-print the data types of the parsed LLVM IR module.
	pretty.Println(m)
}
```

### Output example - Produce LLVM IR assembly

[Example usage in GoDoc](https://godoc.org/github.com/llir/llvm/ir#example-package).

```go
// This example produces LLVM IR code equivalent to the following C code, which
// implements a pseudo-random number generator.
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
package main

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func main() {
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
	abs := m.NewFunc("abs", i32, ir.NewParam("x", i32))

	// Create a global variable definition and append it to the module.
	//
	//    int seed = 0;
	seed := m.NewGlobalDef("seed", zero)

	// Create a function definition and append it to the module.
	//
	//    int rand(void) { ... }
	rand := m.NewFunc("rand", i32)

	// Create an unnamed entry basic block and append it to the `rand` function.
	entry := rand.NewBlock("")

	// Create instructions and append them to the entry basic block.
	tmp1 := entry.NewLoad(i32, seed)
	tmp2 := entry.NewMul(tmp1, a)
	tmp3 := entry.NewAdd(tmp2, c)
	entry.NewStore(tmp3, seed)
	tmp4 := entry.NewCall(abs, tmp3)
	entry.NewRet(tmp4)

	// Print the LLVM IR assembly of the module.
	fmt.Println(m)
}
```

### Analysis example - Process LLVM IR

[Example usage in GoDoc](https://godoc.org/github.com/llir/llvm/ir#example-package--Callgraph).

```go
// This example program analyses an LLVM IR module to produce a callgraph in
// Graphviz DOT format.
package main

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
)

func main() {
	// Parse LLVM IR assembly file.
	m, err := asm.ParseFile("foo.ll")
	if err != nil {
		panic(err)
	}
	// Produce callgraph of module.
	callgraph := genCallgraph(m)
	// Output callgraph in Graphviz DOT format.
	fmt.Println(callgraph)
}

// genCallgraph returns the callgraph in Graphviz DOT format of the given LLVM
// IR module.
func genCallgraph(m *ir.Module) string {
	buf := &strings.Builder{}
	buf.WriteString("digraph {\n")
	// For each function of the module.
	for _, f := range m.Funcs {
		// Add caller node.
		caller := f.Ident()
		fmt.Fprintf(buf, "\t%q\n", caller)
		// For each basic block of the function.
		for _, block := range f.Blocks {
			// For each non-branching instruction of the basic block.
			for _, inst := range block.Insts {
				// Type switch on instruction to find call instructions.
				switch inst := inst.(type) {
				case *ir.InstCall:
					callee := inst.Callee.Ident()
					// Add edges from caller to callee.
					fmt.Fprintf(buf, "\t%q -> %q\n", caller, callee)
				}
			}
			// Terminator of basic block.
			switch term := block.Term.(type) {
			case *ir.TermRet:
				// do something.
				_ = term
			}
		}
	}
	buf.WriteString("}")
	return buf.String()
}
```

## License

The `llir/llvm` project is dual-licensed to the [public domain](UNLICENSE) and under a [zero-clause BSD license](LICENSE). You may choose either license to govern your use of `llir/llvm`.
