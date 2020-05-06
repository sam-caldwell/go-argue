# Makefile
# (c) 2020 Sam Caldwell.  See LICENSE.txt
#
setup:
	./scripts/dev-setup.sh

lint:
	./scripts/run-linters.sh

go_test:
	./scripts/run-tests.sh

test: go_test

