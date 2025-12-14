#!/bin/sh

set -e

DATA_DIR="./data/2018"

go run "./10/10.go" < "${DATA_DIR}/10/sample" | diff "${DATA_DIR}/10/expected_sample" -
go run "./10/10.go" < "${DATA_DIR}/10/input" | diff "${DATA_DIR}/10/expected" -
