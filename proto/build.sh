 protoc -I.  --go_out=:.  --go-grpc_out=require_unimplemented_servers=false:. *.proto

 protoc-go-inject-tag -input="../pb/*.pb.go"