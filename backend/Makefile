.PHONY: swag
swag:
	@echo "\n$(shell tput setaf 2)🚀 Generating Swagger docs...$(shell tput sgr0)"
	@swag init -g ./cmd/main.go -o ./docs

.PHONY: test
test:
	@echo "\n$(shell tput setaf 6)🧪 Running Tests with Coverage...$(shell tput sgr0)"
	@go test -v -cover ./internal/tests/...

.PHONY: docker
docker:
	@echo "\n$(shell tput setaf 3)🐳 Building Docker Image...$(shell tput sgr0)"
	@docker build -t goboard:test .