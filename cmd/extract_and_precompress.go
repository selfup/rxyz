package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	/*
		ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789./:@#!$*-_+=
	*/

	inputCharMap := make(map[string]string)

	inputCharMap["a"] = "\x00"
	inputCharMap["A"] = "\x01"
	inputCharMap["b"] = "\x02"
	inputCharMap["B"] = "\x03"
	inputCharMap["c"] = "\x04"
	inputCharMap["C"] = "\x05"
	inputCharMap["d"] = "\x06"
	inputCharMap["D"] = "\x07"
	inputCharMap["e"] = "\x08"
	inputCharMap["E"] = "\x09"
	inputCharMap["f"] = "\x10"
	inputCharMap["F"] = "\x11"
	inputCharMap["g"] = "\x12"
	inputCharMap["G"] = "\x13"
	inputCharMap["h"] = "\x14"
	inputCharMap["H"] = "\x15"
	inputCharMap["i"] = "\x16"
	inputCharMap["I"] = "\x17"
	inputCharMap["j"] = "\x18"
	inputCharMap["J"] = "\x19"
	inputCharMap["k"] = "\x20"
	inputCharMap["K"] = "\x21"
	inputCharMap["l"] = "\x22"
	inputCharMap["L"] = "\x23"
	inputCharMap["m"] = "\x24"
	inputCharMap["M"] = "\x25"
	inputCharMap["n"] = "\x26"
	inputCharMap["N"] = "\x27"
	inputCharMap["o"] = "\x28"
	inputCharMap["O"] = "\x29"
	inputCharMap["p"] = "\x30"
	inputCharMap["P"] = "\x31"
	inputCharMap["q"] = "\x32"
	inputCharMap["Q"] = "\x33"
	inputCharMap["r"] = "\x34"
	inputCharMap["R"] = "\x35"
	inputCharMap["s"] = "\x36"
	inputCharMap["S"] = "\x37"
	inputCharMap["t"] = "\x38"
	inputCharMap["T"] = "\x39"
	inputCharMap["u"] = "\x40"
	inputCharMap["U"] = "\x41"
	inputCharMap["v"] = "\x42"
	inputCharMap["V"] = "\x43"
	inputCharMap["w"] = "\x44"
	inputCharMap["W"] = "\x45"
	inputCharMap["x"] = "\x46"
	inputCharMap["X"] = "\x47"
	inputCharMap["y"] = "\x48"
	inputCharMap["Y"] = "\x49"
	inputCharMap["z"] = "\x50"
	inputCharMap["Z"] = "\x51"
	inputCharMap["."] = "\x52"
	inputCharMap["/"] = "\x53"
	inputCharMap[":"] = "\x54"
	inputCharMap["@"] = "\x55"
	inputCharMap["#"] = "\x56"
	inputCharMap["!"] = "\x57"
	inputCharMap["$"] = "\x58"
	inputCharMap["*"] = "\x59"
	inputCharMap["-"] = "\x60"
	inputCharMap["_"] = "\x61"
	inputCharMap["+"] = "\x62"
	inputCharMap["="] = "\x63"
	inputCharMap["1"] = "\x64"
	inputCharMap["2"] = "\x65"
	inputCharMap["3"] = "\x66"
	inputCharMap["4"] = "\x67"
	inputCharMap["5"] = "\x68"
	inputCharMap["6"] = "\x69"
	inputCharMap["7"] = "\x70"
	inputCharMap["8"] = "\x71"
	inputCharMap["9"] = "\x72"
	inputCharMap["0"] = "\x73"

	file, ferr := ioutil.ReadFile("fixtures/test_data.txt")

	if ferr != nil {
		log.Fatal(ferr)
	}

	fileChars := string(file)

	output := ""

	for _, fileChar := range strings.Split(fileChars, "") {
		output += inputCharMap[fileChar]
	}

	filePath := "fixtures/precompressed_data.txt"
	fileBytes := []byte(output)

	var fileMode fs.FileMode = 0666

	werr := ioutil.WriteFile(filePath, fileBytes, fileMode)

	if werr != nil {
		log.Fatal(werr)
	}
}
