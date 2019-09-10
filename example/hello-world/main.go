package main
import (
	"context"
	"log"
	"net"

	health "github.com/yourrepo/hello-world/server/health"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	//TODO: Register your servers here
	healthServer := health.NewServer()
	health.RegisterHealthServer(s, healthServer)
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
