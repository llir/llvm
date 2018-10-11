language llvm(go);

lang = "llvm"
package = "github.com/mewmew/l-tm/parser"

# TODO: check when to use Fooopt and when to use Foo? (as based on the AST
# they produce)

# ### [ Lexical part ] #########################################################

:: lexer

# TODO: fix def of float_lit_tok

float_lit_tok : /3\.14/

# TODO: move to the right place.

int_type_tok : /i[0-9]+/

# TODO: fix def of int_lit_tok

#   Integer           [-]?[0-9]+

int_lit_tok : /[-]?[0-9]+/

# TODO: fix proper definition of _escape_name, _name and _id.
_escape_name = /bar/
_name = /foo/
_id = /42/

# TODO: add {_quoted_name|_name} to places where {_name} is used.

global_ident_tok : /{_global_name}|{_global_id}/

_global_name = /[@]{_name}/
_global_id = /[@]{_id}/

local_ident_tok : /{_local_name}|{_local_id}/

_local_name = /[%]{_name}/
_local_id = /[%]{_id}/

#   Label             [-a-zA-Z$._0-9]+:

# TODO: add _quoted_string version of label_ident_tok

label_ident_tok : /[-a-zA-Z$\._0-9]+:/

attr_group_id_tok : /[#]{_id}/

comdat_name_tok : /[$]{_name}/

metadata_name_tok : /[!]{_escape_name}/

metadata_id_tok : /[!]{_id}/

'add' : /add/
'addrspace' : /addrspace/
'addrspacecast' : /addrspacecast/
'alias' : /alias/
'align' : /align/
'alignstack' : /alignstack/
'allocsize' : /allocsize/
'alwaysinline' : /alwaysinline/
'amdgpu_cs' : /amdgpu_cs/
'amdgpu_es' : /amdgpu_es/
'amdgpu_gs' : /amdgpu_gs/
'amdgpu_hs' : /amdgpu_hs/
'amdgpu_kernel' : /amdgpu_kernel/
'amdgpu_ls' : /amdgpu_ls/
'amdgpu_ps' : /amdgpu_ps/
'amdgpu_vs' : /amdgpu_vs/
'and' : /and/
'any' : /any/
'anyregcc' : /anyregcc/
'appending' : /appending/
'argmemonly' : /argmemonly/
'arm_aapcs_vfpcc' : /arm_aapcs_vfpcc/
'arm_aapcscc' : /arm_aapcscc/
'arm_apcscc' : /arm_apcscc/
'ashr' : /ashr/
'asm' : /asm/
'attributes' : /attributes/
'available_externally' : /available_externally/
'avr_intrcc' : /avr_intrcc/
'avr_signalcc' : /avr_signalcc/
'bitcast' : /bitcast/
'blockaddress' : /blockaddress/
'builtin' : /builtin/
'byval' : /byval/
'c' : /c/
'cc' : /cc/
'ccc' : /ccc/
'cold' : /cold/
'coldcc' : /coldcc/
'comdat' : /comdat/
'common' : /common/
'constant' : /constant/
'convergent' : /convergent/
'cxx_fast_tlscc' : /cxx_fast_tlscc/
'datalayout' : /datalayout/
'declare' : /declare/
'default' : /default/
'define' : /define/
'dereferenceable_or_null' : /dereferenceable_or_null/
'dereferenceable' : /dereferenceable/
'distinct' : /distinct/
'dllexport' : /dllexport/
'dllimport' : /dllimport/
'double' : /double/
'dso_local' : /dso_local/
'dso_preemptable' : /dso_preemptable/
'eq' : /eq/
'exact' : /exact/
'exactmatch' : /exactmatch/
'extern_weak' : /extern_weak/
'external' : /external/
'externally_initialized' : /externally_initialized/
'extractelement' : /extractelement/
'extractvalue' : /extractvalue/
'fadd' : /fadd/
'false' : /false/
'fastcc' : /fastcc/
'fcmp' : /fcmp/
'fdiv' : /fdiv/
'float' : /float/
'fmul' : /fmul/
'fp128' : /fp128/
'fpext' : /fpext/
'fptosi' : /fptosi/
'fptoui' : /fptoui/
'fptrunc' : /fptrunc/
'frem' : /frem/
'fsub' : /fsub/
'gc' : /gc/
'getelementptr' : /getelementptr/
'ghccc' : /ghccc/
'global' : /global/
'half' : /half/
'hhvm_ccc' : /hhvm_ccc/
'hhvmcc' : /hhvmcc/
'hidden' : /hidden/
'icmp' : /icmp/
'ifunc' : /ifunc/
'inaccessiblemem_or_argmemonly' : /inaccessiblemem_or_argmemonly/
'inaccessiblememonly' : /inaccessiblememonly/
'inalloca' : /inalloca/
'inbounds' : /inbounds/
'initialexec' : /initialexec/
'inlinehint' : /inlinehint/
'inrange' : /inrange/
'inreg' : /inreg/
'insertelement' : /insertelement/
'insertvalue' : /insertvalue/
'intel_ocl_bicc' : /intel_ocl_bicc/
'inteldialect' : /inteldialect/
'internal' : /internal/
'inttoptr' : /inttoptr/
'jumptable' : /jumptable/
'label' : /label/
'largest' : /largest/
'linkonce_odr' : /linkonce_odr/
'linkonce' : /linkonce/
'local_unnamed_addr' : /local_unnamed_addr/
'localdynamic' : /localdynamic/
'localexec' : /localexec/
'lshr' : /lshr/
'metadata' : /metadata/
'minsize' : /minsize/
'module' : /module/
'msp430_intrcc' : /msp430_intrcc/
'mul' : /mul/
'naked' : /naked/
'ne' : /ne/
'nest' : /nest/
'noalias' : /noalias/
'nobuiltin' : /nobuiltin/
'nocapture' : /nocapture/
'noduplicate' : /noduplicate/
'noduplicates' : /noduplicates/
'noimplicitfloat' : /noimplicitfloat/
'noinline' : /noinline/
'none' : /none/
'nonlazybind' : /nonlazybind/
'nonnull' : /nonnull/
'norecurse' : /norecurse/
'noredzone' : /noredzone/
'noreturn' : /noreturn/
'nounwind' : /nounwind/
'nsw' : /nsw/
'null' : /null/
'nuw' : /nuw/
'oeq' : /oeq/
'oge' : /oge/
'ogt' : /ogt/
'ole' : /ole/
'olt' : /olt/
'one' : /one/
'opaque' : /opaque/
'optnone' : /optnone/
'optsize' : /optsize/
'or' : /or/
'ord' : /ord/
'personality' : /personality/
'ppc_fp128' : /ppc_fp128/
'prefix' : /prefix/
'preserve_allcc' : /preserve_allcc/
'preserve_mostcc' : /preserve_mostcc/
'private' : /private/
'prologue' : /prologue/
'protected' : /protected/
'ptrtoint' : /ptrtoint/
'ptx_device' : /ptx_device/
'ptx_kernel' : /ptx_kernel/
'readnone' : /readnone/
'readonly' : /readonly/
'returned' : /returned/
'returns_twice' : /returns_twice/
'safestack' : /safestack/
'samesize' : /samesize/
'sanitize_address' : /sanitize_address/
'sanitize_hwaddress' : /sanitize_hwaddress/
'sanitize_memory' : /sanitize_memory/
'sanitize_thread' : /sanitize_thread/
'sdiv' : /sdiv/
'section' : /section/
'select' : /select/
'sext' : /sext/
'sge' : /sge/
'sgt' : /sgt/
'shl' : /shl/
'shufflevector' : /shufflevector/
'sideeffect' : /sideeffect/
'signext' : /signext/
'sitofp' : /sitofp/
'sle' : /sle/
'slt' : /slt/
'source_filename' : /source_filename/
'speculatable' : /speculatable/
'spir_func' : /spir_func/
'spir_kernel' : /spir_kernel/
'srem' : /srem/
'sret' : /sret/
'ssp' : /ssp/
'sspreq' : /sspreq/
'sspstrong' : /sspstrong/
'strictfp' : /strictfp/
'sub' : /sub/
'swiftcc' : /swiftcc/
'swifterror' : /swifterror/
'swiftself' : /swiftself/
'target' : /target/
'thread_local' : /thread_local/
'to' : /to/
'token' : /token/
'triple' : /triple/
'true' : /true/
'trunc' : /trunc/
'type' : /type/
'udiv' : /udiv/
'ueq' : /ueq/
'uge' : /uge/
'ugt' : /ugt/
'uitofp' : /uitofp/
'ule' : /ule/
'ult' : /ult/
'undef' : /undef/
'une' : /une/
'unnamed_addr' : /unnamed_addr/
'uno' : /uno/
'urem' : /urem/
'uselistorder_bb' : /uselistorder_bb/
'uselistorder' : /uselistorder/
'uwtable' : /uwtable/
'void' : /void/
'weak_odr' : /weak_odr/
'weak' : /weak/
'webkit_jscc' : /webkit_jscc/
'win64cc' : /win64cc/
'writeonly' : /writeonly/
'x' : /x/
'x86_64_sysvcc' : /x86_64_sysvcc/
'x86_fastcallcc' : /x86_fastcallcc/
'x86_fp80' : /x86_fp80/
'x86_intrcc' : /x86_intrcc/
'x86_mmx' : /x86_mmx/
'x86_regcallcc' : /x86_regcallcc/
'x86_stdcallcc' : /x86_stdcallcc/
'x86_thiscallcc' : /x86_thiscallcc/
'x86_vectorcallcc' : /x86_vectorcallcc/
'xor' : /xor/
'zeroext' : /zeroext/
'zeroinitializer' : /zeroinitializer/
'zext' : /zext/

# TODO: remove placeholders.
placeholder1 : /placeholder1/
placeholder2 : /placeholder2/
placeholder3 : /placeholder3/

',' : /[,]/
'!' : /[!]/
'...' : /\.\.\./
'(' : /[(]/
')' : /[)]/
'[' : /[\[]/
']' : /[\]]/
'{' : /[{]/
'}' : /[}]/
'*' : /[*]/
'<' : /[<]/
'=' : /[=]/
'>' : /[>]/

# TODO: figure out how to handle string_lit_tok correctly.
string_lit_tok : /"[^"]"/

