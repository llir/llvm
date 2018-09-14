package ir

import (
	"strings"
	"testing"

	"github.com/llir/l/ir/types"
)

func TestModuleString(t *testing.T) {
	golden := []struct {
		in   *Module
		want string
	}{
		// Empty module.
		{
			in:   &Module{},
			want: "",
		},
		// Type definition.
		{
			in: &Module{
				TypeDefs: []types.Type{&types.StructType{
					Alias:  "foo",
					Fields: []types.Type{types.I32},
				}},
			},
			want: "%foo = type { i32 }",
		},
	}
	for _, g := range golden {
		got := strings.TrimSpace(g.in.String())
		if g.want != got {
			t.Errorf("module mismatch; expected `%v`, got `%v`", g.want, got)
		}
	}
}
