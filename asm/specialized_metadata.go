package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/ir/metadata"
	"github.com/pkg/errors"
)

// === [ SpecializedMDNode ] ===================================================

func (gen *generator) irSpecializedMDNode(old ast.SpecializedMDNode) (metadata.SpecializedMDNode, error) {
	switch old := old.(type) {
	case *ast.DIBasicType:
		return gen.irDIBasicType(old)
	case *ast.DICompileUnit:
		return gen.irDICompileUnit(old)
	case *ast.DICompositeType:
		return gen.irDICompositeType(old)
	case *ast.DIDerivedType:
		return gen.irDIDerivedType(old)
	case *ast.DIEnumerator:
		return gen.irDIEnumerator(old)
	case *ast.DIExpression:
		return gen.irDIExpression(old)
	case *ast.DIFile:
		return gen.irDIFile(old)
	case *ast.DIGlobalVariable:
		return gen.irDIGlobalVariable(old)
	case *ast.DIGlobalVariableExpression:
		return gen.irDIGlobalVariableExpression(old)
	case *ast.DIImportedEntity:
		return gen.irDIImportedEntity(old)
	case *ast.DILabel:
		return gen.irDILabel(old)
	case *ast.DILexicalBlock:
		return gen.irDILexicalBlock(old)
	case *ast.DILexicalBlockFile:
		return gen.irDILexicalBlockFile(old)
	case *ast.DILocalVariable:
		return gen.irDILocalVariable(old)
	case *ast.DILocation:
		return gen.irDILocation(old)
	case *ast.DIMacro:
		return gen.irDIMacro(old)
	case *ast.DIMacroFile:
		return gen.irDIMacroFile(old)
	case *ast.DIModule:
		return gen.irDIModule(old)
	case *ast.DINamespace:
		return gen.irDINamespace(old)
	case *ast.DIObjCProperty:
		return gen.irDIObjCProperty(old)
	case *ast.DISubprogram:
		return gen.irDISubprogram(old)
	case *ast.DISubrange:
		return gen.irDISubrange(old)
	case *ast.DISubroutineType:
		return gen.irDISubroutineType(old)
	case *ast.DITemplateTypeParameter:
		return gen.irDITemplateTypeParameter(old)
	case *ast.DITemplateValueParameter:
		return gen.irDITemplateValueParameter(old)
	case *ast.GenericDINode:
		return gen.irGenericDINode(old)
	default:
		panic(fmt.Errorf("support for %T not yet implemented", old))
	}
}

// --- [ DIBasicType ] ---------------------------------------------------------

