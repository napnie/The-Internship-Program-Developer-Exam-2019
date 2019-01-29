package bin

import (
	"bufio"
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"../model"
)

//ListCSVFile List all csv files in dict folder
func ListCSVFile() []string {
	var fileList []string

	files, err := ioutil.ReadDir("dict")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fileList = append(fileList, file.Name())
	}

	return fileList
}

//SelWrd Select the word from the csv file
func SelWrd(csvPath string) model.HangmanWord {
	csvFile, _ := os.Open("dict/" + csvPath)
	r := bufio.NewReader(csvFile)
	reader := csv.NewReader(r)
	var hangwords []model.HangmanWord
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		hangwords = append(hangwords,
			model.NewHangmanWord(line[0], line[1]),
		)
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	wordQuantity := len(hangwords)
	sel := random.Intn(wordQuantity)
	return hangwords[sel]
}
