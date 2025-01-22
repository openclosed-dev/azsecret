.DEFAULT_GOAL := build

.PHONY: build clean

build:
	go build -trimpath -o bin/ ./cmd/azsecret

clean:
	@rm -rf bin

