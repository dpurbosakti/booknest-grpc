package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ENVIRONMENT          string        `mapstructure:"ENVIRONMENT"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	EchoServerAddress    string        `mapstructure:"ECHO_SERVER_ADDRESS"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	//google configuration
	GoogleCientID        string `mapstructure:"GOOGLECLIENTID"`
	GoogleClientSecret   string `mapstructure:"GOOGLECLIENTSECRET"`
	GoogleRedirectURL    string `mapstructure:"GOOGLEREDIRECTURL"`
	GoogleState          string `mapstructure:"GOOGLESTATE"`
	GoogleTokenAccessURL string `mapstructure:"GOOGLETOKENACCESSURL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
