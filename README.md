WIP
---

This project is a *work in progress*. The implementation is *incomplete* and
subject to change. The documentation can be inaccurate.

llvm
====

The aim of this project is to provide access to the various LLVM IR
representations; which includes the [LLVM bitcode][] file format, the
[LLVM assembly][] language, and an in-memory representation similar to that of
[go/ast][]. It should be possible to convert between the various
representations; possibly using the interfaces defined in the [encoding][]
package.

[LLVM bitcode]: http://llvm.org/docs/BitCodeFormat.html
[LLVM assembly]: http://llvm.org/docs/LangRef.html
[go/ast]: http://golang.org/pkg/go/ast/
[encoding]: http://golang.org/pkg/encoding/

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
