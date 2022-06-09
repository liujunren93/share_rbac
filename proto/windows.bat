 protoc -I. -ID:\1mywork\protobuf\src\  --go_out=plugins=grpc:. *.proto

 protoc-go-inject-tag -input="../pb/*.pb.go"
