# dotcopy

Dotcopy tries to solve the following problem:

> I want to have the same dotfile templates across all of my machines, but with a couple change between them

It allows you to make your dotfiles DRY across mutliple machines.

## CLI

```txt
NAME:
   dotcopy - Builds your dotfiles. See https://dotcopy.firesquid.co

USAGE:
   dotcopy [global options] command [command options]

COMMANDS:
   init     Initializes a basic localconfig
   version  Prints the version of dotcopy
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --silent -s -d, -s -s -d         Silence all output to stdout. Use -s -d to silence all out
put. (default: false)
   --disable-notifications, -d, -n  Disable system notifications (default: false)
   --help, -h                       show help
```

# Install

You can install dotcopy using an install script like so:

```
curl -fsSL https://dotcopy.firesquid.co/install.sh | bash
```

Additionally, dotcopy has a nixkpg called `dotcopy`. You can install that however you'd like.

# Walkthrough

## Initializing

Dotcopy depends on a file called `localconfig.yaml` in `~/.config/dotcopy`. It should look something like this:

```yaml
root_filepath: /etc/nixos/dotfiles # where your dotfiles are stored. I keep mine in my nixos folder.
machine_directory: kotoko # Your machine's hostname most likely
```

This file should be different on each of your machines. If it is not found, dotcopy uses this config:

```
root_filepath: <HOME_DIRECTORY>/dotfiles
machine_directory: <HOSTNAME>
```

**Note: all yaml files need to use `.yaml` and not `.yml`. I am rejected the idea that one file format needs two extensions**

## The Dotfiles Directory

The dotfiles directory should have a heirchy that looks something like this:

```
dotfiles/
  dotcopy.yaml

  i3
  kitty.conf
  < All of my other dotfile templates >
  <machine-name>/
    i3.slot
    kitty.slot.conf
```

All of your dotfiles go in the root, while the "slots" for them go inside separate folders for each machine

### The `dotcopy.yaml` file

The `dotcopy.yaml` file tells dotcopy where your dotfiles go. It should look like:

```yaml
- template: i3 # filepath relative to the root
  slotfile: i3.slot # filepath relative to the machine directory
  location: /home/firesquid/.config/i3/config # where to put the compiled file
- template: kitty.conf
  slotfile: "" # my kitty.conf doesn't need to change, so the slotfile is nothing
  location: /home/firesquid/.config/kitty/kitty.conf
# ...
# more dotfiles if necessary
```

## Compiling Dotfiles

Dotcopy will analyze a template file for slots and then insert into them. Let's see an example.

Imagine I have this dotfile called `font-file.txt` in my root:

```txt
# font-file
# a magical dotfile that changes your font size

font-size: 10

# pretend there's a bunch of complicated
```

On my machine `chisato`, I want to have `font-size` be 10, but on `kotoko`, I want it to be 12. First, I change the template file in the root:

```
# font-file
# a magical dotfile that changes your font size

font-size: {{font-size}}

# pretend there's a bunch of complicated
```

I also define this in my `dotcopy.yaml`:

```
- template: font-file.txt
  slotfile: font-file.slot.txt
  location: ~/.config/font-file/config.txt
```

Now in my `chisato` machine directory, I create a file called `font-file.slot.txt`. It looks like:

```
--- {{font-size}}
10
---

# I could add more values using that same syntax if I wanted to
```

You can repeat this process for all of your dotfiles.

# FAQ

## Why use this when home manger exists?

Home manager is great! However, it has some things that make it not perfect for everyone:

- it only works on nix
- it forces everything to be in nix (not a bad thing, just not what everyone wants)
- you can't do the different stuff for different machines with home-manager (ig you could import stuff from a different file in gitignore maybe but that's a pain)

## Does this waste disk space?

Yes - dotcopy does waste some disk space by copying your dotfiles. However, it shouldn't be more than a couple of megabytes.
