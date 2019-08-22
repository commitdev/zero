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

## What does it generate?

The generation will create a folder with 3 repos within it.

* A rep for the IDL's
* A repo that has the generated artifacts from the IDL
* A repo that implements the interfaces of the generated artifacts

`NOTE: It only creates the folders for these repos, you will still need to create the git repos on your respected platform. Aswell as initialise each folder as a git repo and push when there have been changes. (if there is a strong desire we can look at how to make this process easier.)`

## The development cycle

1) Setup sprout config & run generation
2) Start adding your desired methods to the protobuf files generated
3) Rerun generation
4) Push the idl and the language generated repo
5) Implement these methods on the main application repo
6) When you feel the need to add more services add them to the sprout config
7) Repeat steps 1 - 5

## Dependencies

In order to use this you need ensure you have these installed.

* prototool
* protoc-gen-go

## Building locally

As the templates are embeded into the binary you will need to ensure packr2 is installed.

You can run `make deps` to install this.

