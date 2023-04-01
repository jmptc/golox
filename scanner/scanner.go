package scanner

import (
	"github.com/jmptc/golox/token"
)

type Scanner struct {
	source  string
	start   int
	current int
	line    int
	tokens  []token.Token
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: source}
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.AtEnd() {
		s.scanToken()
		s.start = s.current
	}

	s.addToken(token.EOF)
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {
	case '(':
		s.addToken(token.LEFT_PAREN)
	case ')':
		s.addToken(token.RIGHT_PAREN)
	case '{':
		s.addToken(token.LEFT_BRACE)
	case '}':
		s.addToken(token.RIGHT_BRACE)
	case ',':
		s.addToken(token.COMMA)
	case '.':
		s.addToken(token.DOT)
	case '-':
		s.addToken(token.MINUS)
	case '+':
		s.addToken(token.PLUS)
	case ';':
		s.addToken(token.SEMICOLON)
	case '*':
		s.addToken(token.STAR)
	}
}

func (s *Scanner) advance() byte {
	b := s.source[s.current]
	s.current += 1
	return b
}

func (s *Scanner) AtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) addToken(tokenType string) {
	text := s.source[s.start:s.current]
	token := token.Token{TokenType: tokenType, Lexeme: text, Line: s.line}
	s.tokens = append(s.tokens, token)
}
