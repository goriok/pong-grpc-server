start:
	go run server/server.go

create-proto:
	protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative proto/pong.proto

