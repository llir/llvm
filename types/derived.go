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
func (typ *Func) Result() Type {
	return typ.result
}

// Params returns the function parameter types.
func (typ *Func) Params() []Type {
	return typ.params
}

// IsVariadic returns true if the function takes a variable number of arguments,
// and false otherwise.
func (typ *Func) IsVariadic() bool {
	return typ.variadic
}

func (typ *Func) String() string {
	// i32 (i8*, ...)
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v (", typ.result)
	for i, param := range typ.params {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(param.String())
	}
	if typ.variadic {
		if len(typ.params) > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("...")
	}
	buf.WriteString(")")
	return buf.String()
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
func (typ *Pointer) Elem() Type {
	return typ.elem
}

func (typ *Pointer) String() string {
	return fmt.Sprintf("%v*", typ.elem)
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
func (typ Vector) Elem() Type {
	return typ.elem
}

// Len returns the length of the vector in number of elements.
func (typ Vector) Len() int {
	return typ.n
}

func (typ *Vector) String() string {
	return fmt.Sprintf("<%d x %v>", typ.n, typ.elem)
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
func (typ Array) Elem() Type {
	return typ.elem
}

// Len returns the length of the array in number of elements.
func (typ Array) Len() int {
	return typ.n
}

func (typ *Array) String() string {
	return fmt.Sprintf("[%d x %v]", typ.n, typ.elem)
}

// Struct represents a structure type.
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
func (typ *Struct) Fields() []Type {
	return typ.fields
}

// IsPacked returns true if the structure is 1 byte aligned.
func (typ *Struct) IsPacked() bool {
	return typ.packed
}

func (typ *Struct) String() string {
	buf := new(bytes.Buffer)
	if typ.packed {
		buf.WriteString("<")
	}
	buf.WriteString("{")
	for i, field := range typ.fields {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(field.String())
	}
	buf.WriteString("}")
	if typ.packed {
		buf.WriteString(">")
	}
	return buf.String()
}

// isType ensures that only types can be assigned to the Type interface.
func (*Func) isType()    {}
func (*Pointer) isType() {}
func (*Vector) isType()  {}
func (*Array) isType()   {}
func (*Struct) isType()  {}
