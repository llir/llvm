package lexer_test

import (
	"fmt"
	"log"

	"github.com/mewlang/llvm/asm/lexer"
)

func ExampleParse() {
	tokens, err := lexer.ParseFile("../testdata/c4.ll")
	if err != nil {
		log.Fatalln(err)
	}
	for i, tok := range tokens {
		if i > 10 {
			break
		}
		fmt.Printf("%-12v %q\n", tok.Kind, tok.Val)
	}
	// Output:
	// Comment      " ModuleID = 'c4.ll'"
	// KwTarget     "target"
	// KwDatalayout "datalayout"
	// Equal        "="
	// String       "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
	// KwTarget     "target"
	// KwTriple     "triple"
	// Equal        "="
	// String       "x86_64-unknown-linux-gnu"
	// GlobalVar    "p"
	// Equal        "="
}

func ExampleParseString() {
	tokens := lexer.ParseString("%bar = icmp sge i32 %foo, 65 ; bar = (foo >= 'A')")
	for i, tok := range tokens {
		fmt.Printf("=== [ token %d ] ===\n", i)
		fmt.Println("pos: ", tok.Pos)
		fmt.Println("kind:", tok.Kind)
		fmt.Println("val: ", tok.Val)
		fmt.Println()
	}
	// Output:
	// === [ token 0 ] ===
	// pos:  0
	// kind: LocalVar
	// val:  bar
	//
	// === [ token 1 ] ===
	// pos:  5
	// kind: Equal
	// val:  =
	//
	// === [ token 2 ] ===
	// pos:  7
	// kind: KwIcmp
	// val:  icmp
	//
	// === [ token 3 ] ===
	// pos:  12
	// kind: KwSge
	// val:  sge
	//
	// === [ token 4 ] ===
	// pos:  16
	// kind: Type
	// val:  i32
	//
	// === [ token 5 ] ===
	// pos:  20
	// kind: LocalVar
	// val:  foo
	//
	// === [ token 6 ] ===
	// pos:  24
	// kind: Comma
	// val:  ,
	//
	// === [ token 7 ] ===
	// pos:  26
	// kind: Int
	// val:  65
	//
	// === [ token 8 ] ===
	// pos:  29
	// kind: Comment
	// val:   bar = (foo >= 'A')
	//
	// === [ token 9 ] ===
	// pos:  49
	// kind: EOF
	// val:
}
