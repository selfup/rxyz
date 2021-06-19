package rxyz

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
)

func Compress() {
	file, _ := ioutil.ReadFile("fixtures/test_data.txt")

	var b bytes.Buffer

	w := gzip.NewWriter(&b)
	w.Write(file)
	w.Close()

	finalBytes := b.Bytes()

	ioutil.WriteFile("fixtures/test_output.txt.gz", finalBytes, 0666)

	log.Println("--- Total payload byte size:", len(file))

	log.Println("--- Total output size:", len(finalBytes))
}
