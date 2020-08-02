package mediator

type Behavior interface {
  Execute(h IMediatorCommandHandler, m interface{}, r interface{}) error
  SetNext(n Behavior)
  HasNext() bool
}
