package rxyz

import (
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Generate(charSet string, totalBytes int, inputFilePath string) {
	MkdirDashPErr := os.MkdirAll("fixtures", os.ModePerm)
	if MkdirDashPErr != nil {
		log.Fatal(MkdirDashPErr)
	}

	rand.Seed(time.Now().UnixNano())

	inputChars := strings.Split(charSet, "")

	log.Println("--- input chars length:", len(inputChars))

	charsLen := len(inputChars)

	output := ""

	for i := 0; i < totalBytes; i++ {
		idx := rand.Intn(charsLen)

		char := inputChars[idx]

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
