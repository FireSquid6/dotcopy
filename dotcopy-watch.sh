#!/usr/bin/env bash

on_update () {
  notify-send "Dotcopy: File change detected, rebuilding..."

  echo "\n"
  ./dotcopy build
  echo "\n"

  notify-send "Dotcopy: Rebuilt dotfiles. Make sure to reload i3, polybar, and other programs to see the changes."
}

while true; do
inotifywait -e modify,create,delete -r $1 && on_update 
done


