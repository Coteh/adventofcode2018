#!/bin/sh

set -e

go run "./01/01-1.go" < "./01/sample" | diff "./01/expected_sample-1" -
go run "./01/01-1.go" < "./01/input" | diff "./01/expected-1" -

go run "./01/01-2.go" < "./01/sample" | diff "./01/expected_sample-2" -
go run "./01/01-2.go" < "./01/input" | diff "./01/expected-2" -
