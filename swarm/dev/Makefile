.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/codesferum/go-sferum/cmd/swarm
	go build -o bin/geth     github.com/codesferum/go-sferum/cmd/geth
	go build -o bin/bootnode github.com/codesferum/go-sferum/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/codesferum/go-sferum/swarm/...
