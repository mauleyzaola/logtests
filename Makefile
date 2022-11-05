.PHONY:mock-install
mock-install:
	go install github.com/matryer/moq@latest

.PHONY:mock-gen
mock-gen:
	rm -rf mocks
	mkdir -p mocks
	go generate ./pkg/...

.PHONY: test
test:
	go test -v -cover -count 1 ./...