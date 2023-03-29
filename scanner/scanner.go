package scanner

import (
    "github.com/jmptc/golox/token"
)

type Scanner struct {
    source string
    start int
    current int
    line int
}

func NewScanner(source string) *Scanner {
    return &Scanner{ source: source }
}

func (s *Scanner) ScanTokens() []token.Token {
    tokens := make([]token.Token, 0)
    return tokens
}

func (s *Scanner) AtEnd() bool {
    return s.current >= len(s.source)
}
