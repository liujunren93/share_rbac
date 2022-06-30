package test

import (
	"fmt"
	"net"
	"sync"
	"testing"
	"time"

	re "github.com/go-redis/redis/v8"
)

var r *re.Client

var a []int
var mu = sync.RWMutex{}

func TestList(t *testing.T) {

	conn, err := net.Listen("tcp", "0.0.0.0:0")
	fmt.Println(conn.Addr(), err)

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
