all: build test

test:
	dist/cmd --config example/config.yaml

build:
	GO111MODULE=on go build -o dist/cmd cmd/cmd.go
