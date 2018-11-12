package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
)

// === [ Aliases ] =============================================================

// Alias is an alias of a global variable or constant expression.
type Alias struct {
	// Alias name (without '@' prefix).
	GlobalName string
	// Aliasee.
	Aliasee constant.Constant

	// Pointer type of aliasee.
	Typ *types.PointerType
	// (optional) Linkage; zero value if not present.
	Linkage enum.Linkage
	// (optional) Preemption; zero value if not present.
	Preemption enum.Preemption
	// (optional) Visibility; zero value if not present.
	Visibility enum.Visibility
	// (optional) DLL storage class; zero value if not present.
	DLLStorageClass enum.DLLStorageClass
	// (optional) Thread local storage model; zero value if not present.
	TLSModel enum.TLSModel
	// (optional) Unnamed address; zero value if not present.
	UnnamedAddr enum.UnnamedAddr
}

// NewAlias returns a new alias based on the given alias name and aliasee.
func NewAlias(name string, aliasee constant.Constant) *Alias {
	alias := &Alias{GlobalName: name, Aliasee: aliasee}
	// Compute type.
	alias.Type()
	return alias
}

// String returns the LLVM syntax representation of the alias as a type-value
// pair.
func (a *Alias) String() string {
	return fmt.Sprintf("%s %s", a.Type(), a.Ident())
}

// Type returns the type of the alias.
func (a *Alias) Type() types.Type {
	// Cache type if not present.
	if a.Typ == nil {
		typ, ok := a.Aliasee.Type().(*types.PointerType)
		if !ok {
			panic(fmt.Errorf("invalid aliasee type of %q; expected *types.PointerType, got %T", a.Ident(), a.Aliasee.Type()))
		}
		a.Typ = typ
	}
	return a.Typ
}

// Ident returns the identifier associated with the alias.
func (a *Alias) Ident() string {
	return enc.Global(a.GlobalName)
}

// Name returns the name of the alias.
func (a *Alias) Name() string {
	return a.GlobalName
}

// SetName sets the name of the alias.
func (a *Alias) SetName(name string) {
	a.GlobalName = name
}

// Def returns the LLVM syntax representation of the alias definition.
func (a *Alias) Def() string {
	// Name=GlobalIdent '=' (ExternLinkage | Linkageopt) Preemptionopt
	// Visibilityopt DLLStorageClassopt ThreadLocalopt UnnamedAddropt 'alias'
	// ContentType=Type ',' Aliasee=TypeConst
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s =", a.Ident())
	if a.Linkage != enum.LinkageNone {
		fmt.Fprintf(buf, " %s", a.Linkage)
	}
	if a.Preemption != enum.PreemptionNone {
		fmt.Fprintf(buf, " %s", a.Preemption)
	}
	if a.Visibility != enum.VisibilityNone {
		fmt.Fprintf(buf, " %s", a.Visibility)
	}
	if a.DLLStorageClass != enum.DLLStorageClassNone {
		fmt.Fprintf(buf, " %s", a.DLLStorageClass)
	}
	if a.TLSModel != enum.TLSModelNone {
		fmt.Fprintf(buf, " %s", a.TLSModel)
	}
	if a.UnnamedAddr != enum.UnnamedAddrNone {
		fmt.Fprintf(buf, " %s", a.UnnamedAddr)
	}
	buf.WriteString(" alias")
	fmt.Fprintf(buf, " %s, %s", a.Typ.ElemType, a.Aliasee)
	return buf.String()
}
