// === [ Simple constants ] ====================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#simple-constants

package constant

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/llir/llvm/ir/types"
)

// --- [ integer ] -------------------------------------------------------------

// Int represents an integer constant.
type Int struct {
	// Integer type.
	Typ *types.IntType
	// Integer value.
	X *big.Int
}

// NewInt returns a new integer constant based on the given integer value and
// type.
func NewInt(x int64, typ types.Type) *Int {
	t, ok := typ.(*types.IntType)
	if !ok {
		panic(fmt.Errorf("invalid integer constant type; expected *types.IntType, got %T", typ))
	}
	return &Int{Typ: t, X: big.NewInt(x)}
}

// NewIntFromString returns a new integer constant based on the given integer
// string and type.
func NewIntFromString(s string, typ types.Type) *Int {
	// Parse boolean integer constants.
	c := NewInt(0, typ)
	if types.IsBool(c.Typ) {
		switch s {
		case "0", "false":
			c.X.SetInt64(0)
		case "1", "true":
			c.X.SetInt64(1)
		default:
			panic(fmt.Errorf("invalid integer constant %q for type i1", s))
		}
		return c
	} else if s == "true" || s == "false" {
		panic(fmt.Errorf("invalid integer constant %q for type %s", s, typ))
	}

	// Parse decimal or hexadecimal integer constant.
	base := 10
	// Hexadecimal integer format
	//
	//    [us]0x[0-9A-Fa-f]+
	switch {
	case strings.HasPrefix(s, "u0x"), strings.HasPrefix(s, "s0x"):
		s = s[len("u0x"):]
		base = 16
	case strings.HasPrefix(s, "0x"):
		s = s[len("0x"):]
		base = 16
	}
	if _, ok := c.X.SetString(s, base); !ok {
		panic(fmt.Errorf("unable to parse constant %q of type %s", s, typ))
	}
	return c
}

// Type returns the type of the constant.
func (c *Int) Type() types.Type {
	return c.Typ
}

// Ident returns the string representation of the constant.
func (c *Int) Ident() string {
	if c.Typ.Size == 1 {
		switch c.Int64() {
		case 0:
			return "false"
		case 1:
			return "true"
		}
	}
	return c.X.String()
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Int) Immutable() {}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*Int) MetadataNode() {}

// Int64 returns the int64 representation of the integer constant.
func (c *Int) Int64() int64 {
	return c.X.Int64()
}

// --- [ null pointer ] --------------------------------------------------------

// Null represents a null pointer constant.
type Null struct {
	// Pointer type.
	Typ *types.PointerType
}

// NewNull returns a new null pointer constant based on the given pointer type.
func NewNull(typ types.Type) *Null {
	t, ok := typ.(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid null pointer constant type; expected *types.PointerType, got %T", typ))
	}
	return &Null{Typ: t}
}

// Type returns the type of the constant.
func (c *Null) Type() types.Type {
	return c.Typ
}

// Ident returns the string representation of the constant.
func (c *Null) Ident() string {
	return "null"
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Null) Immutable() {}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*Null) MetadataNode() {}
