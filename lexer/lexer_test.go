package lexer

import (
	"compile/token"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
let result = add(five, ten);
!-/*5;
5 < 10 > 5;
if (5 < 10) {
	return true;
} else {
	return false;
}
10 == 10;
10 != 9;
	`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.TypeLet, "let"},
		{token.TypeIdent, "five"},
		{token.TypeAssign, "="},
		{token.TypeInt, "5"},
		{token.TypeSemicolon, ";"},
		{token.TypeLet, "let"},
		{token.TypeIdent, "ten"},
		{token.TypeAssign, "="},
		{token.TypeInt, "10"},
		{token.TypeSemicolon, ";"},
		{token.TypeLet, "let"},
		{token.TypeIdent, "add"},
		{token.TypeAssign, "="},
		{token.TypeFunction, "fn"},
		{token.TypeLParen, "("},
		{token.TypeIdent, "x"},
		{token.TypeComma, ","},
		{token.TypeIdent, "y"},
		{token.TypeRParen, ")"},
		{token.TypeLBrace, "{"},
		{token.TypeIdent, "x"},
		{token.TypePlus, "+"},
		{token.TypeIdent, "y"},
		{token.TypeSemicolon, ";"},
		{token.TypeRBrace, "}"},
		{token.TypeSemicolon, ";"},
		{token.TypeLet, "let"},
		{token.TypeIdent, "result"},
		{token.TypeAssign, "="},
		{token.TypeIdent, "add"},
		{token.TypeLParen, "("},
		{token.TypeIdent, "five"},
		{token.TypeComma, ","},
		{token.TypeIdent, "ten"},
		{token.TypeRParen, ")"},
		{token.TypeSemicolon, ";"},
		{token.TypeBang, "!"},
		{token.TypeMinus, "-"},
		{token.TypeSlash, "/"},
		{token.TypeAsterisk, "*"},
		{token.TypeInt, "5"},
		{token.TypeSemicolon, ";"},
		{token.TypeInt, "5"},
		{token.TypeLT, "<"},
		{token.TypeInt, "10"},
		{token.TypeGT, ">"},
		{token.TypeInt, "5"},
		{token.TypeSemicolon, ";"},
		{token.TypeIf, "if"},
		{token.TypeLParen, "("},
		{token.TypeInt, "5"},
		{token.TypeLT, "<"},
		{token.TypeInt, "10"},
		{token.TypeRParen, ")"},
		{token.TypeLBrace, "{"},
		{token.TypeReturn, "return"},
		{token.TypeTrue, "true"},
		{token.TypeSemicolon, ";"},
		{token.TypeRBrace, "}"},
		{token.TypeElse, "else"},
		{token.TypeLBrace, "{"},
		{token.TypeReturn, "return"},
		{token.TypeFalse, "false"},
		{token.TypeSemicolon, ";"},
		{token.TypeRBrace, "}"},
		{token.TypeInt, "10"},
		{token.TypeEQ, "=="},
		{token.TypeInt, "10"},
		{token.TypeSemicolon, ";"},
		{token.TypeInt, "10"},
		{token.TypeNotEQ, "!="},
		{token.TypeInt, "9"},
		{token.TypeSemicolon, ";"},
		{token.TypeEOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
