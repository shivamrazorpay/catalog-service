
.PHONY: unit-test # Run tests
unit-test:
	go test -v ./internal/...

# Run tests and generate a coverage report
coverage:
	go test -coverprofile=coverage.out ./internal/...
	go tool cover -html=coverage.out

# Tidy up dependencies
tidy:
	go mod tidy

# Clean up generated files
clean:
	rm -f coverage.out

