package config

type Config struct {
	//  gRPC
	GRPCNetworkType string
	GRPCHost        string
	GRPCPort        string

	// REST Gateway
	RESTHost        string
	RESTPort        string
	AbsoluteLogPath string
	// database
	DBHost     string
	DBName     string
	DBUserName string
	DBPassword string
	DBPort     string
	DBSLLMode  string
}

func NewConfig() *Config {
	return &Config{}
}
