let
  unstable = import <nixos-unstable> { config = { allowUnfree = true; }; };
in
{ pkgs ? import <nixpkgs> { } }:
with pkgs; mkShell {
  buildInputs = [
    nodejs_20
    inotify-tools
    libnotify
    go
    unstable.bun
    wget
  ];
}
