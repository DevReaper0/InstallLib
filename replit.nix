{ pkgs }: {
    deps = [
        pkgs.go_1_18
        pkgs.gopls
        pkgs.pkg-config
        pkgs.xlibsWrapper
        pkgs.xorg.libX11.dev
        pkgs.xorg.libXcursor
        pkgs.xorg.libXrandr
        pkgs.xorg.libXinerama
        pkgs.xorg.libXi
        pkgs.xorg.libXxf86vm
        pkgs.libGL
        pkgs.glfw
        pkgs.unzip
    ];
}