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
	// Comment      " ModuleID = 'c4.c'"
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
