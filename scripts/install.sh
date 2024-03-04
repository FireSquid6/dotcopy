#!/usr/bin/env bash

if ! [ -x "$(command -v wget)" ]; then
  echo 'Error: wget is not installed. Either install it or just manually download the binary' >&2
  exit 1
fi

# we assume linux by default. If we're on mac then we change it.
PLATFORM="linux"
if [ $(uname -s) == "Darwin" ]; then
  PLATFORM="darwin"
fi

ARCHITECTURE=""
case $(uname -m) in
  i386) ARCHITECTURE="386" ;;
  i686) ARCHITECTURE="386" ;;
  x86_64) ARCHITECTURE="amd64" ;;
  arm) ARCHITECTURE="arm" ;;
  aarch64) ARCHITECTURE="arm" ;;
esac

# if this script is broken then this line is probably the reason
# sorry future self
VERSION=$(curl -s https://api.github.com/repos/FireSquid6/dotcopy/releases/latest | grep tag_name | cut -d ":" -f 2 | sed 's/\"//g' | sed 's/,//g' | sed 's/ //g')

echo $VERSION 

tarball="dotcopy-$VERSION-$PLATFORM-$ARCHITECTURE.tar.gz"

# https://github.com/FireSquid6/dotcopy/releases/download/v0.2.8/dotcopy-v0.2.8-linux-amd64.tar.gz
wget https://github.com/FireSquid6/dotcopy/releases/download/$VERSION/$tarball
mkdir -p ~/.dotcopy
tar -xvf $tarball -C ~/.dotcopy
rm $tarball

echo -e "\n\n\n"
echo -e "\033[32m Dotcopy installed successfully to ~/.dotcopy \033[0m"
echo -e "\033[32m Add ~/.dotcopy to your PATH by putting: \033[0m"
echo -e "     export PATH=\$PATH:~/.dotcopy"
echo -e "\033[32m in your .bashrc or .zshrc \033[0m"
