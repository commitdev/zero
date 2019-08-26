package health

import (
	"context"
	api "github.com/yourrepo/hello-world-go/health"
)

type HealthServer struct {

}

func (s *HealthServer) Check(ctx context.Context, req *api.HealthCheckRequest) (*api.HealthCheckResponse, error) {
	resp := &api.HealthCheckResponse{
		Status: api.HealthCheckResponse_SERVING,
	}
	return resp,nil
}

func (s *HealthServer) Watch(req *api.HealthCheckRequest, server api.Health_WatchServer) error {
	return nil
}