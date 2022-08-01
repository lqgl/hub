package mongohub

import (
	"context"
	"fmt"
	"github.com/lqgl/hub"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

type testOwner struct {
	c *Client
}

func (t testOwner) Launch() {
	t.c.Launch()
}

func (t testOwner) Stop() {
	t.c.Stop()
}

func TestClient(t *testing.T) {
	ctx := context.Background()
	to := &testOwner{}
	tc := &Client{
		BaseComponent: hub.NewBaseComponent(),
		RealCli: NewClient(ctx, &Config{
			URI:         "mongodb://root:mongo@localhost:27017",
			MinPoolSize: 3,
			MaxPoolSize: 3000,
		}),
	}
	to.c = tc
	to.Launch()
	fn := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		res, err := tc.InsertOne(ctx, "strike", "movies",
			bson.D{{"title", "Moon Light"}, {"title", "Tiller"}})
		if err != nil {
			fmt.Println(err)
		}
		id := res.InsertedID
		fmt.Println(id)
	}
	op := hub.Operation{
		CB:  fn,
		Ret: make(chan interface{}),
	}
	to.c.Resolve(op)
	<-op.Ret
	fmt.Println("op success")
	time.Sleep(time.Second * 5)
	to.Stop()
}
