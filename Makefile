all:
	go run cmd/main.go

.PHONY: ci-test	

ci-test: clean vet
	go test ./... -cover -race -bench=.

vet:
	go list ./... | grep -v "vendor/" | xargs go vet

clean:
	go clean ./...