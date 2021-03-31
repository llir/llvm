package ir

import (
	"github.com/llir/llvm/ir/types"
)

type Value interface {
	isValue()
}

type LocalIdent2 struct {
	Value
	Name string
	types.Type
}

type Inst interface {
	isInst()
}

type InstAdd2 struct {
	LocalIdent2
	X, Y Value
}
