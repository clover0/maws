
bin:
	mkdir -p bin

bin/moq: bin
	GOBIN=$(PWD)/bin go install github.com/matryer/moq@latest && chmod +x $(PWD)/bin/moq

mock:
	PATH=$(PWD)/bin:$$PATH go generate ./...
