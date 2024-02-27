let
  unstable = import (fetchTarball https://nixos.org/channels/nixos-unstable/nixexprs.tar.xz) { };
in
{ pkgs ? import <nixpkgs> { } }:
with pkgs; mkShell {
  buildInputs = [
    nodejs_20
    inotify-tools
    libnotify
    go
  ];
}
