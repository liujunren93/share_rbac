package test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	re "github.com/go-redis/redis/v8"
	"github.com/liujunren93/share_utils/databases/redis"
)

var r *re.Client

func init() {

	r = redis.NewRedis(&re.Options{
		Network: "tcp",
		Addr:    "node1:6379",
	})

	if sc := r.Ping(context.Background()); sc.Err() != nil {
		panic(sc.Err().Error())
	}
}

var a []int
var mu = sync.RWMutex{}

func TestList(t *testing.T) {
	go func() {
		for i := 0; i < 10; i++ {
			a = append(a, 1)
		}

	}()
	go func() {
		for i := 0; i < 10; i++ {
			a = append(a, 2)
		}

	}()
	time.Sleep(time.Second * 1)
	fmt.Println(a)
}

func ttt(c <-chan int) {
	for {
		select {
		case m := <-c:
			fmt.Println(111, m)
		}
		time.Sleep(time.Second)
	}
}