# ### [ Syntax part ] ##########################################################

# The LLVM IR grammar has been based on the source code of the official LLVM
# project, as of 2018-02-19 (rev db070bbdacd303ae7da129f59beaf35024d94c53).
#
#    * lib/AsmParser/LLParser.cpp

:: parser

input : Module;

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
	| ComdatDef
	| GlobalDecl
	| GlobalDef
	| IndirectSymbolDef
	| FunctionDecl
	| FunctionDef
	| AttrGroupDef
	| NamedMetadataDef
	| MetadataDef
	| UseListOrder
	| UseListOrderBB
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

# ~~~ [ Comdat Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#langref-comdats

# ref: parseComdat

ComdatDef
	: ComdatName '=' 'comdat' SelectionKind
;

SelectionKind
	: 'any'
	| 'exactmatch'
	| 'largest'
	| 'noduplicates'
	| 'samesize'
;

# ~~~ [ Global Variable Declaration ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#global-variables

# ref: ParseUnnamedGlobal
#
#   OptionalVisibility (ALIAS | IFUNC) ...
#   OptionalLinkage OptionalPreemptionSpecifier OptionalVisibility
#   OptionalDLLStorageClass
#                                                     ...   -> global variable
#   GlobalID '=' OptionalVisibility (ALIAS | IFUNC) ...
#   GlobalID '=' OptionalLinkage OptionalPreemptionSpecifier OptionalVisibility
#                OptionalDLLStorageClass
#                                                     ...   -> global variable

