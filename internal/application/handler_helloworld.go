package application

import (
  "fmt"
  "github.com/go-pg/pg/v10"
  //"github.com/jacobscunn07/golang-web-template/pkg/mediator"
  "log"
)

// Command
type SayHelloCommand struct {
  Message string
}

// Result
type SayHelloCommandResult struct {
  Result string
}

// Handler
func NewSayHelloCommandHandler(db *pg.DB) *SayHelloCommandHandler {
  return &SayHelloCommandHandler{db: db}
}

type SayHelloCommandHandler struct {
  db *pg.DB
}

func (c *SayHelloCommandHandler) Handle(m interface{}, ret interface{}) {
  origin := ret.(*SayHelloCommandResult)

  var n int
  _, err := c.db.QueryOne(pg.Scan(&n), "SELECT 1")
  if err != nil {
    log.Fatal(err)
  }

  *origin = SayHelloCommandResult{
    Result: fmt.Sprint(n),
  }

  ret = origin
}

// Validator
func NewSayHelloCommandValidator() *SayHelloCommandValidator {
 return &SayHelloCommandValidator{}
}

type SayHelloCommandValidator struct {}

func (c *SayHelloCommandValidator) Validate(m interface{}) bool {
  return true
}
