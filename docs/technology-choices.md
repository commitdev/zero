![zero](https://raw.githubusercontent.com/commitdev/zero/main/docs/img/logo-small.png)

## Technology Choices
As we add features to Zero, we rely heavily on our years of experience with founding and growing startups, and judge tools and technologies based on the axes of:
- Quality - Is it the best tool for the job? Will it allow a project to start small and scale big?
- Simplicity - Is it easy to set up automatically and in a way that is easy to understand and maintain for the developers and operators of the project?
- Price - Ideally we look for open source tools, but there are some tools we integrate with that have a cost if we see it as providing enough value to justify that cost. In this case we will often try to also include an open source alternative.

When there are multiple technologies that we consider to be front-runners, we try to add multiple options, along with documentation and use cases to describe in which situations each tool might be the right choice for a project.

### Infrastructure
#### **Cloud Provider**

[AWS](https://aws.amazon.com)

AWS has been around much longer than the other large cloud providers, and while GCP or Azure may have some additional niceties in terms of developer experience, ML functionality, etc., most of the features and functionality are pretty common across all platforms. AWS is fairly low-cost, has many features, good API and CLI support, easy to set up, and with full support for infrastructure-as-code using Terraform. It can also be quite easy for startups to get credits to offset some of their initial costs.

GCP and Azure support are planned but not implemented.

#### **Orchestration**

[Kubernetes](https://kubernetes.io)

Initially developed by Google as a scaled-down, open source version of the Borg system that runs all their applications. It has become the de facto choice for container orchestration due to its huge community adoption, flexibility of deployment, and giant feature set. Between its native functionality, and various tools that we include by default, it's an incredibly powerful platform for easily deploying secure, stable, highly-available applications. It also has first-class support for managed clusters in all the major cloud providers.

We have used Kubernetes at startups in the past and have found it to be an incredible tool to allow you to start small, but scale up a huge amount. It offers a lot of native functionality that enables great developer flows, zero-downtime deploys, auto scaling, high availability, easy visibility, and more.

Compared to something like EC2, there are lots of benefits to deploying containers as immutable artifacts, good integration with load balancers, fast deploys, etc.

Compared to serverless, it is much easier to control, monitor, and have visibility, in addition to having much more flexibility, while giving you quite a similar experience in many ways.

#### **Database**

[RDS MySQL](https://aws.amazon.com/rds/mysql/) / [RDS Postgresql](https://aws.amazon.com/rds/postgresql/)

RDS MySQL and Postgres give you all the benefits of these RDBMS tools without the burden of managing them yourself. Either one of these database technologies are great for most startups, and allow you to scale quite large, depending on your schema design and data set.

For large or complex data you can also potentially move to [RDS Aurora](https://aws.amazon.com/rds/aurora/) which has additional functionality which allows you to scale further.

#### **Logging**

[Elasticsearch + Kibana](https://aws.amazon.com/elasticsearch-service/)

AWS Managed Elasticsearch with Kibana is a great tool for centralized logging from your system. We use a tool called [Fluentd](https://www.fluentd.org/) to collect all the logs from your system and send them to Elasticsearch. Then Kibana is used to run queries against that data, set up alerts, create dashboards, etc.

This is a great choice if you want more flexibility and more advanced features. It adds some complexity due to requiring an Elasticsearch cluster, but AWS' managed ES offering allows you to offload a lot of work that would typically be required.

_- OR -_

[Cloudwatch](https://aws.amazon.com/cloudwatch/)

Potentially a cheaper and lighter-weight option, Cloudwatch is an AWS tool that supports many (but not all) of the features of Kibana.

This is a good choice if you don't have too many requirements about logging yet.

#### **Monitoring**

[Prometheus](https://prometheus.io/) + [Grafana](https://grafana.com/)

Prometheus and Grafana run in the Kubernetes cluster. Prometheus provides metrics collection from your applications and Grafana provides graphing, dashboards, alerting, etc. on that data. It can also pull in data from many other sources.

This is a great choice if you want more flexibility and more advanced features. It adds some complexity due to requiring these tools to run in your cluster but does add a significant amount of features over Cloudwatch.

_- OR -_

[Cloudwatch](https://aws.amazon.com/cloudwatch/)

Potentially a lighter-weight option, Cloudwatch is an AWS tool that supports many (but not all) of the features of Grafana.

This is a good choice if you don't have too many requirements about monitoring yet.

#### **Access**
[VPN (Wireguard)](https://www.wireguard.com/)

It is a good policy to keep as much of your infrastructure as possible private, exposing only what is necessary to the public internet. When we set up the Kubernetes cluster and other resources, they are all on private VPC Subnets so they can reach each other but are not otherwise reachable. When dealing with Kubernetes there is already great tooling for being able to access your application (via [`kubectl exec`](https://kubernetes.io/docs/tasks/debug-application-cluster/get-shell-running-container/)) but you may find cases where you want to access other resources as if you were on the private network yourself.

To this end we are using Wireguard for VPN to allow your personal computer to act as if it were on the private network, having access to the things running inside your Kubernetes cluster, and within AWS.

Wireguard is a great secure, light-weight tool that runs well on Kubernetes, so it doesn't require us to set up any other infrastructure. You may be able to get by without it if you are comfortable with the Kubernetes tooling, but it would be very useful if you are using something like Kibana above, so you can access it directly through your web browser.

### Application

#### **User Management**
[Ory Kratos](https://github.com/ory/kratos)

Kratos is an open-source identity and user management tool. It can take the place of something like Auth0, or the user management that many startups end up writing themselves. By delegating this work to Kratos, you can save time and effort that you would otherwise have to put in to build a stable, secure user management system.

If you are building a system that will require authenticated users, it may make sense to use Kratos with Oathkeeper (below) to save a lot of effort that may otherwise be spent building user management and auth features that have been built a million times before.

#### **Authentication / Authorization**
[Ory Oathkeeper](https://github.com/ory/oathkeeper)

Oathkeeper is an open-source Identity & Access Proxy that can securely act as a gateway to your system, preventing unauthorized access. It uses Oauth & JWT to log a user in and maintain a session, and proxies traffic to your services for users with valid sessions.

This allows you to save time and effort by not having to write auth and session management code. If a specific header exists, the specified user is logged in. That's it.

#### **Dev Experience**
[Telepresence](https://www.telepresence.io/)

Telepresence is a great tool for creating a fast and effective "hybrid cloud/local" developer experience using Kubernetes. You can have your services, databases, and other dependencies running in the cloud, and then run the service you are working on locally, and have it override the one running in the cloud. In this way you can use all the same IDEs, debuggers, and local tools you usually do, but have your code running in a realistic environment much more like staging or production, complete with all the dependencies that can make local dev a pain.

It's a great tool even when starting a project to quickly and easily test things in a realistic environment but can also really shine if you are working in a system that has many dependencies (multiple state stores, microservices, etc.)

### Backend
While we offer ways to bootstrap your backend application in various languages, you can also choose to use whichever language you like and just make use of the infrastructure we set up, or use your custom backend with a bootstrapped frontend.

The backend application will have a build and deploy pipeline that will build a docker image and push it to AWS' ECR image repository. The image will be used to deploy containers in the Kubernetes cluster.

**Language**

Whichever language you choose will be automatically set up, Dockerized, and deployed into the Kubernetes cluster.

[Golang](https://golang.org/)

Go is a great language for backend application development, especially in a microservices environment. It is a typed, compiled language with powerful support for parallel processing, and can create small, self-contained binaries for any operating system. It has a huge community, and is used by some very significant open source projects such as Kubernetes, and the HashiCorp tools.

Being a compiled language may slightly raise the barrier to entry for newer developers. Lack of generic types and traditional object-oriented features, and preference for code generation may deter more experienced but unfamiliar developers.

_- OR -_

[Node.js](https://nodejs.org/en/)

Node.js is an extremely popular language for backend development. It has a huge community and tons of libraries available, since it is based on JavaScript. It can also be beneficial to use JavaScript for both the backend and frontend, in case developers are not familiar with other backend languages.

Being a [single-threaded](https://nodejs.org/en/docs/guides/event-loop-timers-and-nexttick/), interpreted language, Node may not be ideal for high-performance applications.


### Frontend
While we offer ways to bootstrap your frontend application in various languages, you can also choose to use whichever language you like and just make use of the infrastructure we set up, or use your custom frontend with a bootstrapped backend.

The frontend application will have a build and deploy pipeline that will build an artifact that will be uploaded to S3, and then hosted through AWS' Cloudfront CDN. This allows your application to be stable and fast from all over the world.

**Language**

[React.js](https://reactjs.org/)

One of the most popular modern libraries for building dynamic frontend applications. React is hugely popular, powerful, and is an obvious choice for most frontends.


### CI

We advocate for (and configure by default) build pipelines for deploying your application. As opposed to deploying manually via the command line which may work fine for a single developer, when you are working on a team there are huge benefits to making GitHub your source of truth, and relying on your version control system to trigger builds and deployments. Our preferred pipeline looks something like this:
```
merge to main branch -> build artifacts -> run tests -> deploy to staging -> ask for user input -> deploy to production
```
But you can also configure it to fit your desired workflow.

[CircleCI](https://circleci.com/)

CircleCI is a powerful tool for creating CI/CD pipelines. They have a free tier which is fully featured, with paid plans that add more concurrency and speed for builds. It requires a bit more setup at the beginning, as you need to create an account with them, but it does have some additional features compared to GitHub Actions.

_- OR -_

[GitHub Actions](https://github.com/features/actions)

GitHub Actions is newer than CircleCI, so it's features are a bit more limited, but the fact that it is integrated with GitHub makes it much easier to get started with. Also similar, it has a free tier with paid versions that add more concurrency. GHA would be useful for projects that don't require a lot of advanced features, visualization, etc.

### Fundamentals
**[Twelve-Factor App](https://12factor.net/)**

The Twelve-Factor App is a methodology for building Software-as-a-Service that encourages building applications in a way that inherently makes them more stable, secure, failure-safe, easier to measure, and easier to reason about. As much as possible we try to stick with these principles when building Zero, as we have seen it to be a very effective way to build scalable, manageable applications.

**[CNCF](https://www.cncf.io/)**

The Cloud Native Computing Foundation hosts and promotes a bunch of great open source software, including Kubernetes, Prometheus, Helm, Fluentd, and Envoy. Their principles and the design of many of the CNCF projects really resonate with the team behind Zero, and we try to make use of CNCF-backed tools whenever possible.
