package asm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// === [ Identifiers ] =========================================================

// --- [ Global identifiers ] --------------------------------------------------

// globalIdent returns the identifier (without '@' prefix) of the given global
// identifier.
func globalIdent(old ast.GlobalIdent) ir.GlobalIdent {
	ident := old.Text()
	const prefix = "@"
	if !strings.HasPrefix(ident, prefix) {
		panic(fmt.Errorf("invalid global identifier %q; missing '%s' prefix", ident, prefix))
	}
	ident = ident[len(prefix):]
	if id, err := strconv.ParseInt(ident, 10, 64); err == nil {
		return ir.GlobalIdent{GlobalID: id}
	}
	// Unquote after trying to parse as ID, since @"42" is recognized as named
	// and not unnamed.
	ident = unquote(ident)
	return ir.GlobalIdent{GlobalName: ident}
}

// --- [ Local identifiers ] ---------------------------------------------------

// localIdent returns the identifier (without '%' prefix) of the given local
// identifier.
func localIdent(old ast.LocalIdent) ir.LocalIdent {
	ident := old.Text()
	const prefix = "%"
	if !strings.HasPrefix(ident, prefix) {
		panic(fmt.Errorf("invalid local identifier %q; missing '%s' prefix", ident, prefix))
	}
	ident = ident[len(prefix):]
	if id, err := strconv.ParseInt(ident, 10, 64); err == nil {
		return ir.LocalIdent{LocalID: id}
	}
	// Unquote after trying to parse as ID, since %"42" is recognized as named
	// and not unnamed.
	ident = unquote(ident)
	return ir.LocalIdent{LocalName: ident}
}

// --- [ Label identifiers ] ---------------------------------------------------

// labelIdent returns the identifier (without ':' suffix) of the given label
// identifier.
func labelIdent(old ast.LabelIdent) ir.LocalIdent {
	ident := old.Text()
	const suffix = ":"
	if !strings.HasSuffix(ident, suffix) {
		panic(fmt.Errorf("invalid label identifier %q; missing '%s' suffix", ident, suffix))
	}
	ident = ident[:len(ident)-len(suffix)]
	if id, err := strconv.ParseInt(ident, 10, 64); err == nil {
		return ir.LocalIdent{LocalID: id}
	}
	// Unquote after trying to parse as ID, since %"42" is recognized as named
	// and not unnamed.
	ident = unquote(ident)
	return ir.LocalIdent{LocalName: ident}
}

// --- [ Attribute group identifiers ] -----------------------------------------

// attrGroupID returns the ID (without '#' prefix) of the given attribute group
// ID.
func attrGroupID(old ast.AttrGroupID) int64 {
	text := old.Text()
	const prefix = "#"
	if !strings.HasPrefix(text, prefix) {
		panic(fmt.Errorf("invalid attribute group ID %q; missing '%s' prefix", text, prefix))
	}
	text = text[len(prefix):]
	id, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		panic(fmt.Errorf("unable to parse attribute group ID %q; %v", text, err))
	}
	return id
}

// --- [ Comdat identifiers ] --------------------------------------------------

// comdatName returns the name (without '%' prefix) of the given comdat name.
func comdatName(old ast.ComdatName) string {
	name := old.Text()
	const prefix = "$"
	if !strings.HasPrefix(name, prefix) {
		panic(fmt.Errorf("invalid comdat name %q; missing '%s' prefix", name, prefix))
	}
	name = name[len(prefix):]
	return unquote(name)
}

// --- [ Metadata identifiers ] ------------------------------------------------

// metadataName returns the name (without '!' prefix) of the given metadata
// name.
func metadataName(old ast.MetadataName) string {
	name := old.Text()
	const prefix = "!"
	if !strings.HasPrefix(name, prefix) {
		panic(fmt.Errorf("invalid metadata name %q; missing '%s' prefix", name, prefix))
	}
	name = name[len(prefix):]
	return string(enc.Unescape(name))
}

// metadataID returns the ID (without '!' prefix) of the given metadata ID.
func metadataID(old ast.MetadataID) int64 {
	text := old.Text()
	const prefix = "!"
	if !strings.HasPrefix(text, prefix) {
		panic(fmt.Errorf("invalid metadata ID %q; missing '%s' prefix", text, prefix))
	}
	text = text[len(prefix):]
	id, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		panic(fmt.Errorf("unable to parse metadata ID %q; %v", text, err))
	}
	return id
}

