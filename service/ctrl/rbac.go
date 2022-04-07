package ctrl

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_rbac/pb"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/netHelper"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 11:44
 */
type Rbac struct {
}

func (*Rbac) DomainList(ctx *gin.Context) {
	var req pb.DomainListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := client.DomainList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}
func (*Rbac) DomainInfo(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	info, err := client.DomainInfo(ctx, &req)
	netHelper.Response(ctx, info, err, nil)
}
func (*Rbac) DomainDel(ctx *gin.Context) {

}
func (*Rbac) DomainCreate(ctx *gin.Context) {
	var req pb.DomainCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	create, err := client.DomainCreate(ctx, &req)
	netHelper.Response(ctx, create, err, nil)
}
func (*Rbac) DomainUpdate(ctx *gin.Context) {}

func (*Rbac) AdminList(ctx *gin.Context) {
	var req pb.AdminListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := client.AdminList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)

}

func (*Rbac) AdminCreate(ctx *gin.Context) {
	var req pb.AdminCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	create, err := client.AdminCreate(ctx, &req)
	netHelper.Response(ctx, create, err, nil)
}

func (*Rbac) AdminUpdate(ctx *gin.Context) {
	var req pb.AdminUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	update, err := client.AdminUpdate(ctx, &req)
	netHelper.Response(ctx, update, err, nil)

}

func (*Rbac) AdminDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	update, err := client.AdminDel(ctx, &req)
	netHelper.Response(ctx, update, err, nil)
}

func (*Rbac) PermissionList(ctx *gin.Context) {
	var req pb.PermissionListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	list, err := client.PermissionList(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) PermissionCreate(ctx *gin.Context) {
	var req pb.PermissionCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	list, err := client.PermissionCreate(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) PermissionUpdate(ctx *gin.Context) {
	var req pb.PermissionUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	list, err := client.PermissionUpdate(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) PermissionDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	list, err := client.PermissionDel(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) PermissionPathList(ctx *gin.Context) {
	var req pb.PermissionPathListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := client.PermissionPathList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) PermissionPathSet(ctx *gin.Context) {
	var req pb.PermissionPathSetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := client.PermissionPathSet(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) RoleList(ctx *gin.Context) {
	var req pb.RoleListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := client.RoleList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) RoleCreate(ctx *gin.Context) {
	var req pb.RoleCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := client.RoleCreate(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) RoleUpdate(ctx *gin.Context) {
	var req pb.RoleUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	list, err := client.RoleUpdate(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) RoleDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}

	list, err := client.RoleDel(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) RolePermissionList(ctx *gin.Context) {
	var req pb.RolePermissionListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := client.RolePermissionList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) RolePermissionSet(ctx *gin.Context) {
	var req pb.RolePermissionSetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := client.RolePermissionSet(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) PathList(ctx *gin.Context) {
	var req pb.PathListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := client.PathList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) PathCreate(ctx *gin.Context) {
	var req pb.PathCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := client.PathCreate(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) PathUpdate(ctx *gin.Context) {
	var req pb.PathUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	list, err := client.PathUpdate(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) PathDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}

	list, err := client.PathDel(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) UserMenu(ctx *gin.Context) {
	var req *pb.AdminMenuTreeReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	tree, err := client.AdminMenuTree(ctx, req)
	netHelper.Response(ctx, tree, err, nil)
}
