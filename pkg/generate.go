package rxyz

import (
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func Generate(charSet string, totalBytes int, inputFilePath string) {
	rand.Seed(time.Now().UnixNano())

	inputChars := strings.Split(charSet, "")

	log.Println("--- input chars length:", len(inputChars))

	var chars = []string{}

	chars = append(chars, inputChars...)

	charsLen := len(chars)

	output := ""

	for i := 0; i < totalBytes; i++ {
		idx := rand.Intn(charsLen)

		char := chars[idx]

		output += char
	}

	filePath := "fixtures/" + inputFilePath
	fileBytes := []byte(output)

	var fileMode fs.FileMode = 0666

	err := ioutil.WriteFile(filePath, fileBytes, fileMode)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("--- ", totalBytes, " random bytes written")
}
