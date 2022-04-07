package service

import (
	"github.com/liujunren93/share/server"
	"github.com/liujunren93/share_rbac/pb"
	"github.com/liujunren93/share_rbac/service/rpc"
)

/**
* @Author: liujunren
* @Date: 2022/3/7 17:06
 */

func NewGrpcService()  {
	grpcServer := server.NewGrpcServer(server.WithAddress("0.0.0.0:29091"))
	pb.RegisterRbacServer(grpcServer.Server(),rpc.Rbac{})
	grpcServer.Run()
}