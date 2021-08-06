---
title: Core Concepts
sidebar_label: Core Concepts
sidebar_position: 1
---

## Project
A project defines a set of **modules** along with the module's parameters input by the user, and these information are stored in `zero-project.yml`. Which you can generated using `zero init`.

The Project manifest(`zero-project.yml`) is the source of truth for the commands (`create` and `apply`). It determines where to fetch the modules, the execution order of modules, whether it will push your project to version control, and other project information. You can provision both staging and production environment using the same manifest to ensure the environments are reproducible and controlled.

## Module
A module is useful bundle of code and/or resources that can be templated out with the Project Manifest, then executes a provision flow, this could be templating out terraform infrastructure as code then provisioning the resources, or creating a backend application and deploying it.

A module is defined by it's **Module manifest**(`zero-module.yml`) in it's root folder, it contains all the parameters it requires, and declares it's requirements and execution.

Modules can declare it's dependencies, for example a backend that will be deployed can declare its dependency on the infrastructure repository, so that it will execute the infrastructure module's flow before the backend itself.
