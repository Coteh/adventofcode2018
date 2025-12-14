#!/bin/sh

set -e

DATA_DIR="./data/2018"

go test 18/18*

go run "./18/18.go" < "${DATA_DIR}/18/sample" | diff "${DATA_DIR}/18/expected_sample" -
go run "./18/18.go" < "${DATA_DIR}/18/input" | diff "${DATA_DIR}/18/expected" -
