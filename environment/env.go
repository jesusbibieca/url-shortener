package environment

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Configuration struct {
	AppPort    string `mapstructure:"APP_PORT"`
	AppAddress string `mapstructure:"APP_ADDRESS"`

	// REDIS
	RedisPort    string `mapstructure:"REDIS_PORT"`
	RedisDb      int    `mapstructure:"REDIS_DB"`
	RedisAddress string `mapstructure:"REDIS_ADDRESS"`

	// POSTGRES
	DbDriver string `mapstructure:"DB_DRIVER"`
	DbSource string `mapstructure:"DB_SOURCE"`

	// AUTHENTICATION
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

var Config Configuration

func LoadConfig(path string) (*Configuration, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Println("Error reading config file: ", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &Config, nil
}
