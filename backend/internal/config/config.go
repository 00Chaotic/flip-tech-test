package config

// Config defines the configuration object for the server.
type Config struct {
	DatabaseUrl string `envconfig:"DB_URL"`
	ServerPort  int    `envconfig:"SERVER_PORT" default:"8080"`
}
