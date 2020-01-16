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

As there some dependencies it will be easier to use this tool within the provided image, clone the repo and then run `make build-docker-local`.
The best way then to use this is to add an alias, then you can use the CLI as if it was installed as usual on your machine:
`alias commit0='docker run -it -v "$(pwd):/project" -v "${HOME}/.aws:/root/.aws" commit0:v0'`

## Usage

1) To create a project run `commit0 create [PROJECT_NAME]`
2) A folder will be created and within that update the `commit0.yml` and then run `commit0 generate -c <commit0.yml>`
3) You will see that there is now an idl folder created.
4) Within the idl folder modify the the protobuf services generated with your desired methods
5) Go up to the parent directory and re run `commit0 generate -c <commit0.yml>`
6) You will now see a `server` folder navigate to your service folder within that directory and implement the methods generated for it
7) Once you have tested your implementation and are happy with it return to the idl repo push that directory up to git
8) Return to the parent directory and check the depency file, for go it will be the go.mod file remove the lines that point it to your local directory, this will now point it to the version on git that was pushed up previously
10) Test and push up your implementation!
9) When you feel the need to add more services add them to the commit0 config and re-run `commit0 generate` and repeat steps 4 - 7.


## Development
We are looking for contributors!

Building from the source
```
make install-go
```
Compile a new `commit0` binary in the working directory
```
make build
```

Now you can either add your project directory to your path or just execute it directly
```
mkdir tmp
cd tmp
../commit0 create test-app
cd test-app
../../commit0 generate -c commit0.yml
```

Example how run a single test for development
```
go test -run TestGenerateModules "github.com/commitdev/commit0/internal/generate" -v
```

Usage:
User either launch the web UI or start the command line
- `commit0 create <project_name>` create a `commit0.yml` in the project directory
  - prompts for the infrastructure configs as well as the credentials
  - creates a `secrets.yml`
- `commit0 generate` generates all the selected projects as defined in the `commit0.yml`
- `commit0 ui` launches the locally hosted web UI (Static SPA) and the API server

### Building locally
As there are some dependencies it will be easier to use this tool within the provided image, clone the repo and then run `make build-docker-local`.

The best way then to use this is to add an alias, then you can use the CLI as if it was installed as usual on your machine:
`alias commit0='docker run -it -v "$(pwd):/project" commit0:v0'`
