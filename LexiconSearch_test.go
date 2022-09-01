package main

import "testing"

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

	dir := t.TempDir()

	for _, tc := range tests {
		if got, _ := LoadLexicon(dir); got[tc.n] != tc.want {
			t.Errorf("got %s, want %s", got[tc.n], tc.want)
		}
	}
}
