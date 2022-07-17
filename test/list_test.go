package test

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)





func TestList(t *testing.T) {
	fmt.Println("start")
	ch := make(chan os.Signal, 100)
	sign := []os.Signal{
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,
	}
	signal.Notify(ch, sign...)

	for c := range ch {
		switch c {
		case syscall.SIGTERM:
			fmt.Println("SIGTERM")
		case syscall.SIGINT:
			fmt.Println("SIGINT")
		case syscall.SIGQUIT:
			fmt.Println("SIGQUIT")
		case syscall.SIGKILL:
			fmt.Println("SIGKILL")
		}
		fmt.Println("3333", c)
	}
	time.Sleep(time.Second * 10)
}
