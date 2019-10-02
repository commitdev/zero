PROTOC_VERSION := 3.9.2
PROTOC_WEB_VERSION := 1.0.6

PROTO_SOURCES := -I /usr/local/include
PROTO_SOURCES += -I .
PROTO_SOURCES += -I ${GOPATH}/src
PROTO_SOURCES += -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
PROTO_SOURCES += -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway

VERSION:= 0.0.1

deps-linux: deps-go deps-protoc-linux deps-grpc-web-linux

deps-protoc-linux:
	curl -OL https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-linux-x86_64.zip
	unzip protoc-$(PROTOC_VERSION)-linux-x86_64.zip -d protoc3
	sudo mv protoc3/bin/* /usr/local/bin/
	sudo mv protoc3/include/* /usr/local/include/
	rm -rf protoc3 protoc-$(PROTOC_VERSION)-linux-x86_64.zip

deps-grpc-web-linux:
	curl -OL https://github.com/grpc/grpc-web/releases/download/$(PROTOC_WEB_VERSION)/protoc-gen-grpc-web-$(PROTOC_WEB_VERSION)-linux-x86_64
	sudo mv protoc-gen-grpc-web-$(PROTOC_WEB_VERSION)-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web
	chmod +x /usr/local/bin/protoc-gen-grpc-web

deps-go:
	go get -u github.com/gobuffalo/packr/v2/packr2
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go

fmt:
	go fmt ./...

run:
	go run main.go

build:
	packr2 build -o sprout
	packr2 clean

build-example: build clean-example
	mkdir -p example
	cd example && ../sprout create "hello-world"
	cd example/hello-world && ../../sprout generate -l go

clean-example:
	rm -rf example

install-linux: build
	mkdir -p ${HOME}/bin
	cp sprout ${HOME}/bin/sprout
	chmod +x ${HOME}/bin/sprout
	@printf "\n Please run 'source ~/.profile' to add this installation to the path."
