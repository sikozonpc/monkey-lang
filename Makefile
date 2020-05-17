.PHONY: run-tests

# Runs the unit tests
run-tests:
	go test ./... -v

# Shows all of the files that have 0 test coverage
cov:
	bash -c "grep -v -e ' 1$' .testCoverage.txt";