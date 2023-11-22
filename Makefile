user = $(shell whoami)

.PHONY: all generate permissions docker docker-image docker-generate

all: generate permissions

generate:
	sudo go generate -run go run . -o ./pkg/

permissions:
	sudo chown -R $(user):$(user) pkg

docker: docker-image docker-generate permissions

docker-image:
	docker build --progress "plain" -t "go-flatpak-builder" $(CURDIR)

docker-generate:
	docker run -it --rm -v "$(CURDIR):/work" "go-flatpak-builder" /usr/lib/go-1.19/bin/go generate
