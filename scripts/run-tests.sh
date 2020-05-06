#!/bin/bash -e
# scripts/run-tests.sh
# (c) 2020 Sam Caldwell.  See LICENSE.txt
#
# Run all tests

(
    cd tests
    go test -v
)