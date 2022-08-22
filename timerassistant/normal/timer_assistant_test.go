package normal

import (
	"fmt"
	"github.com/lqgl/hub/timerassistant"
	"testing"
	"time"
)

type TOwner struct {
	resumeCh chan func()
	o        timerassistant.TimerAssistant
}

// Execute callback func
func (t *TOwner) Execute(fn func()) {
	fmt.Println("Executed!")
	fn()
}

func (t TOwner) loop() {
	for {
		select {
		case fn := <-t.resumeCh:
			t.Execute(fn)
		}
	}
}

func TestTimerAssistant(t *testing.T) {
	tn := &TOwner{
		resumeCh: make(chan func(), 1),
		o:        NewTimerAssistant(time.Second),
	}
	onceCall := &timerassistant.CallInfo{
		Category: &timerassistant.Once{
			Hour: 11,
			Min:  45,
			Sec:  0,
		},
		ResumeCallCh: tn.resumeCh,
	}
	onceCall.Fn = func() {
		fmt.Println("inner fn execute!")
		tn.o.DelCallBack(onceCall) // because in this case only run once
	}
	tn.o.AddCallBack(onceCall)
	go tn.loop()
	go tn.o.Loop()
	select {}
}
