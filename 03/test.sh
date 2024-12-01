#!/bin/sh

set -e

go run "./03/03.go" < "./03/sample" | diff "./03/expected_sample" -
go run "./03/03.go" < "./03/input" | diff "./03/expected" -
