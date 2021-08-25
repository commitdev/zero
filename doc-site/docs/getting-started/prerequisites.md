---
title: Prerequisites
sidebar_label: Prerequisites
sidebar_position: 2
---


Using Zero to spin up your infrastructure and application is easy and straightforward. Using just a few commands, you can configure and deploy your very own scalable, high-performance, production-ready infrastructure.

A few caveats before getting started:

- For Zero to provision resources, you will need to be [authenticated with the AWS CLI tool ](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html#cli-configure-files-methods).

- It is recommended practice to [create a GitHub org](https://docs.github.com/en/github/setting-up-and-managing-organizations-and-teams/creating-a-new-organization-from-scratch) where your code is going to live. If you choose, after creating your codebases, Zero will automatically create repositories and check in your code for you. You will need to [create a Personal Access Token](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token) to enable this.

<details>
  <summary>If using CircleCI as your build pipeline...</summary>
  <ul>
    <li>
    Grant <a href="https://github.com/settings/connections/applications/78a2ba87f071c28e65bb">CircleCi Organization access</a> to your repositories to allow pulling the code during the build pipeline.
    </li>
    <li>
    You will need to <a href="https://circleci.com/docs/2.0/managing-api-tokens/">create a CircleCi access token</a> and enter it during the setup process; you should store your generated tokens securely.
    </li>
    <li>
    For your CI build to work, you need to opt into the use of third-party orbs. You can find this in your CircleCi Org Setting &gt; Security &gt; Allow Uncertified Orbs.
    </li>
  </ul>
</details>


### `zero check`
In order to use Zero, run the `zero check` command on your system to find out which other tools / dependencies you might need to install.

<img src="/img/docs/zero-check.png" width="400" />

[AWS CLI], [Kubectl], [Terraform], [jq], [Git], [Wget]

You need to [register a new domain](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-register.html) / [host a registered domain](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingDNS.html) you will use to access your infrastructure on [Amazon Route 53](https://aws.amazon.com/route53/).

:::tip
 We recommended you have two domains - one for staging and another for production. For example, mydomain.com and mydomain-staging.com. This will lead to environments that are more similar, rather than trying to use a subdomain like staging.mydomain.com for staging which may cause issues in your app later on.
:::

[AWS CLI]:  https://aws.amazon.com/cli/
[git]:      https://git-scm.com
[kubectl]:  https://kubernetes.io/docs/tasks/tools/install-kubectl/
[terraform]:https://www.terraform.io/downloads.html
[jq]:       https://github.com/stedolan/jq
[Wget]: https://stackoverflow.com/questions/33886917/how-to-install-wget-in-macos
