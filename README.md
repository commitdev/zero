# Sprout [POC]

Status: currently poc

Sprout intends to be a multi-language service generator.

Based on specified config it will generate:
  * Proto files [Done]
  * Proto libraries [Done]
  * GraphQL files [In progress]
  * GraphQL libraries [In progress]
  * Http grpc gateway [Later]
  * grpc web [Later]
  * Layout [In progress]
  * Kubernetes manifests [In progress]

It will also live with your project, when you add a new service to the config it will generate everything needed for that new service.

## Dependencies

In order to use this you need ensure you have these installed.

* prototool
* protoc-gen-go

## Building locally

As the templates are embeded into the binary you will need to ensure packr2 is installed.

You can run `make deps` to install this.

