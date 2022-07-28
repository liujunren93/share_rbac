package test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/liujunren93/share_rbac/rbac_pb"
	"google.golang.org/protobuf/encoding/prototext"
)

func TestList(t *testing.T) {
	session := rbac_pb.Session{
		UID:      1,
		DomainID: 3,
		PL:       4,
		RoleIDs:  []int64{1, 2},
	}

	var session2 rbac_pb.Session

	err := prototext.Unmarshal([]byte(session.String()), &session2)
	fmt.Println(&session2, err)

	// data := []uint8{8, 1, 16, 1, 34, 2, 2, 1}
	// fmt.Println(string(data))

}

func aa() (er error) {
	defer func() {
		fmt.Println(er)
	}()
	return errors.New("111")
}
