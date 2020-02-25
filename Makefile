all:
	go run cmd/main.go

.PHONY: ci-test	

ci-test: dep 
	go test ./... -cover -race -bench=.

vet:
	go list ./... | grep -v "vendor/" | xargs go vet

dep:
	go get -u ./...