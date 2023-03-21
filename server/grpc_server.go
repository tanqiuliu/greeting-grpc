package server

import (
	"crypto/tls"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/tanqiuliu/greeting-grpc/api/greetingpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"math"
	"net"
)

const (
	maxRecvBytes = 1 * 1024 * 1024
	maxSendBytes = math.MaxInt32
)

func grpcServer(tlsConfig *tls.Config) *grpc.Server {
	var opts []grpc.ServerOption

	// set up tls if configured
	if tlsConfig != nil {
		opts = append(opts, grpc.Creds(credentials.NewTLS(tlsConfig)))
	}

	chainUnaryInterceptors := []grpc.UnaryServerInterceptor{
		grpc_prometheus.UnaryServerInterceptor,
	}

	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(chainUnaryInterceptors...)))
	opts = append(opts, grpc.MaxRecvMsgSize(maxRecvBytes))
	opts = append(opts, grpc.MaxSendMsgSize(maxSendBytes))

	grpcServer := grpc.NewServer(opts...)

	greetingpb.RegisterGreetingServer(grpcServer, NewGreetingServerImpl())

	grpc_prometheus.Register(grpcServer)
	return grpcServer
}

func Start(port int) error {
	// start grpc server
	var gs *grpc.Server
	defer func() {
		if gs != nil {
			gs.GracefulStop()
			fmt.Println("gRPC server gracefully stopped.")
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	gs = grpcServer(nil)
	fmt.Println("Starting gRPC server...")
	err = gs.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start serving: %v", err)
		return err
	}
	return nil
}
