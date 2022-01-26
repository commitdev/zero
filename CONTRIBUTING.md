# Contributing to Zero

Thanks for your interest in Zero!

Zero is a fully open source project started by [Commit](https://commit.dev) but we are happy to be receiving support from our users and community.

If you have used Zero or are just interested in helping out there are a few key ways to contribute:

## Contribute some code
Zero is made up of a number of different, modular components. Eventually the idea is to make these more composable and discoverable with a module repository for anyone to supply their own, but for now we mostly organize the modules into predefined "stacks" combining the layers of infrastructure, backend, frontend, and static site.

Each module is in its own repo and has its own focus, language, etc. so there's plenty to contribute, regardless of your language of choice.
Here's a list of the core repositories:
|_Repo_|_Language_|_Description_|
|--|--|--|
| [zero](https://github.com/commitdev/zero) | Go | This repo - the application used to do prompting, module fetching, template rendering, and executing the commands of each module |
| [zero-aws-eks-stack](https://github.com/commitdev/zero-aws-eks-stack) | Terraform | The terraform code to create all infrastructure required to host the backend and frontend  applications in AWS (primarily using EKS) |
| [zero-backend-go](https://github.com/commitdev/zero-backend-go) | Go | A deployable backend service providing an API written in Go |
| [zero-backend-node](https://github.com/commitdev/zero-backend-node) | Node.js | A deployable backend service providing an API written in Node |
| [zero-frontend-react](https://github.com/commitdev/zero-frontend-react) | Javascript / React | A deployable web frontend meant to communicate with one of the zero backend services |
| [zero-static-site-gatsby](https://github.com/commitdev/zero-static-site-gatsby) | Javascript / Gatsby | A deployable static site / marketing site for your application |
| [zero-notification-service](https://github.com/commitdev/zero-notification-service) | Go | A service to abstract away some concepts around sending notifications (via Email, SMS, Slack, etc.) |
| [terraform-aws-zero](https://github.com/commitdev/zero-notification-service) | Terraform | Terraform modules that are exposed via the [Terrform Registry](https://registry.terraform.io/modules/commitdev/zero/aws/latest) and used in Zero modules. Typically functionality that is reusable and standalone |

There is a [GitHub Projects board](https://github.com/orgs/commitdev/projects/6/views/2) to aggregate the issues across all these repositories. This is a great way to get a sense of the work that is available and in the backlog across all the various modules and languages.

We are trying to make sure to keep a good amount of issues in the backlog with the "[good first issue](https://github.com/orgs/commitdev/projects/6/views/2?filterQuery=label%3A%22good+first+issue%22)" label and any issues with this label could be a good place to start either because they are relatively easy or have few dependencies. We also try to include an estimate (Fibonacci where 1 is trivial, probably just a couple lines of code, and 8 or 13 would be a huge undertaking that likely needs to be split up into smaller issues.)

### Pull Requests
If you're interested in taking on an issue please comment on it to let other people know that someone is working on it. Then you can fork the repo and start local development.

When committing code, please sign your commits and try to include relevant commit messages, starting with a tag indicating what type of change you are making. Some of the repositories use these tags to automatically generate changelogs. (`feat`, `fix`, `enhancement`, `docs`, etc.)
For example:
`fix: add proper encoding of billing parameters`
or
`enhancement: support new terraform kubernetes provider`

When submitting a pull request to one of the projects please try to follow any PR and style guidelines, and make sure any relevant GitHub Actions tests are passing. If one of the tests is failing and you don't think it's related to your change, please let us know.



## Contribute documentation
Any place you find the documentation to be lacking or incorrect we would be happy for someone to contribute a change. The documentation at our [public docs site](https://getzero.dev/docs/) is auto-generated using [Docusaurus](https://docusaurus.io/), and is updated automatically when any PR is merged into the main branch in one of the repositories.

## Write up a bug report
If you run into a problem when using Zero, please feel free to create a ticket in the relevant repository. Ideally it will be clear which repo the issue should be in, but if you're not sure you can create it in the main zero repo and we will move it around if necessary. Please include as much detail as possible and any reproduction steps that might be necessary.

## Request a feature
If there's something you think should be part of Zero but you don't see it yet, you can join the conversation in the [Zero community Slack](https://slack.getzero.dev) or [GitHub Discussions](https://github.com/commitdev/zero/discussions) and if we think it fits and it's not already covered in our roadmap we will add an issue for it.

## Tell a friend
If you like what we're doing, please share it with your network!
