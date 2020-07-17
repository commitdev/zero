![zero](https://github.com/commitdev/zero/blob/master/docs/img/logo-small.png?raw=true)

Zero is a developer cli tool which aims to make it quick and easy for developers to bootstrap a production ready infrasturcture and get to writing code.
***

## Getting Started

### Download and Install Zero
Download the latest [zero binary] based on your local system archetecture. unzip your download add copy the zero binary to the desired location then add it to your system path.

Zero curretnly supports:
| System | Support|
| --------|:-----:|
| MacOS   |  ✅   |
| Linux   |  ✅   |
| Windows |  ❌   |

Before you can use zero there are

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
[press-release]: https://docs.google.com/document/d/1YNRNgCfCHCxmIpD5ZsLYG2xCBxJLFd6CBI0DS_NFqoY/edit
[zero binary]: https://github.com/commitdev/zero/releases/tag/v0.0.1