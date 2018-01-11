.PHONY: build
build:
	./scripts/build.sh

.PHONY: deps
deps:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/golang/lint/golint
	go get -u google.golang.org/grpc
	go get -u github.com/golang/mock/gomock
	go get -u github.com/golang/mock/mockgen

.PHONY: image
image:
	./scripts/image.sh
