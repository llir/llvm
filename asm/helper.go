package asm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/l/ir"
	"github.com/llir/l/ir/enum"
	"github.com/llir/l/ir/types"
	asmenum "github.com/mewmew/l-tm/asm/enum"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/mewmew/l-tm/internal/enc"
	"github.com/pkg/errors"
)

// === [ Identifiers ] =========================================================

// --- [ Global Identifiers ] --------------------------------------------------

// global returns the name (without '@' prefix) of the given global identifier.
func global(n ast.GlobalIdent) string {
	text := n.Text()
	const prefix = "@"
	if !strings.HasPrefix(text, prefix) {
		// NOTE: Panic instead of returning error as this case should not be
		// possible given the grammar.
		panic(fmt.Errorf("invalid global identifier %q; missing '%s' prefix", text, prefix))
	}
	text = text[len(prefix):]
	return unquote(text)
}

// --- [ Local Identifiers ] ---------------------------------------------------

// local returns the name (without '%' prefix) of the given local identifier.
func local(n ast.LocalIdent) string {
	text := n.Text()
	const prefix = "%"
	if !strings.HasPrefix(text, prefix) {
		// NOTE: Panic instead of returning error as this case should not be
		// possible given the grammar.
		panic(fmt.Errorf("invalid local identifier %q; missing '%s' prefix", text, prefix))
	}
	text = text[len(prefix):]
	return unquote(text)
}

// optLocal returns the name (without '%' prefix) of the given optional local
// identifier.
func optLocal(n *ast.LocalIdent) string {
	if n == nil {
		return ""
	}
	return local(*n)
}

// --- [ Label Identifiers ] ---------------------------------------------------

// label returns the name (without ':' suffix) of the given label identifier.
func label(n ast.LabelIdent) string {
	text := n.Text()
	const suffix = ":"
	if !strings.HasSuffix(text, suffix) {
		// NOTE: Panic instead of returning error as this case should not be
		// possible given the grammar.
		panic(fmt.Errorf("invalid label identifier %q; missing '%s' suffix", text, suffix))
	}
	text = text[:len(text)-len(suffix)]
	return unquote(text)
}

// optLabel returns the name (without ':' suffix) of the given optional label
// identifier.
func optLabel(n *ast.LabelIdent) string {
	if n == nil {
		return ""
	}
	return label(*n)
}

// --- [ Attribute Group Identifiers ] -----------------------------------------

// --- [ Comdat Identifiers ] --------------------------------------------------

// --- [ Metadata Identifiers ] ------------------------------------------------

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
		// NOTE: Panic instead of returning error as this case should not be
		// possible given the grammar.
		panic(fmt.Errorf("invalid boolean literal; expected `true` or `false`, got `%v`", text))
	}
}

