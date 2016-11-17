package constant

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/llir/llvm/ir/internal/enc"
	"github.com/llir/llvm/ir/types"
	"github.com/mewkiz/pkg/errutil"
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
	// Vector type.
	typ *types.Vector
	// Vector elements.
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
		return nil, fmt.Errorf("invalid type %q for vector constant", typ)
	}

	// Verify vector element types.
	if len(elems) != v.typ.Len() {
		return nil, fmt.Errorf("incorrect number of elements in vector constant; expected %d, got %d", v.typ.Len(), len(elems))
	}
	for _, elem := range elems {
		got, want := elem.Type(), v.typ.Elem()
		if !got.Equal(want) {
			return nil, fmt.Errorf("invalid vector element type; expected %q, got %q", want, got)
		}
	}
	v.elems = elems

	return v, nil
}

// Type returns the type of the value.
func (v *Vector) Type() types.Type {
	return v.typ
}

// String returns a string representation of the vector; e.g.
//
//    <i32 42, i32 -13>
func (v *Vector) String() string {
	buf := new(bytes.Buffer)
	for i, elem := range v.elems {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%s %s", elem.Type(), elem)
	}
	return fmt.Sprintf("<%s>", buf)
}

// ValueString returns a string representation of the value.
func (v *Vector) ValueString() string {
	return v.String()
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
	// Array type.
	typ *types.Array
	// TODO: Be clever about data layout later (e.g. use []byte instead of
	// []Constant when applicable). Strive for correctness and simplicity first,
	// optimize later. The same goes for Vector and maybe Struct.

	// Array elements.
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
		return nil, fmt.Errorf("invalid type %q for array constant", typ)
	}

	// Verify array element types.
	if len(elems) != v.typ.Len() {
		return nil, fmt.Errorf("incorrect number of elements in array constant; expected %d, got %d", v.typ.Len(), len(elems))
	}
	for _, elem := range elems {
		got, want := elem.Type(), v.typ.Elem()
		if !got.Equal(want) {
			return nil, fmt.Errorf("invalid array element type; expected %q, got %q", want, got)
		}
	}
	v.elems = elems

	return v, nil
}

// NewCharArray returns a character array constant based on the given array type
// and string.
func NewCharArray(typ types.Type, s string) (*Array, error) {
	// Verify array type.
	v := new(Array)
	var ok bool
	v.typ, ok = typ.(*types.Array)
	if !ok {
		return nil, fmt.Errorf("invalid type %q for array constant", typ)
	}
	var err error
	s, err = unquote(s)
	if err != nil {
		return nil, errutil.Err(err)
	}

	// Verify array element types.
	if len(s) != v.typ.Len() {
		return nil, fmt.Errorf("incorrect number of elements in character array constant; expected %d, got %d", v.typ.Len(), len(s))
	}
	elemType := v.typ.Elem()
	if !types.Equal(elemType, types.I8) {
		return nil, fmt.Errorf("invalid character array element type; expected %q, got %q", types.I8, elemType)
	}
	for i := 0; i < len(s); i++ {
		elem, err := NewInt(elemType, strconv.Itoa(int(s[i])))
		if err != nil {
			return nil, errutil.Err(err)
		}
		v.elems = append(v.elems, elem)
	}

	return v, nil
}

// unquote interprets s as a double-quoted LLVM IR string literal, returning the
// string value that s quotes.
func unquote(s string) (string, error) {
	if !strings.HasPrefix(s, `"`) {
		return "", errutil.Newf(`invalid prefix of quoted string %q; expected '"'`, s)
	}
	s = s[1:]
	if !strings.HasSuffix(s, `"`) {
		return "", errutil.Newf(`invalid suffix of quoted string %q; expected '"'`, s)
	}
	s = s[:len(s)-1]
	return enc.Unescape(s), nil
}

// Type returns the type of the value.
func (v *Array) Type() types.Type {
	return v.typ
}

// String returns a string representation of the array; e.g.
//
//    [i32 42, i32 -13]
func (v *Array) String() string {
	// Pretty print character arrays; e.g.
	//    c"hello world\0A\00"
	if v.typ.Elem().Equal(types.I8) {
		// TODO: Cleanup once the array data structure has been refined.
		buf := new(bytes.Buffer)
		for _, elem := range v.elems {
			x := elem.(*Int).x.Int64()
			buf.WriteByte(byte(x))
		}
		return `c"` + escape(buf.String()) + `"`
	}

	buf := new(bytes.Buffer)
	for i, elem := range v.elems {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%s %s", elem.Type(), elem)
	}

	return fmt.Sprintf("[%s]", buf)
}

// ValueString returns a string representation of the value.
func (v *Array) ValueString() string {
	return v.String()
}

// escape replaces any characters which are not printable with corresponding
// hexadecimal escape sequence (\XX).
func escape(s string) string {
	// Check if a replacement is required.
	extra := 0
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if utf8.ValidRune(r) && unicode.IsPrint(r) {
			i += size
			continue
		}
		// Two extra bytes are required for each non-printable byte; e.g.
		//    "\n" -> `\0A`
		//    "\x00" -> `\00`
		extra += 2
		i++
	}
	if extra == 0 {
		return s
	}

	// Replace non-printable bytes.
	const hextable = "0123456789ABCDEF"
	buf := make([]byte, len(s)+extra)
	j := 0
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if utf8.ValidRune(r) && unicode.IsPrint(r) {
			for k := 0; k < size; k++ {
				buf[j+k] = s[i+k]
			}
			i += size
			j += size
			continue
		}
		b := s[i]
		buf[j] = '\\'
		buf[j+1] = hextable[b>>4]
		buf[j+2] = hextable[b&0x0F]
		i++
		j += 3
	}
	return string(buf)
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
	// Struct type.
	typ *types.Struct
	// Struct fields.
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
		return nil, fmt.Errorf("invalid type %q for structure constant", typ)
	}

	// Verify structure field types.
	fieldTypes := v.typ.Fields()
	if len(fields) != len(fieldTypes) {
		return nil, fmt.Errorf("incorrect number of fields in structure constant; expected %d, got %d", len(fieldTypes), len(fields))
	}
	for i := range fields {
		got, want := fields[i].Type(), fieldTypes[i]
		if !got.Equal(want) {
			return nil, fmt.Errorf("invalid structure field (%d) type; expected %q, got %q", i, want, got)
		}
	}
	v.fields = fields

	return v, nil
}

// Type returns the type of the value.
func (v *Struct) Type() types.Type {
	return v.typ
}

// String returns a string representation of the structure; e.g.
//
//    {i32 -13, i8 3}
func (v *Struct) String() string {
	buf := new(bytes.Buffer)
	for i, field := range v.fields {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%s %s", field.Type(), field)
	}

	return fmt.Sprintf("{%s}", buf)
}

// ValueString returns a string representation of the value.
func (v *Struct) ValueString() string {
	return v.String()
}

// isConst ensures that only constant values can be assigned to the Constant
// interface.
func (*Vector) isConst() {}
func (*Array) isConst()  {}
func (*Struct) isConst() {}
