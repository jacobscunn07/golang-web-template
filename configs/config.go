package configs

import (
  "fmt"
  "github.com/spf13/viper"
)

type configuration struct {
  HttpPort         string                        `mapstructure:"http_port"`
  ConnectionString configurationConnectionString `mapstructure:"connection_string"`
}

func GetConfiguration() *configuration {
  viper.SetConfigName("development")
  viper.AddConfigPath("/app/configs")
  viper.SetEnvPrefix("todo")
  viper.AutomaticEnv()
  viper.SetConfigType("yaml")

  if err := viper.ReadInConfig(); err != nil {
    fmt.Printf("Error reading config file, %s", err)
  }

  var configuration configuration

  if err := viper.Unmarshal(&configuration); err != nil {
    fmt.Printf("Unable to decode into struct, %v", err)
  }

  return &configuration
}

type configurationConnectionString struct {
  User string `mapstructure:"user"`
  Password string `mapstructure:"password"`
  Host string `mapstructure:"host"`
  Port string `mapstructure:"port"`
  Database string `mapstructure:"database"`
}

func (cs *configurationConnectionString) ToString() string {
  return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
    cs.User,
    cs.Password,
    cs.Host,
    cs.Port,
    cs.Database)
}
