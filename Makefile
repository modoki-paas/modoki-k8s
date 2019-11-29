PROTOC = protoc
GO111MODULE = on
DOCKER_APISERVER_IMAGE = modokipaas/modoki-k8s
DOCKER_YAMLER_IMAGE = modokipaas/modoki-yamler
DOCKER_BUILDKIT = 1

.DEFAULT_GOAL := modokid

.PHONY: modokid
modokid: 
	go build -o modokid $(wildcard ./apiserver/*.go)

.PHONY: all
all: modokid docker test

.PHONY: docker
docker:
	docker build -t $(DOCKER_APISERVER_IMAGE) .
	docker build -t $(DOCKER_YAMLER_IMAGE) .

.PHONY: test
test:
	go test -race -tags use_external_db -v ./...

.PHONY: docker-push
docker-push:
	if [ "$(CIRCLE_BRANCH)" = "master" ]; then\
		docker push $(DOCKER_APISERVER_IMAGE);\
		docker push $(DOCKER_YAMLER_IMAGE);\
	fi

	docker tag $(DOCKER_APISERVER_IMAGE) $(DOCKER_APISERVER_IMAGE):$(CIRCLE_SHA1)
	docker push $(DOCKER_APISERVER_IMAGE):$(CIRCLE_SHA1)

	docker tag $(DOCKER_YAMLER_IMAGE) $(DOCKER_YAMLER_IMAGE):$(CIRCLE_SHA1)
	docker push $(DOCKER_YAMLER_IMAGE):$(CIRCLE_SHA1)

.PHONY: generate
generate: clean
	cd ./design && $(PROTOC) --go_out=plugins=grpc:../api *.proto

.PHONY: clean
clean:
	rm ./api/*.pb.go || true