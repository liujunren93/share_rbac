package rpc

import (
	"context"
	"fmt"
	"github.com/liujunren93/share_rbac/intenal/dao"
	"github.com/liujunren93/share_rbac/pb"
	"github.com/liujunren93/share_utils/netHelper"
)

/**
* @Author: liujunren
* @Date: 2022/3/7 13:36
 */
type Rbac struct {
}

//admin
func (r Rbac) AdminList(ctx context.Context, req *pb.AdminListReq) (*pb.DefaultRes, error) {
	list := dao.Admin{}.List(req)
	var res pb.DefaultRes

	return &res, netHelper.RpcResponseMarshal(&res, nil, list)
}

func (r Rbac) AdminInfo(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	//TODO implement me
	panic("implement me")
}

func (r Rbac) AdminCreate(ctx context.Context, req *pb.AdminCreateReq) (*pb.DefaultRes, error) {
	er := dao.Admin{}.Create(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, er, nil)
}

func (r Rbac) AdminUpdate(ctx context.Context, req *pb.AdminUpdateReq) (*pb.DefaultRes, error) {
	err := dao.Admin{}.Update(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r Rbac) CheckPwd(ctx context.Context, req *pb.CheckPwdReq) (*pb.DefaultRes, error) {
	err := dao.Admin{}.CheckPwd(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r Rbac) AdminDel(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	err := dao.Admin{}.Del(req.Pk.(*pb.DefaultPkReq_ID).ID)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r Rbac) RoleList(ctx context.Context, req *pb.RoleListReq) (*pb.DefaultRes, error) {
	list := dao.Role{}.List(req)
	var res pb.DefaultRes

	return &res, netHelper.RpcResponseMarshal(&res, nil, list)
}

func (r Rbac) RoleInfo(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	//TODO implement me
	panic("implement me")
}

func (r Rbac) RoleCreate(ctx context.Context, req *pb.RoleCreateReq) (*pb.DefaultRes, error) {
	err := dao.Role{}.Create(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r Rbac) RoleUpdate(ctx context.Context, req *pb.RoleUpdateReq) (*pb.DefaultRes, error) {
	err := dao.Role{}.Update(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r Rbac) RoleDel(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	err := dao.Role{}.Del(req.Pk.(*pb.DefaultPkReq_ID).ID)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r Rbac) RolePermissionList(ctx context.Context, req *pb.RolePermissionListReq) (*pb.DefaultRes, error) {
	list := dao.Role{}.RolePermissionList(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseMarshal(&res, nil, list)
}

func (r Rbac) RolePermissionSet(ctx context.Context, req *pb.RolePermissionSetReq) (*pb.DefaultRes, error) {
	err := dao.Role{}.RolePermissionSet(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r Rbac) PermissionList(ctx context.Context, req *pb.PermissionListReq) (*pb.DefaultRes, error) {
	list := dao.Permission{}.List(req)
	var res pb.DefaultRes

	return &res, netHelper.RpcResponseMarshal(&res, nil, list)
}

func (r Rbac) PermissionInfo(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	//TODO implement me
	panic("implement me")
}

func (r Rbac) PermissionCreate(ctx context.Context, req *pb.PermissionCreateReq) (*pb.DefaultRes, error) {
	err := dao.Permission{}.Create(req)
	var res pb.DefaultRes

	return &res, netHelper.RpcResponseMarshal(&res, err, nil)
}

func (r Rbac) PermissionUpdate(ctx context.Context, req *pb.PermissionUpdateReq) (*pb.DefaultRes, error) {
	err := dao.Permission{}.Update(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseMarshal(&res, err, nil)
}

func (r Rbac) PermissionDel(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	//err := dao.Permission{}.Del(req.Pk.(*pb.DefaultPkReq_ID).ID)
	//var res pb.DefaultRes
	//return &res, netHelper.RpcResponse(&res, err, nil)
	panic(11)
}

func (r Rbac) PermissionPathList(ctx context.Context, req *pb.PermissionPathListReq) (*pb.DefaultRes, error) {
	list := dao.Permission{}.PathList(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, nil, list)
}

func (r Rbac) PermissionPathSet(ctx context.Context, req *pb.PermissionPathSetReq) (*pb.DefaultRes, error) {
	err := dao.Permission{}.PathSet(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r Rbac) PathList(ctx context.Context, req *pb.PathListReq) (*pb.DefaultRes, error) {
	list := dao.Path{}.List(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseMarshal(&res, nil, list)
}

func (r Rbac) PathInfo(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	//TODO implement me
	panic("implement me")
}

func (r Rbac) PathCreate(ctx context.Context, req *pb.PathCreateReq) (*pb.DefaultRes, error) {
	err := dao.Path{}.Create(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r Rbac) PathUpdate(ctx context.Context, req *pb.PathUpdateReq) (*pb.DefaultRes, error) {
	err := dao.Path{}.Update(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r Rbac) PathDel(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	err := dao.Path{}.Del(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}
func (r Rbac) DomainList(ctx context.Context, req *pb.DomainListReq) (*pb.DefaultRes, error) {
	list := dao.Domain{}.List(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseMarshal(&res, nil, list)
}

func (r Rbac) DomainCreate(ctx context.Context, req *pb.DomainCreateReq) (*pb.DefaultRes, error) {
	list := dao.Domain{}.Create(req)
	var res pb.DefaultRes
	fmt.Println(list)
	return &res, netHelper.RpcResponseMarshal(&res, nil, list)
}

func (r Rbac) DomainUpdate(ctx context.Context, req *pb.DomainUpdateReq) (*pb.DefaultRes, error) {
	err := dao.Domain{}.Update(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r Rbac) DomainDel(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	err := dao.Domain{}.Del(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r Rbac) DomainInfo(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	info := dao.Domain{}.Info(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, nil, info)
}
func (r Rbac) AdminMenuTree(ctx context.Context, req *pb.AdminMenuTreeReq) (*pb.DefaultRes, error) {
	tree := dao.Admin{}.MenuTree(uint(req.UID))
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, nil, tree)
}
