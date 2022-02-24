package main

import (
	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

type Token struct {
	TokenType int
	Lexeme    string
	Match     *machines.Match
}

var TokenIds map[string]int

func NewToken(tokenType string, m *machines.Match) *Token {
	return &Token{
		TokenType: TokenIds[tokenType], // defined above
		Lexeme:    string(m.Bytes),
		Match:     m,
	}
}
func token(tokenType string) func(*lexmachine.Scanner, *machines.Match) (interface{}, error) {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		return NewToken(tokenType, m), nil
	}
}
func must(err error) {
	if err != nil {
		panic(err)
	}
}
func __init__() *lexmachine.Lexer {
	lexer := lexmachine.NewLexer()

	Tokens := []string{
		"SPACE",
		"BANG",
		"INTEGER_LITERAL",
		"BOOLEAN_LITERAL",
		"IDENTIFIER",
	}
	TokenIds = make(map[string]int)
	for i, tok := range Tokens {
		TokenIds[tok] = i
	}
	// LEXER PATTERNS
	lexer.Add([]byte(` +`), token("SPACE"))
	lexer.Add([]byte(`[!]`), token("BANG"))
	lexer.Add([]byte(`-?[1-9][0-9]*`), token("INTEGER_LITERAL"))
	lexer.Add([]byte(`true|false`), token("BOOLEAN_LITERAL"))
	lexer.Add([]byte(`[A-Za-z_][A-Za-z0-9_]*`), token("IDENTIFIER"))

	err := lexer.Compile()
	must(err)

	return lexer
}
