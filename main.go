package main

import (
	"context"
	"flag"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sajanjswl/sandbox-service/config"
	"github.com/sajanjswl/sandbox-service/models"
	"github.com/sajanjswl/sandbox-service/pkg/protocol/grpc"

	"github.com/sajanjswl/sandbox-service/pkg/protocol/rest"
	"github.com/sajanjswl/sandbox-service/pkg/service/v1alpha1"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const service string = "sandbox-service"

func main() {
	cfg := config.NewConfig()

	// grpc-server configs
	flag.StringVar(&cfg.GRPCHost, "grpc-host", "localhost", "grpc host")
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "8000", "grpc port")
	flag.StringVar(&cfg.GRPCNetworkType, "grpc-network-type", "tcp", "grpc network type")

	// rest-server configs
	flag.StringVar(&cfg.RESTHost, "rest-host", "localhost", "rest host")
	flag.StringVar(&cfg.RESTPort, "rest-port", "9000", "rest port")
	flag.StringVar(&cfg.AbsoluteLogPath, "logpath", "logs.log", "application logs path")

	// database configs
	flag.StringVar(&cfg.DBHost, "db-host", "localhost", "database host")
	flag.StringVar(&cfg.DBPort, "db-port", "3305", "database port")
	flag.StringVar(&cfg.DBUserName, "db-user-name", "root", "db username")
	flag.StringVar(&cfg.DBPassword, "db-password", "example", "db password")
	flag.StringVar(&cfg.DBName, "db-name", "user", "database name")

	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	db := InitDb(logger, cfg)
	db.AutoMigrate(&models.User{})

	ctx := context.Background()

	grpcAuthServerApi := v1alpha1.NewAuthServiceServer(db, logger, cfg)
	//passing DB connection to Rest
	restAuthServerApi := rest.NewRestServer(db, cfg, logger)

	// // run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, restAuthServerApi)
	}()
	grpc.RunServer(ctx, grpcAuthServerApi, cfg, logger)

}

func InitDb(logger *zap.Logger, cfg *config.Config) *gorm.DB {
	var err error
	dns := cfg.DBUserName + ":" + cfg.DBPassword + "@tcp" + "(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?" + "charset=utf8mb4&parseTime=True&loc=Local"
	logger.Info("dns config", zap.String("dns", dns))
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		logger.Fatal("error connecting to database", zap.Error(err))
		return nil
	}

	return db
}
