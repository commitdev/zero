FROM golang:1.12.6@sha256:83e8267be041b3ddf6a5792c7e464528408f75c446745642db08cfe4e8d58d18 AS build
WORKDIR /cache
COPY go.mod .
COPY go.sum .
RUN go mod tidy && go mod download
COPY . .
RUN apt-get update && apt-get -y install sudo && apt-get -y install unzip
RUN make deps-linux && mv /go/bin/protoc-* /usr/local/bin/
RUN make install-go && \
  mv /go/bin/sprout /usr/local/bin/


FROM alpine
WORKDIR project
RUN apk add --update make
COPY --from=build /usr/local/bin /usr/local/bin
COPY --from=build /usr/local/include /usr/include
