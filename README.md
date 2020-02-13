# Commit0 [POC]

[![Build Status](https://travis-ci.org/commitdev/commit0.svg)](https://travis-ci.org/commitdev/commit0)
[![Go Report Card](https://goreportcard.com/badge/github.com/commitdev/commit0?style=flat-square)](https://goreportcard.com/report/github.com/commitdev/commit0)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/commitdev/commit0)

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
2) It will prompt you to select a cloud provider and an account profile to use
3) A folder `PROJECT_NAME` will be created. You can `cd [PROJECT_NAME]` and configure the example `commit0.yml` that's generated
4) Run `commit0 generate -c <commit0.yml>` to generate all the all the project repos
5) You can go to each project repo and follow the project readme to start the service
6) `commit0 ui` launches the locally hosted web UI (Static SPA) and the API server


# Configuring Commit0

This is a guide on how to configure your project manually with a single file `commit0.yml`. Simply write this file to the root of your project directory and run the commit0 CLI tool against it to generate your project files.


# Table of Contents

*  [commit0.yml](#commit0-yaml)
	*  [name*](#name)
	*  [context](#context)
	*  [modules*](#modules)
		* [source*](#module-source)
		* [params*](#module-params)
		* [output](#module-output)
*  [commit0.module.yml](#commit0-module-yaml)
	*  [name*](#module-name)
	*  [description](#module-description)
	*  [template](#template)
		*  [extension](#template-extension)
		*  [delimiters](#template-delimiters)
		*  [output](#template-output)

# Commit0.yaml<a name="commit0-yaml"></a>
Your project config file. It describes the project 
Example:
```
name: hello-world
context: 
	cognitoPoolID: xxx
modules: 
	#- source: "../tests/modules/ci"
	- source: "github.com/zthomas/react-mui-kit"	
	- output: "web-app"
``` 

## name<a name="name"></a>
Name of your project. This will be used to name the github repos as well as in other parts of the generated code.

[]() | |
--- | ---
Required | True
Type | String

## context<a name="context"></a>
A key value map of global context parameters to use in the templates. 

[]() | |
--- | ---
Required | False
Type | Map[String]

## modules<a name="modules"></a>
List of modules template modules to import

[]() | |
--- | ---
Required | True
Type | Map[Module]

## source<a name="module-source"></a>
We are using go-getter to parse the sources, we you can use any URL or file formats that [go-getter](https://github.com/hashicorp/go-getter#url-format) supports.

[]() | |
--- | ---
Required | True
Type | String

## module<a name="module-params"></a>
Module parameters to use during templating

[]() | |
--- | ---
Required | True
Type | String

## output<a name="module-output"></a>
Template output directory that the current module should write to.

[]() | |
--- | ---
Required | False
Type | String


# Commit0.module.yaml<a name="commit0-module-yaml"></a>
The module config file. You can configure how the templating engine should process the files in the current repository.
Example:
```
name: react-mui-kit
template: 
	extension: '.tmplt'
	delimiters: 
		- '<%'
		- '%>'
	output: web-app
``` 

## name<a name="module-name"></a>
Name of your module. This will be used as the default module directory as well as a display name in the prompts.

[]() | |
--- | ---
Required | True
Type | String

## description<a name="module-description"></a>
Short description of the module

[]() | |
--- | ---
Required | False
Type | String

## template<a name="template"></a>
Template configurations
[]() | |
--- | ---
Required | False
Type | Map

## extension<a name="template-extension"></a>
File extension to signify that a file is a template. If this is defined, non-template files will not be parsed and will be copied over directly. The default value is `.tmplt`

[]() | |
--- | ---
Required | False
Type | Map

## delimiters<a name="template-delimiters"></a>
An pair of delimiters that the template engine should use. The default values are: `{{`, `}}`

[]() | |
--- | ---
Required | False
Type | Map[String]

## output<a name="template-output"></a>
The default template output directory that you want the template engine to write to. This will be overwritten by the 

[]() | |
--- | ---
Required | False
Type | String


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

To run a single test for development
```
go test -run TestGenerateModules "github.com/commitdev/commit0/internal/generate" -v
```

### Building locally
As there are some dependencies it will be easier to use this tool within the provided image, clone the repo and then run `make build-docker-local`.

The best way then to use this is to add an alias, then you can use the CLI as if it was installed as usual on your machine:
`alias commit0='docker run -it -v "$(pwd):/project" commit0:v0'`
