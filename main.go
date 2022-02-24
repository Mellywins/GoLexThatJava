package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/timtadh/lexmachine"
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
	outputFile, err := os.Create("output.java")
	must(err)
	defer outputFile.Close()
	ioutil.WriteFile(outputFile.Name(), []byte(`package main;`), os.ModePerm)
	openedFile, err := os.OpenFile("output.java", os.O_APPEND|os.O_WRONLY, os.ModePerm)
	must(err)
	openedFile.WriteString("\n")
	for tok, err, eos := scanner.Next(); !eos; tok, err, eos = scanner.Next() {
		if ui, is := err.(*machines.UnconsumedInput); is {
			// skip the error via:
			scanner.TC = ui.FailTC

			// return err
		}
		must(err)
		k, _ := mapkey(TokenIds, tok.(*lexmachine.Token).Type)

		// fmt.Print(k, ":", tok.(*Token).Lexeme, " ")
		if k == "SPACE" {
			_, err := openedFile.WriteString(" ")
			must(err)
			fmt.Print(" ")
		} else if k == "BREAK_LINE" {
			fmt.Println()
			_, err := openedFile.WriteString("\n")
			must(err)
			// } else if k == "TAB" {
			// 	fmt.Print("\t")
			// 	_, err := openedFile.WriteString("\t")
			// 	must(err)
		} else {
			fmt.Print(k)
			_, err := openedFile.WriteString(k)
			must(err)

		}
	}
}
