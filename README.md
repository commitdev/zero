# Commit0 [POC]

[![Build Status](https://travis-ci.org/commitdev/commit0.svg)](https://travis-ci.org/commitdev/commit0)

Status: Proof of Concept

Commit0 is an open source Push-To-Deploy tool designed to provide an amazing deployment process for developers while not compromising on dev ops best practices. Instead of using a Platform as a Service that simplifies your development but locks you in, we recreate the entire seamless workflow using open source technolgies and generate the infrastructure code for you while providing you with a simple interface.

With Commit0:
- You get the same simple Push-To-Deploy workflow that you are accustomed to with premium PaaS offerings
- Based on your configurations we'll generate all the infrastructure code that is needed to deploy and scale your application (Kubenetes manifests, Terraform, CI/CI configs etc.) and deploy to your own cloud provider.
- There's no vendor lock-in. It's all done with open source tools and generated code
- You don't need to know any dev ops to use Commit0 but if you are a dev ops engineer you can rest assured that you have a solid starting point and you can customize it as your project grows.
- We also include a set of commonly used open source microservices for tasks like authentication, user management, image resizing etc. so you can start developing the core application right away.

## Installation

As there alot of dependencies it will be easier to use this tool within the provided image, clone the repo and then run `make build-docker-local`.
The best way then to use this is to add an alias, then you can use the CLI as if it was installed as usual on your machine:
`alias commit0='docker run -v "$(pwd):/project" commit0:v0'`

## Usage

1) To create a project run `commit0 create [PROJECT_NAME]`
2) A folder will be created and within that update the `commit0.yml` and then run `commit0 generate -l=[LANGUAGE OF CHOICE] eg. go`
3) You will see that there is now an idl folder created.
4) Within the idl folder modify the the protobuf services generated with your desired methods
5) Go up to the parent directory and re run `commit0 generate -l=[LANGUAGE OF CHOICE]`
6) You will now see a `server` folder navigate to your service folder within that directory and implement the methods generated for it
7) Once you have tested your implementation and are happy with it return to the idl repo push that directory up to git
8) Return to the parent directory and check the depency file, for go it will be the go.mod file remove the lines that point it to your local directory, this will now point it to the version on git that was pushed up previously
10) Test and push up your implementation!
9) When you feel the need to add more services add them to the commit0 config and re-run `commit0 generate` and repeat steps 4 - 7.


## What does it generate?

The generation will create project folder, within this there will be your implementation and an IDL folder

* A parent directory that implements a skeleton and sets up your service implementation of the generated artifacts
* A child directory for the IDL's, this folder will also contain generated artifacts from the IDL under 'gen'

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


## Development 
We are looking for contributors!

Building from the source
```
make build-deps
make deps-go
```
this will create a commit0 executable in your working direcory. To install install it into your go path use: 
```
make install-go
```

### Architecture
The project is built with GoLang and requires Docker 
- /cmd - the CLI command entry points
- /internal/generate
- /internal/config
- /internal/templator - the templating service

### Building locally

As the templates are embeded into the binary you will need to ensure packr2 is installed.

You can run `make deps-go` to install this.


### Dependencies

In order to use this you need ensure you have these installed.
* protoc
* protoc-gen-go [Go]
* protoc-gen-web [gRPC Web]
* protoc-gen-gateway [Http]
* protoc-gen-swagger [Swagger]


