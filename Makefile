VERSION = 0.0.1
BUILD ?=$(shell git rev-parse --short HEAD)
PKG ?=github.com/commitdev/zero
BUILD_ARGS=-v -trimpath -ldflags=all="-X ${PKG}/cmd.appVersion=${VERSION} -X ${PKG}/cmd.appBuild=${BUILD}"

deps:
	go mod download

check:
	go list -f '{{.Dir}}' ./... | grep -v /tmp/ | xargs go test -v

fmt:
	go fmt ./...

build-docker-local:
	docker build . -t zero:v0

build-example-docker: clean-example
	mkdir -p example
	docker run -v "$(shell pwd)/example:/project" --user $(shell id -u):$(shell id -g) zero:v0 create "hello-world"
	docker run -v "$(shell pwd)/example/hello-world:/project" --user $(shell id -u):$(shell id -g) zero:v0 generate -l go

build:
	go build ${BUILD_ARGS}

# Installs the CLI int your GOPATH
install-go:
	go build -o ${GOPATH}/bin/zero

# CI Commands used on CircleCI
ci-docker-build:
	docker build . -t commitdev/zero:${VERSION_TAG} -t commitdev/zero:latest

ci-docker-push:
	echo "${DOCKERHUB_PASS}" | docker login -u commitdev --password-stdin
	docker push commitdev/zero:${VERSION_TAG}
