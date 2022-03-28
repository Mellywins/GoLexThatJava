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
func NewLexer() *lexmachine.Lexer {
	lexer := lexmachine.NewLexer()

	Tokens := []string{
		"BANG",
		"INTEGER_LITERAL",
		"BOOLEAN_LITERAL",
		"IDENTIFIER",
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
	ReservedWordsTokens := []string{
		"CLASS",
		"PUBLIC",
		"STATIC",
		"VOID",
		"MAIN",
		"STRING",
		"SYSTEMOUTPRINTLN",
		"RETURN",
		"INT",
		"IF",
		"FOR",
		"ELSE",
		"WHILE",
		"THIS",
		"NEW",
		"BOOLEAN",
		"LENGTH",
		"EXTENDS",
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
	LiteralTokens := []string{
		"LEFTBRACKET",
		"RIGHTBRACKET",
		"LEFTANGLEBRACKET",
		"RIGHTANGLEBRACKET",
		"COMMA",
		"SEMICOLON",
		"COLON",
		"LEFTPARENTHESIS",
		"RIGHTPARENTHESIS",
		"PERIOD",
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
	OperatorTokens := []string{
		"PLUS",
		"ASTERIX",
		"DIVISION",
		"MODULO",
		"DOUBLEQUAL",
		"DIFFERENT",
		"LESS",
		"LESSOREQUALS",
		"GREATER",
		"GREATEROREQUALS",
		"LOGICALAND",
		"LOGICALOR",
		"EQUAL",
		"MINUS",
	}
	Tokens = append(Tokens, ReservedWordsTokens...)
	Tokens = append(Tokens, LiteralTokens...)
	Tokens = append(Tokens, OperatorTokens...)
	TokenIds = make(map[string]int)
	for i, tok := range Tokens {
		TokenIds[tok] = i
		fmt.Println(tok, ":", i)

	}

	// LEXER PATTERNS
	lexer.Add([]byte(`( |\t|\n)`), skip)
	lexer.Add([]byte(`[!]`), token("BANG"))
	lexer.Add([]byte(`[1-9][0-9]*`), token("INTEGER_LITERAL"))
	lexer.Add([]byte(`true|false`), token("BOOLEAN_LITERAL"))
	// lexer.Add([]byte(`\(`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
	// 	parenthesisCount := 1
	// 	match.EndLine = match.StartLine
	// 	match.EndColumn = match.StartColumn
	// 	for tc := scan.TC; tc < len(scan.Text)-1; tc++ {
	// 		tokenStr := scan.Text[tc]
	// 		match.EndColumn += 1
	// 		if scan.Text[tc] == '\n' {
	// 			match.EndLine += 1
	// 		}
	// 		if tokenStr == '(' {
	// 			parenthesisCount++
	// 		} else if tokenStr == ')' {
	// 			parenthesisCount--
	// 		}
	// 		if parenthesisCount == 0 {
	// 			return token("LEFTPARENTHESIS")(scan, match)
	// 		}
	// 	}
	// 	// fmt.Println("DEBUG LOG: ", "\n", "Parenthesis Count: ", parenthesisCount, "\n", "Match: ", match.TC, "\n", "Text Length", len(scan.Text))
	// 	return nil, fmt.Errorf(
	// 		"Error: Reached EOF without Unclosed Parenthesis: \"(\" starting at %d, (%d, %d)", match.TC, match.StartLine, match.StartColumn)
	// })
	// lexer.Add([]byte(`\{`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
	// 	bracersCount := 1
	// 	match.EndLine = match.StartLine
	// 	match.EndColumn = match.StartColumn
	// 	for tc := scan.TC; tc < len(scan.Text); tc++ {
	// 		tokenStr := scan.Text[tc]
	// 		match.EndColumn += 1
	// 		if scan.Text[tc] == '\n' {
	// 			match.EndLine += 1
	// 		}
	// 		if tokenStr == '{' {
	// 			bracersCount++
	// 		} else if tokenStr == '}' {
	// 			bracersCount--
	// 		}
	// 		if bracersCount == 0 {
	// 			return token("LEFTANGLEBRACKET")(scan, match)
	// 		}
	// 	}
	// 	// fmt.Println("DEBUG LOG: ", "\n", "Parenthesis Count: ", parenthesisCount, "\n", "Match: ", match.TC, "\n", "Text Length", len(scan.Text))
	// 	return nil, fmt.Errorf(
	// 		"Error: Reached EOF without Unclosed bracket: \"{\" starting at %d, (%d, %d)", match.TC, match.StartLine, match.StartColumn)
	// })
	for i, lit := range Literals {
		r := "\\" + strings.Join(strings.Split(lit, ""), "\\")
		lexer.Add([]byte(r), token(LiteralTokens[i]))
	}
	for i, name := range ReservedWords {
		lexer.Add([]byte(name), token(ReservedWordsTokens[i]))
	}
	for i, op := range Operators {
		formattedOp := "\\" + strings.Join(strings.Split(op, ""), "\\")
		lexer.Add([]byte(formattedOp), token(OperatorTokens[i]))
	}

	lexer.Add([]byte(`[A-Za-z_][A-Za-z0-9_]*`), token("IDENTIFIER"))
	lexer.Add([]byte(`//[^\n]*\n?`), skip)
	lexer.Add([]byte(`/\*([^*]|\r|\n|(\*+([^*/]|\r|\n)))*\*+/`), skip)
	// lexer.Add([]byte(`\/\*`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
	// 	commentIndicatorCount := 1
	// 	match.EndLine = match.StartLine
	// 	match.EndColumn = match.StartColumn
	// 	for tc := scan.TC; tc < len(scan.Text)-1; tc++ {
	// 		tokenStr := fmt.Sprintf("%b%b", scan.Text[tc], scan.Text[tc+1])
	// 		match.EndColumn += 1
	// 		if scan.Text[tc] == '\n' {
	// 			match.EndLine += 1
	// 		}
	// 		if tokenStr == "/*" {
	// 			commentIndicatorCount++
	// 		} else if tokenStr == "*/" {
	// 			commentIndicatorCount--
	// 		}
	// 		if commentIndicatorCount == 0 {
	// 			return token("MULTI_LINE_COMMENT")(scan, match)
	// 		}
	// 	}
	// 	return nil, fmt.Errorf(
	// 		"Error: Reached EOF without Unclosed comment starting at %d, (%d, %d)", match.TC, match.StartLine, match.StartColumn)
	// })
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
