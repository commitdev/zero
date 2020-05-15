![zero](https://github.com/commitdev/zero/blob/master/docs/img/logo-small.png?raw=true)

# Design

**In this document `zero` is a placeholder for some named utility.**

The guiding principle behind `zero` is to **make it easy for developers to ship
to production on day 1**

## Developer Experience (DX)

Developer Experience is the equivalent of User Experience when the primary user
of the product is a developer. DX cares about the developer experience of using
a product, its libs, SDKs, documentation, frameworks, open-source solutions,
general tools, APIs, etc.

**DX Pillars**

- Function: Something is only as useful as it is functional. If it doesn't work
  it's of no value. Maintaining function is a key element of quality and
  percetpion of quality.
- Stability: Performance and reliability
- Ease of Use: The tool alone only forms one part of the developers journey.
  Resoures like github issues, documentation, clear workflows, snippets, etc
  all help reduce the friction in using and learning a new tool.
- Clarity: Clear, concise, actionable information makes the world of difference
  when using a tool.

## DX Guidelines

- easy onboarding. it must be easy for someone to get started with the tool,
  and to join an existing \$setup or create a new one.
- follows [principle of least astonishment][1]
- use short easy to remember commands
- help commands on [every sub command][2]
- pipe-able / machine-readable output
- feedback at the end of a command. every tool and command is guiding someone
  down a path and we should tell them where to go next. for example, a command
  that results in some build artifact; tell them the typical next step. github
  does the same thing with their post remote push message telling you to open a
  PR. any command that takes a long time to run should let the user know though
  some kind of visual feedback. for scripts, this feedback should be able to be
  suppressed, like curl silent mode.
- colors; when do we use them? anything needs to work _without_ colors too, so
  they can't be the primary way of labeling / indicating anything. just an
  enhancement to words or symbols that are already on the screen. For an
  example of things to think about [see this mocha issue][3]

## Implementation Principles

- Discoverable:
- Familiar:
- Alterable:
- Parsable:
- Consistency: as a command line tool we need to establish a convention and
  stick to it. things like `[cmd] [noun] [verb]` with flags. consistent flag options,
  etc. short form flags are possible but we should encourage long form flags
  when scripting usage of the tool so that things are more self-documenting.
- error messages should be easy to understand and steer someone towards an
  answer / next step.
- error early: if we can detect things that will be problems do it as early as
  possible before proceeding with an operation. it makes no sense to take
  expensive actions that may mutate state which will only half-finish
- idempotency

## Users

- Our users are not power users
- typically experienced but not devops masters
- havent built their own full environment before or they have but are missing
  best practices / some clear conventions.
- Users will more often be joining a team or a project much more frequently
  than creating an environment from scratch. That only happens one time.

## Installation

- must be easy to install for users of different platforms with the least
  amount of lift
  linux, osx. should look at mechanisms for installation like brew, apt, curl
- dont have to worry about windows for now but we shouldn't back ourselves into
  a corner

## Documentation

Documentation should be consistent, easy to follow, and follow the user through
their own journey. We should only make use of commonly installed utilities which
come with the prerequisites we define or commands we _know_ are installed. For
example use `psql` vs `pgcli` in shell examples.

Each README should have an outline on:

- what the tool does
- which dependencies to install
- how to install the application
- how to contribute
- communication channels

To ensure the README is easy to read, please keep length limited. Any FAQ
sections or hints / tips / design decisions / etc should go into a separate
documentation path.

Shell commands should be preceeded with a `$` dollar sign to act as a visual
representation of the shell. While the world will be moving towards `zsh` we
should use compatible examples whenever possible. Shell commands should wrap at
the 80 character limit. For long commands use a `\` to move additional arguments
or flags onto separate lines, and omit the preceeding `$` so that it can be
copy-pasted without the dollar sign affecting interpretation.

```shell
$ zero arg \
  --flag-one \
  --flag-two \
  three
```

## Concepts

#### Help

built in help should follow standard conventions. `[option]` with square
brackets. `<argumennts>` which are required with greater-than and lesser-than
signs.

```shell
$ zero --help
$ zero noun [[]noun] verb --help
$ zero noun verb [options] --help
$
$ Usage: noun verb [options] <arg> <arg>
$ # flags list
```

or

```shell
$ zero <command> <subcommand>
```

## Glossary

Terms we'll use repeatedly should have clear definitions or use common knowledge
definitions. Some of these terms may still require definitions. Any frequently
used word should be added to this list if it applies to usage in `zero`.

- **Project**: A project is the top level (root) entity that contains the infrastructure and configuration. A project is compartmentalized and has no awareness of the details of other projects. Environments are built using Modules as part of a project.

- **Module**: This is a git repository that contains everything Zero needs to set up a piece of a project.
  - What’s included in a single Module?
    - Templates
    - Documentation
    - Potentially definition of something to execute (terraform, api calls, etc)
  - Examples of Current modules:
    - Zero-aws-eks-stack
    - Zero-deployable-backend
    - Zero-deployable-react-frontend

- Application all the pieces of your project working together (Frontend, Service, Infrastructure).

- **Infrastructure**: The systems your application runs on and other required dependencies that are not part of the Frontends or Services. Typically provided by a cloud provider and provisioned with Terraform.  (For example AWS EKS, RDS, Cloudfront, etc.)

- **Environment**: A running instance of your entire Application. Examples being Dev / Staging / UAT / Production (one or many set up for a single project)

- **Frontend**: A single page app serving the front-end of a web application. Typically makes requests to a Service.

- **Service**: A backend app serving APIs or providing funtionality for a web application. Typically serves requests from the Frontend App.

- **Pipeline**: CI/CD pipeline responsible for running automated tests and deploying Frontend Apps and/or Services to Environments.


<!-- links go here -->

[1]: https://en.wikipedia.org/wiki/Principle_of_least_astonishment
[2]: https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-help.html
[3]: https://github.com/mochajs/mocha/issues/802
