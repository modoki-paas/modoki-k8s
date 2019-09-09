PROTOC = protoc
GO111MODULE = on

.DEFAULT_GOAL := modokid

modokid:
	go build -o modokid $(wildcard ./daemon/*.go)

.PHONY: docker
docker:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: generate
generate: clean
	cd ./design && $(PROTOC) --go_out=plugins=grpc:../api *.proto

.PHONY: clean
clean:
	rm ./api/*.pb.go