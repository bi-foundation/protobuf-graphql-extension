FROM quay.io/opsee/build-go:16

COPY ./ /gopath/src/github.com/bi-foundation/protobuf-graphql-extension

RUN cd /gopath/src/github.com/bi-foundation/protobuf-graphql-extension && \
  go install ./opseeproto && \
  go install ./plugin/... && \
  go install ./protoc-gen-gogoopsee
