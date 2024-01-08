all: build

GO_MOD := $(shell go list -m)

build:
	go build -o bin/ ${GO_MOD}/cmd/...
