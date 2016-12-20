//+build ignore

// Fix dummy values as follows.
//
// Per module.
//
//    1. Index type definitions.
//    2. Index global variables.
//    3. Index functions.
//    4. Fix body of named types.
//       - i.e. look up and set Def for each *types.NamedType.
//    6. Replace dummy instructions containing dummy Type method
//       implementations; e.g. *dummy.InstGetElementPtr.
//
// Per function.
//
//    1. Force generate local IDs for unnamed basic blocks and instructions.
//    2. Index basic blocks.
//    3. Index function parameters.
//    4. Fix dummy instructions and terminators.
//       - e.g. replace *dummy.InstCall with *ir.InstCall
//       - Replace function and label names with *ir.Function and *ir.BasicBlock
//         values.
//       - Leave dummy operands (e.g. *dummy.Local, *dummy.Global,
//         *dummy.InstPhi and *dummy.InstCall) as these will be replaced
//         in a later stage.
//    5. Index local variables produced by instructions.
//    6. Replace dummy operands of instructions and terminators.

package irx
