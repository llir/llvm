* figure out how to remove backtracking table in lexer (ref: http://textmapper.org/documentation.html#backtracking-and-invalid-tokens)
* remove stutter in package metadata (e.g. metadata.MetadataNode)
* ensure that sumtype interfaces are enforced and implemented.
* check names of fields of instructions against Haskell LLVM library. e.g. name of CleanupPad.Scope. Should it be Parent or From instead of Scope?
* rename TypeDef.Alias to TypeDef.LocalName (also rename Alias= to Name= in TypeDef rule of grammar)?

# Panics in test cases

- [ ] Analysis/BlockFrequencyInfo/loop_with_branch.ll

syntax error at line 17

- [ ] Analysis/ConstantFolding/vectorgep-crash.ll

```
panic: invalid index type for structure element; expected *ast.IntConst, got ast.TypeValue

goroutine 1 [running]:
github.com/llir/llvm/asm.(*generator).gepType(0xc000001b00, 0x6b5bc0, 0xc000123620, 0xc0001250c0, 0x3, 0x3, 0x0, 0xc000123380, 0x0, 0x0)
	/home/u/Desktop/go/src/github.com/llir/llvm/asm/local.go:662 +0x971
github.com/llir/llvm/asm.(*funcGen).newIRValueInst(0xc0000cf898, 0xc00010b10a, 0x9, 0x6b0040, 0xc000120c48, 0xc000120c40, 0x0, 0xc0000cf4b0, 0x43c2fc)
```

- [ ] Analysis/CostModel/AMDGPU/fdiv.ll

```
panic: support for floating-point kind half not yet implemented

goroutine 1 [running]:
github.com/llir/llvm/ir/constant.NewFloatFromString(0xc0003d7520, 0xc000013e17, 0x3, 0x203000, 0x203000, 0xc0000d30d8)
	/home/u/Desktop/go/src/github.com/llir/llvm/ir/constant/const_float.go:187 +0x113f
github.com/llir/llvm/asm.(*generator).irFloatConst(0xc000001380, 0x6b5c00, 0xc0003d7520, 0xc0000114e8, 0x4084eb, 0xc00001e000, 0xc0000d3140)
	/home/u/Desktop/go/src/github.com/llir/llvm/asm/const.go:101 +0x13b
```

# Enable test cases

- [ ] Analysis/BasicAA/pr18573.ll

NaN
