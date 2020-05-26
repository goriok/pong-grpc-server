package main

import (
  "net"
  "log"
  "google.golang.org/grpc"
)

func main(){
  listener, err := net.Listen("tcp", "Localhost:2000")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  server := grpc.NewServer()
  server.Serve(listener)

}
