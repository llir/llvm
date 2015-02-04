package consts

import (
	"bytes"
	"fmt"

	"github.com/mewkiz/pkg/errutil"
	"github.com/mewlang/llvm/types"
	"github.com/mewlang/llvm/values"
)

// Vector represents a vector constant which is a vetor containing only
// constants.
//
// Examples:
//    <i32 37, i32 42>   ; type: <2 x i32>
//
// References:
//    http://llvm.org/docs/LangRef.html#complex-constants
type Vector struct {
	typ   *types.Vector
	elems []Constant
}

// NewVector returns a vector constant based on the given vector type and vector
// elements.
func NewVector(typ types.Type, elems []Constant) (*Vector, error) {
	// Verify vector type.
	v := new(Vector)
	var ok bool
	v.typ, ok = typ.(*types.Vector)
	if !ok {
		return nil, errutil.Newf("invalid type %q for vector constant", typ)
	}

	// Verify vector element types.
	for _, elem := range elems {
		got, want := elem.Type(), v.typ.Elem()
		if !got.Equal(want) {
			return nil, errutil.Newf("invalid vector element type; expected %q, got %q", want, got)
		}
	}
	v.elems = elems

	return v, nil
}

// Type returns the type of the value.
func (v *Vector) Type() types.Type {
	return v.typ
}

// UseList returns a list of all values which uses the value.
func (v *Vector) UseList() []values.Value {
	panic("not yet implemented.")
}

// ReplaceAll replaces all uses of the value with new.
func (v *Vector) ReplaceAll(new values.Value) error {
	panic("not yet implemented.")
}

// String returns a string representation of the vector. The vector string
// representation is preceded by the type of the constant, e.g.
//
//    <2 x i32> <i32 42, i32 -13>
func (v *Vector) String() string {
	buf := new(bytes.Buffer)
	for i, elem := range v.elems {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(elem.String())
	}

	return fmt.Sprintf("%s <%s>", v.typ, buf)
}

// Array represents an array constant which is an array containing only
// constants.
//
// As a special case, character array constants may be represented as a double-
// quoted string using the c prefix.
//
// Examples:
//    [double 3.14, double 1.5]      ; type: [2 x double]
//    [<2 x i32> <i32 15, i32 20>]   ; type: [1 x <2 x i32>]
//    c"hello world\0a\00"           ; type: [13 x i8]
//
// References:
//    http://llvm.org/docs/LangRef.html#complex-constants
type Array struct {
	typ *types.Array
	// TODO: Be clever about data layout later (e.g. use []byte instead of
	// []Constant when applicable). Strive for correctness and simplicity first,
	// optimize later. The same goes for Vector and maybe Struct.
	elems []Constant
}

// NewArray returns an array constant based on the given array type and array
// elements.
func NewArray(typ types.Type, elems []Constant) (*Array, error) {
	// Verify array type.
	v := new(Array)
	var ok bool
	v.typ, ok = typ.(*types.Array)
	if !ok {
		return nil, errutil.Newf("invalid type %q for array constant", typ)
	}

	// Verify array element types.
	for _, elem := range elems {
		got, want := elem.Type(), v.typ.Elem()
		if !got.Equal(want) {
			return nil, errutil.Newf("invalid array element type; expected %q, got %q", want, got)
		}
	}
	v.elems = elems

	return v, nil
}

// Type returns the type of the value.
func (v *Array) Type() types.Type {
	return v.typ
}

// UseList returns a list of all values which uses the value.
func (v *Array) UseList() []values.Value {
	panic("not yet implemented.")
}

// ReplaceAll replaces all uses of the value with new.
func (v *Array) ReplaceAll(new values.Value) error {
	panic("not yet implemented.")
}

// String returns a string representation of the array. The array string
// representation is preceded by the type of the constant, e.g.
//
//    [2 x i32] [i32 42, i32 -13]
func (v *Array) String() string {
	buf := new(bytes.Buffer)
	for i, elem := range v.elems {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(elem.String())
	}

	return fmt.Sprintf("%s [%s]", v.typ, buf)
}

// Struct represents a structure constant which is a structure containing only
// constants.
//
// Examples:
//    {i32 7, i8 3}                    ; type: {i32, i8}
//    {i32 7, {i8, i8} {i8 3, i8 5}}   ; type: {i32, {i8, i8}}
//
// References:
//    http://llvm.org/docs/LangRef.html#complex-constants
type Struct struct {
	typ    *types.Struct
	fields []Constant
}

// NewStruct returns a structure constant based on the given structure type and
// structure fields.
func NewStruct(typ types.Type, fields []Constant) (*Struct, error) {
	// Verify structure type.
	v := new(Struct)
	var ok bool
	v.typ, ok = typ.(*types.Struct)
	if !ok {
		return nil, errutil.Newf("invalid type %q for structure constant", typ)
	}

	// Verify structure field types.
	fieldTypes := v.typ.Fields()
	if len(fields) != len(fieldTypes) {
		return nil, errutil.Newf("incorrect number of fields in structure constant; expected %d, got %d", len(fieldTypes), len(fields))
	}
	for i := range fields {
		got, want := fields[i].Type(), fieldTypes[i]
		if !got.Equal(want) {
			return nil, errutil.Newf("invalid structure field type; expected %q, got %q", want, got)
		}
	}
	v.fields = fields

	return v, nil
}

// Type returns the type of the value.
func (v *Struct) Type() types.Type {
	return v.typ
}

// UseList returns a list of all values which uses the value.
func (v *Struct) UseList() []values.Value {
	panic("not yet implemented.")
}

// ReplaceAll replaces all uses of the value with new.
func (v *Struct) ReplaceAll(new values.Value) error {
	panic("not yet implemented.")
}

// isConst ensures that only constant values can be assigned to the Constant
// interface.
func (*Vector) isConst() {}
func (*Array) isConst()  {}
func (*Struct) isConst() {}
