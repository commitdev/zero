# Sprout [POC]

Status: currently poc

Sprout intends to be a multi-language service generator. The intention is to create a modular monolith, which is easy to seperate at a later stage when the boundries are completely understood.

Based on specified config it will generate:
  * Proto files [Done]
  * Proto libraries [Done]
  * GraphQL files [Later]
  * GraphQL libraries [Later]
  * grpc web [Partial - Libraries generates for typescript]
  * grpc gateway [ Partial  - generates swagger & grpc gateway libraries]
  * Layout [Done for go]
  * Kubernetes manifests [In progress]

It will also live with your project, when you add a new service to the config it will generate everything needed for that new service.

## What does it generate?

The generation will create project folder, within this there will be your implementation and an IDL folder

* A parent directory that implements a skeleton and sets up your service implementation of the generated artifacts
* A child directory for the IDL's, this folder will also contain generated artifacts from the IDL under 'gen'

## The development cycle

1) To create a project run `sprout create [PROJECT_NAME]`
2) A folder will be created and within that update the `sprout.yml` and then run `sprout generate -l=[LANGUAGE OF CHOICE] eg. go`
3) You will see that there is now an idl folder created.
4) Within the idl folder modify the the protobuf services generated with your desired methods
5) Go up to the parrent directory and re run `sprout generate -l=[LANGUAGE OF CHOICE]`
6) Return back to the parent directory and implement the methods
7) Once you have tested your implementation and are happy with it return to the idl repo push that directory up to git
8) Return to the parent directory and check the depency file, for go it will be the go.mod file remove the lines that point it to your local directory, this will now point it to the version on git that was pushed up previously
10) Test and push up your implementation!
9) When you feel the need to add more services add them to the sprout config and re-run `sprout generate` and repeat steps 4 - 7.

## Usage & installation

As there alot of dependencies it will be easier to use this tool within the provided image, clone the repo and then run `make build-docker-local`. The best way then to use this is to alias `docker run -v "$(pwd):/project" --user $(id -u):$(id -g) sprout:v0` as sprout from then you can use the CLI as if it was installed as usual on your machine.

## Dependencies

In order to use this you need ensure you have these installed.
* protoc
* protoc-gen-go [Go]
* protoc-gen-web [gRPC Web]
* protoc-gen-gateway [Http]
* prooc-gen-swagger [Swagger]

## Building locally

As the templates are embeded into the binary you will need to ensure packr2 is installed.

You can run `make deps` to install this.

