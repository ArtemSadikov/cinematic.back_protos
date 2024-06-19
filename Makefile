gen_proto:
	protoc -I proto ./proto/**/*.proto \
	--go_out=./generated/go				--go_opt=paths=source_relative \
	--go-grpc_out=./generated/go	--go-grpc_opt=paths=source_relative
