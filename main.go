package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tm "github.com/buger/goterm"
)

func main() {
	tm.Clear()
	tm.MoveCursor(1, 1)

	dict := Words{}
	dict.LoadWords()

	choice := dict.getChoice()
	// fmt.Println(choice)

	guesses := Guesses{}
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\n\n")
		guesses.Print(choice)

		if len(guesses) >= 6 {
			fmt.Println("you failed: " + choice)

			return
		}

		w, _ := r.ReadString('\n')
		w = strings.Replace(w, "\n", "", -1)
		if len(w) != 5 || !dict.IsInputInDict(w) {
			continue
		}
		if w == choice {
			guesses.Print(choice)
			fmt.Println("congrats")
			return
		}

		guesses.AddWord(Guess(w))

	}
}
