package lexical

import (
	"fmt"
	"strings"

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

// func token(tokenType string) func(*lexmachine.Scanner, *machines.Match) (interface{}, error) {
// 	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
// 		return NewToken(tokenType, m), nil
// 	}
// }
func must(err error) {
	if err != nil {
		panic(err)
	}
}
func __init__() *lexmachine.Lexer {
	lexer := lexmachine.NewLexer()

	Tokens := []string{
		"SPACE",
		"TAB",
		"BANG",
		"INTEGER_LITERAL",
		"BOOLEAN_LITERAL",
		"IDENTIFIER",
		"SINGLE_LINE_COMMENT",
		"MULTI_LINE_COMMENT",
		"BREAK_LINE",
		"RESERVED_WORD",
	}
	ReservedWords := []string{
		"class",
		"public",
		"static",
		"void",
		"main",
		"String",
		"System.out.println",
		"return",
		"int",
		"if",
		"for",
		"else",
		"while",
		"this",
		"new",
		"boolean",
		"length",
		"extends",
	}
	Literals := []string{
		"[",
		"]",
		"{",
		"}",
		",",
		";",
		":",
		"(",
		")",
		".",
	}
	Operators := []string{
		"+",
		"*",
		"/",
		"%",
		"==",
		"!=",
		"<",
		"<=",
		">",
		">=",
		"&&",
		"||",
		"=",
		"-",
	}
	Tokens = append(Tokens, ReservedWords...)
	Tokens = append(Tokens, Literals...)
	Tokens = append(Tokens, Operators...)
	TokenIds = make(map[string]int)
	for i, tok := range Tokens {
		TokenIds[tok] = i
	}
	// LEXER PATTERNS
	lexer.Add([]byte(`	`), token("TAB"))
	lexer.Add([]byte(` `), token("SPACE"))
	lexer.Add([]byte(`[!]`), token("BANG"))
	lexer.Add([]byte(`[1-9][0-9]*`), token("INTEGER_LITERAL"))
	lexer.Add([]byte(`true|false`), token("BOOLEAN_LITERAL"))
	lexer.Add([]byte(`\(`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		parenthesisCount := 1
		match.EndLine = match.StartLine
		match.EndColumn = match.StartColumn
		for tc := scan.TC; tc < len(scan.Text)-1; tc++ {
			tokenStr := scan.Text[tc]
			match.EndColumn += 1
			if scan.Text[tc] == '\n' {
				match.EndLine += 1
			}
			if tokenStr == '(' {
				parenthesisCount++
			} else if tokenStr == ')' {
				parenthesisCount--
			}
			if parenthesisCount == 0 {
				return token("(")(scan, match)
			}
		}
		// fmt.Println("DEBUG LOG: ", "\n", "Parenthesis Count: ", parenthesisCount, "\n", "Match: ", match.TC, "\n", "Text Length", len(scan.Text))
		return nil, fmt.Errorf(
			"Error: Reached EOF without Unclosed Parenthesis: \"(\" starting at %d, (%d, %d)", match.TC, match.StartLine, match.StartColumn)
	})
	lexer.Add([]byte(`\{`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		bracersCount := 1
		match.EndLine = match.StartLine
		match.EndColumn = match.StartColumn
		for tc := scan.TC; tc < len(scan.Text); tc++ {
			tokenStr := scan.Text[tc]
			match.EndColumn += 1
			if scan.Text[tc] == '\n' {
				match.EndLine += 1
			}
			if tokenStr == '{' {
				bracersCount++
			} else if tokenStr == '}' {
				bracersCount--
			}
			if bracersCount == 0 {
				return token("{")(scan, match)
			}
		}
		// fmt.Println("DEBUG LOG: ", "\n", "Parenthesis Count: ", parenthesisCount, "\n", "Match: ", match.TC, "\n", "Text Length", len(scan.Text))
		return nil, fmt.Errorf(
			"Error: Reached EOF without Unclosed bracket: \"{\" starting at %d, (%d, %d)", match.TC, match.StartLine, match.StartColumn)
	})
	for _, lit := range Literals {
		r := "\\" + strings.Join(strings.Split(lit, ""), "\\")
		lexer.Add([]byte(r), token(lit))
	}
	for _, name := range ReservedWords {
		lexer.Add([]byte(name), token(name))
	}
	for _, op := range Operators {
		formattedOp := "\\" + strings.Join(strings.Split(op, ""), "\\")
		lexer.Add([]byte(formattedOp), token(op))
	}

	lexer.Add([]byte(`[A-Za-z_][A-Za-z0-9_]*`), token("IDENTIFIER"))
	lexer.Add([]byte(`//[^\n]*\n?`), token("SINGLE_LINE_COMMENT"))
	lexer.Add([]byte(`/\*([^*]|\r|\n|(\*+([^*/]|\r|\n)))*\*+/`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {

		return token("MULTI_LINE_COMMENT")(scan, match)
	})
	lexer.Add([]byte(`\/\*`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		commentIndicatorCount := 1
		match.EndLine = match.StartLine
		match.EndColumn = match.StartColumn
		for tc := scan.TC; tc < len(scan.Text)-1; tc++ {
			tokenStr := fmt.Sprintf("%b%b", scan.Text[tc], scan.Text[tc+1])
			match.EndColumn += 1
			if scan.Text[tc] == '\n' {
				match.EndLine += 1
			}
			if tokenStr == "/*" {
				commentIndicatorCount++
			} else if tokenStr == "*/" {
				commentIndicatorCount--
			}
			if commentIndicatorCount == 0 {
				return token("MULTI_LINE_COMMENT")(scan, match)
			}
		}
		return nil, fmt.Errorf(
			"Error: Reached EOF without Unclosed comment starting at %d, (%d, %d)", match.TC, match.StartLine, match.StartColumn)
	})
	lexer.Add([]byte(`\n`), token("BREAK_LINE"))
	lexer.Add([]byte(`.`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {

		return nil, fmt.Errorf("error: uknown element found at %d, (%d,%d)", match.TC, match.StartLine, match.StartColumn)
	})

	err := lexer.Compile()
	must(err)

	return lexer
}
func skip(*lexmachine.Scanner, *machines.Match) (interface{}, error) {
	return nil, nil
}

// a lexmachine.Action function with constructs a Token of the given token type by
// the token type's name.
func token(name string) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(TokenIds[name], string(m.Bytes), m), nil
	}
}