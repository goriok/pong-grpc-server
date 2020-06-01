server:
	go run server/server.go

client: 
	go run client/client.go

create-proto:
	protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative proto/pong.proto

