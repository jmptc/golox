package scanner

import (
	"testing"

	"github.com/jmptc/golox/token"
)

func TestAtEndOfSource(t *testing.T) {
	scanner := NewScanner("Hello")

	// set current right before end
	scanner.current = len(scanner.source) - 1
	if scanner.AtEnd() {
		t.Errorf("Not expecting AtEnd to be true here. current: %d len: %d", scanner.current, len(scanner.source))
	}

	// set current after source
	scanner.current = len(scanner.source)
	if !scanner.AtEnd() {
		t.Errorf("Expected AtEnd to be true. current: %d len: %d", scanner.current, len(scanner.source))
	}
}

func TestScanSingleCharacterTokens(t *testing.T) {
	text := "(){},.-+;*"

	expectedTokens := []token.Token{
		{TokenType: token.LEFT_PAREN, Lexeme: "(", Line: 0},
		{TokenType: token.RIGHT_PAREN, Lexeme: ")", Line: 0},
		{TokenType: token.LEFT_BRACE, Lexeme: "{", Line: 0},
		{TokenType: token.RIGHT_BRACE, Lexeme: "}", Line: 0},
		{TokenType: token.COMMA, Lexeme: ",", Line: 0},
		{TokenType: token.DOT, Lexeme: ".", Line: 0},
		{TokenType: token.MINUS, Lexeme: "-", Line: 0},
		{TokenType: token.PLUS, Lexeme: "+", Line: 0},
		{TokenType: token.SEMICOLON, Lexeme: ";", Line: 0},
		{TokenType: token.STAR, Lexeme: "*", Line: 0},
		{TokenType: token.EOF, Lexeme: "", Line: 0},
	}
	compareExpected(t, text, expectedTokens)
}

func TestConditionalTwoCharacterTokens(t *testing.T) {
	text := "! !=  = == < <= > >="

	expectedTokens := []token.Token{
		{TokenType: token.BANG, Lexeme: "!", Line: 0},
		{TokenType: token.BANG_EQUAL, Lexeme: "!=", Line: 0},
		{TokenType: token.EQUAL, Lexeme: "=", Line: 0},
		{TokenType: token.EQUAL_EQUAL, Lexeme: "==", Line: 0},
		{TokenType: token.LESS, Lexeme: "<", Line: 0},
		{TokenType: token.LESS_EQUAL, Lexeme: "<=", Line: 0},
		{TokenType: token.GREATER, Lexeme: ">", Line: 0},
		{TokenType: token.GREATER_EQUAL, Lexeme: ">=", Line: 0},
		{TokenType: token.EOF, Lexeme: "", Line: 0},
	}

	compareExpected(t, text, expectedTokens)
}

func TestComments(t *testing.T) {
	text := `! // This is a comment
		// Here's another comment
		( )`

	expectedTokens := []token.Token{
		{TokenType: token.BANG, Lexeme: "!", Line: 0},
		{TokenType: token.LEFT_PAREN, Lexeme: "(", Line: 2},
		{TokenType: token.RIGHT_PAREN, Lexeme: ")", Line: 2},
		{TokenType: token.EOF, Lexeme: "", Line: 2},
	}

	compareExpected(t, text, expectedTokens)
}

func TestStringTokenization(t *testing.T) {
	text := `"hello world"`

	expectedTokens := []token.Token{
		{TokenType: token.STRING, Lexeme: "hello world", Line: 0},
		{TokenType: token.EOF, Lexeme: "", Line: 0},
	}

	compareExpected(t, text, expectedTokens)
}

func TestNumberTokenization(t *testing.T) {
	text := "12.34 2. .37"

	expectedTokens := []token.Token{
		{TokenType: token.NUMBER, Lexeme: "12.34", Line: 0},
		{TokenType: token.NUMBER, Lexeme: "2", Line: 0},
		{TokenType: token.DOT, Lexeme: ".", Line: 0},
		{TokenType: token.DOT, Lexeme: ".", Line: 0},
		{TokenType: token.NUMBER, Lexeme: "37", Line: 0},
		{TokenType: token.EOF, Lexeme: "", Line: 0},
	}

	compareExpected(t, text, expectedTokens)
}

func TestIdentifierAndKeywordsTokenization(t *testing.T) {
	text := "var hello = \"hi there\" if else testing"

	expectedTokens := []token.Token{
		{TokenType: token.VAR, Lexeme: "var", Line: 0},
		{TokenType: token.IDENTIFIER, Lexeme: "hello", Line: 0},
		{TokenType: token.EQUAL, Lexeme: "=", Line: 0},
		{TokenType: token.STRING, Lexeme: "hi there", Line: 0},
		{TokenType: token.IF, Lexeme: "if", Line: 0},
		{TokenType: token.ELSE, Lexeme: "else", Line: 0},
		{TokenType: token.IDENTIFIER, Lexeme: "testing", Line: 0},
		{TokenType: token.EOF, Lexeme: "", Line: 0},
	}

	compareExpected(t, text, expectedTokens)
}

// creates scanner with source and compares output with expected tokens
func compareExpected(t *testing.T, source string, expectedTokens []token.Token) {
	s := NewScanner(source)
	tokens := s.ScanTokens()

	if len(expectedTokens) != len(tokens) {
		t.Errorf("Token count mismatch. Expected: %d Got: %d", len(expectedTokens), len(tokens))
	}

	for i, tok := range tokens {
		expectedTok := expectedTokens[i]
		if tok.TokenType != expectedTok.TokenType {
			t.Errorf("TokenType mismatch. Expected: %s Got: %s", expectedTok.TokenType, tok.TokenType)
		}

		if tok.Lexeme != expectedTok.Lexeme {
			t.Errorf("Lexeme mismatch. Expected: %s Got: %s", expectedTok.Lexeme, tok.Lexeme)
		}

		if tok.Line != expectedTok.Line {
			t.Errorf("Line mismatch: Expected: %d Got: %d", expectedTok.Line, tok.Line)
		}
	}
}
