package config

type Settings struct {
	PostgresURI string `envconfig:"POSTGRES_URI"`
}
