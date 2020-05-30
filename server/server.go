package main

import (
  "net"
  "log"
  "google.golang.org/grpc"
  pb "github.com/goriok/poc-grpc-pong/proto"
)

const (
  port = ":2000"
)

type server struct{}

func (s *server) Ping(ctx context.Context, in *pb.Ping) (*pb.Pong, error) {
    return &pb.Pong{Response: "PONG"}, nil
}

func main(){
  listener, err := net.Listen("tcp", port)

  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  s := grpc.NewServer()
  pb.RegisterPingServiceServer(s, &server{})
  server.Serve(listener)
}
