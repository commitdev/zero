package main

import (
  "log"
  "context"
  "net/http"

  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"

	health "github.com/yourrepo/hello-world-idl/gen/go/health"
	helloworld "github.com/yourrepo/hello-world-idl/gen/go/helloworld"
)

func run(endpoint string, listening string) error {

  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := health.RegisterHealthHandlerFromEndpoint(ctx, mux,  endpoint, opts)
	err = helloworld.RegisterHelloworldHandlerFromEndpoint(ctx, mux,  endpoint, opts)

  if err != nil {
    return err
  }

  return http.ListenAndServe(listening, mux)
}

func main() {
  endpoint := "0.0.0.0:3000"
  listening := "0.0.0.0:8080"
  log.Printf("Starting http grpc gateway server on %v...", listening)

  if err := run(endpoint, listening); err != nil {
    log.Fatal(err)
  }
}
