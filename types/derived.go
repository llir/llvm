package types

import (
	"bytes"
	"errors"
	"fmt"
)

// Func represents a function type.
//
// Examples:
//     i32 (i8*, ...)   ; Function signature of printf.
//     void ()
//
// References:
//    http://llvm.org/docs/LangRef.html#function-type
type Func struct {
	// Result parameter type.
	result Type
	// Function parameter types.
	params []Type
	// Specifies if the function takes a variable number of arguments or not.
	variadic bool
}

// NewFunc returns a function type based on the given result parameter type and
// function parameter types. The function takes a variable number of arguments
// if variadic is true.
func NewFunc(result Type, params []Type, variadic bool) (*Func, error) {
	// Validate result parameter type (any type except label, metadata and
	// function).
	switch result.(type) {
	case *Void, *Int, *Float, *MMX, *Pointer, *Vector, *Array, *Struct:
		// valid type
	default:
		return nil, fmt.Errorf("invalid result parameter type %q", result)
	}

	// Validate function parameter types (any type except void and function).
	for _, param := range params {
		switch param.(type) {
		case *Int, *Float, *MMX, *Label, *Metadata, *Pointer, *Vector, *Array, *Struct:
			// valid type
		case *Void:
			return nil, errors.New("invalid function parameter type; void type only allowed for function results")
		default:
			return nil, fmt.Errorf("invalid function parameter type %q", param)
		}
	}

	return &Func{result: result, params: params, variadic: variadic}, nil
}

// Result returns the function result type.
func (t *Func) Result() Type {
	return t.result
}

// Params returns the function parameter types.
func (t *Func) Params() []Type {
	return t.params
}

// IsVariadic returns true if the function takes a variable number of arguments,
// and false otherwise.
func (t *Func) IsVariadic() bool {
	return t.variadic
}

// Equal returns true if the given types are equal, and false otherwise.
func (t *Func) Equal(u Type) bool {
	if u, ok := u.(*Func); ok {
		if !t.result.Equal(u.result) {
			return false
		}
		if len(t.params) != len(u.params) {
			return false
		}
		for i := range t.params {
			if !t.params[i].Equal(u.params[i]) {
				return false
			}
		}
		return t.variadic == u.variadic
	}
	return false
}

// String returns a string representation of the function type.
func (t *Func) String() string {
	// e.g. "void ()"
	// e.g. "i32 (i8*, ...)"
	params := t.Params()
	buf := new(bytes.Buffer)
	for i, param := range params {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(param.String())
	}
	if t.IsVariadic() {
		if len(params) > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("...")
	}

	return fmt.Sprintf("%s (%s)", t.Result(), buf)
}

// Pointer represents a pointer type.
//
// Examples:
//    int32*
//
// References:
//    http://llvm.org/docs/LangRef.html#pointer-type
type Pointer struct {
	// Element type.
	elem Type
}

// NewPointer returns a pointer type for the given element type.
func NewPointer(elem Type) (*Pointer, error) {
	// Validate element type (any type except void, label and metadata).
	switch elem.(type) {
	case *Int, *Float, *MMX, *Func, *Pointer, *Vector, *Array, *Struct:
		// valid type
	case *Void:
		return nil, errors.New(`invalid pointer to "void"; use i8* instead`)
	default:
		return nil, fmt.Errorf("invalid pointer to %q", elem)
	}

	return &Pointer{elem: elem}, nil
}

// Elem returns the element type of the pointer.
func (t *Pointer) Elem() Type {
	return t.elem
}

// Equal returns true if the given types are equal, and false otherwise.
func (t *Pointer) Equal(u Type) bool {
	if u, ok := u.(*Pointer); ok {
		return t.elem.Equal(u.elem)
	}
	return false
}

// String returns a string representation of the pointer type.
func (t *Pointer) String() string {
	// e.g. "i32*"
	return t.Elem().String() + "*"
}

// Vector represents a vector type.
//
// Examples:
//    <10 x i32>   ; A vector of 10 32-bit integers.
//
// References:
//    http://llvm.org/docs/LangRef.html#vector-type
type Vector struct {
	// Element type.
	elem Type
	// Number of elements.
	n int
}

// NewVector returns a vector type based on the specified element type and
// length.
func NewVector(elem Type, n int) (*Vector, error) {
	// Validate element type (any type except void, x86_mmx, label, metadata,
	// function, vector, array and struct).
	switch elem.(type) {
	case *Int, *Float, *Pointer:
		// valid type
	case *Void:
		return nil, errors.New("invalid vector element type; void type only allowed for function results")
	default:
		return nil, fmt.Errorf("invalid vector element type %q", elem)
	}

	// Validate vector length.
	if n < 1 {
		return nil, fmt.Errorf("invalid vector length (%d)", n)
	}

	return &Vector{elem: elem, n: n}, nil
}

