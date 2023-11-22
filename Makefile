.PHONY: all generate permissions docker docker-image docker-generate

all: generate permissions

docker: docker-image docker-generate permissions

generate:
	sudo go generate -run go run . -out ./pkg/

permissions:
	sudo chown -R $(shell id -u):$(shell id -g) pkg

docker-image:
	docker build --progress "plain" -t "go-flatpak-builder" $(CURDIR)

docker-generate:
	docker run -it --rm -v "$(CURDIR):/work" "go-flatpak-builder" /usr/lib/go-1.19/bin/go generate
