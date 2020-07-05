### Regenerate grpc

```bash
protoc \
  -I/usr/local/include \
  -I./grpc \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=paths=source_relative:./grpc \
  --go-grpc_out=paths=source_relative:./grpc \
  --grpc-gateway_out=logtostderr=true,paths=source_relative:./grpc \
  **/*.proto
```