// Elem returns the element type of the vector.
func (t *Vector) Elem() Type {
	return t.elem
}

// Len returns the length of the vector in number of elements.
func (t *Vector) Len() int {
	return t.n
}

// Equal returns true if the given types are equal, and false otherwise.
func (t *Vector) Equal(u Type) bool {
	if u, ok := u.(*Vector); ok {
		return t.elem.Equal(u.elem) && t.n == u.n
	}
	return false
}

// String returns a string representation of the vector type.
func (t *Vector) String() string {
	// e.g. "<2 x i32>"
	return fmt.Sprintf("<%d x %v>", t.Len(), t.Elem())
}

// Array represents an array type.
//
// Examples:
//    [10 x i32]   ; An array of 10 32-bit integers.
//
// References:
//    http://llvm.org/docs/LangRef.html#array-type
type Array struct {
	// Element type.
	elem Type
	// Number of elements.
	n int
}

// NewArray returns an array type based on the specified element type and
// length.
func NewArray(elem Type, n int) (*Array, error) {
	// Validate element type (any type except void, label, metadata and
	// function).
	switch elem.(type) {
	case *Int, *Float, *MMX, *Pointer, *Vector, *Array, *Struct:
		// valid type
	case *Void:
		return nil, errors.New("invalid array element type; void type only allowed for function results")
	default:
		return nil, fmt.Errorf("invalid array element type %q", elem)
	}

	// Validate array length.
	if n < 0 {
		return nil, fmt.Errorf("invalid array length (%d)", n)
	}

	return &Array{elem: elem, n: n}, nil
}

// Elem returns the element type of the array.
func (t *Array) Elem() Type {
	return t.elem
}

// Len returns the length of the array in number of elements.
func (t *Array) Len() int {
	return t.n
}

// Equal returns true if the given types are equal, and false otherwise.
func (t *Array) Equal(u Type) bool {
	if u, ok := u.(*Array); ok {
		return t.elem.Equal(u.elem) && t.n == u.n
	}
	return false
}

// String returns a string representation of the array type.
func (t *Array) String() string {
	// e.g. "[2 x float]"
	return fmt.Sprintf("[%d x %v]", t.Len(), t.Elem())
}

// Struct represents a literal structure type.
//
// Examples:
//    {float, i32, i32}   ; Normal structure (padding depends on datalayout).
//    <{i32 i8}>          ; Packed structure (5 bytes in size).
//
// References:
//    http://llvm.org/docs/LangRef.html#structure-type
type Struct struct {
	// Structure fields.
	fields []Type
	// Packed structures use 1 byte alignment.
	packed bool
}

// NewStruct returns a structure type based on the given field types. The
// structure is 1 byte aligned if packed is true.
func NewStruct(fields []Type, packed bool) (*Struct, error) {
	// Validate field types (any type except void, label, metadata and function).
	for _, field := range fields {
		switch field.(type) {
		case *Int, *Float, *MMX, *Pointer, *Vector, *Array, *Struct:
			// valid type
		case *Void:
			return nil, errors.New("invalid structure field type; void type only allowed for function results")
		default:
			return nil, fmt.Errorf("invalid structure field type %q", field)
		}
	}

	return &Struct{fields, packed}, nil
}

// Fields returns the field types of the structure.
func (t *Struct) Fields() []Type {
	return t.fields
}

// IsPacked returns true if the structure is 1 byte aligned.
func (t *Struct) IsPacked() bool {
	return t.packed
}

// Equal returns true if the given types are equal, and false otherwise.
//
// Literal structure types are uniqued by structural identity.
func (t *Struct) Equal(u Type) bool {
	if u, ok := u.(*Struct); ok {
		if len(t.fields) != len(u.fields) {
			return false
		}
		for i := range t.fields {
			if !t.fields[i].Equal(u.fields[i]) {
				return false
			}
		}
		return t.packed == u.packed
	}
	return false
}

// String returns a string representation of the literal structure type.
func (t *Struct) String() string {
	// e.g. "{float, i32, i32}"
	// e.g. "<{i32, i8}>"
	buf := new(bytes.Buffer)
	for i, field := range t.Fields() {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(field.String())
	}

	if t.IsPacked() {
		return fmt.Sprintf("<{%s}>", buf)
	}
	return fmt.Sprintf("{%s}", buf)
}
