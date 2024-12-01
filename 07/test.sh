#!/bin/sh

set -e

go test 07/07*
go run "./07/07.go" < "./07/sample" --num-workers 2 --fixed-time-amount 0 | diff "./07/expected_sample" -
go run "./07/07.go" < "./07/input" | diff "./07/expected" -