// === [ Literals ] ============================================================

// --- [ Integer literals ] ----------------------------------------------------

// boolLit returns the boolean value corresponding to the given boolean literal.
func boolLit(old ast.BoolLit) bool {
	text := old.Text()
	switch text {
	case "true":
		return true
	case "false":
		return false
	default:
		panic(fmt.Errorf(`invalid boolean literal; expected "true" or "false", got %q`, text))
	}
}

// uintLit returns the unsigned integer value corresponding to the given
// unsigned integer literal.
func uintLit(old ast.UintLit) uint64 {
	text := old.Text()
	x, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		panic(fmt.Errorf("unable to parse unsigned integer literal %q; %v", text, err))
	}
	return x
}

// uintSlice returns the slice of unsigned integer value corresponding to the given
// unsigned integer slice.
func uintSlice(olds []ast.UintLit) []uint64 {
	xs := make([]uint64, len(olds))
	for i, old := range olds {
		x := uintLit(old)
		xs[i] = x
	}
	return xs
}

// intLit returns the integer value corresponding to the given integer literal.
func intLit(old ast.IntLit) int64 {
	text := old.Text()
	x, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		panic(fmt.Errorf("unable to parse integer literal %q; %v", text, err))
	}
	return x
}

// --- [ String literals ] -----------------------------------------------------

// stringLit returns the string corresponding to the given string literal.
func stringLit(old ast.StringLit) string {
	return unquote(old.Text())
}

// ___ [ Helpers ] _____________________________________________________________

// irAddrSpace returns the IR address space corresponding to the given AST
// address space.
func irAddrSpace(old ast.AddrSpace) types.AddrSpace {
	return types.AddrSpace(uintLit(old.N()))
}

// irAlign returns the IR alignment corresponding to the given AST alignment.
func irAlign(old ast.Align) ir.Align {
	return ir.Align(uintLit(old.N()))
}

// irArg returns the IR argument corresponding to the given AST argument.
func (fgen *funcGen) irArg(old ast.Arg) (value.Value, error) {
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	switch oldVal := old.Val().(type) {
	case ast.Value:
		x, err := fgen.irValue(typ, oldVal)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		// (optional) Parameter attributes.
		if oldAttrs := old.Attrs(); len(oldAttrs) > 0 {
			attrs := make([]ir.ParamAttribute, len(oldAttrs))
			for i, oldAttr := range old.Attrs() {
				attr, err := fgen.gen.irParamAttribute(oldAttr)
				if err != nil {
					return nil, errors.WithStack(err)
				}
				attrs[i] = attr
			}
			return &ir.Arg{Attrs: attrs, Value: x}, nil
		}
		return x, nil
	case ast.Metadata:
		md, err := fgen.irMetadata(oldVal)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &metadata.Value{Value: md}, nil
	default:
		panic(fmt.Errorf("support for value %T not yet implemented", oldVal))
	}
}

// irBlock returns the IR basic block corresponding to the given AST label.
func (fgen *funcGen) irBlock(old ast.Label) (*ir.Block, error) {
	ident := localIdent(old.Name())
	v, ok := fgen.locals[ident]
	if !ok {
		return nil, errors.Errorf("unable to locate local identifier %q", ident.Ident())
	}
	block, ok := v.(*ir.Block)
	if !ok {
		return nil, errors.Errorf("invalid basic block type; expected *ir.Block, got %T", v)
	}
	return block, nil
}

// irCallingConv returns the IR calling convention corresponding to the given
// AST calling convention.
func irCallingConv(old ast.CallingConv) enum.CallingConv {
	switch old := old.(type) {
	case *ast.CallingConvEnum:
		return asmenum.CallingConvFromString(old.Text())
	case *ast.CallingConvInt:
		cc := uintLit(old.UintLit())
		switch cc {
		case 0:
			// Note, C calling convention is defined as 0 in LLVM. To have the zero-value
			// calling convention mean no calling convention, re-define C calling
			// convention as 1, and use 0 for none.
			return enum.CallingConvC
		default:
			return enum.CallingConv(cc)
		}
	default:
		panic(fmt.Errorf("support for calling convention type %T not yet implemented", old))
	}
}

