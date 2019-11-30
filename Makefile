PROTOC_VERSION := 3.9.2
PROTOC_WEB_VERSION := 1.0.6

PROTO_SOURCES := -I /usr/local/include
PROTO_SOURCES += -I .
PROTO_SOURCES += -I ${GOPATH}/src
PROTO_SOURCES += -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway

VERSION:= 0.0.1

deps-linux: deps-go deps-protoc-linux deps-grpc-web-linux

deps-protoc-linux:
	curl -OL https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-linux-x86_64.zip
	unzip protoc-$(PROTOC_VERSION)-linux-x86_64.zip -d protoc3
	mv protoc3/bin/protoc /usr/local/bin
	mv protoc3/include/* /usr/local/include
	rm -rf protoc3 protoc-$(PROTOC_VERSION)-linux-x86_64.zip

deps-grpc-web-linux:
	curl -OL https://github.com/grpc/grpc-web/releases/download/$(PROTOC_WEB_VERSION)/protoc-gen-grpc-web-$(PROTOC_WEB_VERSION)-linux-x86_64
	mv protoc-gen-grpc-web-$(PROTOC_WEB_VERSION)-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web
	chmod +x /usr/local/bin/protoc-gen-grpc-web

deps-go:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go

build-deps:
	go install github.com/gobuffalo/packr/v2/packr2

check:
	go test ./...

fmt:
	go fmt ./...

build-example: build clean-example
	mkdir -p example
	cd example && ../commit0 create "hello-world"
	cd example/hello-world && ../../commit0 generate -l go

build-docker-local:
	docker build . -t commit0:v0

build-example-docker: clean-example
	mkdir -p example
	docker run -v "$(shell pwd)/example:/project" --user $(shell id -u):$(shell id -g) commit0:v0 create "hello-world"
	docker run -v "$(shell pwd)/example/hello-world:/project" --user $(shell id -u):$(shell id -g) commit0:v0 generate -l go

clean-example:
	rm -rf example

watch-ui:
	cd internal/ui && yarn start

# builds 
build:
	CGO_ENABLED=0 packr2 build -o commit0
	packr2 clean
	cd internal/ui && yarn && yarn build

# Installs the CLI int your GOPATH
install-go:
	CGO_ENABLED=0 packr2 build -o ${GOPATH}/bin/commit0
	packr2 clean
