* figure out how to remove backtracking table in lexer (ref: http://textmapper.org/documentation.html#backtracking-and-invalid-tokens)
* remove stutter in package metadata (e.g. metadata.MetadataNode)
* ensure that sumtype interfaces are enforced and implemented.
* check names of fields of instructions against Haskell LLVM library. e.g. name of CleanupPad.Scope. Should it be Parent or From instead of Scope?
* rename TypeDef.Alias to TypeDef.LocalName (also rename Alias= to Name= in TypeDef rule of grammar)?
* change `NewArray(typ *types.ArrayType, elems ...Constant) *Array` into `NewArray(elems ...Constant) *Array`, move Typ computation to Type method, and invoke Type method from NewArray. Do the same for remaining constants, where applicable?
