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
          shellHook =
            ''
              set -a # automatically export all variables
              source .env
              set +a
            '';
        };

      packages.${system} = {
        default =
          with import nixpkgs { inherit system; };
          pkgs.buildGoPackage {
            inherit name;
            src = ./.;
            goPackagePath = "github.com/rudesome/${name}";
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
            copyToRoot =
              with pkgs.dockerTools;
              pkgs.buildEnv {
                name = "image-root";
                paths = [
                  self.packages.${system}.default
                  caCertificates
                ];
              };
          };
      };
    };
}
