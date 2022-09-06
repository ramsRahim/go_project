package main

import (
	"os"
	"path"
	"testing"
)

func TestLoadLexicon(t *testing.T) {
	tests := []struct {
		n    string
		want string
	}{
		{"COME", "k V m"},
		{"WORDS", "w 3` d z"},
		{"MECCA", "m E k @"},
	}
	dummy := make(map[string]string)
	dummy["COME"] = "k V m"
	dummy["WORDS"] = "w 3` d z"
	dummy["MECCA"] = "m E k @"

	str := ""
	for k, v := range dummy {
		str += k + "\t" + v + "\n"
	}
	lex_path := path.Join(t.TempDir(), "lexicon.txt")
	err := os.WriteFile(lex_path, []byte(str), 0644)

	if err != nil {
		t.Fatalf("unexpected error when writing lexicon file, err=%v", err)
	}

	for _, tc := range tests {
		if got, _ := LoadLexicon(lex_path); got[tc.n] != tc.want {
			t.Errorf("got %s, want %s", got[tc.n], tc.want)
		}
	}
}
