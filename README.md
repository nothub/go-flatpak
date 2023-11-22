# go-flatpak

Autogenerated `libflatpak` bindings for Go.

[![Go Documentation](https://godocs.io/github.com/erazemk/go-flatpak/pkg/flatpak?status.svg)](https://godocs.io/github.com/erazemk/go-flatpak/pkg/flatpak)

## Building

```shell
# Install dependencies (dnf)
sudo dnf install -y atk-devel flatpak-devel gobject-introspection-devel gtk3-devel gtk4-devel
# Install dependencies (apt)
sudo apt install -y libatk1.0-dev libflatpak-dev gobject-introspection libgtk-3-dev libgtk-4-dev

# Generate library
sudo "$(which go)" generate

# Restore directory ownership (since running as root is required to walk /usr/)
sudo chown -R "$(id -u):$(id -g)" pkg
```

### Docker

```shell
# Build builder image
docker build --progress "plain" -t "go-flatpak-builder" .

# Generate library
docker run -it --rm -v "${PWD}:/work" "go-flatpak-builder" /usr/lib/go-1.19/bin/go generate

# Restore directory ownership
sudo chown -R "$(id -u):$(id -g)" .
```
