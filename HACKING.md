# Hacking of llir/llvm

The purpose of this file is to describe how to get started working on the `llir/llvm` code base. It is intended for people who want more insight into the inner workings of the `llir/llvm`, and who may want to contribute to the project, or change the inner workings to suit their own specific needs.

## Installation

Step one is to install the required tools and dependencies of the project.

### llir/ll

These steps describe how to install the tools required to generate the [llir/ll](https://github.com/llir/ll) LLVM IR parser from an [EBNF grammar](https://github.com/llir/grammar/blob/master/ll.tm).

```bash
# Clone repo and submodules.
git clone --recursive https://github.com/llir/ll

# Install textmapper.
cd ll/tools/textmapper/tm-tool
./gen.sh

# Generate LLVM IR parser.
cd ../../..
make
```

### llir/llvm

These steps describe how to install the [llir/llvm](https://github.com/llir/llvm) LLVM IR library and how to download the associated test data.

```bash
# Clone repo and submodules.
git clone --recursive https://github.com/llir/llvm
cd llvm

# Re-generate asm/enum package when making changes to ir/enum.
make -C asm/enum

# Build.
go install ./...

# Run tests.
go test ./...
```

## Directory structure

This section will give a brief introduction to what each directory contains, so that you may know what parts of the code to look closer at or modify.

* `asm`: package responsible for parsing LLVM IR assembly into the data structures defined in `llir/llvm/ir`. This package uses the `llir/llvm/ll` parser under the hood, and is mainly responsible for translating the [Textmapper](https://github.com/inspirer/textmapper) generated AST data types into equivalent IR data types. For instance, it performs type resolution (with support for recursive type definitions), identifier resolution (e.g. the occurrences of an identifier `@foo` are mapped to their associated global value [*ir.Global](https://pkg.go.dev/github.com/llir/llvm/ir#Global)), etc.
   - `asm/enum`: simple Go package containing enumerated definitions. This package mirrors the definitions of `ir/enum` and is automatically generated (see the associated [Makefile](https://github.com/llir/llvm/blob/master/asm/enum/Makefile)).
* `cmd/l-tm`: simple example tool used to profile CPU and memory usage of the LLVM IR parser. (*Note*, this tool is likely to be removed in future releases of `llir/llvm`.)
* `internal/enc`: internal package dealing with encoding/decoding of LLVM IR identifiers (e.g. global identifier `foo` is encoded as `@foo`). Used by both `llir/llvm/asm` and `llir/llvm/ir`.
* `ir`: top-level LLVM IR package, defines the intermediate representation of modules, functions, global variables and other key concepts of LLVM IR.
   - `ir/constant`: implements LLVM IR constants, which act as immutable values.
   - `ir/enum`: simple Go package containing enumerated definitions. This package exists mainly to not proliferate the number of definitions in the top-level `llir/llvm/ir` package.
   - `ir/metadata`: defines the metadata types of LLVM IR, including DWARF debug information.
   - `ir/types`: defines the data types of LLVM IR (e.g. `i32`, `double`, etc).
   - `ir/value`: provides a Go interface definition of LLVM IR values, a core concept in the `llir/llvm/ir` API.
* `testdata`: submodule of https://github.com/llir/testdata containing test data from the official LLVM project and from Coreutils and SQLite.
