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

	inputCharMap["a"] = "0"
	inputCharMap["A"] = "00"
	inputCharMap["b"] = "1"
	inputCharMap["B"] = "01"
	inputCharMap["c"] = "10"
	inputCharMap["C"] = "11"
	inputCharMap["d"] = "2"
	inputCharMap["D"] = "02"
	inputCharMap["e"] = "20"
	inputCharMap["E"] = "22"
	inputCharMap["f"] = "12"
	inputCharMap["F"] = "3"
	inputCharMap["g"] = "03"
	inputCharMap["G"] = "30"
	inputCharMap["h"] = "33"
	inputCharMap["H"] = "13"
	inputCharMap["i"] = "23"
	inputCharMap["I"] = "33"
	inputCharMap["j"] = "31"
	inputCharMap["J"] = "32"
	inputCharMap["k"] = "4"
	inputCharMap["K"] = "04"
	inputCharMap["l"] = "40"
	inputCharMap["L"] = "44"
	inputCharMap["m"] = "14"
	inputCharMap["M"] = "24"
	inputCharMap["n"] = "34"
	inputCharMap["N"] = "41"
	inputCharMap["o"] = "42"
	inputCharMap["O"] = "43"
	inputCharMap["p"] = "5"
	inputCharMap["P"] = "05"
	inputCharMap["q"] = "50"
	inputCharMap["Q"] = "55"
	inputCharMap["r"] = "15"
	inputCharMap["R"] = "25"
	inputCharMap["s"] = "35"
	inputCharMap["S"] = "45"
	inputCharMap["t"] = "51"
	inputCharMap["T"] = "52"
	inputCharMap["u"] = "53"
	inputCharMap["U"] = "73"
	inputCharMap["v"] = "54"
	inputCharMap["V"] = "6"
	inputCharMap["w"] = "06"
	inputCharMap["W"] = "60"
	inputCharMap["x"] = "66"
	inputCharMap["X"] = "16"
	inputCharMap["y"] = "26"
	inputCharMap["Y"] = "36"
	inputCharMap["z"] = "46"
	inputCharMap["Z"] = "61"
	inputCharMap["."] = "7"
	inputCharMap["/"] = "77"
	inputCharMap[":"] = "70"
	inputCharMap["@"] = "07"
	inputCharMap["#"] = "17"
	inputCharMap["!"] = "27"
	inputCharMap["$"] = "37"
	inputCharMap["*"] = "47"
	inputCharMap["-"] = "57"
	inputCharMap["_"] = "67"
	inputCharMap["+"] = "71"
	inputCharMap["="] = "72"

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
