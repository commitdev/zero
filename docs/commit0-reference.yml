# Configuring Commit0

This is a guide on how to configure your project manually with a single file `commit0.yml`. Simply write this file to the root of your project directory and run the commit0 CLI tool against it to generate your project files.

You can see a complete commit0.yml in our [full example](TODO: Create example).


## Table of Contents

* [organization](#ogranization)
* [name](#name)
* [description](#description)
* [maintainers](#maintainers)
* [infrastructure](#infrastructure)
  * [\<cloud provider\>](#provider)
* [frontend](#frontend)
  * [framework](#framework)
  * [ci](#ci)
* [services](#services)
  * [\<service\>](#service)
    * name
    * description
    * language
    * gitRepo
    * dockerRepo
    * ci
* [network](#network)
  * grpc
    * host
    * port
  * http
    * enabled
    * port
  * web
    * enabled
    * port

## organization<a name="organization"></a>

Name of the github organization to store your code in.
[]() | |
--- | ---
Required | True
Type | String

## name<a name="name"></a>

Name of your project. This will be used to name the github repos as well as in other parts of the generated code.
[]() | |
--- | ---
Required | True
Type | String

## description<a name="description"></a>

Description of the project. This will be used to tell others what your project does, but can easily be updated later.
[]() | |
--- | ---
Required | False
Type | String

## maintainers<a name="maintainers"></a>

List of people who are maintaining the project.
[]() | |
--- | ---
Required | False
Type | Map

## infrastructure<a name="infrastructure"></a>

This section describes how we will set up your infrastructure. First you'll pick a cloud provider in which to create the infrastructure, then depending on the provider, you'll need to pick a few other options so we know how to configure it.

### cloud provider<a name="provider"></a>

Select one of the available cloud providers to host your infrastructure. Currently we only support AWS, but more may be added in the future for different use-cases.
[]() | |
--- | ---
Required | True
Type | enum
Options | aws

### AWS<a name="aws"></a>

Key | Required | Type | Description
--- | --- | --- | ---
accountId | True | Number | ID for the amazon account you are using.
region | True | String | This is the geographical region your infrastructure will be hosted in. Depending on who is connecting to your project you may want to base it closer to the majority of your users. See [AWS Regions](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/).
eks | False | Map | Options for using Amazon's hosted Kubernetes, Elastic Kubernetes Service.
cognito | False | Map | 
s3Hosting | False | Map

#### eks

Key | Required | Type | Description
--- | --- | --- | ---
clusterName | True  | String | Name of the cluster

#### cognito

Key | Required | Type | Description
--- | --- | --- | ---
enabled | True  | Boolean | Whether or not to use Cognito


#### s3Hosting

Key | Required | Type | Description
--- | --- | --- | ---
enabled | True  | Boolean | Whether or not to enable S3 Hosting

## frontend<a name="frontend"></a>

This is where you specify which javascript frontend framework you will use (if any) with your application. This will often be a preference as most frameworks will allow you to achieve what you want. It's more about how you structure your frontend components and how you bring in data to be displayed.

### framework<a name="framework"></a>

Frontend framework to use. Currently only React is supported.
[]() | |
--- | ---
Required | True
Type | enum
Options | react


### ci<a name="ci"></a>

Key | Required | Type | Description
--- | --- | --- | ---
system | True | Enum | github, circleci, jenkins, travis
buildImage | False | String | Docker image that you want to build with
buildTag | False | String | Docker image tag, requires an image to be specified too.
buildCommand | False | String | Command to build your application. Default: `make build`
testCommand | False | String | Command to test your application. Default: `make test`

## services<a name="services"></a>

### service<a name="service"></a>

Key | Required | Type | Description
--- | --- | --- | ---
name | True | String | 
description | True | String | 
language | True | Enum | go, java, node
gitRepo | True | String | Name of the repo for this service. This is different from your infrastructure repo.
dockerRepo | True | String | Where to store docker image once built.
ci | True | Enum | github, circleci, jenkins, travis

## network<a name="network"></a>

Network options.