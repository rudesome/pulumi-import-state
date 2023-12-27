{
  description = "Go pulumi importer";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.11";
  };

  outputs = { self, nixpkgs }:
    let
      name = "pulumi-import-state";
      system = "x86_64-linux";
      version = "latest";
    in
    {

      devShells.${system}.default =
        with import nixpkgs { inherit system; };
        pkgs.mkShell {
          buildInputs = with pkgs; [
            (pulumi.withPackages (p: with p; [
              pulumi-language-go
            ]))
          ];
        };

      packages.${system} = {
        default =
          with import nixpkgs { inherit system; };
          pkgs.buildGoPackage {
            inherit name;
            src = ./.;
            goPackagePath = "github.com/rudesome/${name}";
            subPackages = [ "cmd" ];
            vendorHash = null;
          };

        ## build docker image
        docker =
          with import nixpkgs { inherit system; };
          pkgs.dockerTools.buildImage {
            inherit name;
            tag = version;
            config = {
              Cmd = [ "${self.packages.${system}.default}/bin/cmd" ];
            };
          };
      };
    };
}
