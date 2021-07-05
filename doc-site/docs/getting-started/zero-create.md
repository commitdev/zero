---
title: zero create
sidebar_label: zero create
sidebar_position: 4
---

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