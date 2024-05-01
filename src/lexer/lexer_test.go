package lexer

import (
	"belle/src/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "+=(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.LBRACE, "("},
		{token.RBRACE, ")"},
		{token.LSQUIRLY, "{"},
		{token.RSQUIRLY, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestProgram(t *testing.T) {
	input := `let x : Int = -10;
	let mut x2 : UInt = 420;
	x2 = 69;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LetToken, "let"},
		{token.IDENT, "x"},
		{token.COLON, ":"},
		{token.Int, "Int"},
		{token.ASSIGN, "="},
		{token.LITERAL, "-10"},
		{token.SEMICOLON, ";"},

		{token.LetToken, "let"},
		{token.MutToken, "mut"},
		{token.IDENT, "x2"},
		{token.COLON, ":"},
		{token.UInt, "UInt"},
		{token.ASSIGN, "="},
		{token.LITERAL, "420"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "x2"},
		{token.ASSIGN, "="},
		{token.LITERAL, "69"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

/*
func TestFunction(t *testing.T) {
	input := `let Func : Procedure = {
		return "hallo";
	}
	Func();`

	l := New(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LetToken, "let"},
		{token.IDENT, "Func"},
		{token.COLON, ":"},
		{token.Procedure, "Procedure"},
		{token.ASSIGN, "="},
		{token.LSQUIRLY, "{"},
		{token.ReturnToken, "return"},
		{token.DOUBLEQUOTE, "\""},
		{token.LITERAL, "hallo"},
		{token.DOUBLEQUOTE, "\""},
		{token.SEMICOLON, ";"},
		{token.RSQUIRLY, "}"},

		{token.IDENT, "Func"},
		{token.LBRACE, "("},
		{token.RBRACE, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
*/
