1. figure out how to remove backtracking table in lexer (ref: http://textmapper.org/documentation.html#backtracking-and-invalid-tokens)
2. check names of fields of instructions against Haskell LLVM library. e.g. name of CleanupPad.Scope. Should it be Parent or From instead of Scope?
	- check names of exception handling instruction fields; check Haskell and C++ API.
3. report error in translation of global decl if comdat is used
4. void call produce value, should not.
	- %0 = call void @f()
	- check with @pwaller how to reproduce.
