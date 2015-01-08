package lexer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/mewlang/go/token"
)

// TODO(u): Optimize lexString, lexStringRaw, lexLineComment and lexBlockComment
// using strings.IndexAny.

const (
	// whitespace specifies the white space characters (except newline), which
	// include spaces (U+0020), horizontal tabs (U+0009), and carriage returns
	// (U+000D).
	whitespace = " \t\r"
	// decimal specifies the decimal digit characters.
	decimal = "0123456789"
	// octal specifies the octal digit characters.
	octal = "01234567"
	// hex specifies the hexadecimal digit characters.
	hex = "0123456789ABCDEFabcdef"
)

// A stateFn represents the state of the lexer as a function that returns a
// state function.
type stateFn func(l *lexer) stateFn

// lexToken lexes a token of the Go programming language. It is the initial
// state function of the lexer.
func lexToken(l *lexer) stateFn {
	// Ignore white space characters (except newline).
	l.ignoreRun(whitespace)

	r := l.next()
	switch r {
	case eof:
		insertSemicolon(l)
		// Terminate the lexer with a nil state function.
		return nil
	case '\n':
		l.ignore()
		insertSemicolon(l)
		// Update the index to the first token of the current line.
		l.first = len(l.tokens)
		return lexToken
	case '/':
		return lexDivOrComment
	case '!':
		return lexNot
	case '<':
		return lexLessArrowOrShl
	case '>':
		return lexGreaterOrShr
	case '&':
		return lexAndOrClear
	case '|':
		return lexOr
	case '=':
		return lexEqOrAssign
	case ':':
		return lexColonOrDeclAssign
	case '*':
		return lexMul
	case '%':
		return lexMod
	case '^':
		return lexXor
	case '+':
		return lexAddOrInc
	case '-':
		return lexSubOrDec
	case '(':
		l.emit(token.Lparen)
		return lexToken
	case '[':
		l.emit(token.Lbrack)
		return lexToken
	case '{':
		l.emit(token.Lbrace)
		return lexToken
	case ')':
		l.emit(token.Rparen)
		return lexToken
	case ']':
		l.emit(token.Rbrack)
		return lexToken
	case '}':
		l.emit(token.Rbrace)
		return lexToken
	case ',':
		l.emit(token.Comma)
		return lexToken
	case ';':
		l.emit(token.Semicolon)
		return lexToken
	case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		l.backup()
		return lexDotOrNumber
	case '\'':
		return lexRune
	case '"':
		return lexString
	case '`':
		return lexRawString
	}

	// Check if r is a Unicode letter or an underscore character.
	if isLetter(r) {
		return lexKeywordOrIdent
	}

	l.emit(token.Invalid)

	// Append error but continue lexing.
	l.errorf("syntax error: unexpected %#U", r)
	return lexToken
}

