# dotcopy

Dotcopy tries to solve the following problem:

> I want to have the same dotfiles _templates_ across all of my machines, but with a couple change between them

It allows you to compile your dotfiles in one repository to wherever they need to be. For the developer's version of this, see his [dotfiles folder](https://github.com/FireSquid6/nixos-config/tree/main/dotfiles)

# Installation

For a quick install:

```bash
curl -fsSL https://dotcopy.firesquid.co/install.sh | bash
```

You can also manually download and extract the tarball from the latest release.

## Package Managers

We're looking for maintainers! If you'd like to package dotcopy on your distro I would highly encourage you to do so!

# Documentation

See [firesquid.dotcopy.co](https://firesquid.dotcopy.co) for more information. The source is located in the `/site` directory.

# FAQ

## Why use this when home manger exists?

Home manager is great! However, it has some things that make it not perfect for everyone:

- it only works on nix
- it forces everything to be in nix (not a bad thing, just not what everyone wants)
- you can't do the different stuff for different machines with home-manager (ig you could import stuff from a different file in gitignore maybe but that's a pain)

## Doesn't this waste disk space having nearly duplicate files?

cry me a river about your few kilobytes.
