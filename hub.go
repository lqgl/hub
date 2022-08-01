package hub

type Component interface {
	Resolve(opCh Operation)
	Launch()
	Stop()
}

type BaseComponent struct {
	opCh   chan Operation
	stopCh chan struct{}
}

func NewBaseComponent() *BaseComponent {
	return &BaseComponent{
		opCh:   make(chan Operation),
		stopCh: make(chan struct{}),
	}
}

func (b *BaseComponent) Resolve(op Operation) {
	b.opCh <- op
}

func (b *BaseComponent) Launch() {
	go func() {
		for {
			select {
			case op := <-b.opCh:
				b.deal(op)
			case <-b.stopCh:
				break
			}
		}
	}()
}

func (b *BaseComponent) Stop() {
	b.stopCh <- struct{}{}
}

func (b *BaseComponent) deal(operation Operation) {
	fn := func() {
		operation.CB()
		operation.Ret <- struct{}{}
	}
	if !operation.IsAsynchronous { // 同步操作
		fn()
	} else { // 异步操作
		go func() {
			fn()
		}()
	}
}
