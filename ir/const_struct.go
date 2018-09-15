package ir

import (
	"fmt"

	"github.com/llir/l/ir/types"
)

// --- [ Struct constants ] ----------------------------------------------------

// ConstStruct is an LLVM IR struct constant.
type ConstStruct struct {
	// Struct fields.
	Fields []Constant
}

// NewStruct returns a new struct constant based on the given struct fields.
func NewStruct(fields ...Constant) *ConstStruct {
	return &ConstStruct{Fields: fields}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstStruct) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *ConstStruct) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant.
func (c *ConstStruct) Ident() string {
	panic("not yet implemented")
}
