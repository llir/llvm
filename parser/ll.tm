language llvm(go);

lang = "llvm"
package = "github.com/mewmew/l-tm/parser"

# ### [ Lexical part ] #########################################################

:: lexer

# TODO: fix def of int_lit_tok

#   Integer           [-]?[0-9]+

int_lit_tok : /[-]?[0-9]+/

# TODO: fix proper definition of _name and _id.
_name = /foo/
_id = /42/

# TODO: add {_quoted_name|_name} to places where {_name} is used.

global_ident_tok : /{_global_name}|{_global_id}/

_global_name = /[@]{_name}/
_global_id = /[@]{_id}/

local_ident_tok : /{_local_name}|{_local_id}/

_local_name = /[%]{_name}/
_local_id = /[%]{_id}/

attr_group_id_tok : /[#]{_id}/

comdat_name_tok : /[$]{_name}/

'addrspace' : /addrspace/
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
'any' : /any/
'anyregcc' : /anyregcc/
'appending' : /appending/
'argmemonly' : /argmemonly/
'arm_aapcs_vfpcc' : /arm_aapcs_vfpcc/
'arm_aapcscc' : /arm_aapcscc/
'arm_apcscc' : /arm_apcscc/
'asm' : /asm/
'available_externally' : /available_externally/
'avr_intrcc' : /avr_intrcc/
'avr_signalcc' : /avr_signalcc/
'builtin' : /builtin/
'byval' : /byval/
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
'dllexport' : /dllexport/
'dllimport' : /dllimport/
'dso_local' : /dso_local/
'dso_preemptable' : /dso_preemptable/
'exactmatch' : /exactmatch/
'extern_weak' : /extern_weak/
'external' : /external/
'externally_initialized' : /externally_initialized/
'fastcc' : /fastcc/
'gc' : /gc/
'ghccc' : /ghccc/
'global' : /global/
'hhvm_ccc' : /hhvm_ccc/
'hhvmcc' : /hhvmcc/
'hidden' : /hidden/
'ifunc' : /ifunc/
'inaccessiblemem_or_argmemonly' : /inaccessiblemem_or_argmemonly/
'inaccessiblememonly' : /inaccessiblememonly/
'inalloca' : /inalloca/
'initialexec' : /initialexec/
'inlinehint' : /inlinehint/
'inreg' : /inreg/
'intel_ocl_bicc' : /intel_ocl_bicc/
'internal' : /internal/
'jumptable' : /jumptable/
'largest' : /largest/
'linkonce_odr' : /linkonce_odr/
'linkonce' : /linkonce/
'local_unnamed_addr' : /local_unnamed_addr/
'localdynamic' : /localdynamic/
'localexec' : /localexec/
'minsize' : /minsize/
'module' : /module/
'msp430_intrcc' : /msp430_intrcc/
'naked' : /naked/
'nest' : /nest/
'noalias' : /noalias/
'nobuiltin' : /nobuiltin/
'nocapture' : /nocapture/
'noduplicate' : /noduplicate/
'noduplicates' : /noduplicates/
'noimplicitfloat' : /noimplicitfloat/
'noinline' : /noinline/
'nonlazybind' : /nonlazybind/
'nonnull' : /nonnull/
'norecurse' : /norecurse/
'noredzone' : /noredzone/
'noreturn' : /noreturn/
'nounwind' : /nounwind/
'optnone' : /optnone/
'optsize' : /optsize/
'personality' : /personality/
'prefix' : /prefix/
'preserve_allcc' : /preserve_allcc/
'preserve_mostcc' : /preserve_mostcc/
'private' : /private/
'prologue' : /prologue/
'protected' : /protected/
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
'section' : /section/
'signext' : /signext/
'source_filename' : /source_filename/
'speculatable' : /speculatable/
'spir_func' : /spir_func/
'spir_kernel' : /spir_kernel/
'sret' : /sret/
'ssp' : /ssp/
'sspreq' : /sspreq/
'sspstrong' : /sspstrong/
'strictfp' : /strictfp/
'swiftcc' : /swiftcc/
'swifterror' : /swifterror/
'swiftself' : /swiftself/
'target' : /target/
'thread_local' : /thread_local/
'triple' : /triple/
'type' : /type/
'unnamed_addr' : /unnamed_addr/
'uwtable' : /uwtable/
'weak_odr' : /weak_odr/
'weak' : /weak/
'webkit_jscc' : /webkit_jscc/
'win64cc' : /win64cc/
'writeonly' : /writeonly/
'x86_64_sysvcc' : /x86_64_sysvcc/
'x86_fastcallcc' : /x86_fastcallcc/
'x86_intrcc' : /x86_intrcc/
'x86_regcallcc' : /x86_regcallcc/
'x86_stdcallcc' : /x86_stdcallcc/
'x86_thiscallcc' : /x86_thiscallcc/
'x86_vectorcallcc' : /x86_vectorcallcc/
'zeroext' : /zeroext/

# TODO: remove placeholders.
placeholder1 : /placeholder1/
placeholder2 : /placeholder2/

',' : /[,]/
'...' : /\.\.\./
'(' : /[(]/
')' : /[)]/
'{' : /[{]/
'}' : /[}]/
'=' : /[=]/

# TODO: figure out how to handle string_lit_tok correctly.
string_lit_tok : /"[^"]"/

# ### [ Syntax part ] ##########################################################

# The LLVM IR grammar has been based on the source code of the official LLVM
# project, as of 2018-02-19 (rev db070bbdacd303ae7da129f59beaf35024d94c53).
#
#    * lib/AsmParser/LLParser.cpp

:: parser

input : Module;

# TODO: move these to their corresponding place in ll.bnf.
StringLit
   : string_lit_tok
;

GlobalIdent
   : global_ident_tok
;

LocalIdent
   : local_ident_tok
;

AttrGroupID
   : attr_group_id_tok
;

ComdatName
   : comdat_name_tok
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
	| ComdatDef
	| GlobalDecl
	| GlobalDef
	| IndirectSymbolDef
	| FunctionDecl
	| FunctionDef
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
	: '{' BasicBlockList UseListOrders '}'
;

# ///////////////////////////////

BasicBlockList
   : placeholder1
;

UseListOrders
   : placeholder2
;

# TODO: move Constant to where it belongs.

Constant
   : placeholder1
;

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
