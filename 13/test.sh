#!/bin/sh

set -e

DATA_DIR="./data/2018"

go run "./13/13.go" < "${DATA_DIR}/13/sample" | diff "${DATA_DIR}/13/expected_sample" -
go run "./13/13.go" < "${DATA_DIR}/13/input" | diff "${DATA_DIR}/13/expected" -
