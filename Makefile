prepare:
	@([[ ! -e ./app.env ]] && cp ./app.env.sample ./app.env) || echo "app.env exists" 

build: prepare
	@GOOS=${GOOS:-linux} GOARCH=${GOARCH:-amd64} go build -o ./build/goselflearn_${GOOS}_${GOARCH} && \
		tar -czvf ./build/goselflearn_${GOOS}_${GOARCH}.tar.gz ./build/goselflearn_${GOOS}_${GOARCH} && \
		rm -f ./build/goselflearn_${GOOS}_${GOARCH} && \
		shasum -a 256 ./build/goselflearn_${GOOS}_${GOARCH}.tar.gz > ./build/goselflearn_${GOOS}_${GOARCH}.checksum.sha256


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
