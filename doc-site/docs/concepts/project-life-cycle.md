---
title: Project Life Cycle
sidebar_label: Project Life Cycle
sidebar_position: 2
---

## zero init
The goal of the `init` step is to create the project manifest (`zero-project.yml`).

`zero init` will fetch each **module** defined in `zero-project.yml` from their remote repository, and prompt the user through a series of questions to fill in parameters required by each module. In this phase, the module definition will be parsed and provide defaults, options, and extra context to guide users through filling in their project details.

:::note
It's recommended to review the `zero-project.yml` and make adjustments as needed before running `zero create` and `zero apply`.
:::

## zero create
`zero create` is run in the same folder as `zero-project.yml`. It will template out each module specified in the project manifest as the basis of your repositories, then push them to your version control repository (Github).

During the `create` step, Zero will also conditionally include or exclude certain sets of files, as defined by each module. For example, it will not scaffold the authentication examples if you opted not to use this feature.

## zero apply
`zero apply` is the provisioning step that starts to make real-world changes. By default, it runs a command defined by the module to check for any dependencies, and then runs a command to actually apply any changes.

### Check
`check` is run for all the modules in your project before attempting to do the `run` step, so if a dependent's `check` fails it will not start the actual provisioning for any of the modules. This is useful to check for missing binaries or API token permissions before starting to apply any changes.

### Apply
By default, the run step is to execute `make` in the root of the module, but that can be overridden in the module definition. Run should be the one-time setup commands that allow the module to function.
For example, in the infrastructure repository, this would be to **run terraform**, and for the backend repository, this could be to  **make API calls to CircleCI to set up your build pipeline**.
