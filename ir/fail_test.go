package ir

import (
	"fmt"
	"strings"
	"testing"

	"github.com/llir/llvm/ir/types"
)

func TestFailSituation(t *testing.T) {
	golden := []struct {
		do           func()
		errorMessage string
	}{
		{
			do: func() {
				mod := NewModule()
				f := mod.NewFunc("test", types.Void)
				f.NewBlock("")
				mod.String()
			},
			errorMessage: "block do not have terminator",
		},
	}
	for _, g := range golden {
		defer func() {
			if r := recover(); r != nil {
				if !strings.Contains(fmt.Sprintf("%s", r), g.errorMessage) {
					t.Errorf("expected fail with `%s` but got: `%s`", g.errorMessage, r)
				}
			} else if r == nil {
				t.Errorf("expected fail with `%s` but don't", g.errorMessage)
			}
		}()
		g.do()
	}

}
