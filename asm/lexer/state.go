package lexer

import "unicode/utf8"

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
	case eof:
		l.emitEOF()
		return nil
	case utf8.RuneError:
		// Append error but continue lexing.
		l.errorf("illegal UTF-8 encoding")
		return lexToken
	}
	panic("not yet implemented.")
}
