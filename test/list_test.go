package test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestList(t *testing.T) {

	// data := []uint8{8, 1, 16, 1, 34, 2, 2, 1}
	fmt.Println(strconv.FormatFloat(1, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(1.000, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(1.0001, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(1.000002, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(1.0333333333333333333333333333, 'f', -1, 64))

}

func aa() (er error) {
	defer func() {
		fmt.Println(er)
	}()
	return errors.New("111")
}
