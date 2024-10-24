package configuration

import "github.com/Netflix/go-env"

type ServicesConfig struct {
	AuthServiceSocket  string `env:"AUTH_SERVICE_SOCKET"`
	TeamsServiceSocket string `env:"TEAMS_SERVICE_SOCKET"`
}

type DBConfig struct {
	Host string `env:"DB_HOST"`
	User string `env:"DB_USER"`
	Pass string `env:"DB_PASS"`
	Name string `env:"DB_NAME"`
}

type ServerConfig struct {
	HttpSocket string `env:"HTTP_SOCKET"`
	GrpcSocket string `env:"GRPC_SOCKET"`
}

type S3Config struct {
	Endpoint        string `env:"S3_ENDPOINT"`
	AccessKeyID     string `env:"S3_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"S3_SECRET_ACCESS_KEY"`
	BucketName      string `env:"S3_BUCKET_NAME"`
	UseSSL          bool   `env:"S3_USE_SSL"`
}

type Config struct {
	DB       DBConfig
	Server   ServerConfig
	S3       S3Config
	Services ServicesConfig
}

func Load() (*Config, error) {
	var cfg Config
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
