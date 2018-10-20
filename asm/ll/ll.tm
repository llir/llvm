language llvm(go);

lang = "llvm"
package = "github.com/mewmew/l-tm/asm/ll"
eventBased = true
eventFields = true

# TODO: check when to use Fooopt and when to use Foo? (as based on the AST
# they produce)

# ### [ Lexical part ] #########################################################

:: lexer

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

# FullDebug
emission_kind_tok : /(DebugDirectivesOnly)|(FullDebug)|(LineTablesOnly)|(NoDebug)/

# GNU
name_table_kind_tok : /(GNU)|(None)|(Default)/

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

# --- [ Null literals ] --------------------------------------------------------

NullLit -> NullLit
	: 'null'
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
	| GlobalDecl
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

# TODO: Rename `Typ=` to `Type=` once https://github.com/inspirer/textmapper/issues/13
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

GlobalDecl -> GlobalDecl
	: Name=GlobalIdent '=' ExternLinkage Preemptionopt Visibilityopt DLLStorageClassopt ThreadLocalopt UnnamedAddropt AddrSpaceopt ExternallyInitializedopt Immutable ContentType=Type GlobalAttrs=(',' GlobalAttr)+? FuncAttrs=(',' FuncAttr)+?
;

# ~~~ [ Global Variable Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

GlobalDef -> GlobalDef
	: Name=GlobalIdent '=' Linkageopt Preemptionopt Visibilityopt DLLStorageClassopt ThreadLocalopt UnnamedAddropt AddrSpaceopt ExternallyInitializedopt Immutable ContentType=Type Init=Constant GlobalAttrs=(',' GlobalAttr)+? FuncAttrs=(',' FuncAttr)+?
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
	: Name=GlobalIdent '=' (ExternLinkage | Linkageopt) Preemptionopt Visibilityopt DLLStorageClassopt ThreadLocalopt UnnamedAddropt 'alias' Typ=Type ',' AliaseeType=Type Aliasee=Constant
;

IFuncDef -> IFuncDef
	: Name=GlobalIdent '=' (ExternLinkage | Linkageopt) Preemptionopt Visibilityopt DLLStorageClassopt ThreadLocalopt UnnamedAddropt 'ifunc' Typ=Type ',' ResolverType=Type Resolver=Constant
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
	: (Linkage | ExternLinkage)? Preemptionopt Visibilityopt DLLStorageClassopt CallingConvopt ReturnAttrs=ReturnAttr* RetType=Type Name=GlobalIdent '(' Params ')' UnnamedAddropt AddrSpaceopt FuncAttrs=FuncAttr* Sectionopt Comdatopt GCopt Prefixopt Prologueopt Personalityopt
;

# TODO: Rename GCNode to GC when collision with token 'gc' has been resolved.
# Both define an identifier GC, the former in listener.go and the latter in token.go.

# TODO: Create issue in Textmapper to track this upstream.

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

StructType -> StructType
	: '{' Fields=(Type separator ',')+? '}'
	| '<' '{' Fields=(Type separator ',')+? '}' '>'   -> PackedStructType
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
	# TODO: Move InlineAsm from Value to Callee and Invokee?
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

# === [ Constants ] ============================================================

# https://llvm.org/docs/LangRef.html#constants

# ref: ParseValID

%interface Constant;

Constant -> Constant
	: BoolConst
	| IntConst
	| FloatConst
	| NullConst
	| NoneConst
	| StructConst
	| ArrayConst
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

BoolConst -> BoolConst
	: BoolLit
;

# --- [ Integer Constants ] ----------------------------------------------------

# https://llvm.org/docs/LangRef.html#simple-constants

# ref: ParseValID

IntConst -> IntConst
	: IntLit
;

# --- [ Floating-point Constants ] ---------------------------------------------

# https://llvm.org/docs/LangRef.html#simple-constants

# ref: ParseValID

FloatConst -> FloatConst
	: FloatLit
;

# --- [ Null Pointer Constants ] -----------------------------------------------

# https://llvm.org/docs/LangRef.html#simple-constants

# ref: ParseValID

NullConst -> NullConst
	: NullLit
;

# --- [ Token Constants ] ------------------------------------------------------

# https://llvm.org/docs/LangRef.html#simple-constants

# ref: ParseValID

NoneConst -> NoneConst
	: 'none'
;

# --- [ Structure Constants ] --------------------------------------------------

# https://llvm.org/docs/LangRef.html#complex-constants

# ref: ParseValID
#
#  ::= '{' ConstVector '}'
#  ::= '<' '{' ConstVector '}' '>' --> Packed Struct.

StructConst -> StructConst
	: '{' Fields=(TypeConst separator ',')+? '}'
	| '<' '{' Fields=(TypeConst separator ',')+? '}' '>'
;

# --- [ Array Constants ] ------------------------------------------------------

# https://llvm.org/docs/LangRef.html#complex-constants

# ref: ParseValID
#
#  c "foo"

ArrayConst -> ArrayConst
	: '[' Elems=(TypeConst separator ',')* ']'
	| 'c' Val=StringLit                          -> CharArrayConst
;


# --- [ Vector Constants ] -----------------------------------------------------

# https://llvm.org/docs/LangRef.html#complex-constants

# ref: ParseValID
#
#  ::= '<' ConstVector '>'         --> Vector.

VectorConst -> VectorConst
	: '<' Elems=(TypeConst separator ',')* '>'
;

# --- [ Zero Initialization Constants ] ----------------------------------------

# https://llvm.org/docs/LangRef.html#complex-constants

# ref: ParseValID

ZeroInitializerConst -> ZeroInitializerConst
	: 'zeroinitializer'
;

# --- [ Undefined Values ] -----------------------------------------------------

# https://llvm.org/docs/LangRef.html#undefined-values

# ref: ParseValID

UndefConst -> UndefConst
	: 'undef'
;

# --- [ Addresses of Basic Blocks ] --------------------------------------------

# https://llvm.org/docs/LangRef.html#addresses-of-basic-blocks

# ref: ParseValID
#
#  ::= 'blockaddress' '(' @foo ',' %bar ')'

BlockAddressConst -> BlockAddressConst
	: 'blockaddress' '(' Func=GlobalIdent ',' Block=LocalIdent ')'
;

# === [ Constant expressions ] =================================================

# https://llvm.org/docs/LangRef.html#constant-expressions

# ref: ParseValID

%interface ConstantExpr;

ConstantExpr -> ConstantExpr
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

AddExpr -> AddExpr
	: 'add' OverflowFlags '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FAddExpr -> FAddExpr
	: 'fadd' '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

SubExpr -> SubExpr
	: 'sub' OverflowFlags '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FSubExpr -> FSubExpr
	: 'fsub' '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

MulExpr -> MulExpr
	: 'mul' OverflowFlags '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FMulExpr -> FMulExpr
	: 'fmul' '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

UDivExpr -> UDivExpr
	: 'udiv' Exactopt '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

SDivExpr -> SDivExpr
	: 'sdiv' Exactopt '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FDivExpr -> FDivExpr
	: 'fdiv' '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

URemExpr -> URemExpr
	: 'urem' '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

SRemExpr -> SRemExpr
	: 'srem' '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FRemExpr -> FRemExpr
	: 'frem' '(' X=TypeConst ',' Y=TypeConst ')'
;

# --- [ Bitwise expressions ] --------------------------------------------------

# https://llvm.org/docs/LangRef.html#constant-expressions

# ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

ShlExpr -> ShlExpr
	: 'shl' OverflowFlags '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

LShrExpr -> LShrExpr
	: 'lshr' Exactopt '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

AShrExpr -> AShrExpr
	: 'ashr' Exactopt '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

AndExpr -> AndExpr
	: 'and' '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

OrExpr -> OrExpr
	: 'or' '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

XorExpr -> XorExpr
	: 'xor' '(' X=TypeConst ',' Y=TypeConst ')'
