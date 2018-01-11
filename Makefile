.PHONY: build
build:
	./scripts/build.sh

.PHONY: vendor
vendor:
	dep ensure

.PHONY: image
image:
	./scripts/image.sh
