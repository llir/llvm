# How to update LLVM

This document aims to describe the approach taken to update `llir/llvm` to support new versions of the official LLVM release, as in, how to update from LLVM 8.0 to LLVM 9.0.

For now, we collect links to issues, PR and comments which describe the actions taken, and given example commits for adding new enums, updating the grammar, translating the AST to IR, updating the test cases, etc.

## Test cases

To ensure `llir/llvm` is following the LLVM specification, we rely on the test cases of the official LLVM distribution. To update these test cases when a new version of LLVM is released, do as described in [this comment of issue #105](https://github.com/llir/llvm/issues/105#issuecomment-548619916).
