#!/bin/sh

set -e

DATA_DIR="./data/2018"

go run "./04/04.go" < "${DATA_DIR}/04/sample" | diff "${DATA_DIR}/04/expected_sample" -
go run "./04/04.go" < "${DATA_DIR}/04/input" | diff "${DATA_DIR}/04/expected" -
