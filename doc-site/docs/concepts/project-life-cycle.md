---
title: Project Life Cycle
sidebar_label: Project Life Cycle
sidebar_position: 2
---

## zero init
The goal of the `init` step is to create the project manifest (`zero-project.yml`).

`zero init` will fetch each **module(s)** defined in the `zero-project.yml`, and prompt user through a series of questions to fill in parameters required by the module definition. At this phase module definition will be parsed and provide defaults and options to guide users to fill in their project details.

:::note
It's recommended to review the `zero-project.yml` and make adjustments as needed before running `zero create` and `zero apply`
:::

## zero create
`zero create` runs on the same folder as `zero-project.yml`, it will scaffold and template out each module `templates/` folder as the basis of your repository, then push them to your version control repository(github).

During scaffolding Zero will also determine based on project's parameters to optionally scaffold conditioned files, for example it will not scaffold the authentication examples if you opted out from this step.

## zero apply
`zero apply` setups up the final steps of your provisioning, it by default runs the `make` command and modules can specify to change this behavior. Zero expects a two step process of which it will run `check` then `make` to setup the module.

### Check
`check` is ran for all the modules in your project before attempting to do the `run` step, so if one dependent's `check` fails it will not start the actual provisioning for any of the modules, this is useful to check for missing binaries or API token permissions.

### Apply
The run step is by default `make` and can be overriden in the module definition. Run should be the one time setup procedures that allow the module to function.
For example in the infrastructure repository this would be **running terraform**, and for the backend repository this could be **setting up the webhooks for your circleCI to connect to your github**.
