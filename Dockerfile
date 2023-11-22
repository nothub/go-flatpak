FROM debian:12-slim

RUN DEBIAN_FRONTEND=noninteractive             \
    apt-get update                             \
 && apt-get upgrade -y --with-new-pkgs         \
 && apt-get install -y --no-install-recommends \
      ca-certificates                          \
      golang-1.19-go                           \
      gobject-introspection                    \
      libatk1.0-dev                            \
      libflatpak-dev                           \
      libgtk-3-dev                             \
      libgtk-4-dev                             \
 && apt-get clean      -y                      \
 && apt-get autoremove -y                      \
 && rm -rf /var/lib/apt/lists/*

ENV PATH="/usr/lib/go-1.19/bin/go:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

WORKDIR /work
