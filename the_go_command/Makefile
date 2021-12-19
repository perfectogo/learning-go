.DEFAULT_GOAL := build

fmt:
	go fmt hello_world.go
.PHONY:fmt

lint: fmt
	golint hello_world.go

.PHONY:lint

vet: fmt
	go vet hello_world.go
.PHONY:vet

build: vet
	go build hello_world.go
.PHONY:build