#!/bin/sh

set -e

go run "./19/19.go" < "./19/sample" | diff "./19/expected_sample" -
go run "./19/19.go" < "./19/input" | diff "./19/expected" -
