test:
	go test -race ./...

test\:ci:
	go test -coverprofile=coverage.txt -covermode=atomic ./...
