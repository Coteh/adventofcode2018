#!/bin/sh

set -e

go test 16/16*

go run "./16/16.go" < "./16/sample" | diff "./16/expected_sample" -
go run "./16/16.go" < "./16/input" | diff "./16/expected" -
