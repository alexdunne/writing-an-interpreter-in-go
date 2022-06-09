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
			}},
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
