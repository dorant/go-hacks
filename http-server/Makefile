APP ?= go-http
PORT ?= 8000

GOOS ?= linux
GOARCH ?= amd64

PROJECT ?= github.com/dorant/go-http
RELEASE ?= 0.0.1
COMMIT ?= $(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

CONTAINER_IMAGE ?= docker.io/bjornsv/${APP}

build: clean
	cd cmd/go-http && CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
	-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
			-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
	-o ${APP}

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) . || true

clean:
	cd cmd/go-http && rm -rf ${APP}

run: container
	docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm -e "PORT=${PORT}" $(APP):$(RELEASE)

test:
	go test -v -race ./...

push:
	docker push $(CONTAINER_IMAGE):$(RELEASE)

minikube: push
	for t in $(shell find ./k8s -type f -name "*.yaml"); do \
	cat $$t | \
	sed -E "s/\{\{ \.Release \}\}/$(RELEASE)/g" | \
	sed -E "s/\{\{ \.ServiceName \}\}/$(APP)/g"; \
	echo ---; \
	done > tmp.yaml
	kubectl apply -f tmp.yaml
