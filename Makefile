PROTOC = protoc
GO111MODULE = on
DOCKER_APISERVER_DOCKERFILE = Dockerfile-api
DOCKER_YAMLER_DOCKERFILE = Dockerfile-yamler
DOCKER_AUTHSERVER_DOCKERFILE = Dockerfile-auth
DOCKER_APISERVER_IMAGE = modokipaas/modoki-k8s
DOCKER_YAMLER_IMAGE = modokipaas/modoki-yamler
DOCKER_AUTHSERVER_IMAGE = modokipaas/modoki-auth

DOCKER_DOCKERFILE=$(DOCKER_APISERVER_DOCKERFILE)
DOCKER_IMAGE=$(DOCKER_APISERVER_IMAGE)

DOCKER_BUILDKIT = 1

.DEFAULT_GOAL := modokid

.PHONY: install_dependencies
install_dependencies:
	curl -s "https://raw.githubusercontent.com/\
	kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash

.PHONY: modokid
modokid: 
	go build -o modokid $(wildcard ./apiserver/*.go)

.PHONY: all
all: modokid docker-all test

.PHONY: docker
docker:
	docker build -t $(DOCKER_IMAGE) -f $(DOCKER_DOCKERFILE) .

.PHONY: docker-push
docker-push:
	if [ "$(CIRCLE_BRANCH)" = "master" ]; then\
		docker push $(DOCKER_IMAGE);\
	fi

	docker tag $(DOCKER_IMAGE) $(DOCKER_IMAGE):$(CIRCLE_SHA1)

.PHONY: docker-all
docker-all: DOCKER_IMAGE=$(DOCKER_APISERVER_IMAGE) DOCKER_DOCKERFILE=$(DOCKER_APISERVER_DOCKERFILE)
docker-all: docker
docker-all: docker-push
docker-all: DOCKER_IMAGE=$(DOCKER_YAMLER_IMAGE) DOCKER_DOCKERFILE=$(DOCKER_YAMLER_DOCKERFILE)
docker-all: docker
docker-all: docker-push
docker-all: DOCKER_IMAGE=$(DOCKER_AUTHSERVER_IMAGE) DOCKER_DOCKERFILE=$(DOCKER_AUTHSERVER_DOCKERFILE)
docker-all: docker
docker-all: docker-push

.PHONY: test
test:
	go test -race -tags use_external_db -v ./...


.PHONY: generate
generate: clean
	cd ./design && $(PROTOC) --go_out=plugins=grpc:../api *.proto

.PHONY: clean
clean:
	rm ./api/*.pb.go || true