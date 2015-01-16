package lexer

import (
	"fmt"
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
				{Kind: token.GlobalVar, Val: "foo", Line: 1, Col: 1},
				{Kind: token.LocalVar, Val: "bar", Line: 1, Col: 5},
				{Kind: token.ComdatVar, Val: "baz", Line: 1, Col: 9},
				{Kind: token.MetadataVar, Val: "qux", Line: 1, Col: 13},
				{Kind: token.GlobalID, Val: "42", Line: 1, Col: 17},
				{Kind: token.LocalID, Val: "37", Line: 1, Col: 20},
				{Kind: token.AttrID, Val: "7", Line: 1, Col: 23},
				{Kind: token.EOF, Line: 1, Col: 25},
			},
		},
		// i=3
		{
			input: "...=,*[]{}()<>!",
			want: []token.Token{
				{Kind: token.Ellipsis, Val: "...", Line: 1, Col: 1},
				{Kind: token.Equal, Val: "=", Line: 1, Col: 4},
				{Kind: token.Comma, Val: ",", Line: 1, Col: 5},
				{Kind: token.Star, Val: "*", Line: 1, Col: 6},
				{Kind: token.Lbrack, Val: "[", Line: 1, Col: 7},
				{Kind: token.Rbrack, Val: "]", Line: 1, Col: 8},
				{Kind: token.Lbrace, Val: "{", Line: 1, Col: 9},
				{Kind: token.Rbrace, Val: "}", Line: 1, Col: 10},
				{Kind: token.Lparen, Val: "(", Line: 1, Col: 11},
				{Kind: token.Rparen, Val: ")", Line: 1, Col: 12},
				{Kind: token.Less, Val: "<", Line: 1, Col: 13},
				{Kind: token.Greater, Val: ">", Line: 1, Col: 14},
				{Kind: token.Exclaim, Val: "!", Line: 1, Col: 15},
				{Kind: token.EOF, Line: 1, Col: 16},
			},
		},
		// i=4
		{
			input: `"fo\6F":"fo\6F"@"fo\6F"%"fo\6F"$"fo\6F"!fo\6F`,
			want: []token.Token{
				{Kind: token.Label, Val: "foo", Line: 1, Col: 1},
				{Kind: token.String, Val: "foo", Line: 1, Col: 9},
				{Kind: token.GlobalVar, Val: "foo", Line: 1, Col: 16},
				{Kind: token.LocalVar, Val: "foo", Line: 1, Col: 24},
				{Kind: token.ComdatVar, Val: "foo", Line: 1, Col: 32},
				{Kind: token.MetadataVar, Val: "foo", Line: 1, Col: 40},
				{Kind: token.EOF, Line: 1, Col: 46},
			},
		},
		// i=5
		{
			input: "!42.0foo:;foo",
			want: []token.Token{
				{Kind: token.Exclaim, Val: "!", Line: 1, Col: 1},
				{Kind: token.Float, Val: "42.0", Line: 1, Col: 2},
				{Kind: token.Label, Val: "foo", Line: 1, Col: 6},
				{Kind: token.Comment, Val: "foo", Line: 1, Col: 10},
				{Kind: token.EOF, Line: 1, Col: 14},
			},
		},
		// i=6
		{
			input: "i42floatvoidaddxu0x6F",
			want: []token.Token{
				{Kind: token.Type, Val: "i42", Line: 1, Col: 1},
				{Kind: token.Type, Val: "float", Line: 1, Col: 4},
				{Kind: token.Type, Val: "void", Line: 1, Col: 9},
				{Kind: token.KwAdd, Val: "add", Line: 1, Col: 13},
				{Kind: token.KwX, Val: "x", Line: 1, Col: 16},
				{Kind: token.Int, Val: "u0x6F", Line: 1, Col: 17},
				{Kind: token.EOF, Line: 1, Col: 22},
			},
		},
		// i=7
		{
			input: "i42floatvoidaddxu0x6F:",
			want: []token.Token{
				{Kind: token.Label, Val: "i42floatvoidaddxu0x6F", Line: 1, Col: 1},
				{Kind: token.EOF, Line: 1, Col: 23},
			},
		},
		// i=8
		{
			input: "42:-foo:0x1e",
			want: []token.Token{
				{Kind: token.Label, Val: "42", Line: 1, Col: 1},
				{Kind: token.Label, Val: "-foo", Line: 1, Col: 4},
				{Kind: token.Float, Val: "0x1e", Line: 1, Col: 9},
				{Kind: token.EOF, Line: 1, Col: 13},
			},
		},
		// i=9
		{
			input: "0xK1e 0xL1e 0xM1e 0xH1e",
			want: []token.Token{
				{Kind: token.Float, Val: "0xK1e", Line: 1, Col: 1},
				{Kind: token.Float, Val: "0xL1e", Line: 1, Col: 7},
				{Kind: token.Float, Val: "0xM1e", Line: 1, Col: 13},
				{Kind: token.Float, Val: "0xH1e", Line: 1, Col: 19},
				{Kind: token.EOF, Line: 1, Col: 24},
			},
		},
		// i=10
		{
			input: "37-42",
			want: []token.Token{
				{Kind: token.Int, Val: "37", Line: 1, Col: 1},
				{Kind: token.Int, Val: "-42", Line: 1, Col: 7},
				{Kind: token.EOF, Line: 1, Col: 24},
			},
		},
		// i=11
		{
			input: "....foo:",
			want: []token.Token{
				{Kind: token.Ellipsis, Val: "...", Line: 1, Col: 1},
				{Kind: token.Label, Val: ".foo", Line: 1, Col: 4},
				{Kind: token.EOF, Line: 1, Col: 9},
			},
		},
	}
	for i, g := range golden {
		got := Parse(g.input)
		if !reflect.DeepEqual(got, g.want) {
			t.Errorf("i=%d: expected %#v, got %#v", i, g.want, got)
			continue
		}
		fmt.Println("PASS:", i) // TODO: Remove when all test cases passes.
	}
}
