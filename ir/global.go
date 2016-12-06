// === [ Global variables ] ====================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#global-variables

package ir

import (
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// A Global represents an LLVM IR global variable definition or external global
// variable declaration.
//
// Global variables always define a pointer to their "content" type because they
// describe a region of memory, and all memory objects in LLVM are accessed
// through pointers.
//
// Global variables may be referenced from instructions (e.g. load), and are
// thus considered LLVM IR values of pointer type.
type Global struct {
	// Global variable name.
	name string
	// Content type.
	content types.Type
	// Initial value; or nil if defined externally.
	init constant.Constant
	// Global variable type.
	typ *types.PointerType
	// Immutability of the global variable.
	isConst bool
}

// NewGlobalDecl returns a new external global variable declaration based on the
// given global variable name and content type.
func NewGlobalDecl(name string, content types.Type) *Global {
	typ := types.NewPointer(content)
	return &Global{name: name, content: content, typ: typ}
}

// NewGlobalDef returns a new global variable definition based on the given
// global variable name and initial value.
func NewGlobalDef(name string, init constant.Constant) *Global {
	content := init.Type()
	typ := types.NewPointer(content)
	return &Global{name: name, content: content, init: init, typ: typ}
}

// Type returns the type of the global variable.
func (global *Global) Type() types.Type {
	return global.typ
}

// Ident returns the identifier associated with the global variable.
func (global *Global) Ident() string {
	return enc.Global(global.name)
}

// Name returns the name of the global variable.
func (global *Global) Name() string {
	return global.name
}

// SetName sets the name of the global variable.
func (global *Global) SetName(name string) {
	global.name = name
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Global) Immutable() {}

// String returns the LLVM syntax representation of the global variable.
func (global *Global) String() string {
	imm := "global"
	if global.Const() {
		imm = "constant"
	}
	content := global.ContentType()
	if init, ok := global.Init(); ok {
		// Global variable definition.
		return fmt.Sprintf("%s = %s %s %s",
			global.Ident(),
			imm,
			content,
			init.Ident())
	}
	// External global variable declaration.
	return fmt.Sprintf("%s = external %s %s",
		global.Ident(),
		imm,
		content)
}

// ContentType returns the content type of the global variable.
func (global *Global) ContentType() types.Type {
	return global.content
}

// Init returns the initial value of the global variable and a boolean
// indicating if an initializer was present.
func (global *Global) Init() (constant.Constant, bool) {
	if global.init != nil {
		return global.init, true
	}
	return nil, false
}

// SetInit sets the initial value of the global variable.
func (global *Global) SetInit(init constant.Constant) {
	global.init = init
}

// Const reports whether the global variable is a constant.
func (global *Global) Const() bool {
	return global.isConst
}

// SetConst sets the immutability of the global variable.
func (global *Global) SetConst(isConst bool) {
	global.isConst = isConst
}