// isLetter returns true if r is a Unicode letter or an underscore, and false
// otherwise.
func isLetter(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

// isValid returns true if r is a valid Unicode code point in a Go source text,
// and false otherwise.
func isValid(r rune) bool {
	return r != utf8.RuneError && r != bom && r != nul
}

// lexDivOrComment lexes a division operator (/), a division assignment operator
// (/=), a line comment (//), or a block comment (/*). A slash character (/) has
// already been consumed.
func lexDivOrComment(l *lexer) stateFn {
	switch l.next() {
	case '=':
		// Division assignment operator (/=).
		l.emit(token.DivAssign)
		return lexToken
	case '/':
		// Line comment (//).
		return lexLineComment
	case '*':
		// Block comment (/*).
		return lexBlockComment
	default:
		// Division operator (/).
		l.backup()
		l.emit(token.Div)
		return lexToken
	}
}

// lexLineComment lexes a line comment. A line comment acts like a newline.
func lexLineComment(l *lexer) stateFn {
	kind := token.Comment
	insertSemicolon(l)
	for {
		r := l.next()
		switch r {
		case eof:
			// Strip carriage returns.
			s := strings.Replace(l.input[l.start:l.pos], "\r", "", -1)
			l.emitCustom(kind, s)

			// Terminate the lexer with a nil state function.
			return nil
		case '\n':
			// Strip carriage returns and trailing newline.
			s := strings.Replace(l.input[l.start:l.pos-1], "\r", "", -1)
			l.emitCustom(kind, s)

			// Update the index to the first token of the current line.
			l.first = len(l.tokens)
			return lexToken
		default:
			if !isValid(r) {
				kind |= token.Invalid
			}
		}
	}
}

// lexBlockComment lexes a block comment. A block comment containing one or more
// newlines acts like a newline, otherwise it acts like a space.
func lexBlockComment(l *lexer) stateFn {
	hasNewline := false
	kind := token.Comment
	for !strings.HasSuffix(l.input[l.start:l.pos], "*/") {
		r := l.next()
		switch r {
		case eof:
			insertSemicolon(l)

			// Strip carriage returns.
			s := strings.Replace(l.input[l.start:l.pos], "\r", "", -1)
			l.emitCustom(token.Comment|token.Invalid, s)

			// Terminate the lexer with a nil state function.
			l.errorf("unexpected eof in comment")
			return nil
		case '\n':
			hasNewline = true
		default:
			if !isValid(r) {
				kind |= token.Invalid
			}
		}
	}
	if hasNewline {
		insertSemicolon(l)
	}

	// Strip carriage returns.
	s := strings.Replace(l.input[l.start:l.pos], "\r", "", -1)
	l.emitCustom(kind, s)

	if hasNewline {
		// Update the index to the first token of the current line.
		l.first = len(l.tokens)
	}

	return lexToken
}

// lexNot lexes a logical not operator (!), or a not equal comparison operator
// (!=). An exclamation mark character (!) has already been consumed.
func lexNot(l *lexer) stateFn {
	switch l.next() {
	case '=':
		// Not equal comparison operator (!=).
		l.emit(token.Neq)
	default:
		// Logical not operator (!).
		l.backup()
		l.emit(token.Not)
	}
	return lexToken
}

// lexLessArrowOrShl lexes a less than comparison operator (<), a less than or
// equal comparison operator (<=), a left shift operator (<<), a left shift
// assignment operator (<<=), or a channel communication operator (<-). A
// less-than sign character (<) has already been consumed.
func lexLessArrowOrShl(l *lexer) stateFn {
	switch l.next() {
	case '-':
		// Channel communication operator (<-).
		l.emit(token.Arrow)
	case '<':
		if l.accept("=") {
			// Left shift assignment operator (<<=).
			l.emit(token.ShlAssign)
		} else {
			// Left shift operator (<<).
			l.emit(token.Shl)
		}
	case '=':
		// Less than or equal comparison operator (<=).
		l.emit(token.Lte)
	default:
		// Less than comparison operator (<).
		l.backup()
		l.emit(token.Lt)
	}
	return lexToken
}

// lexGreaterOrShr lexes a greater than comparison operator (>), a greater than
// or equal comparison operator (>=), a right shift operator (>>), or a right
// shift assignment operator (>>=). A greater-than sign character (>) has
// already been consumed.
func lexGreaterOrShr(l *lexer) stateFn {
	switch l.next() {
	case '>':
		if l.accept("=") {
			// Right shift assignment operator (>>=).
			l.emit(token.ShrAssign)
		} else {
			// Right shift operator (>>).
			l.emit(token.Shr)
		}
	case '=':
		// Greater than or equal comparison operator (>=).
		l.emit(token.Gte)
	default:
		// Greater than comparison operator (>).
		l.backup()
		l.emit(token.Gt)
	}
	return lexToken
}

// lexAndOrClear lexes a bitwise AND operator (&), a bitwise AND assignment
// operator (&=), a bit clear operator (&^), a bit clear assignment operator
// (&^=), or a logical AND operator (&&). An ampersand character (&) has already
// been consumed.
func lexAndOrClear(l *lexer) stateFn {
	switch l.next() {
	case '^':
		if l.accept("=") {
			// Bit clear assignment operator (&^=).
			l.emit(token.ClearAssign)
		} else {
			// Bit clear operator (&^).
			l.emit(token.Clear)
		}
	case '&':
		// Logical AND operator (&&).
		l.emit(token.Land)
	case '=':
		// Bitwise AND assignment operator (&=).
		l.emit(token.AndAssign)
	default:
		// Bitwise AND operator (&).
		l.backup()
		l.emit(token.And)
	}
	return lexToken
}

// lexOr lexes a bitwise OR operator (|), a bitwise OR assignment operator (|=),
// or a logical OR operator (||). A vertical bar character (|) has already been
// consumed.
func lexOr(l *lexer) stateFn {
	switch l.next() {
	case '|':
		// Logical OR operator (||).
		l.emit(token.Lor)
	case '=':
		// Bitwise OR assignment operator (|=).
		l.emit(token.OrAssign)
	default:
		// Bitwise OR operator (|).
		l.backup()
		l.emit(token.Or)
	}
	return lexToken
}

// lexEqOrAssign lexes an equal comparison operator (==), or an assignment
// operator (=). An equal sign character (=) has already been consumed.
func lexEqOrAssign(l *lexer) stateFn {
	switch l.next() {
	case '=':
		// Equal comparison operator (==).
		l.emit(token.Eq)
	default:
		// Assignment operator (=).
		l.backup()
		l.emit(token.Assign)
	}
	return lexToken
}

// lexColonOrDeclAssign lexes a colon delimiter (:), or a declare and initialize
// operator (:=). A colon character (:) has already been consumed.
func lexColonOrDeclAssign(l *lexer) stateFn {
	switch l.next() {
	case '=':
		// Declare and initialize operator (:=).
		l.emit(token.DeclAssign)
	default:
		// Colon delimiter (:).
		l.backup()
		l.emit(token.Colon)
	}
	return lexToken
}

// lexMul lexes a multiplication operator (*), or a multiplication assignment
// operator (*=). An asterisk character (*) has already been consumed.
func lexMul(l *lexer) stateFn {
	switch l.next() {
	case '=':
		// Multiplication assignment operator (*=).
		l.emit(token.MulAssign)
	default:
		// Multiplication operator (*). The semantical analysis will determine if
		// the token is part of a pointer dereference expression.
		l.backup()
		l.emit(token.Mul)
	}
	return lexToken
}

// lexMod lexes a modulo operator (%), or a modulo assignment operator (%=). A
// percent sign character (%) has already been consumed.
func lexMod(l *lexer) stateFn {
	switch l.next() {
	case '=':
		// Modulo assignment operator (%=).
		l.emit(token.ModAssign)
	default:
		// Modulo operator (%).
		l.backup()
		l.emit(token.Mod)
	}
	return lexToken
}

// lexXor lexes a bitwise XOR operator (^), or a bitwise XOR assignment operator
// (^=). A caret character (^) has already been consumed.
func lexXor(l *lexer) stateFn {
	switch l.next() {
	case '=':
		// Bitwise XOR assignment operator (^=).
		l.emit(token.XorAssign)
	default:
		// Bitwise XOR operator (^). The semantical analysis will determine if the
		// token is part of a bitwise complement expression.
		l.backup()
		l.emit(token.Xor)
	}
	return lexToken
}

// lexAddOrInc lexes an addition operator (+), an addition assignment operator
// (+=), or an increment statement operator (++). A plus character (+) has
// already been consumed.
func lexAddOrInc(l *lexer) stateFn {
	switch l.next() {
	case '=':
		// Addition assignment operator (+=).
		l.emit(token.AddAssign)
	case '+':
		// Increment statement operator (++).
		l.emit(token.Inc)
	default:
		// Addition operator (+). The semantical analysis will determine if the
		// token is part of a positive number as an unary operator.
		l.backup()
		l.emit(token.Add)
	}
	return lexToken
}

// lexSubOrDec lexes a subtraction operator (-), a subtraction assignment
// operator (-=), or a decrement statement operator (--). A minus character (-)
// has already been consumed.
func lexSubOrDec(l *lexer) stateFn {
	switch l.next() {
	case '=':
		// Subtraction assignment operator (-=).
		l.emit(token.SubAssign)
	case '-':
		// Decrement statement operator (--).
		l.emit(token.Dec)
	default:
		// Subtraction operator (-). The semantical analysis will determine if the
		// token is part of a negative number as an unary operator.
		l.backup()
		l.emit(token.Sub)
	}
	return lexToken
}

// lexDotOrNumber lexes a dot delimiter (.), an ellipsis delimiter (...), or a
// number (123, 0x7B, 0173, .123, 123.45, 1e-15, 2i).
func lexDotOrNumber(l *lexer) stateFn {
	// Integer part.
	var kind token.Kind
	if l.accept("0") {
		kind = token.Int
		// Early return for hexadecimal constant.
		if l.accept("xX") {
			if !l.acceptRun(hex) {
				l.emit(token.Int | token.Invalid)

				// Append error but continue lexing.
				l.errorf("missing digits in hexadecimal constant")
				return lexToken
			}
			l.emit(token.Int)
			return lexToken
		}
	}
	if l.acceptRun(decimal) {
		kind = token.Int
	}

	// Decimal point.
	if l.accept(".") {
		if kind == token.Int {
			kind = token.Float
		} else {
			kind = token.Dot
		}
	}

	// Fraction part.
	if l.acceptRun(decimal) {
		kind = token.Float
	}

	// Early return for dot or ellipsis delimiter.
	if kind == token.Dot {
		if strings.HasPrefix(l.input[l.pos:], "..") {
			l.pos += 2
			l.col += 2
			l.width = 0
			kind = token.Ellipsis
		}
		l.emit(kind)
		return lexToken
	}

	// Exponent part.
	if l.accept("eE") {
		kind = token.Float

		// Optional sign.
		l.accept("+-")

		if !l.acceptRun(decimal) {
			l.emit(token.Float | token.Invalid)

			// Append error but continue lexing.
			l.errorf("missing digits in floating-point exponent")
			return lexToken
		}
	}

	// Imaginary.
	if l.accept("i") {
		kind = token.Imag
	}

	// Validate octal numbers.
	if kind == token.Int {
		if s := l.input[l.start:l.pos]; s[0] == '0' {
			if pos := strings.IndexAny(s, "89"); pos != -1 {
				l.emit(token.Int | token.Invalid)

				// Append error but continue lexing.
				l.errorf("invalid digit %q in octal constant", s[pos])
				return lexToken
			}
		}
	}

	l.emit(kind)
	return lexToken
}

// lexRune lexes a rune literal ('a'). A single quote character (') has already
// been consumed.
func lexRune(l *lexer) stateFn {
	// Consume one or more characters enclosed in single quotes.
	kind := token.Rune
	for i := 0; ; i++ {
		r := l.next()
		switch r {
		case eof:
			l.emit(token.Rune | token.Invalid)

			insertSemicolon(l)

			// Terminate the lexer with a nil state function.
			l.errorf("unexpected eof in rune literal")
			return nil
		case '\n':
			l.backup()
			l.emit(token.Rune | token.Invalid)

			insertSemicolon(l)
			// Update the index to the first token of the current line.
			l.first = len(l.tokens)

			// Append error but continue lexing.
			l.errorf("unexpected newline in rune literal")
			return lexToken
		case '\\':
			// Consume backslash escape sequence.
			err := consumeEscape(l, '\'')
			if err != nil {
				kind |= token.Invalid

				// Append error but continue lexing the rune literal.
				l.errs = append(l.errs, err)
			}
		case '\'':
			switch i {
			case 0:
				l.emit(token.Rune | token.Invalid)

				// Append error but continue lexing.
				l.errorf("empty rune literal or unescaped ' in rune literal")
				return lexToken
			case 1:
				l.emit(kind)
				return lexToken
			default:
				l.emit(token.Rune | token.Invalid)

				// Append error but continue lexing.
				l.errorf("too many characters in rune literal")
				return lexToken
			}
		default:
			if !isValid(r) {
				kind |= token.Invalid
			}
		}
	}
}

// lexString lexes an interpreted string literal ("foo"). A double quote
// character (") has already been consumed.
func lexString(l *lexer) stateFn {
	kind := token.String
	for {
		r := l.next()
		switch r {
		case eof:
			l.emit(token.String | token.Invalid)

			insertSemicolon(l)

			// Terminate the lexer with a nil state function.
			l.errorf("unexpected eof in string literal")
			return nil
		case '\n':
			l.backup()
			l.emit(token.String | token.Invalid)

			insertSemicolon(l)
			// Update the index to the first token of the current line.
			l.first = len(l.tokens)

			// Append error but continue lexing.
			l.errorf("unexpected newline in string literal")
			return lexToken
		case '\\':
			// Consume backslash escape sequence.
			err := consumeEscape(l, '"')
			if err != nil {
				kind |= token.Invalid
				// Append error but continue lexing the string literal.
				l.errs = append(l.errs, err)
			}
		case '"':
			l.emit(kind)
			return lexToken
		default:
			if !isValid(r) {
				kind |= token.Invalid
			}
		}
	}
}

// lexRawString lexes a raw string literal (`foo`). A back quote character (`)
// has already been consumed.
func lexRawString(l *lexer) stateFn {
	kind := token.String
	for {
		r := l.next()
		switch r {
		case eof:
			l.emit(token.String | token.Invalid)

			insertSemicolon(l)

			// Terminate the lexer with a nil state function.
			l.errorf("unexpected eof in raw string literal")
			return nil
		case '`':
			// Strip carriage returns.
			s := strings.Replace(l.input[l.start:l.pos], "\r", "", -1)
			l.emitCustom(kind, s)
			return lexToken
		default:
			if !isValid(r) {
				kind |= token.Invalid
			}
		}
	}
}

// keywords specifies the reserved keywords of the Go programming language.
var keywords = map[string]token.Kind{
	"break":       token.Break,
	"case":        token.Case,
	"chan":        token.Chan,
	"const":       token.Const,
	"continue":    token.Continue,
	"default":     token.Default,
	"defer":       token.Defer,
	"else":        token.Else,
	"fallthrough": token.Fallthrough,
	"for":         token.For,
	"func":        token.Func,
	"go":          token.Go,
	"goto":        token.Goto,
	"if":          token.If,
	"import":      token.Import,
	"interface":   token.Interface,
	"map":         token.Map,
	"package":     token.Package,
	"range":       token.Range,
	"return":      token.Return,
	"select":      token.Select,
	"struct":      token.Struct,
	"switch":      token.Switch,
	"type":        token.Type,
	"var":         token.Var,
}

// lexKeywordOrIdent lexes a keyword, or an identifier. A Unicode letter or an
// underscore character (_) has already been consumed.
func lexKeywordOrIdent(l *lexer) stateFn {
	for {
		r := l.next()
		if r == eof {
			break
		}
		if !isLetter(r) && !unicode.IsDigit(r) {
			l.backup()
			break
		}
	}
	s := l.input[l.start:l.pos]
	if kind, ok := keywords[s]; ok {
		l.emit(kind)
	} else {
		l.emit(token.Ident)
	}
	return lexToken
}

// consumeEscape consumes an escape sequence. A valid single-character escape
// sequence is specified by valid. Single quotes are only valid within rune
// literals and double quotes are only valid within string literals. A backslash
// character (\) has already been consumed.
//
// Several backslash escapes allow arbitrary values to be encoded as ASCII text.
// There are four ways to represent the integer value as a numeric constant: \x
// followed by exactly two hexadecimal digits; \u followed by exactly four
// hexadecimal digits; \U followed by exactly eight hexadecimal digits, and a
// plain backslash \ followed by exactly three octal digits. In each case the
// value of the literal is the value represented by the digits in the
// corresponding base.
//
// Although these representations all result in an integer, they have different
// valid ranges. Octal escapes must represent a value between 0 and 255
// inclusive. Hexadecimal escapes satisfy this condition by construction. The
// escapes \u and \U represent Unicode code points so within them some values
// are illegal, in particular those above 0x10FFFF and surrogate halves.
//
// After a backslash, certain single-character escapes represent special values:
//    \a   U+0007 alert or bell
//    \b   U+0008 backspace
//    \f   U+000C form feed
//    \n   U+000A line feed or newline
//    \r   U+000D carriage return
//    \t   U+0009 horizontal tab
//    \v   U+000b vertical tab
//    \\   U+005c backslash
//    \'   U+0027 single quote  (valid escape only within rune literals)
//    \"   U+0022 double quote  (valid escape only within string literals)
//
// All other sequences starting with a backslash are illegal inside rune and
// string literals.
//
// ref: http://golang.org/ref/spec#Rune_literals
func consumeEscape(l *lexer, valid rune) error {
	r := l.next()
	switch r {
	case eof:
		return errors.New("unexpected eof in escape sequence")
	case '0', '1', '2', '3':
		// Octal escape.
		for i := 0; i < 2; i++ {
			if !l.accept(octal) {
				r := l.next()
				switch r {
				case eof:
					return errors.New("unexpected eof in octal escape")
				case valid:
					return fmt.Errorf("too few digits in octal escape; expected 3, got %d", 1+i)
				}
				return fmt.Errorf("non-octal character %#U in octal escape", r)
			}
		}
		s := l.input[l.pos-3 : l.pos]
		_, err := strconv.ParseUint(s, 8, 8)
		if err != nil {
			return fmt.Errorf("invalid octal escape; %v", err)
		}
	case 'x':
		// Hexadecimal escape.
		for i := 0; i < 2; i++ {
			if !l.accept(hex) {
				r := l.next()
				switch r {
				case eof:
					return errors.New("unexpected eof in hex escape")
				case valid:
					return fmt.Errorf("too few digits in hex escape; expected 2, got %d", i)
				}
				return fmt.Errorf("non-hex character %#U in hex escape", r)
			}
		}
	case 'u', 'U':
		// Unicode escape.
		n := 4
		if r == 'U' {
			n = 8
		}
		for i := 0; i < n; i++ {
			if !l.accept(hex) {
				r := l.next()
				switch r {
				case eof:
					return errors.New("unexpected eof in Unicode escape")
				case valid:
					return fmt.Errorf("too few digits in Unicode escape; expected %d, got %d", n, i)
				}
				return fmt.Errorf("non-hex character %#U in Unicode escape", r)
			}
		}
		s := l.input[l.pos-n : l.pos]
		x, err := strconv.ParseUint(s, 16, 32)
		if err != nil {
			return fmt.Errorf("invalid Unicode escape; %v", err)
		}
		r := rune(x)
		if !utf8.ValidRune(r) {
			return fmt.Errorf("invalid Unicode code point %#U in escape sequence", r)
		}
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', valid:
		// Single-character escape.
	default:
		return fmt.Errorf("unknown escape sequence %#U", r)
	}
	return nil
}

// insertSemicolon inserts a semicolon if the correct conditions have been met.
//
// When the input is broken into tokens, a semicolon is automatically inserted
// into the token stream at the end of a non-blank line if the line's final
// token is
//    * an identifier
//    * an integer, floating-point, imaginary, rune, or string literal
//    * one of the keywords break, continue, fallthrough, or return
//    * one of the operators and delimiters ++, --, ), ], or }
//
// ref: http://golang.org/ref/spec#Semicolons
var insertSemicolon = func(l *lexer) {
	insert := false
	trailingComments := false
	tok := token.Token{
		Kind: token.Semicolon,
		Val:  ";",
	}
	var pos int
	for pos = len(l.tokens) - 1; pos >= l.first; pos-- {
		last := l.tokens[pos]
		switch last.Kind {
		case token.Comment:
			// Ignore trailing comments.
			trailingComments = true
			continue
		case token.Ident:
			// * an identifier
		case token.Int, token.Float, token.Imag, token.Rune, token.String:
			// * an integer, floating-point, imaginary, rune, or string literal
		case token.Break, token.Continue, token.Fallthrough, token.Return:
			// * one of the keywords break, continue, fallthrough, or return
		case token.Inc, token.Dec, token.Rparen, token.Rbrack, token.Rbrace:
			// * one of the operators and delimiters ++, --, ), ], or }
		default:
			return
		}
		insert = true
		tok.Line = last.Line
		tok.Col = last.Col + utf8.RuneCountInString(last.Val)
		break
	}

	// Insert a semicolon.
	if insert {
		l.tokens = append(l.tokens, tok)

		if trailingComments {
			// Move trailing comments to the end.
			copy(l.tokens[pos+2:], l.tokens[pos+1:])
			// Insert a semicolon before the trailing comments.
			l.tokens[pos+1] = tok
		}
	}
}
