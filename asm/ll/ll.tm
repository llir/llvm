language llvm(go);

lang = "llvm"
package = "github.com/mewmew/l-tm/asm/ll"
eventBased = true
eventFields = true

# TODO: check when to use Fooopt and when to use Foo? (as based on the AST
# they produce)

# ### [ Lexical part ] #########################################################

:: lexer

# TODO: remove placeholders.

placeholder1 : /placeholder1/
placeholder2 : /placeholder2/
placeholder3 : /placeholder3/

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

# ref: DWKEYWORD

# GNU
NameTableKind : /(GNU)|(None)|(Default)/

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
'!DILabel' : /!DILabel/
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
'aarch64_vector_pcs' : /aarch64_vector_pcs/
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
'isImplicitCode:' : /isImplicitCode:/
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
'nameTableKind:' : /nameTableKind:/
'nand' : /nand/
'ne' : /ne/
'nest' : /nest/
'ninf' : /ninf/
'nnan' : /nnan/
'noalias' : /noalias/
'nobuiltin' : /nobuiltin/
'nocapture' : /nocapture/
'nocf_check' : /nocf_check/
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
'optforfuzzing' : /optforfuzzing/
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
'retainedNodes:' : /retainedNodes:/
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
'shadowcallstack' : /shadowcallstack/
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
'source:' : /source:/
'speculatable' : /speculatable/
'speculative_load_hardening' : /speculative_load_hardening/
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
# project, version 7.0

:: parser

%input Module;

# === [ Identifiers ] ==========================================================

# --- [ Global Identifiers ] ---------------------------------------------------

GlobalIdent -> GlobalIdent
	: global_ident_tok
;

# --- [ Local Identifiers ] ----------------------------------------------------

LocalIdent -> LocalIdent
	: local_ident_tok
;

# --- [ Label Identifiers ] ----------------------------------------------------

LabelIdent -> LabelIdent
	: label_ident_tok
;

# --- [ Attribute Group Identifiers ] ------------------------------------------

AttrGroupID -> AttrGroupID
	: attr_group_id_tok
;

# --- [ Comdat Identifiers ] ---------------------------------------------------

ComdatName -> ComdatName
	: comdat_name_tok
;

# --- [ Metadata Identifiers ] -------------------------------------------------

MetadataName -> MetadataName
	: metadata_name_tok
;

MetadataID -> MetadataID
	: metadata_id_tok
;

# === [ Literals ] =============================================================

# --- [ Integer literals ] -----------------------------------------------------

BoolLit -> BoolLit
	: 'true'
	| 'false'
;

IntLit -> IntLit
	: int_lit_tok
;

UintLit -> UintLit
	: int_lit_tok
;

# --- [ Floating-point literals ] ----------------------------------------------

FloatLit -> FloatLit
	: float_lit_tok
;

# --- [ String literals ] ------------------------------------------------------

StringLit -> StringLit
	: string_lit_tok
;

# === [ Module ] ===============================================================

# https://llvm.org/docs/LangRef.html#module-structure

# ref: Run
#
#   module ::= toplevelentity*

Module -> Module
	: TopLevelEntities=TopLevelEntity*
;

# --- [ Top-level Entities ] ---------------------------------------------------

# ref: ParseTopLevelEntities

%interface TopLevelEntity;

TopLevelEntity -> TopLevelEntity
	: SourceFilename
	| TargetDef
	| ModuleAsm
	| TypeDef
	| ComdatDef
	| GlobalDef
	| IndirectSymbolDef
	| FuncDecl
	| FuncDef
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

SourceFilename -> SourceFilename
	: 'source_filename' '=' Name=StringLit
;

# ~~~ [ Target Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#target-triple
# https://llvm.org/docs/LangRef.html#data-layout

# ref: ParseTargetDefinition
#
#   ::= 'target' 'triple' '=' STRINGCONSTANT
#   ::= 'target' 'datalayout' '=' STRINGCONSTANT

%interface TargetDef;

TargetDef -> TargetDef
	: TargetDataLayout
	| TargetTriple
;

TargetDataLayout -> TargetDataLayout
	: 'target' 'datalayout' '=' DataLayout=StringLit
;

TargetTriple -> TargetTriple
	: 'target' 'triple' '=' TargetTriple=StringLit
;

# ~~~ [ Module-level Inline Assembly ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#module-level-inline-assembly

# ref: ParseModuleAsm
#
#   ::= 'module' 'asm' STRINGCONSTANT

ModuleAsm -> ModuleAsm
	: 'module' 'asm' Asm=StringLit
;

