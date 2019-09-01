PROTOC = protoc

.PHONY: clean build

modokid:
	go build -o modokid $(wildcard ./daemon/*.go)

docker:
	docker build -t $(DOCKER_IMAGE) .

generate: clean
	cd ./design && $(PROTOC) --go_out=plugins=grpc:../api *.proto

clean:
	rm ./api/*.pb.go