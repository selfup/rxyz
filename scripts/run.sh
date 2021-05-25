#!/usr/bin/env bash

set -eou pipefail

echo '--- building bins'

go build -o tdg cmd/test_data_generator.go
go build -o tdc cmd/test_data_compressor.go
go build -o eap cmd/extract_and_precompress.go
go build -o pc cmd/precompressed_compressor.go

echo '--- bins built'

echo '--- executing bins'

./tdg && ./tdc && ./eap && ./pc

echo '--- removing bins'

rm tdg
rm tdc
rm eap
rm pc

echo '--- done!'
