package main

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/ramsRahim/go_project/lexicon"
)

func writeDummyLexicon(t *testing.T) string {
	t.Helper()

	dummy := make(map[string]string)
	dummy["COME"] = "k V m"
	dummy["WORDS"] = "w 3` d z"
	dummy["MECCA"] = "m E k @"

	str := ""
	for k, v := range dummy {
		str += fmt.Sprintf("%s\t%s\n", k, v)
	}
	lexPath := path.Join(t.TempDir(), "lexicon.txt")
	err := os.WriteFile(lexPath, []byte(str), 0644)

	if err != nil {
		t.Fatalf("unexpected error when writing lexicon file, err=%v", err)
	}

	return lexPath
}

func TestLoadLexicon(t *testing.T) {
	tests := []struct {
		n    string
		want string
	}{
		{"COME", "k V m"},
		{"WORDS", "w 3` d z"},
		{"MECCA", "m E k @"},
	}

	lexPath := writeDummyLexicon(t)
	for _, tc := range tests {
		if got, _ := lexicon.LoadLexicon(lexPath); got[tc.n] != tc.want {
			t.Errorf("got %s, want %s", got[tc.n], tc.want)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"MECCA", true},
		{"RAHIM", false},
	}
	lexPath := writeDummyLexicon(t)
	lexWords, _ := lexicon.LoadLexicon(lexPath)
	l := lexicon.Lexicon{Lex: lexWords}
	for _, tc := range tests {
		if got := l.Remove(tc.word); got != tc.want {
			t.Errorf("got %t, want %t", got, tc.want)
		}
	}
}

func TestLookUp(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"COME", "k V m"},
		{"WORDS", "w 3` d z"},
		{"MECCA", "m E k @"},
	}
	lexPath := writeDummyLexicon(t)
	lexWords, _ := lexicon.LoadLexicon(lexPath)
	l := lexicon.Lexicon{Lex: lexWords}
	for _, tc := range tests {
		if got := l.LookUp(tc.word); got != tc.want {
			t.Errorf("got %s, want %s", got, tc.want)
		}
	}
}
