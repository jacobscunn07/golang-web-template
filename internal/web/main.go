package web

import (
  "fmt"
  "github.com/go-pg/pg/v10"
  _ "github.com/golang-migrate/migrate/v4/database/postgres"
  _ "github.com/golang-migrate/migrate/v4/source/file"
  "github.com/jacobscunn07/golang-web-template/internal/application"
  "github.com/jacobscunn07/golang-web-template/internal/data"
  "github.com/jacobscunn07/golang-web-template/pkg/mediator"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/spf13/viper"
  "net/http"
  "reflect"
)

func Run() {
  viper.SetConfigName("development")
  viper.AddConfigPath("/app/configs")
  viper.SetEnvPrefix("todo")
  viper.AutomaticEnv()
  viper.SetConfigType("yaml")

  if err := viper.ReadInConfig(); err != nil {
    fmt.Printf("Error reading config file, %s", err)
  }

  var configuration Configuration

  if err := viper.Unmarshal(&configuration); err != nil {
    fmt.Printf("Unable to decode into struct, %v", err)
  }

  connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
    configuration.ConnectionString.User,
    configuration.ConnectionString.Password,
    configuration.ConnectionString.Host,
    configuration.ConnectionString.Port,
    configuration.ConnectionString.Database)

  data.Migrate(connectionString)

  db := pg.Connect(&pg.Options{
    Addr:     fmt.Sprintf("%v:%v", configuration.ConnectionString.Host, configuration.ConnectionString.Port),
    User:     configuration.ConnectionString.User,
    Password: configuration.ConnectionString.Password,
    Database: configuration.ConnectionString.User,
  })
  defer db.Close()

  m := mediator.NewMediator()
  m.Use(application.NewLoggingBehavior())
  m.Use(application.NewTimerBehavior())

  m.Register(reflect.TypeOf(application.SayHelloCommand{}), application.NewSayHelloCommandHandler(db))

  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Route => handler
  e.GET("/", func(c echo.Context) error {
    var r application.SayHelloCommandResult
    m.Send(application.SayHelloCommand{Message: "hello"}, &r)
    return c.String(http.StatusOK, r.Result)
  })

  // Start server
  e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", configuration.HttpPort)))
}

type Configuration struct {
  HttpPort string `mapstructure:"http_port"`
  ConnectionString ConfigurationConnectionString `mapstructure:"connection_string"`
}

type ConfigurationConnectionString struct {
  User string `mapstructure:"user"`
  Password string `mapstructure:"password"`
  Host string `mapstructure:"host"`
  Port string `mapstructure:"port"`
  Database string `mapstructure:"database"`
}
