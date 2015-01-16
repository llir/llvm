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
	// identifier (i.e. all characters in the identifier except the first). All
	// characters of a label may be from the tail set, even the first character.
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
		return lexAt // @foo, @"fo\6F", @42
	case '%':
		return lexPercent // %foo, %"fo\6F", %42
	case '!':
		return lexExclaim // !, !foo, !foo\2A
	case '$':
		return lexDollar // $foo, $"fo\6F", $foo:
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
		return lexQuote // "fo\6F":, "fo\6F"
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
			s := strings.TrimRight(l.input[l.start:l.cur], "\r")
			l.emitCustom(token.Comment, s)
			l.emitEOF()
			// Terminate the lexer with a nil state function.
			return nil
		case '\n':
			// Ignore trailing carriage return and newline characters.
			s := strings.TrimRight(l.input[l.start:l.cur], "\r\n")
			l.emitCustom(token.Comment, s)
			return lexToken
		}
	}
}

// lexAt lexes a global variable (@foo, @"fo\6F") or a global ID (@42). An at
// character (@) has already been consumed.
//
//    GlobalVar = @[-a-zA-Z$._][-a-zA-Z$._0-9]*
//    GlobalVar = @"[^"]*"   (may contain hex escapes)
//    GlobalID  = @[0-9]+
func lexAt(l *lexer) stateFn {
	// Store start position to skip the leading at character (@).
	start := l.cur

	switch {
	// @foo
	case l.accept(head):
		l.acceptRun(tail)
		s := l.input[start:l.cur]
		l.emitCustom(token.GlobalVar, s)

	// @"foo", @"fo\6F"
	case l.accept(`"`):
		s, ok := readString(l)
		if !ok {
			// Terminate the lexer with a nil state function.
			return nil
		}
		l.emitCustom(token.GlobalVar, s)

	// @42
	case l.acceptRun(decimal):
		s := l.input[start:l.cur]
		l.emitCustom(token.GlobalID, s)

	default:
		// Emit error token but continue lexing next token.
		l.emitErrorf("unexpected '@'")
	}

	return lexToken
}

// lexPercent lexes a local variable (%foo, %"fo\6F") or a local ID (%42). A
// percent character (%) has already been consumed.
//
//    LocalVar = %[-a-zA-Z$._][-a-zA-Z$._0-9]*
//    LocalVar = %"[^"]*"   (may contain hex escapes)
//    LocalID  = %[0-9]+
func lexPercent(l *lexer) stateFn {
	log.Println("lexPercent: not yet implemented.")
	return nil
}

// lexExclaim lexes an exclamation mark (!) or a metadata variable (!foo,
// !foo\2A). An exclamation mark (!) has already been consumed.
//
//    Exclaim     = !
//    MetadataVar = ![-a-zA-Z$._][-a-zA-Z$._0-9]*   (may contain hex escapes)
func lexExclaim(l *lexer) stateFn {
	log.Println("lexExclaim: not yet implemented.")
	return nil
}

// lexDollar lexes an COMDAT variable ($foo, $"fo\6F") or a label ($foo:). A
// dollar sign ($) has already been consumed.
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
	// Store the current token position.
	cur := l.cur

	// end and kind tracks the end position and token type of the longest
	// candidate token respectively.
	end, kind := l.cur, token.Error

	// Try lexing an ellipsis (...).
	if l.accept(".") && l.accept(".") {
		end, kind = l.cur, token.Ellipsis
	}

	// Restore the current token position.
	l.cur = cur

	// Try lexing a label starting with a dot (.foo:).
	l.acceptRun(tail)
	if l.accept(":") && l.cur > end {
		end, kind = l.cur, token.Label
	}

	// Set the current position to the end position of the longest candidate
	// token.
	l.cur = end

	switch kind {
	case token.Ellipsis:
		l.emit(token.Ellipsis)
	case token.Label:
		s := l.input[l.start : l.cur-1] // skip trailing colon (:)
		l.emitCustom(token.Label, s)
	default:
		// Emit error token but continue lexing next token.
		l.emitErrorf("unexpected '.'")
	}

	return lexToken
}

// lexQuote lexes a string constant ("fo\6F") or a quoted label ("fo\6F":). A
// double quote (") has already been consumed.
//
//    Label  = "[^"]+":   (may contain hex escapes)
//    String = "[^"]*"   (may contain hex escapes)
func lexQuote(l *lexer) stateFn {
	// Consume a string constant ("foo", "fo\6F").
	s, ok := readString(l)
	if !ok {
		// Terminate the lexer with a nil state function.
		return nil
	}

	switch {
	case l.accept(":"):
		l.emitCustom(token.Label, s)
	default:
		l.emitCustom(token.String, s)
	}

	return lexToken
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

// readString consumes a string constant ("foo", "fo\6F") and returns its
// unescaped string value. The returned boolean is false in the case of an
// unexpected EOF. A double quote (") has already been consumed.
func readString(l *lexer) (s string, ok bool) {
	// Store start position to skip leading double quote (") and any token
	// specific characters (e.g. @, %).
	start := l.cur
	for {
		switch l.next() {
		case eof:
			l.emitErrorf("unexpected eof in quoted string")
			return "", false
		case utf8.RuneError:
			// Append error but continue lexing string constant.
			l.errorf("illegal UTF-8 encoding")
		case '"':
			s := l.input[start : l.cur-1] // skip leading and trailing double quotes (")
			return unescape(s), true
		}
	}
}

// unescape replaces hexadecimal escape sequences (\xx) in s with their
// corresponding characters.
func unescape(s string) string {
	if !strings.ContainsRune(s, '\\') {
		return s
	}
	j := 0
	buf := []byte(s)
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == '\\' && i+2 < len(s) {
			x1, ok := unhex(s[i+1])
			if ok {
				x2, ok := unhex(s[i+2])
				if ok {
					b = x1<<4 | x2
					i += 2
				}
			}
		}
		if i != j {
			buf[j] = b
		}
		j++
	}
	return string(buf[:j]) // TODO: Check that the end offset j isn't off-by-one.
}

// unhex returns the numeric value represented by the hexadecimal digit b. It
// returns false if b is not a hexadecimal digit.
func unhex(b byte) (v byte, ok bool) {
	// This is an adapted copy of the unhex function from the strconv package,
	// which is goverend by a BSD-style license.
	switch {
	case '0' <= b && b <= '9':
		return b - '0', true
	case 'a' <= b && b <= 'f':
		return b - 'a' + 10, true
	case 'A' <= b && b <= 'F':
		return b - 'A' + 10, true
	}
	return 0, false
}

// isDigit returns true if r is a digit (0-9), and false otherwise.
func isDigit(r rune) bool {
	return strings.ContainsRune(decimal, r)
}

// isAlpha returns true if r is an alphabetic character, and false otherwise.
func isAlpha(r rune) bool {
	return strings.ContainsRune(alpha, r)
}
