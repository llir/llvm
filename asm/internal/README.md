# Unofficial LLVM IR Specification

The aim of this project is to define a formal grammar for LLVM IR. To achieve this aim, a set of parser generators will be used to evaluate and validate the grammar as it evolves. The goal of each release is to define the grammar required to parse a specific set of LLVM IR example files (contained within the [testdata](testdata/) directory).

## Changes

* Version 0.4 (2016-05-27)
    - Add support for the c4 compiler example (see [#7](https://github.com/llir/spec/issues/7)).
    - Now capable of generating IR for [c4.ll](testdata/c4/c4.ll).

* Version 0.3 (2016-05-26)
    - Generate IR for version 0.2 LLVM IR language constructs (see [#6](https://github.com/llir/spec/issues/6)).
    - Now capable of generating IR for the LLVM IR files in [testdata/uc](testdata/uc).

* Version 0.2 (2016-05-10)
    - Add support for Clang output of µC test cases (see [#5](https://github.com/llir/spec/issues/5)).
    - Now capable of parsing LLVM IR files in [testdata/uc](testdata/uc).

* Version 0.1 (2016-04-18)
    - Initial release of Gocc grammar for LLVM IR.
    - Add support for minimal interesting subset of LLVM IR (see [#1](https://github.com/llir/spec/issues/1)).
    - Now capable of parsing [crt.ll](testdata/crt.ll), [rand.ll](testdata/rand.ll), and [main1.ll](testdata/main1.ll).

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
