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

.DEFAULT_GOAL := build

.PHONY: install_dependencies
install_dependencies:
	curl -s "https://raw.githubusercontent.com/\
	kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash

.PHONY: apiserver
apiserver: 
	go build -o modoki-apiserver $(wildcard ./apiserver/*.go)

.PHONY: authserver
authserver: 
	go build -o modoki-authserver $(wildcard ./authserver/*.go)

.PHONY: yamler
yamler: 
	go build -o modoki-yamler $(wildcard ./yamler/*.go)

build: apiserver authserver yamler

.PHONY: all
all: build test docker-all

define docker
	docker build -t $1 -f $2 .
endef

define docker-push
	if [ "$(CIRCLE_BRANCH)" = "master" ]; then\
		docker push $1;\
	fi

	docker tag $1 $1:$(CIRCLE_SHA1)
	docker push $1:$(CIRCLE_SHA1)
endef


.PHONY: docker-api docker-auth docker-yamler
docker-api:
	@$(call docker,$(DOCKER_APISERVER_IMAGE),$(DOCKER_APISERVER_DOCKERFILE))
	@$(call docker-push,$(DOCKER_APISERVER_IMAGE))

docker-auth:
	@$(call docker,$(DOCKER_AUTHSERVER_IMAGE),$(DOCKER_AUTHSERVER_DOCKERFILE))
	@$(call docker-push,$(DOCKER_AUTHSERVER_IMAGE))

docker-yamler:
	@$(call docker,$(DOCKER_YAMLER_IMAGE),$(DOCKER_YAMLER_DOCKERFILE))
	@$(call docker-push,$(DOCKER_YAMLER_IMAGE))

.PHONY: docker-all
docker-all: docker-api docker-auth docker-yamler

.PHONY: test
test:
	go test -race -tags use_external_db -v ./...

.PHONY: generate
generate: clean
	cd ./design && $(PROTOC) --go_out=plugins=grpc:../api *.proto

.PHONY: clean
clean:
	rm ./api/*.pb.go || true