# ~~~ [ Type Defintion ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#structure-type

# ref: ParseUnnamedType
#
#   ::= LocalVarID '=' 'type' type

# ref: ParseNamedType
#
#   ::= LocalVar '=' 'type' type

# TODO: Remove `Typ=` once https://github.com/inspirer/textmapper/issues/13
# is resolved.

TypeDef -> TypeDef
	: Alias=LocalIdent '=' 'type' Typ=OpaqueType
	| Alias=LocalIdent '=' 'type' Typ=Type
;

# ~~~ [ Comdat Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#langref-comdats

# ref: parseComdat

ComdatDef -> ComdatDef
	: Name=ComdatName '=' 'comdat' Kind=SelectionKind
;

SelectionKind -> SelectionKind
	: 'any'
	| 'exactmatch'
	| 'largest'
	| 'noduplicates'
	| 'samesize'
;

# ~~~ [ Global Variable Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

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

GlobalDef -> GlobalDef
	: Name=GlobalIdent '=' Linkageopt PreemptionSpecifieropt Visibilityopt DLLStorageClassopt ThreadLocalopt UnnamedAddropt AddrSpaceopt ExternallyInitializedopt Immutable Typ=Type Init=Constantopt GlobalAttrs=(',' GlobalAttr)+? FuncAttrs=(',' FuncAttr)+?
;

# TODO: Check if ExternallyInitialized can be inlined or handled in a cleaner way. ref: https://github.com/inspirer/textmapper/issues/14

ExternallyInitialized -> ExternallyInitialized
	: 'externally_initialized'
;

# ref: ParseGlobalType
#
#   ::= 'constant'
#   ::= 'global'

# TODO: Check if Immutable can be inlined or handled in a cleaner way. ref: https://github.com/inspirer/textmapper/issues/14

Immutable -> Immutable
	: 'constant'
	| 'global'
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

%interface IndirectSymbolDef;

IndirectSymbolDef -> IndirectSymbolDef
	: AliasDef
	| IFuncDef
;

AliasDef -> AliasDef
	: Name=GlobalIdent '=' Linkageopt PreemptionSpecifieropt Visibilityopt DLLStorageClassopt ThreadLocalopt UnnamedAddropt 'alias' Typ=Type ',' AliaseeType=Type Aliasee=Constant
;

IFuncDef -> IFuncDef
	: Name=GlobalIdent '=' Linkageopt PreemptionSpecifieropt Visibilityopt DLLStorageClassopt ThreadLocalopt UnnamedAddropt 'ifunc' Typ=Type ',' ResolverType=Type Resolver=Constant
;

# ~~~ [ Function Declaration ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#functions

# ref: ParseDeclare
#
#   ::= 'declare' FunctionHeader

FuncDecl -> FuncDecl
	: 'declare' Metadata=FuncMetadata Header=FuncHeader
;

# ~~~ [ Function Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#functions

# ref: ParseDefine
#
#   ::= 'define' FunctionHeader (!dbg !56)* '{' ...

FuncDef -> FuncDef
	: 'define' Header=FuncHeader Metadata=FuncMetadata Body=FuncBody
;

# ref: ParseFunctionHeader
#
#   ::= OptionalLinkage OptionalPreemptionSpecifier OptionalVisibility
#       OptionalCallingConv OptRetAttrs OptUnnamedAddr Type GlobalName
#       '(' ArgList ')' OptAddrSpace OptFuncAttrs OptSection OptionalAlign
#       OptGC OptionalPrefix OptionalPrologue OptPersonalityFn

# TODO: Add OptAlignment before OptGC once the LR-1 conflict has been resolved.
# The shift/reduce conflict is present since FuncAttr also contains 'align'.

FuncHeader -> FuncHeader
	: Linkageopt PreemptionSpecifieropt Visibilityopt DLLStorageClassopt CallingConvopt ReturnAttr* RetType=Type Name=GlobalIdent '(' Params ')' UnnamedAddropt AddrSpaceopt FuncAttrs=FuncAttr* Sectionopt Comdatopt GCopt Prefixopt Prologueopt Personalityopt
;

# TODO: Rename GCNode to GC when collision with token 'gc' has been resolved.
# Both define an identifier GC, the former in listener.go and the latter in token.go.

GC -> GCNode
	: 'gc' Name=StringLit
;

Prefix -> Prefix
	: 'prefix' Typ=Type Val=Constant
;

Prologue -> Prologue
	: 'prologue' Typ=Type Val=Constant
;

Personality -> Personality
	: 'personality' Typ=Type Val=Constant
;

