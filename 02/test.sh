#!/bin/sh

set -e

go run "./02/02-1.go" < "./02/sample" | diff "./02/expected_sample-1" -
go run "./02/02-1.go" < "./02/input" | diff "./02/expected-1" -

go run "./02/02-2.go" < "./02/sample2" | diff "./02/expected_sample2-2" -
go run "./02/02-2.go" < "./02/input" | diff "./02/expected-2" -
