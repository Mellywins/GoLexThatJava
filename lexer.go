package main

import (
	"fmt"

	lexical "github.com/Mellywins/GoLexThatJava/lexical"
	"github.com/timtadh/lexmachine"
)

type golex struct {
	*lexmachine.Scanner
	stmts []*Node
}

func newGoLex(lexer *lexmachine.Lexer, text []byte) (*golex, error) {
	scan, err := lexer.Scanner(text)
	if err != nil {
		return nil, err
	}
	return &golex{Scanner: scan}, nil
}

// Lex implements yyLexer's interface for getting the next token. It returns the
// token type as an integer. The tokens should be defined in the $parser.y file.
// The actual number returned will be >= yyPrivate - 1 which is the range for
// custom token names.
func (g *golex) Lex(lval *yySymType) (tokenType int) {
	s := g.Scanner
	tok, err, eof := s.Next()
	if err != nil {
		fmt.Println("token that failed: ", tok)
		g.Error(err.Error())
	} else if eof {
		return -1 // signals EOF to goyacc's yyParse
	}
	lval.token = tok.(*lexmachine.Token)
	// To return the correct number for goyacc you must add yyPrivate - 1 to
	// put the value into the correct range.
	return lval.token.Type + yyPrivate + 2
}

// Error implements the error handling for if there is a parse error of any
// kind. This implementation panics. There may be no better way to hand errors
// from goyacc. I recommend you use defer ... recover() to handle this where
// you call into the parser.
func (l *golex) Error(message string) {
	// is there a better way to handle this in the context of goyacc?
	panic(fmt.Errorf(message))
}
func newLexer() *lexmachine.Lexer {
	return lexical.NewLexer()
}
