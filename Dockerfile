FROM golang:1.12.12-alpine3.10 as builder

ENV GO111MODULE=on
ENV GOLANG_PROTOBUF_VERSION=1.3.1
ENV GRPC_GATEWAY_VERSION=1.11.3
ENV GRPC_WEB_VERSION=1.0.6
ENV PROTOBUF_VERSION=3.6.1

RUN apk add --update --no-cache build-base curl git upx && \
  rm -rf /var/cache/apk/*

RUN apk add --update nodejs npm

RUN go get \
  github.com/golang/protobuf/protoc-gen-go@v${GOLANG_PROTOBUF_VERSION} &&\
  mv /go/bin/protoc-gen-go* /usr/local/bin/

RUN curl -sSL \
  https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v${GRPC_GATEWAY_VERSION}/protoc-gen-grpc-gateway-v${GRPC_GATEWAY_VERSION}-linux-x86_64 \
  -o /usr/local/bin/protoc-gen-grpc-gateway && \
  curl -sSL \
  https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v${GRPC_GATEWAY_VERSION}/protoc-gen-swagger-v${GRPC_GATEWAY_VERSION}-linux-x86_64 \
  -o /usr/local/bin/protoc-gen-swagger

RUN curl -sSL \
  https://github.com/grpc/grpc-web/releases/download/${GRPC_WEB_VERSION}/protoc-gen-grpc-web-${GRPC_WEB_VERSION}-linux-x86_64 \
  -o /usr/local/bin/protoc-gen-grpc-web

RUN mkdir -p /tmp/protoc && \
  curl -sSL \
  https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip \
  -o /tmp/protoc/protoc.zip && \
  cd /tmp/protoc && \
  unzip protoc.zip && \
  mv /tmp/protoc/include /usr/local/include

RUN GO111MODULE=off go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

RUN curl -sSLo /tmp/terraform.zip "https://releases.hashicorp.com/terraform/0.12.12/terraform_0.12.12_linux_amd64.zip" && \
unzip -d /usr/local/bin/ /tmp/terraform.zip

RUN chmod +x /usr/local/bin/* && \
  upx --lzma /usr/local/bin/*

WORKDIR /tmp/commit0
COPY . .

RUN make build-deps && make build && \
  mv commit0 /usr/local/bin && \
  upx --lzma /usr/local/bin/commit0

FROM alpine:3.10
ENV \
  PROTOBUF_VERSION=3.6.1-r1 \
  GOPATH=/proto-libs 

RUN apk add --update bash ca-certificates git python && \
apk add --update -t deps make py-pip

RUN mkdir ${GOPATH}
COPY --from=builder /usr/local/bin /usr/local/bin
COPY --from=builder /usr/local/include /usr/local/include
COPY --from=builder /go/src/github.com/grpc-ecosystem/grpc-gateway ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway
WORKDIR /project
RUN apk add --update --no-cache make protobuf=${PROTOBUF_VERSION} && \
  rm -rf /var/cache/apk/*
ENTRYPOINT ["/usr/local/bin/commit0"]
