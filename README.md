## WIP

This project is a *work in progress*. The implementation is *incomplete* and subject to change. The documentation may be inaccurate.

# llvm

[![Build Status](https://travis-ci.org/llir/llvm.svg?branch=master)](https://travis-ci.org/llir/llvm)
[![Coverage Status](https://coveralls.io/repos/github/llir/llvm/badge.svg?branch=master)](https://coveralls.io/github/llir/llvm?branch=master)
[![GoDoc](https://godoc.org/github.com/llir/llvm?status.svg)](https://godoc.org/github.com/llir/llvm)

The aim of this project is to provide a pure Go library for interacting with [LLVM IR](http://llvm.org/docs/LangRef.html).

## Status

Updated: 2016-12-02

- [x] Write support of LLVM IR assembly files.
    - [Example usage in GoDoc](https://godoc.org/github.com/llir/llvm/ir#example-package).
- [x] Preliminary read support of LLVM IR assmebly files.
    - A lexer and parser for LLVM IR assembly is generated from a [BNF grammar](https://github.com/llir/llvm/blob/master/asm/internal/ll.bnf) using [gocc](https://github.com/goccmack/gocc).

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
