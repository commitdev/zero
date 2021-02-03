![zero](https://raw.githubusercontent.com/commitdev/zero/main/docs/img/logo-small.png)

## What is Zero

Zero is an open source tool which makes it quick and easy for startup technical founders and developers to build everything they need to launch and grow high-quality SaaS applications faster and more cost-effectively.

Zero sets up everything you need so you can immediately start building your product.

## Why is Zero good for startups

As a technical founder or the first technical hire at a startup, your sole focus is to build the logic for your application and get it into customers’ hands as quickly and reliably as possible. Yet you immediately face multiple hurdles before even writing the first line of code. You’re forced to make many tech trade-offs, leading to decision fatigue. You waste countless hours building boilerplate SaaS features not adding direct value to your customers. You spend precious time picking up unfamiliar tech, make wrong choices that result in costly refactoring or rebuilding in the future, and are unaware of tools and best practices that would speed up your product iteration.

Zero was built by a team of engineers with many years of experience in building and scaling startups. We have faced all the problems you will and want to provide a way for new startups to avoid all those pitfalls. We also want to help you learn about the tech choices we made so your team can become proficient in some of the great tools we have included. The system you get starts small but allows you to scale well into the future when you need to.

Everything built by Zero is yours. After using Zero to generate your infrastructure, backend, and frontend, all the code is checked into your source control repositories and becomes the basis for your new system. We provide constant updates and new modules that you can pull in on an ongoing basis, but you can also feel free to customize as much as you like with no strings attached. If you do happen to make a change to core functionality and feel like contributing it back to the project, we'd love that too!

It's easy to get started, the only thing you'll need is an AWS account. Just enter your AWS CLI tokens or choose your existing profile during the setup process and everything is built for you automatically using infrastructure-as-code so you can see exactly what's happening and easily modify it if necessary.

[Read about the day-to-day experience of using a system set up using Zero](docs/real-world-usage.md)


## Why is Zero Reliable, Scalable, Performant, and Secure

Reliability: Your infrastructure will be set up in multiple availability zones making it highly available and fault tolerant. All production workloads will run with multiple instances by default, using AWS ELB and Nginx to load balance traffic. All infrastructure is represented with code using [HashiCorp Terraform][terraform] so your environments are reproducible, auditable, and easy to configure.

Scalability: Your services will be running in Kubernetes, with the EKS nodes running in an AWS [Auto Scaling Group][asg]. Both the application workloads and cluster size are ready to scale whenever the need arises. Your frontend assets will be stored in S3 and served from AWS' Cloudfront CDN which operates at global scale.

Security: Properly configured access-control to resources/security groups, using secret storage systems (AWS Secret Manager, Kubernetes secrets), and following best practices provides great security out of the box. Our practices are built on top of multiple security audits and penetration tests. Automatic certificate management using [Let's Encrypt][letsencrypt], database encryption, VPN support, and more means your traffic will always be encrypted. Built-in application features like user authentication help you bullet-proof your application by using existing, tested tools rather than reinventing the wheel when it comes to features like user management and auth.


## What do you get out of the box?
[Read about why we made these technology choices and where they are most applicable.](docs/technology-choices.md)

[Check out some resources for learning more about these technologies.](docs/learning-resources.md)

### Infrastructure
- Fully configured infrastructure-as-code AWS environment including:
  - VPCs per environment (staging, production) with pre-configured subnets, security groups, etc.
  - EKS Kubernetes cluster per environment, pre-configured with helpful tools like cert-manager, external-dns, nginx-ingress-controller
  - RDS database for your application (Postgres or MySQL)
  - S3 buckets and Cloudfront distributions to serve your assets
- Logging and Metrics collected automatically using either Cloudwatch or Prometheus + Grafana, Elasticsearch + Kibana
- VPN using [Wireguard][wireguard] (Optional)
- User management and Identity / Access Proxy using Ory [Kratos][kratos] and [Oathkeeper][oathkeeper] (Optional)
- Tooling to make it easy to set up secure access for your dev team
- Local/Cloud Hybrid developer environment using Telepresence (Optional)

### Backend
- Golang or Node.js example project automatically set up, Dockerized, and deployed to your new Kubernetes cluster
- CI pipeline built with [CircleCI][circleci] or GitHub Actions. Just merge a PR and a deploy will start. Your code will be built and tested, deployed to staging, then prompt you to push to production
- File upload / download support using signed Cloudfront URLs (Optional)
- Email support using [SendGrid][sendgrid] or AWS SES (Optional)
- Notification support for sending and receiving messages in your application (web, mobile, SMS, Email, etc.) (Optional) (In Progress)
- User management integration with Kratos and Oathkeeper - No need to handle login, signup, authentication yourself (Optional)

### Frontend
- React example project automatically set up, deployed and served securely to your customers
- CI pipeline built with CircleCI or GitHub Actions. Just merge a PR and a deploy will start. Your code will be built and tested, deployed to staging, then prompt you to push to production
- File upload / download support using signed Cloudfront URLs (Optional)
- User management integration with Kratos - Just style the example login / signup flow to look the way you want (Optional)
- Static site example project using Gatsby to easily make a landing page, also set up with a CI Pipeline using CircleCI (Optional)

___

## Getting Started

### How to Install and Configure Zero

There are multiple ways to install Zero:

- Install Zero using your systems package manager.

```
# MacOS
brew tap commitdev/zero
brew install zero
```

- Install Zero by downloading the binary.

Download the latest [Zero binary] for your systems architecture. Unzip your downloaded package and copy the Zero binary to the desired location and add it to your system PATH.

Zero currently supports:
| System | Support|  Package Manager |
|---------|:-----:|:------:|
| MacOS   |  ✅   | `brew` |
| Linux   |  ✅   |   `deb, rpm, apk`  |
| Windows |  ❌   |   n/a  |

### Prerequisites

In order to use Zero, run the `zero check` command on your system to find out which other tools / dependencies you might need to install.

![zero-check](./docs/img/zero-check.png)

[AWS CLI], [Kubectl], [Terraform], [jq], [Git], [Wget]

You need to [register a new domain](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-register.html) / [host a registered domain](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingDNS.html) you will use to access your infrastructure on [Amazon Route 53](https://aws.amazon.com/route53/).

> We recommended you have two domains - one for staging and another for production. For example, mydomain.com and mydomain-staging.com. This will lead to environments that are more similar, rather than trying to use a subdomain like staging.mydomain.com for staging which may cause issues in your app later on.

___

## Using zero to spin up your own stack

Using Zero to spin up your infrastructure and application is easy and straightforward. Using just a few commands, you can configure and deploy your very own scalable, high-performance, production-ready infrastructure.

A few caveats before getting started:

- For Zero to provision resources, you will need to be authenticated with AWS [(authenticate with aws-cli)](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html#cli-configure-files-methods).

- It is recommended practice to [create a GitHub org](https://docs.github.com/en/github/setting-up-and-managing-organizations-and-teams/creating-a-new-organization-from-scratch) where your code is going to live. If you choose, after creating your codebases, Zero will automatically create repositories and check in your code for you. You will need to [create a Personal Access Token](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token) to enable this.

<details>
  <summary>If using CircleCI as your build pipeline ...</summary>

  - Grant [CircleCi Organization access](https://github.com/settings/connections/applications/78a2ba87f071c28e65bb) to your repositories to allow pulling the code during the build pipeline.

  - You will need to [create a CircleCi access token](https://circleci.com/docs/2.0/managing-api-tokens/) and enter it during the setup process; you should store your generated tokens securely.

  - For your CI build to work, you need to opt into the use of third-party orbs. You can find this in your CircleCi Org Setting > Security > Allow Uncertified Orbs.
</details>

### zero init

The `zero init` command creates a new project and outputs an infrastructure configuration file with user input prompted responses into a file.  -> 📁 `YOUR_PROJECT_NAME/zero-project.yml`

```shell
# To create and customize a new project you run
$ zero init

## Sample project initialization
✔ Project Name: myapp-infra
🎉  Initializing project
✔ EKS + Go + React + Gatsby
✔ Should the created projects be checked into github automatically? (y/n): y
✔ What's the root of the github org to create repositories in?: github.com/myapp-org
✔ Existing AWS Profiles
✔ default

Github personal access token: used for creating repositories for your project
Requires the following permissions: [repo::public_repo, admin::orgread:org]
The token can be created at https://github.com/settings/tokens
✔ Github Personal Access Token with access to the above organization: <MY_GITHUB_ORG_ACCESS_TOKEN>

CircleCI api token: used for setting up CI/CD for your project
The token can be created at https://app.circleci.com/settings/user/tokens
✔ Circleci api key for CI/CD: <MY_CIRCLE_CI_ACCESS_TOKEN>
✔ us-west-2
✔ Production Root Host Name (e.g. mydomain.com) - this must be the root of the chosen domain, not a subdomain.: commitzero.com
✔ Production Frontend Host Name (e.g. app.): app.
✔ Production Backend Host Name (e.g. api.): api.
✔ Staging Root Host Name (e.g. mydomain-staging.com) - this must be the root of the chosen domain, not a subdomain.: commitzero-stage.com
✔ Staging Frontend Host Name (e.g. app.): app.
✔ Staging Backend Host Name (e.g. api.): api.
✔ What do you want to call the zero-aws-eks-stack project?: infrastructure
✔ What do you want to call the zero-deployable-backend project?: backend-service
✔ What do you want to call the zero-deployable-react-frontend project?: frontend

```

### zero create

The `zero create` command renders the infrastructure modules you've configured into your project folder and pushes your code to GitHub.

```shell
# Template the selected modules and configuration specified in zero-project.yml and push to the repository.
$ cd zero-init   # change your working dir to YOUR_PROJECT_NAME
$ zero create

## Sample Output
🕰  Fetching Modules
📝  Rendering Modules
  Finished templating : backend-service/.circleci/README.md
✅  Finished templating : backend-service/.circleci/config.yml
✅  Finished templating : backend-service/.gitignore
...
...
✅  Finished templating : infrastructure/terraform/modules/vpc/versions.tf
⬆  Done Rendering - committing repositories to version control.
✅  Repository created: github.com/myapp-org/infrastructure
✅  Repository created: github.com/myapp-org/backend-service
✅  Repository created: github.com/myapp-org/frontend
✅  Done - run zero apply to create any required infrastructure or execute any other remote commands to prepare your environments.


```

### zero apply

The `zero apply` command takes the templated modules generated based on your input and spins up a scalable & performant infrastructure for you!

_Note that this can take 20 minutes or more depending on your choices, as it is waiting for all the provisioned infrastructure to be created_
```shell
$ zero apply

# Sample Output
Choose the environments to apply. This will create infrastructure, CI pipelines, etc.
At this point, real things will be generated that may cost money!
Only a single environment may be suitable for an initial test, but for a real system we suggest setting up both staging and production environments.
✔ Production
🎉  Bootstrapping project zero-init. Please use the zero-project.yml file to modify the project as needed.
Cloud provider: AWS
Runtime platform: Kubernetes
Infrastructure executor: Terraform

...
...


✅  Done.
Your projects and infrastructure have been successfully created.  Here are some useful links and commands to get you started:
zero-aws-eks-stack:
- Repository URL: github.com/myapp-org/infrastructure
- To see your kubernetes clusters, run: 'kubectl config get-contexts'
- To switch to a cluster, use the following commands:
- for production use: kubectl config use-context arn:aws:eks:us-west-2:123456789:cluster/myapp-infra-production-us-west-2

- To inspect the selected cluster, run 'kubectl get node,service,deployment,pods'
zero-deployable-react-frontend:
- Repository URL: github.com/myapp-org/frontend
- Deployment Pipeline URL: https://app.circleci.com/pipelines/github/myapp-org/frontend
- Production Landing Page: app.commitzero.com

zero-deployable-backend:
- Repository URL: github.com/myapp-org/backend-service
- Deployment Pipeline URL: https://app.circleci.com/pipelines/github/myapp-org/backend-service
- Production API: api.commitzero.com
```

***Your stack is now up and running, follow the links in your terminal to visit your application 🎉***


## Zero Default Stack

[System Architecture Diagram](https://raw.githubusercontent.com/commitdev/zero-aws-eks-stack/main/docs/architecture-overview.svg)

The core zero modules currently available are:
| Project | URL |
|---|---|
| AWS Infrastructure | [https://github.com/commitdev/zero-aws-eks-stack](https://github.com/commitdev/zero-aws-eks-stack) |
| Backend (Go)  | [https://github.com/commitdev/zero-deployable-backend](https://github.com/commitdev/zero-deployable-backend) |
| Backend (Node.js)  | [https://github.com/commitdev/zero-deployable-node-backend](https://github.com/commitdev/zero-deployable-node-backend) |
| Frontend (React) | [https://github.com/commitdev/zero-deployable-react-frontend](https://github.com/commitdev/zero-deployable-react-frontend) |
| Static Site (Gatsby) | [https://github.com/commitdev/zero-deployable-landing-page](https://github.com/commitdev/zero-deployable-landing-page) |

___

## Contributing to Zero

Zero welcomes collaboration from the community; you can open new issues in our GitHub repo, Submit PRs' for bug fixes or browse through the tickets currently open to see what you can contribute too.

### Building this tool

```shell
$ git clone git@github.com:commitdev/zero.git
$ cd zero && make build
```

### Running the tool locally

To install the CLI into your GOPATH and test it, run:

```shell
$ make install-go
$ zero --help
```

### Releasing a new version on GitHub and Brew

We are using a tool called `goreleaser` which you can get from brew if you're on MacOS:
`brew install goreleaser`

After you have the tool, you can follow these steps:
```
export GITHUB_TOKEN=<your token with access to write to the zero repo>
git tag -s -a <version number like v0.0.1> -m "Some message about this release"
git push origin <version number>
goreleaser release
```

This will create a new release in GitHub and automatically collect all the commits since the last release into a changelog.
It will also build binaries for various OSes and attach them to the release and push them to brew.
The configuration for goreleaser is in [.goreleaser.yml](.goreleaser.yml)


___
## FAQ

Why is my deployed application not yet accessible?

- It takes about 20 - 35 mins for your deployed application to be globally available through AWS CloudFront CDN.

___

## Planning and Process

Zero's documents are stored in the [Commit Zero Google Drive][drive]

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
[Wget]: https://stackoverflow.com/questions/33886917/how-to-install-wget-in-macos
[and more]: https://github.com/commitdev/zero-aws-eks-stack/blob/master/docs/resources.md
[terraform]: https://terraform.io
[letsencrypt]: https://letsencrypt.org/
[kratos]: https://www.ory.sh/kratos/
[oathkeeper]: https://www.ory.sh/oathkeeper/
[wireguard]: https://wireguard.com/
[circleci]: https://circleci.com/
[sendgrid]: https://sendgrid.com/
[launchdarkly]: https://launchdarkly.com/
