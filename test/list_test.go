package test

import (
	"errors"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	aa()

}

func aa() (er error) {
	defer func() {
		fmt.Println(er)
	}()
	return errors.New("111")
}
