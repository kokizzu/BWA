
# gRPC example

1. install dependencies, start dependencies
```
go mod tidy
go mod vendor

docker-compose up
```

2. install [protobuf and grpc](https://grpc.io/docs/languages/go/quickstart/#prerequisites)
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

3. create a proto file (eg. user.proto) then compile it
```
cd rpcp
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    user.proto
```

4. register a new service on `main.go`, implement the `func (*hndlr) *Grpc(ctx,in) (out,err)` method on the handler, refactor the old method so common function can be called from REST version and GRPC version
5. create client to test `rpcp/client/user_test.go`

If you don't need gRPC streaming, it's better to use: [twirp](https://blog.twitch.tv/en/2018/01/16/twirp-a-sweet-new-rpc-framework-for-go-5f2febbf35f/) instead of gRPC, it can handle both REST and gRPC-like (still using protobuf) in one method, so can be called from web browser.
