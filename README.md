# Sprout [POC]

Status: currently poc

Sprout intends to be a multi-language service generator. The intention is to create a modular monolith, which is easy to seperate at a later stage when the boundries are completely understood.

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

The generation will create 2 folders.

* A rep for the IDL's, this folder will also contain generated artifacts from the IDL under 'gen'
* A repo that implements the interfaces of the generated artifacts

`NOTE: It only creates the folders for these repos, you will still need to create the git repos on your respected platform. Aswell as initialise each folder as a git repo and push when there have been changes. (if there is a strong desire we can look at how to make this process easier.)`

## The development cycle

1) Make folder and within that folder execute `sprout create [PROJECT_NAME]`
2) A folder will be created and within that update the `sprout.yml` and then run `sprout generate -l=[LANGUAGE OF CHOICE]`
3) Move back to the root folder and you will see that there is now an idl folder created.
4) Modify the the protobuf services generated with your desired methods
5) Either run `make generate` or return to the application folder and re run `sprout generate`
6) Push up the IDL repo
6) Implement these methods on the main application repo
7) When you feel the need to add more services add them to the sprout config and re-run `sprout generate` and repeat steps 4 - 6.

## Dependencies

In order to use this you need ensure you have these installed.
* protoc
* protoc-gen-go [Go]

## Building locally

As the templates are embeded into the binary you will need to ensure packr2 is installed.

You can run `make deps` to install this.

