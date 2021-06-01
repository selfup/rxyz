#!/usr/bin/env bash

set -eou pipefail

mkdir -p tmp

echo '--- building bins'

go build -o tmp/tdg cmd/test_data_generator.go
go build -o tmp/tdc cmd/test_data_compressor.go
go build -o tmp/eap cmd/extract_and_precompress.go
go build -o tmp/pc cmd/precompressed_compressor.go

echo '--- bins built'

echo '--- executing bins'

./tmp/tdg
./tmp/tdc
./tmp/eap
./tmp/pc

echo '--- removing bins'

rm -rf tmp

echo '--- done'