;

# --- [ Vector expressions ] ---------------------------------------------------

# https://llvm.org/docs/LangRef.html#constant-expressions

# ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

ExtractElementExpr -> ExtractElementExpr
	: 'extractelement' '(' X=TypeConst ',' Index=TypeConst ')'
;

# ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

InsertElementExpr -> InsertElementExpr
	: 'insertelement' '(' X=TypeConst ',' Elem=TypeConst ',' Index=TypeConst ')'
;

# ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

ShuffleVectorExpr -> ShuffleVectorExpr
	: 'shufflevector' '(' X=TypeConst ',' Y=TypeConst ',' Mask=TypeConst ')'
;

# --- [ Aggregate expressions ] ------------------------------------------------

# https://llvm.org/docs/LangRef.html#constant-expressions

# ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

ExtractValueExpr -> ExtractValueExpr
	: 'extractvalue' '(' X=TypeConst Indices=(',' UintLit)* ')'
;

# ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

InsertValueExpr -> InsertValueExpr
	: 'insertvalue' '(' X=TypeConst ',' Elem=TypeConst Indices=(',' UintLit)* ')'
;

# --- [ Memory expressions ] ---------------------------------------------------

# https://llvm.org/docs/LangRef.html#constant-expressions

# ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

GetElementPtrExpr -> GetElementPtrExpr
	: 'getelementptr' InBoundsopt '(' ElemType=Type ',' Src=TypeConst Indices=(',' GEPIndex)* ')'
;

# ref: ParseGlobalValueVector
#
#   ::= empty
#   ::= [inrange] TypeAndValue (',' [inrange] TypeAndValue)*

GEPIndex -> GEPIndex
	: InRangeopt Index=TypeConst
;

InRange -> InRange
	: 'inrange'
;

# --- [ Conversion expressions ] -----------------------------------------------

# https://llvm.org/docs/LangRef.html#constant-expressions

# ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

TruncExpr -> TruncExpr
	: 'trunc' '(' From=TypeConst 'to' To=Type ')'
;

# ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

ZExtExpr -> ZExtExpr
	: 'zext' '(' From=TypeConst 'to' To=Type ')'
;

# ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

SExtExpr -> SExtExpr
	: 'sext' '(' From=TypeConst 'to' To=Type ')'
;

# ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FPTruncExpr -> FPTruncExpr
	: 'fptrunc' '(' From=TypeConst 'to' To=Type ')'
;

# ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FPExtExpr -> FPExtExpr
	: 'fpext' '(' From=TypeConst 'to' To=Type ')'
;

# ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FPToUIExpr -> FPToUIExpr
	: 'fptoui' '(' From=TypeConst 'to' To=Type ')'
;

# ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FPToSIExpr -> FPToSIExpr
	: 'fptosi' '(' From=TypeConst 'to' To=Type ')'
;

# ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

UIToFPExpr -> UIToFPExpr
	: 'uitofp' '(' From=TypeConst 'to' To=Type ')'
;

# ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

SIToFPExpr -> SIToFPExpr
	: 'sitofp' '(' From=TypeConst 'to' To=Type ')'
;

# ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

PtrToIntExpr -> PtrToIntExpr
	: 'ptrtoint' '(' From=TypeConst 'to' To=Type ')'
;

# ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

IntToPtrExpr -> IntToPtrExpr
	: 'inttoptr' '(' From=TypeConst 'to' To=Type ')'
;

# ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

BitCastExpr -> BitCastExpr
	: 'bitcast' '(' From=TypeConst 'to' To=Type ')'
;

# ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

AddrSpaceCastExpr -> AddrSpaceCastExpr
	: 'addrspacecast' '(' From=TypeConst 'to' To=Type ')'
;

# --- [ Other expressions ] ----------------------------------------------------

# https://llvm.org/docs/LangRef.html#constant-expressions

# ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

ICmpExpr -> ICmpExpr
	: 'icmp' Pred=IPred '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

FCmpExpr -> FCmpExpr
	: 'fcmp' Pred=FPred '(' X=TypeConst ',' Y=TypeConst ')'
;

# ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseValID

SelectExpr -> SelectExpr
	: 'select' '(' Cond=TypeConst ',' X=TypeConst ',' Y=TypeConst ')'
;

# === [ Basic Blocks ] =========================================================

# ref: ParseBasicBlock
#
#   ::= LabelStr? Instruction*

BasicBlock -> BasicBlock
	: Name=LabelIdent? Insts=Instruction* Term=Terminator
;

# === [ Instructions ] =========================================================

# https://llvm.org/docs/LangRef.html#instruction-reference

# ref: ParseInstruction

%interface Instruction;

Instruction -> Instruction
	# Instructions not producing values.
	: StoreInst
	| FenceInst
	| CmpXchgInst
	| AtomicRMWInst
	# Instructions producing values.
	| LocalDef
	| ValueInstruction
;

LocalDef -> LocalDef
	: Name=LocalIdent '=' Inst=ValueInstruction
;

%interface ValueInstruction;

ValueInstruction -> ValueInstruction
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

AddInst -> AddInst
	: 'add' OverflowFlags X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fadd-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

FAddInst -> FAddInst
	: 'fadd' FastMathFlags=FastMathFlag* X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#sub-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

SubInst -> SubInst
	: 'sub' OverflowFlags X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fsub-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

FSubInst -> FSubInst
	: 'fsub' FastMathFlags=FastMathFlag* X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#mul-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

MulInst -> MulInst
	: 'mul' OverflowFlags X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fmul-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

FMulInst -> FMulInst
	: 'fmul' FastMathFlags=FastMathFlag* X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#udiv-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

UDivInst -> UDivInst
	: 'udiv' Exactopt X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#sdiv-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

SDivInst -> SDivInst
	: 'sdiv' Exactopt X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fdiv-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

FDivInst -> FDivInst
	: 'fdiv' FastMathFlags=FastMathFlag* X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#urem-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

URemInst -> URemInst
	: 'urem' X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#srem-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

SRemInst -> SRemInst
	: 'srem' X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#frem-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

FRemInst -> FRemInst
	: 'frem' FastMathFlags=FastMathFlag* X=TypeValue ',' Y=Value InstMetadata
;

# --- [ Bitwise instructions ] -------------------------------------------------

# ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#shl-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

ShlInst -> ShlInst
	: 'shl' OverflowFlags X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#lshr-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

LShrInst -> LShrInst
	: 'lshr' Exactopt X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#ashr-instruction

# ref: ParseArithmetic
#
#  ::= ArithmeticOps TypeAndValue ',' Value

AShrInst -> AShrInst
	: 'ashr' Exactopt X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#and-instruction

# ref: ParseLogical
#
#  ::= ArithmeticOps TypeAndValue ',' Value {

AndInst -> AndInst
	: 'and' X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#or-instruction

# ref: ParseLogical
#
#  ::= ArithmeticOps TypeAndValue ',' Value {

OrInst -> OrInst
	: 'or' X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#xor-instruction

# ref: ParseLogical
#
#  ::= ArithmeticOps TypeAndValue ',' Value {

XorInst -> XorInst
	: 'xor' X=TypeValue ',' Y=Value InstMetadata
;

# --- [ Vector instructions ] --------------------------------------------------

# ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#extractelement-instruction

# ref: ParseExtractElement
#
#   ::= 'extractelement' TypeAndValue ',' TypeAndValue

ExtractElementInst -> ExtractElementInst
	: 'extractelement' X=TypeValue ',' Index=TypeValue InstMetadata
;

# ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#insertelement-instruction

# ref: ParseInsertElement
#
#   ::= 'insertelement' TypeAndValue ',' TypeAndValue ',' TypeAndValue

InsertElementInst -> InsertElementInst
	: 'insertelement' X=TypeValue ',' Elem=TypeValue ',' Index=TypeValue InstMetadata