# ref: ParseFunctionBody
#
#   ::= '{' BasicBlock+ UseListOrderDirective* '}'

FuncBody -> FuncBody
	: '{' Blocks=BasicBlock+ UseListOrders=UseListOrder* '}'
;

# ~~~ [ Attribute Group Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#attribute-groups

# ref: ParseUnnamedAttrGrp
#
#   ::= 'attributes' AttrGrpID '=' '{' AttrValPair+ '}'

AttrGroupDef -> AttrGroupDef
	: 'attributes' Name=AttrGroupID '=' '{' Attrs=FuncAttr* '}'
;

# ~~~ [ Named Metadata Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#named-metadata

# ref: ParseNamedMetadata
#
#   !foo = !{ !1, !2 }

NamedMetadataDef -> NamedMetadataDef
	: Name=MetadataName '=' '!' '{' MDNodes=(MetadataNode separator ',')* '}'
;

%interface MetadataNode;

MetadataNode -> MetadataNode
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

MetadataDef -> MetadataDef
	: Name=MetadataID '=' Distinctopt MDNode=MDTuple
	| Name=MetadataID '=' Distinctopt MDNode=SpecializedMDNode
;

Distinct -> Distinct
	: 'distinct'
;

# ~~~ [ Use-list Order Directives ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#use-list-order-directives

# ref: ParseUseListOrder
#
#   ::= 'uselistorder' Type Value ',' UseListOrderIndexes
#  UseListOrderIndexes
#   ::= '{' uint32 (',' uint32)+ '}'

UseListOrder -> UseListOrder
	: 'uselistorder' Typ=Type Val=Value ',' '{' Indicies=(UintLit separator ',')+ '}'
;

# ref: ParseUseListOrderBB
#
#   ::= 'uselistorder_bb' @foo ',' %bar ',' UseListOrderIndexes

UseListOrderBB -> UseListOrderBB
	: 'uselistorder_bb' Func=GlobalIdent ',' Block=LocalIdent ',' '{' Indicies=(UintLit separator ',')+ '}'
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

%interface Type;

Type -> Type
	: VoidType
	| FuncType
	| FirstClassType
;

%interface FirstClassType;

FirstClassType -> FirstClassType
	: ConcreteType
	| MetadataType
;

%interface ConcreteType;

ConcreteType -> ConcreteType
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

VoidType -> VoidType
	: 'void'
;

# --- [ Function Types ] -------------------------------------------------------

# ref: ParseFunctionType
#
#  ::= Type ArgumentList OptionalAttrs

FuncType -> FuncType
	: RetType=Type '(' Params ')'
;

# --- [ Integer Types ] --------------------------------------------------------

IntType -> IntType
	: int_type_tok
;

# --- [ Floating-point Types ] -------------------------------------------------

FloatType -> FloatType
	: FloatKind
;

FloatKind -> FloatKind
	: 'half'
	| 'float'
	| 'double'
	| 'x86_fp80'
	| 'fp128'
	| 'ppc_fp128'
;

# --- [ MMX Types ] ------------------------------------------------------------

MMXType -> MMXType
	: 'x86_mmx'
;

# --- [ Pointer Types ] --------------------------------------------------------

PointerType -> PointerType
	: Elem=Type AddrSpaceopt '*'
;

# --- [ Vector Types ] ---------------------------------------------------------

# ref: ParseArrayVectorType
#
#     ::= '<' APSINTVAL 'x' Types '>'

VectorType -> VectorType
	: '<' Len=UintLit 'x' Elem=Type '>'
;

# --- [ Label Types ] ----------------------------------------------------------

LabelType -> LabelType
	: 'label'
;

# --- [ Token Types ] ----------------------------------------------------------

TokenType -> TokenType
	: 'token'
;

# --- [ Metadata Types ] -------------------------------------------------------

MetadataType -> MetadataType
	: 'metadata'
;

# --- [ Array Types ] ----------------------------------------------------------

# ref: ParseArrayVectorType
#
#     ::= '[' APSINTVAL 'x' Types ']'

ArrayType -> ArrayType
	: '[' Len=UintLit 'x' Elem=Type ']'
;

# --- [ Structure Types ] ------------------------------------------------------

# ref: ParseStructBody
#
#   StructType
#     ::= '{' '}'
#     ::= '{' Type (',' Type)* '}'
#     ::= '<' '{' '}' '>'
#     ::= '<' '{' Type (',' Type)* '}' '>'

# TODO: Figure out how to represent packed; ref: https://github.com/inspirer/textmapper/issues/14

