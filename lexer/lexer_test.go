package lexer

import (
	"testing"

	"github.com/alexdunne/writing-an-interpreter-in-go/token"
)

func TestNextToken(t *testing.T) {

	type expectedToken struct {
		tokenType token.TokenType
		literal   string
	}

	tests := []struct {
		name           string
		input          string
		expectedTokens []expectedToken
	}{
		{
			name:  "symbols",
			input: `=+(){},;`,
			expectedTokens: []expectedToken{
				{tokenType: token.ASSIGN, literal: "="},
				{tokenType: token.PLUS, literal: "+"},
				{tokenType: token.LPAREN, literal: "("},
				{tokenType: token.RPAREN, literal: ")"},
				{tokenType: token.LBRACE, literal: "{"},
				{tokenType: token.RBRACE, literal: "}"},
				{tokenType: token.COMMA, literal: ","},
				{tokenType: token.SEMICOLON, literal: ";"},
			},
		},
		{
			name: "simple example",
			input: `let five = 5;
			let ten = 10;
			   let add = fn(x, y) {
				 x + y;
			};
			   let result = add(five, ten);
			`,
			expectedTokens: []expectedToken{
				{tokenType: token.LET, literal: "let"},
				{tokenType: token.IDENT, literal: "five"},
				{tokenType: token.ASSIGN, literal: "="},
				{tokenType: token.INT, literal: "5"},
				{tokenType: token.SEMICOLON, literal: ";"},
				{tokenType: token.LET, literal: "let"},
				{tokenType: token.IDENT, literal: "ten"},
				{tokenType: token.ASSIGN, literal: "="},
				{tokenType: token.INT, literal: "10"},
				{tokenType: token.SEMICOLON, literal: ";"},
				{tokenType: token.LET, literal: "let"},
				{tokenType: token.IDENT, literal: "add"},
				{tokenType: token.ASSIGN, literal: "="},
				{tokenType: token.FUNCTION, literal: "fn"},
				{tokenType: token.LPAREN, literal: "("},
				{tokenType: token.IDENT, literal: "x"},
				{tokenType: token.COMMA, literal: ","},
				{tokenType: token.IDENT, literal: "y"},
				{tokenType: token.RPAREN, literal: ")"},
				{tokenType: token.LBRACE, literal: "{"},
				{tokenType: token.IDENT, literal: "x"},
				{tokenType: token.PLUS, literal: "+"},
				{tokenType: token.IDENT, literal: "y"},
				{tokenType: token.SEMICOLON, literal: ";"},
				{tokenType: token.RBRACE, literal: "}"},
				{tokenType: token.SEMICOLON, literal: ";"},
				{tokenType: token.LET, literal: "let"},
				{tokenType: token.IDENT, literal: "result"},
				{tokenType: token.ASSIGN, literal: "="},
				{tokenType: token.IDENT, literal: "add"},
				{tokenType: token.LPAREN, literal: "("},
				{tokenType: token.IDENT, literal: "five"},
				{tokenType: token.COMMA, literal: ","},
				{tokenType: token.IDENT, literal: "ten"},
				{tokenType: token.RPAREN, literal: ")"},
				{tokenType: token.SEMICOLON, literal: ";"},
				{tokenType: token.EOF, literal: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.input)

			for _, expectedToken := range tt.expectedTokens {
				tok := l.NextToken()

				if tok.Type != expectedToken.tokenType {
					t.Fatalf("expected type %s, got %s", expectedToken.tokenType, tok.Type)
				}

				if tok.Literal != expectedToken.literal {
					t.Fatalf("expected literal %s, got %s", expectedToken.literal, tok.Type)
				}
			}

		})
	}
}