# ref: ParseNamedGlobal
#
#   GlobalVar '=' OptionalVisibility (ALIAS | IFUNC) ...
#   GlobalVar '=' OptionalLinkage OptionalPreemptionSpecifier
#                 OptionalVisibility OptionalDLLStorageClass
#                                                     ...   -> global variable

# ref: ParseGlobal
#
#   ::= GlobalVar '=' OptionalLinkage OptionalPreemptionSpecifier
#       OptionalVisibility OptionalDLLStorageClass
#       OptionalThreadLocal OptionalUnnamedAddr OptionalAddrSpace
#       OptionalExternallyInitialized GlobalType Type Const OptionalAttrs
#   ::= OptionalLinkage OptionalPreemptionSpecifier OptionalVisibility
#       OptionalDLLStorageClass OptionalThreadLocal OptionalUnnamedAddr
#       OptionalAddrSpace OptionalExternallyInitialized GlobalType Type
#       Const OptionalAttrs

GlobalDecl
	: GlobalIdent '=' ExternLinkage PreemptionSpecifieropt Visibilityopt DLLStorageClassopt ThreadLocalopt UnnamedAddropt AddrSpaceopt ExternallyInitializedopt Immutable Type (',' GlobalAttr)+? (',' FuncAttr)+?
;

# ~~~ [ Global Variable Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

GlobalDef
	: GlobalIdent '=' Linkageopt PreemptionSpecifieropt Visibilityopt DLLStorageClassopt ThreadLocalopt UnnamedAddropt AddrSpaceopt ExternallyInitializedopt Immutable Type Constant (',' GlobalAttr)+? (',' FuncAttr)+?
;

ExternallyInitialized
	: 'externally_initialized'
;

# ref: ParseGlobalType
#
#   ::= 'constant'
#   ::= 'global'

Immutable
	: 'constant'
	| 'global'
;

# NOTE: GlobalAttr should contain Alignment. However, using LALR(1) this
# produces a reduce/reduce conflict as FuncAttr also contains Alignment.
#
# Since GlobalAttr is only used in GlobalDecl and GlobalDef, both of which
# include a comma separated list of GlobalAttr and FuncAttr, we can simply
# remove Alignment from GlobalAttr to resolve the reduce/reduce conflict.

GlobalAttr
	: Section
	| Comdat
	#| Alignment # NOTE: removed to resolve reduce/reduce conflict, see above.
	#   ::= !dbg !57
	| MetadataAttachment
;

# ~~~ [ Indirect Symbol Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#aliases
# https://llvm.org/docs/LangRef.html#ifuncs

# ref: parseIndirectSymbol
#
#   ::= GlobalVar '=' OptionalLinkage OptionalPreemptionSpecifier
#                     OptionalVisibility OptionalDLLStorageClass
#                     OptionalThreadLocal OptionalUnnamedAddr
#                     'alias|ifunc' IndirectSymbol
#
#  IndirectSymbol
#   ::= TypeAndValue

IndirectSymbolDef
	: GlobalIdent '=' (ExternLinkage | Linkageopt) PreemptionSpecifieropt Visibilityopt DLLStorageClassopt ThreadLocalopt UnnamedAddropt IndirectSymbolKind Type ',' Type Constant
;

IndirectSymbolKind
	: 'alias'
	| 'ifunc'
;

# ~~~ [ Function Declaration ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#functions

# ref: ParseDeclare
#
#   ::= 'declare' FunctionHeader

FunctionDecl
	: 'declare' MetadataAttachment* ExternLinkageopt FunctionHeader
;

# ~~~ [ Function Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#functions

# ref: ParseDefine
#
#   ::= 'define' FunctionHeader (!dbg !56)* '{' ...

FunctionDef
	: 'define' Linkageopt FunctionHeader MetadataAttachment* FunctionBody
;

# ref: ParseFunctionHeader
#
#   ::= OptionalLinkage OptionalPreemptionSpecifier OptionalVisibility
#       OptionalCallingConv OptRetAttrs OptUnnamedAddr Type GlobalName
#       '(' ArgList ')' OptFuncAttrs OptSection OptionalAlign OptGC
#       OptionalPrefix OptionalPrologue OptPersonalityFn

# TODO: Add OptAlignment before OptGC once the LR-1 conflict has been resolved,
# as FuncAttrs also contains "align".

FunctionHeader
	: PreemptionSpecifieropt Visibilityopt DLLStorageClassopt CallingConvopt ReturnAttr* Type GlobalIdent '(' Params ')' UnnamedAddropt FuncAttr* Sectionopt Comdatopt GCopt Prefixopt Prologueopt Personalityopt
;

GC
	: 'gc' StringLit
;

Prefix
	: 'prefix' Type Constant
;

Prologue
	: 'prologue' Type Constant
;

Personality
	: 'personality' Type Constant
;

