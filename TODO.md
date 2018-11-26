* figure out how to remove backtracking table in lexer (ref: http://textmapper.org/documentation.html#backtracking-and-invalid-tokens)
* remove stutter in package metadata (e.g. metadata.MetadataNode)
* ensure that sumtype interfaces are enforced and implemented.
* check names of fields of instructions against Haskell LLVM library. e.g. name of CleanupPad.Scope. Should it be Parent or From instead of Scope?
* void call produce value, should not.
	- %0 = call void @f()
* report error in translation of global decl if comdat is used
