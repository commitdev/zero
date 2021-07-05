---
title: zero apply
sidebar_label: zero apply
sidebar_position: 5
---

The `zero apply` command takes the templated modules generated based on your input and spins up a scalable & performant infrastructure for you!

_Note that this can take 20 minutes or more depending on your choices, as it is waiting for all the provisioned infrastructure to be created_
```shell
$ zero apply

# Sample Output
Choose the environments to apply. This will create infrastructure, CI pipelines, etc.
At this point, real things will be generated that may cost money!
Only a single environment may be suitable for an initial test, but for a real system we suggest setting up both staging and production environments.
✔ Production
🎉  Bootstrapping project zero-init. Please use the zero-project.yml file to modify the project as needed.
Cloud provider: AWS
Runtime platform: Kubernetes
Infrastructure executor: Terraform

...
...


✅  Done.
Your projects and infrastructure have been successfully created.  Here are some useful links and commands to get you started:
zero-aws-eks-stack:
- Repository URL: github.com/myapp-org/infrastructure
- To see your kubernetes clusters, run: 'kubectl config get-contexts'
- To switch to a cluster, use the following commands:
- for production use: kubectl config use-context arn:aws:eks:us-west-2:123456789:cluster/myapp-infra-production-us-west-2

- To inspect the selected cluster, run 'kubectl get node,service,deployment,pods'
zero-deployable-react-frontend:
- Repository URL: github.com/myapp-org/frontend
- Deployment Pipeline URL: https://app.circleci.com/pipelines/github/myapp-org/frontend
- Production Landing Page: app.commitzero.com

zero-deployable-backend:
- Repository URL: github.com/myapp-org/backend-service
- Deployment Pipeline URL: https://app.circleci.com/pipelines/github/myapp-org/backend-service
- Production API: api.commitzero.com
```

***Your stack is now up and running, follow the links in your terminal to visit your application 🎉***
