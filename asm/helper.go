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

// irExternallyInitialized returns the externally initialized boolean
// corresponding to the given optional AST externally initialized.
func irExternallyInitialized(n *ast.ExternallyInitialized) bool {
	// TODO: check why ExternallyInitialized is non-nil, when reduced as \empty.
	return n.Text() == "externally_initialized"
}

// irImmutable returns the immutable (constant or global) boolean corresponding
// to the given optional AST immutable.
func irImmutable(n ast.Immutable) bool {
	// TODO: check why Immutable is non-nil, when reduced as \empty.
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

// irLinkage returns the IR linkage corresponding to the given optional AST
// linkage.
func irLinkage(text string) ll.Linkage {
	// TODO: when ExternLinkage and Linkage are merged in grammar, update
	// irLinkage to take `n *ast.Linkage` instead of `text string`.

	//text := n.Text()
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

// irSelectionKind returns the IR Comdat selection kind corresponding to the
// given optional AST Comdat selection kind.
func irSelectionKind(n *ast.SelectionKind) ll.SelectionKind {
	text := n.Text()
	switch text {
	case "any":
		return ll.SelectionKindAny
	case "exactmatch":
		return ll.SelectionKindExactMatch
	case "largest":
		return ll.SelectionKindLargest
	case "noduplicates":
		return ll.SelectionKindNoDuplicates
	case "samesize":
		return ll.SelectionKindSameSize
	default:
		panic(fmt.Errorf("support for Comdat selection kind %q not yet implemented", text))
	}
}

// irTLSModelFromThreadLocal returns the IR TLS model corresponding to the given
// optional AST thread local storage.
func irTLSModelFromThreadLocal(n *ast.ThreadLocal) ll.TLSModel {
	if n.Text() != "" {
		model := irTLSModel(n.Model())
		if model == ll.TLSModelNone {
			// If no explicit model is given, the "general dynamic" model is used.
			//    thread_local
			return ll.TLSModelGeneric
		}
		// e.g. thread_local(initialexec)
		return model
	}
	return ll.TLSModelNone
}

// irTLSModel returns the IR TLS model corresponding to the given optional AST
// TLS model.
func irTLSModel(n *ast.TLSModel) ll.TLSModel {
	text := n.Text()
	switch text {
	case "":
		// \empty is used when TLS model not present.
		return ll.TLSModelNone
	case "initialexec":
		return ll.TLSModelInitialExec
	case "localdynamic":
		return ll.TLSModelLocalDynamic
	case "localexec":
		return ll.TLSModelLocalExec
	default:
		panic(fmt.Errorf("support for TLS model %q not yet implemented", text))
	}
}

// irUnnamedAddr returns the IR unnamed address corresponding to the given
// optional AST unnamed address.
func irUnnamedAddr(n *ast.UnnamedAddr) ll.UnnamedAddr {
	text := n.Text()
	switch text {
	case "":
		// \empty is used when unnamed address not present.
		return ll.UnnamedAddrNone
	case "local_unnamed_addr":
		return ll.UnnamedAddrLocalUnnamedAddr
	case "unnamed_addr":
		return ll.UnnamedAddrUnnamedAddr
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
func irVisibility(n *ast.Visibility) ll.Visibility {
	text := n.Text()
	switch text {
	case "":
		// \empty is used when visibility kind not present.
		return ll.VisibilityNone
	case "default":
		return ll.VisibilityDefault
	case "hidden":
		return ll.VisibilityHidden
	case "protected":
		return ll.VisibilityProtected
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
