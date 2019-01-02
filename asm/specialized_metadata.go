package asm

import (
	"fmt"
	"strconv"

	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/metadata"
	"github.com/pkg/errors"
)

// === [ SpecializedMDNode ] ===================================================

// irSpecializedMDNode returns the IR specialized metadata node corresponding to
// the given AST specialized metadata node.
func (gen *generator) irSpecializedMDNode(new metadata.Definition, old ast.SpecializedMDNode) (metadata.SpecializedNode, error) {
	switch old := old.(type) {
	case *ast.DIBasicType:
		return gen.irDIBasicType(new, old)
	case *ast.DICompileUnit:
		return gen.irDICompileUnit(new, old)
	case *ast.DICompositeType:
		return gen.irDICompositeType(new, old)
	case *ast.DIDerivedType:
		return gen.irDIDerivedType(new, old)
	case *ast.DIEnumerator:
		return gen.irDIEnumerator(new, old)
	case *ast.DIExpression:
		return gen.irDIExpression(new, old)
	case *ast.DIFile:
		return gen.irDIFile(new, old)
	case *ast.DIGlobalVariable:
		return gen.irDIGlobalVariable(new, old)
	case *ast.DIGlobalVariableExpression:
		return gen.irDIGlobalVariableExpression(new, old)
	case *ast.DIImportedEntity:
		return gen.irDIImportedEntity(new, old)
	case *ast.DILabel:
		return gen.irDILabel(new, old)
	case *ast.DILexicalBlock:
		return gen.irDILexicalBlock(new, old)
	case *ast.DILexicalBlockFile:
		return gen.irDILexicalBlockFile(new, old)
	case *ast.DILocalVariable:
		return gen.irDILocalVariable(new, old)
	case *ast.DILocation:
		return gen.irDILocation(new, old)
	case *ast.DIMacro:
		return gen.irDIMacro(new, old)
	case *ast.DIMacroFile:
		return gen.irDIMacroFile(new, old)
	case *ast.DIModule:
		return gen.irDIModule(new, old)
	case *ast.DINamespace:
		return gen.irDINamespace(new, old)
	case *ast.DIObjCProperty:
		return gen.irDIObjCProperty(new, old)
	case *ast.DISubprogram:
		return gen.irDISubprogram(new, old)
	case *ast.DISubrange:
		return gen.irDISubrange(new, old)
	case *ast.DISubroutineType:
		return gen.irDISubroutineType(new, old)
	case *ast.DITemplateTypeParameter:
		return gen.irDITemplateTypeParameter(new, old)
	case *ast.DITemplateValueParameter:
		return gen.irDITemplateValueParameter(new, old)
	case *ast.GenericDINode:
		return gen.irGenericDINode(new, old)
	default:
		panic(fmt.Errorf("support for %T not yet implemented", old))
	}
}

// --- [ DIBasicType ] ---------------------------------------------------------

