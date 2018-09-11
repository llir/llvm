package constant

import "github.com/llir/l/ir/types"

// Struct is a struct constant.
type Struct struct {
	// Struct fields.
	Fields []Constant
}

// NewStruct returns a new struct constant based on the given struct fields.
func NewStruct(fields ...Constant) *Struct {
	return &Struct{Fields: fields}
}

// Type returns the type of the constant.
func (c *Struct) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant.
func (c *Struct) Ident() string {
	panic("not yet implemented")
}
