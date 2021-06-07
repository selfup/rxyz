package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	var err error = nil

	err = os.MkdirAll("fixtures", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- running files")

	runTmp("cmd/test_data_generator.go")
	runTmp("cmd/test_data_compressor.go")
	runTmp("cmd/extract_and_precompress.go")
	runTmp("cmd/precompressed_compressor.go")

	fmt.Println("--- done")
}

func runTmp(file string) {
	var err error = nil

	cmd := exec.Command("go", "run", file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
