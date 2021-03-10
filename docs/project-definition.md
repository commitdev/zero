### Project Definition: `zero-project.yml`
Each project is defined by this project definition file, this manifest contains your project details, and is the source of truth for the templating(`zero create`) and provision(`zero apply`) steps. 

`zero-project.yml` is created from `zero-init`, and defines which modules are part of the project and what each module's parameters are for the subsequence steps.

| Parameters               | Type         | Description                                    |
|--------------------------|--------------|------------------------------------------------|
| `name`                   | string       | name of the project                            |
| `shouldPushRepositories` | boolean      | whether to push the modules to version control |
| `modules`                | map(modules) | a map containing modules of your project       |


### Modules
| Parameters   | Type            | Description                                                             |
|--------------|-----------------|-------------------------------------------------------------------------|
| `parameters` | map(string)     | key-value map of all the parameters to run the module                   |
| `files`      | File            | Stores information such as source-module location and destination       |
| `dependsOn`  | list(string)    | a list of dependencies that should be fulfilled before this module      |
| `conditions` | list(condition) | conditions to apply while templating out the module based on parameters |

### Condition
| Parameters   | Type         | Description                                                                                                                                           |
|--------------|--------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|
| `action`     | enum(string) | type of condition, currently supports [`ignoreFile`]                                                                                                  |
| `matchField` | string       | Allows you to condition prompt based on another parameter's value                                                                                     |
| `WhenValue`  | string       | Matches for this value to satisfy the condition                                                                                                       |
| `data`       | list(string) | Supply extra data for condition to run   `ignoreFile`: provide list of paths (file or directory path) to omit from module when condition is satisfied |