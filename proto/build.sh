 protoc -I.  --go_out=plugins=grpc:. *.proto

 protoc-go-inject-tag -input="../pb/*.pb.go"