;

# ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#shufflevector-instruction

# ref: ParseShuffleVector
#
#   ::= 'shufflevector' TypeAndValue ',' TypeAndValue ',' TypeAndValue

ShuffleVectorInst -> ShuffleVectorInst
	: 'shufflevector' X=TypeValue ',' Y=TypeValue ',' Mask=TypeValue InstMetadata
;

# --- [ Aggregate instructions ] -----------------------------------------------

# ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#extractvalue-instruction

# ref: ParseExtractValue
#
#   ::= 'extractvalue' TypeAndValue (',' uint32)+

ExtractValueInst -> ExtractValueInst
   : 'extractvalue' X=TypeValue Indices=(',' UintLit)+ InstMetadata
;

# ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#insertvalue-instruction

# ref: ParseInsertValue
#
#   ::= 'insertvalue' TypeAndValue ',' TypeAndValue (',' uint32)+

InsertValueInst -> InsertValueInst
   : 'insertvalue' X=TypeValue ',' Elem=TypeValue Indices=(',' UintLit)+ InstMetadata
;

# --- [ Memory instructions ] --------------------------------------------------

# ~~~ [ alloca ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#alloca-instruction

# ref: ParseAlloc
#
#   ::= 'alloca' 'inalloca'? 'swifterror'? Type (',' TypeAndValue)?
#       (',' 'align' i32)? (',', 'addrspace(n))?

AllocaInst -> AllocaInst
	: 'alloca' InAllocaopt SwiftErroropt ElemType=Type NElems=(',' TypeValue)? (',' Alignment)? (',' AddrSpace)? InstMetadata
;

InAlloca -> InAlloca
	: 'inalloca'
;

SwiftError -> SwiftError
	: 'swifterror'
;

# ~~~ [ load ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#load-instruction

# ref: ParseLoad
#
#   ::= 'load' 'volatile'? TypeAndValue (',' 'align' i32)?
#   ::= 'load' 'atomic' 'volatile'? TypeAndValue
#       'singlethread'? AtomicOrdering (',' 'align' i32)?

LoadInst -> LoadInst
	# Load.
	: 'load' Volatileopt ElemType=Type ',' Src=TypeValue (',' Alignment)? InstMetadata
	# Atomic load.
	| 'load' 'atomic' Volatileopt ElemType=Type ',' Src=TypeValue SyncScopeopt AtomicOrdering (',' Alignment)? InstMetadata
;

# ~~~ [ store ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#store-instruction

# ref: ParseStore
#
#   ::= 'store' 'volatile'? TypeAndValue ',' TypeAndValue (',' 'align' i32)?
#   ::= 'store' 'atomic' 'volatile'? TypeAndValue ',' TypeAndValue
#       'singlethread'? AtomicOrdering (',' 'align' i32)?

StoreInst -> StoreInst
	: 'store' Volatileopt Src=TypeValue ',' Dst=TypeValue (',' Alignment)? InstMetadata
	| 'store' 'atomic' Volatileopt Src=TypeValue ',' Dst=TypeValue SyncScopeopt AtomicOrdering (',' Alignment)? InstMetadata
;

# ~~~ [ fence ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fence-instruction

# ref: ParseFence
#
#   ::= 'fence' 'singlethread'? AtomicOrdering

FenceInst -> FenceInst
	: 'fence' SyncScopeopt AtomicOrdering InstMetadata
;

# ~~~ [ cmpxchg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#cmpxchg-instruction

# ref: ParseCmpXchg
#
#   ::= 'cmpxchg' 'weak'? 'volatile'? TypeAndValue ',' TypeAndValue ','
#       TypeAndValue 'singlethread'? AtomicOrdering AtomicOrdering

CmpXchgInst -> CmpXchgInst
	: 'cmpxchg' Weakopt Volatileopt Ptr=TypeValue ',' Cmp=TypeValue ',' New=TypeValue SyncScopeopt Success=AtomicOrdering Failure=AtomicOrdering InstMetadata
;

Weak -> Weak
	: 'weak'
;

# ~~~ [ atomicrmw ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#atomicrmw-instruction

# ref: ParseAtomicRMW
#
#   ::= 'atomicrmw' 'volatile'? BinOp TypeAndValue ',' TypeAndValue
#       'singlethread'? AtomicOrdering

AtomicRMWInst -> AtomicRMWInst
	: 'atomicrmw' Volatileopt Op=BinOp Ptr=TypeValue ',' X=TypeValue SyncScopeopt AtomicOrdering InstMetadata
;

BinOp -> BinOp
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

GetElementPtrInst -> GetElementPtrInst
	: 'getelementptr' InBoundsopt ElemType=Type ',' Src=TypeValue Indices=(',' TypeValue)* InstMetadata
;

# --- [ Conversion instructions ] ----------------------------------------------

# ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#trunc-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

TruncInst -> TruncInst
	: 'trunc' From=TypeValue 'to' To=Type InstMetadata
;

# ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#zext-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

ZExtInst -> ZExtInst
	: 'zext' From=TypeValue 'to' To=Type InstMetadata
;

# ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#sext-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

SExtInst -> SExtInst
	: 'sext' From=TypeValue 'to' To=Type InstMetadata
;

# ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fptrunc-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

FPTruncInst -> FPTruncInst
	: 'fptrunc' From=TypeValue 'to' To=Type InstMetadata
;

# ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fpext-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

FPExtInst -> FPExtInst
	: 'fpext' From=TypeValue 'to' To=Type InstMetadata
;

# ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fptoui-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

FPToUIInst -> FPToUIInst
	: 'fptoui' From=TypeValue 'to' To=Type InstMetadata
;

# ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fptosi-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

FPToSIInst -> FPToSIInst
	: 'fptosi' From=TypeValue 'to' To=Type InstMetadata
;

# ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#uitofp-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

UIToFPInst -> UIToFPInst
	: 'uitofp' From=TypeValue 'to' To=Type InstMetadata
;

# ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#sitofp-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

SIToFPInst -> SIToFPInst
	: 'sitofp' From=TypeValue 'to' To=Type InstMetadata
;

# ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#ptrtoint-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

PtrToIntInst -> PtrToIntInst
	: 'ptrtoint' From=TypeValue 'to' To=Type InstMetadata
;

# ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#inttoptr-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

IntToPtrInst -> IntToPtrInst
	: 'inttoptr' From=TypeValue 'to' To=Type InstMetadata
;

# ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#bitcast-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

BitCastInst -> BitCastInst
	: 'bitcast' From=TypeValue 'to' To=Type InstMetadata
;

# ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#addrspacecast-instruction

# ref: ParseCast
#
#   ::= CastOpc TypeAndValue 'to' Type

AddrSpaceCastInst -> AddrSpaceCastInst
	: 'addrspacecast' From=TypeValue 'to' To=Type InstMetadata
;

# --- [ Other instructions ] ---------------------------------------------------

# ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#icmp-instruction

# ref: ParseCompare
#
#  ::= 'icmp' IPredicates TypeAndValue ',' Value

ICmpInst -> ICmpInst
	: 'icmp' Pred=IPred X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#fcmp-instruction

# ref: ParseCompare
#
#  ::= 'fcmp' FPredicates TypeAndValue ',' Value

FCmpInst -> FCmpInst
	: 'fcmp' FastMathFlags=FastMathFlag* Pred=FPred X=TypeValue ',' Y=Value InstMetadata
;

# ~~~ [ phi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#phi-instruction

# ref: ParsePHI
#
#   ::= 'phi' Type '[' Value ',' Value ']' (',' '[' Value ',' Value ']')*

PhiInst -> PhiInst
	: 'phi' Typ=Type Incs=(Inc separator ',')+ InstMetadata
;

Inc -> Inc
	: '[' X=Value ',' Pred=LocalIdent ']'
;

# ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#select-instruction

