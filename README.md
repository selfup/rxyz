# rxyz

No third party deps. All std lib :)

It doesn't actually stand for anything..

Just an experiment!

### Example Use in this repo

```
$ go run main.go
2021/06/19 10:27:34 --- input chars length: 74
2021/06/19 10:27:34 --- 11000 random bytes written
2021/06/19 10:27:34 --- fixtures/test.txt : 11000
2021/06/19 10:27:34 --- fixtures/test.txt.gz : 8654
2021/06/19 10:27:34 --- fixtures/precompressed.txt : 11000
2021/06/19 10:27:34 --- fixtures/precompressed.txt.gz : 8647
```

### Example use as a lib

Same as `main.go` in this repo:

```go
package main

import rxyz "github.com/selfup/rxyz/pkg"

func main() {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789./:@#!$*-_+="

	rxyz.Generate(chars, 11000, "test.txt")
	rxyz.Gzip("test.txt")
	rxyz.Precompress(chars, "test.txt", "precompressed.txt")
	rxyz.Gzip("precompressed.txt")
}
```
