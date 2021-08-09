---
title: Core Concepts
sidebar_label: Core Concepts
sidebar_position: 1
---

## Project
A project defines a set of **modules** to use, along with a full set of config parameters for each module which were entered by the user during the `zero init` step. These config options are stored in `zero-project.yml`.

The project manifest (`zero-project.yml`) is the source of truth for the commands `create` and `apply`. It determines from where to fetch the modules, the execution order of modules, whether it will push your project to version control, module configuration, and other project information. Both staging and production environments are provisioned using the same manifest to ensure the environments are reproducible and controlled.

## Module
A module is a useful bundle of code and/or resources that can be templated out during the `create` step, then executes a provisioning flow. This could be templating out terraform infrastructure as code then provisioning the resources, or creating a backend application, making API calls to set up a build pipeline, and then deploying it.

A module is defined by the **module manifest** (`zero-module.yml`) in its root folder. It contains all the parameters required to render the templated files (during `zero create`) and execute any provisioning steps, and declares it's requirements and the commands to execute during provisioning (`zero apply`).

Modules can declare their dependencies, for example a backend that will be deployed can declare its dependency on the infrastructure repository so that the infrastructure will already exist by the time we want to deploy to it.
