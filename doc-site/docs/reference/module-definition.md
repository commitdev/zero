---
title: Module Definition
sidebar_label: Module Definition
sidebar_position: 1
---

### `zero-module.yml`
This file is the definition of a Zero module. It contains a list of all the required parameters to be able to prompt a user for choices during `zero init`, information about how to template the contents of the module during `zero create`, and the information needed for the module to run (`zero apply`).
It also declares the module's dependencies to determine the order of execution in relation to other modules.

| Parameters    | type                 | Description                                      |
|---------------|----------------------|--------------------------------------------------|
| `name`        | `string`             | Name of module                                   |
| `description` | `string`             | Description of the module                        |
| `template`    | `template`           | default settings for templating out the module   |
| `author`      | `string`             | Author of the module                             |
| `icon`        | `string`             | Path to logo image                               |
| `parameters`  | `list(Parameter)`    | Parameters to prompt users                       |
| `commands`    | `Commands`           | Commands to use instead of makefile defaults     |
| `zeroVersion` | string([go-semver])  | Zero versions it's compatible with               |


### Commands
:::note
Commands are the lifecycle of `zero apply`, it will run all modules' `check` phase, then once satisfied, run in sequence the `apply` phase, then if successful run the `summary` phase.
:::

| Parameters | Type     | Default        | Description                                                              |
|------------|----------|----------------|--------------------------------------------------------------------------|
| `check`    | `string` | `make check`   | Command to check module requirements. check is satisfied if exit code is 0 eg: `sh check-token.sh`, `zero apply` will check all modules before executing |
| `apply`    | `string` | `make`         | Command to execute the project provisioning.                             |
| `summary`  | `string` | `make summary` | Command to summarize to users the module's output and next steps.        |

#### Template
Control how module templates will be parsed during the `zero create` command.

| Parameters   | Type      | Description                                                           |
|--------------|-----------|-----------------------------------------------------------------------|
| `strictMode` | `boolean` | whether strict mode is enabled                                        |
| `delimiters` | `tuple`   | A tuple of open delimiter and ending delimiter eg: `<%` and `%>`      |
| `inputDir`   | `string`  | Folder to template from the module, becomes the module root for users |
| `outputDir`  | `string`  | local directory name for the module, gets commited to version control |

### Condition (module)
Module conditions are considered during the templatint phase (`zero create`), based on parameters supplied from the project definition.
Modules can decide to have specific files or directories excluded from the user's project.
For example if the user picks `userAuth: no`, we can exclude all the auth resources via templating.

| Parameters   | Type           | Description                                                                                                                   |
|--------------|----------------|-------------------------------------------------------------------------------------------------------------------------------|
| `action`     | `enum(string)` | type of condition, currently supports [`ignoreFile`]                                                                          |
| `matchField` | `string`       | Allows us to check the contents of another parameter's value                                                                  |
| `WhenValue`  | `string`       | Matches for this value to satisfy the condition                                                                               |
| `data`       | `list(string)` | Supply extra data for the condition action. `ignoreFile`: list of paths (file or directory) to omit from the rendered project |

### Parameter
Parameter defines how the user will be prompted during the interview phase of `zero init`.
There are multiple ways of obtaining the value for each parameter.
Parameters may have `Conditions` that must be fulfilled, otherwise it skips the field entirely.

The precedence for different types of parameter prompts are as follow.
1. `execute`: If this parameter is supplied, the command will be executed and the value will be recorded
2. `type`: zero-defined special ways of obtaining values (for example `AWSProfilePicker` which will set 2 values to the map)
3. `value`: directly assigns a static value to a parameter
4. `prompt`: requires users to select an option OR input a string
Note: `default` specifies the value that will appear initially for that prompt, but the user could still choose a new string or entirely erase it

| Parameters            | Type              | Description                                                                                                               |
|-----------------------|-------------------|---------------------------------------------------------------------------------------------------------------------------|
| `field`               | `string`          | Key to store result for project definition                                                                                |
| `label`               | `string`          | Displayed name for the prompt                                                                                             |
| `options`             | `map`             | A map of `value: display name` pairs for users to pick from                                                               |
| `default`             | `string`          | Defaults to this value during prompt                                                                                      |
| `value`               | `string`          | Skips prompt entirely when set                                                                                            |
| `info`                | `string`          | Displays during prompt as extra information at the top of the screen guiding user's input                                 |
| `fieldValidation`     | `Validation`      | Validations for the prompt value                                                                                          |
| `type`                | `enum(string)`    | Built-in custom prompts: currently supports [`AWSProfilePicker`]                                                          |
| `execute`             | `string`          | Executes commands and takes stdout as prompt result                                                                       |
| `omitFromProjectFile` | `bool`            | Field is skipped from adding to project definition                                                                        |
| `conditions`          | `list(Condition)` | Conditions for prompt to run. If supplied, all conditions must pass. See below.                                           |
| `envVarName`          | `string`          | During `zero apply` parameters are available as env vars. Defaults to field name but can be overwritten with `envVarName` |

### Condition (parameter)
Parameter conditions are considered while running user prompts. Prompts are
executed in order of the yml, and will be skipped if conditions are not satisfied.
For example if a user decides to not use circleCI, a condition can be used to skip the circleCI_api_key prompt.

| Parameters   | Type           | Description                                                       |
|--------------|----------------|-------------------------------------------------------------------|
| `action`     | `enum(string)` | type of condition, currently supports [`KeyMatchCondition`]       |
| `matchField` | `string`       | The name of the parameter to check the value of                   |
| `whenValue`  | `string`       | The exact value to match                                          |
| `elseValue`  | `string`       | The value that will be set for this parameter if the condition match fails. Otherwise the value will be set to an empty string.  |
| `data`       | `list(string)` | Supply extra data for condition to run                            |

### Validation
Allows forcing the user to adhere to a certain format of value for a parameter, defined as a regular expression.

| Parameters     | type           | Description                         |
|----------------|----------------|-------------------------------------|
| `type`         | `enum(string)` | Currently supports [[regex](https://github.com/google/re2/wiki/Syntax)] |
| `value`        | `string`       | Regular expression string           |
| `errorMessage` | `string`       | Error message when validation fails |

[go-semver]: https://github.com/hashicorp/go-version/blob/master/README.md
