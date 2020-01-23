
# Configuring Commit0

This is a guide on how to configure your project manually with a single file `commit0.yml`. Simply write this file to the root of your project directory and run the commit0 CLI tool against it to generate your project files.

You can see a complete commit0.yml in our [full example](TODO: Create example). 

# Table of Contents

*  [commit0.yml](#commit0-yaml)
	*  [name*](#name)
	*  [context](#context)
	*  [modules*](#modules)
		* [source*](#module-source)
		* [params*](#module-params)
*  [commit0.module.yml](#commit0-module-yaml)
	*  [name*](#module-name)
	*  [description](#module-description)
	*  [template](#template)
		*  [extension](#template-extension)
		*  [delimiters](#template-delimiters)
		*  [output](#template-output)

# Commit0.yaml<a name="commit0-yaml"></a>
Your project config file. It describes the project 
Example:
```
name: hello-world
context: 
	cognitoPoolID: xxx
modules: 
	#- source: "../tests/modules/ci"
	- source: "github.com/zthomas/react-mui-kit"	
``` 

## name<a name="name"></a>
Name of your project. This will be used to name the github repos as well as in other parts of the generated code.

[]() | |
--- | ---
Required | True
Type | String

## context<a name="context"></a>
A key value map of global context parameters to use in the templates. 

[]() | |
--- | ---
Required | False
Type | Map[String]

## modules<a name="modules"></a>
List of modules template modules to import

[]() | |
--- | ---
Required | True
Type | Map[Module]

## source<a name="module-source"></a>
We are using go-getter to parse the sources, we you can use any URL or file formats that [go-getter](https://github.com/hashicorp/go-getter#url-format) supports.

[]() | |
--- | ---
Required | True
Type | String

## module<a name="module-params"></a>
Module parameters to use during templating

[]() | |
--- | ---
Required | True
Type | String

# Commit0.module.yaml<a name="commit0-module-yaml"></a>
The module config file. You can configure how the templating engine should process the files in the current repository.
Example:
```
name: react-mui-kit
template: 
	extension: '.tmplt'
	delimiters: 
		- '<%'
		- '%>'
	output: web-app
``` 

## name<a name="module-name"></a>
Name of your module. This will be used as the default module directory as well as a display name in the prompts.

[]() | |
--- | ---
Required | True
Type | String

## description<a name="module-description"></a>
Short description of the module

[]() | |
--- | ---
Required | False
Type | String

## template<a name="template"></a>
Template configurations
[]() | |
--- | ---
Required | False
Type | Map

## extension<a name="template-extension"></a>
File extension to signify that a file is a template. If this is defined, non-template files will not be parsed and will be copied over directly. The default value is `.tmplt`

[]() | |
--- | ---
Required | False
Type | Map

## delimiters<a name="template-delimiters"></a>
An pair of delimiters that the template engine should use. The default values are: `{{`, `}}`

[]() | |
--- | ---
Required | False
Type | Map[String]

## output<a name="template-output"></a>
Template output directory that you want the template engine to write to.

[]() | |
--- | ---
Required | False
Type | String
