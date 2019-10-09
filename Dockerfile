FROM golang:1.12.4-alpine3.9 as builder

ENV GO111MODULE=on
ENV GOLANG_PROTOBUF_VERSION=1.3.1
ENV GRPC_GATEWAY_VERSION=1.11.3
ENV GRPC_WEB_VERSION=1.0.6
ENV PROTOBUF_VERSION=3.9.2

RUN apk add --update --no-cache build-base curl git upx && \
  rm -rf /var/cache/apk/*

RUN go get \
  github.com/golang/protobuf/protoc-gen-go@v${GOLANG_PROTOBUF_VERSION} &&\
  mv /go/bin/protoc-gen-go* /usr/local/bin/

RUN curl -sSL \
  https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v${GRPC_GATEWAY_VERSION}/protoc-gen-grpc-gateway-v${GRPC_GATEWAY_VERSION}-linux-x86_64 \
  -o /usr/local/bin/protoc-gen-grpc-gateway && \
  curl -sSL \
  https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v${GRPC_GATEWAY_VERSION}/protoc-gen-swagger-v${GRPC_GATEWAY_VERSION}-linux-x86_64 \
  -o /usr/local/bin/protoc-gen-swagger && \
  chmod +x /usr/local/bin/protoc-gen-grpc-gateway && \
  chmod +x /usr/local/bin/protoc-gen-swagger

RUN curl -sSL \
  https://github.com/grpc/grpc-web/releases/download/${GRPC_WEB_VERSION}/protoc-gen-grpc-web-${GRPC_WEB_VERSION}-linux-x86_64 \
  -o /usr/local/bin/protoc-gen-grpc-web && \
  chmod +x /usr/local/bin/protoc-gen-grpc-web

RUN mkdir -p /tmp/protoc && \
  curl -sSL \
  https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip \
  -o /tmp/protoc/protoc.zip && \
  cd /tmp/protoc && \
  unzip protoc.zip && \
  mv /tmp/protoc/include /usr/local/include

RUN GO111MODULE=off go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

WORKDIR tmp/commit0
COPY . .
RUN make build-deps && make build && \
  mv commit0 /usr/local/bin
RUN upx --lzma /usr/local/bin/*

FROM alpine:edge
ENV \
  PROTOTOOL_PROTOC_BIN_PATH=/usr/bin/protoc \
  PROTOTOOL_PROTOC_WKT_PATH=/usr/local/include \
  PROTOBUF_VERSION=3.9.2 \
  ALPINE_PROTOBUF_VERSION_SUFFIX=r0 \
  GOPATH=/proto-libs
RUN mkdir ${GOPATH}
COPY --from=builder /usr/local/bin /usr/local/bin
COPY --from=builder /usr/local/include /usr/local/include
COPY --from=builder /go/src/github.com/grpc-ecosystem/grpc-gateway ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway
WORKDIR /project
RUN apk add --update --no-cache make protobuf=${PROTOBUF_VERSION}-${ALPINE_PROTOBUF_VERSION_SUFFIX} && \
  rm -rf /var/cache/apk/*
ENTRYPOINT ["commit0"]
