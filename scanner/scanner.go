package scanner

import (
	"unicode"

	"github.com/jmptc/golox/token"
)

var (
	keywords = make(map[string]string)
)

func init() {
	keywords["and"] = token.AND
	keywords["class"] = token.CLASS
	keywords["else"] = token.ELSE
	keywords["false"] = token.FALSE
	keywords["fun"] = token.FUN
	keywords["for"] = token.FOR
	keywords["if"] = token.IF
	keywords["nil"] = token.NIL
	keywords["or"] = token.OR
	keywords["print"] = token.PRINT
	keywords["return"] = token.RETURN
	keywords["super"] = token.SUPER
	keywords["this"] = token.THIS
	keywords["true"] = token.TRUE
	keywords["var"] = token.VAR
	keywords["while"] = token.WHILE

}

type Scanner struct {
	//source  string
	source  []rune
	start   int
	current int
	line    int
	tokens  []token.Token
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: []rune(source)}
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
	case '!':
		if s.match('=') {
			s.addToken(token.BANG_EQUAL)
		} else {
			s.addToken(token.BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(token.EQUAL_EQUAL)
		} else {
			s.addToken(token.EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(token.LESS_EQUAL)
		} else {
			s.addToken(token.LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(token.GREATER_EQUAL)
		} else {
			s.addToken(token.GREATER)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.AtEnd() {
				s.advance()
			}
		}
	case '\n':
		s.line += 1
	case ' ', '\t', '\r':
	case '"':
		s.tokenizeString()
	default:
		if unicode.IsDigit(c) {
			s.tokenizeNumber()
		} else if isAlphanumeric(c) {
			s.tokenizeIdentifier()
		}

	}
}

func (s *Scanner) advance() rune {
	b := s.source[s.current]
	s.current += 1
	return b
}

func (s *Scanner) AtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) addToken(tokenType string) {
	text := s.source[s.start:s.current]
	token := token.Token{TokenType: tokenType, Lexeme: string(text), Line: s.line}
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) addTokenTypeAndVal(tokenType, val string) {
	token := token.Token{TokenType: tokenType, Lexeme: val, Line: s.line}
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) match(expected rune) bool {
	if s.AtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}

	s.current += 1
	return true
}

func (s *Scanner) peek() rune {
	if s.AtEnd() {
		return '\x00'
	} else {
		return s.source[s.current]
	}

}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return '\x00'
	}

	return s.source[s.current+1]
}

func (s *Scanner) tokenizeString() {
	for s.peek() != '"' && !s.AtEnd() {
		if s.peek() == '\n' {
			s.line += 1
		}
		s.advance()
	}

	if s.AtEnd() {
		return
	}

	s.advance()

	value := s.source[s.start+1 : s.current-1]
	s.addTokenTypeAndVal(token.STRING, string(value))
}

func (s *Scanner) tokenizeNumber() {
	for unicode.IsDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && unicode.IsDigit(s.peekNext()) {
		s.advance()

		for unicode.IsDigit(s.peek()) {
			s.advance()
		}
	}

	s.addToken(token.NUMBER)
}

func (s *Scanner) tokenizeIdentifier() {
	for isAlphanumeric(s.peek()) {
		s.advance()
	}

	text := string(s.source[s.start : s.current])
	keyword, ok := keywords[text]
	if !ok {
		s.addTokenTypeAndVal(token.IDENTIFIER, text)
	} else {
		s.addTokenTypeAndVal(keyword, text)
	}
}

func isAlphanumeric(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '_'
}
