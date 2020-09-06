.PHONY: all test examples
all: test examples

build:
	docker build -t test .
test: build
	docker run test go test
examples: build
	docker run test go run examples/prefix/main.go
	docker run test go run examples/suffix/main.go
	docker run test go run examples/allfix/main.go