# ref: ParseFunctionBody
#
#   ::= '{' BasicBlock+ UseListOrderDirective* '}'

FunctionBody
	: '{' BasicBlock+ UseListOrder* '}'
;

# ~~~ [ Attribute Group Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#attribute-groups

# ref: ParseUnnamedAttrGrp
#
#   ::= 'attributes' AttrGrpID '=' '{' AttrValPair+ '}'

AttrGroupDef
	: 'attributes' AttrGroupID '=' '{' FuncAttr* '}'
;

# ~~~ [ Named Metadata Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#named-metadata

# ref: ParseNamedMetadata
#
#   !foo = !{ !1, !2 }

NamedMetadataDef
	: MetadataName '=' '!' '{' (MetadataNode separator ',')* '}'
;

MetadataNode
	: MetadataID
	# Parse DIExpressions inline as a special case. They are still MDNodes, so
	# they can still appear in named metadata. Remove this logic if they become
	# plain Metadata.
	| DIExpression
;

# ~~~ [ Metadata Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#metadata-nodes-and-metadata-strings

# ref: ParseStandaloneMetadata
#
#   !42 = !{...}

MetadataDef
	: MetadataID '=' Distinctopt MDTuple
	| MetadataID '=' Distinctopt SpecializedMDNode
;

Distinct
	: 'distinct'
;

# ~~~ [ Use-list Order Directives ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#use-list-order-directives

# ref: ParseUseListOrder
#
#   ::= 'uselistorder' Type Value ',' UseListOrderIndexes
#  UseListOrderIndexes
#   ::= '{' uint32 (',' uint32)+ '}'

UseListOrder
	: 'uselistorder' Type Value ',' '{' (int_lit_tok separator ',')+ '}' # TODO: use unsigned int lit?
;

# ref: ParseUseListOrderBB
#
#   ::= 'uselistorder_bb' @foo ',' %bar ',' UseListOrderIndexes

UseListOrderBB
	: 'uselistorder_bb' GlobalIdent ',' LocalIdent ',' '{' (int_lit_tok separator ',')+ '}' # TODO: use unsigned int lit?
;

# === [ Identifiers ] ==========================================================

# --- [ Global Identifiers ] ---------------------------------------------------

GlobalIdent
	: global_ident_tok
;

# --- [ Local Identifiers ] ----------------------------------------------------

LocalIdent
	: local_ident_tok
;

# --- [ Label Identifiers ] ----------------------------------------------------

LabelIdent
	: label_ident_tok
;

# --- [ Attribute Group Identifiers ] ------------------------------------------

AttrGroupID
	: attr_group_id_tok
;

# --- [ Comdat Identifiers ] ---------------------------------------------------

ComdatName
	: comdat_name_tok
;

# --- [ Metadata Identifiers ] -------------------------------------------------

MetadataName
	: metadata_name_tok
;

MetadataID
	: metadata_id_tok
;

# === [ Types ] ================================================================

# ref: ParseType
#
#  TYPEKEYWORD("void",      Type::getVoidTy(Context));
#  TYPEKEYWORD("half",      Type::getHalfTy(Context));
#  TYPEKEYWORD("float",     Type::getFloatTy(Context));
#  TYPEKEYWORD("double",    Type::getDoubleTy(Context));
#  TYPEKEYWORD("x86_fp80",  Type::getX86_FP80Ty(Context));
#  TYPEKEYWORD("fp128",     Type::getFP128Ty(Context));
#  TYPEKEYWORD("ppc_fp128", Type::getPPC_FP128Ty(Context));
#  TYPEKEYWORD("label",     Type::getLabelTy(Context));
#  TYPEKEYWORD("metadata",  Type::getMetadataTy(Context));
#  TYPEKEYWORD("x86_mmx",   Type::getX86_MMXTy(Context));
#  TYPEKEYWORD("token",     Type::getTokenTy(Context));

Type
	: VoidType
	| FuncType
	| FirstClassType
;

FirstClassType
	: ConcreteType
	| MetadataType
;

ConcreteType
	: IntType
	# Type ::= 'float' | 'void' (etc)
	| FloatType
	# Type ::= Type '*'
	# Type ::= Type 'addrspace' '(' uint32 ')' '*'
	| PointerType
	# Type ::= '<' ... '>'
	| VectorType
	| LabelType
	# Type ::= '[' ... ']'
	| ArrayType
	# Type ::= StructType
	| StructType
	# Type ::= %foo
	# Type ::= %4
	| NamedType
	| MMXType
	| TokenType
;

# --- [ Void Types ] -----------------------------------------------------------

VoidType
	: 'void'
;

# --- [ Function Types ] -------------------------------------------------------

# ref: ParseFunctionType
#
#  ::= Type ArgumentList OptionalAttrs

FuncType
	: Type '(' Params ')'
;

# --- [ Integer Types ] --------------------------------------------------------

IntType
	: int_type_tok
;

# --- [ Floating-point Types ] -------------------------------------------------

FloatType
	: FloatKind
;

FloatKind
	: 'half'
	| 'float'
	| 'double'
	| 'x86_fp80'
	| 'fp128'
	| 'ppc_fp128'
;

# --- [ MMX Types ] ------------------------------------------------------------

MMXType
	: 'x86_mmx'
;

# --- [ Pointer Types ] --------------------------------------------------------

PointerType
	: Type AddrSpaceopt '*'
;

# --- [ Vector Types ] ---------------------------------------------------------

# ref: ParseArrayVectorType
#
#     ::= '<' APSINTVAL 'x' Types '>'

VectorType
	: '<' int_lit_tok 'x' Type '>' # TODO: unsigned int lit?
;

