package provider

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	sync "sync"
	"syscall"

	common "github.com/cble-platform/cble-provider-grpc/pkg/common"
	"github.com/sirupsen/logrus"
	"golang.org/x/mod/semver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type DefaultProviderServer struct {
	UnimplementedProviderServer
}

type ProviderFeature string

const (
	ProviderFeature_DEPLOY  ProviderFeature = "DEPLOY"
	ProviderFeature_DESTROY ProviderFeature = "DESTORY"
)

type ProviderServerOptions struct {
	TLS      bool
	CertFile string
	KeyFile  string
	Port     int
}

// Serve is a blocking call which returns an error if unable to serve
func Serve(provider ProviderServer, options *ProviderServerOptions) error {
	if options == nil {
		return fmt.Errorf("options must not be nil")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", options.Port))
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
	RegisterProviderServer(grpcServer, provider)

	// Setup graceful shutdown signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-sigCh
		logrus.Warnf("Received signal %v, attempting graceful shutdown...", s)
		grpcServer.GracefulStop()
		wg.Done()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	wg.Wait()
	return nil
}

func (DefaultProviderServer) Handshake(ctx context.Context, request *common.HandshakeRequest) (*common.HandshakeReply, error) {
	if semver.Major(request.ClientVersion) != semver.Major(VERSION) {
		return nil, fmt.Errorf("major version mismatch: server version is %s and client version is %s", VERSION, request.ClientVersion)
	}
	logrus.Debugf("Client (v%s) connected", request.ClientVersion)
	return &common.HandshakeReply{
		ServerVersion: VERSION,
	}, nil
}