// irDIBasicType returns the IR specialized metadata node DIBasicType
// corresponding to the given AST specialized metadata node DIBasicType. A new
// IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDIBasicType(new metadata.SpecializedNode, old *ast.DIBasicType) (*metadata.DIBasicType, error) {
	md, ok := new.(*metadata.DIBasicType)
	if new == nil {
		md = &metadata.DIBasicType{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DIBasicType, got %T", new))
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TagField:
			md.Tag = irDwarfTag(oldField.Tag())
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.SizeField:
			md.Size = uintLit(oldField.Size())
		case *ast.AlignField:
			md.Align = uintLit(oldField.Align())
		case *ast.EncodingField:
			md.Encoding = irDwarfAttEncoding(oldField.Encoding())
		case *ast.FlagsField:
			md.Flags = irDIFlags(oldField.Flags())
		default:
			panic(fmt.Errorf("support for DIBasicType field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DICompileUnit ] -------------------------------------------------------

// irDICompileUnit returns the IR specialized metadata node DICompileUnit
// corresponding to the given AST specialized metadata node DICompileUnit. A new
// IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDICompileUnit(new metadata.SpecializedNode, old *ast.DICompileUnit) (*metadata.DICompileUnit, error) {
	md, ok := new.(*metadata.DICompileUnit)
	if new == nil {
		md = &metadata.DICompileUnit{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DICompileUnit, got %T", new))
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.LanguageField:
			md.Language = irDwarfLang(oldField.Language())
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch file := file.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DIFile:
				md.File = file
			default:
				panic(fmt.Errorf("support for metadata DICompileUnit file field type %T not yet implemented", file))
			}
		case *ast.ProducerField:
			md.Producer = stringLit(oldField.Producer())
		case *ast.IsOptimizedField:
			md.IsOptimized = boolLit(oldField.IsOptimized())
		case *ast.FlagsStringField:
			md.Flags = stringLit(oldField.Flags())
		case *ast.RuntimeVersionField:
			md.RuntimeVersion = uintLit(oldField.RuntimeVersion())
		case *ast.SplitDebugFilenameField:
			md.SplitDebugFilename = stringLit(oldField.SplitDebugFilename())
		case *ast.EmissionKindField:
			md.EmissionKind = irEmissionKind(oldField.EmissionKind())
		case *ast.EnumsField:
			enums, err := gen.irMDField(oldField.Enums())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch enums := enums.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.Tuple:
				md.Enums = enums
			default:
				panic(fmt.Errorf("support for metadata DICompileUnit enums field type %T not yet implemented", enums))
			}
		case *ast.RetainedTypesField:
			retainedTypes, err := gen.irMDField(oldField.RetainedTypes())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch retainedTypes := retainedTypes.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.Tuple:
				md.RetainedTypes = retainedTypes
			default:
				panic(fmt.Errorf("support for metadata DICompileUnit retainedTypes field type %T not yet implemented", retainedTypes))
			}
		case *ast.GlobalsField:
			globals, err := gen.irMDField(oldField.Globals())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch globals := globals.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.Tuple:
				md.Globals = globals
			default:
				panic(fmt.Errorf("support for metadata DICompileUnit globals field type %T not yet implemented", globals))
			}
		case *ast.ImportsField:
			imports, err := gen.irMDField(oldField.Imports())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch imports := imports.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.Tuple:
				md.Imports = imports
			default:
				panic(fmt.Errorf("support for metadata DICompileUnit imports field type %T not yet implemented", imports))
			}
		case *ast.MacrosField:
			macros, err := gen.irMDField(oldField.Macros())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch macros := macros.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.Tuple:
				md.Macros = macros
			default:
				panic(fmt.Errorf("support for metadata DICompileUnit macros field type %T not yet implemented", macros))
			}
		case *ast.DwoIdField:
			md.DwoID = uintLit(oldField.DwoId())
		case *ast.SplitDebugInliningField:
			md.SplitDebugInlining = boolLit(oldField.SplitDebugInlining())
		case *ast.DebugInfoForProfilingField:
			md.DebugInfoForProfiling = boolLit(oldField.DebugInfoForProfiling())
		case *ast.NameTableKindField:
			md.NameTableKind = irNameTableKind(oldField.NameTableKind())
		default:
			panic(fmt.Errorf("support for DICompileUnit field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DICompositeType ] -----------------------------------------------------

// irDICompositeType returns the IR specialized metadata node DICompositeType
// corresponding to the given AST specialized metadata node DICompositeType. A
// new IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDICompositeType(new metadata.SpecializedNode, old *ast.DICompositeType) (*metadata.DICompositeType, error) {
	md, ok := new.(*metadata.DICompositeType)
	if new == nil {
		md = &metadata.DICompositeType{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DICompositeType, got %T", new))
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TagField:
			md.Tag = irDwarfTag(oldField.Tag())
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
			switch file := file.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DIFile:
				md.File = file
			default:
				panic(fmt.Errorf("support for metadata DICompositeType file field type %T not yet implemented", file))
			}
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.BaseTypeField:
			baseType, err := gen.irMDField(oldField.BaseType())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.BaseType = baseType
		case *ast.SizeField:
			md.Size = uintLit(oldField.Size())
		case *ast.AlignField:
			md.Align = uintLit(oldField.Align())
		case *ast.OffsetField:
			md.Offset = uintLit(oldField.OffsetField())
		case *ast.FlagsField:
			md.Flags = irDIFlags(oldField.Flags())
		case *ast.ElementsField:
			elements, err := gen.irMDField(oldField.Elements())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch elements := elements.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.Tuple:
				md.Elements = elements
			default:
				panic(fmt.Errorf("support for metadata DICompositeType elements field type %T not yet implemented", elements))
			}
		case *ast.RuntimeLangField:
			md.RuntimeLang = irDwarfLang(oldField.RuntimeLang())
		case *ast.VtableHolderField:
			vtableHolder, err := gen.irMDField(oldField.VtableHolder())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch vtableHolder := vtableHolder.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DICompositeType:
				md.VtableHolder = vtableHolder
			default:
				panic(fmt.Errorf("support for metadata DICompositeType vtableHolder field type %T not yet implemented", vtableHolder))
			}
		case *ast.TemplateParamsField:
			templateParams, err := gen.irMDField(oldField.TemplateParams())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch templateParams := templateParams.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.Tuple:
				md.TemplateParams = templateParams
			default:
				panic(fmt.Errorf("support for metadata DICompositeType templateParams field type %T not yet implemented", templateParams))
			}
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

