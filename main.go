package main

import rxyz "github.com/selfup/rxyz/pkg"

func main() {
	rxyz.Generate(11000, "test.txt")
	rxyz.Gzip("test.txt")
	rxyz.Precompress("test.txt", "precompressed.txt")
	rxyz.Gzip("precompressed.txt")
}
