package value

import "github.com/llir/l/ir/types"

type Value interface {
	Ident() string
	Type() types.Type
}

type NamedValue interface {
	Value
	Name() string
	SetName() string
}
