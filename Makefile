.PHONY: default build image check publish-images

default: fmt test build

test:
	go test -v -cover ./...

build:
	CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o cobd

image:
	docker build -t banst/cobd .

fmt:
	go fmt .
