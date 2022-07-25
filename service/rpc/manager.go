package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/liujunren93/share_rbac/intenal/dao"
	"github.com/liujunren93/share_rbac/log"
	pb "github.com/liujunren93/share_rbac/rbac_pb"
	"github.com/liujunren93/share_utils/netHelper"
)

const (
	REDISKEY_MQ_DOMAIN_PERSMISSION = "rbac_mq_permission_change_domain"
)

func (r *Rbac) Publish(ctx context.Context, key string, val interface{}) error {
	if val == nil {
		return nil
	}
	ctx, cf := context.WithTimeout(ctx, time.Second*3)
	defer cf()
	err := r.mq.Publish(ctx, key, val)
	if err != nil {
		log.Logger.Error(err, key, val)
		return err
	}
	return nil
}

func (r *Rbac) MAdminRoleList(_ context.Context, req *pb.AdminRoleListReq) (*pb.DefaultRes, error) {
	list := dao.Admin{}.RoleList(req)
	var res pb.DefaultRes

	return &res, netHelper.RpcResponseString(&res, nil, list)
}

func (r *Rbac) MAdminRoleSet(ctx context.Context, req *pb.AdminRoleSetReq) (*pb.DefaultRes, error) {
	err := dao.Admin{}.SetRole(req)
	var res pb.DefaultRes
	r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.DomainID)
	return &res, netHelper.RpcResponse(&res, err, nil)
}

//admin
func (r *Rbac) MAdminList(ctx context.Context, req *pb.AdminListReq) (*pb.DefaultRes, error) {
	list := dao.Admin{}.List(req)
	var res pb.DefaultRes

	return &res, netHelper.RpcResponseString(&res, nil, list)
}

func (r *Rbac) MAdminInfo(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	ra := dao.Admin{}.Info(req)
	var res pb.DefaultRes

	return &res, netHelper.RpcResponseString(&res, nil, ra)
}

func (r *Rbac) MAdminCreate(ctx context.Context, req *pb.AdminCreateReq) (*pb.DefaultRes, error) {
	er := dao.Admin{}.Create(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, er, nil)
}

func (r *Rbac) MAdminUpdate(ctx context.Context, req *pb.AdminUpdateReq) (*pb.DefaultRes, error) {
	err := dao.Admin{}.Update(req)
	r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.DomainID)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r *Rbac) MAdminDel(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	err := dao.Admin{}.Del(req)
	var res pb.DefaultRes
	r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.DomainID)
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r *Rbac) MRoleList(ctx context.Context, req *pb.RoleListReq) (*pb.DefaultRes, error) {
	list := dao.Role{}.List(req)
	var res pb.DefaultRes

	return &res, netHelper.RpcResponseString(&res, nil, list)
}

func (r *Rbac) MRoleInfo(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	info := dao.Role{}.Info(req)
	var res pb.DefaultRes

	return &res, netHelper.RpcResponseString(&res, nil, info)
}

func (r *Rbac) MRoleCreate(ctx context.Context, req *pb.RoleCreateReq) (*pb.DefaultRes, error) {
	err := dao.Role{}.Create(req)
	var res pb.DefaultRes
	r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.DomainID)
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r *Rbac) MRoleUpdate(ctx context.Context, req *pb.RoleUpdateReq) (*pb.DefaultRes, error) {
	var res pb.DefaultRes
	err := dao.Role{}.Update(req)
	if err != nil {
		return &res, netHelper.RpcResponse(&res, err, nil)
	}

	err1 := r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.DomainID)
	if err1 != nil {
		fmt.Println(err1)
	}
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r *Rbac) MRoleDel(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	err := dao.Role{}.Del(req.Pk.(*pb.DefaultPkReq_ID).ID, req.DomainID)
	var res pb.DefaultRes
	r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.DomainID)
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r *Rbac) MRolePermissionList(ctx context.Context, req *pb.RolePermissionListReq) (*pb.DefaultRes, error) {
	list := dao.Role{}.RolePermissionList(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseString(&res, nil, list)
}

func (r *Rbac) MRolePermissionSet(ctx context.Context, req *pb.RolePermissionSetReq) (*pb.DefaultRes, error) {
	err := dao.Role{}.RolePermissionSet(req)
	var res pb.DefaultRes
	r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.DomainID)
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r *Rbac) MPermissionList(ctx context.Context, req *pb.PermissionListReq) (*pb.DefaultRes, error) {
	list := dao.Permission{}.List(req)
	var res pb.DefaultRes

	return &res, netHelper.RpcResponseString(&res, nil, list)
}

