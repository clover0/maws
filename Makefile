
bin:
	mkdir -p bin

bin/moq: bin
	GOBIN=$(pwd)/bin go install github.com/matryer/moq@latest && chmod +x $(pwd)/bin/mock

mock:
	PATH=$(PWD)/bin:$$PATH go generate ./...
