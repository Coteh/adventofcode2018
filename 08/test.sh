#!/bin/sh

set -e

DATA_DIR="./data/2018"

go run "./08/08.go" < "${DATA_DIR}/08/sample" | diff "${DATA_DIR}/08/expected_sample" -
go run "./08/08.go" < "${DATA_DIR}/08/input" | diff "${DATA_DIR}/08/expected" -
