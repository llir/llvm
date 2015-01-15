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
		// i=0
		{
			input: ",",
			want: []token.Token{
				{Kind: token.Comma, Val: ",", Line: 1, Col: 1},
				{Kind: token.EOF, Line: 1, Col: 2},
			},
		},
		// i=1
		{
			input: "+0.314e+1",
			want: []token.Token{
				{Kind: token.Float, Val: "+0.314e+1", Line: 1, Col: 1},
				{Kind: token.EOF, Line: 1, Col: 10},
			},
		},
		// i=2
		{
			input: "@foo%bar$baz!qux@42%37#7",
			want: []token.Token{
				{Kind: token.GlobalVar, Val: "@foo", Line: 1, Col: 1},
				{Kind: token.LocalVar, Val: "%bar", Line: 1, Col: 5},
				{Kind: token.ComdatVar, Val: "$baz", Line: 1, Col: 9},
				{Kind: token.MetadataVar, Val: "!qux", Line: 1, Col: 13},
				{Kind: token.GlobalID, Val: "@42", Line: 1, Col: 17},
				{Kind: token.LocalID, Val: "%37", Line: 1, Col: 20},
				{Kind: token.AttrID, Val: "#7", Line: 1, Col: 23},
				{Kind: token.EOF, Line: 1, Col: 25},
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
