.PHONY: build
build:
	@go get -d -v ./...
	@go build

.PHONY: container
container:
	@docker build -t cmattoon/dir2cm:latest .
