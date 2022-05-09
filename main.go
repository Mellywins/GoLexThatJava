package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/Mellywins/GoLexThatJava/lexical"
	"github.com/timtadh/getopt"
	"github.com/timtadh/lexmachine"
)

func main() {
	short := "hv"
	long := []string{
		"help",
		"verbose",
	}
	_, optargs, err := getopt.GetOpt(os.Args[1:], short, long)
	if err != nil {
		log.Print(err)
		log.Println("try --help")
		os.Exit(1)
	}

	for _, oa := range optargs {
		switch oa.Opt() {
		case "-h", "--help":
			fmt.Println("parse a java file")
			os.Exit(0)
		case "-v", "--verbose":
			lexical.PrintLexing()
		}

	}

	lexer := newLexer()

	stmts, err := parse(lexer, os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, stmt := range stmts {
		fmt.Println(stmt)
	}
	for _, v := range CheckVariableHierarchyQueue {
		FindHealthyOccurenceInParentContexts(v, v.Token.Value.(string))
	}
}

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
	text, err := ioutil.ReadFile("test/example.java")
	if err != nil {
		return nil, err
	}
	scanner, err := newGoLex(lexer, text)
	if err != nil {
		return nil, err
	}
	yyParse(scanner)
	return scanner.stmts, nil
}
