REPO ?= quay.io/arschles/2017-kubecon-eu-watcher
TAG ?= $(shell git rev-parse --short HEAD)
IMAGE := $(REPO):$(TAG)
MUTABLE_IMAGE := $(REPO):devel

build:
	GOOS=linux GOARCH=amd64 go build -o rootfs/bin/watcher
docker-build:
	docker build -t $(IMAGE) rootfs
	docker tag $(IMAGE) $(MUTABLE_IMAGE)
docker-push:
	docker push $(IMAGE)
