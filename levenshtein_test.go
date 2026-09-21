package main

import (
	"testing"
)

func TestLevenshtein(t *testing.T) {
	got := Levenshtein("когда", "когда")
	want := 0
	if got != want {
		t.Errorf("Levenshtein('когда', 'когда') = %d; want %d", got, want)
	}
}