// irDIDerivedType returns the IR specialized metadata node DIDerivedType
// corresponding to the given AST specialized metadata node DIDerivedType. A new
// IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDIDerivedType(new metadata.SpecializedNode, old *ast.DIDerivedType) (*metadata.DIDerivedType, error) {
	md, ok := new.(*metadata.DIDerivedType)
	if new == nil {
		md = &metadata.DIDerivedType{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DIDerivedType, got %T", new))
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TagField:
			md.Tag = irDwarfTag(oldField.Tag())
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
			switch file := file.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DIFile:
				md.File = file
			default:
				panic(fmt.Errorf("support for metadata DIDerivedType file field type %T not yet implemented", file))
			}
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.BaseTypeField:
			baseType, err := gen.irMDField(oldField.BaseType())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.BaseType = baseType
		case *ast.SizeField:
			md.Size = uintLit(oldField.Size())
		case *ast.AlignField:
			md.Align = uintLit(oldField.Align())
		case *ast.OffsetField:
			// TODO: rename OffsetField method to Offset once https://github.com/inspirer/textmapper/issues/13 is resolved.
			md.Offset = uintLit(oldField.OffsetField())
		case *ast.FlagsField:
			md.Flags = irDIFlags(oldField.Flags())
		case *ast.ExtraDataField:
			extraData, err := gen.irMDField(oldField.ExtraData())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.ExtraData = extraData
		case *ast.DwarfAddressSpaceField:
			md.DwarfAddressSpace = uintLit(oldField.DwarfAddressSpace())
		default:
			panic(fmt.Errorf("support for DIDerivedType field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIEnumerator ] --------------------------------------------------------

// irDIEnumerator returns the IR specialized metadata node DIEnumerator
// corresponding to the given AST specialized metadata node DIEnumerator. A new
// IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDIEnumerator(new metadata.SpecializedNode, old *ast.DIEnumerator) (*metadata.DIEnumerator, error) {
	md, ok := new.(*metadata.DIEnumerator)
	if new == nil {
		md = &metadata.DIEnumerator{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DIEnumerator, got %T", new))
	}
	isUnsigned := false
	for _, oldField := range old.Fields() {
		if oldField, ok := oldField.(*ast.IsUnsignedField); ok {
			isUnsigned = boolLit(oldField.IsUnsigned())
			break
		}
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.ValueIntField:
			if isUnsigned {
				text := oldField.Value().Text()
				x, err := strconv.ParseUint(text, 10, 64)
				if err != nil {
					panic(fmt.Errorf("unable to parse unsigned integer literal %q; %v", text, err))
				}
				md.Value = int64(x)
			} else {
				md.Value = intLit(oldField.Value())
			}
		case *ast.IsUnsignedField:
			md.IsUnsigned = boolLit(oldField.IsUnsigned())
		default:
			panic(fmt.Errorf("support for DIEnumerator field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIExpression ] --------------------------------------------------------

// irDIExpression returns the IR specialized metadata node DIExpression
// corresponding to the given AST specialized metadata node DIExpression. A new
// IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDIExpression(new metadata.SpecializedNode, old *ast.DIExpression) (*metadata.DIExpression, error) {
	md, ok := new.(*metadata.DIExpression)
	if new == nil {
		md = &metadata.DIExpression{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DIExpression, got %T", new))
	}
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

// irDIExpressionField returns the IR DIExpression field corresponding to the
// given AST DIExpression field.
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

// irDIFile returns the IR specialized metadata node DIFile corresponding to the
// given AST specialized metadata node DIFile. A new IR specialized metadata
// node correspoding to the AST specialized metadata node is created if new is
// nil, otherwise the body of new is populated.
func (gen *generator) irDIFile(new metadata.SpecializedNode, old *ast.DIFile) (*metadata.DIFile, error) {
	md, ok := new.(*metadata.DIFile)
	if new == nil {
		md = &metadata.DIFile{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DIFile, got %T", new))
	}
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

// irDIGlobalVariable returns the IR specialized metadata node DIGlobalVariable
// corresponding to the given AST specialized metadata node DIGlobalVariable. A
// new IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDIGlobalVariable(new metadata.SpecializedNode, old *ast.DIGlobalVariable) (*metadata.DIGlobalVariable, error) {
	md, ok := new.(*metadata.DIGlobalVariable)
	if new == nil {
		md = &metadata.DIGlobalVariable{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DIGlobalVariable, got %T", new))
	}
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
			switch file := file.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DIFile:
				md.File = file
			default:
				panic(fmt.Errorf("support for metadata DIGlobalVariable file field type %T not yet implemented", file))
			}
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
			switch templateParams := templateParams.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.Tuple:
				md.TemplateParams = templateParams
			default:
				panic(fmt.Errorf("support for metadata DIGlobalVariable templateParams field type %T not yet implemented", templateParams))
			}
		case *ast.DeclarationField:
			declaration, err := gen.irMDField(oldField.Declaration())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Declaration = declaration
		case *ast.AlignField:
			md.Align = uintLit(oldField.Align())
		default:
			panic(fmt.Errorf("support for DIGlobalVariable field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIGlobalVariableExpression ] ------------------------------------------

// irDIGlobalVariableExpression returns the IR specialized metadata node
// DIGlobalVariableExpression corresponding to the given AST specialized
// metadata node DIGlobalVariableExpression. A new IR specialized metadata node
// correspoding to the AST specialized metadata node is created if new is nil,
// otherwise the body of new is populated.
func (gen *generator) irDIGlobalVariableExpression(new metadata.SpecializedNode, old *ast.DIGlobalVariableExpression) (*metadata.DIGlobalVariableExpression, error) {
	md, ok := new.(*metadata.DIGlobalVariableExpression)
	if new == nil {
		md = &metadata.DIGlobalVariableExpression{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DIGlobalVariableExpression, got %T", new))
	}
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
			switch expr := expr.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DIExpression:
				md.Expr = expr
			default:
				panic(fmt.Errorf("support for metadata DIGlobalVariableExpression expr field type %T not yet implemented", expr))
			}
		default:
			panic(fmt.Errorf("support for DIGlobalVariableExpression field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIImportedEntity ] ----------------------------------------------------

// irDIImportedEntity returns the IR specialized metadata node DIImportedEntity
// corresponding to the given AST specialized metadata node DIImportedEntity. A
// new IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDIImportedEntity(new metadata.SpecializedNode, old *ast.DIImportedEntity) (*metadata.DIImportedEntity, error) {
	md, ok := new.(*metadata.DIImportedEntity)
	if new == nil {
		md = &metadata.DIImportedEntity{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DIImportedEntity, got %T", new))
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TagField:
			md.Tag = irDwarfTag(oldField.Tag())
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
			switch file := file.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DIFile:
				md.File = file
			default:
				panic(fmt.Errorf("support for metadata DIImportedEntity file field type %T not yet implemented", file))
			}
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

// irDILabel returns the IR specialized metadata node DILabel corresponding to
// the given AST specialized metadata node DILabel. A new IR specialized
// metadata node correspoding to the AST specialized metadata node is created if
// new is nil, otherwise the body of new is populated.
func (gen *generator) irDILabel(new metadata.SpecializedNode, old *ast.DILabel) (*metadata.DILabel, error) {
	md, ok := new.(*metadata.DILabel)
	if new == nil {
		md = &metadata.DILabel{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DILabel, got %T", new))
	}
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
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch file := file.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DIFile:
				md.File = file
			default:
				panic(fmt.Errorf("support for metadata DILabel file field type %T not yet implemented", file))
			}
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		default:
			panic(fmt.Errorf("support for DILabel field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DILexicalBlock ] ------------------------------------------------------

// irDILexicalBlock returns the IR specialized metadata node DILexicalBlock
// corresponding to the given AST specialized metadata node DILexicalBlock. A
// new IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDILexicalBlock(new metadata.SpecializedNode, old *ast.DILexicalBlock) (*metadata.DILexicalBlock, error) {
	md, ok := new.(*metadata.DILexicalBlock)
	if new == nil {
		md = &metadata.DILexicalBlock{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DILexicalBlock, got %T", new))
	}
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
			switch file := file.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DIFile:
				md.File = file
			default:
				panic(fmt.Errorf("support for metadata DILexicalBlock file field type %T not yet implemented", file))
			}
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

// irDILexicalBlockFile returns the IR specialized metadata node
// DILexicalBlockFile corresponding to the given AST specialized metadata node
// DILexicalBlockFile. A new IR specialized metadata node correspoding to the
// AST specialized metadata node is created if new is nil, otherwise the body of
// new is populated.
func (gen *generator) irDILexicalBlockFile(new metadata.SpecializedNode, old *ast.DILexicalBlockFile) (*metadata.DILexicalBlockFile, error) {
	md, ok := new.(*metadata.DILexicalBlockFile)
	if new == nil {
		md = &metadata.DILexicalBlockFile{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DILexicalBlockFile, got %T", new))
	}
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
			switch file := file.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DIFile:
				md.File = file
			default:
				panic(fmt.Errorf("support for metadata DILexicalBlockFile file field type %T not yet implemented", file))
			}
		case *ast.DiscriminatorIntField:
			md.Discriminator = uintLit(oldField.Discriminator())
		default:
			panic(fmt.Errorf("support for DILexicalBlockFile field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DILocalVariable ] -----------------------------------------------------

// irDILocalVariable returns the IR specialized metadata node DILocalVariable
// corresponding to the given AST specialized metadata node DILocalVariable. A
// new IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDILocalVariable(new metadata.SpecializedNode, old *ast.DILocalVariable) (*metadata.DILocalVariable, error) {
	md, ok := new.(*metadata.DILocalVariable)
	if new == nil {
		md = &metadata.DILocalVariable{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DILocalVariable, got %T", new))
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.ArgField:
			md.Arg = uintLit(oldField.Arg())
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
			switch file := file.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DIFile:
				md.File = file
			default:
				panic(fmt.Errorf("support for metadata DILocalVariable file field type %T not yet implemented", file))
			}
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
			md.Align = uintLit(oldField.Align())
		default:
			panic(fmt.Errorf("support for DILocalVariable field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DILocation ] ----------------------------------------------------------

// irDILocation returns the IR specialized metadata node DILocation
// corresponding to the given AST specialized metadata node DILocation. A new IR
// specialized metadata node correspoding to the AST specialized metadata node
// is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDILocation(new metadata.SpecializedNode, old *ast.DILocation) (*metadata.DILocation, error) {
	md, ok := new.(*metadata.DILocation)
	if new == nil {
		md = &metadata.DILocation{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DILocation, got %T", new))
	}
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

// irDIMacro returns the IR specialized metadata node DIMacro corresponding to
// the given AST specialized metadata node DIMacro. A new IR specialized
// metadata node correspoding to the AST specialized metadata node is created if
// new is nil, otherwise the body of new is populated.
func (gen *generator) irDIMacro(new metadata.SpecializedNode, old *ast.DIMacro) (*metadata.DIMacro, error) {
	md, ok := new.(*metadata.DIMacro)
	if new == nil {
		md = &metadata.DIMacro{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DIMacro, got %T", new))
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TypeMacinfoField:
			md.Type = irDwarfMacinfo(oldField.Typ())
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

// irDIMacroFile returns the IR specialized metadata node DIMacroFile
// corresponding to the given AST specialized metadata node DIMacroFile. A new
// IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDIMacroFile(new metadata.SpecializedNode, old *ast.DIMacroFile) (*metadata.DIMacroFile, error) {
	md, ok := new.(*metadata.DIMacroFile)
	if new == nil {
		md = &metadata.DIMacroFile{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DIMacroFile, got %T", new))
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TypeMacinfoField:
			md.Type = irDwarfMacinfo(oldField.Typ())
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch file := file.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DIFile:
				md.File = file
			default:
				panic(fmt.Errorf("support for metadata DIMacroFile file field type %T not yet implemented", file))
			}
		case *ast.NodesField:
			nodes, err := gen.irMDField(oldField.Nodes())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch nodes := nodes.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.Tuple:
				md.Nodes = nodes
			default:
				panic(fmt.Errorf("support for metadata DIMacroFile nodes field type %T not yet implemented", nodes))
			}
		default:
			panic(fmt.Errorf("support for DIMacroFile field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DIModule ] ------------------------------------------------------------

// irDIModule returns the IR specialized metadata node DIModule corresponding to
// the given AST specialized metadata node DIModule. A new IR specialized
// metadata node correspoding to the AST specialized metadata node is created if
// new is nil, otherwise the body of new is populated.
func (gen *generator) irDIModule(new metadata.SpecializedNode, old *ast.DIModule) (*metadata.DIModule, error) {
	md, ok := new.(*metadata.DIModule)
	if new == nil {
		md = &metadata.DIModule{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DIModule, got %T", new))
	}
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
		case *ast.ConfigMacrosField:
			md.ConfigMacros = stringLit(oldField.ConfigMacros())
		case *ast.IncludePathField:
			md.IncludePath = stringLit(oldField.IncludePath())
		case *ast.IsysrootField:
			md.Isysroot = stringLit(oldField.Isysroot())
		default:
			panic(fmt.Errorf("support for DIModule field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DINamespace ] ---------------------------------------------------------

// irDINamespace returns the IR specialized metadata node DINamespace
// corresponding to the given AST specialized metadata node DINamespace. A new
// IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDINamespace(new metadata.SpecializedNode, old *ast.DINamespace) (*metadata.DINamespace, error) {
	md, ok := new.(*metadata.DINamespace)
	if new == nil {
		md = &metadata.DINamespace{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DINamespace, got %T", new))
	}
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

// irDIObjCProperty returns the IR specialized metadata node DIObjCProperty
// corresponding to the given AST specialized metadata node DIObjCProperty. A
// new IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDIObjCProperty(new metadata.SpecializedNode, old *ast.DIObjCProperty) (*metadata.DIObjCProperty, error) {
	md, ok := new.(*metadata.DIObjCProperty)
	if new == nil {
		md = &metadata.DIObjCProperty{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DIObjCProperty, got %T", new))
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.FileField:
			file, err := gen.irMDField(oldField.File())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch file := file.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DIFile:
				md.File = file
			default:
				panic(fmt.Errorf("support for metadata DIObjCProperty file field type %T not yet implemented", file))
			}
		case *ast.LineField:
			md.Line = intLit(oldField.Line())
		case *ast.SetterField:
			md.Setter = stringLit(oldField.Setter())
		case *ast.GetterField:
			md.Getter = stringLit(oldField.Getter())
		case *ast.AttributesField:
			md.Attributes = uintLit(oldField.Attributes())
		case *ast.TypeField:
			typ, err := gen.irMDField(oldField.Typ())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Type = typ
		default:
			panic(fmt.Errorf("support for DIObjCProperty field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DISubprogram ] --------------------------------------------------------

// irDISubprogram returns the IR specialized metadata node DISubprogram
// corresponding to the given AST specialized metadata node DISubprogram. A new
// IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDISubprogram(new metadata.SpecializedNode, old *ast.DISubprogram) (*metadata.DISubprogram, error) {
	md, ok := new.(*metadata.DISubprogram)
	if new == nil {
		md = &metadata.DISubprogram{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DISubprogram, got %T", new))
	}
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
			switch file := file.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.DIFile:
				md.File = file
			default:
				panic(fmt.Errorf("support for metadata DISubprogram file field type %T not yet implemented", file))
			}
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
			md.Virtuality = irDwarfVirtuality(oldField.Virtuality())
		case *ast.VirtualIndexField:
			md.VirtualIndex = uintLit(oldField.VirtualIndex())
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
			switch templateParams := templateParams.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.Tuple:
				md.TemplateParams = templateParams
			default:
				panic(fmt.Errorf("support for metadata DISubprogram templateParams field type %T not yet implemented", templateParams))
			}
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
			switch retainedNodes := retainedNodes.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.Tuple:
				md.RetainedNodes = retainedNodes
			default:
				panic(fmt.Errorf("support for metadata DISubprogram retainedNodes field type %T not yet implemented", retainedNodes))
			}
		case *ast.ThrownTypesField:
			thrownTypes, err := gen.irMDField(oldField.ThrownTypes())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch thrownTypes := thrownTypes.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.Tuple:
				md.ThrownTypes = thrownTypes
			default:
				panic(fmt.Errorf("support for metadata DISubprogram thrownTypes field type %T not yet implemented", thrownTypes))
			}
		default:
			panic(fmt.Errorf("support for DISubprogram field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DISubrange ] ----------------------------------------------------------

// irDISubrange returns the IR specialized metadata node DISubrange
// corresponding to the given AST specialized metadata node DISubrange. A new IR
// specialized metadata node correspoding to the AST specialized metadata node
// is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDISubrange(new metadata.SpecializedNode, old *ast.DISubrange) (*metadata.DISubrange, error) {
	md, ok := new.(*metadata.DISubrange)
	if new == nil {
		md = &metadata.DISubrange{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DISubrange, got %T", new))
	}
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

// irDISubroutineType returns the IR specialized metadata node DISubroutineType
// corresponding to the given AST specialized metadata node DISubroutineType. A
// new IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irDISubroutineType(new metadata.SpecializedNode, old *ast.DISubroutineType) (*metadata.DISubroutineType, error) {
	md, ok := new.(*metadata.DISubroutineType)
	if new == nil {
		md = &metadata.DISubroutineType{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DISubroutineType, got %T", new))
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.FlagsField:
			md.Flags = irDIFlags(oldField.Flags())
		case *ast.CCField:
			md.CC = irDwarfCC(oldField.CC())
		case *ast.TypesField:
			ts, err := gen.irMDField(oldField.Types())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch ts := ts.(type) {
			case *metadata.NullLit:
				// nothing to do.
			case *metadata.Tuple:
				md.Types = ts
			default:
				panic(fmt.Errorf("support for metadata DISubroutineType types field type %T not yet implemented", ts))
			}
		default:
			panic(fmt.Errorf("support for DISubroutineType field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DITemplateTypeParameter ] ---------------------------------------------

// irDITemplateTypeParameter returns the IR specialized metadata node
// DITemplateTypeParameter corresponding to the given AST specialized metadata
// node DITemplateTypeParameter. A new IR specialized metadata node correspoding
// to the AST specialized metadata node is created if new is nil, otherwise the
// body of new is populated.
func (gen *generator) irDITemplateTypeParameter(new metadata.SpecializedNode, old *ast.DITemplateTypeParameter) (*metadata.DITemplateTypeParameter, error) {
	md, ok := new.(*metadata.DITemplateTypeParameter)
	if new == nil {
		md = &metadata.DITemplateTypeParameter{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DITemplateTypeParameter, got %T", new))
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.TypeField:
			typ, err := gen.irMDField(oldField.Typ())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Type = typ
		default:
			panic(fmt.Errorf("support for DITemplateTypeParameter field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ DITemplateValueParameter ] --------------------------------------------

// irDITemplateValueParameter returns the IR specialized metadata node
// DITemplateValueParameter corresponding to the given AST specialized metadata
// node DITemplateValueParameter. A new IR specialized metadata node
// correspoding to the AST specialized metadata node is created if new is nil,
// otherwise the body of new is populated.
func (gen *generator) irDITemplateValueParameter(new metadata.SpecializedNode, old *ast.DITemplateValueParameter) (*metadata.DITemplateValueParameter, error) {
	md, ok := new.(*metadata.DITemplateValueParameter)
	if new == nil {
		md = &metadata.DITemplateValueParameter{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.DITemplateValueParameter, got %T", new))
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TagField:
			md.Tag = irDwarfTag(oldField.Tag())
		case *ast.NameField:
			md.Name = stringLit(oldField.Name())
		case *ast.TypeField:
			typ, err := gen.irMDField(oldField.Typ())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Type = typ
		case *ast.ValueField:
			value, err := gen.irMDField(oldField.Value())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			md.Value = value
		default:
			panic(fmt.Errorf("support for DITemplateValueParameter field %T not yet implemented", old))
		}
	}
	return md, nil
}

// --- [ GenericDINode ] -------------------------------------------------------

// irGenericDINode returns the IR specialized metadata node GenericDINode
// corresponding to the given AST specialized metadata node GenericDINode. A new
// IR specialized metadata node correspoding to the AST specialized metadata
// node is created if new is nil, otherwise the body of new is populated.
func (gen *generator) irGenericDINode(new metadata.SpecializedNode, old *ast.GenericDINode) (*metadata.GenericDINode, error) {
	md, ok := new.(*metadata.GenericDINode)
	if new == nil {
		md = &metadata.GenericDINode{MetadataID: -1}
	} else if !ok {
		panic(fmt.Errorf("invalid IR specialized metadata node for AST specialized metadata node; expected *metadata.GenericDINode, got %T", new))
	}
	for _, oldField := range old.Fields() {
		switch oldField := oldField.(type) {
		case *ast.TagField:
			md.Tag = irDwarfTag(oldField.Tag())
		case *ast.HeaderField:
			md.Header = stringLit(oldField.Header())
		case *ast.OperandsField:
			for _, field := range oldField.Operands() {
				operand, err := gen.irMDField(field)
				if err != nil {
					return nil, errors.WithStack(err)
				}
				md.Operands = append(md.Operands, operand)
			}
		default:
			panic(fmt.Errorf("support for GenericDINode field %T not yet implemented", old))
		}
	}
	return md, nil
}

// ### [ Helper functions ] ####################################################

// irMDFieldOrInt returns the IR metadata field or integer corresponding to the
// given AST metadata field or integer.
func (gen *generator) irMDFieldOrInt(old ast.MDFieldOrInt) (metadata.FieldOrInt, error) {
	switch old := old.(type) {
	case ast.MDField:
		return gen.irMDField(old)
	case *ast.IntLit:
		return metadata.IntLit(intLit(*old)), nil
	default:
		panic(fmt.Errorf("support for metadata field %T not yet implemented", old))
	}
}

// irDIFlags returns the IR debug info flags corresponding to the given AST
// debug info flags.
func irDIFlags(old ast.DIFlags) enum.DIFlag {
	var flags enum.DIFlag
	for _, oldFlag := range old.Flags() {
		flag := irDIFlag(oldFlag)
		flags |= flag
	}
	return flags
}

// irDIFlag returns the IR debug info flag corresponding to the given AST debug
// info flag.
func irDIFlag(old ast.DIFlag) enum.DIFlag {
	switch old := old.(type) {
	case *ast.DIFlagEnum:
		return asmenum.DIFlagFromString(old.Text())
	case *ast.DIFlagInt:
		return enum.DIFlag(uintLit(old.UintLit()))
	default:
		panic(fmt.Errorf("support for debug info flag %T not yet implemented", old))
	}
}

// irDwarfAttEncoding returns the IR Dwarf attribute encoding corresponding to
// the given AST Dwarf attribute encoding.
func irDwarfAttEncoding(old ast.DwarfAttEncoding) enum.DwarfAttEncoding {
	switch old := old.(type) {
	case *ast.DwarfAttEncodingEnum:
		return asmenum.DwarfAttEncodingFromString(old.Text())
	case *ast.DwarfAttEncodingInt:
		return enum.DwarfAttEncoding(uintLit(old.UintLit()))
	default:
		panic(fmt.Errorf("support for Dwarf attribute encoding %T not yet implemented", old))
	}
}

// irDwarfCC returns the IR Dwarf calling convention corresponding to the given
// AST Dwarf calling convention.
func irDwarfCC(old ast.DwarfCC) enum.DwarfCC {
	switch old := old.(type) {
	case *ast.DwarfCCEnum:
		return asmenum.DwarfCCFromString(old.Text())
	case *ast.DwarfCCInt:
		return enum.DwarfCC(uintLit(old.UintLit()))
	default:
		panic(fmt.Errorf("support for Dwarf calling convention %T not yet implemented", old))
	}
}

// irDwarfLang returns the IR Dwarf language corresponding to the given AST
// Dwarf language.
func irDwarfLang(old ast.DwarfLang) enum.DwarfLang {
	switch old := old.(type) {
	case *ast.DwarfLangEnum:
		return asmenum.DwarfLangFromString(old.Text())
	case *ast.DwarfLangInt:
		return enum.DwarfLang(uintLit(old.UintLit()))
	default:
		panic(fmt.Errorf("support for Dwarf language %T not yet implemented", old))
	}
}

// irDwarfMacinfo returns the IR Dwarf Macinfo corresponding to the given AST
// Dwarf Macinfo.
func irDwarfMacinfo(old ast.DwarfMacinfo) enum.DwarfMacinfo {
	switch old := old.(type) {
	case *ast.DwarfMacinfoEnum:
		return asmenum.DwarfMacinfoFromString(old.Text())
	case *ast.DwarfMacinfoInt:
		return enum.DwarfMacinfo(uintLit(old.UintLit()))
	default:
		panic(fmt.Errorf("support for Dwarf Macinfo %T not yet implemented", old))
	}
}

// irDwarfTag returns the IR Dwarf tag corresponding to the given AST Dwarf tag.
func irDwarfTag(old ast.DwarfTag) enum.DwarfTag {
	switch old := old.(type) {
	case *ast.DwarfTagEnum:
		return asmenum.DwarfTagFromString(old.Text())
	case *ast.DwarfTagInt:
		return enum.DwarfTag(uintLit(old.UintLit()))
	default:
		panic(fmt.Errorf("support for Dwarf tag %T not yet implemented", old))
	}
}

// irDwarfVirtuality returns the IR Dwarf virtuality corresponding to the given
// AST Dwarf virtuality.
func irDwarfVirtuality(old ast.DwarfVirtuality) enum.DwarfVirtuality {
	switch old := old.(type) {
	case *ast.DwarfVirtualityEnum:
		return asmenum.DwarfVirtualityFromString(old.Text())
	case *ast.DwarfVirtualityInt:
		return enum.DwarfVirtuality(uintLit(old.UintLit()))
	default:
		panic(fmt.Errorf("support for Dwarf virtuality %T not yet implemented", old))
	}
}

// irEmissionKind returns the IR emission kind corresponding to the given AST
// emission kind.
func irEmissionKind(old ast.EmissionKind) enum.EmissionKind {
	switch old := old.(type) {
	case *ast.EmissionKindEnum:
		return asmenum.EmissionKindFromString(old.Text())
	case *ast.EmissionKindInt:
		return enum.EmissionKind(uintLit(old.UintLit()))
	default:
		panic(fmt.Errorf("support for emission kind %T not yet implemented", old))
	}
}

// irNameTableKind returns the IR name table kind corresponding to the given AST
// name table kind.
func irNameTableKind(old ast.NameTableKind) enum.NameTableKind {
	switch old := old.(type) {
	case *ast.NameTableKindEnum:
		return asmenum.NameTableKindFromString(old.Text())
	case *ast.NameTableKindInt:
		return enum.NameTableKind(uintLit(old.UintLit()))
	default:
		panic(fmt.Errorf("support for name table kind %T not yet implemented", old))
	}
}
