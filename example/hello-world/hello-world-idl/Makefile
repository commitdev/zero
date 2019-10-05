PROTOC_VERSION := 3.9.2
PROTOC_WEB_VERSION := 1.0.6
PROTO_SOURCES := -I /usr/local/include
PROTO_SOURCES += -I .
PROTO_SOURCES += -I ${GOPATH}/src
PROTO_SOURCES += -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

deps-linux: deps-protoc-linux deps-grpc-web-linux

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

generate: generate-grpc generate-web

generate-grpc:
	mkdir -p gen/go
	protoc ${PROTO_SOURCES} --go_out=plugins=grpc:./gen/go ./proto/health/*.proto
	protoc ${PROTO_SOURCES} --go_out=plugins=grpc:./gen/go ./proto/helloworld/*.proto
	cp -f -rv gen/go/proto/* gen/go
	rm -rf gen/go/proto
generate-web:
	mkdir -p gen/web
	protoc ${PROTO_SOURCES} --grpc-web_out=import_style=typescript,mode=grpcwebtext:gen/web ./proto/health/*.proto
	protoc ${PROTO_SOURCES} --grpc-web_out=import_style=typescript,mode=grpcwebtext:gen/web ./proto/helloworld/*.proto
	cp -f -rv gen/web/proto/* gen/web
	rm -rf gen/web/Proto gen/web/proto

