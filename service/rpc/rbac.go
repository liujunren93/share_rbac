package rpc

import (
	"context"

	"github.com/liujunren93/share_rbac/intenal/dao"
	pb "github.com/liujunren93/share_rbac/rbac_pb"
	"github.com/liujunren93/share_utils/common/mq"
	"github.com/liujunren93/share_utils/netHelper"
)

/**
* @Author: liujunren
* @Date: 2022/3/7 13:36
 */

type Rbac struct {
	mq mq.Mqer
}

func NewRbacServer(mq mq.Mqer) *Rbac {
	return &Rbac{mq: mq}
}

func (r *Rbac) AdminMenuTree(ctx context.Context, req *pb.AdminMenuTreeReq) (*pb.DefaultRes, error) {
	tree := dao.Admin{}.MenuTree(req.RoleIDs)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseString(&res, nil, tree)
}

func (r *Rbac) RolePermission(_ context.Context, req *pb.RolePermissionReq) (*pb.DefaultRes, error) {
	rp := dao.Path{}.GetPathActionsByRoles(1, req.RoleIDs)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseString(&res, nil, rp)
}

func (r *Rbac) AdminInfo(_ context.Context, req *pb.DefaultPkReq) (*pb.LoginRes, error) {
	info := dao.Admin{}.AdminInfo(req)
	var res pb.LoginRes
	return &res, netHelper.RpcResponse(&res, nil, info)
}

func (r *Rbac) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	data, err := dao.Admin{}.Login(req)
	var res pb.LoginRes
	return &res, netHelper.RpcResponse(&res, err, data)
}

func (r *Rbac) AccountEdit(ctx context.Context, req *pb.AccountEditReq) (*pb.DefaultRes, error) {

	err := dao.Admin{}.Update(&pb.AdminUpdateReq{
		ID:       req.UID,
		Name:     req.Name,
		Password: req.Password,
		DomainID: req.DomainID,
	})
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r *Rbac) GetDomainPolicy(ctx context.Context, req *pb.GetDomainPolicyReq) (*pb.DefaultRes, error) {
	userList := dao.Role{}.GetDomainPolicy(req.DomainID)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseString(&res, nil, userList)
}

func (r *Rbac) AccountInfo(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	ra := dao.Admin{}.Info(req, "name")
	var res pb.DefaultRes

	return &res, netHelper.RpcResponseString(&res, nil, ra)
}
