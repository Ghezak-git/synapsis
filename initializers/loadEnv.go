package initializers

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"HOST"`
	DBUserName     string `mapstructure:"USER"`
	DBUserPassword string `mapstructure:"PASSWORD"`
	DBName         string `mapstructure:"DB"`
	DBPort         string `mapstructure:"PORTDB"`
	ServerPort     string `mapstructure:"PORT"`

	DBInitial string `mapstructure:"DB_INITIAL"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
	Domain       string `mapstructure:"DOMAIN"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
