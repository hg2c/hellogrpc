.PHONY: build
build:
	@./scripts/binary.sh

.PHONY: dev
dev:
	@./scripts/docker.sh

.PHONY: vendor
vendor:
	dep ensure

.PHONY: image
image:
	./scripts/image.sh

.PHONY: deploy-image
deploy-image:
	./scripts/deploy-image.sh

.PHONY: cross
cross:
	@./scripts/cross.sh

.PHONY: proto
proto:
	protoc -I helloworld/ --go_out=plugins=grpc:helloworld helloworld/helloworld.proto
	protoc -I customer/proto/ --go_out=plugins=grpc:customer/proto customer/proto/customer.proto
