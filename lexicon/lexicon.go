package lexicon

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Lexicon struct {
	Lex map[string]string
}

func (l Lexicon) LookUp(word string) string {
	elem, ok := l.Lex[strings.ToUpper(word)]
	if ok {
		fmt.Printf("The pronunciation of the word %s is %s\n", word, elem)
	} else {
		fmt.Printf("The word doesn't exist\n")
	}

	return elem
}

func (l Lexicon) Remove(word string) bool {
	_, ok := l.Lex[strings.ToUpper(word)]
	if ok {
		delete(l.Lex, strings.ToUpper(word))
		return true
	} else {
		fmt.Printf("The word doesn't exist\n")
		return false
	}
}

func (l *Lexicon) Add(word string, pronunciation string) bool {
	l.Lex[strings.ToUpper(word)] = pronunciation
	return true
}

func LoadLexicon(lexiconPath string) (map[string]string, error) {
	m := make(map[string]string)
	lexiconPath = strings.TrimSuffix(lexiconPath, "\n")
	readFile, err := os.Open(lexiconPath)
	if err != nil {
		fmt.Println(err)
		return m, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	lex_words := make(map[string]string)

	for _, line := range fileLines {
		temp := strings.Split(line, "\t")
		lex_words[strings.ToUpper(temp[0])] = temp[1]
	}

	return lex_words, err

}
