#!/bin/sh

set -e

DATA_DIR="./data/2018"

go test 16/16*

go run "./16/16.go" < "${DATA_DIR}/16/sample" | diff "${DATA_DIR}/16/expected_sample" -
go run "./16/16.go" < "${DATA_DIR}/16/input" | diff "${DATA_DIR}/16/expected" -
