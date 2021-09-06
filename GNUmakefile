TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
default: build

build:
	go build .

test:
	 go test $(TEST) -v $(TESTARGS)

testacc:
	TEST_ACC=1 \
	 go test $(TEST) -v $(TESTARGS)

.PHONY: build test testacc