func (r *Rbac) MPermissionInfo(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	rp := dao.Permission{}.Info(req)
	var res pb.DefaultRes

	return &res, netHelper.RpcResponseString(&res, nil, rp)

}

func (r *Rbac) MPermissionCreate(ctx context.Context, req *pb.PermissionCreateReq) (*pb.DefaultRes, error) {
	id, err := dao.Permission{}.Create(req)
	var res pb.DefaultRes

	return &res, netHelper.RpcResponseString(&res, err, map[string]interface{}{"pk": id})
}

func (r *Rbac) MPermissionUpdate(ctx context.Context, req *pb.PermissionUpdateReq) (*pb.DefaultRes, error) {
	err := dao.Permission{}.Update(req)
	r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.DomainID)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseString(&res, err, nil)
}

func (r *Rbac) MPermissionDel(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	err := dao.Permission{}.Del(req)
	var res pb.DefaultRes
	r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.DomainID)
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r *Rbac) MPermissionPathList(ctx context.Context, req *pb.PermissionPathListReq) (*pb.DefaultRes, error) {
	list := dao.Permission{}.PathList(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseString(&res, nil, list)
}

func (r *Rbac) MPermissionPathSet(ctx context.Context, req *pb.PermissionPathSetReq) (*pb.DefaultRes, error) {
	err := dao.Permission{}.PathSet(req)
	var res pb.DefaultRes
	r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.DomainID)
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r *Rbac) MPathList(ctx context.Context, req *pb.PathListReq) (*pb.DefaultRes, error) {
	list := dao.Path{}.List(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseString(&res, nil, list)
}

func (r *Rbac) MPathInfo(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	rp := dao.Path{}.Info(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseString(&res, nil, rp)
}

func (r *Rbac) MPathCreate(ctx context.Context, req *pb.PathCreateReq) (*pb.DefaultRes, error) {
	id, err := dao.Path{}.Create(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseString(&res, err, map[string]interface{}{"pk": id})
}

func (r *Rbac) MPathUpdate(ctx context.Context, req *pb.PathUpdateReq) (*pb.DefaultRes, error) {
	err := dao.Path{}.Update(req)
	var res pb.DefaultRes
	r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.DomainID)
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r *Rbac) MPathDel(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	err := dao.Path{}.Del(req)
	var res pb.DefaultRes
	r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.DomainID)
	return &res, netHelper.RpcResponse(&res, err, nil)
}
func (r *Rbac) MDomainList(ctx context.Context, req *pb.DomainListReq) (*pb.DefaultRes, error) {
	list := dao.Domain{}.List(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseString(&res, nil, list)
}

func (r *Rbac) MDomainCreate(ctx context.Context, req *pb.DomainCreateReq) (*pb.DefaultRes, error) {
	err := dao.Domain{}.Create(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseString(&res, err, nil)
}

func (r *Rbac) MDomainUpdate(ctx context.Context, req *pb.DomainUpdateReq) (*pb.DefaultRes, error) {
	err := dao.Domain{}.Update(req)
	var res pb.DefaultRes
	r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.ID)
	return &res, netHelper.RpcResponse(&res, err, nil)
}

func (r *Rbac) MDomainDel(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	err := dao.Domain{}.Del(req)
	var res pb.DefaultRes
	r.Publish(ctx, REDISKEY_MQ_DOMAIN_PERSMISSION, req.Pk.(*pb.DefaultPkReq_ID).ID)
	err1 := netHelper.RpcResponse(&res, err, nil)

	return &res, err1
}

func (r *Rbac) MDomainInfo(ctx context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	info := dao.Domain{}.Info(req)
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseString(&res, nil, info)
}
func (r *Rbac) MAdminPermission(_ context.Context, req *pb.DefaultPkReq) (*pb.DefaultRes, error) {
	rp := dao.Path{}.GetPathWithPermissionByUid(1, uint(req.Pk.(*pb.DefaultPkReq_ID).ID))
	var res pb.DefaultRes
	return &res, netHelper.RpcResponseString(&res, nil, rp)
}