# ref: ParseSelect
#
#   ::= 'select' TypeAndValue ',' TypeAndValue ',' TypeAndValue

SelectInst -> SelectInst
	: 'select' Cond=TypeValue ',' X=TypeValue ',' Y=TypeValue InstMetadata
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

CallInst -> CallInst
	: Tailopt 'call' FastMathFlags=FastMathFlag* CallingConvopt ReturnAttrs=ReturnAttr* AddrSpaceopt Typ=Type Callee=Value '(' Args ')' FuncAttrs=FuncAttr* OperandBundles InstMetadata
;

Tail -> Tail
	: 'musttail'
	| 'notail'
	| 'tail'
;

# ~~~ [ va_arg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#va_arg-instruction

# ref: ParseVA_Arg
#
#   ::= 'va_arg' TypeAndValue ',' Type

VAArgInst -> VAArgInst
	: 'va_arg' ArgList=TypeValue ',' ArgType=Type InstMetadata
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

LandingPadInst -> LandingPadInst
	: 'landingpad' Typ=Type Cleanupopt Clauses=Clause* InstMetadata
;

Cleanup -> Cleanup
	: 'cleanup'
;

%interface Clause;

Clause -> Clause
	: CatchClause
	| FilterClause
;

CatchClause -> CatchClause
	: 'catch' X=TypeValue
;

FilterClause -> FilterClause
	: 'filter' Typ=Type Val=ArrayConst
;

# ~~~ [ catchpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseCatchPad
#
#   ::= 'catchpad' ParamList 'to' TypeAndValue 'unwind' TypeAndValue

CatchPadInst -> CatchPadInst
	: 'catchpad' 'within' Scope=LocalIdent '[' Args=(ExceptionArg separator ',')* ']' InstMetadata
;

# ~~~ [ cleanuppad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ref: ParseCleanupPad
#
#   ::= 'cleanuppad' within Parent ParamList

CleanupPadInst -> CleanupPadInst
	: 'cleanuppad' 'within' Scope=ExceptionScope '[' Args=(ExceptionArg separator ',')* ']' InstMetadata
;

# === [ Terminators ] ==========================================================

# https://llvm.org/docs/LangRef.html#terminator-instructions

# ref: ParseInstruction

%interface Terminator;

Terminator -> Terminator
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

RetTerm -> RetTerm
	# Void return.
	: 'ret' XTyp=VoidType InstMetadata
	# Value return.
	| 'ret' XTyp=ConcreteType X=Value InstMetadata
;

# --- [ br ] -------------------------------------------------------------------

# https://llvm.org/docs/LangRef.html#br-instruction

# ref: ParseBr
#
#   ::= 'br' TypeAndValue
#   ::= 'br' TypeAndValue ',' TypeAndValue ',' TypeAndValue

# Unconditional branch.
BrTerm -> BrTerm
	: 'br' Target=Label InstMetadata
;

# TODO: replace `IntType Value` with TypeValue when the parser generator
# is capable of handling the shift/reduce conflict.

# Conditional branch.
CondBrTerm -> CondBrTerm
	: 'br' CondTyp=IntType Cond=Value ',' TargetTrue=Label ',' TargetFalse=Label InstMetadata
;

# --- [ switch ] ---------------------------------------------------------------

# https://llvm.org/docs/LangRef.html#switch-instruction

# ref: ParseSwitch
#
#    ::= 'switch' TypeAndValue ',' TypeAndValue '[' JumpTable ']'
#  JumpTable
#    ::= (TypeAndValue ',' TypeAndValue)*

SwitchTerm -> SwitchTerm
	: 'switch' X=TypeValue ',' Default=Label '[' Cases=Case* ']' InstMetadata
;

Case -> Case
	: X=TypeConst ',' Target=Label
;

# --- [ indirectbr ] -----------------------------------------------------------

# https://llvm.org/docs/LangRef.html#indirectbr-instruction

# ref: ParseIndirectBr
#
#    ::= 'indirectbr' TypeAndValue ',' '[' LabelList ']'

IndirectBrTerm -> IndirectBrTerm
	: 'indirectbr' Addr=TypeValue ',' '[' Targets=(Label separator ',')+ ']' InstMetadata
;

# --- [ invoke ] ---------------------------------------------------------------

# https://llvm.org/docs/LangRef.html#invoke-instruction

# ref: ParseInvoke
#
#   ::= 'invoke' OptionalCallingConv OptionalAttrs Type Value ParamList
#       OptionalAttrs 'to' TypeAndValue 'unwind' TypeAndValue

InvokeTerm -> InvokeTerm
	: 'invoke' CallingConvopt ReturnAttrs=ReturnAttr* AddrSpaceopt Invokee=TypeValue '(' Args ')' FuncAttrs=FuncAttr* OperandBundles 'to' Normal=Label 'unwind' Exception=Label InstMetadata
;

# --- [ resume ] ---------------------------------------------------------------

# https://llvm.org/docs/LangRef.html#resume-instruction

# ref: ParseResume
#
#   ::= 'resume' TypeAndValue

ResumeTerm -> ResumeTerm
	: 'resume' X=TypeValue InstMetadata
;

# --- [ catchswitch ] ----------------------------------------------------------

# https://llvm.org/docs/LangRef.html#catchswitch-instruction

# ref: ParseCatchSwitch
#
#   ::= 'catchswitch' within Parent

CatchSwitchTerm -> CatchSwitchTerm
	: 'catchswitch' 'within' Scope=ExceptionScope '[' Handlers=(Label separator ',')+ ']' 'unwind' UnwindTarget InstMetadata
;

# --- [ catchret ] -------------------------------------------------------------

# https://llvm.org/docs/LangRef.html#catchret-instruction

# ref: ParseCatchRet
#
#   ::= 'catchret' from Parent Value 'to' TypeAndValue

CatchRetTerm -> CatchRetTerm
	: 'catchret' 'from' From=Value 'to' To=Label InstMetadata
;

# --- [ cleanupret ] -----------------------------------------------------------

# https://llvm.org/docs/LangRef.html#cleanupret-instruction

# ref: ParseCleanupRet
#
#   ::= 'cleanupret' from Value unwind ('to' 'caller' | TypeAndValue)

CleanupRetTerm -> CleanupRetTerm
	: 'cleanupret' 'from' From=Value 'unwind' UnwindTarget InstMetadata
;

# --- [ unreachable ] ----------------------------------------------------------

# https://llvm.org/docs/LangRef.html#unreachable-instruction

# ref: ParseInstruction

UnreachableTerm -> UnreachableTerm
	: 'unreachable' InstMetadata
;

# === [ Metadata Nodes and Metadata Strings ] ==================================

# https://llvm.org/docs/LangRef.html#metadata-nodes-and-metadata-strings

# --- [ Metadata Tuple ] -------------------------------------------------------

# ref: ParseMDTuple

MDTuple -> MDTuple
	: '!' MDFields
;

# ref: ParseMDNodeVector
#
#   ::= { Element (',' Element)* }
#  Element
#   ::= 'null' | TypeAndValue

# ref: ParseMDField(MDFieldList &)

MDFields -> MDFields
	: '{' MDFields=(MDField separator',')* '}'
;

# ref: ParseMDField(MDField &)

%interface MDField;

MDField -> MDField
	# Null is a special case since it is typeless.
	: NullLit
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

%interface Metadata;

Metadata -> Metadata
	: TypeValue
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

MDString -> MDString
	: '!' Val=StringLit
;

# --- [ Metadata Attachment ] --------------------------------------------------

# ref: ParseMetadataAttachment
#
#   ::= !dbg !42

MetadataAttachment -> MetadataAttachment
	: Name=MetadataName MDNode
;

# --- [ Metadata Node ] --------------------------------------------------------

# ref: ParseMDNode
#
#  ::= !{ ... }
#  ::= !7
#  ::= !DILocation(...)

