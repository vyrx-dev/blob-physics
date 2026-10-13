{pkgs ? import <nixpkgs> {}}: let
  libraries = with pkgs; [
    libffi
    libGL
    libxkbcommon
    wayland
    libx11
    libxcursor
    libxi
    libxinerama
    libxrandr
  ];
in
  pkgs.mkShell {
    packages = with pkgs; [
      go
      pkg-config
    ];

    buildInputs = libraries;

    shellHook = ''
      export LD_LIBRARY_PATH=${pkgs.lib.makeLibraryPath libraries}:$LD_LIBRARY_PATH
      echo "blob-physics dev shell ready."
      echo "  CGO_ENABLED=0 go run .   (Pure Go, no C compiler needed)"
      echo "  go test -tags test ./... (Run tests)"
    '';
  }
