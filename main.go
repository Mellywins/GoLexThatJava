package main

import (
	"fmt"

	"github.com/timtadh/lexmachine/machines"
)

func main() {
	fmt.Println("---------- START ----------")
	myLexer := __init__()
	scanner, err := myLexer.Scanner([]byte("wild!     "))
	must(err)
	for tok, err, eos := scanner.Next(); !eos; tok, err, eos = scanner.Next() {
		if ui, is := err.(*machines.UnconsumedInput); is {
			// skip the error via:
			scanner.TC = ui.FailTC

			// return err
		}
		must(err)
		fmt.Println(tok)
	}
}
