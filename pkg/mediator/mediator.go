package mediator

import (
  "log"
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
  Behaviors []IBehavior
}

type IBehavior interface {
  Execute(func() error) error
  SetNext(n IBehavior)
}

type LoggingBehavior struct {
  next IBehavior
}

func (b *LoggingBehavior) SetNext(n IBehavior) {
  b.next = n
}

func (b *LoggingBehavior) Execute(f func() error) error {
  log.Println("Entering Logging Behavior")
  var r error
  if b.next == nil {
    log.Println("Executing command or query")
    if err := f(); err == nil {
      r = nil
    } else {
      r = err
    }
    log.Println("Finished executing command or query")
  } else {
    r = b.next.Execute(f)
  }

  log.Println("Exiting Logging Behavior")
  return r
}

type IMediator interface {
  Send(c interface{}, ret interface{})
  Register(t reflect.Type, c IMediatorCommandHandler)
}

func (m *Mediator) Register(t reflect.Type, c IMediatorCommandHandler) {
  m.Handlers[t.Name()] = c
}

func (m *Mediator) Send(c interface{}, ret interface{}) {

  b := new(LoggingBehavior)
  //m.Handlers[reflect.TypeOf(c).Name()].Handle(c, ret)
  f := func() error {
    return m.Handlers[reflect.TypeOf(c).Name()].Handle(c, ret)
  }
  if err := b.Execute(f); err != nil {
    log.Fatal(err)
  }
}

// Mediator Command Handler
type IMediatorCommandHandler interface {
  Handle(m interface{}, ret interface{}) error
}

type IMediatorCommandValidator interface {
  Validate(m interface{}) bool
}
