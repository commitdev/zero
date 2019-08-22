VERSION:= 0.0.1

deps:
	go get -u github.com/gobuffalo/packr/v2/packr2

fmt:
	go fmt ./...

run:
	go run main.go

build:
	packr2 build -o sprout
	packr2 clean

build-example:
	./sprout generate -c example/sprout.yml -l go -o example

clean-example:
	rm -rf example/example