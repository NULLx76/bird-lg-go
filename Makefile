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

.PHONY: run_frontend
run_frontend: generate
	cd ./services/frontend && go run .

.PHONY: docker_build_frontend
docker_build_frontend:
	docker build -f ./services/frontend/Dockerfile . -t frontend

.PHONY: docker_run_frontend
docker_run_frontend: docker_build_frontend
	docker run -p 8080:8080 frontend

.PHONY: run_api
run_api:
	go run ./services/api

.PHONY: docker_build_api
docker_build_api:
	docker build -f ./services/api/Dockerfile . -t api

.PHONY: docker_run_api
docker_run_api: docker_build_api
	docker run -p 8000:8000 api

.PHONY: run_proxy
run_proxy:
	sudo go run ./services/proxy