%interface MDNode;

MDNode -> MDNode
	# !{ ... }
	: MDTuple
	# !42
	| MetadataID
	| SpecializedMDNode
;

# --- [ Specialized Metadata Nodes ] -------------------------------------------

# https://llvm.org/docs/LangRef.html#specialized-metadata-nodes

# ref: ParseSpecializedMDNode

%interface SpecializedMDNode;

SpecializedMDNode -> SpecializedMDNode
	: DIBasicType
	| DICompileUnit
	| DICompositeType
	| DIDerivedType
	| DIEnumerator
	| DIExpression
	| DIFile
	| DIGlobalVariable
	| DIGlobalVariableExpression # not in spec as of 2018-02-21
	| DIImportedEntity
	| DILabel # not in spec as of 2018-10-14
	| DILexicalBlock
	| DILexicalBlockFile
	| DILocalVariable
	| DILocation
	| DIMacro
	| DIMacroFile
	| DIModule # not in spec as of 2018-02-21
	| DINamespace
	| DIObjCProperty
	| DISubprogram
	| DISubrange
	| DISubroutineType
	| DITemplateTypeParameter
	| DITemplateValueParameter
	| GenericDINode # not in spec as of 2018-02-21
;

# ~~~ [ DIBasicType ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dibasictype

# ref: ParseDIBasicType
#
#   ::= !DIBasicType(tag: DW_TAG_base_type, name: "int", size: 32, align: 32,
#                    encoding: DW_ATE_encoding, flags: 0)
#
#  OPTIONAL(tag, DwarfTagField, (dwarf::DW_TAG_base_type));
#  OPTIONAL(name, MDStringField, );
#  OPTIONAL(size, MDUnsignedField, (0, UINT64_MAX));
#  OPTIONAL(align, MDUnsignedField, (0, UINT32_MAX));
#  OPTIONAL(encoding, DwarfAttEncodingField, );
#  OPTIONAL(flags, DIFlagField, );

DIBasicType -> DIBasicType
	: '!DIBasicType' '(' Fields=(DIBasicTypeField separator ',')* ')'
;

%interface DIBasicTypeField;

DIBasicTypeField -> DIBasicTypeField
	: TagField
	| NameField
	| SizeField
	| AlignField
	| EncodingField
	| FlagsField
;

# ~~~ [ DICompileUnit ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dicompileunit

# ref: ParseDICompileUnit
#
#   ::= !DICompileUnit(language: DW_LANG_C99, file: !0, producer: "clang",
#                      isOptimized: true, flags: "-O2", runtimeVersion: 1,
#                      splitDebugFilename: "abc.debug",
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
#  OPTIONAL(nameTableKind, NameTableKindField, );


DICompileUnit -> DICompileUnit
	: '!DICompileUnit' '(' Fields=(DICompileUnitField separator ',')* ')'
;

%interface DICompileUnitField;

DICompileUnitField -> DICompileUnitField
	: LanguageField
	| FileField
	| ProducerField
	| IsOptimizedField
	| FlagsStringField
	| RuntimeVersionField
	| SplitDebugFilenameField
	| EmissionKindField
	| EnumsField
	| RetainedTypesField
	| GlobalsField
	| ImportsField
	| MacrosField
	| DwoIdField
	| SplitDebugInliningField
	| DebugInfoForProfilingField
	| NameTableKindField
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

DICompositeType -> DICompositeType
	: '!DICompositeType' '(' Fields=(DICompositeTypeField separator ',')* ')'
;

%interface DICompositeTypeField;

DICompositeTypeField -> DICompositeTypeField
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
	| ElementsField
	| RuntimeLangField
	| VtableHolderField
	| TemplateParamsField
	| IdentifierField
	| DiscriminatorField
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

DIDerivedType -> DIDerivedType
	: '!DIDerivedType' '(' Fields=(DIDerivedTypeField separator ',')* ')'
;

%interface DIDerivedTypeField;

DIDerivedTypeField -> DIDerivedTypeField
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
	| ExtraDataField
	| DwarfAddressSpaceField
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

DIEnumerator -> DIEnumerator
	: '!DIEnumerator' '(' Fields=(DIEnumeratorField separator ',')* ')'
;

%interface DIEnumeratorField;

DIEnumeratorField -> DIEnumeratorField
	: NameField
	| ValueIntField
	| IsUnsignedField
;

# ~~~ [ DIExpression ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#diexpression

# ref: ParseDIExpression
#
#   ::= !DIExpression(0, 7, -1)

DIExpression -> DIExpression
	: '!DIExpression' '(' Fields=(DIExpressionField separator ',')* ')'
;

%interface DIExpressionField;

DIExpressionField -> DIExpressionField
	: IntLit
	| DwarfOp
;

# ~~~ [ DIFile ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#difile

# ref: ParseDIFileType
#
#   ::= !DIFileType(filename: "path/to/file", directory: "/path/to/dir",
#                   checksumkind: CSK_MD5,
#                   checksum: "000102030405060708090a0b0c0d0e0f",
#                   source: "source file contents")
#
#  REQUIRED(filename, MDStringField, );
#  REQUIRED(directory, MDStringField, );
#  OPTIONAL(checksumkind, ChecksumKindField, (DIFile::CSK_MD5));
#  OPTIONAL(checksum, MDStringField, );
#  OPTIONAL(source, MDStringField, );

DIFile -> DIFile
	: '!DIFile' '(' Fields=(DIFileField separator ',')* ')'
;

%interface DIFileField;

DIFileField -> DIFileField
	: FilenameField
	| DirectoryField
	| ChecksumkindField
	| ChecksumField
	| SourceField
;

# ~~~ [ DIGlobalVariable ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#diglobalvariable

# ref: ParseDIGlobalVariable
#
#   ::= !DIGlobalVariable(scope: !0, name: "foo", linkageName: "foo",
#                         file: !1, line: 7, type: !2, isLocal: false,
#                         isDefinition: true, templateParams: !3,
#                         declaration: !4, align: 8)
#
#  REQUIRED(name, MDStringField, (AllowEmpty false));
#  OPTIONAL(scope, MDField, );
#  OPTIONAL(linkageName, MDStringField, );
#  OPTIONAL(file, MDField, );
#  OPTIONAL(line, LineField, );
#  OPTIONAL(type, MDField, );
#  OPTIONAL(isLocal, MDBoolField, );
#  OPTIONAL(isDefinition, MDBoolField, (true));
#  OPTIONAL(templateParams, MDField, );                                         \
#  OPTIONAL(declaration, MDField, );
#  OPTIONAL(align, MDUnsignedField, (0, UINT32_MAX));

DIGlobalVariable -> DIGlobalVariable
	: '!DIGlobalVariable' '(' Fields=(DIGlobalVariableField separator ',')* ')'
;

%interface DIGlobalVariableField;

DIGlobalVariableField -> DIGlobalVariableField
	: NameField
	| ScopeField
	| LinkageNameField
	| FileField
	| LineField
	| TypeField
	| IsLocalField
	| IsDefinitionField
	| TemplateParamsField
	| DeclarationField
	| AlignField
;

# ~~~ [ DIGlobalVariableExpression ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# TODO: add link to LangRef.html.

# ref: ParseDIGlobalVariableExpression
#
#   ::= !DIGlobalVariableExpression(var: !0, expr: !1)
#
#  REQUIRED(var, MDField, );
#  REQUIRED(expr, MDField, );

DIGlobalVariableExpression -> DIGlobalVariableExpression
	: '!DIGlobalVariableExpression' '(' Fields=(DIGlobalVariableExpressionField separator ',')* ')'
;

%interface DIGlobalVariableExpressionField;

DIGlobalVariableExpressionField -> DIGlobalVariableExpressionField
	: VarField
	| ExprField
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

DIImportedEntity -> DIImportedEntity
	: '!DIImportedEntity' '(' Fields=(DIImportedEntityField separator ',')* ')'
