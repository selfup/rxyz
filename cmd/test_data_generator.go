package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	inputCharsFile, icerr := ioutil.ReadFile("chars/input.txt")

	if icerr != nil {
		log.Fatal(icerr)
	}

	trimmedInputChars := strings.Trim(string(inputCharsFile), "\n")

	inputChars := strings.Split(trimmedInputChars, "")

	log.Println("--- input chars length:", len(inputChars))

	var chars = []string{}

	for _, inputChar := range inputChars {
		chars = append(chars, inputChar)
	}

	charsLen := len(chars)

	output := ""

	for i := 0; i < 11000; i++ {
		idx := rand.Intn(charsLen)

		char := chars[idx]

		output += char
	}

	filePath := "fixtures/test_data.txt"
	fileBytes := []byte(output)

	var fileMode fs.FileMode = 0666

	err := ioutil.WriteFile(filePath, fileBytes, fileMode)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("--- 11k of random random bytes written")
}
