package ir

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// === [ Functions ] ===========================================================

// Function is an LLVM IR function. The body of a function definition consists
// of a set of basic blocks, interconnected by terminator control flow
// instructions.
type Function struct {
	// Function name (without '@' prefix).
	GlobalName string
	// Function signature.
	Sig *types.FuncType
	// Function parameters.
	Params []*Param
	// Basic blocks.
	Blocks []*BasicBlock // nil if declaration.

	// extra.

	// Pointer type to function, including an optional address space. If Typ is
	// nil, the first invocation of Type stores a pointer type with Sig as
	// element.
	Typ *types.PointerType
	// (optional) Linkage.
	Linkage enum.Linkage
	// (optional) Preemption; zero value if not present.
	Preemption enum.Preemption
	// (optional) Visibility; zero value if not present.
	Visibility enum.Visibility
	// (optional) DLL storage class; zero value if not present.
	DLLStorageClass enum.DLLStorageClass
	// (optional) Calling convention; zero value if not present.
	CallingConv enum.CallingConv
	// (optional) Return attributes.
	ReturnAttrs []ReturnAttribute
	// (optional) Unnamed address.
	UnnamedAddr enum.UnnamedAddr
	// (optional) Function attributes.
	FuncAttrs []FuncAttribute
	// (optional) Section name; empty if not present.
	Section string
	// (optional) Comdat; nil if not present.
	Comdat *ComdatDef
	// (optional) Garbage collection; empty if not present.
	GC string
	// (optional) Prefix; nil if not present.
	Prefix constant.Constant
	// (optional) Prologue; nil if not present.
	Prologue constant.Constant
	// (optional) Personality; nil if not present.
	Personality constant.Constant
	// (optional) Use list orders.
	UseListOrders []*UseListOrder
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment

	// mu prevents races on AssignIDs.
	mu sync.Mutex
}

// NewFunc returns a new function based on the given function name, return type
// and function parameters.
func NewFunc(name string, retType types.Type, params ...*Param) *Function {
	paramTypes := make([]types.Type, len(params))
	for i, param := range params {
		paramTypes[i] = param.Type()
	}
	sig := types.NewFunc(retType, paramTypes...)
	f := &Function{GlobalName: name, Sig: sig, Params: params}
	// Compute type.
	f.Type()
	return f
}

// String returns the LLVM syntax representation of the function as a type-value
// pair.
func (f *Function) String() string {
	return fmt.Sprintf("%s %s", f.Type(), f.Ident())
}

// Type returns the type of the function.
func (f *Function) Type() types.Type {
	// Cache type if not present.
	if f.Typ == nil {
		f.Typ = types.NewPointer(f.Sig)
	}
	return f.Typ
}

// Ident returns the identifier associated with the function.
func (f *Function) Ident() string {
	return enc.Global(f.GlobalName)
}

// Name returns the name of the function.
func (f *Function) Name() string {
	return f.GlobalName
}

// SetName sets the name of the function.
func (f *Function) SetName(name string) {
	f.GlobalName = name
}

// Def returns the LLVM syntax representation of the function definition or
// declaration.
func (f *Function) Def() string {
	// Function declaration.
	//
	//    'declare' Metadata=MetadataAttachment* Header=FuncHeader
	//
	// Function definition.
	//
	//    'define' Header=FuncHeader Metadata=MetadataAttachment* Body=FuncBody
	buf := &strings.Builder{}
	if len(f.Blocks) == 0 {
		// Function declaration.
		buf.WriteString("declare")
		for _, md := range f.Metadata {
			fmt.Fprintf(buf, " %s", md)
		}
		if f.Linkage != enum.LinkageNone {
			fmt.Fprintf(buf, " %s", f.Linkage)
		}
		buf.WriteString(headerString(f))
		return buf.String()
	}
	// Function definition.
	if err := f.AssignIDs(); err != nil {
		panic(fmt.Errorf("unable to assign IDs of function %q; %v", f.Ident(), err))
	}
	buf.WriteString("define")
	if f.Linkage != enum.LinkageNone {
		fmt.Fprintf(buf, " %s", f.Linkage)
	}
	buf.WriteString(headerString(f))
	for _, md := range f.Metadata {
		fmt.Fprintf(buf, " %s", md)
	}
	fmt.Fprintf(buf, " %s", bodyString(f))
	return buf.String()
}

