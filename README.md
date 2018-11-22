# grpc-testing
gRPC Testing TLS Auth

### Usage

#### Start Server
```
make server
```

#### Run Client
```
make client
```

### Example

In one terminal
```
make server
```

In another terminal
```
make client
```

#### Example Combined Output

```
$ make server
artifacts/build/debug/darwin/amd64/grpc-testing server
2018/11/22 17:30:01 Received: world
```

```
$ make client
artifacts/build/debug/darwin/amd64/grpc-testing client world
2018/11/22 17:30:01 Greeting: Hello world
```