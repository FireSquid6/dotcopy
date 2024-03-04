{ lib
, buildGoModule
, fetchurl
}:

buildGoModule rec {
  pname = "dotcopy";
  version = "0.2.8";

  src = fetchurl {
    url = "https://github.com/firesquid6/dotcopy/releases/download/v${version}/dotcopy-v${version}-linux-amd64.tar.gz";
    hash = "sha256-cfd5e9d0634fec0ea965dc656dfd21be";
  };

  vendorHash = "sha256-cTpDJhcw0JUClpZVEGchWGNSKPX2zhZW+MVGgWrxcrY=";

  buildPhase = ''
    tar -xvf dotcopy-v${version}-linux-amd64.tar.gz -C ./
  '';
  installPhase = ''
    mkdir -p $out/bin
    cp dotcopy $out/bin
  '';


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
