{ self, ... }:
{
  perSystem =
    { pkgs, ... }:
    let
      version = "1.1.1";
      ldflags = [
        "-s"
        "-w"
        "-X github.com/z3co/prot/cmd.Version=${version}"
        "-X github.com/z3co/prot/cmd.Commit=${self.rev or "dirty"}"
      ];
    in
    {
      packages.dev = pkgs.buildGoModule {
        pname = "prot";
        inherit version;
        src = ./.;
        vendorHash = "sha256-1xQSQTUZUzykz8YXVnIp5bImU9cJCODiA3cWeb852w0=";
      };
			packages.default = pkgs.buildGoModule {
				pname = "prot";
				inherit version ldflags;
				src = pkgs.fetchFromGitHub {
					owner = "z3co";
					repo = "protv2";
					tag = "v1.0.1";
					hash = "sha256-YSxmIzbgAyFO9kVlLA6wUt1vwO2pIiaN970rGp+8/cw=";
				};
				vendorHash = "sha256-1xQSQTUZUzykz8YXVnIp5bImU9cJCODiA3cWeb852w0=A";
			};
    };
}
