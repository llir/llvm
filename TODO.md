* figure out how to remove backtracking table in lexer (ref: http://textmapper.org/documentation.html#backtracking-and-invalid-tokens)
* ensure that sumtype interfaces are enforced and implemented.
* check names of fields of instructions against Haskell LLVM library. e.g. name of CleanupPad.Scope. Should it be Parent or From instead of Scope?
* void call produce value, should not.
	- %0 = call void @f()
* report error in translation of global decl if comdat is used
* rename Def to LLString (or LLVMString) analogous to fmt.GoStringer
* rethink sumtypes to allow for user-defined types; e.g. currently ir.Instruction requires `isInstruction`, but there are valid uses cases where users may wish to define their own instructions to put in basic blocks. One such use case seen in the wild is the pseudo-instruction `type Comment { Data string }` which prints itself as `"; data..."`
* check names of exception handling instruction fields; check Haskell and C++ API.
* add type as first argument to NewArray, NewStruct and NewVector?