StructType -> StructType
	: '{' Fields=(Type separator ',')+? '}'
	| '<' '{' Fields=(Type separator ',')+? '}' '>'
;

OpaqueType -> OpaqueType
	: 'opaque'
;

# --- [ Named Types ] ----------------------------------------------------------

NamedType -> NamedType
	: Name=LocalIdent
;

# === [ Values ] ===============================================================

# ref: ParseValue

%interface Value;

Value -> Value
	: Constant
	# %42
	# %foo
	| LocalIdent
	# Inline assembler expressions may only be used as the callee operand of a
	# call or an invoke instruction.
	| InlineAsm
;

# --- [ Inline Assembler Expressions ] -----------------------------------------

# https://llvm.org/docs/LangRef.html#inline-assembler-expressions

# ref: ParseValID
#
#  ::= 'asm' SideEffect? AlignStack? IntelDialect? STRINGCONSTANT ','
#             STRINGCONSTANT

InlineAsm -> InlineAsm
	: 'asm' SideEffectopt AlignStackopt IntelDialectopt Asm=StringLit ',' Constraints=StringLit
;

SideEffect -> SideEffect
	: 'sideeffect'
;

AlignStack -> AlignStack
	: 'alignstack'
;

IntelDialect -> IntelDialect
	: 'inteldialect'
;

# //////////////////////////////////////////////////////////////////////////////

# ref: ParseOptionalAddrSpace
#
#   := empty
#   := 'addrspace' '(' uint32 ')'

AddrSpace -> AddrSpace
	: 'addrspace' '(' N=UintLit ')'
;

# ref: ParseOptionalAlignment
#
#   ::= empty
#   ::= 'align' 4

# TODO: Rename Alignment to Align.

Alignment -> Alignment
	: 'align' N=UintLit
;

AlignPair -> AlignPair
	: 'align' '=' N=UintLit
;

AlignStack -> AlignStack
	: 'alignstack' '=' N=UintLit
;

# ref: parseAllocSizeArguments

AllocSize -> AllocSize
	: 'allocsize' '(' ElemSize=UintLit ')'
	| 'allocsize' '(' ElemSize=UintLit ',' N=UintLit ')'
;

AttrPair -> AttrPair
	: Key=StringLit '=' Val=StringLit
;

AttrString -> AttrString
	: Val=StringLit
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
#   ::= 'aarch64_vector_pcs'
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

CallingConv -> CallingConv
	: 'aarch64_vector_pcs'
	| 'amdgpu_cs'
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

# ref: parseOptionalComdat

Comdat -> Comdat
	: 'comdat'
	| 'comdat' '(' Name=ComdatName ')'
;

# ref: ParseOptionalDerefAttrBytes
#
#   ::= empty
#   ::= AttrKind '(' 4 ')'

Dereferenceable -> Dereferenceable
	: 'dereferenceable' '(' N=UintLit ')'
	| 'dereferenceable_or_null' '(' N=UintLit ')'
;

# https://llvm.org/docs/LangRef.html#dll-storage-classes

# ref: ParseOptionalDLLStorageClass
#
#   ::= empty
#   ::= 'dllimport'
#   ::= 'dllexport'

DLLStorageClass -> DLLStorageClass
	: 'dllexport'
	| 'dllimport'
;

# ref: ParseFnAttributeValuePairs
#
#   ::= <attr> | <attr> '=' <value>

# NOTE: FuncAttr should contain Alignment. However, using LALR(1) this
# produces a reduce/reduce conflict as GlobalAttr also contains Alignment.
#
# To handle these ambiguities, (FuncAttr | Alignment) is used in those places
# where FuncAttr is used outside of GlobalDef and GlobalDecl (which also has
# GlobalAttr).

%interface FuncAttr;

FuncAttr -> FuncAttr
	: AttrString
	| AttrPair
	# not used in attribute groups.
	| AttrGroupID
	# used in attribute groups.
	| AlignPair
	| AlignStack
	# used in functions.
	#| Alignment # NOTE: removed to resolve reduce/reduce conflict, see above.
	| AllocSize
	| StackAlignment
	| FuncAttribute
;

FuncAttribute -> FuncAttribute
	: 'alwaysinline'
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
	| 'nocf_check'
	| 'noduplicate'
	| 'noimplicitfloat'
	| 'noinline'
	| 'nonlazybind'
	| 'norecurse'
	| 'noredzone'
	| 'noreturn'
	| 'nounwind'
	| 'optforfuzzing'
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
	| 'shadowcallstack'
	| 'speculatable'
	| 'speculative_load_hardening'
	| 'ssp'
	| 'sspreq'
	| 'sspstrong'
	| 'strictfp'
	| 'uwtable'
	| 'writeonly'
