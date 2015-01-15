package lexer

import (
	"log"
	"strings"
	"unicode/utf8"

	"github.com/mewlang/llvm/asm/token"
)

const (
	// whitespace specifies the white space characters.
	whitespace = "\x00 \t\r\n"
	// decimal specifies the decimal digit characters.
	decimal = "0123456789"
	// hex specifies the hexadecimal digit characters.
	hex = "0123456789ABCDEFabcdef"
	// alpha specifies the alphabetic characters.
	alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	// head is the set of valid characters for the first character of an
	// identifier.
	head = alpha + "$-._"
	// tail is the set of valid characters for the remaining characters of an
	// identifier (i.e. all characters in the identifier except the first).
	tail = head + decimal
)

// A stateFn represents the state of the lexer as a function that returns a
// state function.
type stateFn func(l *lexer) stateFn

// lexToken lexes a token of the LLVM IR assembly language. It is the initial
// state function of the lexer.
func lexToken(l *lexer) stateFn {
	l.ignoreRun(whitespace)

	r := l.next()
	switch r {
	// Special tokens.
	case utf8.RuneError:
		// Emit error token but continue lexing next token.
		l.emitErrorf("illegal UTF-8 encoding")
		return lexToken
	case eof:
		l.emitEOF()
		// Terminate the lexer with a nil state function.
		return nil
	case ';':
		return lexComment // ; foo

	// Identifiers.
	case '@':
		return lexAt // @foo, @"foo", @42
	case '%':
		return lexPercent // %foo, %"foo", %42
	case '!':
		return lexExclaim // !, !foo
	case '$':
		return lexDollar // $foo, $foo:
	case '#':
		return lexHash // #42

	// Operators and delimiters.
	case '.':
		return lexDot // .foo:, ...
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

	// Constants.
	case '"':
		return lexQuote // "foo:", "foo"
	}

	// Lex label, integer constant, floating-point constant or hexadecimal
	// floating-point constant.
	if r == '+' || r == '-' || isDigit(r) {
		l.backup()
		return lexDigitOrSign // 42:, -foo:, 42, +0.314e+1, 0x1e
	}

	// Lex label, type, keyword or hexadecimal integer constant.
	if r == '_' || isAlpha(r) {
		l.backup()
		return lexLetter // foo:, _foo:, i32, void, add, u0x10
	}

	// Emit error token but continue lexing next token.
	l.emitErrorf("invalid token starting with %q", r)
	return lexToken
}

