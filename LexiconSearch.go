package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func main() {

	var lexiconPath string
	fmt.Scanln(&lexiconPath)

	lex_words, err := LoadLexicon(lexiconPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var word string

	fmt.Scanln(&word)
	word = strings.TrimSuffix(word, "\n")

	elem, ok := lex_words[strings.ToUpper(word)]

	if ok {
		fmt.Printf("The pronunciation of the word %s is %s\n", word, elem)
	} else {
		fmt.Printf("The word doesn't exist\n")
	}

}
