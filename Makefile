prepare:
	@([[ ! -e ./app.env ]] && cp ./app.env.sample ./app.env) || echo "app.env exists" 

build: prepare
	@go build -o ./build/goselflearn

run: build
	@go run ./main.go

dev:
	air .

format:
	@gofmt -l -s -w .

run_compose:
	@docker-compose -f ./deployments/docker-compose.yml up
