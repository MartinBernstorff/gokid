.PHONY: verify

verify:
	go test ./...
	golangci-lint run
	go install .
