---
title: Prerequisites
sidebar_label: Prerequisites
sidebar_position: 2
---

### `zero check`
In order to use Zero, run the `zero check` command on your system to find out which other tools / dependencies you might need to install.

<img src="/img/docs/zero-check.png" width="400" />

[AWS CLI], [Kubectl], [Terraform], [jq], [Git], [Wget]

You need to [register a new domain](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-register.html) / [host a registered domain](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingDNS.html) you will use to access your infrastructure on [Amazon Route 53](https://aws.amazon.com/route53/).

> We recommended you have two domains - one for staging and another for production. For example, mydomain.com and mydomain-staging.com. This will lead to environments that are more similar, rather than trying to use a subdomain like staging.mydomain.com for staging which may cause issues in your app later on.

