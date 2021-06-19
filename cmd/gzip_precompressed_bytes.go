package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
)

func main() {
	file, _ := ioutil.ReadFile("fixtures/precompressed_data.txt")

	var b bytes.Buffer

	w := gzip.NewWriter(&b)
	w.Write(file)
	w.Close()

	finalBytes := b.Bytes()

	ioutil.WriteFile("fixtures/precompressed_output.txt.gz", finalBytes, 0666)

	log.Println("--- Total precompressed payload byte size:", len(file))

	log.Println("--- Total precompressed output size:", len(finalBytes))
}
