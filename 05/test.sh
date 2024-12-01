#!/bin/sh

set -e

go run "./05/05.go" < "./05/sample" | diff "./05/expected_sample" -
go run "./05/05.go" < "./05/input" | diff "./05/expected" -
