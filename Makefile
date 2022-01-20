# Makefile

.PHONY: build
build:
	goreleaser release --snapshot --rm-dist

.PHONY: test
test:
	go test -v -race ./cmd/main.go

.PHONY: release
release:
	git bump
	goreleaser release --rm-dist
