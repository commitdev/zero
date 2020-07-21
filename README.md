![zero](https://github.com/commitdev/zero/blob/master/docs/img/logo-small.png?raw=true)

## What is Zero

Zero is an open-source developer platform CLI tool which makes it quick and easy for technical founders & developers to build quality and reliable infrastructure to launch, grow and scale high-quality SaaS applications faster and more cost-effectively.


## Why is zero good for startups

As a technical founder or the first technical hire at a startup, your sole focus is to build the logic for your application, and get it into customers’ hands as quickly and reliably as possible. Yet you immediately face multiple hurdles before even writing the first line of code. You’re forced to make many tech trade offs, leading to decision fatigue. You waste countless hours building boilerplate SaaS features not adding direct value to your customers. You spend precious time picking up unfamiliar tech, make wrong choices that result in costly refactoring or rebuilding in the future, and are unaware of tools and best practices that would speed up your product iteration.

## Why is zero reliable, scalable, performant and secure

Zero leverages Amazons’ Elastic Kubernetes Service. EKS is amazon managed Kubernetes service where you can build and deploy your applications / containers. EKS is deeply integrated with other AWS services such as:

- [Amazon Virtual Private Cloud][vpc]
- [AWS Identity and Access Management][iam]
- [Amazon Cloud Watch][acw].
- [Auto Scaling Groups][asg].
- etc.
<!-- TODO: link to list of servies that zero porvieds out of the box -->

<!-- TODO: need some help on explaning why it's performant and secure  -->
<!-- Zero levrages  -->
<!-- Zero is archiected from the ground-up to be reliable  -->
___

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

Zero requires some dependencies to function, run the `zero check` command on your system to find out which other tools/dependencies you might need to install.

![zero-check](./docs/img/zero-check.png)

[AWS CLI], [Kubectl], [Terraform], [jq], [Git]

**Notes:**

1. **For Zero to communicate with your AWS account make sure you [authenticate AWS CLI with your account credentials](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html#cli-configure-files-methods)**

    - **You can also configure your aws cli during the zero porject initilization**


2. **You need to [register a new domain](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-register.html) / [host a registered domain](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingDNS.html) you would like to use to host  your infrastructure on [Amazon Route 53](https://aws.amazon.com/route53/)**

      - **We recommended you have two domain names one for staging another from production**


___

## Using zero to spin up your own stack
Using Zero to spin up your infrastructure is easy and straightforward; using a few commands, you can configure and deploy your very own scalable high-performant infrastructure that is production-ready.

### zero init

```
# To create and customize a new project you run
$ zero init

## Sample project initilization
✔ Project Name: zero-test
🎉  Initializing project
✔ EKS + Go + React
✔ Should the created projects be checked into github automatically? (y/n): y
✔ What's the root of the github org to create repositories in?: github.com/zero-test-org
✔ Existing AWS Profiles 
✔ default

Github personal access token: used for creating repositories for your project
Requires the following permissions: [repo::public_repo, admin::orgread:org]
The token can be created at https://github.com/settings/tokens
✔ Github Personal Access Token with access to the above organization: <MY_GITHUB_ORG_ACCESS_TOKEN>

CircleCI api token: used for setting up CI/CD for your project
The token can be created at https://app.circleci.com/settings/user/tokens
✔ Circleci api key for CI/CD: <MY_CIRCLE_CI_ACCESS_TOKEN>

✔ Production Root Host Name (e.g. mydomain.com) - this must be the root of the chosen domain, not a subdomain.: commitzero.com
✔ Production Frontend Host Name (e.g. app.): app.c0-dtoki.commitzero.com
✔ Production Backend Host Name (e.g. api.): api.c0-dtoki.commitzero.com
✔ Staging Root Host Name (e.g. mydomain-staging.com) - this must be the root of the chosen domain, not a sub✔ Staging Root Host Name (e.g. mydomain-staging.com) - this must be the root of the chosen domain, not a subdomain.:cmtzerostage.com
✔ Staging Frontend Host Name (e.g. app.): app.c0-dtoki.cmtzerostage.com
✔ Staging Backend Host Name (e.g. api.): api.c0-dtoki.cmtzerostage.com
✔ What do you want to call the zero-aws-eks-stack project?: infrastructure
✔ What do you want to call the zero-deployable-backend project?: backend-service
✔ What do you want to call the zero-deployable-react-frontend project?: frontend

```

### zero create
```
# Template the selected modules and configuration specified in zero-project.yml and push to repository.
zero create
```

### zero apply
```
zero apply
```


## Zeros' stack
![systerm-architecture](https://raw.githubusercontent.com/commitdev/zero-aws-eks-stack/master/templates/docs/architecture-overview.svg)

___

## Contributing to Zero 

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


___
## Learn More about Zero

Zeros' documents are stored in the [Commit Zero Google Drive][drive]

- [UX Design Components][ux]
- [Zeros Vision]
- [Project Board]

License: N/A

<!-- links -->
[drive]:    https://drive.google.com/drive/u/0/folders/1_b8qqy5iN5envfWvIYPW5SNR_ektt5kJ
[ux]:       https://docs.google.com/document/d/1yQ4bZ5z0slL9PpmduItEiCXYKIor0nX-nnGT3J-JOFw
[old]:      https://github.com/commitdev/zero-old
[git]:      https://git-scm.com
[kubectl]:  https://kubernetes.io/docs/tasks/tools/install-kubectl/
[terraform]:https://www.terraform.io/downloads.html
[jq]:       https://github.com/stedolan/jq
[AWS CLI]:  https://aws.amazon.com/cli/
[acw]:      https://aws.amazon.com/cloudwatch/
[vpc]:      https://aws.amazon.com/vpc/
[iam]:      https://aws.amazon.com/iam/
[asg]:      https://aws.amazon.com/autoscaling/
[press-release]: https://docs.google.com/document/d/1YNRNgCfCHCxmIpD5ZsLYG2xCBxJLFd6CBI0DS_NFqoY/edit
[zero binary]: https://github.com/commitdev/zero/releases/tag/v0.0.1
[zeros vision]: https://docs.google.com/document/d/1YNRNgCfCHCxmIpD5ZsLYG2xCBxJLFd6CBI0DS_NFqoY/edit
[project board]: [https://app.zenhub.com/workspaces/commit-zero-5da8decc7046a60001c6db44/board?filterLogic=any&repos=203630543,247773730,257676371,258369081]