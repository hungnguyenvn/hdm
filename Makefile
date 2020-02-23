all:
	go run -race cmd/main.go

.PHONY: test	

test:
	go test -race