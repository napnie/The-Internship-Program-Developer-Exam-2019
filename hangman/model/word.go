package model

import (
	"strings"
	"unicode"
)

//HangmanWord keep data of a word
type HangmanWord struct {
	Word  string
	Hint  string
	Tried int
	Known []rune
}

//NewHangmanWord create new HangmanWord instance
func NewHangmanWord(word, hint string) HangmanWord {
	return HangmanWord{
		Word:  strings.ToLower(word),
		Hint:  hint,
		Tried: 0,
	}
}

//Guess guess the letter of the word
func (word *HangmanWord) Guess(check string) bool {
	word.Tried++
	if len(check) != 1 {
		return false
	}
	return word.guessRune([]rune(check)[0])
}

//GuessRune helper method of Guess
func (word *HangmanWord) guessRune(r rune) bool {
	if !unicode.IsLetter(r) {
		return false
	}

	if strings.ContainsRune(word.Word, r) {
		if word.IsKnown(string(r)) { // check if already know this letter
			return false
		}
		word.Known = append(word.Known, r)
		return true
	}
	return false
}

//IsKnown check if s letter is already known
func (word HangmanWord) IsKnown(s string) bool {
	for _, r := range word.Known {
		if strings.ContainsRune(s, r) || !unicode.IsLetter([]rune(s)[0]) /* check if this is non character */ {
			return true
		}
	}
	return false
}

//Show show a state of a word
func (word HangmanWord) Show() string {
	show := strings.Split(word.Word, "")
	for i := 0; i < len(show); i++ {
		if !word.IsKnown(show[i]) {
			show[i] = "_"
		}
	}

	return strings.Join(show, " ")
}

//IsComplete check if hangman word is complete
func (word HangmanWord) IsComplete() (bool, string) {
	show := word.Show()
	if !strings.ContainsAny(show, "_") {
		return true, show
	}
	return false, show
}