# --- [ Label Types ] ----------------------------------------------------------

LabelType
	: 'label'
;

# --- [ Token Types ] ----------------------------------------------------------

TokenType
	: 'token'
;

# --- [ Metadata Types ] -------------------------------------------------------

MetadataType
	: 'metadata'
;

# --- [ Array Types ] ----------------------------------------------------------

# ref: ParseArrayVectorType
#
#     ::= '[' APSINTVAL 'x' Types ']'

ArrayType
	: '[' int_lit_tok 'x' Type ']' # TODO: unsigned int lit?
;

# --- [ Structure Types ] ------------------------------------------------------

# ref: ParseStructBody
#
#   StructType
#     ::= '{' '}'
#     ::= '{' Type (',' Type)* '}'
#     ::= '<' '{' '}' '>'
#     ::= '<' '{' Type (',' Type)* '}' '>'

StructType
	: '{' (Type separator ',')+? '}'
	| '<' '{' (Type separator ',')+? '}' '>'
;

OpaqueType
	: 'opaque'
;

# --- [ Named Types ] ----------------------------------------------------------

NamedType
	: LocalIdent
;

# === [ Values ] ===============================================================

# ref: ParseValue

Value
	: Constant
	# %42
	# %foo
	| LocalIdent
	| InlineAsm
;

# --- [ Inline Assembler Expressions ] -----------------------------------------

# https://llvm.org/docs/LangRef.html#inline-assembler-expressions

# ref: ParseValID
#
#  ::= 'asm' SideEffect? AlignStack? IntelDialect? STRINGCONSTANT ','
#             STRINGCONSTANT

InlineAsm
	: 'asm' SideEffectopt AlignStackopt IntelDialectopt StringLit ',' StringLit
;

SideEffect
	: 'sideeffect'
;

AlignStack
	: 'alignstack'
;

IntelDialect
	: 'inteldialect'
;

# === [ Constants ] ============================================================

# https://llvm.org/docs/LangRef.html#constants

# ref: ParseValID

Constant
	: BoolConst
	| IntConst
	| FloatConst
	| NullConst
	| NoneConst
	| StructConst
	| ArrayConst
	| CharArrayConst
	| VectorConst
	| ZeroInitializerConst
	# @42
	# @foo
	| GlobalIdent
	| UndefConst
	| BlockAddressConst
	| ConstantExpr
;

# --- [ Boolean Constants ] ----------------------------------------------------

# https://llvm.org/docs/LangRef.html#simple-constants

# ref: ParseValID

BoolConst
	: BoolLit
;

BoolLit
	: 'true'
	| 'false'
;

# --- [ Integer Constants ] ----------------------------------------------------

# https://llvm.org/docs/LangRef.html#simple-constants

# ref: ParseValID

IntConst
	: int_lit_tok
;

IntLit
	: int_lit_tok
;

# --- [ Floating-point Constants ] ---------------------------------------------

# https://llvm.org/docs/LangRef.html#simple-constants

# ref: ParseValID

FloatConst
	: float_lit_tok
;

# --- [ Null Pointer Constants ] -----------------------------------------------

# https://llvm.org/docs/LangRef.html#simple-constants

# ref: ParseValID

NullConst
	: 'null'
;

# --- [ Token Constants ] ------------------------------------------------------

# https://llvm.org/docs/LangRef.html#simple-constants

# ref: ParseValID

NoneConst
	: 'none'
;

# --- [ Structure Constants ] --------------------------------------------------

# https://llvm.org/docs/LangRef.html#complex-constants

# ref: ParseValID
#
#  ::= '{' ConstVector '}'
#  ::= '<' '{' ConstVector '}' '>' --> Packed Struct.

StructConst
	: '{' (Type Constant separator ',')+? '}'
	| '<' '{' (Type Constant separator ',')+? '}' '>'
;

# --- [ Array Constants ] ------------------------------------------------------

# https://llvm.org/docs/LangRef.html#complex-constants

# ref: ParseValID
#
#  c "foo"

ArrayConst
	: '[' (Type Constant separator ',')* ']'
;

CharArrayConst
	: 'c' StringLit
;

StringLit
	: string_lit_tok
;

# --- [ Vector Constants ] -----------------------------------------------------

# https://llvm.org/docs/LangRef.html#complex-constants

# ref: ParseValID
#
#  ::= '<' ConstVector '>'         --> Vector.

VectorConst
	: '<' (Type Constant separator ',')* '>'
;

# --- [ Zero Initialization Constants ] ----------------------------------------

# https://llvm.org/docs/LangRef.html#complex-constants

# ref: ParseValID

ZeroInitializerConst
	: 'zeroinitializer'
;

# --- [ Undefined Values ] -----------------------------------------------------

# https://llvm.org/docs/LangRef.html#undefined-values

# ref: ParseValID

UndefConst
	: 'undef'
;

# --- [ Addresses of Basic Blocks ] --------------------------------------------

# https://llvm.org/docs/LangRef.html#addresses-of-basic-blocks

# ref: ParseValID
#
#  ::= 'blockaddress' '(' @foo ',' %bar ')'

BlockAddressConst
	: 'blockaddress' '(' GlobalIdent ',' LocalIdent ')'
;

# === [ Constant expressions ] =================================================

# https://llvm.org/docs/LangRef.html#constant-expressions

# ref: ParseValID

