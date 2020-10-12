GO111MODULE=on

.PHONY: build
build:
	go build ./...

.PHONY: test
test:
	go test ./...

.PHONY: generate
generate:
	go generate ./...

.PHONY: docker_build_frontend
docker_build_frontend:
	docker build -f ./services/frontend/Dockerfile . -t frontend

.PHONY: docker_run_frontend
docker_run_frontend: docker_build_frontend
	docker run -p 8080:8080 frontend

.PHONY: run_frontend
run_frontend: generate
	go run ./services/frontend

.PHONY: run_api
run_api:
	go run ./services/api

.PHONY: run_proxy
run_proxy:
	sudo go run ./services/proxy

