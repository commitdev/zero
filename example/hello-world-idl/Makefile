PROTOC_VERSION := 3.9.2
PROTO_SOURCES := -I /usr/local/include
PROTO_SOURCES += -I .
PROTO_SOURCES += -I ${GOPATH}/src
PROTO_SOURCES += -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis



deps-linux:
	curl -OL https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-linux-x86_64.zip
	unzip protoc-$(PROTOC_VERSION)-linux-x86_64.zip -d protoc3
	sudo mv protoc3/bin/* /usr/local/bin/
	sudo mv protoc3/include/* /usr/local/include/
	rm -rf protoc3 protoc-$(PROTOC_VERSION)-linux-x86_64.zip

generate:
	mkdir -p gen/go
	protoc ${PROTO_SOURCES} --go_out=plugins=grpc:./gen/go ./proto/health/*.proto
	protoc ${PROTO_SOURCES} --go_out=plugins=grpc:./gen/go ./proto/helloworld/*.proto
	cp -f -rv gen/go/proto/* gen/go
	rm -rf gen/go/proto

