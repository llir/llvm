* figure out how to remove backtracking table in lexer (ref: http://textmapper.org/documentation.html#backtracking-and-invalid-tokens)
* remove stutter in package metadata (e.g. metadata.MetadataNode)
* Consider making []*metadata.MetadataAttachment into a map map[string]*metadata.MDNode as used by instructions, global variables, etc.
* ensure that sumtype interfaces are enforced and implemented.

* swap order of name and typ in NewParam, use NewParam(name, typ) to be consistent with NewFunction, NewGlobal, etc.
