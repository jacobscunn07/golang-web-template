package mediator

import (
  "log"
  "reflect"
)

// Mediator
func NewMediator() *Mediator {
  return &Mediator{handlers: make(map[string]IMediatorCommandHandler)}
}

type Mediator struct {
  handlers map[string]IMediatorCommandHandler
  behavior Behavior
}

type IMediator interface {
  Send(c interface{}, ret interface{})
  Register(t reflect.Type, c IMediatorCommandHandler)
  Use(behaviors Behavior)
}

func (m *Mediator) Use(behavior Behavior) {
  if m.behavior == nil {
    m.behavior = behavior
  } else {
    findLast(m.behavior).SetNext(behavior)
  }
}

func findLast(b Behavior) Behavior {
  if b.HasNext() {
    return findLast(b)
  }

  return b
}

func (m *Mediator) Register(t reflect.Type, c IMediatorCommandHandler) {
  m.handlers[t.Name()] = c
}

func (m *Mediator) Send(c interface{}, ret interface{}) {
  f := func() error {
    return m.handlers[reflect.TypeOf(c).Name()].Handle(c, ret)
  }
  if err := m.behavior.Execute(f); err != nil {
    log.Fatal(err)
  }
}

// Mediator Command Handler
type IMediatorCommandHandler interface {
  Handle(m interface{}, ret interface{}) error
}
