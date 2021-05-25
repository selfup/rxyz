# RxYz

It doesn't actually stand for anything..

Just an experiment!

### Example Use

---

Generate data and compress with regular gzip:

```bash
$ go run cmd/test_data_generator.go && go run cmd/test_data_compressor.go
2021/05/25 06:28:26 --- input chars length: 74
2021/05/25 06:28:26 --- 11k of random random bytes written
2021/05/25 06:28:26 --- Total payload byte size: 11000
2021/05/25 06:28:26 --- Total output size: 8661
```

---

Grab generated data (from disk), expand size but reduce complexity, then gzip:

```bash
$ go run cmd/extract_and_precompress.go && go run cmd/precompressed_compressor.go
2021/05/25 06:28:29 --- Total payload byte size: 20679
2021/05/25 06:28:29 --- Total output size: 9324
```

### Conclusion

#### 74 char range

`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789./:@#!$*-_+=`

11k to 8.6k with a randomly generated payload

---

#### 9 char range

`012345678`

20k to 9.3k with the "precompressor"
