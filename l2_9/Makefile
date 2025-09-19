.PHONY: lint vet test check

vet:
	go vet ./...

lint:
	golint ./...

golangci:
	golangci-lint run

test:
	go test -v ./...

check: vet lint test
	@echo "All checks passed!"