;

%interface DIImportedEntityField;

DIImportedEntityField -> DIImportedEntityField
	: TagField
	| ScopeField
	| EntityField
	| FileField
	| LineField
	| NameField
;

# ~~~ [ DILabel ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# TODO: add link to LangRef.html.

# ref: ParseDILabel:
#
#   ::= !DILabel(scope: !0, name: "foo", file: !1, line: 7)
#
#  REQUIRED(scope, MDField, (/* AllowNull */ false));                           \
#  REQUIRED(name, MDStringField, );                                             \
#  REQUIRED(file, MDField, );                                                   \
#  REQUIRED(line, LineField, );

DILabel -> DILabel
	: '!DILabel' '(' Fields=(DILabelField separator ',')* ')'
;

%interface DILabelField;

DILabelField -> DILabelField
	: ScopeField
	| NameField
	| FileField
	| LineField
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

DILexicalBlock -> DILexicalBlock
	: '!DILexicalBlock' '(' Fields=(DILexicalBlockField separator ',')* ')'
;

%interface DILexicalBlockField;

DILexicalBlockField -> DILexicalBlockField
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

DILexicalBlockFile -> DILexicalBlockFile
	: '!DILexicalBlockFile' '(' Fields=(DILexicalBlockFileField separator ',')* ')'
;

%interface DILexicalBlockFileField;

DILexicalBlockFileField -> DILexicalBlockFileField
	: ScopeField
	| FileField
	| DiscriminatorIntField
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

DILocalVariable -> DILocalVariable
	: '!DILocalVariable' '(' Fields=(DILocalVariableField separator ',')* ')'
;

%interface DILocalVariableField;

DILocalVariableField -> DILocalVariableField
	: NameField
	| ArgField
	| ScopeField
	| FileField
	| LineField
	| TypeField
	| FlagsField
	| AlignField
;

# ~~~ [ DILocation ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#dilocation

# ref: ParseDILocation
#
#   ::= !DILocation(line: 43, column: 8, scope: !5, inlinedAt: !6,
#   isImplicitCode: true)
#
#  OPTIONAL(line, LineField, );
#  OPTIONAL(column, ColumnField, );
#  REQUIRED(scope, MDField, (AllowNull false));
#  OPTIONAL(inlinedAt, MDField, );
#  OPTIONAL(isImplicitCode, MDBoolField, (false));

DILocation -> DILocation
	: '!DILocation' '(' Fields=(DILocationField separator ',')* ')'
;

%interface DILocationField;

DILocationField -> DILocationField
	: LineField
	| ColumnField
	| ScopeField
	| InlinedAtField
	| IsImplicitCodeField
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

DIMacro -> DIMacro
	: '!DIMacro' '(' Fields=(DIMacroField separator ',')* ')'
;

%interface DIMacroField;

DIMacroField -> DIMacroField
	: TypeMacinfoField
	| LineField
	| NameField
	| ValueStringField
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

DIMacroFile -> DIMacroFile
	: '!DIMacroFile' '(' Fields=(DIMacroFileField separator ',')* ')'
;

%interface DIMacroFileField;

DIMacroFileField -> DIMacroFileField
	: TypeMacinfoField
	| LineField
	| FileField
	| NodesField
;

# ~~~ [ DIModule ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# TODO: add link to LangRef.html.

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

DIModule -> DIModule
	: '!DIModule' '(' Fields=(DIModuleField separator ',')* ')'
;

%interface DIModuleField;

DIModuleField -> DIModuleField
	: ScopeField
	| NameField
	| ConfigMacrosField
	| IncludePathField
	| IsysrootField
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

DINamespace -> DINamespace
	: '!DINamespace' '(' Fields=(DINamespaceField separator ',')* ')'
;

%interface DINamespaceField;

DINamespaceField -> DINamespaceField
	: ScopeField
	| NameField
	| ExportSymbolsField
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

DIObjCProperty -> DIObjCProperty
	: '!DIObjCProperty' '(' Fields=(DIObjCPropertyField separator ',')* ')'
;

%interface DIObjCPropertyField;

DIObjCPropertyField -> DIObjCPropertyField
	: NameField
	| FileField
	| LineField
	| SetterField
	| GetterField
	| AttributesField
	| TypeField
;

# ~~~ [ DISubprogram ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#disubprogram

# ref: ParseDISubprogram
#
#   ::= !DISubprogram(scope: !0, name: "foo", linkageName: "_Zfoo",
#                     file: !1, line: 7, type: !2, isLocal: false,
#                     isDefinition: true, scopeLine: 8, containingType: !3,
#                     virtuality: DW_VIRTUALTIY_pure_virtual,
#                     virtualIndex: 10, thisAdjustment: 4, flags: 11,
#                     isOptimized: false, templateParams: !4, declaration: !5,
#                     retainedNodes: !6, thrownTypes: !7)
#
#  OPTIONAL(scope, MDField, );                                                  \
#  OPTIONAL(name, MDStringField, );                                             \
#  OPTIONAL(linkageName, MDStringField, );                                      \
#  OPTIONAL(file, MDField, );                                                   \
#  OPTIONAL(line, LineField, );                                                 \
#  OPTIONAL(type, MDField, );                                                   \
#  OPTIONAL(isLocal, MDBoolField, );                                            \
#  OPTIONAL(isDefinition, MDBoolField, (true));                                 \
#  OPTIONAL(scopeLine, LineField, );                                            \
#  OPTIONAL(containingType, MDField, );                                         \
#  OPTIONAL(virtuality, DwarfVirtualityField, );                                \
#  OPTIONAL(virtualIndex, MDUnsignedField, (0, UINT32_MAX));                    \
#  OPTIONAL(thisAdjustment, MDSignedField, (0, INT32_MIN, INT32_MAX));          \
#  OPTIONAL(flags, DIFlagField, );                                              \
#  OPTIONAL(isOptimized, MDBoolField, );                                        \
#  OPTIONAL(unit, MDField, );                                                   \
#  OPTIONAL(templateParams, MDField, );                                         \
#  OPTIONAL(declaration, MDField, );                                            \
#  OPTIONAL(retainedNodes, MDField, );                                              \
#  OPTIONAL(thrownTypes, MDField, );

DISubprogram -> DISubprogram
	: '!DISubprogram' '(' Fields=(DISubprogramField separator ',')* ')'
;

%interface DISubprogramField;

DISubprogramField -> DISubprogramField
	: ScopeField
	| NameField
	| LinkageNameField
	| FileField
	| LineField
	| TypeField
	| IsLocalField
	| IsDefinitionField
	| ScopeLineField
	| ContainingTypeField
	| VirtualityField
	| VirtualIndexField
	| ThisAdjustmentField
	| FlagsField
	| IsOptimizedField
	| UnitField
	| TemplateParamsField
	| DeclarationField
	| RetainedNodesField
	| ThrownTypesField
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

DISubrange -> DISubrange
	: '!DISubrange' '(' Fields=(DISubrangeField separator ',')* ')'
;

%interface DISubrangeField;

DISubrangeField -> DISubrangeField
	: CountField
	| LowerBoundField
;

# ~~~ [ DISubroutineType ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#disubroutinetype

# ref: ParseDISubroutineType
#
#  OPTIONAL(flags, DIFlagField, );
#  OPTIONAL(cc, DwarfCCField, );
#  REQUIRED(types, MDField, );

DISubroutineType -> DISubroutineType
	: '!DISubroutineType' '(' Fields=(DISubroutineTypeField separator ',')* ')'
;

%interface DISubroutineTypeField;

DISubroutineTypeField -> DISubroutineTypeField
	: FlagsField
	| CCField
	| TypesField
;

# ~~~ [ DITemplateTypeParameter ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# https://llvm.org/docs/LangRef.html#ditemplatetypeparameter

