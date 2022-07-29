generate:
	@protoc --proto_path=proto --go_out=proto/gen --go_opt=paths=source_relative proto/upvote.proto