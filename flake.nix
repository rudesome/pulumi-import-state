{
  description = "Go pulumi importer";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [ "aarch64-darwin" "aarch64-linux" "x86_64-darwin" "x86_64-linux" ];
      perSystem = { config, self', inputs', pkgs, system, ... }:
        let
          name = "pulumi-import-state";
        in
        {
          devShells = {
            default = pkgs.mkShell {
              buildInputs = with pkgs; [
                (pkgs.pulumi.withPackages (p: with p; [
                  pulumi-language-go
                ]))
              ];
            };
          };
          packages = {
            default = pkgs.buildGoPackage {
              inherit name;
              src = ./.;
              goPackagePath = "github.com/rudesome/${name}";
              vendorHash = null;
            };
          };
        };
    };
}

