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

### Busybox `time` benchmark

On compiled build `main`

```
$ time ./main
2021/06/21 11:54:09 --- input chars length: 74
2021/06/21 11:54:09 --- 11000 random bytes written
2021/06/21 11:54:09 --- fixtures/test.txt : 11000
2021/06/21 11:54:09 --- fixtures/test.txt.gz : 8654
2021/06/21 11:54:09 --- fixtures/precompressed.txt : 11000
2021/06/21 11:54:09 --- fixtures/precompressed.txt.gz : 8647

real    0m0.030s
user    0m0.026s
sys     0m0.017s
```
