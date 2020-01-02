VERSION:= 0.0.1

check:
	go test ./...

fmt:
	go fmt ./...

build-docker-local:
	docker build . -t commit0:v0

build-example-docker: clean-example
	mkdir -p example
	docker run -v "$(shell pwd)/example:/project" --user $(shell id -u):$(shell id -g) commit0:v0 create "hello-world"
	docker run -v "$(shell pwd)/example/hello-world:/project" --user $(shell id -u):$(shell id -g) commit0:v0 generate -l go

build:
	go build -o commit0

# Installs the CLI int your GOPATH
install-go:
	go build -o ${GOPATH}/bin/commit0
