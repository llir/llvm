# --- [ Specialized Metadata Nodes ] -------------------------------------------

# https://llvm.org/docs/LangRef.html#specialized-metadata-nodes

# ref: ParseSpecializedMDNode

SpecializedMDNode
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
	| 'nameTableKind:' NameTableKindField
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

DIFile
	: '!DIFile' '(' (DIFileField separator ',')* ')'
;

DIFileField
	: 'filename:' StringLit
	| 'directory:' StringLit
	| 'checksumkind:' ChecksumKind
	| 'checksum:' StringLit
	| 'source:' StringLit
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

DIBasicType
	: '!DIBasicType' '(' (DIBasicTypeField separator ',')* ')'
;

DIBasicTypeField
	: TagField
	| NameField
	| SizeField
	| AlignField
	| 'encoding:' DwarfAttEncoding
	| FlagsField
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
	| TemplateParamsField
	| DeclarationField
	| AlignField
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

DISubprogram
	: '!DISubprogram' '(' (DISubprogramField separator ',')* ')'
;

DISubprogramField
	: ScopeField
	| NameField
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
	| 'retainedNodes:' MDField
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
#   ::= !DILocation(line: 43, column: 8, scope: !5, inlinedAt: !6,
#   isImplicitCode: true)
#
#  OPTIONAL(line, LineField, );
#  OPTIONAL(column, ColumnField, );
#  REQUIRED(scope, MDField, (AllowNull false));
#  OPTIONAL(inlinedAt, MDField, );
#  OPTIONAL(isImplicitCode, MDBoolField, (false));

DILocation
	: '!DILocation' '(' (DILocationField separator ',')* ')'
;

DILocationField
	: LineField
	| ColumnField
	| ScopeField
	| 'inlinedAt:' MDField
	| 'isImplicitCode:' BoolLit
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

DILabel
	: '!DILabel' '(' (DILabelField separator ',')* ')'
;

DILabelField
	: ScopeField
	| NameField
	| FileField
	| LineField
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
	: 'flags:' (DIFlag separator '|')+
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

# ref: bool LLParser::ParseMDField(NameTableKindField &)

NameTableKindField
	: IntLit
	| NameTableKind
;
