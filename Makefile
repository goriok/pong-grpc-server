protobuff:
	protoc --proto_path=src --go_out=build/gen --go_opt=paths=source_relative src/main/pong.proto

start:
	go run src/main/server.go
