package value

import "github.com/llir/l/types"

type Value interface {
	Ident() string
	Type() types.Type
}

type NamedValue interface {
	Value
	Name() string
	SetName() string
}
