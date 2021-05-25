# RxYz

It doesn't actually stand for anything..

Just an experiment!

### Example Use

---

Generate data and compress with regular gzip:

```bash
$ go run cmd/test_data_generator.go && go run cmd/test_data_compressor.go /*
2021/05/25 07:04:55 --- input chars length: 74
2021/05/25 07:04:55 --- 11k of random random bytes written
2021/05/25 07:04:55 --- Total payload byte size: 11000
2021/05/25 07:04:55 --- Total output size: 8662
```

---

Grab generated data (from disk), expand size but reduce complexity, then gzip:

```bash
go run cmd/extract_and_precompress.go && go run cmd/precompressed_compressor.go
2021/05/25 07:05:00 --- Total payload byte size: 11000
2021/05/25 07:05:00 --- Total output size: 8663
```