// irCase returns the IR switch case corresponding to the given AST switch case.
func (fgen *funcGen) irCase(n ast.Case) (*ir.Case, error) {
	// Case comparand.
	x, err := fgen.gen.irTypeConst(n.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Case target branch.
	target, err := fgen.irBlock(n.Target())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewCase(x, target), nil
}

// irClause returns the IR clause corresponding to the given AST clause.
func (fgen *funcGen) irClause(n ast.Clause) (*ir.Clause, error) {
	x, err := fgen.irTypeValue(n.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	clauseType := asmenum.ClauseTypeFromString(n.ClauseType().Text())
	return ir.NewClause(clauseType, x), nil
}

// irExceptionArg returns the IR exception argument corresponding to the given
// AST exception argument.
func (fgen *funcGen) irExceptionArg(old ast.ExceptionArg) (value.Value, error) {
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	switch val := old.Val().(type) {
	case ast.Value:
		return fgen.irValue(typ, val)
	case ast.Metadata:
		md, err := fgen.irMetadata(val)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &metadata.Value{Value: md}, nil
	default:
		panic(fmt.Errorf("spport for exception argument value %T not yet implemented", val))
	}
}

// irExceptionScope returns the IR exception scope corresponding to the given
// AST exception scope.
func (fgen *funcGen) irExceptionScope(old ast.ExceptionScope) (ir.ExceptionScope, error) {
	switch old := old.(type) {
	case *ast.NoneConst:
		return constant.None, nil
	case *ast.LocalIdent:
		ident := localIdent(*old)
		v, ok := fgen.locals[ident]
		if !ok {
			return nil, errors.Errorf("unable to locate local identifier %q", ident.Ident())
		}
		return v, nil
	default:
		panic(fmt.Errorf("spport for exception scope %T not yet implemented", old))
	}
}

// irFastMathFlags returns the IR fast math flags corresponding to the given AST
// fast math flags.
func irFastMathFlags(olds []ast.FastMathFlag) []enum.FastMathFlag {
	if len(olds) == 0 {
		return nil
	}
	flags := make([]enum.FastMathFlag, len(olds))
	for i, old := range olds {
		flag := asmenum.FastMathFlagFromString(old.Text())
		flags[i] = flag
	}
	return flags
}

// irFuncAttribute returns the IR function attribute corresponding to the given
// AST function attribute.
func (gen *generator) irFuncAttribute(old ast.FuncAttribute) ir.FuncAttribute {
	switch old := old.(type) {
	case *ast.AttrString:
		return ir.AttrString(unquote(old.Text()))
	case *ast.AttrPair:
		return ir.AttrPair{
			Key:   unquote(old.Key().Text()),
			Value: unquote(old.Val().Text()),
		}
	case *ast.AttrGroupID:
		id := attrGroupID(*old)
		def, ok := gen.new.attrGroupDefs[id]
		if !ok {
			// Attribute group definition for ID not found.
			//
			// The input file should have contained this definition, but seeing as
			// the LLVM test suite contains several LLVM IR files which omit the
			// attribute group definitions, we will play nice and add an empty
			// definition instead of panicking.
			//
			// This issue is tracked at: https://github.com/llir/llvm/issues/37
			def = &ir.AttrGroupDef{ID: id}
			gen.new.attrGroupDefs[id] = def
		}
		return def
	// TODO: add support for Align.
	//case *ast.Align:
	//	return ir.Align(uintLit(old.N()))
	case *ast.AlignPair:
		return ir.Align(uintLit(old.N()))
	case *ast.AlignStack:
		return ir.AlignStack(uintLit(old.N()))
	case *ast.AlignStackPair:
		return ir.AlignStack(uintLit(old.N()))
	case *ast.AllocSize:
		elemSizeIndex := int(uintLit(old.ElemSizeIndex()))
		if nElemsIndex, ok := old.NElemsIndex(); ok {
			return ir.AllocSize{
				ElemSizeIndex: elemSizeIndex,
				NElemsIndex:   int(uintLit(nElemsIndex)),
			}
		}
		return ir.AllocSize{
			ElemSizeIndex: elemSizeIndex,
			NElemsIndex:   -1,
		}
	case *ast.FuncAttr:
		return asmenum.FuncAttrFromString(old.Text())
	default:
		panic(fmt.Errorf("support for function attribute %T not yet implemented", old))
	}
}

// irImmutable returns the immutable boolean (constant or global) corresponding
// to the given AST immutable.
func irImmutable(old ast.Immutable) bool {
	text := old.Text()
	switch text {
	case "constant":
		return true
	case "global":
		return false
	default:
		panic(fmt.Errorf("support for immutable %q not yet implemented", text))
	}
}

// irIncoming returns the incoming value corresponding to the given AST incoming
// value.
func (fgen *funcGen) irIncoming(xType types.Type, oldX ast.Value, oldPred ast.LocalIdent) (*ir.Incoming, error) {
	x, err := fgen.irValue(xType, oldX)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	predIdent := localIdent(oldPred)
	v, ok := fgen.locals[predIdent]
	if !ok {
		return nil, errors.Errorf("unable to locate local identifier %q", predIdent.Ident())
	}
	pred, ok := v.(*ir.Block)
	if !ok {
		return nil, errors.Errorf("invalid basic block type; expected *ir.Block, got %T", v)
	}
	return ir.NewIncoming(x, pred), nil
}

// irIndirectSymbol returns the IR indirect symbol corresponding to the given
// AST indirect symbol.
func (gen *generator) irIndirectSymbol(typ *types.PointerType, old ast.IndirectSymbol) (constant.Constant, error) {
	switch old := old.(type) {
	case *ast.TypeConst:
		symbol, err := gen.irTypeConst(*old)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return symbol, nil
	case *ast.BitCastExpr:
		symbol, err := gen.irConstant(typ, old)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return symbol, nil
	case *ast.GetElementPtrExpr:
		symbol, err := gen.irConstant(typ, old)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return symbol, nil
	case *ast.AddrSpaceCastExpr:
		symbol, err := gen.irConstant(typ, old)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return symbol, nil
	case *ast.IntToPtrExpr:
		symbol, err := gen.irConstant(typ, old)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return symbol, nil
	default:
		panic(fmt.Errorf("support for indirect symbol %T not yet implemented", old))
	}
}

// irInlineAsm translates the AST inline assembler expression into an equivalent
// IR inline assembler expression.
func irInlineAsm(typ types.Type, old *ast.InlineAsm) *ir.InlineAsm {
	// Assembly instructions.
	asm := stringLit(old.Asm())
	// Constraints.
	constraint := stringLit(old.Constraints())
	v := ir.NewInlineAsm(typ, asm, constraint)
	// (optional) Side effect.
	_, v.SideEffect = old.SideEffect()
	// (optional) Stack alignment.
	_, v.AlignStack = old.AlignStackTok()
	// (optional) Intel dialect.
	_, v.IntelDialect = old.IntelDialect()
	return v
}

// irOperandBundle returns the IR operand bundle corresponding to the given AST
// operand bundle.
func (fgen *funcGen) irOperandBundle(old ast.OperandBundle) (*ir.OperandBundle, error) {
	// Tag.
	tag := stringLit(old.Tag())
	// Inputs.
	var inputs []value.Value
	for _, oldInput := range old.Inputs() {
		input, err := fgen.irTypeValue(oldInput)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		inputs = append(inputs, input)
	}
	return ir.NewOperandBundle(tag, inputs...), nil
}

// irOverflowFlags returns the IR overflow flags corresponding to the given AST
// overflow flags.
func irOverflowFlags(olds []ast.OverflowFlag) []enum.OverflowFlag {
	if len(olds) == 0 {
		return nil
	}
	flags := make([]enum.OverflowFlag, len(olds))
	for i, old := range olds {
		flag := asmenum.OverflowFlagFromString(old.Text())
		flags[i] = flag
	}
	return flags
}

// irParamAttribute returns the IR parameter attribute corresponding to the given
// AST parameter attribute.
func (gen *generator) irParamAttribute(old ast.ParamAttribute) (ir.ParamAttribute, error) {
	switch old := old.(type) {
	case *ast.AttrString:
		return ir.AttrString(unquote(old.Text())), nil
	case *ast.AttrPair:
		return ir.AttrPair{
			Key:   unquote(old.Key().Text()),
			Value: unquote(old.Val().Text()),
		}, nil
	case *ast.Align:
		return ir.Align(uintLit(old.N())), nil
	case *ast.Byval:
		if t, ok := old.Typ(); ok {
			typ, err := gen.irType(t)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			return ir.Byval{Typ: typ}, nil
		}
		return ir.Byval{}, nil
	case *ast.Dereferenceable:
		return ir.Dereferenceable{N: uintLit(old.N())}, nil
	case *ast.DereferenceableOrNull:
		return ir.Dereferenceable{
			N:           uintLit(old.N()),
			DerefOrNull: true,
		}, nil
	case *ast.ParamAttr:
		return asmenum.ParamAttrFromString(old.Text()), nil
	default:
		panic(fmt.Errorf("support for parameter attribute %T not yet implemented", old))
	}
}

// irReturnAttribute returns the IR return attribute corresponding to the given
// AST return attribute.
func irReturnAttribute(old ast.ReturnAttribute) ir.ReturnAttribute {
	switch old := old.(type) {
	// TODO: add support for AttrString.
	//case *ast.AttrString:
	//	return ir.AttrString(unquote(old.Text()))
	// TODO: add support for AttrPair.
	//case *ast.AttrPair:
	//	return ir.AttrPair{
	//		Key:   unquote(old.Key().Text()),
	//		Value: unquote(old.Val().Text()),
	//	}
	case *ast.Align:
		return ir.Align(uintLit(old.N()))
	case *ast.Dereferenceable:
		return ir.Dereferenceable{N: uintLit(old.N())}
	case *ast.DereferenceableOrNull:
		return ir.Dereferenceable{
			N:           uintLit(old.N()),
			DerefOrNull: true,
		}
	case *ast.ReturnAttr:
		return asmenum.ReturnAttrFromString(old.Text())
	default:
		panic(fmt.Errorf("support for return attribute %T not yet implemented", old))
	}
}

// irTLSModelFromThreadLocal returns the IR TLS model corresponding to the given
// AST thread local storage.
func irTLSModelFromThreadLocal(old ast.ThreadLocal) enum.TLSModel {
	if n, ok := old.Model(); ok {
		// e.g. thread_local(initialexec)
		return asmenum.TLSModelFromString(n.Text())
	}
	// If no explicit model is given, the "general dynamic" model is used.
	//    thread_local
	return enum.TLSModelGeneric
}

// irUnwindTarget returns the IR unwind target corresponding to the given AST
// unwind target.
func (fgen *funcGen) irUnwindTarget(n ast.UnwindTarget) (ir.UnwindTarget, error) {
	switch n := n.(type) {
	case *ast.Label:
		return fgen.irBlock(*n)
	case *ast.UnwindToCaller:
		return ir.UnwindToCaller{}, nil
	default:
		panic(fmt.Errorf("support for unwind target %T not yet implemented", n))
	}
}

// irUseListOrder returns the IR use-list order corresponding to the given AST
// use-list order.
func (fgen *funcGen) irUseListOrder(old ast.UseListOrder) (*ir.UseListOrder, error) {
	// Value.
	val, err := fgen.irTypeValue(old.Val())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Indices.
	indices := uintSlice(old.Indices())
	useListOrder := &ir.UseListOrder{
		Value:   val,
		Indices: indices,
	}
	return useListOrder, nil
}

// ### [ Helpers ] #############################################################

// unquote returns the unquoted version of s if quoted, and the original string
// otherwise.
func unquote(s string) string {
	if len(s) >= 2 && strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		return string(enc.Unquote(s))
	}
	return s
}

// text returns the text of the given node.
func text(n ast.LlvmNode) string {
	if n := n.LlvmNode(); n.IsValid() {
		return n.Text()
	}
	return ""
}

// findBlock returns the basic block with the given local identifier in the
// function.
func findBlock(f *ir.Func, blockIdent ir.LocalIdent) (*ir.Block, error) {
	for _, block := range f.Blocks {
		if block.LocalIdent == blockIdent {
			return block, nil
		}
	}
	return nil, errors.Errorf("unable to locate basic block %q of function %q", blockIdent.Ident(), f.Ident())
}
