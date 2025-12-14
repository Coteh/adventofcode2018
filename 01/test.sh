#!/bin/sh

set -e

DATA_DIR="./data/2018"

go run "./01/01-1.go" < "${DATA_DIR}/01/sample" | diff "${DATA_DIR}/01/expected_sample-1" -
go run "./01/01-1.go" < "${DATA_DIR}/01/input" | diff "${DATA_DIR}/01/expected-1" -

go run "./01/01-2.go" < "${DATA_DIR}/01/sample" | diff "${DATA_DIR}/01/expected_sample-2" -
go run "./01/01-2.go" < "${DATA_DIR}/01/input" | diff "${DATA_DIR}/01/expected-2" -
