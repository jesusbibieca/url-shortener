package environment

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type config struct {
	AppPort string `mapstructure:"APP_PORT"`
	AppAddress string `mapstructure:"APP_ADDRESS"`

	//redis
	RedisPort string `mapstructure:"REDIS_PORT"`
	RedisDb int `mapstructure:"REDIS_DB"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisAddress string `mapstructure:"REDIS_ADDRESS"`
}

var Config config

func LoadConfig() (*config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	
  // Find and read the config file
  err := viper.ReadInConfig()
	
  if err != nil {
		  log.Fatalf("Error while reading config file %s", err)
		}

	fmt.Println(viper.ConfigFileUsed())

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}

	fmt.Println(Config)
return &Config, nil
}
