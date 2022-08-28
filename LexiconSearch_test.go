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
	for _, tc := range tests {
		if got, _ := LoadLexicon("nil"); got[tc.n] != tc.want {
			t.Errorf("got %s, want %s", got[tc.n], tc.want)
		}
	}
}
