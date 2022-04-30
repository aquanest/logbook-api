# Makefile

.PHONY: build
build:
	goreleaser release --snapshot --rm-dist

.PHONY: test
test:
	go test -v -race ./cmd/main.go

.PHONY: start
start:
	./scripts/deploy.sh
