{
  description = "Description for the project";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
      ];
      systems = [ "x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin" ];
      perSystem = {pkgs, ... }: {
				devShells.default = pkgs.mkShell {
					name = "go";
					packages = with pkgs; [
						go just zsh sqlc cobra-cli zip gopls gofumpt
					];
					shellHook = ''
						export SHELL=${pkgs.zsh}/bin/zsh
						exec $SHELL
						'';
				};
      };
    };
}
