#!/bin/sh

set -e

go test 18/18*

go run "./18/18.go" < "./18/sample" | diff "./18/expected_sample" -
go run "./18/18.go" < "./18/input" | diff "./18/expected" -
