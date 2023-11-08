# cble-provider-grpc

This package is designed to be a communication method for the CBLE server to talk to providers.

## Minimal Provider Implementation

```go
package main

import (
  "context"

  cbleGRPC "github.com/cble-platform/cble-provider-grpc/pkg/cble"
  commonGRPC "github.com/cble-platform/cble-provider-grpc/pkg/common"
  providerGRPC "github.com/cble-platform/cble-provider-grpc/pkg/provider"
  "github.com/google/uuid"
  "github.com/sirupsen/logrus"
)

var (
  id      = uuid.New().String()
  name    = "example-provider"
  version = "v0.1"
)

type ExampleProvider struct {
  providerGRPC.DefaultProviderServer
}

func (ExampleProvider) Deploy(ctx context.Context, request *providerGRPC.DeployRequest) (*providerGRPC.DeployReply, error) {
  logrus.Infof("Deploying something with example provider")
  return &providerGRPC.DeployReply{
    DeploymentId:    request.DeploymentId,
    Status:          commonGRPC.RPCStatus_SUCCESS,
    DeploymentState: &structpb.Struct{},
    DeploymentVars:  &structpb.Struct{},
  }, nil
}

func main() {
  conn, err := cbleGRPC.DefaultConnect()
  if err != nil {
    logrus.Fatalf("failed to connect to CBLE gRPC server: %v", err)
  }
  defer conn.Close()

  ctx := context.Background()

  client, err := cbleGRPC.NewClient(ctx, conn)
  if err != nil {
    logrus.Fatalf("failed to connect client: %v", err)
  }

  registerReply, err := client.RegisterProvider(ctx, &cbleGRPC.RegistrationRequest{
    Id:      id,
    Name:    name,
    Version: version,
    Features: map[string]bool{
      string(providerGRPC.ProviderFeature_DEPLOY):  true,
      string(providerGRPC.ProviderFeature_DESTROY): false,
    },
  })
  if err != nil || registerReply.Status == commonGRPC.RPCStatus_FAILURE {
    logrus.Fatalf("registration failed: %v", err)
  } else if registerReply.Status == commonGRPC.RPCStatus_SUCCESS {
    logrus.Printf("Registration success! Starting provider server on port %d", registerReply.Port)
  } else {
    logrus.Fatalf("unknown error occurred: %v", err)
  }

  // Gracefully deregister on provider shutdown
  defer func() {
    // Time to shutdown
    unregisterReply, err := client.UnregisterProvider(ctx, &cbleGRPC.UnregistrationRequest{
      Id:      id,
      Name:    name,
      Version: version,
    })
    if err != nil || unregisterReply.Status == commonGRPC.RPCStatus_FAILURE {
      logrus.Fatalf("unregistration failed: %v", err)
    } else if unregisterReply.Status == commonGRPC.RPCStatus_SUCCESS {
      logrus.Print("Unregistration success! Shutting down...")
    } else {
      logrus.Fatalf("unknown error occurred: %v", err)
    }
  }()

  // Set up the provider gRPC server
  providerOpts := &providerGRPC.ProviderServerOptions{
    TLS:      false,
    CertFile: "",
    KeyFile:  "",
    Port:     int(registerReply.Port),
  }

  // Serve the provider gRPC server (blocking call)
  if err := providerGRPC.Serve(ExampleProvider{}, providerOpts); err != nil {
    logrus.Fatalf("failed to server provider gRPC server: %v", err)
  }

  // Provider is now ready to receive communications from CBLE
  //   (sending SIGINT/SIGTERM will shutdown the server)
}

```
