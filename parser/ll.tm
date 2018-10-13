language llvm(go);

lang = "llvm"
package = "github.com/mewmew/l-tm/parser"

# TODO: check when to use Fooopt and when to use Foo? (as based on the AST
# they produce)

# ### [ Lexical part ] #########################################################

:: lexer

# ### [ Lexical part ] #########################################################

_ascii_letter_upper = /[A-Z]/

_ascii_letter_lower = /[a-z]/

_ascii_letter = /{_ascii_letter_upper}|{_ascii_letter_lower}/

_letter = /{_ascii_letter}|[-$\._]/

_escape_letter = /{_letter}|[\\]/

_decimal_digit = /[0-9]/

_hex_digit = /{_decimal_digit}|[A-Fa-f]/

comment : /[;][^\r\n]*/               (space)
whitespace : /[\x00 \t\r\n]+/         (space)

# === [ Identifiers ] ==========================================================

_name = /{_letter}({_letter}|{_decimal_digit})*/

_escape_name = /{_escape_letter}({_escape_letter}|{_decimal_digit})*/

_quoted_name = /{_quoted_string}/

_id = /{_decimals}/

# --- [ Global identifiers ] ---------------------------------------------------

global_ident_tok : /{_global_name}|{_global_id}/

_global_name = /[@]({_name}|{_quoted_name})/

_global_id = /[@]{_id}/

# --- [ Local identifiers ] ----------------------------------------------------

local_ident_tok : /{_local_name}|{_local_id}/

_local_name = /[%]({_name}|{_quoted_name})/

_local_id = /[%]{_id}/

# --- [ Labels ] ---------------------------------------------------------------

#   Label             [-a-zA-Z$._0-9]+:

label_ident_tok : /(({_letter}|{_decimal_digit})({_letter}|{_decimal_digit})*[:])|({_quoted_string}[:])/   (class)

# --- [ Attribute group identifiers ] ------------------------------------------

