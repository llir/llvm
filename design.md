# API design considerations

## Merge ir and ir/constant

### Pro

1. `Func` and `Block` of `BlockAddress` would have correct types (`*ir.Function`
   and `*ir.BasicBlock`, respectively).

   - This is not possible when `ir` and `ir/constant` are split as
     `ir.IndirectBr.Addr` is of type `*constant.BlockAddress`, and
     `constant.BlockAddress.Func` is of type `*ir.Function`, thus creating an
     import cycle.

2. The `IsConstant` dummy method could be unexported.

### Con

1. Name collisions between `ir.Add` and `constant.Add`.

   - Resolved by renaming `constant.Add` to `ir.ExprAdd`.

2. The GoDoc becomes more difficult to navigate.

   - Add a prefix to each type category to ease navigation and grouping. E.g.
     prefix the types of instructions with Inst, terminators with Term, constant
     expressions with Expr and constants with Const.

3. Without renaming types, e.g. `BlockAddress` of package `constant`, the
   constant property would be lost in the name, if simply moving these type
   definitions to package ir; e.g. `ir.BlockAddress`.

   - Resolved in the same fashion as 2), namely by prefixing constants with
     Const.

### Decision

~Keep package `ir` and `ir/constant` separate.~

Merge package `ir/constant` with `ir`.
