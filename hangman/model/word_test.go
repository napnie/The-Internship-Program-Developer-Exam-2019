package model

import (
	"testing"
)

func TestWord(t *testing.T) {
	word := "banana"
	hint := "tropical fruit"

	hang := NewHangmanWord(word, hint)

	if hang.Word != word {
		t.Errorf("expected word: %s, but get: %s", word, hang.Word)
	}
	if hang.Hint != hint {
		t.Errorf("expected hint: %s, but get: %s", word, hang.Hint)
	}
	if hang.Tried != 0 {
		t.Errorf("expected hangword tried: %d, but get: %d", 0, hang.Tried)
	}

	if hang.Show() != "_ _ _ _ _ _" {
		t.Errorf("expected hangword to show: %s, but get: %s", "_ _ _ _ _ _", hang.Show())
	}

	hang.Guess("a")
	if hang.Show() != "_ a _ a _ a" {
		t.Errorf("expected hangword to show: %s, but get: %s", "_ a _ a _ a", hang.Show())
	}
}
