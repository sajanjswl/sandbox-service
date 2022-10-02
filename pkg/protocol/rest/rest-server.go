package rest

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/sajanjswl/sandbox-service/config"
	v1 "github.com/sajanjswl/sandbox-service/gen/go/sandbox/v1alpha1"
	"gorm.io/gorm"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type RestServer struct {
	Db     *gorm.DB
	Mux    *http.ServeMux
	cfg    *config.Config
	logger *zap.Logger
}

func NewRestServer(db *gorm.DB, cfg *config.Config, logger *zap.Logger) *RestServer {
	return &RestServer{
		Db:     db,
		cfg:    cfg,
		logger: logger,
	}
}

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, restServer *RestServer) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := v1.RegisterSandboxServiceHandlerFromEndpoint(ctx, rmux, restServer.cfg.GRPCHost+":"+restServer.cfg.GRPCPort, opts); err != nil {

		restServer.logger.Error("failed to start gRPC-Rest gatewy", zap.Error(err))
		return err
	}

	restServer.Mux = http.NewServeMux()
	restServer.Mux.Handle("/", rmux)

	srv := &http.Server{
		Addr:    ":" + restServer.cfg.RESTPort,
		Handler: restServer.Mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			restServer.logger.Warn("shutting down gRPC-Rest Gateway....!!!")

		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	restServer.logger.Info("stating gRPC-Rest Gateway....!!!")

	return srv.ListenAndServe()

}
