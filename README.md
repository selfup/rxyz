# RxYz

It doesn't actually stand for anything..

Just an experiment!

### Example Use

---

Generate data and compress with regular gzip:

```bash
$ go run cmd/test_data_generator.go && go run cmd/test_data_compressor.go
2021/05/24 22:31:34 --- input chars length: 74
2021/05/24 22:31:34 --- 11k of random random bytes written
2021/05/24 22:31:35 --- Total payload byte size: 11000
2021/05/24 22:31:35 --- Total output size: 8658
```

---

Grab generated data (from disk), expand size but reduce complexity, then gzip:

```bash
$ go run cmd/extract_and_precompress.go && go run cmd/precompressed_compressor.go
2021/05/24 22:31:37 --- Total payload byte size: 17872
2021/05/24 22:31:37 --- Total output size: 7854
```

### Conclusion

#### 74 char range

`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789./:@#!$*-_+=`

11k to 8.6k with a randomly generated payload

---

#### 8 char range

`01234567`

17.8k to 7.8k with the "precompressor"
