package main

import (
	"fmt"

	tm "github.com/buger/goterm"
)

type Guess string

type Guesses []Guess

func (g *Guesses) AddWord(w Guess) {
	if len(*g) >= 6 {
		return
	}

	*g = append(*g, w)
}

func (guesses *Guesses) Print(choice string) {
	for _, g := range *guesses {
		out := ""
		for i, ch := range g {
			if string(choice[i]) == string(ch) {
				out += tm.Color(string(ch)+" ", tm.GREEN)
			} else if isLetterInWord(string(ch), choice) {
				out += tm.Color(string(ch)+" ", tm.YELLOW)
			} else {
				out += tm.Color(string(ch)+" ", tm.WHITE)
			}
		}
		fmt.Println(out)
	}
}
