#!/bin/sh

set -e

go run "./12/12.go" < "./12/sample" | diff "./12/expected_sample" -
go run "./12/12.go" < "./12/input" | diff "./12/expected" -
