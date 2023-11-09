package cble

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	sync "sync"
	"syscall"

	"github.com/cble-platform/cble-provider-grpc/pkg/common"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/mod/semver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type DefaultCBLEServer struct {
	UnimplementedCBLEServer

	// Maps provider name@version to port it's running on
	RegisteredProviders map[string]RegisteredProvider
}

type RegisteredProvider struct {
	ID       string
	SocketID string
	Features map[string]bool
}

type CBLEServerOptions struct {
	TLS      bool
	CertFile string
	KeyFile  string
	Port     int
}

var defaultServerOptions = &CBLEServerOptions{
	TLS:      false,
	CertFile: "",
	KeyFile:  "",
	Port:     50051,
}

func DefaultServe(server CBLEServer) error {
	return Serve(server, defaultServerOptions)
}

// Serve is a blocking call which returns an error if unable to serve
func Serve(server CBLEServer, options *CBLEServerOptions) error {
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
	RegisterCBLEServer(grpcServer, server)

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

func (s DefaultCBLEServer) Handshake(ctx context.Context, request *common.HandshakeRequest) (*common.HandshakeReply, error) {
	if semver.Major(request.ClientVersion) != semver.Major(VERSION) {
		return nil, fmt.Errorf("major version mismatch: server version is %s and client version is %s", VERSION, request.ClientVersion)
	}
	logrus.Debugf("Client (v%s) connected", request.ClientVersion)
	return &common.HandshakeReply{
		ServerVersion: VERSION,
	}, nil
}

func (s DefaultCBLEServer) RegisterProvider(ctx context.Context, request *RegistrationRequest) (*RegistrationReply, error) {
	logrus.Debugf("Registration request from %s@%s (%s)", request.Name, request.Version, request.Id)
	providerKey := fmt.Sprintf("%s@%s", request.Name, request.Version)
	// Check if a provider with this version is already registered
	if _, exist := s.RegisteredProviders[providerKey]; exist {
		return nil, fmt.Errorf("provider with same name and version (%s) already registered", providerKey)
	}
	// Generate random UUID for socket
	socketId := uuid.NewString()
	// Map the port
	s.RegisteredProviders[providerKey] = RegisteredProvider{
		ID:       request.Id,
		SocketID: socketId,
		Features: request.Features,
	}
	// Reply to the provider
	return &RegistrationReply{
		Status:   common.RPCStatus_SUCCESS,
		SocketId: socketId,
	}, nil
}

func (s DefaultCBLEServer) UnregisterProvider(ctx context.Context, request *UnregistrationRequest) (*UnregistrationReply, error) {
	logrus.Debugf("Unregistration request from %s@%s (%s)", request.Name, request.Version, request.Id)
	providerKey := fmt.Sprintf("%s@%s", request.Name, request.Version)
	// Check to make sure this provider is registered
	prov, exists := s.RegisteredProviders[providerKey]
	if !exists {
		return &UnregistrationReply{
			Status: common.RPCStatus_FAILURE,
		}, nil
	}
	// Make sure the unregister request is coming with the right ID... super basic security check :)
	if prov.ID != request.Id {
		return &UnregistrationReply{
			Status: common.RPCStatus_FAILURE,
		}, nil
	}
	// If all that passes, unregister the provider
	delete(s.RegisteredProviders, providerKey)
	return &UnregistrationReply{
		Status: common.RPCStatus_SUCCESS,
	}, nil
}
