package behavior

type Behavior interface {
  Execute() error
}

type BehaviorPipeline struct {
  behaviors []*Behavior
}

func (b *BehaviorPipeline) SendHTTP(f func() error) error {
  return f()
}
