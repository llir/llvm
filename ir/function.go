// === [ Functions ] ===========================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#functions

package ir

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/llvm/ir/internal/enc"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// A Function represents an LLVM IR function definition or external function
// declaration. The body of a function definition consists of a set of basic
// blocks, interconnected by control flow instructions.
//
// Functions may be referenced from terminators (e.g. call), and are thus
// considered LLVM IR values of function type.
type Function struct {
	// Parent module of the function.
	parent *Module
	// Function name.
	name string
	// Function type.
	typ *types.FuncType
	// Basic blocks of the function; or nil if external function declaration.
	blocks []*BasicBlock
}

// NewFunction returns a new function based on the given function name, return
// type and parameters.
func NewFunction(name string, ret types.Type, params ...*types.Param) *Function {
	typ := types.NewFunc(ret, params...)
	return &Function{name: name, typ: typ}
}

// Type returns the type of the function.
func (f *Function) Type() types.Type {
	return f.typ
}

// Ident returns the identifier associated with the function.
func (f *Function) Ident() string {
	return enc.Global(f.name)
}

// String returns the LLVM syntax representation of the function.
func (f *Function) String() string {
	// Assign unique local IDs to unnamed basic blocks and local variables.
	assignIDs(f)
	// Function signature.
	sig := &bytes.Buffer{}
	fmt.Fprintf(sig, "%s %s(",
		f.RetType(),
		f.Ident())
	params := f.Params()
	for i, p := range params {
		if i != 0 {
			sig.WriteString(", ")
		}
		fmt.Fprintf(sig, "%s %s",
			p.Type(),
			p.Ident())
	}
	if f.Variadic() {
		if len(params) > 0 {
			sig.WriteString(", ")
		}
		sig.WriteString("...")
	}
	sig.WriteString(")")
	if blocks := f.Blocks(); len(blocks) > 0 {
		// Function definition.
		buf := &bytes.Buffer{}
		fmt.Fprintf(buf, "define %s {\n", sig)
		for _, block := range blocks {
			fmt.Fprintln(buf, block)
		}
		buf.WriteString("}")
		return buf.String()
	}
	// External function declaration.
	return fmt.Sprintf("declare %s", sig)
}

// Parent returns the parent module of the function.
func (f *Function) Parent() *Module {
	return f.parent
}

// SetParent sets the parent module of the function.
func (f *Function) SetParent(parent *Module) {
	f.parent = parent
}

// RetType returns the return type of the function.
func (f *Function) RetType() types.Type {
	return f.typ.RetType()
}

// Params returns the function parameters of the function.
func (f *Function) Params() []*types.Param {
	return f.typ.Params()
}

// Variadic reports whether the function is variadic.
func (f *Function) Variadic() bool {
	return f.typ.Variadic()
}

// SetVariadic sets the variadicity of the function.
func (f *Function) SetVariadic(variadic bool) {
	f.typ.SetVariadic(variadic)
}

// Blocks returns the basic blocks of the function.
func (f *Function) Blocks() []*BasicBlock {
	return f.blocks
}

// AppendParam appends the given function parameter to the function.
func (f *Function) AppendParam(p *types.Param) {
	f.typ.AppendParam(p)
}

// AppendBlock appends the given basic block to the function.
func (f *Function) AppendBlock(block *BasicBlock) {
	block.SetParent(f)
	f.blocks = append(f.blocks, block)
}

// NewParam appends a new function parameter to the function based on the given
// parameter name and type.
func (f *Function) NewParam(name string, typ types.Type) *types.Param {
	return f.typ.NewParam(name, typ)
}

// NewBlock appends a new basic block to the function based on the given label
// name. An empty label name indicates an unnamed basic block.
func (f *Function) NewBlock(name string) *BasicBlock {
	block := NewBlock(name)
	f.AppendBlock(block)
	return block
}

// assignIDs assigns unique local IDs to unnamed basic blocks and local
// variables of the function.
func assignIDs(f *Function) {
	// identifiable represents a basic block or local variable.
	type identifiable interface {
		value.Value
		// SetIdent sets the identifier associated with the value.
		SetIdent(ident string)
	}
	id := 0
	setName := func(n identifiable) {
		got := n.Ident()
		switch {
		case isUnnamed(got):
			n.SetIdent(strconv.Itoa(id))
			id++
		case isLocalID(got):
			want := enc.Local(strconv.Itoa(id))
			if got != want {
				panic(fmt.Sprintf("invalid local ID; expected %s, got %s", want, got))
			}
			id++
		}
	}
	for _, block := range f.blocks {
		setName(block)
		for _, inst := range block.insts {
			n, ok := inst.(identifiable)
			if !ok {
				continue
			}
			if n.Type().Equal(types.Void) {
				continue
			}
			setName(n)
		}
	}
}

// isUnnamed reports whether the given identifier is unnamed.
func isUnnamed(ident string) bool {
	return len(ident) == 0 || ident == "%" || ident == "@"
}

// isLocalID reports whether the given identifier is a local ID (e.g. "%42").
func isLocalID(ident string) bool {
	if !strings.HasPrefix(ident, "%") {
		return false
	}
	ident = ident[1:]
	for _, r := range ident {
		if strings.IndexRune("0123456789", r) == -1 {
			return false
		}
	}
	return len(ident) > 0
}
