VANITY_DIR = opseeproto
EXAMPLE_DIR = examples
PROTO_DIR = proto

all: EXAMPLE_DIR

VANITY_DIR:
	docker run --rm -it -v $$(pwd):/build quay.io/opsee/build-go:16 /bin/bash -c 'cd /build/$(VANITY_DIR) && make generate'

docker: clean VANITY_DIR
	docker build -t quay.io/opsee/build-go:proto16 .

EXAMPLE_DIR: docker
	docker run --rm -it -v $$(pwd):/gopath/src/github.com/bi-foundation/protobuf-graphql-extension quay.io/opsee/build-go:proto16 /bin/bash -c 'cd /gopath/src/github.com/bi-foundation/protobuf-graphql-extension/$(EXAMPLE_DIR) && make generate && go test -v ./...'

push:
	docker push quay.io/opsee/build-go:proto16

clean:
	$(MAKE) -C $(EXAMPLE_DIR) clean

.PHONY:
	docker
	clean
	push
	all
