// === [ Global variables ] ====================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#global-variables

package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
)

type LinkageType int

const (
	// Extern at position 0 so it's used as the default

	External LinkageType = iota
	Private
	Internal
	AvailableExternally
	LinkOnce
	Weak
	Common
	Appending
	ExternWeak
	LinkOnceODR
	WeakODR
)

var linkageNames = []string{
	External:            "external",
	Private:             "private",
	Internal:            "internal",
	AvailableExternally: "available_externally",
	LinkOnce:            "linkonce",
	Weak:                "weak",
	Common:              "common",
	Appending:           "appending",
	ExternWeak:          "extern_weak",
	LinkOnceODR:         "linkonce_odr",
	WeakODR:             "weak_odr",
}

func (t LinkageType) String() string {
	return linkageNames[t]
}

type UnnamedAddrType int

const (
	UnnamedAddr = iota + 1
	LocalUnnamedAddr
)

var unnamedAddrNames = []string{
	0:                "",
	UnnamedAddr:      "unnamed_addr",
	LocalUnnamedAddr: "local_unnamed_addr",
}

func (t UnnamedAddrType) String() string {
	return unnamedAddrNames[t]
}

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
	Name string
	// Global variable type.
	Typ *types.PointerType
	// Content type.
	Content types.Type
	// Initial value; or nil if defined externally.
	Init constant.Constant
	// Immutability of the global variable.
	IsConst bool
	// Linkage type of the global variable.
	LinkageType LinkageType
	UnnamedAddr UnnamedAddrType
	// Map from metadata identifier (e.g. !dbg) to metadata associated with the
	// global.
	Metadata map[string]*metadata.Metadata
}

// NewGlobalDecl returns a new external global variable declaration based on the
// given global variable name and content type.
func NewGlobalDecl(name string, content types.Type) *Global {
	typ := types.NewPointer(content)
	return &Global{
		Name:     name,
		Typ:      typ,
		Content:  content,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// NewGlobalDef returns a new global variable definition based on the given
// global variable name and initial value.
func NewGlobalDef(name string, init constant.Constant) *Global {
	content := init.Type()
	typ := types.NewPointer(content)
	return &Global{
		Name:     name,
		Typ:      typ,
		Content:  content,
		Init:     init,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// Type returns the type of the global variable.
func (global *Global) Type() types.Type {
	return global.Typ
}

// Ident returns the identifier associated with the global variable.
func (global *Global) Ident() string {
	return enc.Global(global.Name)
}

// GetName returns the name of the global variable.
func (global *Global) GetName() string {
	return global.Name
}

// SetName sets the name of the global variable.
func (global *Global) SetName(name string) {
	global.Name = name
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Global) Immutable() {}

// MetadataNode ensures that only metadata nodes can be assigned to the
// metadata.Node interface.
func (*Global) MetadataNode() {}

// String returns the LLVM syntax representation of the global variable.
func (global *Global) String() string {
	// @<GlobalVarName> = [Linkage] [Visibility] [DLLStorageClass] [ThreadLocal]
	//                [(unnamed_addr|local_unnamed_addr)] [AddrSpace]
	//                [ExternallyInitialized]
	//                <global | constant> <Type> [<InitializerConstant>]
	//                [, section "name"] [, comdat [($name)]]
	//                [, align <Alignment>] (, !name !N)*

	md := metadataString(global.Metadata, ",")
	var attributes []string
	if global.LinkageType != External || global.Init == nil {
		attributes = append(attributes, global.LinkageType.String())
	}
	if global.Typ.AddrSpace != 0 {
		attributes = append(attributes, fmt.Sprintf("addrspace(%d)", global.Typ.AddrSpace))
	}
	if s := global.UnnamedAddr.String(); s != "" {
		attributes = append(attributes, s)
	}
	if global.IsConst {
		attributes = append(attributes, "constant")
	} else {
		attributes = append(attributes, "global")
	}
	if global.Init != nil {
		// Global variable definition.
		attributes = append(attributes, global.Init.Type().String(), global.Init.Ident())
	} else {
		// External global variable declaration.
		attributes = append(attributes, global.Content.String())
	}

	attrString := strings.Join(attributes, " ")
	return fmt.Sprintf("%s = %s%s",
		global.Ident(),
		attrString,
		md)
}
