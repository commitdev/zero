---
title: Working on Zero
sidebar_label: Working on Zero
sidebar_position: 3
---

### Building the tool

```shell
$ git clone git@github.com:commitdev/zero.git
$ cd zero && make build
```

### Releasing a new version on GitHub and Brew

We are using a tool called `goreleaser` which you can get from brew if you're on MacOS:
`brew install goreleaser`

After you have the tool, you can follow these steps:
```
export GITHUB_TOKEN=<your token with access to write to the zero repo>
git tag -s -a <version number like v0.0.1> -m "Some message about this release"
git push origin <version number>
goreleaser release
```