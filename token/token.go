package token

type Type string

const (
	TypeIllegal = "ILLEGAL"
	TypeEOF     = "EOF"

	// Identifiers + literals
	TypeIdent = "IDENT" // add, foobar, x, y, ...
	TypeInt   = "INT"   // 1343456

	// Operators
	TypeAssign   = "="
	TypePlus     = "+"
	TypeMinus    = "-"
	TypeBang     = "!"
	TypeAsterisk = "*"
	TypeSlash    = "/"
	TypeLT       = "<"
	TypeGT       = ">"

	// Delimiters
	TypeComma     = ","
	TypeSemicolon = ";"

	TypeLParen = "("
	TypeRParen = ")"
	TypeLBrace = "{"
	TypeRBrace = "}"

	// Keywords
	TypeFunction = "FUNCTION"
	TypeLet      = "LET"
	TypeTrue     = "TRUE"
	TypeFalse    = "FALSE"
	TypeIf       = "IF"
	TypeElse     = "ELSE"
	TypeReturn   = "RETURN"

	TypeEQ    = "=="
	TypeNotEQ = "!="
)

var (
	keywords = map[string]Type{
		"fn":     TypeFunction,
		"let":    TypeLet,
		"true":   TypeTrue,
		"false":  TypeFalse,
		"if":     TypeIf,
		"else":   TypeElse,
		"return": TypeReturn,
	}
)

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return TypeIdent
}

type Token struct {
	Type    Type
	Literal string
}

func NewToken(tokenType Type, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
