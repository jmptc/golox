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
	/*
	scanner := NewScanner(text)
	tokens := scanner.ScanTokens()

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
	*/
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
