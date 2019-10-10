{ nixpkgs ? import (builtins.fetchTarball {
              name = "nixos-19.09-2019-10-10";
               url = "https://github.com/nixos/nixpkgs/archive/9bbad4c6254513fa62684da57886c4f988a92092.tar.gz";
               sha256 = "00dhkkmar3ynfkx9x0h7hzjpcqvwsfmgz3j0xj80156kbw7zq4bb";
             }) {}
, buildGoPackage ? nixpkgs.buildGoPackage
}:
let
  filter = path: type:
    !(type == "directory" && baseNameOf path == "example") &&
    !(type == "file" && baseNameOf path == "commit0");
in
buildGoPackage rec {
  name = "commit0-${version}";
  version = "0";
  goPackagePath = "github.com/commitdev/commit0";
  src = builtins.filterSource filter ../..;
  goDeps = ./deps.nix;
}
