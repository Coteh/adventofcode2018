#!/bin/sh

set -e

go run "./04/04.go" < "./04/sample" | diff "./04/expected_sample" -
go run "./04/04.go" < "./04/input" | diff "./04/expected" -
