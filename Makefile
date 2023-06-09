prepare:
	@([[ ! -e ./app.env ]] && cp ./app.env.sample ./app.env) || echo "app.env exists" 

build: prepare
	@bash ./scripts/build_cli.sh

run: build
	@go run ./main.go

setup_dev:
	@! command -v air &>/dev/null && echo "air not found, installing air" && go install github.com/cosmtrek/air@latest || echo "air exists"

dev: setup_dev
	air .

format:
	@gofmt -l -s -w .

run_compose:
	@docker-compose -f ./deployments/docker-compose.yml up
