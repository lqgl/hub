package example

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

type ConsumerHandle struct {
	handlers map[uint64]func(message *nsq.Message) error
}

// Register fn callback
func (c *ConsumerHandle) Register(id uint64, fn func(message *nsq.Message) error) {
	c.handlers = make(map[uint64]func(message *nsq.Message) error)
	c.handlers[id] = fn
}

// HandleMessage Message handler
func (c *ConsumerHandle) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}
	fmt.Println(string(m.Body)) // just print message body
	err := c.processMessage(m)
	return err
}

func (c *ConsumerHandle) processMessage(m *nsq.Message) error {
	f, ok := c.handlers[0] // 调用第一个 handler
	if ok {
		return f(m)
	}
	return nil
}