# ref: ParseDITemplateTypeParameter
#
#   ::= !DITemplateTypeParameter(name: 'Ty', type: !1)
#
#  OPTIONAL(name, MDStringField, );
#  REQUIRED(type, MDField, );

DITemplateTypeParameter -> DITemplateTypeParameter
	: '!DITemplateTypeParameter' '(' Fields=(DITemplateTypeParameterField separator ',')* ')'
;

%interface DITemplateTypeParameterField;

DITemplateTypeParameterField -> DITemplateTypeParameterField
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

DITemplateValueParameter -> DITemplateValueParameter
	: '!DITemplateValueParameter' '(' Fields=(DITemplateValueParameterField separator ',')* ')'
;

%interface DITemplateValueParameterField;

DITemplateValueParameterField -> DITemplateValueParameterField
	: TagField
	| NameField
	| TypeField
	| ValueField
;

# ~~~ [ GenericDINode ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# TODO: add link to LangRef.html.

# ref: ParseGenericDINode
#
#   ::= !GenericDINode(tag: 15, header: '...', operands: {...})
#
#  REQUIRED(tag, DwarfTagField, );
#  OPTIONAL(header, MDStringField, );
#  OPTIONAL(operands, MDFieldList, );

GenericDINode -> GenericDINode
	: '!GenericDINode' '(' Fields=(GenericDINodeField separator ',')* ')'
;

%interface GenericDINodeField;

GenericDINodeField -> GenericDINodeField
	: TagField
	| HeaderField
	| OperandsField
;

# ___ [ Specialized metadata fields ] __________________________________________

AlignField -> AlignField
	: 'align:' IntLit
;

ArgField -> ArgField
	: 'arg:' IntLit
;

AttributesField -> AttributesField
	: 'attributes:' IntLit
;

BaseTypeField -> BaseTypeField
	: 'baseType:' MDField
;

CCField -> CCField
	: 'cc:' DwarfCC
;

ChecksumField -> ChecksumField
	: 'checksum:' StringLit
;

ChecksumkindField -> ChecksumkindField
	: 'checksumkind:' ChecksumKind
;

ColumnField -> ColumnField
	: 'column:' IntLit
;

ConfigMacrosField -> ConfigMacrosField
	: 'configMacros:' StringLit
;

ContainingTypeField -> ContainingTypeField
	: 'containingType:' MDField
;

CountField -> CountField
	: 'count:' MDFieldOrInt
;

DebugInfoForProfilingField -> DebugInfoForProfilingField
	: 'debugInfoForProfiling:' BoolLit
;

DeclarationField -> DeclarationField
	: 'declaration:' MDField
;

DirectoryField -> DirectoryField
	: 'directory:' StringLit
;

DiscriminatorField -> DiscriminatorField
	: 'discriminator:' MDField
;

DiscriminatorIntField -> DiscriminatorIntField
	: 'discriminator:' IntLit
;

DwarfAddressSpaceField -> DwarfAddressSpaceField
	: 'dwarfAddressSpace:' IntLit
;

DwoIdField -> DwoIdField
	: 'dwoId:' IntLit
;

ElementsField -> ElementsField
	: 'elements:' MDField
;

EmissionKindField -> EmissionKindField
	: 'emissionKind:' EmissionKind
;

EncodingField -> EncodingField
	: 'encoding:' DwarfAttEncoding
;

EntityField -> EntityField
	: 'entity:' MDField
;

EnumsField -> EnumsField
	: 'enums:' MDField
;

ExportSymbolsField -> ExportSymbolsField
	: 'exportSymbols:' BoolLit
;

ExprField -> ExprField
	: 'expr:' MDField
;

ExtraDataField -> ExtraDataField
	: 'extraData:' MDField
;

FileField -> FileField
	: 'file:' MDField
;

FilenameField -> FilenameField
	: 'filename:' StringLit
;

FlagsField -> FlagsField
	: 'flags:' DIFlags
;

FlagsStringField -> FlagsStringField
	: 'flags:' StringLit
;

GetterField -> GetterField
	: 'getter:' StringLit
;

GlobalsField -> GlobalsField
	: 'globals:' MDField
;

HeaderField -> HeaderField
	: 'header:' StringLit
;

IdentifierField -> IdentifierField
	: 'identifier:' StringLit
;

ImportsField -> ImportsField
	: 'imports:' MDField
;

IncludePathField -> IncludePathField
	: 'includePath:' StringLit
;

InlinedAtField -> InlinedAtField
	: 'inlinedAt:' MDField
;

IsDefinitionField -> IsDefinitionField
	: 'isDefinition:' BoolLit
;

IsImplicitCodeField -> IsImplicitCodeField
	: 'isImplicitCode:' BoolLit
;

IsLocalField -> IsLocalField
	: 'isLocal:' BoolLit
;

IsOptimizedField -> IsOptimizedField
	: 'isOptimized:' BoolLit
;

IsUnsignedField -> IsUnsignedField
	: 'isUnsigned:' BoolLit
;

IsysrootField -> IsysrootField
	: 'isysroot:' StringLit
;

LanguageField -> LanguageField
	: 'language:' DwarfLang
;

LineField -> LineField
	: 'line:' IntLit
;

LinkageNameField -> LinkageNameField
	: 'linkageName:' StringLit
;

LowerBoundField -> LowerBoundField
	: 'lowerBound:' IntLit
;

MacrosField -> MacrosField
	: 'macros:' MDField
;

NameField -> NameField
	: 'name:' StringLit
;

NameTableKindField -> NameTableKindField
	: 'nameTableKind:' NameTableKind
;

NodesField -> NodesField
	: 'nodes:' MDField
;

OffsetField -> OffsetField
	: 'offset:' IntLit
;

OperandsField -> OperandsField
	: 'operands:' MDFields
;

ProducerField -> ProducerField
	: 'producer:' StringLit
;

RetainedNodesField -> RetainedNodesField
	: 'retainedNodes:' MDField
;

RetainedTypesField -> RetainedTypesField
	: 'retainedTypes:' MDField
;

RuntimeLangField -> RuntimeLangField
	: 'runtimeLang:' DwarfLang
;

RuntimeVersionField -> RuntimeVersionField
	: 'runtimeVersion:' IntLit
;

ScopeField -> ScopeField
	: 'scope:' MDField
;

ScopeLineField -> ScopeLineField
	: 'scopeLine:' IntLit
;

SetterField -> SetterField
	: 'setter:' StringLit
;

SizeField -> SizeField
	: 'size:' IntLit
;

SourceField -> SourceField
	: 'source:' StringLit
;

SplitDebugFilenameField -> SplitDebugFilenameField
	: 'splitDebugFilename:' StringLit
;

SplitDebugInliningField -> SplitDebugInliningField
	: 'splitDebugInlining:' BoolLit
;

TagField -> TagField
	: 'tag:' DwarfTag
;

TemplateParamsField -> TemplateParamsField
	: 'templateParams:' MDField
;

ThisAdjustmentField -> ThisAdjustmentField
	: 'thisAdjustment:' IntLit
;

ThrownTypesField -> ThrownTypesField
	: 'thrownTypes:' MDField
;

TypeField -> TypeField
	: 'type:' MDField
;

TypeMacinfoField -> TypeMacinfoField
	: 'type:' DwarfMacinfo
;

TypesField -> TypesField
	: 'types:' MDField
;

UnitField -> UnitField
	: 'unit:' MDField
;

ValueField -> ValueField
	: 'value:' MDField
;

ValueIntField -> ValueIntField
	: 'value:' IntLit
;

ValueStringField -> ValueStringField
	: 'value:' StringLit
;

VarField -> VarField
	: 'var:' MDField
;

VirtualIndexField -> VirtualIndexField
	: 'virtualIndex:' IntLit
;

VirtualityField -> VirtualityField
	: 'virtuality:' DwarfVirtuality
