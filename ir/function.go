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

	"github.com/llir/llvm/internal/enc"
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
	typ *types.PointerType
	// Function type.
	sig *types.FuncType
	// Function parameters.
	params []*Param
	// Basic blocks of the function; or nil if defined externally.
	blocks []*BasicBlock
	// Track uses of the value.
	used
}

// NewFunction returns a new function based on the given function name, return
// type and parameters.
func NewFunction(name string, ret types.Type, params ...*Param) *Function {
	sig := types.NewFunc(ret)
	for _, param := range params {
		sig.AppendParam(param.Param)
	}
	typ := types.NewPointer(sig)
	return &Function{name: name, typ: typ, sig: sig, params: params}
}

// Type returns the type of the function.
func (f *Function) Type() types.Type {
	return f.typ
}

// Ident returns the identifier associated with the function.
func (f *Function) Ident() string {
	return enc.Global(f.name)
}

// Name returns the name of the function.
func (f *Function) Name() string {
	return f.name
}

// SetName sets the name of the function.
func (f *Function) SetName(name string) {
	f.name = name
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Function) Immutable() {}

// String returns the LLVM syntax representation of the function.
func (f *Function) String() string {
	// Assign unique local IDs to unnamed function parameters, basic blocks and
	// local variables.
	assignIDs(f)

	// Function signature.
	sig := &bytes.Buffer{}
	fmt.Fprintf(sig, "%s %s(",
		f.RetType(),
		f.Ident())
	params := f.Params()
	for i, param := range params {
		if i != 0 {
			sig.WriteString(", ")
		}
		if len(param.Name()) > 0 {
			fmt.Fprintf(sig, "%s %s",
				param.Type(),
				param.Ident())
		} else {
			sig.WriteString(param.Type().String())
		}
	}
	if f.Variadic() {
		if len(params) > 0 {
			sig.WriteString(", ")
		}
		sig.WriteString("...")
	}
	sig.WriteString(")")

	// Function definition.
	if blocks := f.Blocks(); len(blocks) > 0 {
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

// Sig returns the signature of the function.
func (f *Function) Sig() *types.FuncType {
	return f.sig
}

// RetType returns the return type of the function.
func (f *Function) RetType() types.Type {
	return f.sig.RetType()
}

// Params returns the function parameters of the function.
func (f *Function) Params() []*Param {
	return f.params
}

// Variadic reports whether the function is variadic.
func (f *Function) Variadic() bool {
	return f.sig.Variadic()
}

// SetVariadic sets the variadicity of the function.
func (f *Function) SetVariadic(variadic bool) {
	f.sig.SetVariadic(variadic)
}

// Blocks returns the basic blocks of the function.
func (f *Function) Blocks() []*BasicBlock {
	return f.blocks
}

// AppendParam appends the given function parameter to the function.
func (f *Function) AppendParam(param *Param) {
	f.sig.AppendParam(param.Param)
	f.params = append(f.params, param)
}

// AppendBlock appends the given basic block to the function.
func (f *Function) AppendBlock(block *BasicBlock) {
	block.SetParent(f)
	f.blocks = append(f.blocks, block)
}

// NewParam appends a new function parameter to the function based on the given
// parameter name and type.
func (f *Function) NewParam(name string, typ types.Type) *Param {
	param := NewParam(name, typ)
	f.AppendParam(param)
	return param
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
	id := 0
	setName := func(n value.Named) {
		name := n.Name()
		switch {
		case isUnnamed(name):
			n.SetName(strconv.Itoa(id))
			id++
		case isLocalID(name):
			want := strconv.Itoa(id)
			if name != want {
				panic(fmt.Sprintf("invalid local ID; expected %s, got %s", enc.Local(want), enc.Local(name)))
			}
			id++
		}
	}
	for _, param := range f.Params() {
		// Assign local IDs to unnamed parameters of function definitions.
		if len(f.blocks) > 0 {
			setName(param)
		}
	}
	for _, block := range f.blocks {
		// Assign local IDs to unnamed basic blocks.
		setName(block)
		for _, inst := range block.insts {
			n, ok := inst.(value.Named)
			if !ok {
				continue
			}
			if n.Type().Equal(types.Void) {
				continue
			}
			// Assign local IDs to unnamed local variables.
			setName(n)
		}
	}
}

// isUnnamed reports whether the given identifier is unnamed.
func isUnnamed(name string) bool {
	return len(name) == 0
}

// isLocalID reports whether the given identifier is a local ID (e.g. "%42").
func isLocalID(name string) bool {
	for _, r := range name {
		if strings.IndexRune("0123456789", r) == -1 {
			return false
		}
	}
	return len(name) > 0
}
