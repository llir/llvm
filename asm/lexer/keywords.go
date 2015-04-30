package lexer

import (
	"sort"

	"github.com/llir/llvm/asm/token"
)

// keywords is the set of valid keywords in LLVM IR
var keywords []string

func init() {
	keywords = make([]string, 0, len(token.Keywords))
	for keyword := range token.Keywords {
		keywords = append(keywords, keyword)
	}
	sort.Strings(keywords)
}
