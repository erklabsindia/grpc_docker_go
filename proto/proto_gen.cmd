protoc --go_out=./ --go-grpc_out=require_unimplemented_servers=false:./ ./posts.proto
protoc --go_out=./ --go-grpc_out=require_unimplemented_servers=false:./ ./news.proto