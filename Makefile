PROTO_DIR = pkg/pb

PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
proto:
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_opt=require_unimplemented_servers=false \
	--go-grpc_out=. ${PROTO_DIR}/*.proto

server:
	go run cmd/main.go