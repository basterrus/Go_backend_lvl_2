package config

type Config struct {
	DatabaseURL string `env:"DATABASE_URL"`
	RESTPort    int    `env:"REST_SERVER_PORT" envDefault:"3000"`
	GRPCPort    int    `env:"GRPC_SERVER_PORT" envDefault:"50051"`
}
