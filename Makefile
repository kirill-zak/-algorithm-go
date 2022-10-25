.SILENT: help unit-tests
.DEFAULT_GOAL = help

.PHONY: clean
clean:
	@go clean

.PHONY: pre-push
pre-push: lint static-check unit-tests

.PHONY: lint
lint:
	@echo "Running GoLang CI lint..."
	@golangci-lint run

.PHONY: unit-tests
unit-tests:
	@echo "Running tests..."
	@go test ./... -cover -short -count=1

.PHONY: static-check
static-check:
	@echo "Running Static Check..."
	@staticcheck ./...

# Help
COLOUR_HEADER=\e[34;01m
COLOUR=\033[0;33m
END=\033[0m
help:
	printf "$(COLOUR)make clean$(END)		clean the generated files\n"
	printf "$(COLOUR)make pre-push$(END)		run pre push validate action for package\n"
	printf "$(COLOUR)make lint$(END)		run lint for package\n"
	printf "$(COLOUR)make unit-tests$(END)		run unit tests for package\n"
	printf "$(COLOUR)make static-check$(END)	run staticcheck for package\n"