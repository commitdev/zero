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
	mkdir -p example
	cd example && ../sprout create -p "hello-world"
	cd example/hello-world && ../../sprout generate -l go

clean-example:
	rm -rf example

install-linux: build
	mkdir -p ${HOME}/bin
	cp sprout ${HOME}/bin/sprout
	chmod +x ${HOME}/bin/sprout
