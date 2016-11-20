package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// A Function represents an LLVM IR function definition or external function
// declaration. The body of a function definition consists of a set of basic
// blocks, interconnected by control flow instructions.
//
// Functions may be referenced from terminator instructions (e.g. call), and are
// thus considered LLVM IR values of function type.
type Function struct {
	// Parent module of the function.
	parent *Module
	// Function name.
	name string
	// Return type.
	ret types.Type
	// Function parameters.
	params []*Param
	// Function type.
	typ *types.FuncType
	// Basic blocks of the function.
	blocks []*BasicBlock
}

// NewFunction returns a new LLVM IR function based on the given name, return
// type and parameters.
func NewFunction(name string, ret types.Type, params ...*Param) *Function {
	var ps []types.Type
	for _, param := range params {
		ps = append(ps, param.typ)
	}
	typ := types.NewFuncType(ret, ps...)
	return &Function{name: name, ret: ret, params: params, typ: typ}
}

// Type returns the type of the function.
func (f *Function) Type() types.Type {
	return f.typ
}

// Ident returns the identifier associated with the function.
func (f *Function) Ident() string {
	// TODO: Encode name if containing special characters.
	return "@" + f.name
}

// LLVMString returns the LLVM syntax representation of the function.
func (f *Function) LLVMString() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%v %v(", f.ret.LLVMString(), f.Ident())
	for i, param := range f.params {
		if i != 0 {
			buf.WriteString(", ")
			fmt.Fprintf(buf, param.typ.LLVMString(), param.name)
		}
	}
	buf.WriteString(")")
	if len(f.blocks) > 0 {
		fmt.Fprintln(buf, " {")
		for _, block := range f.blocks {
			fmt.Fprintln(buf, block.LLVMString())
		}
		buf.WriteString("}")
	}
	return buf.String()
}

// NewParam appends a new parameter to the function based on the given parameter
// name and type.
func (f *Function) NewParam(name string, typ types.Type) *Param {
	param := &Param{name: name, typ: typ}
	f.params = append(f.params, param)
	return param
}

// NewBlock appends a new basic block to the function based on the given basic
// block label name.
func (f *Function) NewBlock(name string) *BasicBlock {
	block := &BasicBlock{name: name}
	block.SetParent(f)
	f.blocks = append(f.blocks, block)
	return block
}

// Parent returns the parent module of the function.
func (f *Function) Parent() *Module {
	return f.parent
}

// SetParent sets the parent module of the function.
func (f *Function) SetParent(parent *Module) {
	f.parent = parent
}

// A Param represents a function parameter.
type Param struct {
	// Parameter name.
	name string
	// Parameter type.
	typ types.Type
}

// NewParam returns a new function parameter based on the given parameter name
// and type.
func NewParam(name string, typ types.Type) *Param {
	return &Param{name: name, typ: typ}
}

// Type returns the type of the function parameter.
func (p *Param) Type() types.Type {
	return p.typ
}

// Ident returns the identifier associated with the function parameter.
func (p *Param) Ident() string {
	// TODO: Encode name if containing special characters.
	return "%" + p.name
}