// AssignIDs assigns IDs to unnamed local variables.
func (f *Function) AssignIDs() error {
	if len(f.Blocks) == 0 {
		return nil
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	id := int64(0)
	setName := func(n local) error {
		if n.IsUnnamed() {
			if n.ID() != 0 && id != n.ID() {
				want := strconv.FormatInt(id, 10)
				got := strconv.FormatInt(n.ID(), 10)
				return errors.Errorf("invalid local ID in function %q, expected %s, got %s", enc.Global(f.GlobalName), enc.Local(want), enc.Local(got))
			}
			n.SetID(id)
			id++
		}
		return nil
	}
	for _, param := range f.Params {
		// Assign local IDs to unnamed parameters of function definitions.
		if err := setName(param); err != nil {
			return errors.WithStack(err)
		}
	}
	for _, block := range f.Blocks {
		// Assign local IDs to unnamed basic blocks.
		if err := setName(block); err != nil {
			return errors.WithStack(err)
		}
		for _, inst := range block.Insts {
			n, ok := inst.(local)
			if !ok {
				continue
			}
			// Skip void instructions.
			// TODO: Check if any other value instructions than call may have void
			// type.
			if isVoidValue(n) {
				continue
			}
			// Assign local IDs to unnamed local variables.
			if err := setName(n); err != nil {
				return errors.WithStack(err)
			}
		}
		n, ok := block.Term.(local)
		if !ok {
			continue
		}
		if isVoidValue(n) {
			continue
		}
		if err := setName(n); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// ### [ Helper functions ] ####################################################

// headerString returns the string representation of the function header.
func headerString(f *Function) string {
	// (Linkage | ExternLinkage)? Preemptionopt Visibilityopt DLLStorageClassopt
	// CallingConvopt ReturnAttrs=ReturnAttribute* RetType=Type Name=GlobalIdent
	// '(' Params ')' UnnamedAddropt AddrSpaceopt FuncAttrs=FuncAttribute*
	// Sectionopt Comdatopt GCopt Prefixopt Prologueopt Personalityopt
	buf := &strings.Builder{}
	if f.Preemption != enum.PreemptionNone {
		fmt.Fprintf(buf, " %s", f.Preemption)
	}
	if f.Visibility != enum.VisibilityNone {
		fmt.Fprintf(buf, " %s", f.Visibility)
	}
	if f.DLLStorageClass != enum.DLLStorageClassNone {
		fmt.Fprintf(buf, " %s", f.DLLStorageClass)
	}
	if f.CallingConv != enum.CallingConvNone {
		fmt.Fprintf(buf, " %s", callingConvString(f.CallingConv))
	}
	for _, attr := range f.ReturnAttrs {
		fmt.Fprintf(buf, " %s", attr)
	}
	fmt.Fprintf(buf, " %s", f.Sig.RetType)
	fmt.Fprintf(buf, " %s(", enc.Global(f.GlobalName))
	for i, param := range f.Params {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(param.Def())
	}
	if f.Sig.Variadic {
		if len(f.Params) > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("...")
	}
	buf.WriteString(")")
	if f.UnnamedAddr != enum.UnnamedAddrNone {
		fmt.Fprintf(buf, " %s", f.UnnamedAddr)
	}
	for _, attr := range f.FuncAttrs {
		fmt.Fprintf(buf, " %s", attr)
	}
	if len(f.Section) > 0 {
		fmt.Fprintf(buf, " section %s", quote(f.Section))
	}
	if f.Comdat != nil {
		if f.Comdat.Name == f.GlobalName {
			buf.WriteString(" comdat")
		} else {
			fmt.Fprintf(buf, " %s", f.Comdat)
		}
	}
	if len(f.GC) > 0 {
		fmt.Fprintf(buf, " gc %s", quote(f.GC))
	}
	if f.Prefix != nil {
		fmt.Fprintf(buf, " prefix %s", f.Prefix)
	}
	if f.Prologue != nil {
		fmt.Fprintf(buf, " prologue %s", f.Prologue)
	}
	if f.Personality != nil {
		fmt.Fprintf(buf, " personality %s", f.Personality)
	}
	return buf.String()
}

// bodyString returns the string representation of the function body.
func bodyString(body *Function) string {
	// '{' Blocks=BasicBlock+ UseListOrders=UseListOrder* '}'
	buf := &strings.Builder{}
	buf.WriteString("{\n")
	for i, block := range body.Blocks {
		if i != 0 {
			buf.WriteString("\n")
		}
		fmt.Fprintf(buf, "%s\n", block.Def())
	}
	if len(body.UseListOrders) > 0 {
		buf.WriteString("\n")
	}
	for _, u := range body.UseListOrders {
		fmt.Fprintf(buf, "\t%s\n", u)
	}
	buf.WriteString("}")
	return buf.String()
}

// isVoidValue reports whether the given named value is a non-value (i.e. a call
// instruction or invoke terminator with void-return type).
func isVoidValue(n value.Named) bool {
	switch n.(type) {
	case *InstCall, *TermInvoke:
		return n.Type().Equal(types.Void)
	}
	return false
}

// local is a local variable.
type local interface {
	value.Named
	// ID returns the ID of the local identifier.
	ID() int64
	// SetID sets the ID of the local identifier.
	SetID(id int64)
	// IsUnnamed reports whether the local identifier is unnamed.
	IsUnnamed() bool
}
