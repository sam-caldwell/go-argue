#!/bin/bash -e
# scripts/run-linters.sh
# (c) 2020 Sam Caldwell.  See LICENSE.txt
#
# This ensures `make lint` runs and passes before we push.

# shellcheck disable=SC2044
for f in $(find ./ -name "*.y*ml" -type f); do
    yamllint "$f" ||{
        echo "Lint $f fails"
        exit 1
     }
    echo "Lint $f passes"
done

# shellcheck disable=SC2044
for f in $(find ./ -name "*.sh" -type f); do
    shellcheck "$f" ||{
        echo "Lint $f fails"
        exit 1
     }
    echo "Lint $f passes"
done

# shellcheck disable=SC2044
for f in $(find ./ -name "*.go"); do
    golint -set_exit_status "$f" || exit 1
    echo "$f passes"
done
