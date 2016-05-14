package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir/types"
)

// A Function declaration specifies the name and type of a function. A function
// definition contains a set of basic blocks, interconnected by control flow
// instructions (such as br), which forms the nodes in a Control Flow Graph of
// the function [1,2].
//
//    [1]: http://llvm.org/docs/LangRef.html#functions
//    [2]: http://llvm.org/docs/LangRef.html#terminators
type Function struct {
	// Function name.
	name string
	// Function signature.
	sig *types.Func
	// Basic blocks of the function, or nil if function declaration.
	blocks []*BasicBlock
}

// NewFunction returns a new function based on the given name and function
// signature.
func NewFunction(name string, sig *types.Func) *Function {
	return &Function{name: name, sig: sig}
}

// Name returns the name of the function.
func (f *Function) Name() string {
	return f.name
}

// Type returns the function signature.
func (f *Function) Type() *types.Func {
	return f.sig
}

// TODO: Try to figure out a better API for accessing the function body.

// Blocks returns the basic blocks of the function.
func (f *Function) Blocks() []*BasicBlock {
	return f.blocks
}

// SetBlocks sets the basic blocks of the function body to the given basic
// blocks.
func (f *Function) SetBlocks(blocks []*BasicBlock) {
	f.blocks = blocks
}

// AppendBlock appends the given block to the basic blocks of the function body.
func (f *Function) AppendBlock(block *BasicBlock) {
	f.blocks = append(f.blocks, block)
}

// String returns the string representation of the function declaration.
func (f *Function) String() string {
	// Function signature; e.g.
	//    "void @foo()"
	//    "i32 @printf(i8*, ...)"
	paramsBuf := new(bytes.Buffer)
	for i, param := range f.sig.Params() {
		if i > 0 {
			paramsBuf.WriteString(", ")
		}
		paramsBuf.WriteString(param.String())
	}
	if f.sig.IsVariadic() {
		if len(f.sig.Params()) > 0 {
			paramsBuf.WriteString(", ")
		}
		paramsBuf.WriteString("...")
	}
	sig := fmt.Sprintf("%s %s(%s)", f.sig.Result(), asm.EncGlobal(f.Name()), paramsBuf)

	// Function declaration; e.g.
	//    declare i32 @printf(i8*, ...)
	if f.blocks == nil {
		return fmt.Sprintf("declare %s", sig)
	}

	// Function definition; e.g.
	//     define i32 @main() {
	//      ret i32 42
	//   }
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "define %s {\n", sig)
	for _, block := range f.blocks {
		fmt.Fprintln(buf, block)
	}
	buf.WriteString("}\n")
	return buf.String()
}
