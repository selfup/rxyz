package rxyz

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
)

func Compress(inputFilePath string) {
	filePath := "fixtures/" + inputFilePath

	file, _ := ioutil.ReadFile(filePath)

	var b bytes.Buffer

	w := gzip.NewWriter(&b)
	w.Write(file)
	w.Close()

	finalBytes := b.Bytes()

	outputFilePath := "fixtures/" + inputFilePath + ".gz"

	ioutil.WriteFile(outputFilePath, finalBytes, 0666)

	log.Println("---", filePath, ":", len(file))

	log.Println("---", outputFilePath, ":", len(finalBytes))
}
