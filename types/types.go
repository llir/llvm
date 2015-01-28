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
	// isType ensures that only types can be assigned to the Type interface.
	isType()
}

// IsInt returns true if typ is an integer type, and false otherwise.
func IsInt(typ Type) bool {
	_, ok := typ.(*Int)
	return ok
}

// IsInts returns true if typ is an integer type or a vector of integers type,
// and false otherwise.
func IsInts(typ Type) bool {
	if t, ok := typ.(*Vector); ok {
		return IsInts(t.Elem())
	}
	return IsInt(typ)
}

// IsFloat returns true if typ is a floating point type, and false otherwise.
func IsFloat(typ Type) bool {
	_, ok := typ.(*Float)
	return ok
}

// IsFloats returns true if typ is a floating point type or a vector of floating
// points type, and false otherwise.
func IsFloats(typ Type) bool {
	if t, ok := typ.(*Vector); ok {
		return IsFloats(t.Elem())
	}
	return IsFloat(typ)
}