func (gen *generator) irDIBasicType(old *ast.DIBasicType) (*metadata.DIBasicType, error) {
	md := &metadata.DIBasicType{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TagField:
			md.Tag = asmenum.DwarfTagFromString(oldField.Tag().Text())
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.SizeField:
			md.Size = intLit(oldField.Size())
		case *ast.AlignField:
			md.Align = intLit(oldField.Align())
		case *ast.EncodingField:
			md.Encoding = asmenum.DwarfAttEncodingFromString(oldField.Encoding().Text())
		case *ast.FlagsField:
			md.Flags = irDIFlags(oldField.Flags())
		default:
			panic(fmt.Errorf("support for DIBasicType field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DICompileUnit ] -------------------------------------------------------

func (gen *generator) irDICompileUnit(old *ast.DICompileUnit) (*metadata.DICompileUnit, error) {
	md := &metadata.DICompileUnit{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.LanguageField:
			md.Language = asmenum.DwarfLangFromString(oldField.Language().Text())
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.File = file
		case *ast.ProducerField:
			md.Producer = stringLit(oldField.Producer())
		case *ast.IsOptimizedField:
			md.IsOptimized = boolLit(oldField.IsOptimized())
		case *ast.FlagsStringField:
			md.Flags = stringLit(oldField.Flags())
		case *ast.RuntimeVersionField:
			md.RuntimeVersion = intLit(oldField.RuntimeVersion())
		case *ast.SplitDebugFilenameField:
			md.SplitDebugFilename = stringLit(oldField.SplitDebugFilename())
		case *ast.EmissionKindField:
			md.EmissionKind = asmenum.EmissionKindFromString(oldField.EmissionKind().Text())
		case *ast.EnumsField:
			enums, err := gen.irMDField(oldField.Enums())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Enums = enums
		case *ast.RetainedTypesField:
			retainedTypes, err := gen.irMDField(oldField.RetainedTypes())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.RetainedTypes = retainedTypes
		case *ast.GlobalsField:
			globals, err := gen.irMDField(oldField.Globals())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Globals = globals
		case *ast.ImportsField:
			imports, err := gen.irMDField(oldField.Imports())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Imports = imports
		case *ast.MacrosField:
			macros, err := gen.irMDField(oldField.Macros())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Macros = macros
		case *ast.DwoIdField:
			md.DwoID = intLit(oldField.DwoId())
		case *ast.SplitDebugInliningField:
			md.SplitDebugInlining = boolLit(oldField.SplitDebugInlining())
		case *ast.DebugInfoForProfilingField:
			md.DebugInfoForProfiling = boolLit(oldField.DebugInfoForProfiling())
		case *ast.NameTableKindField:
			md.NameTableKind = asmenum.NameTableKindFromString(oldField.NameTableKind().Text())
		default:
			panic(fmt.Errorf("support for DICompileUnit field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DICompositeType ] -----------------------------------------------------

func (gen *generator) irDICompositeType(old *ast.DICompositeType) (*metadata.DICompositeType, error) {
	md := &metadata.DICompositeType{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TagField:
			md.Tag = asmenum.DwarfTagFromString(oldField.Tag().Text())
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.ScopeField:
			scope, err := gen.irMDField(oldField.Scope())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Scope = scope
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.File = file
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.BaseTypeField:
			baseType, err := gen.irMDField(oldField.BaseType())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.BaseType = baseType
		case *ast.SizeField:
			md.Size = intLit(oldField.Size())
		case *ast.AlignField:
			md.Align = intLit(oldField.Align())
		case *ast.OffsetField:
			md.Offset = intLit(oldField.OffsetField())
		case *ast.FlagsField:
			md.Flags = irDIFlags(oldField.Flags())
		case *ast.ElementsField:
			elements, err := gen.irMDField(oldField.Elements())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Elements = elements
		case *ast.RuntimeLangField:
			md.RuntimeLang = asmenum.DwarfLangFromString(oldField.RuntimeLang().Text())
		case *ast.VtableHolderField:
			vtableHolder, err := gen.irMDField(oldField.VtableHolder())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.VtableHolder = vtableHolder
		case *ast.TemplateParamsField:
			templateParams, err := gen.irMDField(oldField.TemplateParams())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.TemplateParams = templateParams
		case *ast.IdentifierField:
			md.Identifier = stringLit(oldField.Identifier())
		case *ast.DiscriminatorField:
			discriminator, err := gen.irMDField(oldField.Discriminator())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Discriminator = discriminator
		default:
			panic(fmt.Errorf("support for DICompositeType field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIDerivedType ] -------------------------------------------------------

func (gen *generator) irDIDerivedType(old *ast.DIDerivedType) (*metadata.DIDerivedType, error) {
	md := &metadata.DIDerivedType{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TagField:
			md.Tag = asmenum.DwarfTagFromString(oldField.Tag().Text())
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.ScopeField:
			scope, err := gen.irMDField(oldField.Scope())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Scope = scope
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.File = file
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.BaseTypeField:
			baseType, err := gen.irMDField(oldField.BaseType())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.BaseType = baseType
		case *ast.SizeField:
			md.Size = intLit(oldField.Size())
		case *ast.AlignField:
			md.Align = intLit(oldField.Align())
		case *ast.OffsetField:
			// TODO: rename OffsetField method to Offset once https://github.com/inspirer/textmapper/issues/13 is resolved.
			md.Offset = intLit(oldField.OffsetField())
		case *ast.FlagsField:
			md.Flags = irDIFlags(oldField.Flags())
		case *ast.ExtraDataField:
			extraData, err := gen.irMDField(oldField.ExtraData())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.ExtraData = extraData
		case *ast.DwarfAddressSpaceField:
			md.DwarfAddressSpace = intLit(oldField.DwarfAddressSpace())
		default:
			panic(fmt.Errorf("support for DIDerivedType field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIEnumerator ] --------------------------------------------------------

func (gen *generator) irDIEnumerator(old *ast.DIEnumerator) (*metadata.DIEnumerator, error) {
	md := &metadata.DIEnumerator{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.ValueIntField:
			md.Value = intLit(oldField.Value())
		case *ast.IsUnsignedField:
			md.IsUnsigned = boolLit(oldField.IsUnsigned())
		default:
			panic(fmt.Errorf("support for DIEnumerator field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIExpression ] --------------------------------------------------------

func (gen *generator) irDIExpression(old *ast.DIExpression) (*metadata.DIExpression, error) {
	md := &metadata.DIExpression{}
	for _, oldField := range old.Fields() {
		field, err := gen.irDIExpressionField(oldField)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		md.Fields = append(md.Fields, field)
	}
	return md, nil
}

// ~~~ [ DIExpressionField ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irDIExpressionField(old ast.DIExpressionField) (metadata.DIExpressionField, error) {
	switch old := old.(type) {
	case *ast.UintLit:
		return metadata.UintLit(uintLit(*old)), nil
	case *ast.DwarfOp:
		return asmenum.DwarfOpFromString(old.Text()), nil
	default:
		panic(fmt.Errorf("support for DIExpression field %T not yet implemented", old))
	}
}

// --- [ DIFile ] --------------------------------------------------------------

func (gen *generator) irDIFile(old *ast.DIFile) (*metadata.DIFile, error) {
	md := &metadata.DIFile{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.FilenameField:
			md.Filename = stringLit(oldField.Filename())
		case *ast.DirectoryField:
			md.Directory = stringLit(oldField.Directory())
		case *ast.ChecksumkindField:
			md.Checksumkind = asmenum.ChecksumKindFromString(oldField.Checksumkind().Text())
		case *ast.ChecksumField:
			md.Checksum = stringLit(oldField.Checksum())
		case *ast.SourceField:
			md.Source = stringLit(oldField.Source())
		default:
			panic(fmt.Errorf("support for DIFile field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIGlobalVariable ] ----------------------------------------------------

func (gen *generator) irDIGlobalVariable(old *ast.DIGlobalVariable) (*metadata.DIGlobalVariable, error) {
	md := &metadata.DIGlobalVariable{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.ScopeField:
			scope, err := gen.irMDField(oldField.Scope())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Scope = scope
		case *ast.LinkageNameField:
			md.LinkageName = stringLit(oldField.LinkageName())
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.File = file
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.TypeField:
			typ, err := gen.irMDField(oldField.Typ())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Type = typ
		case *ast.IsLocalField:
			md.IsLocal = boolLit(oldField.IsLocal())
		case *ast.IsDefinitionField:
			md.IsDefinition = boolLit(oldField.IsDefinition())
		case *ast.TemplateParamsField:
			templateParams, err := gen.irMDField(oldField.TemplateParams())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.TemplateParams = templateParams
		case *ast.DeclarationField:
			declaration, err := gen.irMDField(oldField.Declaration())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Declaration = declaration
		case *ast.AlignField:
			md.Align = intLit(oldField.Align())
		default:
			panic(fmt.Errorf("support for DIGlobalVariable field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIGlobalVariableExpression ] ------------------------------------------

func (gen *generator) irDIGlobalVariableExpression(old *ast.DIGlobalVariableExpression) (*metadata.DIGlobalVariableExpression, error) {
	md := &metadata.DIGlobalVariableExpression{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.VarField:
			v, err := gen.irMDField(oldField.Var())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Var = v
		case *ast.ExprField:
			expr, err := gen.irMDField(oldField.Expr())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Expr = expr
		default:
			panic(fmt.Errorf("support for DIGlobalVariableExpression field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIImportedEntity ] ----------------------------------------------------

func (gen *generator) irDIImportedEntity(old *ast.DIImportedEntity) (*metadata.DIImportedEntity, error) {
	md := &metadata.DIImportedEntity{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TagField:
			md.Tag = asmenum.DwarfTagFromString(oldField.Tag().Text())
		case *ast.ScopeField:
			scope, err := gen.irMDField(oldField.Scope())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Scope = scope
		case *ast.EntityField:
			entity, err := gen.irMDField(oldField.Entity())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Entity = entity
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.File = file
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		default:
			panic(fmt.Errorf("support for DIImportedEntity field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DILabel ] -------------------------------------------------------------

func (gen *generator) irDILabel(old *ast.DILabel) (*metadata.DILabel, error) {
	panic("support for *ast.DILabel not yet implemented")
}

// --- [ DILexicalBlock ] ------------------------------------------------------

func (gen *generator) irDILexicalBlock(old *ast.DILexicalBlock) (*metadata.DILexicalBlock, error) {
	md := &metadata.DILexicalBlock{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.ScopeField:
			scope, err := gen.irMDField(oldField.Scope())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Scope = scope
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.File = file
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.ColumnField:
			md.Column = intLit(oldField.Column())
		default:
			panic(fmt.Errorf("support for DILexicalBlock field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DILexicalBlockFile ] --------------------------------------------------

func (gen *generator) irDILexicalBlockFile(old *ast.DILexicalBlockFile) (*metadata.DILexicalBlockFile, error) {
	md := &metadata.DILexicalBlockFile{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.ScopeField:
			scope, err := gen.irMDField(oldField.Scope())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Scope = scope
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.File = file
		case *ast.DiscriminatorIntField:
			md.Discriminator = intLit(oldField.Discriminator())
		default:
			panic(fmt.Errorf("support for DILexicalBlockFile field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DILocalVariable ] -----------------------------------------------------

func (gen *generator) irDILocalVariable(old *ast.DILocalVariable) (*metadata.DILocalVariable, error) {
	md := &metadata.DILocalVariable{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.ArgField:
			md.Arg = intLit(oldField.Arg())
		case *ast.ScopeField:
			scope, err := gen.irMDField(oldField.Scope())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Scope = scope
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.File = file
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.TypeField:
			typ, err := gen.irMDField(oldField.Typ())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Type = typ
		case *ast.FlagsField:
			md.Flags = irDIFlags(oldField.Flags())
		case *ast.AlignField:
			md.Align = intLit(oldField.Align())
		default:
			panic(fmt.Errorf("support for DILocalVariable field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DILocation ] ----------------------------------------------------------

func (gen *generator) irDILocation(old *ast.DILocation) (*metadata.DILocation, error) {
	md := &metadata.DILocation{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.ColumnField:
			md.Column = intLit(oldField.Column())
		case *ast.ScopeField:
			scope, err := gen.irMDField(oldField.Scope())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Scope = scope
		case *ast.InlinedAtField:
			inlinedAt, err := gen.irMDField(oldField.InlinedAt())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.InlinedAt = inlinedAt
		case *ast.IsImplicitCodeField:
			md.IsImplicitCode = boolLit(oldField.IsImplicitCode())
		default:
			panic(fmt.Errorf("support for DILocation field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIMacro ] -------------------------------------------------------------

func (gen *generator) irDIMacro(old *ast.DIMacro) (*metadata.DIMacro, error) {
	md := &metadata.DIMacro{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TypeMacinfoField:
			md.Type = asmenum.DwarfMacinfoFromString(oldField.Typ().Text())
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.ValueStringField:
			md.Value = stringLit(oldField.Value())
		default:
			panic(fmt.Errorf("support for DIMacro field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIMacroFile ] ---------------------------------------------------------

func (gen *generator) irDIMacroFile(old *ast.DIMacroFile) (*metadata.DIMacroFile, error) {
	md := &metadata.DIMacroFile{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TypeMacinfoField:
			md.Type = asmenum.DwarfMacinfoFromString(oldField.Typ().Text())
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.File = file
		case *ast.NodesField:
			nodes, err := gen.irMDField(oldField.Nodes())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Nodes = nodes
		default:
			panic(fmt.Errorf("support for DIMacroFile field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIModule ] ------------------------------------------------------------

func (gen *generator) irDIModule(old *ast.DIModule) (*metadata.DIModule, error) {
	panic("support for *ast.DIModule not yet implemented")
}

// --- [ DINamespace ] ---------------------------------------------------------

func (gen *generator) irDINamespace(old *ast.DINamespace) (*metadata.DINamespace, error) {
	md := &metadata.DINamespace{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.ScopeField:
			scope, err := gen.irMDField(oldField.Scope())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Scope = scope
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.ExportSymbolsField:
			md.ExportSymbols = boolLit(oldField.ExportSymbols())
		default:
			panic(fmt.Errorf("support for DINamespace field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIObjCProperty ] ------------------------------------------------------

func (gen *generator) irDIObjCProperty(old *ast.DIObjCProperty) (*metadata.DIObjCProperty, error) {
	panic("support for *ast.DIObjCProperty not yet implemented")
}

// --- [ DISubprogram ] --------------------------------------------------------

func (gen *generator) irDISubprogram(old *ast.DISubprogram) (*metadata.DISubprogram, error) {
	md := &metadata.DISubprogram{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.ScopeField:
			scope, err := gen.irMDField(oldField.Scope())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Scope = scope
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.LinkageNameField:
			md.LinkageName = stringLit(oldField.LinkageName())
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.File = file
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.TypeField:
			typ, err := gen.irMDField(oldField.Typ())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Type = typ
		case *ast.IsLocalField:
			md.IsLocal = boolLit(oldField.IsLocal())
		case *ast.IsDefinitionField:
			md.IsDefinition = boolLit(oldField.IsDefinition())
		case *ast.ScopeLineField:
			md.ScopeLine = intLit(oldField.ScopeLine())
		case *ast.ContainingTypeField:
			containingType, err := gen.irMDField(oldField.ContainingType())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.ContainingType = containingType
		case *ast.VirtualityField:
			md.Virtuality = asmenum.DwarfVirtualityFromString(oldField.Virtuality().Text())
		case *ast.VirtualIndexField:
			md.VirtualIndex = intLit(oldField.VirtualIndex())
		case *ast.ThisAdjustmentField:
			md.ThisAdjustment = intLit(oldField.ThisAdjustment())
		case *ast.FlagsField:
			md.Flags = irDIFlags(oldField.Flags())
		case *ast.IsOptimizedField:
			md.IsOptimized = boolLit(oldField.IsOptimized())
		case *ast.UnitField:
			unit, err := gen.irMDField(oldField.Unit())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Unit = unit
		case *ast.TemplateParamsField:
			templateParams, err := gen.irMDField(oldField.TemplateParams())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.TemplateParams = templateParams
		case *ast.DeclarationField:
			declaration, err := gen.irMDField(oldField.Declaration())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Declaration = declaration
		case *ast.RetainedNodesField:
			retainedNodes, err := gen.irMDField(oldField.RetainedNodes())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.RetainedNodes = retainedNodes
		case *ast.ThrownTypesField:
			thrownTypes, err := gen.irMDField(oldField.ThrownTypes())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.ThrownTypes = thrownTypes
		default:
			panic(fmt.Errorf("support for DISubprogram field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DISubrange ] ----------------------------------------------------------

func (gen *generator) irDISubrange(old *ast.DISubrange) (*metadata.DISubrange, error) {
	md := &metadata.DISubrange{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.CountField:
			count, err := gen.irMDFieldOrInt(oldField.Count())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Count = count
		case *ast.LowerBoundField:
			md.LowerBound = intLit(oldField.LowerBound())
		default:
			panic(fmt.Errorf("support for DISubrange field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DISubroutineType ] ----------------------------------------------------

func (gen *generator) irDISubroutineType(old *ast.DISubroutineType) (*metadata.DISubroutineType, error) {
	md := &metadata.DISubroutineType{}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.FlagsField:
			md.Flags = irDIFlags(oldField.Flags())
		case *ast.CCField:
			md.CC = asmenum.DwarfCCFromString(oldField.CC().Text())
		case *ast.TypesField:
			ts, err := gen.irMDField(oldField.Types())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Types = ts
		default:
			panic(fmt.Errorf("support for DISubroutineType field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DITemplateTypeParameter ] ---------------------------------------------

func (gen *generator) irDITemplateTypeParameter(old *ast.DITemplateTypeParameter) (*metadata.DITemplateTypeParameter, error) {
	panic("support for *ast.DITemplateTypeParameter not yet implemented")
}

// --- [ DITemplateValueParameter ] --------------------------------------------

func (gen *generator) irDITemplateValueParameter(old *ast.DITemplateValueParameter) (*metadata.DITemplateValueParameter, error) {
	panic("support for *ast.DITemplateValueParameter not yet implemented")
}

// --- [ GenericDINode ] -------------------------------------------------------

func (gen *generator) irGenericDINode(old *ast.GenericDINode) (*metadata.GenericDINode, error) {
	panic("support for *ast.GenericDINode not yet implemented")
}

// ### [ Helper functions ] ####################################################

func (gen *generator) irMDFieldOrInt(old ast.MDFieldOrInt) (metadata.MDFieldOrInt, error) {
	switch old := old.(type) {
	case ast.MDField:
		return gen.irMDField(old)
	case *ast.IntLit:
		return metadata.IntLit(intLit(*old)), nil
	default:
		panic(fmt.Errorf("support for metadata field %T not yet implemented", old))
	}
}
