
{ 
  lib
, buildGo121Module
, fetchFromGitHub
}:

buildGo121Module rec {
  pname = "dotcopy";
  version = "0.2.0";

  src = fetchFromGitHub {
    owner = "firesquid6";
    repo = pname;
    rev = "refs/tags/v${version}";
    sha256 = "WlrBG12SF1a+PpxArcgixZkWLa7t8bq59uCRWzQQtow=";
  };

  subPackages = [ "." ];
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