;

VtableHolderField -> VtableHolderField
	: 'vtableHolder:' MDField
;

# ___ [ Specialized metadata values ] __________________________________________

# ref: ParseMDField(MDSignedOrMDField &)

%interface MDFieldOrInt;

MDFieldOrInt -> MDFieldOrInt
	: MDField
	| IntLit
;

# ___ [ Specialized metadata enums ] ___________________________________________

ChecksumKind -> ChecksumKind
	# CSK_foo
	: checksum_kind_tok
;

# ref: ParseMDField(DIFlagField &)
#
#  ::= uint32
#  ::= DIFlagVector
#  ::= DIFlagVector '|' DIFlagFwdDecl '|' uint32 '|' DIFlagPublic

DIFlags -> DIFlags
	: Flags=(DIFlag separator '|')+
;

DIFlag -> DIFlag
	: UintLit
	# DIFlagFoo
	| di_flag_tok
;

# ref: ParseMDField(DwarfAttEncodingField &)

DwarfAttEncoding -> DwarfAttEncoding
	: UintLit
	# DW_ATE_foo
	| dwarf_att_encoding_tok
;

# ref: ParseMDField(DwarfCCField &Result)

DwarfCC -> DwarfCC
	: UintLit
	# DW_CC_foo
	| dwarf_cc_tok
;

# ref: ParseMDField(DwarfLangField &)

DwarfLang -> DwarfLang
	: UintLit
	# DW_LANG_foo
	| dwarf_lang_tok
;

# ref: ParseMDField(DwarfMacinfoTypeField &)

DwarfMacinfo -> DwarfMacinfo
	: UintLit
	# DW_MACINFO_foo
	| dwarf_macinfo_tok
;

DwarfOp -> DwarfOp
	# DW_OP_foo
	: dwarf_op_tok
;

# ref: ParseMDField(DwarfTagField &)

DwarfTag -> DwarfTag
	: UintLit
	# DW_TAG_foo
	| dwarf_tag_tok
;

# ref: ParseMDField(DwarfVirtualityField &)

DwarfVirtuality -> DwarfVirtuality
	: UintLit
	# DW_VIRTUALITY_foo
	| dwarf_virtuality_tok
;

# ref bool LLParser::ParseMDField(EmissionKindField &)

EmissionKind -> EmissionKind
	: UintLit
	# FullDebug
	| emission_kind_tok
;

# ref: bool LLParser::ParseMDField(NameTableKindField &)

NameTableKind -> NameTableKind
	: UintLit
	# GNU
	| name_table_kind_tok
;

# ___ [ Helpers ] ______________________________________________________________

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

AlignStackPair -> AlignStackPair
	: 'alignstack' '=' N=UintLit
;

# ref: parseAllocSizeArguments

AllocSize -> AllocSize
	: 'allocsize' '(' ElemSize=UintLit ')'
	| 'allocsize' '(' ElemSize=UintLit ',' N=UintLit ')'
;

# ref: ParseParameterList
#
#    ::= '(' ')'
#    ::= '(' Arg (',' Arg)* ')'
#  Arg
#    ::= Type OptionalAttributes Value OptionalAttributes

# NOTE: Args may contain '...'. The ellipsis is purely for readability.

Args -> Args
	: '...'?
	| Args=(Arg separator ',')+ (',' '...')?
;

# ref: ParseMetadataAsValue
#
#  ::= metadata i32 %local
#  ::= metadata i32 @global
#  ::= metadata i32 7
#  ::= metadata !0
#  ::= metadata !{...}
#  ::= metadata !"string"

Arg -> Arg
	: Typ=ConcreteType Attrs=ParamAttr* Val=Value
	| Typ=MetadataType Val=Metadata
;

# ref: ParseOrdering
#
#   ::= AtomicOrdering

AtomicOrdering -> AtomicOrdering
	: 'acq_rel'
	| 'acquire'
	| 'monotonic'
	| 'release'
	| 'seq_cst'
	| 'unordered'
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
	| 'cc' UintLit # TODO: Check how the AST looks like for this case.
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

Ellipsis -> Ellipsis
	: '...'
;

Exact -> Exact
	: 'exact'
;

# ref: ParseExceptionArgs

ExceptionArg -> ExceptionArg
	: Typ=ConcreteType Val=Value
	| Typ=MetadataType Val=Metadata
;

ExceptionScope -> ExceptionScope
	: NoneConst
	| LocalIdent
;

# ref: EatFastMathFlagsIfPresent

FastMathFlag -> FastMathFlag
	: 'afn'
	| 'arcp'
	| 'contract'
	| 'fast'
	| 'ninf'
	| 'nnan'
	| 'nsz'
	| 'reassoc'
;

# ref: ParseCmpPredicate

FPred -> FPred
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
	| AlignStackPair
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

# ref: ParseOptionalFunctionMetadata
#
#   ::= (!dbg !57)*

FuncMetadata -> FuncMetadata
	: MetadataAttachments=MetadataAttachment*
;

# TODO: consider removing remove GlobalAttr in favour of using (',' Section)?
# (',' Comdat)? (',' Alignment)? (',' MetadataAttachment)* as was used in the
# LangRef spec of LLVM IR. Note that the LLVM C++ parser is implemented using
# GlobalAttr, and does not follow the LangRef spec.

%interface GlobalAttr;

GlobalAttr -> GlobalAttr
	: Section
	| Comdat
	| Alignment
	#   ::= !dbg !57
	| MetadataAttachment
;

InBounds -> InBounds
	: 'inbounds'
;

# ref: ParseInstructionMetadata
#
#   ::= !dbg !42 (',' !dbg !57)*

InstMetadata -> InstMetadata
   : MetadataAttachments=(',' MetadataAttachment)+?
;

# ref: ParseCmpPredicate

IPred -> IPred
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

Label -> Label
	: Typ=LabelType Name=LocalIdent
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
;

ExternLinkage -> ExternLinkage
	: 'extern_weak'
	| 'external'
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

# TODO: inline OperandBundles to avoid OperandBundles as a node in the AST?

OperandBundles -> OperandBundles
	: ('[' OperandBundles=(OperandBundle separator ',')+ ']')?
;

OperandBundle -> OperandBundle
	: Tag=StringLit '(' Inputs=(TypeValue separator ',')* ')'
;

OverflowFlags -> OverflowFlags
	: ('nsw' | 'nuw')*
;

# ref: ParseArgumentList
#
#   ::= '(' ArgTypeListI ')'
#  ArgTypeListI
#   ::= empty
#   ::= '...'
#   ::= ArgTypeList ',' '...'
#   ::= ArgType (',' ArgType)*

# NOTE: The grammar for Params of FuncType contains Attrs and Name. However, the
# semantic check will report an error if any of those are present in the input.

Params -> Params
	: Variadic=Ellipsisopt
	| Params=(Param separator ',')+ Variadic=(',' Ellipsis)?
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

Preemption -> Preemption
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

# TODO: StackAlignment rename to AlignStack?

# ref: ParseOptionalStackAlignment
#
#   ::= empty
#   ::= 'alignstack' '(' 4 ')'
StackAlignment -> StackAlignment
	: 'alignstack' '(' N=UintLit ')'
;

# ref: ParseScope
#
#   ::= syncscope("singlethread" | "<target scope>")?

SyncScope -> SyncScope
	: 'syncscope' '(' Scope=StringLit ')'
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

TypeConst -> TypeConst
	: Typ=Type Val=Constant
;

TypeValue -> TypeValue
	: Typ=Type Val=Value
;

# ref: ParseOptionalUnnamedAddr

UnnamedAddr -> UnnamedAddr
	: 'local_unnamed_addr'
	| 'unnamed_addr'
;

UnwindTarget -> UnwindTarget
	: 'to' 'caller'
	| Label
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

Volatile -> Volatile
	: 'volatile'
;
