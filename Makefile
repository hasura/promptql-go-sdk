.PHONY: openapi-gen
openapi-gen:
	./scripts/openapi-gen.sh

# Install golangci-lint tool to run lint locally
# https://golangci-lint.run/usage/install
.PHONY: lint
lint:
	golangci-lint run --fix

.PHONY: format
format:
	golangci-lint fmt

.PHONY: test
test:
	go test -v -coverpkg=./... -race -timeout 3m -coverprofile=coverage.out ./...