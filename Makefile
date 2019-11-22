SHELL := /bin/bash
PROTO_FILES := $(wildcard protobuf/*.proto)
PROTO_DEFS  := $(PROTO_FILES:.proto=.pb.go)
README := $(wildcard *.md)

main:
	go build -o main

.PHONY: proto
proto: $(PROTO_DEFS)

protobuf/%.pb.go: protobuf/%.proto
	protoc --twirp_out=. --go_out=. $<

test: $(README)
	echo $<



.PHONY: clean
clean:
	rm -f protobuf/*.go
