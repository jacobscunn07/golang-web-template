package mediator

type Behavior interface {
  Execute(func() error) error
  SetNext(n Behavior)
  HasNext() bool
}
