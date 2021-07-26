# Release history

## Version 0.3.4 (2021-07-26)

Update `llir/llvm` to include new language concepts of LLVM 12.0 (thanks [@dannypsnl](https://github.com/dannypsnl)).

[Full list of changes since v0.3.3](https://github.com/llir/llvm/compare/v0.3.3...v0.3.4).

## Version 0.3.3 (2021-03-25)

Update `llir/llvm` to include new language concepts of LLVM 11.0 (thanks [@dannypsnl](https://github.com/dannypsnl)).

[Full list of changes since v0.3.2](https://github.com/llir/llvm/compare/v0.3.2...v0.3.3).

## Version 0.3.2 (2020-06-09)

Re-add `ir.NewLocalIdent` which was unintentionally removed in v0.3.1.

[Full list of changes since v0.3.1](https://github.com/llir/llvm/compare/v0.3.1...v0.3.2).

## Version 0.3.1 (2020-05-04)

Update `llir/llvm` to include new language concepts of LLVM 10.0 (thanks [@dannypsnl](https://github.com/dannypsnl)).

[Full list of changes since v0.3.0](https://github.com/llir/llvm/compare/v0.3.0...v0.3.1).

## Version 0.3 (2019-12-29)

Primary focus of version 0.3: *grammar covering the entire LLVM IR language*.

The grammar for LLVM IR is now complete and covers the entire LLVM IR language (as of LLVM 9.0).

Lexers and parsers for LLVM IR assembly are automatically generated from an [EBNF grammar](https://github.com/llir/grammar/blob/master/ll.tm) using [Textmapper](https://github.com/inspirer/textmapper).

The Textmapper generated source code has been split into a [dedicated repository](https://github.com/llir/ll) to better handle repository size.

## Version 0.2 (2017-06-24)

Primary focus of version 0.2: *read and write support of LLVM IR assembly*.

Lexers and parsers for LLVM IR assembly are automatically generated from a [BNF grammar](https://github.com/llir/llvm/blob/28149269dab73cc63915a9c2c6c7b25dbd4db027/asm/internal/ll.bnf) using [Gocc](https://github.com/goccmack/gocc).

A high-level API for parsing LLVM IR assembly is provided by [llvm/asm](https://godoc.org/github.com/llir/llvm/asm).

The [llvm/ir](https://godoc.org/github.com/llir/llvm/ir) package supports all instructions of LLVM IR, except the instructions used for concurrency and exception handling.

The llir/llvm packages are now go-getable, as the Gocc generated source code has been added to the source tree.

## Version 0.1 (2015-04-19)

Initial release.

Preliminary work on the `llvm/ir` package which provides an in-memory representation of LLVM IR in pure Go.

Hand-written lexer and preliminary work on a recursive descent parser for LLVM IR assembly.