attr_group_id_tok : /[#]{_id}/

# --- [ Comdat identifiers ] ---------------------------------------------------

comdat_name_tok : /[$]({_name}|{_quoted_name})/

# --- [ Metadata identifiers ] -------------------------------------------------

metadata_name_tok : /[!]{_escape_name}/   (class)

metadata_id_tok : /[!]{_id}/

# DW_TAG_foo
dwarf_tag_tok : /DW_TAG_({_ascii_letter}|{_decimal_digit}|[_])*/

# DW_ATE_foo
dwarf_att_encoding_tok : /DW_ATE_({_ascii_letter}|{_decimal_digit}|[_])*/

# DIFlagFoo
di_flag_tok : /DIFlag({_ascii_letter}|{_decimal_digit}|[_])*/

# DW_LANG_foo
dwarf_lang_tok : /DW_LANG_({_ascii_letter}|{_decimal_digit}|[_])*/

# DW_CC_foo
dwarf_cc_tok : /DW_CC_({_ascii_letter}|{_decimal_digit}|[_])*/

# CSK_foo
checksum_kind_tok : /CSK_({_ascii_letter}|{_decimal_digit}|[_])*/

# DW_VIRTUALITY_foo
dwarf_virtuality_tok : /DW_VIRTUALITY_({_ascii_letter}|{_decimal_digit}|[_])*/

# DW_MACINFO_foo
dwarf_macinfo_tok : /DW_MACINFO_({_ascii_letter}|{_decimal_digit}|[_])*/

# DW_OP_foo
dwarf_op_tok : /DW_OP_({_ascii_letter}|{_decimal_digit}|[_])*/

# === [ Integer literals ] =====================================================

#   Integer           [-]?[0-9]+

int_lit_tok : /[-]?[0-9]+/

_decimal_lit = /[-]?{_decimals}/

_decimals = /{_decimal_digit}+/

# === [ Floating-point literals ] ==============================================

#   FPConstant        [-+]?[0-9]+[.][0-9]*([eE][-+]?[0-9]+)?

float_lit_tok : /{_frac_lit}|{_sci_lit}|{_float_hex_lit}/

_frac_lit = /{_sign}?{_decimals}[\.]{_decimal_digit}*/

_sign = /[+-]/

_sci_lit = /{_frac_lit}[eE]{_sign}?{_decimals}/

#   HexFPConstant     0x[0-9A-Fa-f]+     // 16 hex digits
#   HexFP80Constant   0xK[0-9A-Fa-f]+    // 20 hex digits
#   HexFP128Constant  0xL[0-9A-Fa-f]+    // 32 hex digits
#   HexPPC128Constant 0xM[0-9A-Fa-f]+    // 32 hex digits
#   HexHalfConstant   0xH[0-9A-Fa-f]+    // 4 hex digits

_float_hex_lit = /0x[KLMH]?[0-9A-Fa-f]+/

# === [ String literals ] ======================================================

string_lit_tok : /{_quoted_string}/

_quoted_string = /["][^"]*["]/

# === [ Types ] ================================================================

int_type_tok : /i[0-9]+/

'!DIBasicType' : /!DIBasicType/
'!DICompileUnit' : /!DICompileUnit/
'!DICompositeType' : /!DICompositeType/
'!DIDerivedType' : /!DIDerivedType/
'!DIEnumerator' : /!DIEnumerator/
'!DIExpression' : /!DIExpression/
'!DIFile' : /!DIFile/
'!DIGlobalVariable' : /!DIGlobalVariable/
'!DIGlobalVariableExpression' : /!DIGlobalVariableExpression/
'!DIImportedEntity' : /!DIImportedEntity/
'!DILexicalBlock' : /!DILexicalBlock/
'!DILexicalBlockFile' : /!DILexicalBlockFile/
'!DILocalVariable' : /!DILocalVariable/
'!DILocation' : /!DILocation/
'!DIMacro' : /!DIMacro/
'!DIMacroFile' : /!DIMacroFile/
'!DIModule' : /!DIModule/
'!DINamespace' : /!DINamespace/
'!DIObjCProperty' : /!DIObjCProperty/
'!DISubprogram' : /!DISubprogram/
'!DISubrange' : /!DISubrange/
'!DISubroutineType' : /!DISubroutineType/
'!DITemplateTypeParameter' : /!DITemplateTypeParameter/
'!DITemplateValueParameter' : /!DITemplateValueParameter/
'!GenericDINode' : /!GenericDINode/
'acq_rel' : /acq_rel/
'acquire' : /acquire/
'add' : /add/
'addrspace' : /addrspace/
'addrspacecast' : /addrspacecast/
'afn' : /afn/
'alias' : /alias/
'align:' : /align:/
'align' : /align/
'alignstack' : /alignstack/
'alloca' : /alloca/
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
'arcp' : /arcp/
'arg:' : /arg:/
'argmemonly' : /argmemonly/
'arm_aapcs_vfpcc' : /arm_aapcs_vfpcc/
'arm_aapcscc' : /arm_aapcscc/
'arm_apcscc' : /arm_apcscc/
'ashr' : /ashr/
'asm' : /asm/
'atomic' : /atomic/
'atomicrmw' : /atomicrmw/
'attributes:' : /attributes:/
'attributes' : /attributes/
'available_externally' : /available_externally/
'avr_intrcc' : /avr_intrcc/
'avr_signalcc' : /avr_signalcc/
'baseType:' : /baseType:/
'bitcast' : /bitcast/
'blockaddress' : /blockaddress/
'br' : /br/
'builtin' : /builtin/
'byval' : /byval/
'c' : /c/
'call' : /call/
'caller' : /caller/
'catch' : /catch/
'catchpad' : /catchpad/
'catchret' : /catchret/
'catchswitch' : /catchswitch/
'cc:' : /cc:/
'cc' : /cc/
'ccc' : /ccc/
'checksum:' : /checksum:/
'checksumkind:' : /checksumkind:/
'cleanup' : /cleanup/
'cleanuppad' : /cleanuppad/
'cleanupret' : /cleanupret/
'cmpxchg' : /cmpxchg/
'cold' : /cold/
'coldcc' : /coldcc/
'column:' : /column:/
'comdat' : /comdat/
'common' : /common/
'configMacros:' : /configMacros:/
'constant' : /constant/
'containingType:' : /containingType:/
'contract' : /contract/
'convergent' : /convergent/
'count:' : /count:/
'cxx_fast_tlscc' : /cxx_fast_tlscc/
'datalayout' : /datalayout/
'debugInfoForProfiling:' : /debugInfoForProfiling:/
'declaration:' : /declaration:/
'declare' : /declare/
'default' : /default/
'define' : /define/
'dereferenceable_or_null' : /dereferenceable_or_null/
'dereferenceable' : /dereferenceable/
'directory:' : /directory:/
'discriminator:' : /discriminator:/
'distinct' : /distinct/
'dllexport' : /dllexport/
'dllimport' : /dllimport/
'double' : /double/
'dso_local' : /dso_local/
'dso_preemptable' : /dso_preemptable/
'dwarfAddressSpace:' : /dwarfAddressSpace:/
'dwoId:' : /dwoId:/
'elements:' : /elements:/
'emissionKind:' : /emissionKind:/
'encoding:' : /encoding:/
'entity:' : /entity:/
'enums:' : /enums:/
'eq' : /eq/
'exact' : /exact/
'exactmatch' : /exactmatch/
'exportSymbols:' : /exportSymbols:/
'expr:' : /expr:/
'extern_weak' : /extern_weak/
'external' : /external/
'externally_initialized' : /externally_initialized/
'extractelement' : /extractelement/
'extractvalue' : /extractvalue/
'extraData:' : /extraData:/
'fadd' : /fadd/
'false' : /false/
'fast' : /fast/
'fastcc' : /fastcc/
'fcmp' : /fcmp/
'fdiv' : /fdiv/
'fence' : /fence/
'file:' : /file:/
'filename:' : /filename:/
'filter' : /filter/
'flags:' : /flags:/
'float' : /float/
'fmul' : /fmul/
'fp128' : /fp128/
'fpext' : /fpext/
'fptosi' : /fptosi/
'fptoui' : /fptoui/
'fptrunc' : /fptrunc/
'frem' : /frem/
'from' : /from/
'fsub' : /fsub/
'FullDebug' : /FullDebug/
'gc' : /gc/
'getelementptr' : /getelementptr/
'getter:' : /getter:/
'ghccc' : /ghccc/
'global' : /global/
'globals:' : /globals:/
'gnuPubnames:' : /gnuPubnames:/
'half' : /half/
'header:' : /header:/
'hhvm_ccc' : /hhvm_ccc/
'hhvmcc' : /hhvmcc/
'hidden' : /hidden/
'icmp' : /icmp/
'identifier:' : /identifier:/
'ifunc' : /ifunc/
'imports:' : /imports:/
'inaccessiblemem_or_argmemonly' : /inaccessiblemem_or_argmemonly/
'inaccessiblememonly' : /inaccessiblememonly/
'inalloca' : /inalloca/
'inbounds' : /inbounds/
'includePath:' : /includePath:/
'indirectbr' : /indirectbr/
'initialexec' : /initialexec/
'inlinedAt:' : /inlinedAt:/
'inlinehint' : /inlinehint/
'inrange' : /inrange/
'inreg' : /inreg/
'insertelement' : /insertelement/
'insertvalue' : /insertvalue/
'intel_ocl_bicc' : /intel_ocl_bicc/
'inteldialect' : /inteldialect/
'internal' : /internal/
'inttoptr' : /inttoptr/
'invoke' : /invoke/
'isDefinition:' : /isDefinition:/
'isLocal:' : /isLocal:/
'isOptimized:' : /isOptimized:/
'isUnsigned:' : /isUnsigned:/
'isysroot:' : /isysroot:/
'jumptable' : /jumptable/
'label' : /label/
'landingpad' : /landingpad/
'language:' : /language:/
'largest' : /largest/
'line:' : /line:/
'LineTablesOnly' : /LineTablesOnly/
'linkageName:' : /linkageName:/
'linkonce_odr' : /linkonce_odr/
'linkonce' : /linkonce/
'load' : /load/
'local_unnamed_addr' : /local_unnamed_addr/
'localdynamic' : /localdynamic/
'localexec' : /localexec/
'lowerBound:' : /lowerBound:/
'lshr' : /lshr/
'macros:' : /macros:/
'max' : /max/
'metadata' : /metadata/
'min' : /min/
'minsize' : /minsize/
'module' : /module/
'monotonic' : /monotonic/
'msp430_intrcc' : /msp430_intrcc/
'mul' : /mul/
'musttail' : /musttail/
'naked' : /naked/
'name:' : /name:/
'nand' : /nand/
'ne' : /ne/
'nest' : /nest/
'ninf' : /ninf/
'nnan' : /nnan/
'noalias' : /noalias/
'nobuiltin' : /nobuiltin/
'nocapture' : /nocapture/
'NoDebug' : /NoDebug/
'nodes:' : /nodes:/
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
'notail' : /notail/
'nounwind' : /nounwind/
'nsw' : /nsw/
'nsz' : /nsz/
'null' : /null/
'nuw' : /nuw/
'oeq' : /oeq/
'offset:' : /offset:/
'oge' : /oge/
'ogt' : /ogt/
'ole' : /ole/
'olt' : /olt/
'one' : /one/
'opaque' : /opaque/
'operands:' : /operands:/
'optnone' : /optnone/
'optsize' : /optsize/
'or' : /or/
'ord' : /ord/
'personality' : /personality/
'phi' : /phi/
'ppc_fp128' : /ppc_fp128/
'prefix' : /prefix/
'preserve_allcc' : /preserve_allcc/
'preserve_mostcc' : /preserve_mostcc/
'private' : /private/
'producer:' : /producer:/
'prologue' : /prologue/
'protected' : /protected/
'ptrtoint' : /ptrtoint/
'ptx_device' : /ptx_device/
'ptx_kernel' : /ptx_kernel/
'readnone' : /readnone/
'readonly' : /readonly/
'reassoc' : /reassoc/
'release' : /release/
'resume' : /resume/
'ret' : /ret/
'retainedTypes:' : /retainedTypes:/
'returned' : /returned/
'returns_twice' : /returns_twice/
'runtimeLang:' : /runtimeLang:/
'runtimeVersion:' : /runtimeVersion:/
'safestack' : /safestack/
'samesize' : /samesize/
'sanitize_address' : /sanitize_address/
'sanitize_hwaddress' : /sanitize_hwaddress/
'sanitize_memory' : /sanitize_memory/
'sanitize_thread' : /sanitize_thread/
'scope:' : /scope:/
'scopeLine:' : /scopeLine:/
'sdiv' : /sdiv/
'section' : /section/
'select' : /select/
'seq_cst' : /seq_cst/
'setter:' : /setter:/
'sext' : /sext/
'sge' : /sge/
'sgt' : /sgt/
'shl' : /shl/
'shufflevector' : /shufflevector/
'sideeffect' : /sideeffect/
'signext' : /signext/
'singlethread' : /singlethread/
'sitofp' : /sitofp/
'size:' : /size:/
'sle' : /sle/
'slt' : /slt/
'source_filename' : /source_filename/
'speculatable' : /speculatable/
'spir_func' : /spir_func/
'spir_kernel' : /spir_kernel/
'splitDebugFilename:' : /splitDebugFilename:/
'splitDebugInlining:' : /splitDebugInlining:/
'srem' : /srem/
'sret' : /sret/
'ssp' : /ssp/
'sspreq' : /sspreq/
'sspstrong' : /sspstrong/
'store' : /store/
'strictfp' : /strictfp/
'sub' : /sub/
'swiftcc' : /swiftcc/
'swifterror' : /swifterror/
'swiftself' : /swiftself/
'switch' : /switch/
'syncscope' : /syncscope/
'tag:' : /tag:/
'tail' : /tail/
'target' : /target/
'templateParams:' : /templateParams:/
'thisAdjustment:' : /thisAdjustment:/
'thread_local' : /thread_local/
'thrownTypes:' : /thrownTypes:/
'to' : /to/
'token' : /token/
'triple' : /triple/
'true' : /true/
'trunc' : /trunc/
'type:' : /type:/
'type' : /type/
'types:' : /types:/
'udiv' : /udiv/
'ueq' : /ueq/
'uge' : /uge/
'ugt' : /ugt/
'uitofp' : /uitofp/
'ule' : /ule/
'ult' : /ult/
'umax' : /umax/
'umin' : /umin/
'undef' : /undef/
'une' : /une/
'unit:' : /unit:/
'unnamed_addr' : /unnamed_addr/
'uno' : /uno/
'unordered' : /unordered/
'unreachable' : /unreachable/
'unwind' : /unwind/
'urem' : /urem/
'uselistorder_bb' : /uselistorder_bb/
'uselistorder' : /uselistorder/
'uwtable' : /uwtable/
'va_arg' : /va_arg/
'value:' : /value:/
'var:' : /var:/
'variables:' : /variables:/
'virtualIndex:' : /virtualIndex:/
'virtuality:' : /virtuality:/
'void' : /void/
'volatile' : /volatile/
'vtableHolder:' : /vtableHolder:/
'weak_odr' : /weak_odr/
'weak' : /weak/
'webkit_jscc' : /webkit_jscc/
'win64cc' : /win64cc/
'within' : /within/
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
'xchg' : /xchg/
'xor' : /xor/
'zeroext' : /zeroext/
'zeroinitializer' : /zeroinitializer/
'zext' : /zext/

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
'|' : /[|]/

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

GlobalAttr
	: Section
	| Comdat
	| Alignment
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
	: 'declare' FunctionMetadata ExternLinkageopt FunctionHeader
;

# ~~~ [ Function Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#functions

# ref: ParseDefine
#
#   ::= 'define' FunctionHeader (!dbg !56)* '{' ...

FunctionDef
	: 'define' Linkageopt FunctionHeader FunctionMetadata FunctionBody
;

# ref: ParseFunctionHeader
#
#   ::= OptionalLinkage OptionalPreemptionSpecifier OptionalVisibility
#       OptionalCallingConv OptRetAttrs OptUnnamedAddr Type GlobalName
#       '(' ArgList ')' OptFuncAttrs OptSection OptionalAlign OptGC
#       OptionalPrefix OptionalPrologue OptPersonalityFn

# TODO: Add OptAlignment before OptGC once the LR-1 conflict has been resolved.
# The shift/reduce conflict is present since FuncAttr also contains 'align'.

FunctionHeader
	: PreemptionSpecifieropt Visibilityopt DLLStorageClassopt CallingConvopt ReturnAttr* Type GlobalIdent '(' Params ')' UnnamedAddropt (FuncAttr | Alignment)* Sectionopt Comdatopt GCopt Prefixopt Prologueopt Personalityopt
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
	: 'attributes' AttrGroupID '=' '{' (FuncAttr | Alignment)* '}'
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
	: 'uselistorder' Type Value ',' '{' (UintLit separator ',')+ '}'
;

# ref: ParseUseListOrderBB
#
#   ::= 'uselistorder_bb' @foo ',' %bar ',' UseListOrderIndexes

UseListOrderBB
	: 'uselistorder_bb' GlobalIdent ',' LocalIdent ',' '{' (UintLit separator ',')+ '}'
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
	: '<' UintLit 'x' Type '>'
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
	: '[' UintLit 'x' Type ']'
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

UintLit
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

# === [ Instructions ] =========================================================

# https://llvm.org/docs/LangRef.html#instruction-reference

# ref: ParseInstruction

Instruction
	# Instructions not producing values.
	: StoreInst
	| FenceInst
	| CmpXchgInst
	| AtomicRMWInst
	# Instructions producing values.
	| LocalIdent '=' ValueInstruction
	| ValueInstruction
;

ValueInstruction
	# Binary instructions
	: AddInst
	| FAddInst
	| SubInst
	| FSubInst
	| MulInst
	| FMulInst
	| UDivInst
	| SDivInst
	| FDivInst
	| URemInst
	| SRemInst
	| FRemInst
	# Bitwise instructions
	| ShlInst
	| LShrInst
	| AShrInst
	| AndInst
	| OrInst
	| XorInst
	# Vector instructions
	| ExtractElementInst
	| InsertElementInst
	| ShuffleVectorInst
	# Aggregate instructions
	| ExtractValueInst
	| InsertValueInst
	# Memory instructions
	| AllocaInst
	| LoadInst
	| GetElementPtrInst
	# Conversion instructions
	| TruncInst
	| ZExtInst
	| SExtInst
	| FPTruncInst
	| FPExtInst
	| FPToUIInst
	| FPToSIInst
	| UIToFPInst
	| SIToFPInst
	| PtrToIntInst
	| IntToPtrInst
	| BitCastInst
	| AddrSpaceCastInst
	# Other instructions
	| ICmpInst
	| FCmpInst
	| PhiInst
	| SelectInst
	| CallInst
	| VAArgInst
	| LandingPadInst
	| CatchPadInst
	| CleanupPadInst
;

# --- [ Binary instructions ] --------------------------------------------------

# ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#add-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

AddInst
	: 'add' OverflowFlags Type Value ',' Value InstructionMetadata
;

# ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fadd-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

FAddInst
	: 'fadd' FastMathFlag* Type Value ',' Value InstructionMetadata
;

# ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#sub-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

SubInst
	: 'sub' OverflowFlags Type Value ',' Value InstructionMetadata
;

# ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fsub-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

FSubInst
	: 'fsub' FastMathFlag* Type Value ',' Value InstructionMetadata
;

# ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#mul-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

MulInst
	: 'mul' OverflowFlags Type Value ',' Value InstructionMetadata
;

# ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fmul-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

FMulInst
	: 'fmul' FastMathFlag* Type Value ',' Value InstructionMetadata
;

# ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#udiv-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

UDivInst
	: 'udiv' Exactopt Type Value ',' Value InstructionMetadata
;

# ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#sdiv-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

SDivInst
	: 'sdiv' Exactopt Type Value ',' Value InstructionMetadata
;

# ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fdiv-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

FDivInst
	: 'fdiv' FastMathFlag* Type Value ',' Value InstructionMetadata
;

# ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#urem-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

URemInst
	: 'urem' Type Value ',' Value InstructionMetadata
;

# ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#srem-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

SRemInst
	: 'srem' Type Value ',' Value InstructionMetadata
;

# ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#frem-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

FRemInst
	: 'frem' FastMathFlag* Type Value ',' Value InstructionMetadata
;

# --- [ Bitwise instructions ] -------------------------------------------------

# ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#shl-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

ShlInst
	: 'shl' OverflowFlags Type Value ',' Value InstructionMetadata
;

# ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#lshr-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

LShrInst
	: 'lshr' Exactopt Type Value ',' Value InstructionMetadata
;

# ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#ashr-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

AShrInst
	: 'ashr' Exactopt Type Value ',' Value InstructionMetadata
;

# ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#and-instruction

# ref: ParseLogical
#
#  ::= ArithmeticOps TypeAndValue ',' Value {

AndInst
	: 'and' Type Value ',' Value InstructionMetadata
;

# ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#or-instruction

# ref: ParseLogical
#
#  ::= ArithmeticOps TypeAndValue ',' Value {

OrInst
	: 'or' Type Value ',' Value InstructionMetadata
;

# ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#xor-instruction

# ref: ParseLogical
#
#  ::= ArithmeticOps TypeAndValue ',' Value {

XorInst
	: 'xor' Type Value ',' Value InstructionMetadata
;

# --- [ Vector instructions ] --------------------------------------------------

# ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#extractelement-instruction

# ref: ParseExtractElement
#
#   ::= 'extractelement' TypeAndValue ',' TypeAndValue

ExtractElementInst
	: 'extractelement' Type Value ',' Type Value InstructionMetadata
;

# ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#insertelement-instruction

# ref: ParseInsertElement
#
#   ::= 'insertelement' TypeAndValue ',' TypeAndValue ',' TypeAndValue

InsertElementInst
	: 'insertelement' Type Value ',' Type Value ',' Type Value InstructionMetadata
;

# ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#shufflevector-instruction

# ref: ParseShuffleVector
#
#   ::= 'shufflevector' TypeAndValue ',' TypeAndValue ',' TypeAndValue

ShuffleVectorInst
	: 'shufflevector' Type Value ',' Type Value ',' Type Value InstructionMetadata
;

# --- [ Aggregate instructions ] -----------------------------------------------

# ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#extractvalue-instruction

# ref: ParseExtractValue
#
#   ::= 'extractvalue' TypeAndValue (',' uint32)+

ExtractValueInst
   : 'extractvalue' Type Value (',' UintLit)+ InstructionMetadata
;

# ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#insertvalue-instruction

# ref: ParseInsertValue
#
#   ::= 'insertvalue' TypeAndValue ',' TypeAndValue (',' uint32)+

InsertValueInst
   : 'insertvalue' Type Value ',' Type Value (',' UintLit)+ InstructionMetadata
;

# --- [ Memory instructions ] --------------------------------------------------

# ~~~ [ alloca ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#alloca-instruction

# ref: ParseAlloc
#
#   ::= 'alloca' 'inalloca'? 'swifterror'? Type (',' TypeAndValue)?
#       (',' 'align' i32)? (',', 'addrspace(n))?

AllocaInst
	: 'alloca' InAllocaopt SwiftErroropt Type (',' Type Value)? (',' Alignment)? (',' AddrSpace)? InstructionMetadata
;

InAlloca
	: 'inalloca'
;

SwiftError
	: 'swifterror'
;

# ~~~ [ load ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#load-instruction

# ref: ParseLoad
#
#   ::= 'load' 'volatile'? TypeAndValue (',' 'align' i32)?
#   ::= 'load' 'atomic' 'volatile'? TypeAndValue
#       'singlethread'? AtomicOrdering (',' 'align' i32)?

LoadInst
	# Load.
	: 'load' Volatileopt Type ',' Type Value (',' Alignment)? InstructionMetadata
	# Atomic load.
	| 'load' 'atomic' Volatileopt Type ',' Type Value SyncScopeopt AtomicOrdering (',' Alignment)? InstructionMetadata
;

# ~~~ [ store ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#store-instruction

# ref: ParseStore
#
#   ::= 'store' 'volatile'? TypeAndValue ',' TypeAndValue (',' 'align' i32)?
#   ::= 'store' 'atomic' 'volatile'? TypeAndValue ',' TypeAndValue
#       'singlethread'? AtomicOrdering (',' 'align' i32)?

StoreInst
	: 'store' Volatileopt Type Value ',' Type Value (',' Alignment)? InstructionMetadata
	| 'store' 'atomic' Volatileopt Type Value ',' Type Value SyncScopeopt AtomicOrdering (',' Alignment)? InstructionMetadata
;

# ~~~ [ fence ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fence-instruction

# ref: ParseFence
#
#   ::= 'fence' 'singlethread'? AtomicOrdering

FenceInst
	: 'fence' SyncScopeopt AtomicOrdering InstructionMetadata
;

# ~~~ [ cmpxchg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#cmpxchg-instruction

# ref: ParseCmpXchg
#
#   ::= 'cmpxchg' 'weak'? 'volatile'? TypeAndValue ',' TypeAndValue ','
#       TypeAndValue 'singlethread'? AtomicOrdering AtomicOrdering

CmpXchgInst
	: 'cmpxchg' Weakopt Volatileopt Type Value ',' Type Value ',' Type Value SyncScopeopt AtomicOrdering AtomicOrdering InstructionMetadata
;

Weak
	: 'weak'
;

# ~~~ [ atomicrmw ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#atomicrmw-instruction

# ref: ParseAtomicRMW
#
#   ::= 'atomicrmw' 'volatile'? BinOp TypeAndValue ',' TypeAndValue
#       'singlethread'? AtomicOrdering

AtomicRMWInst
	: 'atomicrmw' Volatileopt BinOp Type Value ',' Type Value SyncScopeopt AtomicOrdering InstructionMetadata
;

BinOp
	: 'add'
	| 'and'
	| 'max'
	| 'min'
	| 'nand'
	| 'or'
	| 'sub'
	| 'umax'
	| 'umin'
	| 'xchg'
	| 'xor'
;

# ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#getelementptr-instruction

# ref: ParseGetElementPtr
#
#   ::= 'getelementptr' 'inbounds'? TypeAndValue (',' TypeAndValue)*

GetElementPtrInst
	: 'getelementptr' InBoundsopt Type ',' Type Value (',' Type Value)* InstructionMetadata
;

# --- [ Conversion instructions ] ----------------------------------------------

# ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#trunc-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

TruncInst
	: 'trunc' Type Value 'to' Type InstructionMetadata
;

# ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#zext-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

ZExtInst
	: 'zext' Type Value 'to' Type InstructionMetadata
;

# ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#sext-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

SExtInst
	: 'sext' Type Value 'to' Type InstructionMetadata
;

# ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fptrunc-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

FPTruncInst
	: 'fptrunc' Type Value 'to' Type InstructionMetadata
;

# ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fpext-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

FPExtInst
	: 'fpext' Type Value 'to' Type InstructionMetadata
;

# ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fptoui-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

FPToUIInst
	: 'fptoui' Type Value 'to' Type InstructionMetadata
;

# ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fptosi-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

FPToSIInst
	: 'fptosi' Type Value 'to' Type InstructionMetadata
;

# ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#uitofp-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

UIToFPInst
	: 'uitofp' Type Value 'to' Type InstructionMetadata
;

# ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#sitofp-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

SIToFPInst
	: 'sitofp' Type Value 'to' Type InstructionMetadata
;

# ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#ptrtoint-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

PtrToIntInst
	: 'ptrtoint' Type Value 'to' Type InstructionMetadata
;

# ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#inttoptr-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

IntToPtrInst
	: 'inttoptr' Type Value 'to' Type InstructionMetadata
;

# ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#bitcast-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

BitCastInst
	: 'bitcast' Type Value 'to' Type InstructionMetadata
;

# ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#addrspacecast-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

AddrSpaceCastInst
	: 'addrspacecast' Type Value 'to' Type InstructionMetadata
;

# --- [ Other instructions ] ---------------------------------------------------

# ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#icmp-instruction

# ref: ParseCompare
#
#  ::= 'icmp' IPredicates TypeAndValue ',' Value

ICmpInst
	: 'icmp' IPred Type Value ',' Value InstructionMetadata
;

# ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fcmp-instruction

# ref: ParseCompare
#
#  ::= 'fcmp' FPredicates TypeAndValue ',' Value

FCmpInst
	: 'fcmp' FastMathFlag* FPred Type Value ',' Value InstructionMetadata
;

# ~~~ [ phi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#phi-instruction

# ref: ParsePHI
#
#   ::= 'phi' Type '[' Value ',' Value ']' (',' '[' Value ',' Value ']')*

PhiInst
	: 'phi' Type (Inc separator ',')+ InstructionMetadata
;

Inc
	: '[' Value ',' LocalIdent ']'
;

# ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#select-instruction

# ref: ParseSelect
#
#   ::= 'select' TypeAndValue ',' TypeAndValue ',' TypeAndValue

SelectInst
	: 'select' Type Value ',' Type Value ',' Type Value InstructionMetadata
;

# ~~~ [ call ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#call-instruction

# ref: ParseCall
#
#   ::= 'call' OptionalFastMathFlags OptionalCallingConv
#           OptionalAttrs Type Value ParameterList OptionalAttrs
#   ::= 'tail' 'call' OptionalFastMathFlags OptionalCallingConv
#           OptionalAttrs Type Value ParameterList OptionalAttrs
#   ::= 'musttail' 'call' OptionalFastMathFlags OptionalCallingConv
#           OptionalAttrs Type Value ParameterList OptionalAttrs
#   ::= 'notail' 'call'  OptionalFastMathFlags OptionalCallingConv
#           OptionalAttrs Type Value ParameterList OptionalAttrs

CallInst
	: Tailopt 'call' FastMathFlag* CallingConvopt ReturnAttr* Type Value '(' Args ')' (FuncAttr | Alignment)* OperandBundles InstructionMetadata
;

Tail
	: 'musttail'
	| 'notail'
	| 'tail'
;

# ~~~ [ va_arg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#va_arg-instruction

# ref: ParseVA_Arg
#
#   ::= 'va_arg' TypeAndValue ',' Type

VAArgInst
	: 'va_arg' Type Value ',' Type InstructionMetadata
;

# ~~~ [ landingpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#landingpad-instruction

# ref: ParseLandingPad
#
#   ::= 'landingpad' Type 'personality' TypeAndValue 'cleanup'? Clause+
#  Clause
#   ::= 'catch' TypeAndValue
#   ::= 'filter'
#   ::= 'filter' TypeAndValue ( ',' TypeAndValue )*

LandingPadInst
	: 'landingpad' Type Cleanupopt Clause* InstructionMetadata
;

Cleanup
	: 'cleanup'
;

Clause
	: 'catch' Type Value
	| 'filter' Type ArrayConst
;

# --- [ catchpad ] -------------------------------------------------------------

# ref: ParseCatchPad
#
#   ::= 'catchpad' ParamList 'to' TypeAndValue 'unwind' TypeAndValue

CatchPadInst
	: 'catchpad' 'within' LocalIdent '[' (ExceptionArg separator ',')* ']' InstructionMetadata
;

# --- [ cleanuppad ] -----------------------------------------------------------

# ref: ParseCleanupPad
#
#   ::= 'cleanuppad' within Parent ParamList

CleanupPadInst
	: 'cleanuppad' 'within' ExceptionScope '[' (ExceptionArg separator ',')* ']' InstructionMetadata
;

# === [ Terminators ] ==========================================================

# https://llvm.org/docs/LangRef.html#terminator-instructions

# ref: ParseInstruction

Terminator
	: RetTerm
	| BrTerm
	| CondBrTerm
	| SwitchTerm
	| IndirectBrTerm
	| InvokeTerm
	| ResumeTerm
	| CatchSwitchTerm
	| CatchRetTerm
	| CleanupRetTerm
	| UnreachableTerm
;

# --- [ ret ] ------------------------------------------------------------------

# https://llvm.org/docs/LangRef.html#ret-instruction

# ref: ParseRet
#
#   ::= 'ret' void (',' !dbg, !1)*
#   ::= 'ret' TypeAndValue (',' !dbg, !1)*

RetTerm
	# Void return.
	: 'ret' VoidType InstructionMetadata
	# Value return.
	| 'ret' ConcreteType Value InstructionMetadata
;

# --- [ br ] -------------------------------------------------------------------

# https://llvm.org/docs/LangRef.html#br-instruction

# ref: ParseBr
#
#   ::= 'br' TypeAndValue
#   ::= 'br' TypeAndValue ',' TypeAndValue ',' TypeAndValue

# Unconditional branch.
BrTerm
	: 'br' LabelType LocalIdent InstructionMetadata
;

# Conditional branch.
CondBrTerm
	: 'br' IntType Value ',' LabelType LocalIdent ',' LabelType LocalIdent InstructionMetadata
;

# --- [ switch ] ---------------------------------------------------------------

# https://llvm.org/docs/LangRef.html#switch-instruction

# ref: ParseSwitch
#
#    ::= 'switch' TypeAndValue ',' TypeAndValue '[' JumpTable ']'
#  JumpTable
#    ::= (TypeAndValue ',' TypeAndValue)*

SwitchTerm
	: 'switch' Type Value ',' LabelType LocalIdent '[' Case* ']' InstructionMetadata
;

Case
	: Type IntConst ',' LabelType LocalIdent
;

# --- [ indirectbr ] -----------------------------------------------------------

# https://llvm.org/docs/LangRef.html#indirectbr-instruction

# ref: ParseIndirectBr
#
#    ::= 'indirectbr' TypeAndValue ',' '[' LabelList ']'

IndirectBrTerm
	: 'indirectbr' Type Value ',' '[' LabelList ']' InstructionMetadata
;

LabelList
	: Label
	| LabelList ',' Label
;

Label
	: LabelType LocalIdent
;

# --- [ invoke ] ---------------------------------------------------------------

# https://llvm.org/docs/LangRef.html#invoke-instruction

# ref: ParseInvoke
#
#   ::= 'invoke' OptionalCallingConv OptionalAttrs Type Value ParamList
#       OptionalAttrs 'to' TypeAndValue 'unwind' TypeAndValue

InvokeTerm
	: 'invoke' CallingConvopt ReturnAttr* Type Value '(' Args ')' (FuncAttr | Alignment)* OperandBundles 'to' LabelType LocalIdent 'unwind' LabelType LocalIdent InstructionMetadata
;

# --- [ resume ] ---------------------------------------------------------------

# https://llvm.org/docs/LangRef.html#resume-instruction

# ref: ParseResume
#
#   ::= 'resume' TypeAndValue

ResumeTerm
	: 'resume' Type Value InstructionMetadata
;

# --- [ catchswitch ] ----------------------------------------------------------

# https://llvm.org/docs/LangRef.html#catchswitch-instruction

# ref: ParseCatchSwitch
#
#   ::= 'catchswitch' within Parent

CatchSwitchTerm
	: 'catchswitch' 'within' ExceptionScope '[' LabelList ']' 'unwind' UnwindTarget InstructionMetadata
;

# --- [ catchret ] -------------------------------------------------------------

# https://llvm.org/docs/LangRef.html#catchret-instruction

# ref: ParseCatchRet
#
#   ::= 'catchret' from Parent Value 'to' TypeAndValue

CatchRetTerm
	: 'catchret' 'from' Value 'to' LabelType LocalIdent InstructionMetadata
;

# --- [ cleanupret ] -----------------------------------------------------------

# https://llvm.org/docs/LangRef.html#cleanupret-instruction

# ref: ParseCleanupRet
#
#   ::= 'cleanupret' from Value unwind ('to' 'caller' | TypeAndValue)

CleanupRetTerm
	: 'cleanupret' 'from' Value 'unwind' UnwindTarget InstructionMetadata
;

# --- [ unreachable ] ----------------------------------------------------------

# https://llvm.org/docs/LangRef.html#unreachable-instruction

# ref: ParseInstruction

UnreachableTerm
	: 'unreachable' InstructionMetadata
;

# ___ [ Helpers ] _____________________________________________________________

UnwindTarget
	: 'to' 'caller'
	| LabelType LocalIdent
;

# === [ Metadata Nodes and Metadata Strings ] ==================================

# https://llvm.org/docs/LangRef.html#metadata-nodes-and-metadata-strings

# --- [ Metadata Tuple ] -------------------------------------------------------

# ref: ParseMDTuple

MDTuple
	: '!' MDFields
;

# ref: ParseMDNodeVector
#
#   ::= { Element (',' Element)* }
#  Element
#   ::= 'null' | TypeAndValue

# ref: ParseMDField(MDFieldList &)

MDFields
	: '{' '}'
	| '{' MDFieldList '}'
;

MDFieldList
	: MDField
	| MDFieldList ',' MDField
;

# ref: ParseMDField(MDField &)

MDField
	# Null is a special case since it is typeless.
	: 'null'
	| Metadata
;

# --- [ Metadata ] -------------------------------------------------------------

# ref: ParseMetadata
#
#  ::= i32 %local
#  ::= i32 @global
#  ::= i32 7
#  ::= !42
#  ::= !{...}
#  ::= !'string'
#  ::= !DILocation(...)

Metadata
	: Type Value
	| MDString
	# !{ ... }
	| MDTuple
	# !7
	| MetadataID
	| SpecializedMDNode
;

# --- [ Metadata String ] ------------------------------------------------------

# ref: ParseMDString
#
#   ::= '!' STRINGCONSTANT

MDString
	: '!' StringLit
;

# --- [ Metadata Attachment ] --------------------------------------------------

# ref: ParseMetadataAttachment
#
#   ::= !dbg !42

MetadataAttachment
	: MetadataName MDNode
;

# --- [ Metadata Node ] --------------------------------------------------------

# ref: ParseMDNode
#
#  ::= !{ ... }
#  ::= !7
#  ::= !DILocation(...)

MDNode
	# !{ ... }
	: MDTuple
	# !42
	| MetadataID
	| SpecializedMDNode
;

# ### [ Helper productions ] ##################################################

# ref: ParseOptionalFunctionMetadata
#
#   ::= (!dbg !57)*

FunctionMetadata
	: MetadataAttachment*
;

# --- [ Specialized Metadata Nodes ] -------------------------------------------

# https://llvm.org/docs/LangRef.html#specialized-metadata-nodes

# ref: ParseSpecializedMDNode

SpecializedMDNode
	: DICompileUnit
	| DIFile
	| DIBasicType
	| DISubroutineType
	| DIDerivedType
	| DICompositeType
	| DISubrange
	| DIEnumerator
	| DITemplateTypeParameter
	| DITemplateValueParameter
	| DIModule # not in spec as of 2018-02-21
	| DINamespace
	| DIGlobalVariable
	| DISubprogram
	| DILexicalBlock
	| DILexicalBlockFile
	| DILocation
	| DILocalVariable
	| DIExpression
	| DIGlobalVariableExpression # not in spec as of 2018-02-21
	| DIObjCProperty
	| DIImportedEntity
	| DIMacro
	| DIMacroFile
	| GenericDINode # not in spec as of 2018-02-21
;

# ~~~ [ DICompileUnit ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dicompileunit

# ref: ParseDICompileUnit
#
#   ::= !DICompileUnit(language: DW_LANG_C99, file: !0, producer: 'clang',
#                      isOptimized: true, flags: '-O2', runtimeVersion: 1,
#                      splitDebugFilename: 'abc.debug',
#                      emissionKind: FullDebug, enums: !1, retainedTypes: !2,
#                      globals: !4, imports: !5, macros: !6, dwoId: 0x0abcd)
#
#  REQUIRED(language, DwarfLangField, );
#  REQUIRED(file, MDField, (AllowNull false));
#  OPTIONAL(producer, MDStringField, );
#  OPTIONAL(isOptimized, MDBoolField, );
#  OPTIONAL(flags, MDStringField, );
#  OPTIONAL(runtimeVersion, MDUnsignedField, (0, UINT32_MAX));
#  OPTIONAL(splitDebugFilename, MDStringField, );
#  OPTIONAL(emissionKind, EmissionKindField, );
#  OPTIONAL(enums, MDField, );
#  OPTIONAL(retainedTypes, MDField, );
#  OPTIONAL(globals, MDField, );
#  OPTIONAL(imports, MDField, );
#  OPTIONAL(macros, MDField, );
#  OPTIONAL(dwoId, MDUnsignedField, );
#  OPTIONAL(splitDebugInlining, MDBoolField, = true);
#  OPTIONAL(debugInfoForProfiling, MDBoolField, = false);
#  OPTIONAL(gnuPubnames, MDBoolField, = false);

DICompileUnit
	: '!DICompileUnit' '(' (DICompileUnitField separator ',')* ')'
;

DICompileUnitField
	: 'language:' DwarfLang
	| FileField
	| 'producer:' StringLit
	| IsOptimizedField
	| 'flags:' StringLit
	| 'runtimeVersion:' IntLit
	| 'splitDebugFilename:' StringLit
	| 'emissionKind:' EmissionKind
	| 'enums:' MDField
	| 'retainedTypes:' MDField
	| 'globals:' MDField
	| 'imports:' MDField
	| 'macros:' MDField
	| 'dwoId:' IntLit
	| 'splitDebugInlining:' BoolLit
	| 'debugInfoForProfiling:' BoolLit
	| 'gnuPubnames:' BoolLit
;

# ~~~ [ DIFile ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#difile

# ref: ParseDIFileType
#
#   ::= !DIFileType(filename: 'path/to/file', directory: '/path/to/dir'
#                   checksumkind: CSK_MD5,
#                   checksum: '000102030405060708090a0b0c0d0e0f')
#
#  REQUIRED(filename, MDStringField, );
#  REQUIRED(directory, MDStringField, );
#  OPTIONAL(checksumkind, ChecksumKindField, (DIFile::CSK_MD5));
#  OPTIONAL(checksum, MDStringField, );

DIFile
	: '!DIFile' '(' (DIFileField separator ',')* ')'
;

DIFileField
	: 'filename:' StringLit
	| 'directory:' StringLit
	| 'checksumkind:' ChecksumKind
	| 'checksum:' StringLit
;

# ~~~ [ DIBasicType ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dibasictype

# ref: ParseDIBasicType
#
#   ::= !DIBasicType(tag: DW_TAG_base_type, name: 'int', size: 32, align: 32)
#
#  OPTIONAL(tag, DwarfTagField, (dwarf::DW_TAG_base_type));
#  OPTIONAL(name, MDStringField, );
#  OPTIONAL(size, MDUnsignedField, (0, UINT64_MAX));
#  OPTIONAL(align, MDUnsignedField, (0, UINT32_MAX));
#  OPTIONAL(encoding, DwarfAttEncodingField, );

DIBasicType
	: '!DIBasicType' '(' (DIBasicTypeField separator ',')* ')'
;

DIBasicTypeField
	: TagField
	| NameField
	| SizeField
	| AlignField
	| 'encoding:' DwarfAttEncoding
;

# ~~~ [ DISubroutineType ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#disubroutinetype

# ref: ParseDISubroutineType
#
#  OPTIONAL(flags, DIFlagField, );
#  OPTIONAL(cc, DwarfCCField, );
#  REQUIRED(types, MDField, );

DISubroutineType
	: '!DISubroutineType' '(' (DISubroutineTypeField separator ',')* ')'
;

DISubroutineTypeField
	: FlagsField
	| 'cc:' DwarfCC
	| 'types:' MDField
;

# ~~~ [ DIDerivedType ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#diderivedtype

# ref: ParseDIDerivedType
#
#   ::= !DIDerivedType(tag: DW_TAG_pointer_type, name: 'int', file: !0,
#                      line: 7, scope: !1, baseType: !2, size: 32,
#                      align: 32, offset: 0, flags: 0, extraData: !3,
#                      dwarfAddressSpace: 3)
#
#  REQUIRED(tag, DwarfTagField, );
#  OPTIONAL(name, MDStringField, );
#  OPTIONAL(scope, MDField, );
#  OPTIONAL(file, MDField, );
#  OPTIONAL(line, LineField, );
#  REQUIRED(baseType, MDField, );
#  OPTIONAL(size, MDUnsignedField, (0, UINT64_MAX));
#  OPTIONAL(align, MDUnsignedField, (0, UINT32_MAX));
#  OPTIONAL(offset, MDUnsignedField, (0, UINT64_MAX));
#  OPTIONAL(flags, DIFlagField, );
#  OPTIONAL(extraData, MDField, );
#  OPTIONAL(dwarfAddressSpace, MDUnsignedField, (UINT32_MAX, UINT32_MAX));

DIDerivedType
	: '!DIDerivedType' '(' (DIDerivedTypeField separator ',')* ')'
;

DIDerivedTypeField
	: TagField
	| NameField
	| ScopeField
	| FileField
	| LineField
	| BaseTypeField
	| SizeField
	| AlignField
	| OffsetField
	| FlagsField
	| 'extraData:' MDField
	| 'dwarfAddressSpace:' IntLit
;

# ~~~ [ DICompositeType ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dicompositetype

# ref: ParseDICompositeType
#
#  REQUIRED(tag, DwarfTagField, );
#  OPTIONAL(name, MDStringField, );
#  OPTIONAL(scope, MDField, );
#  OPTIONAL(file, MDField, );
#  OPTIONAL(line, LineField, );
#  OPTIONAL(baseType, MDField, );
#  OPTIONAL(size, MDUnsignedField, (0, UINT64_MAX));
#  OPTIONAL(align, MDUnsignedField, (0, UINT32_MAX));
#  OPTIONAL(offset, MDUnsignedField, (0, UINT64_MAX));
#  OPTIONAL(flags, DIFlagField, );
#  OPTIONAL(elements, MDField, );
#  OPTIONAL(runtimeLang, DwarfLangField, );
#  OPTIONAL(vtableHolder, MDField, );
#  OPTIONAL(templateParams, MDField, );
#  OPTIONAL(identifier, MDStringField, );
#  OPTIONAL(discriminator, MDField, );

DICompositeType
	: '!DICompositeType' '(' (DICompositeTypeField separator ',')* ')'
;

DICompositeTypeField
	: TagField
	| NameField
	| ScopeField
	| FileField
	| LineField
	| BaseTypeField
	| SizeField
	| AlignField
	| OffsetField
	| FlagsField
	| 'elements:' MDField
	| 'runtimeLang:' DwarfLang
	| 'vtableHolder:' MDField
	| TemplateParamsField
	| 'identifier:' StringLit
	| 'discriminator:' MDField
;

# ~~~ [ DISubrange ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#disubrange

# ref: ParseDISubrange
#
#   ::= !DISubrange(count: 30, lowerBound: 2)
#   ::= !DISubrange(count: !node, lowerBound: 2)
#
#  REQUIRED(count, MDSignedOrMDField, (-1, -1, INT64_MAX, false));
#  OPTIONAL(lowerBound, MDSignedField, );

DISubrange
	: '!DISubrange' '(' (DISubrangeField separator ',')* ')'
;

DISubrangeField
	: 'count:' IntOrMDField
	| 'lowerBound:' IntLit
;

# ~~~ [ DIEnumerator ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dienumerator

# ref: ParseDIEnumerator
#
#   ::= !DIEnumerator(value: 30, isUnsigned: true, name: 'SomeKind')
#
#  REQUIRED(name, MDStringField, );
#  REQUIRED(value, MDSignedOrUnsignedField, );
#  OPTIONAL(isUnsigned, MDBoolField, (false));

DIEnumerator
	: '!DIEnumerator' '(' (DIEnumeratorField separator ',')* ')'
;

DIEnumeratorField
	: NameField
	| 'value:' IntLit
	| 'isUnsigned:' BoolLit
;

# ~~~ [ DITemplateTypeParameter ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#ditemplatetypeparameter

# ref: ParseDITemplateTypeParameter
#
#   ::= !DITemplateTypeParameter(name: 'Ty', type: !1)
#
#  OPTIONAL(name, MDStringField, );
#  REQUIRED(type, MDField, );

DITemplateTypeParameter
	: '!DITemplateTypeParameter' '(' (DITemplateTypeParameterField separator ',')* ')'
;

DITemplateTypeParameterField
	: NameField
	| TypeField
;

# ~~~ [ DITemplateValueParameter ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#ditemplatevalueparameter

# ref: ParseDITemplateValueParameter
#
#   ::= !DITemplateValueParameter(tag: DW_TAG_template_value_parameter,
#                                 name: 'V', type: !1, value: i32 7)
#
#  OPTIONAL(tag, DwarfTagField, (dwarf::DW_TAG_template_value_parameter));
#  OPTIONAL(name, MDStringField, );
#  OPTIONAL(type, MDField, );
#  REQUIRED(value, MDField, );

DITemplateValueParameter
	: '!DITemplateValueParameter' '(' (DITemplateValueParameterField separator ',')* ')'
;

DITemplateValueParameterField
	: TagField
	| NameField
	| TypeField
	| 'value:' MDField
;

# ~~~ [ DIModule ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseDIModule
#
#   ::= !DIModule(scope: !0, name: 'SomeModule', configMacros: '-DNDEBUG',
#                 includePath: '/usr/include', isysroot: '/')
#
#  REQUIRED(scope, MDField, );
#  REQUIRED(name, MDStringField, );
#  OPTIONAL(configMacros, MDStringField, );
#  OPTIONAL(includePath, MDStringField, );
#  OPTIONAL(isysroot, MDStringField, );

DIModule
	: '!DIModule' '(' (DIModuleField separator ',')* ')'
;

DIModuleField
	: ScopeField
	| NameField
	| 'configMacros:' StringLit
	| 'includePath:' StringLit
	| 'isysroot:' StringLit
;

# ~~~ [ DINamespace ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dinamespace

# ref: ParseDINamespace
#
#   ::= !DINamespace(scope: !0, file: !2, name: 'SomeNamespace', line: 9)
#
#  REQUIRED(scope, MDField, );
#  OPTIONAL(name, MDStringField, );
#  OPTIONAL(exportSymbols, MDBoolField, );

DINamespace
	: '!DINamespace' '(' (DINamespaceField separator ',')* ')'
;

DINamespaceField
	: ScopeField
	| NameField
	| 'exportSymbols:' BoolLit
;

# ~~~ [ DIGlobalVariable ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#diglobalvariable

# ref: ParseDIGlobalVariable
#
#   ::= !DIGlobalVariable(scope: !0, name: 'foo', linkageName: 'foo',
#                         file: !1, line: 7, type: !2, isLocal: false,
#                         isDefinition: true, declaration: !3, align: 8)
#
#  REQUIRED(name, MDStringField, (AllowEmpty false));
#  OPTIONAL(scope, MDField, );
#  OPTIONAL(linkageName, MDStringField, );
#  OPTIONAL(file, MDField, );
#  OPTIONAL(line, LineField, );
#  OPTIONAL(type, MDField, );
#  OPTIONAL(isLocal, MDBoolField, );
#  OPTIONAL(isDefinition, MDBoolField, (true));
#  OPTIONAL(declaration, MDField, );
#  OPTIONAL(align, MDUnsignedField, (0, UINT32_MAX));

DIGlobalVariable
	: '!DIGlobalVariable' '(' (DIGlobalVariableField separator ',')* ')'
;

DIGlobalVariableField
	: NameField
	| ScopeField
	| LinkageNameField
	| FileField
	| LineField
	| TypeField
	| IsLocalField
	| IsDefinitionField
	| DeclarationField
	| AlignField
;

# ~~~ [ DISubprogram ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#disubprogram

# ref: ParseDISubprogram
#
#   ::= !DISubprogram(scope: !0, name: 'foo', linkageName: '_Zfoo',
#                     file: !1, line: 7, type: !2, isLocal: false,
#                     isDefinition: true, scopeLine: 8, containingType: !3,
#                     virtuality: DW_VIRTUALTIY_pure_virtual,
#                     virtualIndex: 10, thisAdjustment: 4, flags: 11,
#                     isOptimized: false, templateParams: !4, declaration: !5,
#                     variables: !6, thrownTypes: !7)
#
#  OPTIONAL(name, MDStringField, );
#  OPTIONAL(scope, MDField, );
#  OPTIONAL(linkageName, MDStringField, );
#  OPTIONAL(file, MDField, );
#  OPTIONAL(line, LineField, );
#  OPTIONAL(type, MDField, );
#  OPTIONAL(isLocal, MDBoolField, );
#  OPTIONAL(isDefinition, MDBoolField, (true));
#  OPTIONAL(scopeLine, LineField, );
#  OPTIONAL(containingType, MDField, );
#  OPTIONAL(virtuality, DwarfVirtualityField, );
#  OPTIONAL(virtualIndex, MDUnsignedField, (0, UINT32_MAX));
#  OPTIONAL(thisAdjustment, MDSignedField, (0, INT32_MIN, INT32_MAX));
#  OPTIONAL(flags, DIFlagField, );
#  OPTIONAL(isOptimized, MDBoolField, );
#  OPTIONAL(unit, MDField, );
#  OPTIONAL(templateParams, MDField, );
#  OPTIONAL(declaration, MDField, );
#  OPTIONAL(variables, MDField, );
#  OPTIONAL(thrownTypes, MDField, );

DISubprogram
	: '!DISubprogram' '(' (DISubprogramField separator ',')* ')'
;

DISubprogramField
	: NameField
	| ScopeField
	| LinkageNameField
	| FileField
	| LineField
	| TypeField
	| IsLocalField
	| IsDefinitionField
	| 'scopeLine:' IntLit
	| 'containingType:' MDField
	| 'virtuality:' DwarfVirtuality
	| 'virtualIndex:' IntLit
	| 'thisAdjustment:' IntLit
	| FlagsField
	| IsOptimizedField
	| 'unit:' MDField
	| TemplateParamsField
	| DeclarationField
	| 'variables:' MDField
	| 'thrownTypes:' MDField
;

# ~~~ [ DILexicalBlock ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dilexicalblock

# ref: ParseDILexicalBlock
#
#   ::= !DILexicalBlock(scope: !0, file: !2, line: 7, column: 9)
#
#  REQUIRED(scope, MDField, (AllowNull false));
#  OPTIONAL(file, MDField, );
#  OPTIONAL(line, LineField, );
#  OPTIONAL(column, ColumnField, );

DILexicalBlock
	: '!DILexicalBlock' '(' (DILexicalBlockField separator ',')* ')'
;

DILexicalBlockField
	: ScopeField
	| FileField
	| LineField
	| ColumnField
;

# ~~~ [ DILexicalBlockFile ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dilexicalblockfile

# ref: ParseDILexicalBlockFile
#
#   ::= !DILexicalBlockFile(scope: !0, file: !2, discriminator: 9)
#
#  REQUIRED(scope, MDField, (AllowNull false));
#  OPTIONAL(file, MDField, );
#  REQUIRED(discriminator, MDUnsignedField, (0, UINT32_MAX));

DILexicalBlockFile
	: '!DILexicalBlockFile' '(' (DILexicalBlockFileField separator ',')* ')'
;

DILexicalBlockFileField
	: ScopeField
	| FileField
	| 'discriminator:' IntLit
;

# ~~~ [ DILocation ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dilocation

# ref: ParseDILocation
#
#   ::= !DILocation(line: 43, column: 8, scope: !5, inlinedAt: !6)
#
#  OPTIONAL(line, LineField, );
#  OPTIONAL(column, ColumnField, );
#  REQUIRED(scope, MDField, (AllowNull false));
#  OPTIONAL(inlinedAt, MDField, );

DILocation
	: '!DILocation' '(' (DILocationField separator ',')* ')'
;

DILocationField
	: LineField
	| ColumnField
	| ScopeField
	| 'inlinedAt:' MDField
;

# ~~~ [ DILocalVariable ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dilocalvariable

# ref: ParseDILocalVariable
#
#   ::= !DILocalVariable(arg: 7, scope: !0, name: 'foo',
#                        file: !1, line: 7, type: !2, arg: 2, flags: 7,
#                        align: 8)
#   ::= !DILocalVariable(scope: !0, name: 'foo',
#                        file: !1, line: 7, type: !2, arg: 2, flags: 7,
#                        align: 8)
#
#  OPTIONAL(name, MDStringField, );
#  OPTIONAL(arg, MDUnsignedField, (0, UINT16_MAX));
#  REQUIRED(scope, MDField, (AllowNull false));
#  OPTIONAL(file, MDField, );
#  OPTIONAL(line, LineField, );
#  OPTIONAL(type, MDField, );
#  OPTIONAL(flags, DIFlagField, );
#  OPTIONAL(align, MDUnsignedField, (0, UINT32_MAX));

DILocalVariable
	: '!DILocalVariable' '(' (DILocalVariableField separator ',')* ')'
;

DILocalVariableField
	: NameField
	| 'arg:' IntLit
	| ScopeField
	| FileField
	| LineField
	| TypeField
	| FlagsField
	| AlignField
;

# ~~~ [ DIExpression ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#diexpression

# ref: ParseDIExpression
#
#   ::= !DIExpression(0, 7, -1)

DIExpression
	: '!DIExpression' '(' (DIExpressionField separator ',')* ')'
;

DIExpressionField
	: int_lit_tok
	| DwarfOp
;

# ~~~ [ DIGlobalVariableExpression ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseDIGlobalVariableExpression
#
#   ::= !DIGlobalVariableExpression(var: !0, expr: !1)
#
#  REQUIRED(var, MDField, );
#  REQUIRED(expr, MDField, );

DIGlobalVariableExpression
	: '!DIGlobalVariableExpression' '(' (DIGlobalVariableExpressionField separator ',')* ')'
;

DIGlobalVariableExpressionField
	: 'var:' MDField
	| 'expr:' MDField
;

# ~~~ [ DIObjCProperty ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#diobjcproperty

# ref: ParseDIObjCProperty
#
#   ::= !DIObjCProperty(name: 'foo', file: !1, line: 7, setter: 'setFoo',
#                       getter: 'getFoo', attributes: 7, type: !2)
#
#  OPTIONAL(name, MDStringField, );
#  OPTIONAL(file, MDField, );
#  OPTIONAL(line, LineField, );
#  OPTIONAL(setter, MDStringField, );
#  OPTIONAL(getter, MDStringField, );
#  OPTIONAL(attributes, MDUnsignedField, (0, UINT32_MAX));
#  OPTIONAL(type, MDField, );

DIObjCProperty
	: '!DIObjCProperty' '(' (DIObjCPropertyField separator ',')* ')'
;

DIObjCPropertyField
	: NameField
	| FileField
	| LineField
	| 'setter:' StringLit
	| 'getter:' StringLit
	| 'attributes:' IntLit
	| TypeField
;

# ~~~ [ DIImportedEntity ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#diimportedentity

# ref: ParseDIImportedEntity
#
#   ::= !DIImportedEntity(tag: DW_TAG_imported_module, scope: !0, entity: !1,
#                         line: 7, name: 'foo')
#
#  REQUIRED(tag, DwarfTagField, );
#  REQUIRED(scope, MDField, );
#  OPTIONAL(entity, MDField, );
#  OPTIONAL(file, MDField, );
#  OPTIONAL(line, LineField, );
#  OPTIONAL(name, MDStringField, );

DIImportedEntity
	: '!DIImportedEntity' '(' (DIImportedEntityField separator ',')* ')'
;

DIImportedEntityField
	: TagField
	| ScopeField
	| 'entity:' MDField
	| FileField
	| LineField
	| NameField
;

# ~~~ [ DIMacro ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dimacro

# ref: ParseDIMacro
#
#   ::= !DIMacro(macinfo: type, line: 9, name: 'SomeMacro', value: 'SomeValue')
#
#  REQUIRED(type, DwarfMacinfoTypeField, );
#  OPTIONAL(line, LineField, );
#  REQUIRED(name, MDStringField, );
#  OPTIONAL(value, MDStringField, );

DIMacro
	: '!DIMacro' '(' (DIMacroField separator ',')* ')'
;

DIMacroField
	: TypeMacinfoField
	| LineField
	| NameField
	| 'value:' StringLit
;

# ~~~ [ DIMacroFile ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dimacrofile

# ref: ParseDIMacroFile
#
#   ::= !DIMacroFile(line: 9, file: !2, nodes: !3)
#
#  OPTIONAL(type, DwarfMacinfoTypeField, (dwarf::DW_MACINFO_start_file));
#  OPTIONAL(line, LineField, );
#  REQUIRED(file, MDField, );
#  OPTIONAL(nodes, MDField, );

DIMacroFile
	: '!DIMacroFile' '(' (DIMacroFileField separator ',')* ')'
;

DIMacroFileField
	: TypeMacinfoField
	| LineField
	| FileField
	| 'nodes:' MDField
;

# ~~~ [ GenericDINode ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseGenericDINode
#
#   ::= !GenericDINode(tag: 15, header: '...', operands: {...})
#
#  REQUIRED(tag, DwarfTagField, );
#  OPTIONAL(header, MDStringField, );
#  OPTIONAL(operands, MDFieldList, );

GenericDINode
	: '!GenericDINode' '(' (GenericDINodeField separator ',')* ')'
;

GenericDINodeField
	: TagField
	| 'header:' StringLit
	| 'operands:' MDFields
;

# ### [ Helper productions ] ###################################################

FileField
	: 'file:' MDField
;

IsOptimizedField
	: 'isOptimized:' BoolLit
;

TagField
	: 'tag:' DwarfTag
;

NameField
	: 'name:' StringLit
;

SizeField
	: 'size:' IntLit
;

AlignField
	: 'align:' IntLit
;

FlagsField
	: 'flags:' DIFlagList
;

LineField
	: 'line:' IntLit
;

ScopeField
	: 'scope:' MDField
;

BaseTypeField
	: 'baseType:' MDField
;

OffsetField
	: 'offset:' IntLit
;

TemplateParamsField
	: 'templateParams:' MDField
;

# ref: ParseMDField(MDSignedOrMDField &)

IntOrMDField
	: int_lit_tok
	| MDField
;

TypeField
	: 'type:' MDField
;

LinkageNameField
	: 'linkageName:' StringLit
;

IsLocalField
	: 'isLocal:' BoolLit
;

IsDefinitionField
	: 'isDefinition:' BoolLit
;

DeclarationField
	: 'declaration:' MDField
;

ColumnField
	: 'column:' IntLit
;

TypeMacinfoField
	: 'type:' DwarfMacinfo
;

ChecksumKind
	# CSK_foo
	: checksum_kind_tok
;

# ref: ParseMDField(DIFlagField &)
#
#  ::= uint32
#  ::= DIFlagVector
#  ::= DIFlagVector '|' DIFlagFwdDecl '|' uint32 '|' DIFlagPublic

DIFlagList
	: DIFlag
	| DIFlagList '|' DIFlag
;

DIFlag
	: IntLit
	# DIFlagFoo
	| di_flag_tok
;

# ref: ParseMDField(DwarfAttEncodingField &)

DwarfAttEncoding
	: IntLit
	# DW_ATE_foo
	| dwarf_att_encoding_tok
;

# ref: ParseMDField(DwarfCCField &Result)

DwarfCC
	: IntLit
	# DW_CC_foo
	| dwarf_cc_tok
;

# ref: ParseMDField(DwarfLangField &)

DwarfLang
	: IntLit
	# DW_LANG_foo
	| dwarf_lang_tok
;

# ref: ParseMDField(DwarfMacinfoTypeField &)

DwarfMacinfo
	: IntLit
	# DW_MACINFO_foo
	| dwarf_macinfo_tok
;

DwarfOp
	# DW_OP_foo
	: dwarf_op_tok
;

# ref: ParseMDField(DwarfTagField &)

DwarfTag
	: IntLit
	# DW_TAG_foo
	| dwarf_tag_tok
;

# ref: ParseMDField(DwarfVirtualityField &)

DwarfVirtuality
	: IntLit
	# DW_VIRTUALITY_foo
	| dwarf_virtuality_tok
;

EmissionKind
	: IntLit
	| 'FullDebug'
	| 'LineTablesOnly'
	| 'NoDebug'
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
	: 'addrspace' '(' UintLit ')'
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
	: 'align' UintLit
;

# ___ [ Function Attribute ] ___________________________________________________

# ref: ParseFnAttributeValuePairs
#
#   ::= <attr> | <attr> '=' <value>

# NOTE: FuncAttr should contain Alignment. However, using LALR(1) this
# produces a reduce/reduce conflict as GlobalAttr also contains Alignment.
#
# To handle these ambiguities, (FuncAttr | Alignment) is used in those places
# where FuncAttr is used outside of GlobalDef and GlobalDecl (which alos has
# GlobalAttr).

FuncAttr
	# not used in attribute groups.
	: AttrGroupID
	# used in attribute groups.
	| 'align' '=' UintLit
	| 'alignstack' '=' UintLit
	# used in functions.
	#| Alignment # NOTE: removed to resolve reduce/reduce conflict, see above.
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
	: 'allocsize' '(' UintLit ')'
	| 'allocsize' '(' UintLit ',' UintLit ')'
;

# ref: ParseOptionalStackAlignment
#
#   ::= empty
#   ::= 'alignstack' '(' 4 ')'
StackAlignment
	: 'alignstack' '(' UintLit ')'
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
	| 'cc' UintLit
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
	: 'dereferenceable' '(' UintLit ')'
	| 'dereferenceable_or_null' '(' UintLit ')'
;

Exact
	: 'exact'
;

OverflowFlags
	: ('nsw' | 'nuw')*
;

InBounds
	: 'inbounds'
;

# ref: ParseIndexList
#
#    ::=  (',' uint32)+

Indices
	: (UintLit separator ',')*
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

# ref: ParseInstructionMetadata
#
#   ::= !dbg !42 (',' !dbg !57)*

InstructionMetadata
   : (',' MetadataAttachment)+?
;

# ref: EatFastMathFlagsIfPresent

FastMathFlag
	: 'afn'
	| 'arcp'
	| 'contract'
	| 'fast'
	| 'ninf'
	| 'nnan'
	| 'nsz'
	| 'reassoc'
;

Volatile
	: 'volatile'
;

# ref: ParseScope
#
#   ::= syncscope("singlethread" | "<target scope>")?

SyncScope
	: 'syncscope' '(' StringLit ')'
;

# ref: ParseOrdering
#
#   ::= AtomicOrdering

AtomicOrdering
	: 'acq_rel'
	| 'acquire'
	| 'monotonic'
	| 'release'
	| 'seq_cst'
	| 'unordered'
;

# ref: ParseParameterList
#
#    ::= '(' ')'
#    ::= '(' Arg (',' Arg)* ')'
#  Arg
#    ::= Type OptionalAttributes Value OptionalAttributes

Args
	: '...'?
	| (Arg separator ',')+ (',' '...')?
;

# ref: ParseMetadataAsValue
#
#  ::= metadata i32 %local
#  ::= metadata i32 @global
#  ::= metadata i32 7
#  ::= metadata !0
#  ::= metadata !{...}
#  ::= metadata !"string"

Arg
	: ConcreteType ParamAttr* Value
	| MetadataType Metadata
;

# ref: ParseExceptionArgs

ExceptionArg
	: ConcreteType Value
	| MetadataType Metadata
;

# ref: ParseOptionalOperandBundles
#
#    ::= empty
#    ::= '[' OperandBundle [, OperandBundle ]* ']'
#
#  OperandBundle
#    ::= bundle-tag '(' ')'
#    ::= bundle-tag '(' Type Value [, Type Value ]* ')'
#
#  bundle-tag ::= String Constant

OperandBundles
	: ('[' (OperandBundle separator ',')+ ']')?
;

OperandBundle
	: StringLit '(' (Type Value separator ',')* ')'
;

ExceptionScope
	: NoneConst
	| LocalIdent
;
