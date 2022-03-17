package main

import (
	"io"
	"io/ioutil"

	"github.com/timtadh/lexmachine"
)

func parse(lexer *lexmachine.Lexer, fin io.Reader) (stmts []*Node, err error) {
	defer func() {
		if e := recover(); e != nil {
			switch e.(type) {
			case error:
				err = e.(error)
				stmts = nil
			default:
				panic(e)
			}
		}
	}()
	text, err := ioutil.ReadAll(fin)
	if err != nil {
		return nil, err
	}
	scanner, err := newGoLex(lexer, text)
	if err != nil {
		return nil, err
	}
	yyParse(scanner)
	return scanner.stmts, nil
	return nil, nil
}
