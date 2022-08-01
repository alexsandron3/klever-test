generate:
	@go mod download
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/upvote.proto

run:
	@echo "---- Running Server ----"
	@go run server/*.go