package bin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"../model"
)

func Menu() {
	fmt.Println("Hangman Game::")
	fmt.Println("Select Category:")

	// Get Catogory from csv file
	files := ListCSVFile()
	var catList strings.Builder
	catLen := len(files)
	for i := 0; i < catLen; i++ {
		category := strings.TrimSuffix(files[i], ".csv")

		catList.WriteString(
			fmt.Sprintf("%d. %s\n", i+1, category),
		)
	}
	fmt.Println(catList.String())
	fmt.Println(": Please input the number of category you wish to play")

	fmt.Print("\n > ")

	// Select catogory
	var cat int
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		cat, err = strconv.Atoi(input)
		if err == nil {
			// Check if selected cat is more than 0 and less than number of catogory
			// if so then break
			if cat <= catLen && cat >= 0 {
				break
			}
			fmt.Println("Please input the number in show range.")
		} else {
			fmt.Println("Please input the number.")
		}
		fmt.Println(catList.String())
		fmt.Print("\n > ")
	}
	if scanner.Err() != nil {
		// handle error.
		fmt.Print(scanner.Err())
	}

	word := SelWrd(files[cat-1])
	StartGame(word)
}

func StartGame(word model.HangmanWord) {
	fmt.Printf("Hint: \"%s\"\n", word.Hint)
	remain := 10
	fmt.Printf("%s score %d, remaining wrong guess %d\n", word.Show(), getScore(word, remain), remain)

	fmt.Print("\n > ")

	// Guess the letter
	var wrong []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		if !word.Guess(input) {
			// if letter is wrong
			remain--
			if len(input) == 1 && unicode.IsLetter([]rune(input)[0]) {
				wrong = uniqueAppend(wrong, input)
			}
		}

		isEnd, show := word.IsComplete()
		if isEnd {
			fmt.Println(word.Word)
			fmt.Println("Congratulation! You have complete the game!")
			fmt.Printf("Final Score is %d", getScore(word, remain))
			break
		} else {
			fmt.Printf("%s score %d, remaining wrong guess %d, wrong guessed: %s\n", show, getScore(word, remain), remain, strings.Join(wrong, " "))
			if remain < 0 {
				fmt.Println("Game Over!")
				fmt.Println("You have run out of wrong guess!")
				fmt.Printf("Final word is \"%s\"\n", word.Word)
				fmt.Printf("Final Score is %d", getScore(word, remain))
				break
			}
			fmt.Print("\n > ")
		}
	}
}

func uniqueAppend(slice []string, ele string) []string {
	for _, item := range slice {
		if item == ele {
			return slice
		}
	}
	return append(slice, ele)
}

func getScore(word model.HangmanWord, remain int) int {
	miss := 10 - remain
	return len(word.Known)*5 - miss*2
}
