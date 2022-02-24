package main

import (
	"fmt"
	"io/ioutil"

	"github.com/timtadh/lexmachine/machines"
)

func mapkey(m map[string]int, value int) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}
func main() {
	fmt.Println("---------- START ----------")
	defer fmt.Println("\n---------- END ----------")
	myLexer := __init__()
	file, err := ioutil.ReadFile("example.java")
	must(err)
	scanner, err := myLexer.Scanner(file)
	must(err)
	for tok, err, eos := scanner.Next(); !eos; tok, err, eos = scanner.Next() {
		if ui, is := err.(*machines.UnconsumedInput); is {
			// skip the error via:
			scanner.TC = ui.FailTC

			// return err
		}
		must(err)
		k, _ := mapkey(TokenIds, tok.(*Token).TokenType)
		// fmt.Print(k, ":", tok.(*Token).Lexeme, " ")
		if k == "SPACE" {
			fmt.Print(" ")
		} else if k == "BREAK_LINE" {
			fmt.Println()
		} else if k == "TAB" {
			fmt.Print("\t")
		} else {
			fmt.Print(k)
		}
	}
}
