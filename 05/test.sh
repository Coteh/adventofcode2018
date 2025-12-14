#!/bin/sh

set -e

DATA_DIR="./data/2018"

go run "./05/05.go" < "${DATA_DIR}/05/sample" | diff "${DATA_DIR}/05/expected_sample" -
go run "./05/05.go" < "${DATA_DIR}/05/input" | diff "${DATA_DIR}/05/expected" -
