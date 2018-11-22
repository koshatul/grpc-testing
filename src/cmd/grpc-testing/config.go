package main

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func configDefaults() {
	viper.SetDefault("grpc.port", 50061)
	viper.SetDefault("grpc.host", "localhost")
	viper.SetDefault("grpc.client-cert-file", "artifacts/cfssl/localhost.pem")
	viper.SetDefault("grpc.server-cert-file", "artifacts/cfssl/localhost.pem")
	viper.SetDefault("grpc.server-key-file", "artifacts/cfssl/localhost-key.pem")
}

func configInit() {
	viper.SetConfigName("grpc-testing")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./artifacts")
	viper.AddConfigPath("./test")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath("/etc")
	viper.AddConfigPath("/run/secrets")
	viper.AddConfigPath(".")

	configDefaults()

	viper.ReadInConfig()

	configFormatting()
}

func configFormatting() {
}

func zapConfig() zap.Config {
	var cfg zap.Config
	if viper.GetBool("debug") {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}
	return cfg
}
