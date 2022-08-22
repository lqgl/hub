package normal

import (
	"fmt"
	"github.com/lqgl/hub/timerassistant"
	"sync"
	"time"
)

type TimerAssistant struct {
	tickTime  time.Duration
	callbacks sync.Map
}

func NewTimerAssistant(tickTime time.Duration) *TimerAssistant {
	return &TimerAssistant{
		tickTime: tickTime,
	}
}

func (t *TimerAssistant) AddCallBack(info *timerassistant.CallInfo) {
	t.callbacks.Store(info, struct{}{})
}

func (t *TimerAssistant) DelCallBack(info *timerassistant.CallInfo) {
	t.callbacks.Delete(info)
}

func (t *TimerAssistant) Loop() {
	go func() {
		tick := time.NewTicker(t.tickTime)
		defer func() {
			tick.Stop()
			err := recover()
			if err != nil {
				fmt.Println(err)
			}
		}()
		for {
			select {
			case <-tick.C:
				t.Process()
			}
		}
	}()
}

func (t *TimerAssistant) AssertOwner() {
	//TODO implement me
	panic("implement me")
}

// Process callback
func (t *TimerAssistant) Process() {
	t.callbacks.Range(func(key, value any) bool {
		cbInfo := key.(*timerassistant.CallInfo)
		cbInfo.NotifyOwner()
		return true
	})
}
