project_name = basic-api

help: ## This help dialog.
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run-local: ## Run the app locally
	go run cmd/main.go -b 0.0.0.0

requirements-local: ## Generate go.mod & go.sum files
	go mod tidy

clean-packages-local: ## Clean packages
	go clean -modcache

up: ## Up the containers in front
	docker compose up --build

upd: ## Up the containers in background
	docker compose up --build -d

down: ## down the containers
	docker compose down
	
in: ## connect to api container shell
	docker compose run -it basic-api sh