// uintLit returns the unsigned integer value corresponding to the given
// unsigned integer literal.
func uintLit(n ast.UintLit) uint64 {
	text := n.Text()
	x, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		// NOTE: Panic instead of returning error as this case should not be
		// possible given the grammar.

		// TODO: figure out how to update the grammar for UintLit to remove the
		// optional sign.
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

// --- [ Floating-point literals ] ---------------------------------------------

// --- [ String literals ] -----------------------------------------------------

// stringLit returns the string corresponding to the given string literal.
func stringLit(n ast.StringLit) string {
	text := n.Text()
	s := enc.Unquote(text)
	return string(s)
}

// TODO: remove stringLitBytes if not used.

// stringLitBytes returns the byte slice corresponding to the given string literal.
func stringLitBytes(n ast.StringLit) []byte {
	text := n.Text()
	return enc.Unquote(text)
}

// --- [ Null literals ] -------------------------------------------------------

// ___ [ Helpers ] _____________________________________________________________

// irOptAddrSpace returns the IR address space corresponding to the given
// optional AST address space.
func irOptAddrSpace(n *ast.AddrSpace) types.AddrSpace {
	if n == nil {
		return 0
	}
	x := uintLit(n.N())
	return types.AddrSpace(x)
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

// irOptCallingConv returns the IR calling convention corresponding to the given
// optional AST calling convention.
func irOptCallingConv(n ast.CallingConv) enum.CallingConv {
	if n == nil {
		return enum.CallingConvNone
	}
	switch n := n.(type) {
	case *ast.CallingConvEnum:
		return asmenum.CallingConvFromString(n.Text())
	case *ast.CallingConvInt:
		x := uintLit(n.UintLit())
		switch x {
		case 11:
			return enum.CallingConvHiPE
		case 86:
			return enum.CallingConvAVRBuiltin
		case 87:
			return enum.CallingConvAMDGPUVS
		case 88:
			return enum.CallingConvAMDGPUGS
		case 89:
			return enum.CallingConvAMDGPUPS
		case 90:
			return enum.CallingConvAMDGPUCS
		case 91:
			return enum.CallingConvAMDGPUKernel
		case 93:
			return enum.CallingConvAMDGPUHS
		case 94:
			return enum.CallingConvMSP430Builtin
		case 95:
			return enum.CallingConvAMDGPULS
		case 96:
			return enum.CallingConvAMDGPUES
		default:
			panic(fmt.Errorf("support for calling convention %d not yet implemented", x))
		}
	default:
		panic(fmt.Errorf("support for calling convention type %T not yet implemented", n))
	}
}

// irOptDLLStorageClass returns the IR DLL storage class corresponding to the
// given optional AST DLL storage class.
func irOptDLLStorageClass(n *ast.DLLStorageClass) enum.DLLStorageClass {
	if n == nil {
		return enum.DLLStorageClassNone
	}
	return asmenum.DLLStorageClassFromString(n.Text())
}

// irOptExternallyInitialized returns the externally initialized boolean
// corresponding to the given optional AST externally initialized.
func irOptExternallyInitialized(n *ast.ExternallyInitialized) bool {
	return n != nil
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

// irImmutable returns the immutable (constant or global) boolean corresponding
// to the given optional AST immutable.
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

// irOptInBounds returns the in-bounds boolean corresponding to the given
// optional AST in-bounds.
func irOptInBounds(n *ast.InBounds) bool {
	return n != nil
}

// irOptInRange returns the in-range boolean corresponding to the given optional
// AST in-range.
func irOptInRange(n *ast.InRange) bool {
	return n != nil
}

// irOptLinkage returns the IR linkage corresponding to the given optional AST
// linkage.
func irOptLinkage(n ast.LlvmNode) enum.Linkage {
	// TODO: fix implementation of optlinkage.
	return enum.LinkageNone
	if n == nil {
		return enum.LinkageNone
	}
	return asmenum.LinkageFromString(n.LlvmNode().Text())
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

// irOptPreemption returns the IR preemption corresponding to the given optional
// AST preemption.
func irOptPreemption(n *ast.Preemption) enum.Preemption {
	if n == nil {
		return enum.PreemptionNone
	}
	return asmenum.PreemptionFromString(n.Text())
}

// irOptSelectionKind returns the IR Comdat selection kind corresponding to the
// given optional AST Comdat selection kind.
func irOptSelectionKind(n *ast.SelectionKind) enum.SelectionKind {
	if n == nil {
		return enum.SelectionKindAny
	}
	return asmenum.SelectionKindFromString(n.Text())
}

// irOptTLSModelFromThreadLocal returns the IR TLS model corresponding to the
// given optional AST thread local storage.
func irOptTLSModelFromThreadLocal(n *ast.ThreadLocal) enum.TLSModel {
	if n == nil {
		return enum.TLSModelNone
	}
	model := irOptTLSModel(n.Model())
	if model == enum.TLSModelNone {
		// If no explicit model is given, the "general dynamic" model is used.
		//    thread_local
		return enum.TLSModelGeneric
	}
	// e.g. thread_local(initialexec)
	return model
}

// irOptTLSModel returns the IR TLS model corresponding to the given optional
// AST TLS model.
func irOptTLSModel(n *ast.TLSModel) enum.TLSModel {
	if n == nil {
		return enum.TLSModelNone
	}
	return asmenum.TLSModelFromString(n.Text())
}

// irOptUnnamedAddr returns the IR unnamed address corresponding to the given
// optional AST unnamed address.
func irOptUnnamedAddr(n *ast.UnnamedAddr) enum.UnnamedAddr {
	if n == nil {
		return enum.UnnamedAddrNone
	}
	return asmenum.UnnamedAddrFromString(n.Text())
}

// irOptVariadic returns the variadic boolean corresponding to the given
// optional AST ellipsis.
func irOptVariadic(n *ast.Ellipsis) bool {
	return n != nil
}

// irOptVisibility returns the IR visibility kind corresponding to the given
// optional AST visibility kind.
func irOptVisibility(n *ast.Visibility) enum.Visibility {
	if n == nil {
		return enum.VisibilityNone
	}
	return asmenum.VisibilityFromString(n.Text())
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
