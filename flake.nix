{
  description = "Go pulumi importer";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.11";
  };

  outputs = { self, nixpkgs }:
    let
      name = "pulumi-import-state";
      system = "x86_64-linux";
      tag = "latest";
      pkgs = import nixpkgs { system = "x86_64-linux"; };
    in
    {

      devShells.${system}.default =
        with pkgs;
        mkShell
          {
            buildInputs = with pkgs; [
              (pulumi.withPackages (p: with p; [
                pulumi-language-go
              ]))
            ];
            shellHook =
              ''
                # allow to import env variable(s)
                set -a
                # source specific environmental variable(s)
                source .env
                # default (do not export new environmental variables)
                set +a
              '';
          };

      packages.${system} = {
        default =
          with pkgs;
          buildGoPackage {
            inherit name;
            src = ./.;
            goPackagePath = "github.com/rudesome/${name}";
            vendorHash = null;
          };

        ## build docker image
        docker =
          with pkgs.dockerTools;
          buildImage {
            inherit name tag;
            config = {
              Cmd = [ "${self.packages.${system}.default}/bin/cmd" "test" ];
            };
            # https://nixos.org/manual/nixpkgs/stable/#ssec-pkgs-dockerTools-helpers
            copyToRoot =
              with pkgs.dockerTools;
              [ caCertificates ];
          };
      };
    };
}
