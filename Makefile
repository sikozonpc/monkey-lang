.PHONY: run-tests repl

# Runs the unit tests
run-tests:
	go test ./... -v

# Runs a REPL instance
repl:
	go run cmd/repl/main.go

# Shows all of the files that have 0 test coverage
cov:
	bash -c "grep -v -e ' 1$' .testCoverage.txt";
