{
  description = "Go pulumi importer";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.11";
  };

  outputs = { self, nixpkgs }:
    let
      name = "pulumi-import-state";
      system = "x86_64-linux";
    in
    {

      packages.${system}.default =
        with import nixpkgs { inherit system; };
        pkgs.buildGoPackage {
          inherit name;
          src = ./.;
          goPackagePath = "github.com/rudesome/${name}";
          vendorHash = null;
        };

      devShells.${system}.default =
        with import nixpkgs
          {
            inherit system;
          };
        pkgs.mkShell {
          buildInputs = with pkgs; [
            (pulumi.withPackages (p: with p; [
              pulumi-language-go
            ]))
          ];
        };
    };
}
