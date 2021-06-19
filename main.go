package main

import rxyz "github.com/selfup/rxyz/pkg"

func main() {
	rxyz.Generate()
	rxyz.Compress()
	rxyz.Precompress()
	rxyz.Recompress()
}
