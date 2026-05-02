.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

GCFLAGS=
LDFLAGS_EXTRA=

clean:
	rm coverage.out && true

lint:
	go-critic check ./... || true

debug-test: ## Debug test with dlv
	dlv test -- -test.v

connect:
	dlv connect 127.0.0.1:36666

test:
	go test \
	-failfast \
	-race ./... -v

test-watch:
	gotestsum \
	--watch \
	-- \
	-failfast \
	./...

cover:
	go test -coverprofile=coverage.out
	go tool cover -func=coverage.out
	@rm coverage.out

.PHONY: run
