#!/bin/sh

set -e

go run "./08/08.go" < "./08/sample" | diff "./08/expected_sample" -
go run "./08/08.go" < "./08/input" | diff "./08/expected" -
