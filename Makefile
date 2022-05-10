
bin:
	mkdir -p bin

bin/moq: bin
	GOBIN=$(pwd)/bin go install github.com/matryer/moq@latest

mock:
	PATH=$(PWD)/bin:$$PATH go generate ./...
