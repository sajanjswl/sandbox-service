package main

import (
	"flag"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sajanjswl/sandbox-service/config"
	"go.uber.org/zap"
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

	// db := database.InitDb(logger, cfg)
	// db.AutoMigrate(&models.User{})

	// ctx := context.Background()
	// grpcAuthServerApi := grpcv1.NewAuthServiceServer(db, logger, cfg)
	// // //passing DB connection to Rest
	// restAuthServerApi := restv1.NewRestServer(db, cfg, logger)

	// // // // run HTTP gateway
	// go func() {
	// 	_ = rest.RunServer(ctx, restAuthServerApi, cfg, logger)
	// }()
	// grpc.RunServer(ctx, grpcAuthServerApi, cfg, logger)

}
