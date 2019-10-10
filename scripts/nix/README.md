
# Docker-less reproducible builds with Nix

commit0 supports docker-less reproducible builds via the Nix package manager.
Nix is a platform-agnostic way to do reproducible builds without needing a
virtual machine.


## Instructions

Install nix:

    $ curl https://nixos.org/nix/install | sh

Run nix build from the root:

    $ nix build -f scripts/nix

Run commit0:

    $ ./result-bin/bin/commit0 --help

or install it into your local environment:

    $ nix-env -i ./result-bin


## Updating dependencies

go dependencies are determinstically generated from go.mod using
[vgo2nix](https://github.com/adisbladis/vgo2nix). Just run vgo2nix to update the
deps.nix file.
