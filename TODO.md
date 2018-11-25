* figure out how to remove backtracking table in lexer (ref: http://textmapper.org/documentation.html#backtracking-and-invalid-tokens)
* remove stutter in package metadata (e.g. metadata.MetadataNode)
* ensure that sumtype interfaces are enforced and implemented.
* check names of fields of instructions against Haskell LLVM library. e.g. name of CleanupPad.Scope. Should it be Parent or From instead of Scope?
* void call produce value, should not.
	- %0 = call void @f()
* Rename NewFunction to NewFunc?

# Panics in test cases

## Analysis

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

- [ ] Analysis/CostModel/SystemZ/intrinsic-cost-crash.ll

syntax error at line 29

- [ ] Analysis/DivergenceAnalysis/AMDGPU/phi-undef.ll

syntax error at line 16

- [ ] Analysis/DependenceAnalysis/MIVCheckConst.ll

syntax error at line 32

- [ ] Analysis/DependenceAnalysis/NonAffineExpr.ll

syntax error at line 10

- [ ] Analysis/DominanceFrontier/new_pm_test.ll

invalid local ID in function "@a_linear_impl_fig_1", expected %12, got %13

LLVM differentiates between named an unnamed IDs, e.g.
`%42` and `; <label>:42` vs, `%"42"` and `42:`.

Proposed solution:

Have each instruction, function parameter and basic block embed LocalIdent (instead of LocalName), which is defined as follows.

```go
type LocalIdent struct {
	Ident string
	ID int64
}

func (i *LocalIdent) Ident() string {
	ident := i.Ident
	if i.IsUnnamed() {
		ident = strconv.Itoa(i.ID)
	}
	return enc.LocalIdent(ident)
}

func (i *LocalIdent) Name() string {
	return i.Ident
}

func (i *LocalIdent) SetName(name string) {
	i.Ident = names
}

func (i *LocalIdent) ID() int64 {
	return i.ID // TODO: figure out how to resolve name collition between method and field, perhaps have ID be unexported?
}

func (i *LocalIdent) SetID(id int64) {
	i.ID = id
}
```

- [ ] Analysis/MemorySSA/cyclicphi.ll

syntax error at line 7

- [ ] Analysis/RegionInfo/cond_loop.ll

invalid local ID in function "@normal_condition", expected %0, got %5

- [ ] Analysis/RegionInfo/infinite_loop_2.ll

invalid local ID in function "@normal_condition", expected %3, got %5

- [ ] Analysis/RegionInfo/infinite_loop_3.ll

invalid local ID in function "@normal_condition", expected %1, got %7

- [ ] Analysis/RegionInfo/infinite_loop_4.ll

invalid local ID in function "@normal_condition", expected %1, got %7

- [ ] Analysis/RegionInfo/infinite_loop_5_a.ll

invalid local ID in function "@normal_condition", expected %1, got %7

- [ ] Analysis/RegionInfo/infinite_loop_5_b.ll

invalid local ID in function "@normal_condition", expected %1, got %7

- [ ] Analysis/RegionInfo/infinite_loop_5_c.ll

invalid local ID in function "@normal_condition", expected %1, got %7

- [ ] Analysis/RegionInfo/loop_with_condition.ll

invalid local ID in function "@normal_condition", expected %6, got %8

- [ ] Analysis/RegionInfo/mix_1.ll

invalid local ID in function "@a_linear_impl_fig_1", expected %8, got %15

- [ ] Analysis/RegionInfo/paper.ll

invalid local ID in function "@a_linear_impl_fig_1", expected %12, got %13

- [ ] Analysis/ScalarEvolution/2007-11-18-OrInstruction.ll

syntax error at line 7

- [ ] Analysis/ScalarEvolution/2008-02-11-ReversedCondition.ll

syntax error at line 8

- [ ] Analysis/ScalarEvolution/2008-02-15-UMax.ll

syntax error at line 9

- [ ] Analysis/ScalarEvolution/2011-03-09-ExactNoMaxBECount.ll

invalid local ID in function "@bar", expected %0, got %4

- [ ] Analysis/ScalarEvolution/implied-via-division.ll

syntax error at line 15

- [ ] Analysis/ScalarEvolution/pr22674.ll

syntax error at line 14

- [ ] Analysis/TypeBasedAliasAnalysis/PR17620.ll

syntax error at line 16

- [ ] Analysis/ValueTracking/func-ptr-lsb.ll

syntax error at line 15

- [ ] Analysis/ValueTracking/memory-dereferenceable.ll

support for return attribute Dereferenceable not yet implemented

## Assembler

- [ ] Assembler/2003-11-11-ImplicitRename.ll

syntax error at line 3

# Enable test cases

- [ ] Analysis/BasicAA/pr18573.ll

NaN

- [ ] Assembler/2003-11-24-SymbolTableCrash.ll

error reported correctly. check error in test case.

- [ ] Assembler/2004-03-30-UnclosedFunctionCrash.ll

error reported correctly (syntax error). check error in test case
