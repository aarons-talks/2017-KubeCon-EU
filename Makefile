REPO ?= quay.io/arschles/2017-kubecon-eu-watcher
TAG ?= $(shell git rev-parse --short HEAD)
IMAGE := $(REPO):$(TAG)

build:
	GOOS=linux GOARCH=amd64 go build -o rootfs/bin/watcher
docker-build:
	docker build -t $(IMAGE) rootfs
docker-push:
	docker push $(IMAGE)
