VERSION:= 0.0.1

deps:
	go get -u github.com/gobuffalo/packr/v2/packr2

fmt:
	go fmt ./...

run:
	go run main.go

build:
	packr2
	go build -o sprout
