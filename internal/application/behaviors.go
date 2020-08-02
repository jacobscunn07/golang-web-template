package application

import (
  "fmt"
  "github.com/jacobscunn07/golang-web-template/pkg/mediator"
  "github.com/sirupsen/logrus"
)

type LoggingBehavior struct {
  next mediator.Behavior
  logger *logrus.Logger
}

func NewLoggingBehavior(logger *logrus.Logger) *LoggingBehavior {
  return &LoggingBehavior{logger: logger}
}

func (b *LoggingBehavior) SetNext(n mediator.Behavior) {
  b.next = n
}

func (b *LoggingBehavior) Execute(h mediator.IMediatorCommandHandler, m interface{}, r interface{}) error {
  var err error = nil
  if !b.HasNext() {
    err = h.Handle(m, r)
  } else {
    err = b.next.Execute(h, m, r)
  }

  contextLogger := b.logger.
    WithField("command_or_query_type", fmt.Sprintf("%T", m)).
    WithField("command_or_query", m).
    WithField("result_type", fmt.Sprintf("%T", r)).
    WithField("result", r)

  if err == nil {
    contextLogger.Debug("command or query handled successfully")
  } else {
    contextLogger.Error("an error occurred: %v", err)
  }

  return err
}

func (b *LoggingBehavior) HasNext() bool {
  return b.next != nil
}
