package config

type Config struct {
	ServerPort             int      `env:"SERVER_PORT" envDefault:"3000"`
	ElasticsearchAddresses []string `env:"ELASTICSEARCH_ADDRESSES" envSeparator:";"`
}
