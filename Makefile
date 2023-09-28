.PHONY: build
build: ## Build the development docker image.
	docker compose build

.PHONY: up
up: ## Start the development docker container.
	docker compose up

.PHONY: down
down: ## Stop the development docker container.
	docker compose down

.PHONY: app
app: ## Stop the development docker container.
	docker compose exec app bash