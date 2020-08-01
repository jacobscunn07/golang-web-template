package web

import (
  "fmt"
  "github.com/go-pg/pg/v10"
  _ "github.com/golang-migrate/migrate/v4/database/postgres"
  _ "github.com/golang-migrate/migrate/v4/source/file"
  "github.com/jacobscunn07/golang-web-template/configs"
  "github.com/jacobscunn07/golang-web-template/internal/application"
  "github.com/jacobscunn07/golang-web-template/internal/data"
  "github.com/jacobscunn07/golang-web-template/pkg/mediator"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "net/http"
  "reflect"
)

func Run() {
  configuration := configs.GetConfiguration()

  data.Migrate(configuration.ConnectionString.ToString())

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
  //e.Use(middleware.Logger())
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
