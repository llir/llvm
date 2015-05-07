package parser

import (
	"log"

	"github.com/llir/llvm/asm/token"
)

// TODO: Remove the filter function.

// filter filters out token types which are not yet handled by the parser.
func filter(tokens []token.Token) []token.Token {
	subset := make([]token.Token, 0, len(tokens))
	for _, tok := range tokens {
		switch tok.Kind {
		case token.Comment:
			log.Printf("filter: token type %v not yet handled by the parser.", tok.Kind)
			continue
		}
		subset = append(subset, tok)
	}
	return subset
}
