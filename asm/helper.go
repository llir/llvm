package asm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/types"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/mewmew/l-tm/internal/enc"
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

// --- [ Label Identifiers ] ---------------------------------------------------

// --- [ Attribute Group Identifiers ] -----------------------------------------

// --- [ Comdat Identifiers ] --------------------------------------------------

// --- [ Metadata Identifiers ] ------------------------------------------------

// === [ Literals ] ============================================================

// --- [ Integer literals ] ----------------------------------------------------

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

// --- [ Floating-point literals ] ---------------------------------------------

// --- [ String literals ] -----------------------------------------------------

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
func irDLLStorageClass(n *ast.DLLStorageClass) ll.DLLStorageClass {
	text := n.Text()
	switch text {
	case "":
		// \empty is used when DLL storage class not present.
		return ll.DLLStorageClassNone
	case "dllexport":
		return ll.DLLStorageClassDLLExport
	case "dllimport":
		return ll.DLLStorageClassDLLImport
	default:
		panic(fmt.Errorf("support for DLL storage class %q not yet implemented", text))
	}
}

// irLinkage returns the IR linkage corresponding to the given optional AST
// linkage.
func irLinkage(n *ast.Linkage) ll.Linkage {
	text := n.Text()
	switch text {
	case "":
		// \empty is used when linkage not present.
		return ll.LinkageNone
	case "appending":
		return ll.LinkageAppending
	case "available_externally":
		return ll.LinkageAvailableExternally
	case "common":
		return ll.LinkageCommon
	case "internal":
		return ll.LinkageInternal
	case "linkonce":
		return ll.LinkageLinkOnce
	case "linkonce_odr":
		return ll.LinkageLinkOnceODR
	case "private":
		return ll.LinkagePrivate
	case "weak":
		return ll.LinkageWeak
	case "weak_odr":
		return ll.LinkageWeakODR
	case "external":
		return ll.LinkageExternal
	case "extern_weak":
		return ll.LinkageExternWeak
	default:
		panic(fmt.Errorf("support for linkage %q not yet implemented", text))
	}
}

// irPreemption returns the IR preemption corresponding to the given optional
// AST preemption.
func irPreemption(n *ast.Preemption) ll.Preemption {
	text := n.Text()
	switch text {
	case "":
		// \empty is used when preemption not present.
		return ll.PreemptionNone
	case "dso_local":
		return ll.PreemptionDSOLocal
	case "dso_preemptable":
		return ll.PreemptionDSOPreemptable
	default:
		panic(fmt.Errorf("support for preemption %q not yet implemented", text))
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

// ### [ Helpers ] #############################################################

// unquote returns the unquoted version of s if quoted, and the original string
// otherwise.
func unquote(s string) string {
	if len(s) >= 2 && strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		return string(enc.Unquote(s))
	}
	return s
}
