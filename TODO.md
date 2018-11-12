* figure out how to remove backtracking table in lexer (ref: http://textmapper.org/documentation.html#backtracking-and-invalid-tokens)
* remove stutter in package metadata (e.g. metadata.MetadataNode)
* Consider making []*metadata.MetadataAttachment into a map map[string]*metadata.MDNode as used by instructions, global variables, etc.
* ensure that sumtype interfaces are enforced and implemented.
* swap order of name and typ in NewParam, use NewParam(name, typ) to be consistent with NewFunction, NewGlobal, etc.
* replace use of "enc.Quote([]byte" with "quote".
* check names of fields of instructions against Haskell LLVM library. e.g. name of CleanupPad.Scope. Should it be Parent or From instead of Scope?
* move isFoo and IsFoo to sumtype.go
* rename TypeDef.Alias to TypeDef.LocalName (also rename Alias= to Name= in TypeDef rule of grammar)?
* change `NewArray(typ *types.ArrayType, elems ...Constant) *Array` into `NewArray(elems ...Constant) *Array`, move Typ computation to Type method, and invoke Type method from NewArray. Do the same for remaining constants, where applicable?
