// === [ Functions ] ===========================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#functions

package ir

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ Functions ] -----------------------------------------------------------

// A Function represents an LLVM IR function definition or external function
// declaration. The body of a function definition consists of a set of basic
// blocks, interconnected by control flow instructions.
//
// Functions may be referenced from instructions (e.g. call), and are thus
// considered LLVM IR values of function type.
type Function struct {
	// Parent module of the function.
	Parent *Module
	// Function name.
	Name string
	// Function type.
	Typ *types.PointerType
	// Function type.
	Sig *types.FuncType
	// Basic blocks of the function; or nil if defined externally.
	Blocks []*BasicBlock
	// Map from metadata identifier (e.g. !dbg) to metadata associated with the
	// function.
	Metadata map[string]*metadata.Metadata
}

// NewFunction returns a new function based on the given function name, return
// type and parameters.
func NewFunction(name string, ret types.Type, params ...*types.Param) *Function {
	sig := types.NewFunc(ret, params...)
	typ := types.NewPointer(sig)
	return &Function{
		Name:     name,
		Typ:      typ,
		Sig:      sig,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// Type returns the type of the function.
func (f *Function) Type() types.Type {
	return f.Typ
}

// Ident returns the identifier associated with the function.
func (f *Function) Ident() string {
	return enc.Global(f.Name)
}

// GetName returns the name of the function.
func (f *Function) GetName() string {
	return f.Name
}

// SetName sets the name of the function.
func (f *Function) SetName(name string) {
	f.Name = name
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Function) Immutable() {}

// MetadataNode ensures that only metadata nodes can be assigned to the
// metadata.Node interface.
func (*Function) MetadataNode() {}

// String returns the LLVM syntax representation of the function.
func (f *Function) String() string {
	// Assign unique local IDs to unnamed function parameters, basic blocks and
	// local variables.
	assignIDs(f)

	// Function signature.
	sig := &bytes.Buffer{}
	fmt.Fprintf(sig, "%s %s(",
		f.Sig.Ret,
		f.Ident())
	params := f.Params()
	for i, param := range params {
		if i != 0 {
			sig.WriteString(", ")
		}
		// Use same output format as Clang. Don't output local ID for unnamed
		// function parameters.
		if len(param.Name) > 0 && !isLocalID(param.Name) {
			fmt.Fprintf(sig, "%s %s",
				param.Type(),
				param.Ident())
		} else {
			sig.WriteString(param.Type().String())
		}
	}
	if f.Sig.Variadic {
		if len(params) > 0 {
			sig.WriteString(", ")
		}
		sig.WriteString("...")
	}
	sig.WriteString(")")

	// Metadata.
	var keys []string
	for key := range f.Metadata {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	metadata := &bytes.Buffer{}
	for _, key := range keys {
		md := f.Metadata[key]
		fmt.Fprintf(metadata, " %s %s", enc.Metadata(key), md.Ident())
	}

	// Function definition.
	if len(f.Blocks) > 0 {
		buf := &bytes.Buffer{}
		fmt.Fprintf(buf, "define %s%s {\n", sig, metadata)
		for _, block := range f.Blocks {
			fmt.Fprintln(buf, block)
		}
		buf.WriteString("}")
		return buf.String()
	}

	// External function declaration.
	return fmt.Sprintf("declare%s %s", metadata, sig)
}

// Params returns the parameters of the function.
func (f *Function) Params() []*types.Param {
	return f.Sig.Params
}

// AppendParam appends the given function parameter to the function.
func (f *Function) AppendParam(param *types.Param) {
	f.Sig.Params = append(f.Sig.Params, param)
}

// NewParam appends a new function parameter to the function based on the given
// parameter name and type.
func (f *Function) NewParam(name string, typ types.Type) *types.Param {
	param := types.NewParam(name, typ)
	f.AppendParam(param)
	return param
}

// AppendBlock appends the given basic block to the function.
func (f *Function) AppendBlock(block *BasicBlock) {
	block.Parent = f
	f.Blocks = append(f.Blocks, block)
}

// NewBlock appends a new basic block to the function based on the given label
// name. An empty label name indicates an unnamed basic block.
func (f *Function) NewBlock(name string) *BasicBlock {
	block := NewBlock(name)
	f.AppendBlock(block)
	return block
}

// --- [ Function parameters ] -------------------------------------------------

// NewParam returns a new function parameter based on the given parameter name
// and type.
func NewParam(name string, typ types.Type) *types.Param {
	return types.NewParam(name, typ)
}

// ### [ Helper functions ] ####################################################

// assignIDs assigns unique local IDs to unnamed basic blocks and local
// variables of the function.
func assignIDs(f *Function) {
	id := 0
	setName := func(n value.Named) {
		name := n.GetName()
		switch {
		case isUnnamed(name):
			n.SetName(strconv.Itoa(id))
			id++
		case isLocalID(name):
			want := strconv.Itoa(id)
			if name != want {
				panic(fmt.Errorf("invalid local ID in function %s; expected %s, got %s", enc.Global(f.Name), enc.Local(want), enc.Local(name)))
			}
			id++
		}
	}
	for _, param := range f.Params() {
		// Assign local IDs to unnamed parameters of function definitions.
		if len(f.Blocks) > 0 {
			setName(param)
		}
	}
	for _, block := range f.Blocks {
		// Assign local IDs to unnamed basic blocks.
		setName(block)
		for _, inst := range block.Insts {
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
