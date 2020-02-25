APP ?= zipkin-client
RELEASE ?= 0.0.4

GOOS ?= linux
GOARCH ?= amd64

COMMIT ?= $(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

CONTAINER_IMAGE ?= docker.io/bjornsv/${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
	-ldflags="-s -w -X 'main.Release=${RELEASE}' \
			-X 'main.Commit=${COMMIT}' \
			-X 'main.BuildTime=${BUILD_TIME}'" \
	-o ${APP}

clean:
	rm -rf ${APP}

test:
	go test -v -race ./...

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) . || true

push:
	docker push $(CONTAINER_IMAGE):$(RELEASE)
