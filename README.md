## WIP

This project is a *work in progress*. The implementation is *incomplete* and subject to change. The documentation may be inaccurate.

# llvm

[![Build Status](https://travis-ci.org/llir/llvm.svg?branch=master)](https://travis-ci.org/llir/llvm)
[![Coverage Status](https://img.shields.io/coveralls/llir/llvm.svg)](https://coveralls.io/r/llir/llvm?branch=master)
[![GoDoc](https://godoc.org/github.com/llir/llvm?status.svg)](https://godoc.org/github.com/llir/llvm)

The aim of this project is to provide access to the various LLVM IR representations; which includes the [LLVM bitcode] file format, the [LLVM assembly] language, and an in-memory representation similar to that of [go/ssa]. It should be possible to convert between the various representations; possibly using the interfaces defined in the [encoding] package.

[LLVM bitcode]: http://llvm.org/docs/BitCodeFormat.html
[LLVM assembly]: http://llvm.org/docs/LangRef.html
[go/ssa]: https://godoc.org/code.google.com/p/go.tools/go/ssa
[encoding]: http://golang.org/pkg/encoding/

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
