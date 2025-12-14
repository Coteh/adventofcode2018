#!/bin/sh

set -e

DATA_DIR="./data/2018"

go run "./02/02-1.go" < "${DATA_DIR}/02/sample" | diff "${DATA_DIR}/02/expected_sample-1" -
go run "./02/02-1.go" < "${DATA_DIR}/02/input" | diff "${DATA_DIR}/02/expected-1" -

go run "./02/02-2.go" < "${DATA_DIR}/02/sample2" | diff "${DATA_DIR}/02/expected_sample2-2" -
go run "./02/02-2.go" < "${DATA_DIR}/02/input" | diff "${DATA_DIR}/02/expected-2" -
