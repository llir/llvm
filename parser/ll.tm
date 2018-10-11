language llvm(go);

lang = "llvm"
package = "github.com/mewmew/l-tm/parser"

# Lexer

:: lexer

'source_filename' : /source_filename/

'=' : /=/

# TODO: figure out how to handle string_lit_tok correctly.
string_lit_tok : /"[^"]"/

# ### [ Syntax part ] ##########################################################

# The LLVM IR grammar has been based on the source code of the official LLVM
# project, as of 2018-02-19 (rev db070bbdacd303ae7da129f59beaf35024d94c53).
#
#    * lib/AsmParser/LLParser.cpp

:: parser

input : Module;

# TODO: figure out where to place StringLit.
StringLit
   : string_lit_tok
;

# === [ Module ] ===============================================================

# https://llvm.org/docs/LangRef.html#module-structure

# ref: Run
#
#   module ::= toplevelentity*

Module
	: TopLevelEntity*
;

# --- [ Top-level Entities ] ---------------------------------------------------

# ref: ParseTopLevelEntities

TopLevelEntity
	: SourceFilename
	#| TargetDefinition
	#| ModuleAsm
	#| TypeDef
	#| ComdatDef
	#| GlobalDecl
	#| GlobalDef
	#| IndirectSymbolDef
	#| FunctionDecl
	#| FunctionDef
	#| AttrGroupDef
	#| NamedMetadataDef
	#| MetadataDef
	#| UseListOrder
	#| UseListOrderBB
;

# ~~~ [ Source Filename ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#source-filename

# ref: ParseSourceFileName
#
#   ::= 'source_filename' '=' STRINGCONSTANT

SourceFilename
	: 'source_filename' '=' StringLit
;
