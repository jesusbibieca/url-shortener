package environment

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	AppPort    string `mapstructure:"APP_PORT"`
	AppAddress string `mapstructure:"APP_ADDRESS"`

	//redis
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisDb       int    `mapstructure:"REDIS_DB"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisAddress  string `mapstructure:"REDIS_ADDRESS"`

	//postgres
	DbDriver string `mapstructure:"DB_DRIVER"`
	DbSource string `mapstructure:"DB_SOURCE"`
}

var Config config

func LoadConfig(path string) (*config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Error while reading config file %s", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &Config, nil
}
