// TODO: Add support for named types. Implement type equality checks and make
// sure to consider the impact of named types.

// Package types declares the data types of LLVM IR.
package types

import "fmt"

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
	// Equal returns true if the given types are equal, and false otherwise.
	Equal(b Type) bool
}

// Equal returns true if the given types are equal, and false otherwise.
func Equal(t, u Type) bool {
	// TODO: Implement type equality checks for named types.
	return t.Equal(u)
}

// IsInt returns true if t is an integer type, and false otherwise.
func IsInt(t Type) bool {
	_, ok := t.(*Int)
	return ok
}

// IsInts returns true if t is an integer type or a vector of integers type, and
// false otherwise.
func IsInts(t Type) bool {
	if t, ok := t.(*Vector); ok {
		return IsInt(t.Elem())
	}
	return IsInt(t)
}

// IsFloat returns true if t is a floating point type, and false otherwise.
func IsFloat(t Type) bool {
	_, ok := t.(*Float)
	return ok
}

// IsFloats returns true if t is a floating point type or a vector of floating
// points type, and false otherwise.
func IsFloats(t Type) bool {
	if t, ok := t.(*Vector); ok {
		return IsFloat(t.Elem())
	}
	return IsFloat(t)
}

// SameLength returns true if both types are vectors or arrays of the same
// length or if both types are distinct from vectors and arrays, and false
// otherwise.
func SameLength(a, b Type) bool {
	type Lener interface {
		Len() int
	}
	l1, ok1 := a.(Lener)
	l2, ok2 := b.(Lener)

	// Both types are vectors or arrays; verify length.
	if ok1 && ok2 {
		return l1.Len() == l2.Len()
	}

	// Verify that both types are distinct from vectors and arrays.
	return !ok1 && !ok2
}
