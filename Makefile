all:
	go run cmd/main.go

.PHONY: ci-test	

ci-test:  
	env GO111MODULE=on go build
	env GO111MODULE=on go test ./... -cover -race -bench=.

vet:
	go list ./... | grep -v "vendor/" | xargs go vet
