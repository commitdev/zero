## Module Definition: `zero-module.yml`
This file is the definition of a Zero module. It contains a list of all the required parameters to be able to prompt a user for choices during `zero init`, information about how to template the contents of the module during `zero create`, and the information needed for the module to run (`zero apply`).
It also declares the module's  dependencies to determine the order of execution in relation to other modules.

| Parameters    | type            | Description                |
|---------------|-----------------|----------------------------|
| `name`        | string          | Name of module             |
| `description` | string          | Description of the module  |
| `author`      | string          | Author of the module       |
| `icon`        | string          | Path to logo image         |
| `parameters`  | list(Parameter) | Parameters to prompt users |

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
| `options`             | list(string)    | A list of values for users to pick from                                                                                   |
| `default`             | string          | Defaults to this value during prompt                                                                                      |
| `value`               | string          | Skips prompt entirely when set                                                                                            |
| `info`                | string          | Displays during prompt as extra information guiding user's input                                                          |
| `fieldValidation`     | Validation      | Validations for the prompt value                                                                                          |
| `type`                | enum(string)    | Built in custom prompts: currently supports [`AWSProfilePicker`]                                                            |
| `execute`             | string          | executes commands and takes stdout as prompt result                                                                       |
| `omitFromProjectFile` | bool            | Field is skipped from adding to project definition                                                                        |
| `conditions`          | list(Condition) | Conditions for prompt to run, if supplied all conditions must pass                                                        |
| `envVarName`          | string          | During `zero apply` parameters are available as env-vars, defaults to field name but can be overwritten with `envVarName` |

### Condition
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
