# kubernetes tf module

## Introduction

This Terraform module contains configuration to provision kubernetes resources.

## Organization

```
    main.tf - Configuration entrypoint.
    ingress/ - Confioguration required to provision nginx-ingress-controller.
        main.tf
        provider.tf
        variables.tf
    monitoring/ - Configuration required to provision cluster monitoring.
        main.tf
        provider.tf
        variables.tf
        fluentd/
            main.tf
            files/
                ...
        cloudwatch/
            main.tf
            files/
                ...
```