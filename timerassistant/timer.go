package timerassitant

type TimerAssitant interface {
	AddCallBack(*CallInfo)
	DelCallBack(*CallInfo)
	Loop()
	AssertOwner()
}
