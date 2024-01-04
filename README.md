# dotcopy

Dotcopy tries to solve the following problem:

> I want to have the same dotfiles across all of my machines, but with a couple change between them

# Example

Let's imagine the following dotfile (it's my kitty config):

```conf
editor              nvim
font_family 	      FiraCode Nerd Font

background #1a1b26
foreground #c0caf5
selection_background #283457
selection_foreground #c0caf5
url_color #73daca
cursor #c0caf5
cursor_text_color #1a1b26

# Tabs
active_tab_background #7aa2f7
active_tab_foreground #16161e
inactive_tab_background #292e42
inactive_tab_foreground #545c7e
#tab_bar_background #15161e

# Windows
active_border_color #7aa2f7
inactive_border_color #292e42

...
```

Let's also say I have two machines: `laptop` and `desktop`. On my desktop I want to have the font as Hasklug, while on my laptop I want it to be FiraCode. This could be done with two different configs, but I want everything else to be the same and track my dotfiles with git. Dotcopy comes to the rescue!

I define the following dotcopy config:

```yaml
files:
  - ".config/kitty/kitty.conf":
    template: "kitty.conf"
    slots:
      - "{font}":
        laptop: "FiraCode Nerd Font"
        desktop: "Hasklug Nerd Font"
```

Now, if I replace the font parameter in a local kitty.conf, I can run `dotcopy build` to build the kitty.conf and put it in the correct place. Dotcopy will automatically detect which profile to use with either the `-p` tag or the system hostname