ConstantExpr
	# Binary expressions
	: AddExpr
	| FAddExpr
	| SubExpr
	| FSubExpr
	| MulExpr
	| FMulExpr
	| UDivExpr
	| SDivExpr
	| FDivExpr
	| URemExpr
	| SRemExpr
	| FRemExpr
	# Bitwise expressions
	| ShlExpr
	| LShrExpr
	| AShrExpr
	| AndExpr
	| OrExpr
	| XorExpr
	# Vector expressions
	| ExtractElementExpr
	| InsertElementExpr
	| ShuffleVectorExpr
	# Aggregate expressions
	| ExtractValueExpr
	| InsertValueExpr
	# Memory expressions
	| GetElementPtrExpr
	# Conversion expressions
	| TruncExpr
	| ZExtExpr
	| SExtExpr
	| FPTruncExpr
	| FPExtExpr
	| FPToUIExpr
	| FPToSIExpr
	| UIToFPExpr
	| SIToFPExpr
	| PtrToIntExpr
	| IntToPtrExpr
	| BitCastExpr
	| AddrSpaceCastExpr
	# Other expressions
	| ICmpExpr
	| FCmpExpr
	| SelectExpr
;

# --- [ Binary expressions ] --------------------------------------------------

# https://llvm.org/docs/LangRef.html#constant-expressions

# ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

AddExpr
	: 'add' OverflowFlags '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FAddExpr
	: 'fadd' '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

SubExpr
	: 'sub' OverflowFlags '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FSubExpr
	: 'fsub' '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

MulExpr
	: 'mul' OverflowFlags '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FMulExpr
	: 'fmul' '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

UDivExpr
	: 'udiv' Exactopt '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

SDivExpr
	: 'sdiv' Exactopt '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FDivExpr
	: 'fdiv' '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

URemExpr
	: 'urem' '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

SRemExpr
	: 'srem' '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FRemExpr
	: 'frem' '(' Type Constant ',' Type Constant ')'
;

# --- [ Bitwise expressions ] --------------------------------------------------

# https://llvm.org/docs/LangRef.html#constant-expressions

# ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

ShlExpr
	: 'shl' OverflowFlags '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

LShrExpr
	: 'lshr' Exactopt '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

AShrExpr
	: 'ashr' Exactopt '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

AndExpr
	: 'and' '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

OrExpr
	: 'or' '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

XorExpr
	: 'xor' '(' Type Constant ',' Type Constant ')'
;

# --- [ Vector expressions ] ---------------------------------------------------

# https://llvm.org/docs/LangRef.html#constant-expressions

# ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

ExtractElementExpr
	: 'extractelement' '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

InsertElementExpr
	: 'insertelement' '(' Type Constant ',' Type Constant ',' Type Constant ')'
;

# ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

ShuffleVectorExpr
	: 'shufflevector' '(' Type Constant ',' Type Constant ',' Type Constant ')'
;

# --- [ Aggregate expressions ] ------------------------------------------------

# https://llvm.org/docs/LangRef.html#constant-expressions

# ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

ExtractValueExpr
	: 'extractvalue' '(' Type Constant Indices ')'
;

# ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

InsertValueExpr
	: 'insertvalue' '(' Type Constant ',' Type Constant Indices ')'
;

# --- [ Memory expressions ] ---------------------------------------------------

# https://llvm.org/docs/LangRef.html#constant-expressions

# ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

GetElementPtrExpr
	: 'getelementptr' InBoundsopt '(' Type ',' Type Constant (',' GEPConstIndex)* ')'
;

# ref: ParseGlobalValueVector
#
#   ::= empty
#   ::= [inrange] TypeAndValue (',' [inrange] TypeAndValue)*

# TODO: update name to Index?

GEPConstIndex
	: Inrangeopt Type Constant
;

Inrange
	: 'inrange'
;

# --- [ Conversion expressions ] -----------------------------------------------

# https://llvm.org/docs/LangRef.html#constant-expressions

# ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

TruncExpr
	: 'trunc' '(' Type Constant 'to' Type ')'
;

# ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

ZExtExpr
	: 'zext' '(' Type Constant 'to' Type ')'
;

# ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

SExtExpr
	: 'sext' '(' Type Constant 'to' Type ')'
;

# ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FPTruncExpr
	: 'fptrunc' '(' Type Constant 'to' Type ')'
;

# ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FPExtExpr
	: 'fpext' '(' Type Constant 'to' Type ')'
;

# ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FPToUIExpr
	: 'fptoui' '(' Type Constant 'to' Type ')'
;

# ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FPToSIExpr
	: 'fptosi' '(' Type Constant 'to' Type ')'
;

# ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

UIToFPExpr
	: 'uitofp' '(' Type Constant 'to' Type ')'
;

# ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

SIToFPExpr
	: 'sitofp' '(' Type Constant 'to' Type ')'
;

# ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

PtrToIntExpr
	: 'ptrtoint' '(' Type Constant 'to' Type ')'
;

# ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

IntToPtrExpr
	: 'inttoptr' '(' Type Constant 'to' Type ')'
;

# ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

BitCastExpr
	: 'bitcast' '(' Type Constant 'to' Type ')'
;

# ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

AddrSpaceCastExpr
	: 'addrspacecast' '(' Type Constant 'to' Type ')'
;

# --- [ Other expressions ] ----------------------------------------------------

# https://llvm.org/docs/LangRef.html#constant-expressions

# ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

ICmpExpr
	: 'icmp' IPred '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FCmpExpr
	: 'fcmp' FPred '(' Type Constant ',' Type Constant ')'
;

# ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

SelectExpr
	: 'select' '(' Type Constant ',' Type Constant ',' Type Constant ')'
;

# === [ Basic Blocks ] =========================================================

# ref: ParseBasicBlock
#
#   ::= LabelStr? Instruction*

