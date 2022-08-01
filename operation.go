package hub

type CallBackFn func()

type Operation struct {
	IsAsynchronous bool // 是否异步
	CB             CallBackFn
	Ret            chan interface{}
}
