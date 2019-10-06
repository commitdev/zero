package helloworld

import (
	"context"
	health_api "github.com/yourrepo/hello-world-idl/gen/go/health"
	//api "github.com/yourrepo/hello-world-idl/gen/go/helloworld"
)

type HelloworldServer struct {

}

func NewHelloworldServer() *HelloworldServer {
	return &HelloworldServer{}
}

func (s *HelloworldServer) Check(ctx context.Context, req *health_api.HealthCheckRequest) (*health_api.HealthCheckResponse, error) {
	resp := &health_api.HealthCheckResponse{
		Status: health_api.HealthCheckResponse_SERVING,
	}
	return resp,nil
}
