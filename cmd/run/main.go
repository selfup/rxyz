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

	runTmp("cmd/test_generator/main.go")
	runTmp("cmd/test_compressor/main.go")
	runTmp("cmd/precompress/main.go")
	runTmp("cmd/gzip_precompressed_data/main.go")

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
