package lexer

import (
	"strings"
	"unicode/utf8"

	"github.com/mewlang/llvm/asm/token"
)

// TODO: Optimize lexString and lexComment using strings.IndexAny.

const (
	// TODO: Check which whitespace characters are valid in LLVM IR assembly.

	// whitespace specifies the white space characters (except newline), which
	// include spaces (U+0020), horizontal tabs (U+0009), and carriage returns
	// (U+000D).
	whitespace = " \t\r"
	// decimal specifies the decimal digit characters.
	decimal = "0123456789"
	// hex specifies the hexadecimal digit characters.
	hex = "0123456789ABCDEFabcdef"
)

// A stateFn represents the state of the lexer as a function that returns a
// state function.
type stateFn func(l *lexer) stateFn

// lexToken lexes a token of the LLVM IR assembly language. It is the initial
// state function of the lexer.
func lexToken(l *lexer) stateFn {
	r := l.next()
	switch r {
	// Special tokens.
	case utf8.RuneError:
		// Emit error token but continue lexing next token.
		l.emitErrorf("illegal UTF-8 encoding")
		return lexToken
	case eof:
		l.emitEOF()
		return nil
	case ';':
		return lexComment

		// Identifiers.
	}
	panic("not yet implemented.")
}

// lexComment lexes a line comment which acts like a newline. A semicolon (;)
// has already been consumed.
func lexComment(l *lexer) stateFn {
	for {
		switch l.next() {
		case utf8.RuneError:
			// Append error but continue lexing line comment.
			l.errorf("illegal UTF-8 encoding")
		case eof:
			l.emit(token.Comment)
			// Terminate the lexer with a nil state function.
			return nil
		case '\n':
			// Ignore trailing carriage return and newline characters.
			s := strings.TrimRight(l.input[l.start:l.pos], "\r\n")
			l.emitCustom(token.Comment, s)
			return lexToken
		}
	}
	return lexToken
}
