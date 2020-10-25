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
  "github.com/sirupsen/logrus"
  "net/http"
  "reflect"
)

func Run() {
  configuration := configs.GetConfiguration()

  logger := logrus.StandardLogger()
  logger.SetFormatter(&logrus.JSONFormatter{})
  logger.SetLevel(logrus.DebugLevel)

  data.Migrate(configuration.ConnectionString.ToString())

  db := pg.Connect(&pg.Options{
   Addr:     fmt.Sprintf("%v:%v", configuration.ConnectionString.Host, configuration.ConnectionString.Port),
   User:     configuration.ConnectionString.User,
   Password: configuration.ConnectionString.Password,
   Database: configuration.ConnectionString.User,
  })
  defer db.Close()

  m := mediator.NewMediator()
  m.Use(application.NewLoggingBehavior(logger))

  m.Register(reflect.TypeOf(application.SayHelloCommand{}), application.NewSayHelloCommandHandler(db))

  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
  })

  http.ListenAndServe(":5000", nil)

  //// Echo instance
  //e := echo.New()
  //
  //// Middleware
  ////e.Use(middleware.Logger())
  //e.Use(middleware.Recover())
  //
  //// Route => handler
  //e.GET("/", func(c echo.Context) error {
  //  var r application.SayHelloCommandResult
  //  m.Send(application.SayHelloCommand{Message: "hello"}, &r)
  //  return c.String(http.StatusOK, r.Result)
  //})
  //
  //// Start server
  //e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", configuration.HttpPort)))
}


// create database tables (embedded resources)
// create domain objects
// create repositories for aggregate roots
// create mediator command / queries
// implement validator and metric (prometheus) behaviors
