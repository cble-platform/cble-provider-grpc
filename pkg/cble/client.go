package cble

import (
	"context"
	"fmt"

	common "github.com/cble-platform/cble-provider-grpc/pkg/common"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type CBLEClientOptions struct {
	TLS    bool
	CAFile string
	Port   int
}

var defaultClientOptions = &CBLEClientOptions{
	TLS:    false,
	CAFile: "",
	Port:   50051,
}

func DefaultConnect() (*grpc.ClientConn, error) {
	return Connect(defaultClientOptions)
}

// Connect returns a CBLEServer gRPC connection to the CBLE gRPC server for use with the gRPC client
func Connect(options *CBLEClientOptions) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if options.TLS {
		if options.CAFile == "" {
			return nil, fmt.Errorf("CA file must be provided for TLS")
		}
		creds, err := credentials.NewClientTLSFromFile(options.CAFile, "")
		if err != nil {
			return nil, fmt.Errorf("failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", options.Port), opts...)
	if err != nil {
		return nil, fmt.Errorf("fail to dial: %v", err)
	}

	return conn, nil
}

func NewClient(ctx context.Context, conn grpc.ClientConnInterface) (CBLEServerClient, error) {
	client := NewCBLEServerClient(conn)
	reply, err := client.Handshake(ctx, &common.HandshakeRequest{
		ClientVersion: VERSION,
	})
	if err != nil {
		return client, fmt.Errorf("handshake failed: %v", err)
	}
	logrus.Debugf("Connected to CBLE Server (v%s)", reply.ServerVersion)
	return client, nil
}
