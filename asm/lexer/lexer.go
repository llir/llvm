// The implementation of this package is heavily inspired by Rob Pike's amazing
// talk titled "Lexical Scanning in Go" [1].
//
// [1]: https://www.youtube.com/watch?v=HxaD_trXwRE

// Package lexer implements lexical tokenization of Go source code.
package lexer

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/mewlang/go/token"
)

// Parse lexes the input string into a slice of tokens. The underlying type of
// the returned error is ErrorList, and it contains a list of errors that
// occurred while lexing. ErrorList implements the error interface by returning
// the first error of the list from its Error method. Use type assertion to gain
// access to the entire list of errors.
func Parse(input string) (tokens []token.Token, err error) {
	l := &lexer{
		input: input,
		// The average token size of the Go standard library is approximately 5
		// bytes.
		tokens: make([]token.Token, 0, len(input)/5),
	}

	// Tokenize the input.
	l.lex()

	if len(l.errs) > 0 {
		return l.tokens, l.errs
	}
	return l.tokens, nil
}

// ErrorList is a list of errors which implements the error interface. It does
// so by returning the first error of the list from its Error method.
type ErrorList []error

// Error returns the first error of the list, or an empty string if the list is
// empty.
func (errs ErrorList) Error() string {
	if len(errs) > 0 {
		return errs[0].Error()
	}
	return ""
}

// A lexer lexes an input string into a slice of tokens. While breaking the
// input into tokens, the next token is the longest sequence of characters that
// form a valid token.
type lexer struct {
	// The input string.
	input string
	// Start position of the current token.
	start int
	// Current position in the input.
	pos int
	// Width in byte of the last rune read with next; used by backup.
	width int
	// Start line number of the current token, and current line number in the
	// input.
	startLine, line int
	// Start column number of the current token, and current and previous column
	// number in the input.
	startCol, col, prevCol int
	// A slice of scanned tokens.
	tokens []token.Token
	// Index to the first token of the current line; used by insertSemicolon.
	first int
	// A list of errors that occurred while lexing. It implements the error
	// interface by returning the first error of the list from its Error method.
	errs ErrorList
}

// lex lexes the input by repeatedly executing the active state function until
// it returns a nil state.
func (l *lexer) lex() {
	// lexToken is the initial state function of the lexer.
	for state := lexToken; state != nil; {
		state = state(l)
	}
}

// errorf appends an error to the error list.
func (l *lexer) errorf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	l.errs = append(l.errs, err)
}

// emit emits a token of the specified token type and advances the token start
// position.
func (l *lexer) emit(kind token.Kind) {
	l.emitCustom(kind, l.input[l.start:l.pos])
}

// emitCustom emits a custom token and advances the token start position.
func (l *lexer) emitCustom(kind token.Kind, val string) {
	tok := token.Token{
		Kind: kind,
		Val:  val,
		Line: l.startLine + 1,
		Col:  l.startCol + 1,
	}
	l.tokens = append(l.tokens, tok)
	l.start = l.pos
	l.startLine, l.startCol = l.line, l.col
}

const (
	// eof is the rune returned by next when no more input is available.
	eof = -1
	// bom is the UTF-8-encoded byte order mark.
	bom = '\ufeff'
	// nul is the NUL character.
	nul = '\x00'
)

// next consumes and returns the next rune of the input.
func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	switch r {
	case bom:
		// For compatibility with other tools, a compiler may ignore a
		// UTF-8-encoded byte order mark (U+FEFF) if it is the first Unicode code
		// point in the source text. A byte order mark may be disallowed anywhere
		// else in the source.
		//
		// ref: http://golang.org/ref/spec#Source_code_representation
		if l.pos == 3 {
			// Ignore a UTF-8-encoded byte order mark (U+FEFF) if it is the first
			// Unicode code point in the source text.
			l.ignore()
			return l.next()
		}
		// A byte order mark is disallowed anywhere else in the source.
		l.errorf("illegal byte order mark")
	case nul:
		// For compatibility with other tools, a compiler may disallow the NUL
		// character (U+0000) in the source text.
		l.errorf("illegal NUL character")
	case utf8.RuneError:
		l.errorf("illegal UTF-8 encoding")
	}
	// TODO(u): Find a cleaner way to handle line:column tracking. The current
	// implementation requires five different struct fields.
	if r == '\n' {
		l.line++
		l.col, l.prevCol = 0, l.col
	} else {
		l.col++
	}
	return r
}

// backup backs up one rune in the input. It can only be called once per call to
// next.
func (l *lexer) backup() {
	l.pos -= l.width
	l.width = 0
	if l.col == 0 {
		l.line--
		l.col = l.prevCol
	} else {
		l.col--
	}
}

// accept consumes the next rune if it's from the valid set. It returns true if
// a rune was consumed and false otherwise.
func (l *lexer) accept(valid string) bool {
	r := l.next()
	if r == eof {
		return false
	}
	if strings.IndexRune(valid, r) == -1 {
		l.backup()
		return false
	}
	return true
}

// acceptRun consumes a run of runes from the valid set. It returns true if a
// rune was consumed and false otherwise.
func (l *lexer) acceptRun(valid string) bool {
	consumed := false
	for l.accept(valid) {
		consumed = true
	}
	return consumed
}

// ignore ignores any pending input read since the last token.
func (l *lexer) ignore() {
	l.start = l.pos
	l.startLine, l.startCol = l.line, l.col
}

// ignoreRun ignores a run of valid runes.
func (l *lexer) ignoreRun(valid string) {
	if l.acceptRun(valid) {
		l.ignore()
	}
}
