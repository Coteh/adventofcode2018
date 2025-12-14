#!/bin/sh

set -e

DATA_DIR="./data/2018"

go run "./12/12.go" < "${DATA_DIR}/12/sample" | diff "${DATA_DIR}/12/expected_sample" -
go run "./12/12.go" < "${DATA_DIR}/12/input" | diff "${DATA_DIR}/12/expected" -
