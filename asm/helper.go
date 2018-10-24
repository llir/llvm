package asm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/l/ir/enum"
	"github.com/llir/l/ir/types"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/mewmew/l-tm/internal/enc"
)

// === [ Identifiers ] =========================================================

// --- [ Global Identifiers ] --------------------------------------------------

// global returns the name (without '@' prefix) of the given global identifier.
func global(n ast.GlobalIdent) string {
	text := n.Text()
	if text == "" {
		// \empty is used when global identifier not present.
		return ""
	}
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
	if text == "" {
		// \empty is used when local identifier not present.
		return ""
	}
	const prefix = "%"
	if !strings.HasPrefix(text, prefix) {
		// NOTE: Panic instead of returning error as this case should not be
		// possible given the grammar.
		panic(fmt.Errorf("invalid local identifier %q; missing '%s' prefix", text, prefix))
	}
	text = text[len(prefix):]
	return unquote(text)
}

// --- [ Label Identifiers ] ---------------------------------------------------

// label returns the name (without ':' suffix) of the given label identifier.
func label(n ast.LabelIdent) string {
	text := n.Text()
	if text == "" {
		// \empty is used when label identifier not present.
		return ""
	}
	const suffix = ":"
	if !strings.HasSuffix(text, suffix) {
		// NOTE: Panic instead of returning error as this case should not be
		// possible given the grammar.
		panic(fmt.Errorf("invalid label identifier %q; missing '%s' suffix", text, suffix))
	}
	text = text[:len(text)-len(suffix)]
	return unquote(text)
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

// irAddrSpace returns the IR address space corresponding to the given optional
// AST address space.
func irAddrSpace(n *ast.AddrSpace) types.AddrSpace {
	// \empty is used when address space not present.
	if n.Text() == "" {
		return 0
	}
	x := uintLit(n.N())
	return types.AddrSpace(x)
}

// irDLLStorageClass returns the IR DLL storage class corresponding to the given
// optional AST DLL storage class.
func irDLLStorageClass(n *ast.DLLStorageClass) enum.DLLStorageClass {
	text := n.Text()
	switch text {
	case "":
		// \empty is used when DLL storage class not present.
		return enum.DLLStorageClassNone
	case "dllexport":
		return enum.DLLStorageClassDLLExport
	case "dllimport":
		return enum.DLLStorageClassDLLImport
	default:
		panic(fmt.Errorf("support for DLL storage class %q not yet implemented", text))
	}
}

// irExternallyInitialized returns the externally initialized boolean
// corresponding to the given optional AST externally initialized.
func irExternallyInitialized(n *ast.ExternallyInitialized) bool {
	// TODO: check why ExternallyInitialized is non-nil, when reduced as \empty.
	return n.Text() == "externally_initialized"
}

// irFastMathFlags returns the IR fast math flags corresponding to the given
// optional AST fast math flags.
func irFastMathFlags(ns []ast.FastMathFlag) []enum.FastMathFlag {
	var flags []enum.FastMathFlag
	for _, n := range ns {
		flag := irFastMathFlag(n)
		flags = append(flags, flag)
	}
	return flags
}

// irFastMathFlag returns the IR fast math flag corresponding to the given
// optional AST fast math flag.
func irFastMathFlag(n ast.FastMathFlag) enum.FastMathFlag {
	text := n.Text()
	switch text {
	case "afn":
		return enum.FastMathFlagAFn
	case "arcp":
		return enum.FastMathFlagARcp
	case "contract":
		return enum.FastMathFlagContract
	case "fast":
		return enum.FastMathFlagFast
	case "ninf":
		return enum.FastMathFlagNInf
	case "nnan":
		return enum.FastMathFlagNNaN
	case "nsz":
		return enum.FastMathFlagNSZ
	case "reassoc":
		return enum.FastMathFlagReassoc
	default:
		panic(fmt.Errorf("support for fast math flag %q not yet implemented", text))
	}
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

// irInBounds returns the in-bounds boolean corresponding to the given optional
// AST in-bounds.
func irInBounds(n *ast.InBounds) bool {
	// TODO: check why InBounds is non-nil, when reduced as \empty.
	return n.Text() == "inbounds"
}

// irInRange returns the in-range boolean corresponding to the given optional
// AST in-range.
func irInRange(n *ast.InRange) bool {
	// TODO: check why InRange is non-nil, when reduced as \empty.
	return n.Text() == "inrange"
}

// irLinkage returns the IR linkage corresponding to the given optional AST
// linkage.
func irLinkage(text string) enum.Linkage {
	// TODO: when ExternLinkage and Linkage are merged in grammar, update
	// irLinkage to take `n *ast.Linkage` instead of `text string`.
	//text := n.Text()
	switch text {
	case "":
		// \empty is used when linkage not present.
		return enum.LinkageNone
	case "appending":
		return enum.LinkageAppending
	case "available_externally":
		return enum.LinkageAvailableExternally
	case "common":
		return enum.LinkageCommon
	case "internal":
		return enum.LinkageInternal
	case "linkonce":
		return enum.LinkageLinkOnce
	case "linkonce_odr":
		return enum.LinkageLinkOnceODR
	case "private":
		return enum.LinkagePrivate
	case "weak":
		return enum.LinkageWeak
	case "weak_odr":
		return enum.LinkageWeakODR
	case "external":
		return enum.LinkageExternal
	case "extern_weak":
		return enum.LinkageExternWeak
	default:
		panic(fmt.Errorf("support for linkage %q not yet implemented", text))
	}
}

// irOverflowFlags returns the IR overflow flags corresponding to the given
// optional AST overflow flags.
func irOverflowFlags(ns []ast.OverflowFlag) []enum.OverflowFlag {
	var flags []enum.OverflowFlag
	for _, n := range ns {
		flag := irOverflowFlag(n)
		flags = append(flags, flag)
	}
	return flags
}

// irOverflowFlag returns the IR overflow flag corresponding to the given
// optional AST overflow flag.
func irOverflowFlag(n ast.OverflowFlag) enum.OverflowFlag {
	text := n.Text()
	switch text {
	case "nsw":
		return enum.OverflowFlagNSW
	case "nuw":
		return enum.OverflowFlagNUW
	default:
		panic(fmt.Errorf("support for overflow flag %q not yet implemented", text))
	}
}

// irPreemption returns the IR preemption corresponding to the given optional
// AST preemption.
func irPreemption(n *ast.Preemption) enum.Preemption {
	text := n.Text()
	switch text {
	case "":
		// \empty is used when preemption not present.
		return enum.PreemptionNone
	case "dso_local":
		return enum.PreemptionDSOLocal
	case "dso_preemptable":
		return enum.PreemptionDSOPreemptable
	default:
		panic(fmt.Errorf("support for preemption %q not yet implemented", text))
	}
}

// irSelectionKind returns the IR Comdat selection kind corresponding to the
// given optional AST Comdat selection kind.
func irSelectionKind(n *ast.SelectionKind) enum.SelectionKind {
	text := n.Text()
	switch text {
	case "any":
		return enum.SelectionKindAny
	case "exactmatch":
		return enum.SelectionKindExactMatch
	case "largest":
		return enum.SelectionKindLargest
	case "noduplicates":
		return enum.SelectionKindNoDuplicates
	case "samesize":
		return enum.SelectionKindSameSize
	default:
		panic(fmt.Errorf("support for Comdat selection kind %q not yet implemented", text))
	}
}

// irTLSModelFromThreadLocal returns the IR TLS model corresponding to the given
// optional AST thread local storage.
func irTLSModelFromThreadLocal(n *ast.ThreadLocal) enum.TLSModel {
	if n.Text() != "" {
		model := irTLSModel(n.Model())
		if model == enum.TLSModelNone {
			// If no explicit model is given, the "general dynamic" model is used.
			//    thread_local
			return enum.TLSModelGeneric
		}
		// e.g. thread_local(initialexec)
		return model
	}
	return enum.TLSModelNone
}

// irTLSModel returns the IR TLS model corresponding to the given optional AST
// TLS model.
func irTLSModel(n *ast.TLSModel) enum.TLSModel {
	text := n.Text()
	switch text {
	case "":
		// \empty is used when TLS model not present.
		return enum.TLSModelNone
	case "initialexec":
		return enum.TLSModelInitialExec
	case "localdynamic":
		return enum.TLSModelLocalDynamic
	case "localexec":
		return enum.TLSModelLocalExec
	default:
		panic(fmt.Errorf("support for TLS model %q not yet implemented", text))
	}
}

// irUnnamedAddr returns the IR unnamed address corresponding to the given
// optional AST unnamed address.
func irUnnamedAddr(n *ast.UnnamedAddr) enum.UnnamedAddr {
	text := n.Text()
	switch text {
	case "":
		// \empty is used when unnamed address not present.
		return enum.UnnamedAddrNone
	case "local_unnamed_addr":
		return enum.UnnamedAddrLocalUnnamedAddr
	case "unnamed_addr":
		return enum.UnnamedAddrUnnamedAddr
	default:
		panic(fmt.Errorf("support for unnamed address %q not yet implemented", text))
	}
}

// irVariadic returns the variadic boolean corresponding to the given optional
// AST ellipsis.
func irVariadic(n *ast.Ellipsis) bool {
	// TODO: check why Variadic is non-nil for `Variadic=Ellipsisopt`, regardless
	// of whether the input is (...) or ().
	//
	// It seems that the Variadic.Text simply returns empty string when
	// Ellipsisopt reduces to \empty.
	//
	// Using `n.Text() == "..."` for now, would like to use `n != nil`.
	return n.Text() == "..."
}

// irVisibility returns the IR visibility kind corresponding to the given
// optional AST visibility kind.
func irVisibility(n *ast.Visibility) enum.Visibility {
	text := n.Text()
	switch text {
	case "":
		// \empty is used when visibility kind not present.
		return enum.VisibilityNone
	case "default":
		return enum.VisibilityDefault
	case "hidden":
		return enum.VisibilityHidden
	case "protected":
		return enum.VisibilityProtected
	default:
		panic(fmt.Errorf("support for visibility kind %q not yet implemented", text))
	}
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
