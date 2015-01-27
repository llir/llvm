// Package types declares the data types of LLVM IR.
package types

import "fmt"

// Type represents one of the following types:
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
type Type interface {
	fmt.Stringer
	// isType ensures that only types can be assigned to the Type interface.
	isType()
}

// IsInt returns true if typ is an integer type.
func IsInt(typ Type) bool {
	_, ok := typ.(*Int)
	return ok
}

// IsInts returns true if typ is an integer type or a vector of integers type.
func IsInts(typ Type) bool {
	if t, ok := typ.(*Vector); ok {
		return IsInts(t.Elem())
	}
	return IsInt(typ)
}

// IsFloat returns true if typ is a floating point type.
func IsFloat(typ Type) bool {
	_, ok := typ.(*Float)
	return ok
}

// IsFloats returns true if typ is a floating point type or a vector of floating
// points type.
func IsFloats(typ Type) bool {
	if t, ok := typ.(*Vector); ok {
		return IsFloats(t.Elem())
	}
	return IsFloat(typ)
}
