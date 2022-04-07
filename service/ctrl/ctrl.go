package ctrl

import (
	"github.com/liujunren93/share_rbac/pb"
	"google.golang.org/grpc"
)

/**
* @Author: liujunren
* @Date: 2022/3/7 16:03
 */

var client pb.RbacClient

func init()  {
	//newClient := cli.NewClient()
	//conn, err := newClient.GetGrpcClient("admin")
	//if err != nil {
	//
	//}
	dial, err := grpc.Dial("127.0.0.1:29091",grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client = pb.NewRbacClient(dial)
}