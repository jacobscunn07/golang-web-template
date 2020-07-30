package mediator

import (
  "reflect"
)

// Mediator
func NewMediator() *Mediator {
  m := &Mediator{Handlers: make(map[string]IMediatorCommandHandler)}
  //m.Register(new(SayHelloCommand))
  return m
}

type Mediator struct {
  Handlers map[string]IMediatorCommandHandler
  Validators map[string]IMediatorCommandValidator
}

type IMediator interface {
  Send(c interface{}, ret interface{})
  Register(t reflect.Type, c IMediatorCommandHandler)
}

func (m *Mediator) Register(t reflect.Type, c IMediatorCommandHandler) {
  m.Handlers[t.Name()] = c
}

func (m *Mediator) Send(c interface{}, ret interface{}) {
  m.Handlers[reflect.TypeOf(c).Name()].Handle(c, ret)
  //log.Println(ret)
}

// Mediator Command Handler
type IMediatorCommandHandler interface {
  Handle(m interface{}, ret interface{})
}

type IMediatorCommandValidator interface {
  Validate(m interface{}) bool
}
