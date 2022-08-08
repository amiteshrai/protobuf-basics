# Protocol Buffers and Go

## Setup

Install following:

- Golang
- Protoc Compiler

        brew install protoc-gen-go

- VSCode
- The VSCode Golang [extension](https://code.visualstudio.com/docs/languages/go)
- Vscode-proto3 [extension](https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3) (optional)
- Golang package:

        export GOMODULE=on
        go install google.golang.org/protobuf/cmd/protoc-gen-go
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

## Execute

    make
    ./protobuf-basics

    OR
    protoc -Iproto --go_opt=module=github.com/amiteshrai/protobuf-basics --go_out=. proto/*.proto
    go build .
    ./protobuf-basics
