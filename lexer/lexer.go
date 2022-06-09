package lexer

import "github.com/alexdunne/writing-an-interpreter-in-go/token"

type Lexer struct {
	// input hold the given value to process
	input string

	// position tracks the current position in the input
	position int

	// readPosition trackes the next position in the input
	readPosition int

	// ch holds the current char
	ch byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()

	return l
}

// NextToken converts the current character into a Token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// check if the character is a valid letter in which case we're
		// dealing with an identifier
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}

		// check if the character is a digit in which case we're dealing
		// with a number
		if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}

		// we don't know how to handle this character so mark it as such
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()

	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readChar peaks the next character in the input and updates the internal
// state of the lexer accordingly
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition

	l.readPosition++
}

// readIdentifier iterates through the input from the current position
// until it meets an unsupport indentifier character. The values between the
// start position and end position are returned as the identifier value
func (l *Lexer) readIdentifier() string {
	pos := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

// readNumber iterates through the input from the current position
// until it meets a non digit character. The values between the
// start position and end position are returned as the number value
func (l *Lexer) readNumber() string {
	pos := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

// skipWhitespace iterates through the input until it meets a
// none whitespace character
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
