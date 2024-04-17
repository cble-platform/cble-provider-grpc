package cble

import (
	"context"
	"fmt"
	"net"
	sync "sync"

	"github.com/cble-platform/cble-provider-grpc/pkg/common"
	"github.com/sirupsen/logrus"
	"golang.org/x/mod/semver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type DefaultCBLEServer struct {
	UnimplementedCBLEServer
}

type CBLEServerOptions struct {
	TLS      bool
	CertFile string
	KeyFile  string
	Socket   string
}

var defaultServerOptions = &CBLEServerOptions{
	TLS:      false,
	CertFile: "",
	KeyFile:  "",
	Socket:   "/tmp/cble-server",
}

func DefaultServe(ctx context.Context, server CBLEServer) error {
	return Serve(ctx, server, defaultServerOptions)
}

// Serve is a blocking call which returns an error if unable to serve
func Serve(ctx context.Context, server CBLEServer, options *CBLEServerOptions) error {
	lis, err := net.Listen("unix", options.Socket)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if options.TLS {
		if options.CertFile == "" || options.KeyFile == "" {
			return fmt.Errorf("must provider a certificate and key file if using TLS")
		}
		creds, err := credentials.NewServerTLSFromFile(options.CertFile, options.KeyFile)
		if err != nil {
			return err
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	RegisterCBLEServer(grpcServer, server)

	// Setup graceful shutdown signals
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		<-ctx.Done()
		grpcServer.GracefulStop()
		wg.Done()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	wg.Wait()
	return nil
}

func (s DefaultCBLEServer) Handshake(ctx context.Context, request *common.HandshakeRequest) (*common.HandshakeReply, error) {
	if semver.Major(request.ClientVersion) != semver.Major(VERSION) {
		return nil, fmt.Errorf("major version mismatch: server version is %s and client version is %s", VERSION, request.ClientVersion)
	}
	logrus.Debugf("Client (v%s) connected", request.ClientVersion)
	return &common.HandshakeReply{
		ServerVersion: VERSION,
	}, nil
}
