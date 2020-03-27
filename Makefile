VERSION = 0.0.1
BUILD ?=$(shell git rev-parse --short HEAD)
PKG ?=github.com/commitdev/commit0
BUILD_ARGS=-v -ldflags=all="-X ${PKG}/cmd.appVersion=${VERSION} -X ${PKG}/cmd.appBuild=${BUILD}"

check:
	go test ./...

fmt:
	go fmt ./...

build-docker-local:
	docker build . -t commit0:v0

build-example-docker: clean-example
	mkdir -p example
	docker run -v "$(shell pwd)/example:/project" --user $(shell id -u):$(shell id -g) commit0:v0 create "hello-world"
	docker run -v "$(shell pwd)/example/hello-world:/project" --user $(shell id -u):$(shell id -g) commit0:v0 generate -l go

build:
	go build ${BUILD_ARGS} 

# Installs the CLI int your GOPATH
install-go:
	go build -o ${GOPATH}/bin/commit0

# CI Commands used on CircleCI
ci-docker-build:
	docker build . -t commitdev/commit0:${VERSION_TAG} -t commitdev/commit0:latest
 
ci-docker-push:
	echo "${DOCKERHUB_PASS}" | docker login -u commitdev --password-stdin
	docker push commitdev/commit0:${VERSION_TAG}
