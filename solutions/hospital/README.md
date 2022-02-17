# Hostital information system

## Requirements

**Install protobuf:**

See https://developers.google.com/protocol-buffers/

**Install golang binding of protobuf:**

    go get -d -u github.com/golang/protobuf/protoc-gen-go

**Install tooling for mocking:**

    go get github.com/golang/mock/mockgen

## Building

    go generate ./...
    
    