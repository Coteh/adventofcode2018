#!/bin/sh

set -e

go run "./10/10.go" < "./10/sample" | diff "./10/expected_sample" -
go run "./10/10.go" < "./10/input" | diff "./10/expected" -
