.PHONY: openapi-gen
openapi-gen:
	docker run --rm -v ./:/local openapitools/openapi-generator-cli:latest generate \
		-i https://api.promptql.pro.hasura.io/openapi.json \
		-g go \
		-t /local/template \
		-c /local/openapi.config.json \
		-o /local/api

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