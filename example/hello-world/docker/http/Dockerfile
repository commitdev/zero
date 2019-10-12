FROM golang:1.12.6@sha256:83e8267be041b3ddf6a5792c7e464528408f75c446745642db08cfe4e8d58d18 AS build
WORKDIR /cache
COPY . .
RUN go build -o hello-world-http http/main.go 

FROM alpine
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN apk update && apk add ca-certificates
COPY --from=build /cache/hello-world-http /app/
ENTRYPOINT /app/hello-world-http
