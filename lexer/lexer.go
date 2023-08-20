package lexer

import (
	"compile/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// read the next character and advance our position in the input string
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL" (null)
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok.Type = token.TypeEQ
			tok.Literal = "=="
			l.readChar()
		} else {
			tok = token.NewToken(token.TypeAssign, l.ch)
		}
	case ';':
		tok = token.NewToken(token.TypeSemicolon, l.ch)
	case '(':
		tok = token.NewToken(token.TypeLParen, l.ch)
	case ')':
		tok = token.NewToken(token.TypeRParen, l.ch)
	case ',':
		tok = token.NewToken(token.TypeComma, l.ch)
	case '+':
		tok = token.NewToken(token.TypePlus, l.ch)
	case '-':
		tok = token.NewToken(token.TypeMinus, l.ch)
	case '{':
		tok = token.NewToken(token.TypeLBrace, l.ch)
	case '}':
		tok = token.NewToken(token.TypeRBrace, l.ch)
	case '/':
		tok = token.NewToken(token.TypeSlash, l.ch)
	case '*':
		tok = token.NewToken(token.TypeAsterisk, l.ch)
	case '<':
		tok = token.NewToken(token.TypeLT, l.ch)
	case '>':
		tok = token.NewToken(token.TypeGT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			tok.Type = token.TypeNotEQ
			tok.Literal = "!="
			l.readChar()
		} else {
			tok = token.NewToken(token.TypeBang, l.ch)
		}

	case 0:
		tok.Literal = ""
		tok.Type = token.TypeEOF
	default:

		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.TypeInt
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = token.NewToken(token.TypeIllegal, l.ch)
		}

	}

	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0

	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]

}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'

}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' ||
		ch == '_'
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
