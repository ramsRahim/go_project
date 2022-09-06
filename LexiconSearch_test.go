package main

import (
	"fmt"
	"os"
	"path"
	"testing"
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
		if got, _ := LoadLexicon(lexPath); got[tc.n] != tc.want {
			t.Errorf("got %s, want %s", got[tc.n], tc.want)
		}
	}
}
