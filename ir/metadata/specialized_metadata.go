package metadata

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/enum"
)

// ~~~ [ DIBasicType ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DIBasicType is a specialized metadata node.
type DIBasicType struct {
	Tag      enum.DwarfTag         // optional; zero value if not present.
	Name     string                // optional; empty if not present.
	Size     int64                 // optional; zero value if not present.
	Align    int64                 // optional; zero value if not present.
	Encoding enum.DwarfAttEncoding // optional; zero value if not present.
	Flags    enum.DIFlag           // optional.
}

// String returns a string representation of the specialized metadata node.
func (md *DIBasicType) String() string {
	// '!DIBasicType' '(' Fields=(DIBasicTypeField separator ',')* ')'
	var fields []string
	if md.Tag != 0 {
		field := fmt.Sprintf("tag: %s", md.Tag)
		fields = append(fields, field)
	}
	if len(md.Name) > 0 {
		field := fmt.Sprintf("name: %s", quote(md.Name))
		fields = append(fields, field)
	}
	if md.Size != 0 {
		field := fmt.Sprintf("size: %d", md.Size)
		fields = append(fields, field)
	}
	if md.Align != 0 {
		field := fmt.Sprintf("align: %d", md.Align)
		fields = append(fields, field)
	}
	if md.Encoding != 0 {
		field := fmt.Sprintf("encoding: %s", md.Encoding)
		fields = append(fields, field)
	}
	if md.Flags != 0 {
		field := fmt.Sprintf("flags: %s", diFlagsString(md.Flags))
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DIBasicType(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DICompileUnit ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DICompileUnit is a specialized metadata node.
type DICompileUnit struct {
	Language              enum.DwarfLang     // required.
	File                  MDField            // required.
	Producer              string             // optional; empty if not present.
	IsOptimized           bool               // optional; zero value if not present.
	Flags                 string             // optional; empty if not present.
	RuntimeVersion        int64              // optional; zero value if not present.
	SplitDebugFilename    string             // optional; empty if not present.
	EmissionKind          enum.EmissionKind  // optional; zero value if not present.
	Enums                 MDField            // optional; nil if not present.
	RetainedTypes         MDField            // optional; nil if not present.
	Globals               MDField            // optional; nil if not present.
	Imports               MDField            // optional; nil if not present.
	Macros                MDField            // optional; nil if not present.
	DwoID                 int64              // optional; zero value if not present.
	SplitDebugInlining    bool               // optional; zero value if not present.
	DebugInfoForProfiling bool               // optional; zero value if not present.
	NameTableKind         enum.NameTableKind // optional; zero value if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DICompileUnit) String() string {
	// '!DICompileUnit' '(' Fields=(DICompileUnitField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("language: %s", md.Language)
	fields = append(fields, field)
	field = fmt.Sprintf("file: %s", md.File)
	fields = append(fields, field)
	if len(md.Producer) > 0 {
		field = fmt.Sprintf("producer: %s", quote(md.Producer))
		fields = append(fields, field)
	}
	if md.IsOptimized {
		field = fmt.Sprintf("isOptimized: %t", md.IsOptimized)
		fields = append(fields, field)
	}
	if len(md.Flags) > 0 {
		field = fmt.Sprintf("flags: %s", quote(md.Flags))
		fields = append(fields, field)
	}
	if md.RuntimeVersion != 0 {
		field = fmt.Sprintf("runtimeVersion: %d", md.RuntimeVersion)
		fields = append(fields, field)
	}
	if len(md.SplitDebugFilename) > 0 {
		field = fmt.Sprintf("splitDebugFilename: %s", quote(md.SplitDebugFilename))
		fields = append(fields, field)
	}
	if md.EmissionKind != 0 {
		field = fmt.Sprintf("emissionKind: %s", md.EmissionKind)
		fields = append(fields, field)
	}
	if md.Enums != nil {
		field = fmt.Sprintf("enums: %s", md.Enums)
		fields = append(fields, field)
	}
	if md.RetainedTypes != nil {
		field = fmt.Sprintf("retainedTypes: %s", md.RetainedTypes)
		fields = append(fields, field)
	}
	if md.Globals != nil {
		field = fmt.Sprintf("globals: %s", md.Globals)
		fields = append(fields, field)
	}
	if md.Imports != nil {
		field = fmt.Sprintf("imports: %s", md.Imports)
		fields = append(fields, field)
	}
	if md.Macros != nil {
		field = fmt.Sprintf("macros: %s", md.Macros)
		fields = append(fields, field)
	}
	if md.DwoID != 0 {
		field = fmt.Sprintf("dwoId: %d", md.DwoID)
		fields = append(fields, field)
	}
	if md.SplitDebugInlining {
		field = fmt.Sprintf("splitDebugInlining: %t", md.SplitDebugInlining)
		fields = append(fields, field)
	}
	if md.DebugInfoForProfiling {
		field = fmt.Sprintf("debugInfoForProfiling: %t", md.DebugInfoForProfiling)
		fields = append(fields, field)
	}
	if md.NameTableKind != 0 {
		field = fmt.Sprintf("nameTableKind: %s", md.NameTableKind)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DICompileUnit(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DICompositeType ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DICompositeType is a specialized metadata node.
type DICompositeType struct {
	Tag            enum.DwarfTag  // required.
	Name           string         // optional; empty if not present.
	Scope          MDField        // optional; nil if not present.
	File           MDField        // optional; nil if not present.
	Line           int64          // optional; zero value if not present.
	BaseType       MDField        // optional; nil if not present.
	Size           int64          // optional; zero value if not present.
	Align          int64          // optional; zero value if not present.
	Offset         int64          // optional; zero value if not present.
	Flags          enum.DIFlag    // optional.
	Elements       MDField        // optional; nil if not present.
	RuntimeLang    enum.DwarfLang // optional; zero value if not present.
	VtableHolder   MDField        // optional; nil if not present.
	TemplateParams MDField        // optional; nil if not present.
	Identifier     string         // optional; empty if not present.
	Discriminator  MDField        // optional; nil if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DICompositeType) String() string {
	// '!DICompositeType' '(' Fields=(DICompositeTypeField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("tag: %s", md.Tag)
	fields = append(fields, field)
	if len(md.Name) > 0 {
		field := fmt.Sprintf("name: %s", quote(md.Name))
		fields = append(fields, field)
	}
	if md.Scope != nil {
		field := fmt.Sprintf("scope: %s", md.Scope)
		fields = append(fields, field)
	}
	if md.File != nil {
		field := fmt.Sprintf("file: %s", md.File)
		fields = append(fields, field)
	}
	if md.Line != 0 {
		field := fmt.Sprintf("line: %d", md.Line)
		fields = append(fields, field)
	}
	if md.BaseType != nil {
		field := fmt.Sprintf("baseType: %s", md.BaseType)
		fields = append(fields, field)
	}
	if md.Size != 0 {
		field := fmt.Sprintf("size: %d", md.Size)
		fields = append(fields, field)
	}
	if md.Align != 0 {
		field := fmt.Sprintf("align: %d", md.Align)
		fields = append(fields, field)
	}
	if md.Offset != 0 {
		field := fmt.Sprintf("offset: %d", md.Offset)
		fields = append(fields, field)
	}
	if md.Flags != 0 {
		field = fmt.Sprintf("flags: %s", diFlagsString(md.Flags))
		fields = append(fields, field)
	}
	if md.Elements != nil {
		field := fmt.Sprintf("elements: %s", md.Elements)
		fields = append(fields, field)
	}
	if md.RuntimeLang != 0 {
		field := fmt.Sprintf("runtimeLang: %s", md.RuntimeLang)
		fields = append(fields, field)
	}
	if md.VtableHolder != nil {
		field := fmt.Sprintf("vtableHolder: %s", md.VtableHolder)
		fields = append(fields, field)
	}
	if md.TemplateParams != nil {
		field := fmt.Sprintf("templateParams: %s", md.TemplateParams)
		fields = append(fields, field)
	}
	if len(md.Identifier) > 0 {
		field := fmt.Sprintf("identifier: %s", quote(md.Identifier))
		fields = append(fields, field)
	}
	if md.Discriminator != nil {
		field := fmt.Sprintf("discriminator: %s", md.Discriminator)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DICompositeType(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DIDerivedType ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DIDerivedType is a specialized metadata node.
type DIDerivedType struct {
	Tag               enum.DwarfTag // required.
	Name              string        // optional; empty if not present.
	Scope             MDField       // optional; nil if not present.
	File              MDField       // optional; nil if not present.
	Line              int64         // optional; zero value if not present.
	BaseType          MDField       // required.
	Size              int64         // optional; zero value if not present.
	Align             int64         // optional; zero value if not present.
	Offset            int64         // optional; zero value if not present.
	Flags             enum.DIFlag   // optional.
	ExtraData         MDField       // optional; nil if not present.
	DwarfAddressSpace int64         // optional; zero value if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DIDerivedType) String() string {
	// '!DIDerivedType' '(' Fields=(DIDerivedTypeField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("tag: %s", md.Tag)
	fields = append(fields, field)
	if len(md.Name) > 0 {
		field := fmt.Sprintf("name: %s", quote(md.Name))
		fields = append(fields, field)
	}
	if md.Scope != nil {
		field := fmt.Sprintf("scope: %s", md.Scope)
		fields = append(fields, field)
	}
	if md.File != nil {
		field := fmt.Sprintf("file: %s", md.File)
		fields = append(fields, field)
	}
	if md.Line != 0 {
		field := fmt.Sprintf("line: %d", md.Line)
		fields = append(fields, field)
	}
	field = fmt.Sprintf("baseType: %s", md.BaseType)
	fields = append(fields, field)
	if md.Size != 0 {
		field := fmt.Sprintf("size: %d", md.Size)
		fields = append(fields, field)
	}
	if md.Align != 0 {
		field := fmt.Sprintf("align: %d", md.Align)
		fields = append(fields, field)
	}
	if md.Offset != 0 {
		field := fmt.Sprintf("offset: %d", md.Offset)
		fields = append(fields, field)
	}
	if md.Flags != 0 {
		field = fmt.Sprintf("flags: %s", diFlagsString(md.Flags))
		fields = append(fields, field)
	}
	if md.ExtraData != nil {
		field := fmt.Sprintf("extraData: %s", md.ExtraData)
		fields = append(fields, field)
	}
	if md.DwarfAddressSpace != 0 {
		field := fmt.Sprintf("dwarfAddressSpace: %d", md.DwarfAddressSpace)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DIDerivedType(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DIEnumerator ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DIEnumerator is a specialized metadata node.
type DIEnumerator struct {
	Name       string // required.
	Value      int64  // required.
	IsUnsigned bool   // optional; zero value if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DIEnumerator) String() string {
	// '!DIEnumerator' '(' Fields=(DIEnumeratorField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("name: %s", quote(md.Name))
	fields = append(fields, field)
	if md.IsUnsigned {
		field = fmt.Sprintf("value: %d", uint64(md.Value))
	} else {
		field = fmt.Sprintf("value: %d", md.Value)
	}
	fields = append(fields, field)
	if md.IsUnsigned {
		field := fmt.Sprintf("isUnsigned: %t", md.IsUnsigned)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DIEnumerator(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DIExpression ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DIExpression is a specialized metadata node.
type DIExpression struct {
	Fields []DIExpressionField
}

// String returns a string representation of the specialized metadata node.
func (md *DIExpression) String() string {
	// '!DIExpression' '(' Fields=(DIExpressionField separator ',')* ')'
	buf := &strings.Builder{}
	buf.WriteString("!DIExpression(")
	for i, field := range md.Fields {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(field.String())
	}
	buf.WriteString(")")
	return buf.String()
}

// ~~~ [ DIFile ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DIFile is a specialized metadata node.
type DIFile struct {
	Filename     string            // required.
	Directory    string            // required.
	Checksumkind enum.ChecksumKind // optional; zero value if not present.
	Checksum     string            // optional; empty if not present.
	Source       string            // optional; empty if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DIFile) String() string {
	// '!DIFile' '(' Fields=(DIFileField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("filename: %s", quote(md.Filename))
	fields = append(fields, field)
	field = fmt.Sprintf("directory: %s", quote(md.Directory))
	fields = append(fields, field)
	if md.Checksumkind != 0 {
		field := fmt.Sprintf("checksumkind: %s", md.Checksumkind)
		fields = append(fields, field)
	}
	if len(md.Checksum) > 0 {
		field := fmt.Sprintf("checksum: %s", quote(md.Checksum))
		fields = append(fields, field)
	}
	if len(md.Source) > 0 {
		field := fmt.Sprintf("source: %s", quote(md.Source))
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DIFile(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DIGlobalVariable ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DIGlobalVariable is a specialized metadata node.
type DIGlobalVariable struct {
	Name           string  // required.
	Scope          MDField // optional; nil if not present.
	LinkageName    string  // optional; empty if not present.
	File           MDField // optional; nil if not present.
	Line           int64   // optional; zero value if not present.
	Type           MDField // optional; nil if not present.
	IsLocal        bool    // optional; zero value if not present.
	IsDefinition   bool    // optional; zero value if not present.
	TemplateParams MDField // optional; nil if not present.
	Declaration    MDField // optional; nil if not present.
	Align          int64   // optional; zero value if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DIGlobalVariable) String() string {
	// '!DIGlobalVariable' '(' Fields=(DIGlobalVariableField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("name: %s", quote(md.Name))
	fields = append(fields, field)
	if md.Scope != nil {
		field := fmt.Sprintf("scope: %s", md.Scope)
		fields = append(fields, field)
	}
	if len(md.LinkageName) > 0 {
		field := fmt.Sprintf("linkageName: %s", quote(md.LinkageName))
		fields = append(fields, field)
	}
	if md.File != nil {
		field := fmt.Sprintf("file: %s", md.File)
		fields = append(fields, field)
	}
	if md.Line != 0 {
		field := fmt.Sprintf("line: %d", md.Line)
		fields = append(fields, field)
	}
	if md.Type != nil {
		field := fmt.Sprintf("type: %s", md.Type)
		fields = append(fields, field)
	}
	if md.IsLocal {
		field = fmt.Sprintf("isLocal: %t", md.IsLocal)
		fields = append(fields, field)
	}
	if md.IsDefinition {
		field := fmt.Sprintf("isDefinition: %t", md.IsDefinition)
		fields = append(fields, field)
	}
	if md.TemplateParams != nil {
		field := fmt.Sprintf("templateParams: %s", md.TemplateParams)
		fields = append(fields, field)
	}
	if md.Declaration != nil {
		field := fmt.Sprintf("declaration: %s", md.Declaration)
		fields = append(fields, field)
	}
	if md.Align != 0 {
		field := fmt.Sprintf("align: %d", md.Align)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DIGlobalVariable(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DIGlobalVariableExpression ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DIGlobalVariableExpression is a specialized metadata node.
type DIGlobalVariableExpression struct {
	Var MDField // required.

	// Note, the C++ source code of LLVM states that "expr:" is a required field,
	// however, Clang is known to output DIGlobalVariableExpression specialized
	// metadata nodes only containing "var:"; e.g. from `cat.ll`:
	//
	//    !0 = !DIGlobalVariableExpression(var: !1)

	Expr MDField // required.
}

// String returns a string representation of the specialized metadata node.
func (md *DIGlobalVariableExpression) String() string {
	// '!DIGlobalVariableExpression' '(' Fields=(DIGlobalVariableExpressionField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("var: %s", md.Var)
	fields = append(fields, field)
	// NOTE: Should be required. Thus nil check should not be needed. However,
	// Clang outputs `!0 = !DIGlobalVariableExpression(var: !1)` in cat.ll.
	if md.Expr != nil {
		field = fmt.Sprintf("expr: %s", md.Expr)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DIGlobalVariableExpression(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DIImportedEntity ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DIImportedEntity is a specialized metadata node.
type DIImportedEntity struct {
	Tag    enum.DwarfTag // required.
	Scope  MDField       // required.
	Entity MDField       // optional; nil if not present.
	File   MDField       // optional; nil if not present.
	Line   int64         // optional; zero value if not present.
	Name   string        // optional; empty if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DIImportedEntity) String() string {
	// '!DIImportedEntity' '(' Fields=(DIImportedEntityField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("tag: %s", md.Tag)
	fields = append(fields, field)
	field = fmt.Sprintf("scope: %s", md.Scope)
	fields = append(fields, field)
	if md.Entity != nil {
		field := fmt.Sprintf("entity: %s", md.Entity)
		fields = append(fields, field)
	}
	if md.File != nil {
		field := fmt.Sprintf("file: %s", md.File)
		fields = append(fields, field)
	}
	if md.Line != 0 {
		field := fmt.Sprintf("line: %d", md.Line)
		fields = append(fields, field)
	}
	if len(md.Name) > 0 {
		field := fmt.Sprintf("name: %s", quote(md.Name))
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DIImportedEntity(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DILabel ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DILabel is a specialized metadata node.
type DILabel struct {
	Scope MDField // required.
	Name  string  // required.
	File  MDField // required.
	Line  int64   // required.
}

// String returns a string representation of the specialized metadata node.
func (md *DILabel) String() string {
	// '!DILabel' '(' Fields=(DILabelField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("scope: %s", md.Scope)
	fields = append(fields, field)
	field = fmt.Sprintf("name: %s", quote(md.Name))
	fields = append(fields, field)
	field = fmt.Sprintf("file: %s", md.File)
	fields = append(fields, field)
	field = fmt.Sprintf("line: %d", md.Line)
	fields = append(fields, field)
	return fmt.Sprintf("!DILabel(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DILexicalBlock ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DILexicalBlock is a specialized metadata node.
type DILexicalBlock struct {
	Scope  MDField // required.
	File   MDField // optional; nil if not present.
	Line   int64   // optional; zero value if not present.
	Column int64   // optional; zero value if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DILexicalBlock) String() string {
	// '!DILexicalBlock' '(' Fields=(DILexicalBlockField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("scope: %s", md.Scope)
	fields = append(fields, field)
	if md.File != nil {
		field := fmt.Sprintf("file: %s", md.File)
		fields = append(fields, field)
	}
	if md.Line != 0 {
		field := fmt.Sprintf("line: %d", md.Line)
		fields = append(fields, field)
	}
	if md.Column != 0 {
		field := fmt.Sprintf("column: %d", md.Column)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DILexicalBlock(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DILexicalBlockFile ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DILexicalBlockFile is a specialized metadata node.
type DILexicalBlockFile struct {
	Scope         MDField // required.
	File          MDField // optional; nil if not present.
	Discriminator int64   // required.
}

// String returns a string representation of the specialized metadata node.
func (md *DILexicalBlockFile) String() string {
	// '!DILexicalBlockFile' '(' Fields=(DILexicalBlockFileField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("scope: %s", md.Scope)
	fields = append(fields, field)
	if md.File != nil {
		field := fmt.Sprintf("file: %s", md.File)
		fields = append(fields, field)
	}
	field = fmt.Sprintf("discriminator: %d", md.Discriminator)
	fields = append(fields, field)
	return fmt.Sprintf("!DILexicalBlockFile(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DILocalVariable ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DILocalVariable is a specialized metadata node.
type DILocalVariable struct {
	Name  string      // optional; empty if not present.
	Arg   int64       // optional; zero value if not present.
	Scope MDField     // required.
	File  MDField     // optional; nil if not present.
	Line  int64       // optional; zero value if not present.
	Type  MDField     // optional; nil if not present.
	Flags enum.DIFlag // optional.
	Align int64       // optional; zero value if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DILocalVariable) String() string {
	// '!DILocalVariable' '(' Fields=(DILocalVariableField separator ',')* ')'
	var fields []string
	if len(md.Name) > 0 {
		field := fmt.Sprintf("name: %s", quote(md.Name))
		fields = append(fields, field)
	}
	if md.Arg != 0 {
		field := fmt.Sprintf("arg: %d", md.Arg)
		fields = append(fields, field)
	}
	field := fmt.Sprintf("scope: %s", md.Scope)
	fields = append(fields, field)
	if md.File != nil {
		field := fmt.Sprintf("file: %s", md.File)
		fields = append(fields, field)
	}
	if md.Line != 0 {
		field := fmt.Sprintf("line: %d", md.Line)
		fields = append(fields, field)
	}
	if md.Type != nil {
		field := fmt.Sprintf("type: %s", md.Type)
		fields = append(fields, field)
	}
	if md.Flags != 0 {
		field = fmt.Sprintf("flags: %s", diFlagsString(md.Flags))
		fields = append(fields, field)
	}
	if md.Align != 0 {
		field := fmt.Sprintf("align: %d", md.Align)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DILocalVariable(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DILocation ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DILocation is a specialized metadata node.
type DILocation struct {
	Line           int64   // optional; zero value if not present.
	Column         int64   // optional; zero value if not present.
	Scope          MDField // required.
	InlinedAt      MDField // optional; nil if not present.
	IsImplicitCode bool    // optional; zero value if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DILocation) String() string {
	// '!DILocation' '(' Fields=(DILocationField separator ',')* ')'
	var fields []string
	if md.Line != 0 {
		field := fmt.Sprintf("line: %d", md.Line)
		fields = append(fields, field)
	}
	if md.Column != 0 {
		field := fmt.Sprintf("column: %d", md.Column)
		fields = append(fields, field)
	}
	field := fmt.Sprintf("scope: %s", md.Scope)
	fields = append(fields, field)
	if md.InlinedAt != nil {
		field := fmt.Sprintf("inlinedAt: %s", md.InlinedAt)
		fields = append(fields, field)
	}
	if md.IsImplicitCode {
		field := fmt.Sprintf("isImplicitCode: %t", md.IsImplicitCode)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DILocation(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DIMacro ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DIMacro is a specialized metadata node.
type DIMacro struct {
	Type  enum.DwarfMacinfo // required.
	Line  int64             // optional; zero value if not present.
	Name  string            // required.
	Value string            // optional; empty if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DIMacro) String() string {
	// '!DIMacro' '(' Fields=(DIMacroField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("type: %s", md.Type)
	fields = append(fields, field)
	if md.Line != 0 {
		field := fmt.Sprintf("line: %d", md.Line)
		fields = append(fields, field)
	}
	field = fmt.Sprintf("name: %s", quote(md.Name))
	fields = append(fields, field)
	if len(md.Value) > 0 {
		field := fmt.Sprintf("value: %s", quote(md.Value))
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DIMacro(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DIMacroFile ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DIMacroFile is a specialized metadata node.
type DIMacroFile struct {
	Type  enum.DwarfMacinfo // optional; zero value if not present.
	Line  int64             // optional; zero value if not present.
	File  MDField           // required.
	Nodes MDField           // optional; nil if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DIMacroFile) String() string {
	// '!DIMacroFile' '(' Fields=(DIMacroFileField separator ',')* ')'
	var fields []string
	if md.Type != 0 {
		field := fmt.Sprintf("type: %s", md.Type)
		fields = append(fields, field)
	}
	if md.Line != 0 {
		field := fmt.Sprintf("line: %d", md.Line)
		fields = append(fields, field)
	}
	field := fmt.Sprintf("file: %s", md.File)
	fields = append(fields, field)
	if md.Nodes != nil {
		field := fmt.Sprintf("nodes: %s", md.Nodes)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DIMacroFile(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DIModule ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DIModule is a specialized metadata node.
type DIModule struct {
	Scope        MDField // required.
	Name         string  // required.
	ConfigMacros string  // optional; empty if not present.
	IncludePath  string  // optional; empty if not present.
	Isysroot     string  // optional; empty if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DIModule) String() string {
	// '!DIModule' '(' Fields=(DIModuleField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("scope: %s", md.Scope)
	fields = append(fields, field)
	field = fmt.Sprintf("name: %s", quote(md.Name))
	fields = append(fields, field)
	if len(md.ConfigMacros) > 0 {
		field := fmt.Sprintf("configMacros: %s", quote(md.ConfigMacros))
		fields = append(fields, field)
	}
	if len(md.IncludePath) > 0 {
		field := fmt.Sprintf("includePath: %s", quote(md.IncludePath))
		fields = append(fields, field)
	}
	if len(md.Isysroot) > 0 {
		field := fmt.Sprintf("isysroot: %s", quote(md.Isysroot))
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DIModule(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DINamespace ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DINamespace is a specialized metadata node.
type DINamespace struct {
	Scope         MDField // required.
	Name          string  // optional; empty if not present.
	ExportSymbols bool    // optional; zero value if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DINamespace) String() string {
	// '!DINamespace' '(' Fields=(DINamespaceField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("scope: %s", md.Scope)
	fields = append(fields, field)
	if len(md.Name) > 0 {
		field := fmt.Sprintf("name: %s", quote(md.Name))
		fields = append(fields, field)
	}
	if md.ExportSymbols {
		field := fmt.Sprintf("exportSymbols: %t", md.ExportSymbols)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DINamespace(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DIObjCProperty ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DIObjCProperty is a specialized metadata node.
type DIObjCProperty struct {
	Name       string  // optional; empty if not present.
	File       MDField // optional; nil if not present.
	Line       int64   // optional; zero value if not present.
	Setter     string  // optional; empty if not present.
	Getter     string  // optional; empty if not present.
	Attributes int64   // optional; zero value if not present.
	Type       MDField // optional; nil if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DIObjCProperty) String() string {
	// '!DIObjCProperty' '(' Fields=(DIObjCPropertyField separator ',')* ')'
	var fields []string
	if len(md.Name) > 0 {
		field := fmt.Sprintf("name: %s", quote(md.Name))
		fields = append(fields, field)
	}
	if md.File != nil {
		field := fmt.Sprintf("file: %s", md.File)
		fields = append(fields, field)
	}
	if md.Line != 0 {
		field := fmt.Sprintf("line: %d", md.Line)
		fields = append(fields, field)
	}
	if len(md.Setter) > 0 {
		field := fmt.Sprintf("setter: %s", quote(md.Setter))
		fields = append(fields, field)
	}
	if len(md.Getter) > 0 {
		field := fmt.Sprintf("getter: %s", quote(md.Getter))
		fields = append(fields, field)
	}
	if md.Attributes != 0 {
		field := fmt.Sprintf("attributes: %d", md.Attributes)
		fields = append(fields, field)
	}
	if md.Type != nil {
		field := fmt.Sprintf("type: %s", md.Type)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DIObjCProperty(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DISubprogram ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DISubprogram is a specialized metadata node.
type DISubprogram struct {
	Scope          MDField              // optional; nil if not present.
	Name           string               // optional; empty if not present.
	LinkageName    string               // optional; empty if not present.
	File           MDField              // optional; nil if not present.
	Line           int64                // optional; zero value if not present.
	Type           MDField              // optional; nil if not present.
	IsLocal        bool                 // optional; zero value if not present.
	IsDefinition   bool                 // optional; zero value if not present.
	ScopeLine      int64                // optional; zero value if not present.
	ContainingType MDField              // optional; nil if not present.
	Virtuality     enum.DwarfVirtuality // optional; zero value if not present.
	VirtualIndex   int64                // optional; zero value if not present.
	ThisAdjustment int64                // optional; zero value if not present.
	Flags          enum.DIFlag          // optional.
	IsOptimized    bool                 // optional; zero value if not present.
	Unit           MDField              // optional; nil if not present.
	TemplateParams MDField              // optional; nil if not present.
	Declaration    MDField              // optional; nil if not present.
	RetainedNodes  MDField              // optional; nil if not present.
	ThrownTypes    MDField              // optional; nil if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DISubprogram) String() string {
	// '!DISubprogram' '(' Fields=(DISubprogramField separator ',')* ')'
	var fields []string
	// Note, to match Clang output, the output order is changed to output name
	// before scope.
	if len(md.Name) > 0 {
		field := fmt.Sprintf("name: %s", quote(md.Name))
		fields = append(fields, field)
	}
	// Note, to match Clang output, the output order is changed to output
	// linkageName before scope.
	if len(md.LinkageName) > 0 {
		field := fmt.Sprintf("linkageName: %s", quote(md.LinkageName))
		fields = append(fields, field)
	}
	if md.Scope != nil {
		field := fmt.Sprintf("scope: %s", md.Scope)
		fields = append(fields, field)
	}
	if md.File != nil {
		field := fmt.Sprintf("file: %s", md.File)
		fields = append(fields, field)
	}
	if md.Line != 0 {
		field := fmt.Sprintf("line: %d", md.Line)
		fields = append(fields, field)
	}
	if md.Type != nil {
		field := fmt.Sprintf("type: %s", md.Type)
		fields = append(fields, field)
	}
	if md.IsLocal {
		field := fmt.Sprintf("isLocal: %t", md.IsLocal)
		fields = append(fields, field)
	}
	if md.IsDefinition {
		field := fmt.Sprintf("isDefinition: %t", md.IsDefinition)
		fields = append(fields, field)
	}
	if md.ScopeLine != 0 {
		field := fmt.Sprintf("scopeLine: %d", md.ScopeLine)
		fields = append(fields, field)
	}
	if md.ContainingType != nil {
		field := fmt.Sprintf("containingType: %s", md.ContainingType)
		fields = append(fields, field)
	}
	if md.Virtuality != 0 {
		field := fmt.Sprintf("virtuality: %s", md.Virtuality)
		fields = append(fields, field)
	}
	if md.VirtualIndex != 0 {
		field := fmt.Sprintf("virtualIndex: %d", md.VirtualIndex)
		fields = append(fields, field)
	}
	if md.ThisAdjustment != 0 {
		field := fmt.Sprintf("thisAdjustment: %d", md.ThisAdjustment)
		fields = append(fields, field)
	}
	if md.Flags != 0 {
		field := fmt.Sprintf("flags: %s", diFlagsString(md.Flags))
		fields = append(fields, field)
	}
	if md.IsOptimized {
		field := fmt.Sprintf("isOptimized: %t", md.IsOptimized)
		fields = append(fields, field)
	}
	if md.Unit != nil {
		field := fmt.Sprintf("unit: %s", md.Unit)
		fields = append(fields, field)
	}
	if md.TemplateParams != nil {
		field := fmt.Sprintf("templateParams: %s", md.TemplateParams)
		fields = append(fields, field)
	}
	if md.Declaration != nil {
		field := fmt.Sprintf("declaration: %s", md.Declaration)
		fields = append(fields, field)
	}
	if md.RetainedNodes != nil {
		field := fmt.Sprintf("retainedNodes: %s", md.RetainedNodes)
		fields = append(fields, field)
	}
	if md.ThrownTypes != nil {
		field := fmt.Sprintf("thrownTypes: %s", md.ThrownTypes)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DISubprogram(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DISubrange ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DISubrange is a specialized metadata node.
type DISubrange struct {
	Count      MDFieldOrInt // required.
	LowerBound int64        // optional; zero value if not present.
}

// String returns a string representation of the specialized metadata node.
func (md *DISubrange) String() string {
	// '!DISubrange' '(' Fields=(DISubrangeField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("count: %s", md.Count)
	fields = append(fields, field)
	if md.LowerBound != 0 {
		field := fmt.Sprintf("lowerBound: %d", md.LowerBound)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!DISubrange(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DISubroutineType ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DISubroutineType is a specialized metadata node.
type DISubroutineType struct {
	Flags enum.DIFlag  // optional.
	CC    enum.DwarfCC // optional; zero value if not present.
	Types MDField      // required.
}

// String returns a string representation of the specialized metadata node.
func (md *DISubroutineType) String() string {
	// '!DISubroutineType' '(' Fields=(DISubroutineTypeField separator ',')* ')'
	var fields []string
	if md.Flags != 0 {
		field := fmt.Sprintf("flags: %s", diFlagsString(md.Flags))
		fields = append(fields, field)
	}
	if md.CC != 0 {
		field := fmt.Sprintf("cc: %s", md.CC)
		fields = append(fields, field)
	}
	field := fmt.Sprintf("types: %s", md.Types)
	fields = append(fields, field)
	return fmt.Sprintf("!DISubroutineType(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DITemplateTypeParameter ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DITemplateTypeParameter is a specialized metadata node.
type DITemplateTypeParameter struct {
	Name string  // optional; empty if not present.
	Type MDField // required.
}

// String returns a string representation of the specialized metadata node.
func (md *DITemplateTypeParameter) String() string {
	// '!DITemplateTypeParameter' '(' Fields=(DITemplateTypeParameterField separator ',')* ')'
	var fields []string
	if len(md.Name) > 0 {
		field := fmt.Sprintf("name: %s", quote(md.Name))
		fields = append(fields, field)
	}
	field := fmt.Sprintf("type: %s", md.Type)
	fields = append(fields, field)
	return fmt.Sprintf("!DITemplateTypeParameter(%s)", strings.Join(fields, ", "))
}

// ~~~ [ DITemplateValueParameter ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DITemplateValueParameter is a specialized metadata node.
type DITemplateValueParameter struct {
	Tag   enum.DwarfTag // optional; zero value if not present.
	Name  string        // optional; empty if not present.
	Type  MDField       // optional; nil if not present.
	Value MDField       // required.
}

// String returns a string representation of the specialized metadata node.
func (md *DITemplateValueParameter) String() string {
	// '!DITemplateValueParameter' '(' Fields=(DITemplateValueParameterField separator ',')* ')'
	var fields []string
	if md.Tag != 0 {
		field := fmt.Sprintf("tag: %s", md.Tag)
		fields = append(fields, field)
	}
	if len(md.Name) > 0 {
		field := fmt.Sprintf("name: %s", quote(md.Name))
		fields = append(fields, field)
	}
	if md.Type != nil {
		field := fmt.Sprintf("type: %s", md.Type)
		fields = append(fields, field)
	}
	field := fmt.Sprintf("value: %s", md.Value)
	fields = append(fields, field)
	return fmt.Sprintf("!DITemplateValueParameter(%s)", strings.Join(fields, ", "))
}

// ~~~ [ GenericDINode ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// GenericDINode is a specialized GenericDINode metadata node.
type GenericDINode struct {
	Tag      enum.DwarfTag // required
	Header   string        // optional; empty if not present
	Operands []MDField     // optional
}

// String returns a string representation of the specialized metadata node.
func (md *GenericDINode) String() string {
	// '!GenericDINode' '(' Fields=(GenericDINodeField separator ',')* ')'
	var fields []string
	field := fmt.Sprintf("tag: %s", md.Tag)
	fields = append(fields, field)
	if len(md.Header) > 0 {
		field := fmt.Sprintf("header: %s", quote(md.Header))
		fields = append(fields, field)
	}
	if len(md.Operands) > 0 {
		// TODO: figure out what operands output should look like.
		field = fmt.Sprintf("operands: %v", md.Operands)
		fields = append(fields, field)
	}
	return fmt.Sprintf("!GenericDINode(%s)", strings.Join(fields, ", "))
}
