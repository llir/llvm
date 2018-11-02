// === [ Other types ] =========================================================

package types

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
)

// --- [ void ] ----------------------------------------------------------------

// VoidType represents a void type.
//
// References:
//    http://llvm.org/docs/LangRef.html#void-type
type VoidType struct {
	// Type name alias.
	Name string
}

// String returns the LLVM syntax representation of the type.
func (t *VoidType) String() string {
	if len(t.Name) > 0 {
		return enc.Local(t.Name)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *VoidType) Def() string {
	return "void"
}

// Equal reports whether t and u are of equal type.
func (t *VoidType) Equal(u Type) bool {
	_, ok := u.(*VoidType)
	return ok
}

// GetName returns the name of the type.
func (t *VoidType) GetName() string {
	return t.Name
}

// SetName sets the name of the type.
func (t *VoidType) SetName(name string) {
	t.Name = name
}

// --- [ function ] ------------------------------------------------------------

// FuncType represents a function type.
//
// References:
//    http://llvm.org/docs/LangRef.html#function-type
type FuncType struct {
	// Type name alias.
	Name string
	// Return type.
	Ret Type
	// Function parameters.
	Params []*Param
	// Variadicity of the function type.
	Variadic bool
}

// NewFunc returns a new function type based on the given return type and
// parameters.
func NewFunc(ret Type, params ...*Param) *FuncType {
	return &FuncType{Ret: ret, Params: params}
}

// String returns the LLVM syntax representation of the type.
func (t *FuncType) String() string {
	if len(t.Name) > 0 {
		return enc.Local(t.Name)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *FuncType) Def() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s (", t.Ret)
	for i, param := range t.Params {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(param.Typ.String())
	}
	if t.Variadic {
		if len(t.Params) > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("...")
	}
	buf.WriteString(")")
	return buf.String()
}

// Equal reports whether t and u are of equal type.
func (t *FuncType) Equal(u Type) bool {
	if u, ok := u.(*FuncType); ok {
		if !t.Ret.Equal(u.Ret) {
			return false
		}
		if len(t.Params) != len(u.Params) {
			return false
		}
		for i, tp := range t.Params {
			up := u.Params[i]
			if !tp.Typ.Equal(up.Typ) {
				return false
			}
		}
		return t.Variadic == u.Variadic
	}
	return false
}

// GetName returns the name of the type.
func (t *FuncType) GetName() string {
	return t.Name
}

// SetName sets the name of the type.
func (t *FuncType) SetName(name string) {
	t.Name = name
}

// NewParam appends a new function parameter to the function type based on the
// given parameter name and type.
func (t *FuncType) NewParam(name string, typ Type) *Param {
	param := NewParam(name, typ)
	t.Params = append(t.Params, param)
	return param
}

// A Param represents an LLVM IR function parameter.
//
// Function parameters may be referenced from instructions (e.g. add), and are
// thus considered LLVM IR values.
type Param struct {
	// Parameter name.
	Name string
	// Parameter type.
	Typ Type
}

// NewParam returns a new function parameter based on the given parameter name
// and type.
func NewParam(name string, typ Type) *Param {
	return &Param{Name: name, Typ: typ}
}

// Type returns the type of the function parameter.
func (param *Param) Type() Type {
	return param.Typ
}

// Ident returns the identifier associated with the function parameter.
func (param *Param) Ident() string {
	return enc.Local(param.Name)
}

// GetName returns the name of the function parameter.
func (param *Param) GetName() string {
	return param.Name
}

// SetName sets the name of the function parameter.
func (param *Param) SetName(name string) {
	param.Name = name
}

// --- [ label ] ---------------------------------------------------------------

// LabelType represents a label type, which is used for basic block values.
//
// References:
//    http://llvm.org/docs/LangRef.html#label-type
type LabelType struct {
	// Type name alias.
	Name string
}

// String returns the LLVM syntax representation of the type.
func (t *LabelType) String() string {
	if len(t.Name) > 0 {
		return enc.Local(t.Name)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *LabelType) Def() string {
	return "label"
}

// Equal reports whether t and u are of equal type.
func (t *LabelType) Equal(u Type) bool {
	_, ok := u.(*LabelType)
	return ok
}

// GetName returns the name of the type.
func (t *LabelType) GetName() string {
	return t.Name
}

// SetName sets the name of the type.
func (t *LabelType) SetName(name string) {
	t.Name = name
}

// --- [ metadata ] ------------------------------------------------------------

// MetadataType represents a metadata type.
//
// References:
//    http://llvm.org/docs/LangRef.html#metadata-type
type MetadataType struct {
	// Type name alias.
	Name string
}

// String returns the LLVM syntax representation of the type.
func (t *MetadataType) String() string {
	if len(t.Name) > 0 {
		return enc.Local(t.Name)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *MetadataType) Def() string {
	return "metadata"
}

// Equal reports whether t and u are of equal type.
func (t *MetadataType) Equal(u Type) bool {
	_, ok := u.(*MetadataType)
	return ok
}

// GetName returns the name of the type.
func (t *MetadataType) GetName() string {
	return t.Name
}

// SetName sets the name of the type.
func (t *MetadataType) SetName(name string) {
	t.Name = name
}
