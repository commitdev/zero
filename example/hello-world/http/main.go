package main

import (
  "os"
  "fmt"
  "log"
  "context"
  "net/http"

  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"

	health "github.com/yourrepo/hello-world-idl/gen/go/health"
)

func run(endpoint string, listening string) error {

  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := health.RegisterHealthHandlerFromEndpoint(ctx, mux,  endpoint, opts)

  if err != nil {
    return err
  }

  return http.ListenAndServe(listening, mux)
}

func main() {

  endpoint := fmt.Sprintf("%s:%s", getEnv("APP_HOST", "0.0.0.0"), getEnv("APP_PORT", "3000"))
  listening := fmt.Sprintf("%s:%s", getEnv("HTTP_HOST", "0.0.0.0"), getEnv("HTTP_PORT", "8080"))

  log.Printf("Starting http grpc gateway server on %v...", listening)

  if err := run(endpoint, listening); err != nil {
    log.Fatal(err)
  }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
