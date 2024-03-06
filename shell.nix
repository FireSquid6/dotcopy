let
  unstable = import <nixos-unstable> { config = { allowUnfree = true; }; };
in
{ pkgs ? import <nixpkgs> { } }:
with pkgs; mkShell {
  buildInputs = [
    nodejs_20
    inotify-tools
    libnotify
    go_1_20
    unstable.bun
    wget
  ];
}
