package lexer

import (
	"strings"
	"unicode/utf8"

	"github.com/mewlang/llvm/asm/token"
)

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
		// TODO: Handle identifiers.

	// Operators and delimiters.
	case '.':
		// Try to consume two more dots and restore position if unable.
		pos := l.pos
		if l.accept(".") && l.accept(".") {
			l.emit(token.Ellipsis)
			return lexToken
		}
		l.pos = pos
	case '=':
		l.emit(token.Equal)
		return lexToken
	case ',':
		l.emit(token.Comma)
		return lexToken
	case '*':
		l.emit(token.Star)
		return lexToken
	case '[':
		l.emit(token.Lbrack)
		return lexToken
	case ']':
		l.emit(token.Rbrack)
		return lexToken
	case '{':
		l.emit(token.Lbrace)
		return lexToken
	case '}':
		l.emit(token.Rbrace)
		return lexToken
	case '(':
		l.emit(token.Lparen)
		return lexToken
	case ')':
		l.emit(token.Rparen)
		return lexToken
	case '<':
		l.emit(token.Less)
		return lexToken
	case '>':
		l.emit(token.Greater)
		return lexToken
	case '!':
		return lexExclaim
	case '#':
		return lexAttrID
	}

	// Emit error token but continue lexing next token.
	l.emitErrorf("invalid token starting with %q", r)
	return lexToken
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
}

// lexExclaim lexes an exclamation mark (!) or a metadata variable (!foo). An
// exclamation mark (!) has already been consumed.
func lexExclaim(l *lexer) stateFn {
	panic("not yet implemented.")
}

// lexAttrID lexes an attribute ID (#42). A hash character (#) has already been
// consumed.
func lexAttrID(l *lexer) stateFn {
	panic("not yet implemented.")
}
