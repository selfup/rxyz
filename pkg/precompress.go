package rxyz

import (
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
)

func Precompress(chars string, inputFilePath string, outputFilePath string) {
	inputCharMap := make(map[string]byte)

	inputChars := strings.Split(chars, "")

	// byte remap to lowest possible byte (0..n)
	for idx, inputChar := range inputChars {
		inputCharMap[inputChar] = byte(idx)
	}

	file, ferr := ioutil.ReadFile("fixtures/" + inputFilePath)

	if ferr != nil {
		log.Fatal(ferr)
	}

	fileChars := string(file)

	var output []byte

	for _, fileChar := range strings.Split(fileChars, "") {
		charByte := inputCharMap[fileChar]

		output = append(output, charByte)
	}

	filePath := "fixtures/" + outputFilePath

	var fileMode fs.FileMode = 0666

	werr := ioutil.WriteFile(filePath, output, fileMode)

	if werr != nil {
		log.Fatal(werr)
	}
}
