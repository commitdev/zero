---
title: Module Definition
sidebar_label: Module Definition
sidebar_position: 1
---

### `zero-module.yml`
This file is the definition of a Zero module. It contains a list of all the required parameters to be able to prompt a user for choices during `zero init`, information about how to template the contents of the module during `zero create`, and the information needed for the module to run (`zero apply`).
It also declares the module's  dependencies to determine the order of execution in relation to other modules.

| Parameters    | type               | Description                                      |
|---------------|--------------------|--------------------------------------------------|
| `name`        | string             | Name of module                                   |
| `description` | string             | Description of the module                        |
| `template`    | template           | default settings for templating out the module   |
| `author`      | string             | Author of the module                             |
| `icon`        | string             | Path to logo image                               |
| `parameters`  | list(Parameter)    | Parameters to prompt users                       |
| `commands`    | Commands           | Commands to use instead of makefile defaults     |
| `zeroVersion` | string([go-semver])| Zero versions its compatible with                |


### Commands
Commands are the lifecycle of `zero apply`, it will run all module's `check phase`, then once satisfied run in sequence `apply phase` then if successful run `summary phase`.

| Parameters | Type   | Default        | Description                                                              |
|------------|--------|----------------|--------------------------------------------------------------------------|
| `check`    | string | `make check`   | Command to check module requirements. check is satisfied if exit code is 0 eg: `sh check-token.sh`, `zero apply` will check all modules before executing |
| `apply`    | string | `make`         | Command to execute the project provisioning.                             |
| `summary`  | string | `make summary` | Command to summarize to users the module's output and next steps.        |

#### Template
| Parameters   | Type    | Description                                                           |
|--------------|---------|-----------------------------------------------------------------------|
| `strictMode` | boolean | whether strict mode is enabled                                        |
| `delimiters` | tuple   | A tuple of open delimiter and ending delimiter eg: `<%` and `%>`      |
| `inputDir`   | string  | Folder to template from the module, becomes the module root for users |
| `outputDir`  | string  | local directory name for the module, gets commited to version control |

### Condition(module)
Module conditions are considered during template phase (`zero create`), based on parameters supplied from project-definition,
modules can decide to have specific files ignored from the user's module. For example if user picks `userAuth: no`, we can ignore the auth resources via templating.

| Parameters   | Type         | Description                                                                                                                                           |
|--------------|--------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|
| `action`     | enum(string) | type of condition, currently supports [`ignoreFile`]                                                                                                  |
| `matchField` | string       | Allows you to condition prompt based on another parameter's value                                                                                     |
| `WhenValue`  | string       | Matches for this value to satisfy the condition                                                                                                       |
| `data`       | list(string) | Supply extra data for condition to run   `ignoreFile`: provide list of paths (file or directory path) to omit from module when condition is satisfied |

### Parameter:
Parameter defines the prompt during zero-init.
There are multiple ways of obtaining the value for each parameter.
Parameters may have `Conditions` and must be fulfilled when supplied, otherwise it skips the field entirely.

The precedence for different types of parameter prompts are as follow.
1. Execute
2. type: specific ways of obtaining values (in AWS credential case it will set 2 values to the map)
3. value: directly assigns a value to a parameter
4. prompt: requires users to select an option OR input a string
Note: Default is supplied as the starting point of the user's manual input (Not when value passed in is empty)

| Parameters            | Type            | Description                                                                                                               |
|-----------------------|-----------------|---------------------------------------------------------------------------------------------------------------------------|
| `field`               | string          | key to store result for project definition                                                                                |
| `label`               | string          | displayed name for the prompt                                                                                             |
| `options`             | map             | A map of `value: display name` pairs for users to pick from                                                               |
| `default`             | string          | Defaults to this value during prompt                                                                                      |
| `value`               | string          | Skips prompt entirely when set                                                                                            |
| `info`                | string          | Displays during prompt as extra information guiding user's input                                                          |
| `fieldValidation`     | Validation      | Validations for the prompt value                                                                                          |
| `type`                | enum(string)    | Built in custom prompts: currently supports [`AWSProfilePicker`]                                                          |
| `execute`             | string          | executes commands and takes stdout as prompt result                                                                       |
| `omitFromProjectFile` | bool            | Field is skipped from adding to project definition                                                                        |
| `conditions`          | list(Condition) | Conditions for prompt to run, if supplied all conditions must pass                                                        |
| `envVarName`          | string          | During `zero apply` parameters are available as env-vars, defaults to field name but can be overwritten with `envVarName` |

### Condition(paramters)
Parameters conditions are considered while running user prompts, prompts are
executed in order of the yml, and will be skipped if conditions are not satisfied.
For example if a user decide to not use circleCI, condition can be used to skip the circleCI_api_key prompt.

| Parameters   | Type         | Description                                                       |
|--------------|--------------|-------------------------------------------------------------------|
| `action`     | enum(string) | type of condition, currently supports [`KeyMatchCondition`]         |
| `matchField` | string       | Allows you to condition prompt based on another parameter's value |
| `WhenValue`  | string       | Matches for this value to satisfy the condition                   |
| `data`       | list(string) | Supply extra data for condition to run                            |

### Validation

| Parameters     | type         | Description                         |
|----------------|--------------|-------------------------------------|
| `type`         | enum(string) | Currently supports [`regex`]          |
| `value`        | string       | Regular expression string           |
| `errorMessage` | string       | Error message when validation fails |

[go-semver]: https://github.com/hashicorp/go-version/blob/master/README.md