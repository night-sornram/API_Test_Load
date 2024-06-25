# How to run
1. `go mod tidy`
2. Set environment for port using `HTTP1_PORT`, `HTTP2_PORT`, `GRPC_PORT`
3. `go run main.go`

# Example
HTTP 1.1
```shell
curl --location 'localhost:8081/phone?number=0711555001'
```

HTTP 2 (Needs upgrade from 1.1 using h2c)
```shell
curl --location 'localhost:8082/phone?number=0711555001'
```

gRPC
- get spec from `protos/lookup.proto`
- Generate grpc client using protoc
- Implement business logic for client
- Default gRPC server port is 8083