package lexer

import (
	"reflect"
	"testing"

	"github.com/mewlang/llvm/asm/token"
)

func TestParse(t *testing.T) {
	golden := []struct {
		input string
		want  []token.Token
	}{
		{
			input: ",",
			want: []token.Token{
				{Kind: token.Comma, Val: ",", Line: 1, Col: 1},
				{Kind: token.EOF, Line: 1, Col: 2},
			},
		},
	}
	for i, g := range golden {
		got := Parse(g.input)
		if !reflect.DeepEqual(got, g.want) {
			t.Errorf("i=%d: expected %#v, got %#v", i, g.want, got)
		}
	}
}
