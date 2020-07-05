### Regenerate grpc

```bash
protoc -I grpc --go_out=paths=source_relative:grpc --go-grpc_out=paths=source_relative:grpc **/*.proto
```
