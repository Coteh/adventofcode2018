#!/bin/sh

set -e

DATA_DIR="./data/2018"

go run "./03/03.go" < "${DATA_DIR}/03/sample" | diff "${DATA_DIR}/03/expected_sample" -
go run "./03/03.go" < "${DATA_DIR}/03/input" | diff "${DATA_DIR}/03/expected" -
