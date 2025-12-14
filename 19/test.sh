#!/bin/sh

set -e

DATA_DIR="./data/2018"

go run "./19/19.go" < "${DATA_DIR}/19/sample" | diff "${DATA_DIR}/19/expected_sample" -
go run "./19/19.go" < "${DATA_DIR}/19/input" | diff "${DATA_DIR}/19/expected" -
