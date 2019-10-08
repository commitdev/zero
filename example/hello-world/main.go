package main
import (
	"log"
	"net"

	healthpb "github.com/yourrepo/hello-world-idl/gen/go/health"
	
	health "github.com/yourrepo/hello-world/server/health"

	"google.golang.org/grpc"
)

func main() {
	grpcAddr := "0.0.0.0:3000"
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	//Server initialization & registration
	healthServer := health.NewHealthServer()
	healthpb.RegisterHealthServer(s, healthServer)


	log.Printf("Starting grpc server on %v...", grpcAddr)
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