BasicBlock
	: LabelIdent? Instruction* Terminator
;

# ///////////////////////////////

Instruction
   : placeholder1
;

Terminator
   : placeholder2
;

# TODO: fix placeholders.

ConstantExpr
   : placeholder3
;

MDTuple
   : placeholder1
;

SpecializedMDNode
   : placeholder2
;

DIExpression
   : placeholder3
;

# TODO: move Constant to where it belongs.

# === [ Metadata Nodes and Metadata Strings ] ==================================

# TODO: add proper implementation of metadata attachments.
MetadataAttachment
   : placeholder1
;

# ### [ Helper productions ] ###################################################

# https://llvm.org/docs/LangRef.html#linkage-types

# ref: ParseOptionalLinkage
#
#   ::= empty
#   ::= 'private'
#   ::= 'internal'
#   ::= 'weak'
#   ::= 'weak_odr'
#   ::= 'linkonce'
#   ::= 'linkonce_odr'
#   ::= 'available_externally'
#   ::= 'appending'
#   ::= 'common'
#   ::= 'extern_weak'
#   ::= 'external'

Linkage
	: 'appending'
	| 'available_externally'
	| 'common'
	| 'internal'
	| 'linkonce'
	| 'linkonce_odr'
	| 'private'
	| 'weak'
	| 'weak_odr'
;

ExternLinkage
	: 'extern_weak'
	| 'external'
;

# https://llvm.org/docs/LangRef.html#runtime-preemption-model

# ref: ParseOptionalDSOLocal

PreemptionSpecifier
	: 'dso_local'
	| 'dso_preemptable'
;

# https://llvm.org/docs/LangRef.html#visibility-styles

# ref: ParseOptionalVisibility
#
#   ::= empty
#   ::= 'default'
#   ::= 'hidden'
#   ::= 'protected'

Visibility
	: 'default'
	| 'hidden'
	| 'protected'
;

# https://llvm.org/docs/LangRef.html#dll-storage-classes

# ref: ParseOptionalDLLStorageClass
#
#   ::= empty
#   ::= 'dllimport'
#   ::= 'dllexport'

DLLStorageClass
	: 'dllexport'
	| 'dllimport'
;

# ref: ParseOptionalThreadLocal
#
#   := empty
#   := 'thread_local'
#   := 'thread_local' '(' tlsmodel ')'

ThreadLocal
	: 'thread_local'
	| 'thread_local' '(' TLSModel ')'
;

# ref: ParseTLSModel
#
#   := 'localdynamic'
#   := 'initialexec'
#   := 'localexec'

TLSModel
	: 'initialexec'
	| 'localdynamic'
	| 'localexec'
;

# ref: ParseOptionalUnnamedAddr

UnnamedAddr
	: 'local_unnamed_addr'
	| 'unnamed_addr'
;

# ref: ParseOptionalAddrSpace
#
#   := empty
#   := 'addrspace' '(' uint32 ')'

AddrSpace
	: 'addrspace' '(' int_lit_tok ')' # TODO: use unsigned int lit.
;

Section
	: 'section' StringLit
;

# ref: parseOptionalComdat

Comdat
	: 'comdat'
	| 'comdat' '(' ComdatName ')'
;

# ref: ParseOptionalAlignment
#
#   ::= empty
#   ::= 'align' 4

Alignment
	: 'align' int_lit_tok # TODO: use unsigned int lit.
;

# ___ [ Function Attribute ] ___________________________________________________

# ref: ParseFnAttributeValuePairs
#
#   ::= <attr> | <attr> '=' <value>

FuncAttr
	# not used in attribute groups.
	: AttrGroupID
	# used in attribute groups.
	| 'align' '=' int_lit_tok # TODO: use unsigned int lit?
	| 'alignstack' '=' int_lit_tok # TODO: use unsigned int lit?
	# used in functions.
	| Alignment
	| AllocSize
	| StackAlignment
	| AttrString
	| AttrPair
	| 'alwaysinline'
	| 'argmemonly'
	| 'builtin'
	| 'cold'
	| 'convergent'
	| 'inaccessiblemem_or_argmemonly'
	| 'inaccessiblememonly'
	| 'inlinehint'
	| 'jumptable'
	| 'minsize'
	| 'naked'
	| 'nobuiltin'
	| 'noduplicate'
	| 'noimplicitfloat'
	| 'noinline'
	| 'nonlazybind'
	| 'norecurse'
	| 'noredzone'
	| 'noreturn'
	| 'nounwind'
	| 'optnone'
	| 'optsize'
	| 'readnone'
	| 'readonly'
	| 'returns_twice'
	| 'safestack'
	| 'sanitize_address'
	| 'sanitize_hwaddress'
	| 'sanitize_memory'
	| 'sanitize_thread'
	| 'speculatable'
	| 'ssp'
	| 'sspreq'
	| 'sspstrong'
	| 'strictfp'
	| 'uwtable'
	| 'writeonly'
;

AttrString
	: StringLit
;

AttrPair
	: StringLit '=' StringLit
;

# ref: parseAllocSizeArguments

AllocSize
	: 'allocsize' '(' int_lit_tok ')' # TODO: use unsigned int lit?
	| 'allocsize' '(' int_lit_tok ',' int_lit_tok ')' # TODO: use unsigned int lit?
;

# ref: ParseOptionalStackAlignment
#
#   ::= empty
#   ::= 'alignstack' '(' 4 ')'
StackAlignment
	: 'alignstack' '(' int_lit_tok ')' # TODO: use unsigned int lit?
;

