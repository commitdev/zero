![zero](https://github.com/commitdev/zero/blob/master/docs/img/logo-small.png?raw=true)

## What is Zero

Zero is an open-source developer platform CLI tool which makes it quick and easy for technical founders & developers to build quality and reliable infrastructure to launch, grow and scale high-quality SaaS applications faster and more cost-effectively.


## Why is zero good for startups

As a technical founder or the first technical hire at a startup, your sole focus is to build the logic for your application, and get it into customers’ hands as quickly and reliably as possible. Yet you immediately face multiple hurdles before even writing the first line of code. You’re forced to make many tech trade offs, leading to decision fatigue. You waste countless hours building boilerplate SaaS features not adding direct value to your customers. You spend precious time picking up unfamiliar tech, make wrong choices that result in costly refactoring or rebuilding in the future, and are unaware of tools and best practices that would speed up your product iteration.

## Why is zero reliable, scalable, performant and secure

<!-- TODO: need some help on how to phrase this section  -->
<!-- Zero is archiected from the ground-up to be reliable  -->

## 
***

## Getting Started

### How to Install and Configure Zero

There are multiples ways to install zero:

- Install Zero using your systems package manager

```
# MacOS
brew tap commitdev/zero
brew install zero
```

- Install Zero using the binary binary

Download the latest [zero binary] for your systems archetecture. unzip your download add copy the zero binary to the desired location then add it to your system path.

Zero curretnly supports:
| System | Support|  Package Manager |
|---------|:-----:|:------:|
| MacOS   |  ✅   | `brew` |
| Linux   |  ✅   |   n/a  |
| Windows |  ❌   |   n/a  |

### Configure zero dependencies

Zero requires some dependencies to function, run the `zero check` command on your system to find out which dependencies you might need to install.

![zero](https://github.com/commitdev/zero/blob/master/docs/img/zero-check.png?raw=true)




<!-- ### Prerequisites
- [git]
- [kubectl]
- [terraform]
- [jq]
- [aws-cli] -->

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
[press-release]: https://docs.google.com/document/d/1YNRNgCfCHCxmIpD5ZsLYG2xCBxJLFd6CBI0DS_NFqoY/edit
[zero binary]: https://github.com/commitdev/zero/releases/tag/v0.0.1