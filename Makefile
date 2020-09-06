.PHONY: all test examples uses
all: test examples uses

build:
	docker build -t test .
test: build
	docker run test go test
examples: build
	docker run test go run examples/prefix/main.go
	docker run test go run examples/suffix/main.go
	docker run test go run examples/allfix/main.go
uses:
	docker build -t uses uses
	docker run uses go run cmd/uses.go
