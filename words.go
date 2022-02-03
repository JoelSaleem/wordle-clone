package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Words struct {
	arr []string
	set map[string]bool
}

func (words *Words) LoadWords() {
	bytes, _ := ioutil.ReadFile("dictionary")
	data := string(bytes)

	words.arr = strings.Split(data, "\n")
	for i, w := range words.arr {
		words.arr[i] = strings.ToLower(w)
	}

	words.set = make(map[string]bool)
	for _, w := range words.arr {
		words.set[w] = true
	}
}

func (words *Words) getChoice() string {
	w := words.arr

	rand.Seed(time.Now().UnixNano())
	return w[rand.Intn(len(w))]
}

func (words *Words) IsInputInDict(input string) bool {
	return words.set[input]
}