;

%interface GlobalAttr;

GlobalAttr -> GlobalAttr
	: Section
	| Comdat
	| Alignment
	#   ::= !dbg !57
	| MetadataAttachment
;

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

Linkage -> Linkage
	: 'appending'
	| 'available_externally'
	| 'common'
	| 'internal'
	| 'linkonce'
	| 'linkonce_odr'
	| 'private'
	| 'weak'
	| 'weak_odr'
	# External linkage.
	| 'extern_weak'
	| 'external'
;

# ref: ParseArgumentList
#
#   ::= '(' ArgTypeListI ')'
#  ArgTypeListI
#   ::= empty
#   ::= '...'
#   ::= ArgTypeList ',' '...'
#   ::= ArgType (',' ArgType)*

# TODO: Figure out how to handle variadic. ref: https://github.com/inspirer/textmapper/issues/14

Params -> Params
	: '...'?
	| Params=(Param separator ',')+ (',' '...')?
;

Param -> Param
	: Typ=Type Attrs=ParamAttr* Name=LocalIdent?
;

# ref: ParseOptionalParamAttrs

%interface ParamAttr;

ParamAttr -> ParamAttr
	: AttrString
	| AttrPair
	| Alignment
	| Dereferenceable
	| ParamAttribute
;

# TODO: Figure out a cleaner way of handling ParamAttribute.
# Written this way as a workaround for `'byval' cannot be used as an interface`
# which happens when the alternatives of ParamAttribute is inlined
# with ParamAttr in the grammar.

ParamAttribute -> ParamAttribute
	: 'byval'
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

# https://llvm.org/docs/LangRef.html#runtime-preemption-model

# ref: ParseOptionalDSOLocal

PreemptionSpecifier -> PreemptionSpecifier
	: 'dso_local'
	| 'dso_preemptable'
;

# ref: ParseOptionalReturnAttrs

%interface ReturnAttr;

ReturnAttr -> ReturnAttr
	# TODO: Figure out how to re-enable without getting these errors in FuncHeader:
	#    - two unnamed fields share the same type `AttrPair`: ReturnAttr -vs- FuncAttr
	#    - `AttrPair` occurs in both named and unnamed fields
	#    - `ReturnAttrs` cannot be nullable, since it precedes FuncAttrs
	#: AttrString
	#| AttrPair
	: Alignment
	| Dereferenceable
	| ReturnAttribute
;

ReturnAttribute -> ReturnAttribute
	: 'inreg'
	| 'noalias'
	| 'nonnull'
	| 'signext'
	| 'zeroext'
;

Section -> Section
	: 'section' Name=StringLit
;

# ref: ParseOptionalStackAlignment
#
#   ::= empty
#   ::= 'alignstack' '(' 4 ')'
StackAlignment -> StackAlignment
	: 'alignstack' '(' N=UintLit ')'
;

# ref: ParseOptionalThreadLocal
#
#   := empty
#   := 'thread_local'
#   := 'thread_local' '(' tlsmodel ')'

ThreadLocal -> ThreadLocal
	: 'thread_local'
	| 'thread_local' '(' Model=TLSModel ')'
;

# ref: ParseTLSModel
#
#   := 'localdynamic'
#   := 'initialexec'
#   := 'localexec'

TLSModel -> TLSModel
	: 'initialexec'
	| 'localdynamic'
	| 'localexec'
;

# ref: ParseOptionalUnnamedAddr

UnnamedAddr -> UnnamedAddr
	: 'local_unnamed_addr'
	| 'unnamed_addr'
;

# https://llvm.org/docs/LangRef.html#visibility-styles

# ref: ParseOptionalVisibility
#
#   ::= empty
#   ::= 'default'
#   ::= 'hidden'
#   ::= 'protected'

Visibility -> Visibility
	: 'default'
	| 'hidden'
	| 'protected'
;

# TODO: fix metadata.

MetadataAttachment -> MetadataAttachment
	: placeholder1
;

# TODO: fix Constant.

Constant -> Constant
	: placeholder2
;

UseListOrder -> UseListOrder
	: placeholder3
;

BasicBlock -> BasicBlock
	: placeholder2
;

FuncMetadata -> FuncMetadata
	: placeholder1
;

MDTuple -> MDTuple
	: placeholder2
;

SpecializedMDNode -> SpecializedMDNode
	: placeholder3
;

DIExpression -> DIExpression
	: placeholder1
;
