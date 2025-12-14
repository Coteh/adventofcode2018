#!/bin/sh

set -e

DATA_DIR="./data/2018"

go run "./06/06.go" --max-desired-distance 32 < "${DATA_DIR}/06/sample" | diff "${DATA_DIR}/06/expected_sample" -
go run "./06/06.go" < "${DATA_DIR}/06/input" | diff "${DATA_DIR}/06/expected" -
