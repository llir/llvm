package types

type Type interface {
	Equal(u Type) bool
}
