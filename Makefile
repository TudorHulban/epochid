.PHONY: test
test: 
	@go test -failfast -count=1 ./... -json -cover -race | tparse -smallscreen

# for when golangci lint does not work
lint:  
	@fieldalignment ./...
	@gocritic check ./...
	@nilerr ./...
	@staticcheck ./...
	@unconvert -v ./...
	@ineffassign ./...
	@gocyclo -top 10 app
	@gocyclo -avg .
	@whitespace ./...