# ref: ParseOptionalCallingConv
#
#   ::= empty
#   ::= 'ccc'
#   ::= 'fastcc'
#   ::= 'intel_ocl_bicc'
#   ::= 'coldcc'
#   ::= 'x86_stdcallcc'
#   ::= 'x86_fastcallcc'
#   ::= 'x86_thiscallcc'
#   ::= 'x86_vectorcallcc'
#   ::= 'arm_apcscc'
#   ::= 'arm_aapcscc'
#   ::= 'arm_aapcs_vfpcc'
#   ::= 'msp430_intrcc'
#   ::= 'avr_intrcc'
#   ::= 'avr_signalcc'
#   ::= 'ptx_kernel'
#   ::= 'ptx_device'
#   ::= 'spir_func'
#   ::= 'spir_kernel'
#   ::= 'x86_64_sysvcc'
#   ::= 'win64cc'
#   ::= 'webkit_jscc'
#   ::= 'anyregcc'
#   ::= 'preserve_mostcc'
#   ::= 'preserve_allcc'
#   ::= 'ghccc'
#   ::= 'swiftcc'
#   ::= 'x86_intrcc'
#   ::= 'hhvmcc'
#   ::= 'hhvm_ccc'
#   ::= 'cxx_fast_tlscc'
#   ::= 'amdgpu_vs'
#   ::= 'amdgpu_ls'
#   ::= 'amdgpu_hs'
#   ::= 'amdgpu_es'
#   ::= 'amdgpu_gs'
#   ::= 'amdgpu_ps'
#   ::= 'amdgpu_cs'
#   ::= 'amdgpu_kernel'
#   ::= 'cc' UINT

CallingConv
	: 'amdgpu_cs'
	| 'amdgpu_es'
	| 'amdgpu_gs'
	| 'amdgpu_hs'
	| 'amdgpu_kernel'
	| 'amdgpu_ls'
	| 'amdgpu_ps'
	| 'amdgpu_vs'
	| 'anyregcc'
	| 'arm_aapcs_vfpcc'
	| 'arm_aapcscc'
	| 'arm_apcscc'
	| 'avr_intrcc'
	| 'avr_signalcc'
	| 'ccc'
	| 'coldcc'
	| 'cxx_fast_tlscc'
	| 'fastcc'
	| 'ghccc'
	| 'hhvm_ccc'
	| 'hhvmcc'
	| 'intel_ocl_bicc'
	| 'msp430_intrcc'
	| 'preserve_allcc'
	| 'preserve_mostcc'
	| 'ptx_device'
	| 'ptx_kernel'
	| 'spir_func'
	| 'spir_kernel'
	| 'swiftcc'
	| 'webkit_jscc'
	| 'win64cc'
	| 'x86_64_sysvcc'
	| 'x86_fastcallcc'
	| 'x86_intrcc'
	| 'x86_regcallcc'
	| 'x86_stdcallcc'
	| 'x86_thiscallcc'
	| 'x86_vectorcallcc'
	| 'cc' int_lit_tok # TODO: use unsigned int lit?
;

# ___ [ Return Attribute ] ___________________________________________________

# ref: ParseOptionalReturnAttrs

ReturnAttr
	: Alignment
	| Dereferenceable
	| AttrString
	| AttrPair
	| 'inreg'
	| 'noalias'
	| 'nonnull'
	| 'signext'
	| 'zeroext'
;

# ref: ParseArgumentList
#
#   ::= '(' ArgTypeListI ')'
#  ArgTypeListI
#   ::= empty
#   ::= '...'
#   ::= ArgTypeList ',' '...'
#   ::= ArgType (',' ArgType)*

Params
	: '...'?
	| (Param separator ',')+ (',' '...')?
;

Param
	: Type ParamAttr* LocalIdent?
;

# ___ [ Parameter Attribute ] __________________________________________________

# ref: ParseOptionalParamAttrs

# ref: ParseOptionalDerefAttrBytes
#
#   ::= empty
#   ::= AttrKind '(' 4 ')'

ParamAttr
	: Alignment
	| Dereferenceable
	| AttrString
	| AttrPair
	| 'byval'
	| 'inalloca'
	| 'inreg'
	| 'nest'
	| 'noalias'
	| 'nocapture'
	| 'nonnull'
	| 'readnone'
	| 'readonly'
	| 'returned'
	| 'signext'
	| 'sret'
	| 'swifterror'
	| 'swiftself'
	| 'writeonly'
	| 'zeroext'
;

Dereferenceable
	: 'dereferenceable' '(' int_lit_tok ')' # TODO: use unsigned int lit?
	| 'dereferenceable_or_null' '(' int_lit_tok ')' # TODO: use unsigned int lit?
;

Exact
	: 'exact'
;

OverflowFlags
	: ('nsw' | 'nuw')* # TODO: use ('nsw' & 'nuw')? if supported.
;

InBounds
	: 'inbounds'
;

# ref: ParseIndexList
#
#    ::=  (',' uint32)+

Indices
	: (int_lit_tok separator ',')*
;

# ref: ParseCmpPredicate

IPred
	: 'eq'
	| 'ne'
	| 'sge'
	| 'sgt'
	| 'sle'
	| 'slt'
	| 'uge'
	| 'ugt'
	| 'ule'
	| 'ult'
;

# ref: ParseCmpPredicate

FPred
	: 'false'
	| 'oeq'
	| 'oge'
	| 'ogt'
	| 'ole'
	| 'olt'
	| 'one'
	| 'ord'
	| 'true'
	| 'ueq'
	| 'uge'
	| 'ugt'
	| 'ule'
	| 'ult'
	| 'une'
	| 'uno'
;
