---
title: zero init
sidebar_label: zero init
sidebar_position: 3
---


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