// lexComment lexes a line comment which acts like a newline. A semicolon (;)
// has already been consumed.
//
//    Comment = ;[^\n]*
func lexComment(l *lexer) stateFn {
	for {
		switch l.next() {
		case utf8.RuneError:
			// Append error but continue lexing line comment.
			l.errorf("illegal UTF-8 encoding")
		case eof:
			// Ignore trailing carriage return characters.
			s := strings.TrimRight(l.input[l.start:l.pos], "\r")
			l.emitCustom(token.Comment, s)
			l.emitEOF()
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

// lexAt lexes a global variable (@foo, @"foo") or a global ID (@42). An at
// character (@) has already been consumed.
//
//    GlobalVar = @[-a-zA-Z$._][-a-zA-Z$._0-9]*
//    GlobalVar = @"[^"]*"   (may contain hex escapes)
//    GlobalID  = @[0-9]+
func lexAt(l *lexer) stateFn {
	log.Println("lexAt: not yet implemented.")
	return nil
}

// lexPercent lexes a local variable (%foo, %"foo") or a local ID (%42). A
// percent character (%) has already been consumed.
//
//    LocalVar = %[-a-zA-Z$._][-a-zA-Z$._0-9]*
//    LocalVar = %"[^"]*"   (may contain hex escapes)
//    LocalID  = %[0-9]+
func lexPercent(l *lexer) stateFn {
	log.Println("lexPercent: not yet implemented.")
	return nil
}

// lexExclaim lexes an exclamation mark (!) or a metadata variable (!foo). An
// exclamation mark (!) has already been consumed.
//
//    Exclaim     = !
//    MetadataVar = ![-a-zA-Z$._][-a-zA-Z$._0-9]*   (may contain hex escapes)
func lexExclaim(l *lexer) stateFn {
	log.Println("lexExclaim: not yet implemented.")
	return nil
}

// lexDollar lexes an COMDAT variable ($foo) or a label ($foo:). A dollar sign
// ($) has already been consumed.
//
//    ComdatVar = $[-a-zA-Z$._][-a-zA-Z$._0-9]*
//    ComdatVar = $"[^"]*"   (may contain hex escapes)
//    Label     = [-a-zA-Z$._0-9]+:
func lexDollar(l *lexer) stateFn {
	log.Println("lexDollar: not yet implemented.")
	return nil
}

// lexHash lexes an attribute ID (#42). A hash character (#) has already been
// consumed.
//
//    AttrID = #[0-9]+
func lexHash(l *lexer) stateFn {
	log.Println("lexHash: not yet implemented.")
	return nil
}

// lexDot lexes an ellipsis (...) or a label (.foo:). A dot (.) has already been
// consumed.
//
//    Ellipsis = ...
//    Label    = [-a-zA-Z$._0-9]+:
func lexDot(l *lexer) stateFn {
	if strings.HasPrefix(l.input[l.pos:], "..") {
		l.accept(".")
		l.accept(".")
		l.emit(token.Ellipsis)
		return lexToken
	}

	// TODO: Lex label (.foo:).

	// Emit error token but continue lexing next token.
	l.emitErrorf("invalid token starting with '.'")
	return lexToken
}

// lexQuote lexes a string constant ("foo") or a quoted label ("foo":). A double
// quote (") has already been consumed.
//
//    Label  = "[^"]+":   (may contain hex escapes)
//    String = "[^"]*"   (may contain hex escapes)
func lexQuote(l *lexer) stateFn {
	log.Println("lexQuote: not yet implemented.")
	return nil
}

// lexLetter lexes a label (foo:, _foo:), a type (i32, float), a keyword (add,
// x) or a hexadecimal integer constant (u0x10). The next character is either an
// alphabetic character (a-z or A-Z) or an underscore (_).
//
//    Label   = [-a-zA-Z$._0-9]+:
//    Type    = i[0-9]+
//    Type    = float, void, …
//    Keyword = add, x, …
//    HexInt  = [us]0x[0-9A-Fa-f]+
func lexLetter(l *lexer) stateFn {
	log.Println("lexLetter: not yet implemented.")
	return nil
}

// lexDigitOrSign lexes a label (42:, -foo:), an integer constant (42, -42), a
// floating-point constant (+0.314e+1) or a hexadecimal floating-point constant
// (0x1e, 0xK1e, 0xL1e, 0xM1e, 0xH1e). The next character is either a digit or a
// sign character (+ or -).
//
//    Label    = [-a-zA-Z$._0-9]+:
//    Int      = [-]?[0-9]+
//    Float    = [-+]?[0-9]+[.][0-9]*([eE][-+]?[0-9]+)?
//    HexFloat = 0x[KLMH]?[0-9A-Fa-f]+
//
// The 80-bit format used by x86 is represented as 0xK followed by 20
// hexadecimal digits. The 128-bit format used by PowerPC (two adjacent doubles)
// is represented by 0xM followed by 32 hexadecimal digits. The IEEE 128-bit
// format is represented by 0xL followed by 32 hexadecimal digits. The IEEE
// 16-bit format (half precision) is represented by 0xH followed by 4
// hexadecimal digits. All hexadecimal formats are big-endian (sign bit at the
// left). [1]
//
//    [1] http://llvm.org/docs/LangRef.html#simple-constants
func lexDigitOrSign(l *lexer) stateFn {
	log.Println("lexDigitOrSign: not yet implemented.")
	return nil
}

// isDigit returns true if r is a digit (0-9), and false otherwise.
func isDigit(r rune) bool {
	return strings.ContainsRune(decimal, r)
}

// isAlpha returns true if r is an alphabetic character, and false otherwise.
func isAlpha(r rune) bool {
	return strings.ContainsRune(alpha, r)
}
