{ lib
, buildGoModule
, fetchzip
}:

buildGoModule rec {
  pname = "dotcopy";
  version = "0.2.8";

  src = fetchzip {
    url = "https://github.com/firesquid6/dotcopy/releases/download/v${version}/dotcopy-v${version}-linux-amd64.tar.gz";
    sha256 = lib.fakeSha256;
  };

  vendorSha256 = lib.fakeSha256;

  meta = with lib; {
    description = "A linux dotfile manager";
    homepage = "https://dotcopy.firesquid.co";
    license = licenses.gpl3;
    longDescription = ''
      Dotcopy is a linux dotfile manager that allows you to "compile" your dotfiles to use the same template for multiple machines.
    '';
    maintainers = with maintainers; [ firesquid6 ];
  };
}
