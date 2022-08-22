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

func (t *TOwner) Execute(func()) {
	fmt.Println("Executed!")
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
	tn.o.AddCallBack(&timerassistant.CallInfo{
		Category: &timerassistant.Once{
			Hour: 0,
			Min:  0,
			Sec:  0,
		},
		Fn: func() {
			fmt.Println("fn inner")
		},
		ResumeCallCh: tn.resumeCh,
	})
	go tn.loop()
	go tn.o.Loop()
	select {}
}
