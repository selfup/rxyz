package main

import rxyz "github.com/selfup/rxyz/pkg"

func main() {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789./:@#!$*-_+="

	rxyz.Generate(chars, 11000, "test.txt")
	rxyz.Gzip("test.txt")
	rxyz.Precompress(chars, "test.txt", "precompressed.txt")
	rxyz.Gzip("precompressed.txt")
}
