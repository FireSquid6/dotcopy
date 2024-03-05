{ lib
, buildGoModule
, fetchFromGitHub
}:

buildGoModule rec {
  pname = "dotcopy";
  version = "0.2.9";

  src = fetchFromGitHub {
    owner = "firesquid6";
    repo = "dotcopy";
    rev = "v${version}";
    hash = "sha256-33cH8Yz2cMZzaoalniRjwy6ooAmy8rhQqf9ZeprpklA=";
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
