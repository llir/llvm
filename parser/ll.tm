language llvm(go);

lang = "llvm"
package = "github.com/mewmew/l-tm/parser"

# ### [ Lexical part ] #########################################################

:: lexer

# TODO: fix proper definition of _local_name and _local_id.
_local_name = /[%]foo/
_local_id = /[%]42/

local_ident_tok : /{_local_name}|{_local_id}/

'asm' : /asm/
'datalayout' : /datalayout/
'module' : /module/
'source_filename' : /source_filename/
'target' : /target/
'triple' : /triple/
'type' : /type/

# TODO: remove placeholders.
placeholder1 : /placeholder1/
placeholder2 : /placeholder2/

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

# TODO: figure out where to place these.
StringLit
   : string_lit_tok
;

LocalIdent
   : local_ident_tok
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
	| TargetDefinition
	| ModuleAsm
	| TypeDef
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

# ~~~ [ Target Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#target-triple
# https://llvm.org/docs/LangRef.html#data-layout

# ref: ParseTargetDefinition
#
#   ::= 'target' 'triple' '=' STRINGCONSTANT
#   ::= 'target' 'datalayout' '=' STRINGCONSTANT

TargetDefinition
	: 'target' 'datalayout' '=' StringLit
	| 'target' 'triple' '=' StringLit
;

# ~~~ [ Module-level Inline Assembly ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#module-level-inline-assembly

# ref: ParseModuleAsm
#
#   ::= 'module' 'asm' STRINGCONSTANT

ModuleAsm
	: 'module' 'asm' StringLit
;

# ~~~ [ Type Defintion ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#structure-type

# ref: ParseUnnamedType
#
#   ::= LocalVarID '=' 'type' type

# ref: ParseNamedType
#
#   ::= LocalVar '=' 'type' type

TypeDef
	: LocalIdent '=' 'type' OpaqueType
	| LocalIdent '=' 'type' Type
;

# TODO: fix placeholders.

OpaqueType
   : placeholder1
;

Type
   : placeholder2
;
