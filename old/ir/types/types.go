// Package types declares the data types of LLVM IR.
package types

import (
	"fmt"
	"log"

	"github.com/mewkiz/pkg/errutil"
)

// A Type represents an LLVM IR type.
//
// Type is one of the following concrete types:
//    *types.Void
//    *types.Int
//    *types.Float
//    *types.MMX
//    *types.Label
//    *types.Metadata
//    *types.Func
//    *types.Pointer
//    *types.Vector
//    *types.Array
//    *types.Struct
//
// References:
//    http://llvm.org/docs/LangRef.html#typesystem
type Type interface {
	fmt.Stringer
	// Equal reports whether t and u are of equal type.
	Equal(u Type) bool
}

// Equal reports whether t and u are of equal type.
func Equal(t, u Type) bool {
	return t.Equal(u)
}

// IsVoid reports whether t is of void type.
func IsVoid(t Type) bool {
	_, ok := t.(*Void)
	return ok
}

// IsBool reports whether t is of boolean type (i.e. integer of size 1).
func IsBool(t Type) bool {
	return t.Equal(I1)
}

// IsBools reports whether t is of boolean or boolean vector type.
func IsBools(t Type) bool {
	if t, ok := t.(*Vector); ok {
		return IsBool(t.Elem())
	}
	return IsBool(t)
}

// IsInt reports whether t is of integer type.
func IsInt(t Type) bool {
	_, ok := t.(*Int)
	return ok
}

// IsInts reports whether t is of integer or integer vector type.
func IsInts(t Type) bool {
	if t, ok := t.(*Vector); ok {
		return IsInt(t.Elem())
	}
	return IsInt(t)
}

// IsFloat reports whether t is of floating point type.
func IsFloat(t Type) bool {
	_, ok := t.(*Float)
	return ok
}

// IsFloats reports whether t is of floating point or floating point vector
// type.
func IsFloats(t Type) bool {
	if t, ok := t.(*Vector); ok {
		return IsFloat(t.Elem())
	}
	return IsFloat(t)
}

// IsLabel reports whether t is of label type.
func IsLabel(t Type) bool {
	_, ok := t.(*Label)
	return ok
}

// IsPointer reports whether t is of pointer type.
func IsPointer(t Type) bool {
	_, ok := t.(*Pointer)
	return ok
}

// IsPointers reports whether t is of pointer or pointer vector type.
func IsPointers(t Type) bool {
	if t, ok := t.(*Vector); ok {
		return IsPointer(t.Elem())
	}
	return IsPointer(t)
}

// IsArray reports whether t is of array type.
func IsArray(t Type) bool {
	_, ok := t.(*Array)
	return ok
}

// SameLength reports whether t and u are both vectors or arrays of the same
// length or both distinct from vectors and arrays.
func SameLength(a, b Type) bool {
	type Lener interface {
		Len() int
	}
	l1, ok1 := a.(Lener)
	l2, ok2 := b.(Lener)

	// Verify if both types are vectors or arrays of the same length.
	if ok1 && ok2 {
		return l1.Len() == l2.Len()
	}

	// Verify if both types are distinct from vectors and arrays.
	return !ok1 && !ok2
}

// Convenience types.
var (
	// I1 represents the i1 type.
	I1 *Int
	// I8 represents the i8 type.
	I8 *Int
	// I32 represents the i32 type.
	I32 *Int
	// I64 represents the i64 type.
	I64 *Int
)

func init() {
	var err error
	I1, err = NewInt(1)
	if err != nil {
		log.Fatal(errutil.Err(err))
	}
	I8, err = NewInt(8)
	if err != nil {
		log.Fatal(errutil.Err(err))
	}
	I32, err = NewInt(32)
	if err != nil {
		log.Fatal(errutil.Err(err))
	}
	I64, err = NewInt(64)
	if err != nil {
		log.Fatal(errutil.Err(err))
	}
}
