.PHONY: build
build:
	@go build

.PHONY: container
container:
	@docker build -t cmattoon/dir2cm:latest .
