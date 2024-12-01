#!/bin/sh

set -e

go run "./06/06.go" --max-desired-distance 32 < "./06/sample" | diff "./06/expected_sample" -
go run "./06/06.go" < "./06/input" | diff "./06/expected" -
