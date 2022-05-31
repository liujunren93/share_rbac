package t1

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	var a = make(map[interface{}]interface{})
	// fmt.Println(aa == bb)
	a[aa] = aa
	a[bb] = bb
	for _, v := range a {
		v.(ftyp)("")
	}
}

type ftyp func(string) string

var aa ftyp = func(string) string {
	fmt.Println(1)
	return ""
}
var bb ftyp = func(string) string {
	fmt.Println(2)
	return ""
}
