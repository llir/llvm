// === [ Other types ] =========================================================

package types

import (
	"bytes"
	"fmt"
)

// --- [ void ] ----------------------------------------------------------------

// VoidType represents a void type.
//
// References:
//    http://llvm.org/docs/LangRef.html#void-type
type VoidType struct {
}

// String returns the LLVM syntax representation of the type.
func (t *VoidType) String() string {
	return "void"
}

// Equal reports whether t and u are of equal type.
func (t *VoidType) Equal(u Type) bool {
	_, ok := u.(*VoidType)
	return ok
}

// --- [ function ] ------------------------------------------------------------

// FuncType represents a function type.
//
// References:
//    http://llvm.org/docs/LangRef.html#function-type
type FuncType struct {
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

// NewParam appends a new function parameter to the function type based on the
// given parameter name and type.
func (t *FuncType) NewParam(name string, typ Type) *Param {
	param := NewParam(name, typ)
	t.Params = append(t.Params, param)
	return param
}

// A Param represents an LLVM IR function parameter.
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

// --- [ label ] ---------------------------------------------------------------

// LabelType represents a label type, which is used for basic block values.
//
// References:
//    http://llvm.org/docs/LangRef.html#label-type
type LabelType struct {
}

// String returns the LLVM syntax representation of the type.
func (t *LabelType) String() string {
	return "label"
}

// Equal reports whether t and u are of equal type.
func (t *LabelType) Equal(u Type) bool {
	_, ok := u.(*LabelType)
	return ok
}

// --- [ metadata ] ------------------------------------------------------------

// MetadataType represents a metadata type.
//
// References:
//    http://llvm.org/docs/LangRef.html#metadata-type
type MetadataType struct {
}

// String returns the LLVM syntax representation of the type.
func (t *MetadataType) String() string {
	return "metadata"
}

// Equal reports whether t and u are of equal type.
func (t *MetadataType) Equal(u Type) bool {
	_, ok := u.(*MetadataType)
	return ok
}
