
GO ?= go
GOTESTSUM ?= gotestsum

.PHONY: build-devops
## Monta o plugin
build-devops:
	@$(GO) mod tidy
	@$(GO) test -run TestLocalDevOps ./cmd/localDevOps

.PHONY: build-proxy
## Monta o plugin
build-proxy:
	@$(GO) mod tidy
	@$(GO) build -o proxyReverse ./cmd/proxyReverse
	@./proxyReverse

.PHONY: build-server
## Monta o plugin
build-server:
	@$(GO) mod tidy
	@$(GO) build -o ./cmd/server/proxyReverse ./cmd/server
	@./cmd/server/proxyReverse

.PHONY: benchmark
## run the benchmark
benchmark:
	@$(GO) mod tidy
	@$(GO) test -bench=. ./cmd/benchmark
