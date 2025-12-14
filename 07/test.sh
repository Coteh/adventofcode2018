#!/bin/sh

set -e

DATA_DIR="./data/2018"

go test 07/07*
go run "./07/07.go" < "${DATA_DIR}/07/sample" --num-workers 2 --fixed-time-amount 0 | diff "${DATA_DIR}/07/expected_sample" -
go run "./07/07.go" < "${DATA_DIR}/07/input" | diff "${DATA_DIR}/07/expected" -
