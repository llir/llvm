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

// --- [ Global Identifiers ] --------------------------------------------------

// globalIdent returns the identifier (without '@' prefix) of the given global
// identifier.
func globalIdent(n ast.GlobalIdent) ir.GlobalIdent {
	ident := n.Text()
	const prefix = "@"
	if !strings.HasPrefix(ident, prefix) {
		panic(fmt.Errorf("invalid global identifier %q; missing '%s' prefix", ident, prefix))
	}
	ident = ident[len(prefix):]
	if id, err := strconv.ParseInt(ident, 10, 64); err == nil {
		return ir.GlobalIdent{GlobalID: id}
	}
	// Unquote after trying to parse as ID, since %"42" is recognized as named
	// and not unnamed.
	ident = unquote(ident)
	return ir.GlobalIdent{GlobalName: ident}
}

// --- [ Local Identifiers ] ---------------------------------------------------

// localIdent returns the identifier (without '%' prefix) of the given local
// identifier.
func localIdent(n ast.LocalIdent) ir.LocalIdent {
	ident := n.Text()
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

// --- [ Label Identifiers ] ---------------------------------------------------

// labelIdent returns the identifier (without ':' suffix) of the given label
// identifier.
func labelIdent(n ast.LabelIdent) ir.LocalIdent {
	ident := n.Text()
	const suffix = ":"
	if !strings.HasSuffix(ident, suffix) {
		panic(fmt.Errorf("invalid label identifier %q; missing '%s' suffix", ident, suffix))
	}
	ident = ident[:len(ident)-len(suffix)]
	// Note, label identifiers are always named if present (i.e. `42:` has the
	// label name 42, not the ID 42).
	ident = unquote(ident)
	return ir.LocalIdent{LocalName: ident}
}

// --- [ Attribute Group Identifiers ] -----------------------------------------

// attrGroupID returns the ID (without '#' prefix) of the given attribute group
// ID.
func attrGroupID(n ast.AttrGroupID) int64 {
	text := n.Text()
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

// --- [ Comdat Identifiers ] --------------------------------------------------

// comdatName returns the name (without '%' prefix) of the given comdat name.
func comdatName(n ast.ComdatName) string {
	name := n.Text()
	const prefix = "$"
	if !strings.HasPrefix(name, prefix) {
		panic(fmt.Errorf("invalid comdat name %q; missing '%s' prefix", name, prefix))
	}
	name = name[len(prefix):]
	return unquote(name)
}

// --- [ Metadata Identifiers ] ------------------------------------------------

// metadataName returns the name (without '!' prefix) of the given metadata
// name.
func metadataName(n ast.MetadataName) string {
	name := n.Text()
	const prefix = "!"
	if !strings.HasPrefix(name, prefix) {
		panic(fmt.Errorf("invalid metadata name %q; missing '%s' prefix", name, prefix))
	}
	name = name[len(prefix):]
	return string(enc.Unescape(name))
}

// metadataID returns the ID (without '!' prefix) of the given metadata ID.
func metadataID(n ast.MetadataID) int64 {
	text := n.Text()
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
func boolLit(n ast.BoolLit) bool {
	text := n.Text()
	switch text {
	case "true":
		return true
	case "false":
		return false
	default:
		panic(fmt.Errorf("invalid boolean literal; expected `true` or `false`, got `%v`", text))
	}
}

// uintLit returns the unsigned integer value corresponding to the given
// unsigned integer literal.
func uintLit(n ast.UintLit) uint64 {
	text := n.Text()
	x, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		panic(fmt.Errorf("unable to parse unsigned integer literal %q; %v", text, err))
	}
	return x
}

// uintSlice returns the slice of unsigned integer value corresponding to the given
// unsigned integer slice.
func uintSlice(ns []ast.UintLit) []uint64 {
	var xs []uint64
	for _, n := range ns {
		x := uintLit(n)
		xs = append(xs, x)
	}
	return xs
}

// intLit returns the integer value corresponding to the given integer literal.
func intLit(n ast.IntLit) int64 {
	text := n.Text()
	x, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		panic(fmt.Errorf("unable to parse integer literal %q; %v", text, err))
	}
	return x
}

// --- [ Floating-point literals ] ---------------------------------------------

// --- [ String literals ] -----------------------------------------------------

// stringLit returns the string corresponding to the given string literal.
func stringLit(n ast.StringLit) string {
	return string(stringLitBytes(n))
}

// stringLitBytes returns the byte slice corresponding to the given string literal.
func stringLitBytes(n ast.StringLit) []byte {
	text := n.Text()
	return enc.Unquote(text)
}

// --- [ Null literals ] -------------------------------------------------------

// ___ [ Helpers ] _____________________________________________________________

// irAddrSpace returns the IR address space corresponding to the given AST
// address space.
func irAddrSpace(n ast.AddrSpace) types.AddrSpace {
	return types.AddrSpace(uintLit(n.N()))
}

// irAlign returns the IR alignment corresponding to the given AST alignment.
func irAlign(n ast.Align) ir.Align {
	return ir.Align(uintLit(n.N()))
}

// irArg translates the given AST argument into an equivalent IR argument.
func (fgen *funcGen) irArg(old ast.Arg) (value.Value, error) {
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	switch oldVal := old.Val().(type) {
	case ast.Value:
		x, err := fgen.astToIRValue(typ, oldVal)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if len(old.Attrs()) > 0 {
			var attrs []ir.ParamAttribute
			for _, oldAttr := range old.Attrs() {
				attr := irParamAttribute(oldAttr)
				attrs = append(attrs, attr)
			}
			arg := &ir.Arg{
				Attrs: attrs,
				Value: x,
			}
			return arg, nil
		}
		return x, nil
	case ast.Metadata:
		md, err := fgen.irMetadata(oldVal)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		v := &metadata.Value{
			Value: md,
		}
		return v, nil
	default:
		panic(fmt.Errorf("support for value %T not yet implemented", oldVal))
	}
}

// irBasicBlock returns the IR basic block corresponding to the given AST label.
func (fgen *funcGen) irBasicBlock(old ast.Label) (*ir.BasicBlock, error) {
	ident := localIdent(old.Name())
	v, ok := fgen.ls[ident]
	if !ok {
		return nil, errors.Errorf("unable to locate local identifier %q", ident.Ident())
	}
	block, ok := v.(*ir.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid basic block type; expected *ir.BasicBlock, got %T", v)
	}
	return block, nil
}

// irCallingConv returns the IR calling convention corresponding to the given
// AST calling convention.
func irCallingConv(n ast.CallingConv) enum.CallingConv {
	switch n := n.(type) {
	case *ast.CallingConvEnum:
		return asmenum.CallingConvFromString(n.Text())
	case *ast.CallingConvInt:
		cc := uintLit(n.UintLit())
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
		panic(fmt.Errorf("support for calling convention type %T not yet implemented", n))
	}
}

// irCase returns the IR switch case corresponding to the given AST switch case.
func (fgen *funcGen) irCase(n ast.Case) (*ir.Case, error) {
	x, err := fgen.gen.irTypeConst(n.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	target, err := fgen.irBasicBlock(n.Target())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewCase(x, target), nil
}

// irClause returns the IR clause corresponding to the given AST clause.
func (fgen *funcGen) irClause(n ast.Clause) (*ir.Clause, error) {
	x, err := fgen.astToIRTypeValue(n.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	clauseType := asmenum.ClauseTypeFromString(n.ClauseType().Text())
	return ir.NewClause(clauseType, x), nil
}

// irExceptionArg returns the IR exception argument corresponding to the given
// AST exception argument.
func (fgen *funcGen) irExceptionArg(n ast.ExceptionArg) (value.Value, error) {
	typ, err := fgen.gen.irType(n.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	switch val := n.Val().(type) {
	case ast.Value:
		return fgen.astToIRValue(typ, val)
	case ast.Metadata:
		md, err := fgen.irMetadata(val)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		v := &metadata.Value{
			Value: md,
		}
		return v, nil
	default:
		panic(fmt.Errorf("spport for exception argument value %T not yet implemented", val))
	}
}

// irExceptionScope returns the IR exception scope corresponding to the given
// AST exception scope.
func (fgen *funcGen) irExceptionScope(n ast.ExceptionScope) (ir.ExceptionScope, error) {
	switch n := n.(type) {
	case *ast.NoneConst:
		return constant.None, nil
	case *ast.LocalIdent:
		ident := localIdent(*n)
		v, ok := fgen.ls[ident]
		if !ok {
			return nil, errors.Errorf("unable to locate local identifier %q", ident.Ident())
		}
		return v, nil
	default:
		panic(fmt.Errorf("spport for exception scope %T not yet implemented", n))
	}
}

// irFastMathFlags returns the IR fast math flags corresponding to the given AST
// fast math flags.
func irFastMathFlags(ns []ast.FastMathFlag) []enum.FastMathFlag {
	var flags []enum.FastMathFlag
	for _, n := range ns {
		flag := asmenum.FastMathFlagFromString(n.Text())
		flags = append(flags, flag)
	}
	return flags
}

// irFuncAttribute returns the IR function attribute corresponding to the given
// AST function attribute.
func (gen *generator) irFuncAttribute(n ast.FuncAttribute) ir.FuncAttribute {
	switch n := n.(type) {
	case *ast.AttrString:
		return ir.AttrString(unquote(n.Text()))
	case *ast.AttrPair:
		return ir.AttrPair{
			Key:   unquote(n.Key().Text()),
			Value: unquote(n.Val().Text()),
		}
	case *ast.AttrGroupID:
		id := attrGroupID(*n)
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
			// Note, we resolve the ID to an empty definition, but to skip output
			// the ID is not added to attrGroupDefOrder.
			//gen.old.attrGroupDefOrder = append(gen.old.attrGroupDefOrder, id)
		}
		return def
	// TODO: add support for Align.
	//case *ast.Align:
	//	return ir.Align(uintLit(n.N()))
	case *ast.AlignPair:
		return ir.Align(uintLit(n.N()))
	case *ast.AlignStack:
		return ir.AlignStack(uintLit(n.N()))
	case *ast.AlignStackPair:
		return ir.AlignStack(uintLit(n.N()))
	case *ast.AllocSize:
		// TODO: add support for AllocSize.
		panic("support for function attribute AllocSize not yet implemented")
	case *ast.FuncAttr:
		return asmenum.FuncAttrFromString(n.Text())
	default:
		panic(fmt.Errorf("support for function attribute %T not yet implemented", n))
	}
}

// irImmutable returns the immutable (constant or global) boolean corresponding
// to the given AST immutable.
func irImmutable(n ast.Immutable) bool {
	text := n.Text()
	switch text {
	case "constant":
		return true
	case "global":
		return false
	default:
		panic(fmt.Errorf("support for immutable %q not yet implemented", text))
	}
}

// irIncoming translates the given AST incoming value into an equivalent IR
// incoming value.
func (fgen *funcGen) irIncoming(xType types.Type, oldX ast.Value, oldPred ast.LocalIdent) (*ir.Incoming, error) {
	x, err := fgen.astToIRValue(xType, oldX)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	predIdent := localIdent(oldPred)
	v, ok := fgen.ls[predIdent]
	if !ok {
		return nil, errors.Errorf("unable to locate local identifier %q", predIdent.Ident())
	}
	pred, ok := v.(*ir.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid basic block type; expected *ir.BasicBlock, got %T", v)
	}
	inc := ir.NewIncoming(x, pred)
	return inc, nil
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

// irOverflowFlags returns the IR overflow flags corresponding to the given AST
// overflow flags.
func irOverflowFlags(ns []ast.OverflowFlag) []enum.OverflowFlag {
	var flags []enum.OverflowFlag
	for _, n := range ns {
		flag := asmenum.OverflowFlagFromString(n.Text())
		flags = append(flags, flag)
	}
	return flags
}

// irParamAttribute returns the IR parameter attribute corresponding to the given
// AST parameter attribute.
func irParamAttribute(n ast.ParamAttribute) ir.ParamAttribute {
	switch n := n.(type) {
	case *ast.AttrString:
		return ir.AttrString(unquote(n.Text()))
	case *ast.AttrPair:
		return ir.AttrPair{
			Key:   unquote(n.Key().Text()),
			Value: unquote(n.Val().Text()),
		}
	case *ast.Align:
		return ir.Align(uintLit(n.N()))
	case *ast.Dereferenceable:
		return ir.Dereferenceable{N: uintLit(n.N())}
	case *ast.DereferenceableOrNull:
		return ir.Dereferenceable{
			N:           uintLit(n.N()),
			DerefOrNull: true,
		}
	case *ast.ParamAttr:
		return asmenum.ParamAttrFromString(n.Text())
	default:
		panic(fmt.Errorf("support for parameter attribute %T not yet implemented", n))
	}
}

// irReturnAttribute returns the IR return attribute corresponding to the given
// AST return attribute.
func irReturnAttribute(n ast.ReturnAttribute) ir.ReturnAttribute {
	switch n := n.(type) {
	// TODO: add support for *ast.AttrString and *ast.AttrPair when supported by grammar.
	//case *ast.AttrString:
	//	return ir.AttrString(unquote(n.Text()))
	//case *ast.AttrPair:
	//	return ir.AttrPair{
	//		Key:   unquote(n.Key().Text()),
	//		Value: unquote(n.Val().Text()),
	//	}
	case *ast.Align:
		return ir.Align(uintLit(n.N()))
	case *ast.Dereferenceable:
		return ir.Dereferenceable{N: uintLit(n.N())}
	case *ast.DereferenceableOrNull:
		return ir.Dereferenceable{
			N:           uintLit(n.N()),
			DerefOrNull: true,
		}
	case *ast.ReturnAttr:
		return asmenum.ReturnAttrFromString(n.Text())
	default:
		panic(fmt.Errorf("support for return attribute %T not yet implemented", n))
	}
}

// irOperandBundle returns the IR operand bundle corresponding to the given AST
// operand bundle.
func (fgen *funcGen) irOperandBundle(n ast.OperandBundle) (*ir.OperandBundle, error) {
	// Tag.
	tag := stringLit(n.Tag())
	// Inputs.
	var inputs []value.Value
	for _, oldInput := range n.Inputs() {
		input, err := fgen.astToIRTypeValue(oldInput)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		inputs = append(inputs, input)
	}
	return ir.NewOperandBundle(tag, inputs...), nil
}

// irTLSModelFromThreadLocal returns the IR TLS model corresponding to the given
// AST thread local storage.
func irTLSModelFromThreadLocal(n ast.ThreadLocal) enum.TLSModel {
	if n := n.Model(); n.IsValid() {
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
		return fgen.irBasicBlock(*n)
	case *ast.UnwindToCaller:
		return ir.UnwindToCaller{}, nil
	default:
		panic(fmt.Errorf("support for unwind target %T not yet implemented", n))
	}
}

// irUseListOrder returns the IR use-list order corresponding to the given AST
// use-list order.
func (fgen *funcGen) irUseListOrder(n ast.UseListOrder) (*ir.UseListOrder, error) {
	// Value.
	val, err := fgen.astToIRTypeValue(n.Val())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Indices.
	var indices []int64
	for _, oldIndex := range n.Indices() {
		index := int64(uintLit(oldIndex))
		indices = append(indices, index)
	}
	useListOrder := &ir.UseListOrder{
		Value:   val,
		Indices: indices,
	}
	return useListOrder, nil
}

// irUseListOrder returns the IR use-list order corresponding to the given AST
// use-list order.
func (gen *generator) irUseListOrder(n ast.UseListOrder) (*ir.UseListOrder, error) {
	// Value.
	typ, err := gen.irType(n.Val().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	oldVal := n.Val().Val()
	oldConst, ok := oldVal.(ast.Constant)
	if !ok {
		return nil, errors.Errorf("unable to resolve value %T in module use-list order", oldVal)
	}
	val, err := gen.irConstant(typ, oldConst)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Indices.
	var indices []int64
	for _, oldIndex := range n.Indices() {
		index := int64(uintLit(oldIndex))
		indices = append(indices, index)
	}
	useListOrder := &ir.UseListOrder{
		Value:   val,
		Indices: indices,
	}
	return useListOrder, nil
}

// irUseListOrderBB returns the IR basic block specific use-list order
// corresponding to the given AST basic block specific use-list order.
func (gen *generator) irUseListOrderBB(n ast.UseListOrderBB) (*ir.UseListOrderBB, error) {
	// Function.
	funcIdent := globalIdent(n.Func())
	v, ok := gen.new.globals[funcIdent]
	if !ok {
		return nil, errors.Errorf("unable to locate global identifier %q", funcIdent.Ident())
	}
	f, ok := v.(*ir.Function)
	if !ok {
		return nil, errors.Errorf("invalid function type of %q; expected *ir.Function, got %T", funcIdent.Ident(), v)
	}
	// Basic block.
	blockIdent := localIdent(n.Block())
	var block *ir.BasicBlock
	for _, bb := range f.Blocks {
		if bb.LocalIdent == blockIdent {
			block = bb
			break
		}
	}
	if block == nil {
		return nil, errors.Errorf("unable to locate basic block %q of function %q", blockIdent.Ident(), funcIdent.Ident())
	}
	// Indices.
	var indices []int64
	for _, oldIndex := range n.Indices() {
		index := int64(uintLit(oldIndex))
		indices = append(indices, index)
	}
	useListOrderBB := &ir.UseListOrderBB{
		Func:    f,
		Block:   block,
		Indices: indices,
	}
	return useListOrderBB, nil
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
