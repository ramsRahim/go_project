package main

import (
	"fmt"
	"strings"

	"github.com/ramsRahim/go_project/lexicon"
)

func main() {
	var lexiconPath string
	fmt.Scanln(&lexiconPath)

	lexWords, err := lexicon.LoadLexicon(lexiconPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var word string
	fmt.Scanln(&word)
	word = strings.TrimSuffix(word, "\n")
	l := lexicon.Lexicon{Lex: lexWords}
	pronunciation := l.LookUp(word)
	_ = l.Remove(word)
	_ = l.LookUp(word)
	_ = l.Add(word, pronunciation)
	_ = l.LookUp(word)
}
