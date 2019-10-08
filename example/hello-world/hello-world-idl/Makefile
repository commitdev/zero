PROTOC_VERSION := 3.9.2
PROTOC_WEB_VERSION := 1.0.6

PROTO_SOURCES := -I /usr/local/include
PROTO_SOURCES += -I .
PROTO_SOURCES += -I ${GOPATH}/src
PROTO_SOURCES += -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
PROTO_SOURCES += -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway

deps-linux: deps-protoc-linux deps-grpc-web-linux deps-go

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
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go

generate: generate-grpc generate-web generate-http

generate-grpc:	
	mkdir -p gen/go
	protoc ${PROTO_SOURCES} --go_out=plugins=grpc,paths=source_relative:proto ./proto/health/*.proto
	cp -f -rv proto/proto/* gen/go
	rm -rf proto/proto
generate-web:
	mkdir -p gen/web
	protoc ${PROTO_SOURCES} --grpc-web_out=import_style=typescript,mode=grpcwebtext:proto ./proto/health/*.proto
	cp -f -rv proto/proto/* gen/web
	cp -f -rv proto/Proto/* gen/web
	rm -rf proto/proto proto/Proto
generate-http:
	mkdir -p gen/go
	protoc ${PROTO_SOURCES} --grpc-gateway_out=logtostderr=true,paths=source_relative:proto --swagger_out=logtostderr=true:proto ./proto/health/*.proto
	cp -f -rv proto/proto/* gen/go
	rm -rf proto/proto

