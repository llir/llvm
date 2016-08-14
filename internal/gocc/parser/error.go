package parser

import (
	"fmt"
	"sort"

	"github.com/llir/spec/gocc/errors"
)

// NewError returns a user-friendly parse error.
func NewError(err *errors.Error) error {
	if err.Err != nil {
		return err.Err
	}
	// TODO: Add line:col positional tracking.
	var expected []string
	for _, tok := range err.ExpectedTokens {
		if tok == "error" {
			// Remove "error" production rule from the set of expected tokens.
			continue
		}
		expected = append(expected, tok)
	}
	sort.Strings(expected)
	return fmt.Errorf("%d: unexpected %q, expected %q", err.ErrorToken.Pos.Offset, string(err.ErrorToken.Lit), expected)
}
