package ir

import (
	"testing"

	"github.com/llir/llvm/ir/types"
)

func TestSample(t *testing.T) {
	x := &LocalIdent2{Name: "X", Type: types.I32}

	inst := &InstAdd2{
		LocalIdent2: LocalIdent2{Name: "X", Type: types.I32},
		X:           x,
		Y:           x,
	}

	inst.Name = "Y"

	println(x.Name)
	println(inst.Name)
}
