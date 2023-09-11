
## Cоздаем **brotocol buffer** для сервера
- protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto
# go-grpc
