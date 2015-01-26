package types

// Func represents a function type.
//
// Example:
//     i32 (i8*, ...)
type Func struct {
	// Function parameter types.
	params []Type
	// Return type.
	ret Type
	// Specifies if the function takes a variadic number of arguments or not.
	variadic bool
}

// NewFunc returns a new function type based on the given return type and
// function parameter types.
func NewFunc(params []Type, ret Type, variadic bool) *Func {
	return &Func{params: params, ret: ret, variadic: variadic}
}

// Ret returns the function return type.
func (typ *Func) Ret() Type {
	return typ.ret
}

// Params returns the function parameter types.
func (typ *Func) Params() []Type {
	return typ.params
}

// Pointer represents a pointer type.
//
// Example:
//    int32*
type Pointer struct {
	// Element type.
	elem Type
}

// NewPointer returns a new pointer type for the given element type.
func NewPointer(elem Type) *Pointer {
	// TODO: Report error for pointers to void and label or change the type of
	// elem to disallow it by design.
	return &Pointer{elem: elem}
}

// Elem returns the element type of the pointer.
func (typ *Pointer) Elem() Type {
	return typ.elem
}

// Vector represents a vector type.
//
// Example:
//    <10 x i32>
type Vector struct {
	// Element type.
	elem Type
	// Number of elements.
	n int
}

// NewVector returns a new vector type based on the specified element type and
// length.
func NewVector(elem Type, n int) *Vector {
	return &Vector{elem: elem, n: n}
}

// Elem returns the element type of the vector.
func (typ Vector) Elem() Type {
	return typ.elem
}

// Len returns the length of the vector in number of elements.
func (typ Vector) Len() int {
	return typ.n
}

// Array represents an array type.
//
// Example:
//    [10 x i32]
type Array struct {
	// Element type.
	elem Type
	// Number of elements.
	n int
}

// NewArray returns a new array type based on the specified element type and
// length.
func NewArray(elem Type, n int) *Array {
	return &Array{elem: elem, n: n}
}

// Elem returns the element type of the array.
func (typ Array) Elem() Type {
	return typ.elem
}

// Len returns the length of the array in number of elements.
func (typ Array) Len() int {
	return typ.n
}

// Struct represents a structure type.
//
// Example:
//    {float, i32, i32}
type Struct struct {
	// Structure fields.
	fields []Type
	// Packed structures use 1 byte alignment.
	packed bool
}

// NewStruct returns a structure type based on the given field types. The
// structure is 1 byte aligned if packed is true.
func NewStruct(fields []Type, packed bool) *Struct {
	// TODO: Report errors for field types with no size or change to type of
	// fields to enforce it through the API.
	return &Struct{fields, packed}
}

// Fields returns the field types of the structure.
func (typ *Struct) Fields() []Type {
	return typ.fields
}

// IsPacked returns true if the structure is 1 byte aligned.
func (typ *Struct) IsPacked() bool {
	return typ.packed
}

// isType ensures that only types can be assigned to the Type interface.
func (*Func) isType()    {}
func (*Pointer) isType() {}
func (*Vector) isType()  {}
func (*Array) isType()   {}
func (*Struct) isType()  {}
