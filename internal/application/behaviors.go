package application

import (
  "github.com/jacobscunn07/golang-web-template/pkg/mediator"
  "log"
)

type LoggingBehavior struct {
  next mediator.Behavior
}

func NewLoggingBehavior() *LoggingBehavior {
  return &LoggingBehavior{}
}

func (b *LoggingBehavior) SetNext(n mediator.Behavior) {
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

func (b *LoggingBehavior) HasNext() bool {
  return b.next != nil
}

type TimerBehavior struct {
  next mediator.Behavior
}

func NewTimerBehavior() *TimerBehavior {
  return &TimerBehavior{}
}

func (b *TimerBehavior) SetNext(n mediator.Behavior) {
  b.next = n
}

func (b *TimerBehavior) Execute(f func() error) error {
  log.Println("Entering Timer Behavior")
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

  log.Println("Exiting Timer Behavior")
  return r
}

func (b *TimerBehavior) HasNext() bool {
  return b.next != nil
}
