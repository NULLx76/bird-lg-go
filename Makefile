GO111MODULE=on

DOCKER_IMAGE_BASE=harbor.xirion.net/library/bird-lg
DOCKER_IMAGE_API=${DOCKER_IMAGE_BASE}/api
DOCKER_IMAGE_FRONTEND=${DOCKER_IMAGE_BASE}/frontend

.PHONY: build
build:
	go build ./...

.PHONY: test
test:
	go test ./...

.PHONY: generate
generate:
	go generate ./...

.PHONY: docker_build
docker_build: docker_build_frontend docker_build_api

.PHONY: docker_push
docker_push: docker_build
	docker push ${DOCKER_IMAGE_API}
	docker push ${DOCKER_IMAGE_FRONTEND}

.PHONY: run_frontend
run_frontend: generate
	cd ./services/frontend && go run .

.PHONY: docker_build_frontend
docker_build_frontend:
	docker build -f ./services/frontend/Dockerfile . -t ${DOCKER_IMAGE_FRONTEND}

.PHONY: docker_run_frontend
docker_run_frontend: docker_build_frontend
	docker run -p 8080:8080 ${DOCKER_IMAGE_FRONTEND}

.PHONY: run_api
run_api:
	go run ./services/api

.PHONY: docker_build_api
docker_build_api:
	docker build -f ./services/api/Dockerfile . -t ${DOCKER_IMAGE_API}

.PHONY: docker_run_api
docker_run_api: docker_build_api
	docker run -p 8000:8000 ${DOCKER_IMAGE_API}

.PHONY: run_proxy
run_proxy:
	sudo go run ./services/proxy
