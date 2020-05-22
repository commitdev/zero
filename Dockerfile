FROM golang:1.12.12-alpine3.10 as builder

ENV GO111MODULE=on
ENV TERRAFORM_VERSION=0.12.13

RUN apk add --update --no-cache build-base curl git upx && \
  rm -rf /var/cache/apk/*

RUN apk add --update nodejs npm

RUN curl -sSL \
  https://amazon-eks.s3-us-west-2.amazonaws.com/1.14.6/2019-08-22/bin/linux/amd64/aws-iam-authenticator \
  -o /usr/local/bin/aws-iam-authenticator

RUN GO111MODULE=off go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

RUN curl -sSLo /tmp/terraform.zip "https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip" && \
unzip -d /usr/local/bin/ /tmp/terraform.zip

RUN chmod +x /usr/local/bin/* && \
  upx --lzma /usr/local/bin/*

# Hydrate the dependency cache. This way, if the go.mod or go.sum files do not
# change we can cache the depdency layer without having to reinstall them.
WORKDIR /tmp/zero
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build && \
  mv zero /usr/local/bin && \
  upx --lzma /usr/local/bin/zero

FROM alpine:3.10
ENV \
  PROTOBUF_VERSION=3.6.1-r1 \
  GOPATH=/proto-libs

RUN apk add --update bash ca-certificates git python && \
  apk add --update -t deps make py-pip

RUN mkdir ${GOPATH}
COPY --from=builder /usr/local/bin /usr/local/bin
COPY --from=builder /go/src/github.com/grpc-ecosystem/grpc-gateway ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway
WORKDIR /project

ENTRYPOINT ["/usr/local/bin/zero"]
