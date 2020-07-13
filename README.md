![zero](https://github.com/commitdev/zero/blob/master/docs/img/logo-small.png?raw=true)

This is a WIP version of our tool Zero which aims to make it quick and easy for developers to bootstrap a new environment and get to writing code.

For the older, simple tool that just rendered templates, see [commitdev/zero-old][old]

## Press Release
Learn more about Zero's vision here:
https://docs.google.com/document/d/1YNRNgCfCHCxmIpD5ZsLYG2xCBxJLFd6CBI0DS_NFqoY/edit

## Project Board
https://app.zenhub.com/workspaces/commit-zero-5da8decc7046a60001c6db44/board?filterLogic=any&repos=203630543,247773730,257676371,258369081

## Getting Started

Before getting started with Zero on your local machine there are a few prerequisites tools required to get up and running.

### Prerequisites
- [git]
- [kubectl]
- [terraform]
- [jq]
- [aws-cli]

### Zero and AWS

Zero leverages Amazons’ Elastic Kubernetes Service. EKS is amazon managed Kubernetes service where you can build and deploy your applications / containers. EKS is deeply integrated with other AWS services such as:

- [Amazon Virtual Private Cloud][vpc]
- [AWS Identity and Access Management][iam]
- [Amazon Cloud Watch][acw].
- [Auto Scaling Groups][asg].
- etc.

<!-- Tight Integrating with  -->


#### Building this tool

```shell
$ git clone git@github.com:commitdev/zero.git
$ cd zero && make
```
#### Running the tool locally

To install the CLI into your GOPATH and test it, run:
```
$ make install-go
$ zero --help
```



## Planning and Process

Documents should all be stored in the [Commit Zero Google Drive][drive]

- [UX Design Components][ux]


<!-- links -->
[drive]:    https://drive.google.com/drive/u/0/folders/1_b8qqy5iN5envfWvIYPW5SNR_ektt5kJ
[ux]:       https://docs.google.com/document/d/1yQ4bZ5z0slL9PpmduItEiCXYKIor0nX-nnGT3J-JOFw
[old]:      https://github.com/commitdev/zero-old
[git]:      https://git-scm.com
[kubectl]:  https://kubernetes.io/docs/tasks/tools/install-kubectl/
[terraform]:https://www.terraform.io/downloads.html
[jq]:       https://github.com/stedolan/jq
[aws-cli]:  https://aws.amazon.com/cli/
[acw]:      https://aws.amazon.com/cloudwatch/
[vpc]:      https://aws.amazon.com/vpc/
[iam]:      https://aws.amazon.com/iam/
[asg]:      https://aws.amazon.com/autoscaling/