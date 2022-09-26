package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"

	"github.com/sajanjswl/sandbox-service/config"
	v1 "github.com/sajanjswl/sandbox-service/gen/go/sandbox/v1alpha1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, authServiceServerApi v1.AuthServiceServer, cfg *config.Config, logger *zap.Logger) error {

	listen, err := net.Listen(cfg.GRPCNetworkType, ":"+cfg.GRPCPort)
	logger.Info("gRPC auth sevice would listen on", zap.String("network-type", cfg.GRPCNetworkType), zap.String("port", cfg.GRPCPort))

	if err != nil {
		logger.Error("gRPC auth service failed to listen", zap.Error(err), zap.String("network-type", cfg.GRPCNetworkType), zap.String("port", cfg.GRPCPort))
		return err
	}

	// register authServiceServerApi with grpc-server
	server := grpc.NewServer()
	v1.RegisterAuthServiceServer(server, authServiceServerApi)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Warn("gRPC auth service shutting down ...!!")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Info("stating gRPC server...")
	return server.Serve(listen)
}
