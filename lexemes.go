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
		"SINGLE_LINE_COMMENT",
		"MULTI_LINE_COMMENT",
		"AFFECTATION",
		"BREAK_LINE",
		"RESERVED_WORD",
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
	lexer.Add([]byte(`class|public|static|void|main|String|extends|int|boolean|if|else|while|System.out.println|length|this|new`), token("RESERVED_WORD"))
	lexer.Add([]byte(`[A-Za-z_][A-Za-z0-9_]*`), token("IDENTIFIER"))
	lexer.Add([]byte(`//[^\n]*`), token("SINGLE_LINE_COMMENT"))
	lexer.Add([]byte(`/\*[^\*]*\*/`), token("MULTI_LINE_COMMENT"))
	lexer.Add([]byte(`[=]`), token("AFFECTATION"))
	lexer.Add([]byte(`\n`), token("BREAK_LINE"))

	err := lexer.Compile()
	must(err)

	return lexer
}
