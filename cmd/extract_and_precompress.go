package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	inputCharMap := make(map[string]byte)

	inputChars, icerr := ioutil.ReadFile("chars/input.txt")

	if icerr != nil {
		log.Fatal(icerr)
	}

	trimmedInputChars := strings.Trim(string(inputChars), "\n")

	for idx, inputChar := range strings.Split(trimmedInputChars, "") {
		inputCharMap[inputChar] = byte(idx)
	}

	file, ferr := ioutil.ReadFile("fixtures/test_data.txt")

	if ferr != nil {
		log.Fatal(ferr)
	}

	fileChars := string(file)

	output := ""

	for _, fileChar := range strings.Split(fileChars, "") {
		output += string(inputCharMap[fileChar])
	}

	filePath := "fixtures/precompressed_data.txt"
	fileBytes := []byte(output)

	var fileMode fs.FileMode = 0666

	werr := ioutil.WriteFile(filePath, fileBytes, fileMode)

	if werr != nil {
		log.Fatal(werr)
	}
}
