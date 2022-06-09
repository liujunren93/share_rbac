package t1

import (
	"fmt"
	"strings"
	"testing"

	"github.com/casbin/casbin/v2"
)

var str = `admin, domain1, /rbac/path/:id, GET
admin, domain1, data1, write
admin, domain2, data2, read
admin, domain2, data2, write`

var prolicy [][]string

func init() {
	str = strings.ReplaceAll(str, " ", "")
	data := strings.Split(str, "\n")
	for _, v := range data {
		da := strings.Split(v, ",")

		prolicy = append(prolicy, da)

	}
	fmt.Println(prolicy)
}
func TestList(t *testing.T) {
	e, err := casbin.NewEnforcer("./model.conf")
	if err != nil {
		panic(err)
	}

	e.AddPolicies(prolicy)
	e.AddGroupingPolicies([]string{"alice", "admin", "domain1"})

	ok, err := e.Enforce("alice", "domain1", "/rbac/path/1", "GET")
	fmt.Println(ok, err)

}
