#!/bin/sh

set -e

go run "./13/13.go" < "./13/sample" | diff "./13/expected_sample" -
go run "./13/13.go" < "./13/input" | diff "./13/expected" -
